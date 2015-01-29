package preparer

import (
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/square/p2/pkg/hooks"
	"github.com/square/p2/pkg/kp"
	"github.com/square/p2/pkg/logging"
	"github.com/square/p2/pkg/pods"
)

type Pod interface {
	hooks.Pod
	Launch(*pods.PodManifest) (bool, error)
	Install(*pods.PodManifest) error
	Halt() (bool, error)
}

type Hooks interface {
	RunHookType(hookType hooks.HookType, pod hooks.Pod, manifest *pods.PodManifest) error
}

type Store interface {
	Pod(string) (*pods.PodManifest, time.Duration, error)
	SetPod(string, pods.PodManifest) (time.Duration, error)
	RegisterService(pods.PodManifest) error
	WatchPods(string, <-chan struct{}, chan<- error, chan<- kp.ManifestResult)
}

type Preparer struct {
	node         string
	store        Store
	hooks        Hooks
	hookListener HookListener
	Logger       logging.Logger
}

func New(nodeName string, consulAddress string, hooksDirectory string, logger logging.Logger) (*Preparer, error) {
	store := kp.NewStore(kp.Options{
		Address: consulAddress,
	})

	listener := HookListener{
		Intent:         store,
		HookPrefix:     kp.HOOK_TREE,
		DestinationDir: path.Join(pods.DEFAULT_PATH, "hooks"),
		ExecDir:        hooksDirectory,
		Logger:         logger,
	}

	return &Preparer{
		node:         nodeName,
		store:        store,
		hooks:        hooks.Hooks(hooksDirectory, &logger),
		hookListener: listener,
		Logger:       logger,
	}, nil
}

func (p *Preparer) WatchForHooks(quit chan struct{}) {
	hookErrCh := make(chan error)
	hookQuitCh := make(chan struct{})

	go p.hookListener.Sync(hookQuitCh, hookErrCh)
	for {
		select {
		case <-quit:
			hookQuitCh <- struct{}{}
			return
		case err := <-hookErrCh:
			p.Logger.WithField("err", err).Errorln("Error updating hooks")
		}
	}
}

func (p *Preparer) WatchForPodManifestsForNode(quitAndAck chan struct{}) {
	pods.Log = p.Logger
	path := kp.IntentPath(p.node)

	// This allows us to signal the goroutine watching consul to quit
	watcherQuit := make(<-chan struct{})
	errChan := make(chan error)
	podChan := make(chan kp.ManifestResult)

	go p.store.WatchPods(path, watcherQuit, errChan, podChan)

	// we will have one long running goroutine for each app installed on this
	// host. We keep a map of podId => podChan so we can send the new manifests
	// that come in to the appropriate goroutine
	podChanMap := make(map[string]chan pods.PodManifest)
	quitChanMap := make(map[string]chan struct{})

	for {
		select {
		case err := <-errChan:
			p.Logger.WithFields(logrus.Fields{
				"inner_err": err,
			}).Errorln("there was an error reading the manifest")
		case result := <-podChan:
			podId := result.Manifest.ID()
			if podChanMap[podId] == nil {
				// No goroutine is servicing this app currently, let's start one
				podChanMap[podId] = make(chan pods.PodManifest)
				quitChanMap[podId] = make(chan struct{})
				go p.handlePods(podChanMap[podId], quitChanMap[podId])
			}
			podChanMap[podId] <- result.Manifest
		case <-quitAndAck:
			for podToQuit, quitCh := range quitChanMap {
				p.Logger.WithField("pod", podToQuit).Infoln("Quitting...")
				quitCh <- struct{}{}
			}
			p.Logger.NoFields().Infoln("Done, acknowledging quit")
			quitAndAck <- struct{}{} // acknowledge quit
			return
		}

	}
}

// no return value, no output channels. This should do everything it needs to do
// without outside intervention (other than being signalled to quit)
func (p *Preparer) handlePods(podChan <-chan pods.PodManifest, quit <-chan struct{}) {
	// install new launchables
	var manifestToLaunch pods.PodManifest

	// used to track if we have work to do (i.e. pod manifest came through channel
	// and we have yet to operate on it)
	working := false
	var manifestLogger logging.Logger
	for {
		select {
		case <-quit:
			return
		case manifestToLaunch = <-podChan:
			sha, err := manifestToLaunch.SHA()
			manifestLogger = p.Logger.SubLogger(logrus.Fields{
				"manifest": manifestToLaunch.ID(),
				"sha":      sha,
				"sha_err":  err,
			})
			manifestLogger.NoFields().Infoln("New manifest received")
			working = true
		case <-time.After(1 * time.Second):
			if working {
				pod := pods.PodFromManifestId(manifestToLaunch.ID())

				// HACK ZONE. When we have better authz, rewrite.
				// Still need to ensure that preparer, intent both launch correctly
				// as root
				if pod.Id == "intent" || pod.Id == "p2-preparer" {
					pod.RunAs = "root"
				}

				ok := p.installAndLaunchPod(&manifestToLaunch, pod, manifestLogger)
				if ok {
					manifestToLaunch = pods.PodManifest{}
					working = false
				}
			}
		}
	}
}

func (p *Preparer) installAndLaunchPod(newManifest *pods.PodManifest, pod Pod, logger logging.Logger) bool {
	// do not remove the logger argument, it's not the same as p.Logger

	// get currently running pod to compare with the new pod
	realityPath := kp.RealityPath(p.node, newManifest.ID())
	currentManifest, _, err := p.store.Pod(realityPath)
	currentSHA := ""
	if currentManifest != nil {
		currentSHA, _ = currentManifest.SHA()
	}
	newSHA, _ := newManifest.SHA()

	// if new or the manifest is different, launch
	newOrDifferent := err == pods.NoCurrentManifest || currentSHA != newSHA
	if newOrDifferent {
		logger.WithFields(logrus.Fields{
			"old_sha":  currentSHA,
			"sha":      newSHA,
			"manifest": newManifest.ID(),
		}).Infoln("SHA is new or different from old, will update")
	}

	// if the old manifest is corrupted somehow, re-launch since we don't know if this is an update.
	problemReadingCurrentManifest := (err != nil && err != pods.NoCurrentManifest)
	if problemReadingCurrentManifest {
		logger.WithFields(logrus.Fields{
			"sha":       newSHA,
			"inner_err": err,
		}).Errorln("Current manifest not readable, will relaunch")
	}

	if newOrDifferent || problemReadingCurrentManifest {
		err := p.hooks.RunHookType(hooks.BEFORE_INSTALL, pod, newManifest)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err":   err,
				"hooks": hooks.BEFORE_INSTALL,
			}).Warnln("Could not run hooks")
		}

		err = pod.Install(newManifest)
		if err != nil {
			// install failed, abort and retry
			logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Install failed")
			return false
		}

		err = p.hooks.RunHookType(hooks.AFTER_INSTALL, pod, newManifest)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err":   err,
				"hooks": hooks.AFTER_INSTALL,
			}).Warnln("Could not run hooks")
		}

		err = p.store.RegisterService(*newManifest)
		if err != nil {
			logger.WithField("err", err).Errorln("Service registration failed")
			return false
		}
		ok, err := pod.Launch(newManifest)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Launch failed")
		} else {
			duration, err := p.store.SetPod(realityPath, *newManifest)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"err":      err,
					"duration": duration,
				}).Errorln("Could not set pod in reality store")
			}

			err = p.hooks.RunHookType(hooks.AFTER_LAUNCH, pod, newManifest)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"err":   err,
					"hooks": hooks.AFTER_LAUNCH,
				}).Warnln("Could not run hooks")
			}
		}
		return err == nil && ok
	}

	// TODO: shut down removed launchables between pod versions.
	return true
}