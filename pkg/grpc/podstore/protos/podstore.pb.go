// Code generated by protoc-gen-go.
// source: pkg/grpc/podstore/protos/podstore.proto
// DO NOT EDIT!

/*
Package podstore is a generated protocol buffer package.

It is generated from these files:
	pkg/grpc/podstore/protos/podstore.proto

It has these top-level messages:
	SchedulePodRequest
	SchedulePodResponse
	WatchPodStatusRequest
	PodStatusResponse
	ProcessStatus
	ExitStatus
	UnschedulePodRequest
	UnschedulePodResponse
*/
package podstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SchedulePodRequest struct {
	Manifest string `protobuf:"bytes,1,opt,name=manifest" json:"manifest,omitempty"`
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *SchedulePodRequest) Reset()                    { *m = SchedulePodRequest{} }
func (m *SchedulePodRequest) String() string            { return proto.CompactTextString(m) }
func (*SchedulePodRequest) ProtoMessage()               {}
func (*SchedulePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SchedulePodRequest) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *SchedulePodRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type SchedulePodResponse struct {
	PodUniqueKey string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
}

func (m *SchedulePodResponse) Reset()                    { *m = SchedulePodResponse{} }
func (m *SchedulePodResponse) String() string            { return proto.CompactTextString(m) }
func (*SchedulePodResponse) ProtoMessage()               {}
func (*SchedulePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SchedulePodResponse) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

type WatchPodStatusRequest struct {
	PodUniqueKey    string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
	StatusNamespace string `protobuf:"bytes,3,opt,name=status_namespace,json=statusNamespace" json:"status_namespace,omitempty"`
	WaitForExists   bool   `protobuf:"varint,4,opt,name=wait_for_exists,json=waitForExists" json:"wait_for_exists,omitempty"`
}

func (m *WatchPodStatusRequest) Reset()                    { *m = WatchPodStatusRequest{} }
func (m *WatchPodStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchPodStatusRequest) ProtoMessage()               {}
func (*WatchPodStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WatchPodStatusRequest) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

func (m *WatchPodStatusRequest) GetStatusNamespace() string {
	if m != nil {
		return m.StatusNamespace
	}
	return ""
}

func (m *WatchPodStatusRequest) GetWaitForExists() bool {
	if m != nil {
		return m.WaitForExists
	}
	return false
}

type PodStatusResponse struct {
	Manifest        string           `protobuf:"bytes,1,opt,name=manifest" json:"manifest,omitempty"`
	PodState        string           `protobuf:"bytes,2,opt,name=pod_state,json=podState" json:"pod_state,omitempty"`
	ProcessStatuses []*ProcessStatus `protobuf:"bytes,3,rep,name=process_statuses,json=processStatuses" json:"process_statuses,omitempty"`
}

func (m *PodStatusResponse) Reset()                    { *m = PodStatusResponse{} }
func (m *PodStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*PodStatusResponse) ProtoMessage()               {}
func (*PodStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PodStatusResponse) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *PodStatusResponse) GetPodState() string {
	if m != nil {
		return m.PodState
	}
	return ""
}

func (m *PodStatusResponse) GetProcessStatuses() []*ProcessStatus {
	if m != nil {
		return m.ProcessStatuses
	}
	return nil
}

type ProcessStatus struct {
	LaunchableId string      `protobuf:"bytes,1,opt,name=launchable_id,json=launchableId" json:"launchable_id,omitempty"`
	EntryPoint   string      `protobuf:"bytes,2,opt,name=entry_point,json=entryPoint" json:"entry_point,omitempty"`
	LastExit     *ExitStatus `protobuf:"bytes,3,opt,name=last_exit,json=lastExit" json:"last_exit,omitempty"`
}

func (m *ProcessStatus) Reset()                    { *m = ProcessStatus{} }
func (m *ProcessStatus) String() string            { return proto.CompactTextString(m) }
func (*ProcessStatus) ProtoMessage()               {}
func (*ProcessStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ProcessStatus) GetLaunchableId() string {
	if m != nil {
		return m.LaunchableId
	}
	return ""
}

func (m *ProcessStatus) GetEntryPoint() string {
	if m != nil {
		return m.EntryPoint
	}
	return ""
}

func (m *ProcessStatus) GetLastExit() *ExitStatus {
	if m != nil {
		return m.LastExit
	}
	return nil
}

type ExitStatus struct {
	ExitTime   int64 `protobuf:"varint,1,opt,name=exit_time,json=exitTime" json:"exit_time,omitempty"`
	ExitCode   int64 `protobuf:"varint,2,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	ExitStatus int64 `protobuf:"varint,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *ExitStatus) Reset()                    { *m = ExitStatus{} }
func (m *ExitStatus) String() string            { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()               {}
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExitStatus) GetExitTime() int64 {
	if m != nil {
		return m.ExitTime
	}
	return 0
}

func (m *ExitStatus) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *ExitStatus) GetExitStatus() int64 {
	if m != nil {
		return m.ExitStatus
	}
	return 0
}

type UnschedulePodRequest struct {
	PodUniqueKey string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
}

func (m *UnschedulePodRequest) Reset()                    { *m = UnschedulePodRequest{} }
func (m *UnschedulePodRequest) String() string            { return proto.CompactTextString(m) }
func (*UnschedulePodRequest) ProtoMessage()               {}
func (*UnschedulePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnschedulePodRequest) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

type UnschedulePodResponse struct {
}

func (m *UnschedulePodResponse) Reset()                    { *m = UnschedulePodResponse{} }
func (m *UnschedulePodResponse) String() string            { return proto.CompactTextString(m) }
func (*UnschedulePodResponse) ProtoMessage()               {}
func (*UnschedulePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*SchedulePodRequest)(nil), "podstore.SchedulePodRequest")
	proto.RegisterType((*SchedulePodResponse)(nil), "podstore.SchedulePodResponse")
	proto.RegisterType((*WatchPodStatusRequest)(nil), "podstore.WatchPodStatusRequest")
	proto.RegisterType((*PodStatusResponse)(nil), "podstore.PodStatusResponse")
	proto.RegisterType((*ProcessStatus)(nil), "podstore.ProcessStatus")
	proto.RegisterType((*ExitStatus)(nil), "podstore.ExitStatus")
	proto.RegisterType((*UnschedulePodRequest)(nil), "podstore.UnschedulePodRequest")
	proto.RegisterType((*UnschedulePodResponse)(nil), "podstore.UnschedulePodResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for P2PodStore service

type P2PodStoreClient interface {
	// Schedules a uuid pod on a host
	SchedulePod(ctx context.Context, in *SchedulePodRequest, opts ...grpc.CallOption) (*SchedulePodResponse, error)
	WatchPodStatus(ctx context.Context, in *WatchPodStatusRequest, opts ...grpc.CallOption) (P2PodStore_WatchPodStatusClient, error)
	UnschedulePod(ctx context.Context, in *UnschedulePodRequest, opts ...grpc.CallOption) (*UnschedulePodResponse, error)
}

type p2PodStoreClient struct {
	cc *grpc.ClientConn
}

func NewP2PodStoreClient(cc *grpc.ClientConn) P2PodStoreClient {
	return &p2PodStoreClient{cc}
}

func (c *p2PodStoreClient) SchedulePod(ctx context.Context, in *SchedulePodRequest, opts ...grpc.CallOption) (*SchedulePodResponse, error) {
	out := new(SchedulePodResponse)
	err := grpc.Invoke(ctx, "/podstore.P2PodStore/SchedulePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PodStoreClient) WatchPodStatus(ctx context.Context, in *WatchPodStatusRequest, opts ...grpc.CallOption) (P2PodStore_WatchPodStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_P2PodStore_serviceDesc.Streams[0], c.cc, "/podstore.P2PodStore/WatchPodStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PodStoreWatchPodStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2PodStore_WatchPodStatusClient interface {
	Recv() (*PodStatusResponse, error)
	grpc.ClientStream
}

type p2PodStoreWatchPodStatusClient struct {
	grpc.ClientStream
}

func (x *p2PodStoreWatchPodStatusClient) Recv() (*PodStatusResponse, error) {
	m := new(PodStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *p2PodStoreClient) UnschedulePod(ctx context.Context, in *UnschedulePodRequest, opts ...grpc.CallOption) (*UnschedulePodResponse, error) {
	out := new(UnschedulePodResponse)
	err := grpc.Invoke(ctx, "/podstore.P2PodStore/UnschedulePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for P2PodStore service

type P2PodStoreServer interface {
	// Schedules a uuid pod on a host
	SchedulePod(context.Context, *SchedulePodRequest) (*SchedulePodResponse, error)
	WatchPodStatus(*WatchPodStatusRequest, P2PodStore_WatchPodStatusServer) error
	UnschedulePod(context.Context, *UnschedulePodRequest) (*UnschedulePodResponse, error)
}

func RegisterP2PodStoreServer(s *grpc.Server, srv P2PodStoreServer) {
	s.RegisterService(&_P2PodStore_serviceDesc, srv)
}

func _P2PodStore_SchedulePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PodStoreServer).SchedulePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podstore.P2PodStore/SchedulePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PodStoreServer).SchedulePod(ctx, req.(*SchedulePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PodStore_WatchPodStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPodStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2PodStoreServer).WatchPodStatus(m, &p2PodStoreWatchPodStatusServer{stream})
}

type P2PodStore_WatchPodStatusServer interface {
	Send(*PodStatusResponse) error
	grpc.ServerStream
}

type p2PodStoreWatchPodStatusServer struct {
	grpc.ServerStream
}

func (x *p2PodStoreWatchPodStatusServer) Send(m *PodStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _P2PodStore_UnschedulePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnschedulePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PodStoreServer).UnschedulePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podstore.P2PodStore/UnschedulePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PodStoreServer).UnschedulePod(ctx, req.(*UnschedulePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _P2PodStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "podstore.P2PodStore",
	HandlerType: (*P2PodStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SchedulePod",
			Handler:    _P2PodStore_SchedulePod_Handler,
		},
		{
			MethodName: "UnschedulePod",
			Handler:    _P2PodStore_UnschedulePod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPodStatus",
			Handler:       _P2PodStore_WatchPodStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/podstore/protos/podstore.proto",
}

func init() { proto.RegisterFile("pkg/grpc/podstore/protos/podstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x6b, 0x84, 0x9c, 0x09, 0x69, 0xca, 0xd2, 0xaa, 0x56, 0x0a, 0x24, 0x32, 0x08, 0xc2,
	0xa5, 0x81, 0x70, 0x84, 0x13, 0xa8, 0x48, 0x08, 0xa8, 0x2c, 0x97, 0x8a, 0xa3, 0xb5, 0xf5, 0x4e,
	0x13, 0xab, 0xb1, 0x77, 0xeb, 0x5d, 0x8b, 0xe6, 0xca, 0x89, 0x23, 0x3f, 0x19, 0xed, 0xae, 0x3f,
	0x92, 0x7e, 0xa8, 0x3d, 0xce, 0x9b, 0x8f, 0x7d, 0xf3, 0xe6, 0xd9, 0xf0, 0x5a, 0x9c, 0xcf, 0x26,
	0xb3, 0x42, 0x24, 0x13, 0xc1, 0x99, 0x54, 0xbc, 0xc0, 0x89, 0x28, 0xb8, 0xe2, 0xb2, 0x89, 0x0f,
	0x4c, 0x4c, 0xbc, 0x3a, 0x0e, 0x7e, 0x00, 0x39, 0x4e, 0xe6, 0xc8, 0xca, 0x05, 0x86, 0x9c, 0x45,
	0x78, 0x51, 0xa2, 0x54, 0x64, 0x00, 0x5e, 0x46, 0xf3, 0xf4, 0x0c, 0xa5, 0xf2, 0x9d, 0x91, 0x33,
	0xee, 0x44, 0x4d, 0x4c, 0xf6, 0xa1, 0x93, 0x73, 0x86, 0x71, 0x4e, 0x33, 0xf4, 0x37, 0x6d, 0x52,
	0x03, 0x47, 0x34, 0xc3, 0xe0, 0x03, 0x3c, 0x59, 0x1b, 0x27, 0x05, 0xcf, 0x25, 0x92, 0x97, 0xb0,
	0x25, 0x38, 0x8b, 0xcb, 0x3c, 0xbd, 0x28, 0x31, 0x3e, 0xc7, 0x65, 0x35, 0xf5, 0x91, 0xe0, 0xec,
	0xc4, 0x80, 0xdf, 0x70, 0x19, 0xfc, 0x73, 0x60, 0xf7, 0x17, 0x55, 0xc9, 0x3c, 0xe4, 0xec, 0x58,
	0x51, 0x55, 0xca, 0x9a, 0xcf, 0xbd, 0xfa, 0xc9, 0x1b, 0xd8, 0x96, 0xa6, 0xcd, 0x70, 0x93, 0x82,
	0x26, 0xe8, 0xbb, 0xa6, 0xae, 0x6f, 0xf1, 0xa3, 0x1a, 0x26, 0xaf, 0xa0, 0xff, 0x9b, 0xa6, 0x2a,
	0x3e, 0xe3, 0x45, 0x8c, 0x97, 0xa9, 0x54, 0xd2, 0x7f, 0x30, 0x72, 0xc6, 0x5e, 0xd4, 0xd3, 0xf0,
	0x17, 0x5e, 0x1c, 0x1a, 0x50, 0x53, 0x7a, 0xbc, 0xc2, 0xa6, 0x5a, 0xe7, 0x0e, 0x79, 0x34, 0x55,
	0xfd, 0x60, 0x23, 0x8f, 0xb0, 0x13, 0x90, 0x7c, 0x82, 0x6d, 0x51, 0xf0, 0x04, 0xa5, 0x8c, 0x2d,
	0x23, 0x94, 0xbe, 0x3b, 0x72, 0xc7, 0xdd, 0xe9, 0xde, 0x41, 0x73, 0xa2, 0xd0, 0x56, 0x54, 0x6f,
	0xf6, 0xc5, 0x6a, 0x88, 0x32, 0xf8, 0xeb, 0x40, 0x6f, 0xad, 0x84, 0xbc, 0x80, 0xde, 0x82, 0x96,
	0x79, 0x32, 0xa7, 0xa7, 0x0b, 0x8c, 0x53, 0x56, 0x8b, 0xd3, 0x82, 0x5f, 0x19, 0x19, 0x42, 0x17,
	0x73, 0x55, 0x2c, 0x63, 0xc1, 0xd3, 0x5c, 0x55, 0xcc, 0xc0, 0x40, 0xa1, 0x46, 0xc8, 0x3b, 0xe8,
	0x2c, 0xa8, 0x54, 0x5a, 0x0e, 0x65, 0x64, 0xeb, 0x4e, 0x77, 0x5a, 0x52, 0x87, 0x97, 0xa9, 0xaa,
	0x18, 0x79, 0xba, 0x4c, 0xc7, 0xc1, 0x0c, 0xa0, 0xc5, 0xf5, 0xe6, 0xba, 0x37, 0x56, 0x69, 0x86,
	0x86, 0x82, 0x1b, 0x79, 0x1a, 0xf8, 0x99, 0x66, 0xd8, 0x24, 0x13, 0xce, 0xac, 0x2c, 0x55, 0xf2,
	0x33, 0x67, 0x68, 0xb8, 0xe9, 0xa4, 0xd5, 0xc4, 0x3c, 0xee, 0x46, 0x80, 0xcd, 0xe8, 0xe0, 0x23,
	0xec, 0x9c, 0xe4, 0xf2, 0xba, 0x4f, 0xef, 0xe7, 0xab, 0x3d, 0xd8, 0xbd, 0xd2, 0x6d, 0xef, 0x38,
	0xfd, 0xb3, 0x09, 0x10, 0x4e, 0xcd, 0x7d, 0x79, 0x81, 0xe4, 0x3b, 0x74, 0x57, 0xcc, 0x4b, 0x9e,
	0xb6, 0xdb, 0x5f, 0xff, 0x44, 0x06, 0xcf, 0x6e, 0xc9, 0xda, 0xd1, 0xc1, 0x06, 0x89, 0x60, 0x6b,
	0xdd, 0xcc, 0x64, 0xd8, 0xb6, 0xdc, 0x68, 0xf3, 0xc1, 0xfe, 0x8a, 0x09, 0xae, 0x9a, 0x2e, 0xd8,
	0x78, 0xeb, 0x90, 0x08, 0x7a, 0x6b, 0x9b, 0x90, 0xe7, 0x6d, 0xc7, 0x4d, 0x02, 0x0d, 0x86, 0xb7,
	0xe6, 0xeb, 0xa9, 0xa7, 0x0f, 0xcd, 0x2f, 0xe1, 0xfd, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x81, 0x01, 0x62, 0x3d, 0x04, 0x00, 0x00,
}
