// Code generated by protoc-gen-go.
// source: seaweed.proto
// DO NOT EDIT!

/*
Package master_pb is a generated protocol buffer package.

It is generated from these files:
	seaweed.proto

It has these top-level messages:
	Heartbeat
	HeartbeatResponse
	VolumeInformationMessage
*/
package master_pb

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

type Heartbeat struct {
	Ip             string                      `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port           uint32                      `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	PublicUrl      string                      `protobuf:"bytes,3,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
	MaxVolumeCount uint32                      `protobuf:"varint,4,opt,name=max_volume_count,json=maxVolumeCount" json:"max_volume_count,omitempty"`
	MaxFileKey     uint64                      `protobuf:"varint,5,opt,name=max_file_key,json=maxFileKey" json:"max_file_key,omitempty"`
	DataCenter     string                      `protobuf:"bytes,6,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Rack           string                      `protobuf:"bytes,7,opt,name=rack" json:"rack,omitempty"`
	AdminPort      uint32                      `protobuf:"varint,8,opt,name=admin_port,json=adminPort" json:"admin_port,omitempty"`
	Volumes        []*VolumeInformationMessage `protobuf:"bytes,9,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Heartbeat) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Heartbeat) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Heartbeat) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func (m *Heartbeat) GetMaxVolumeCount() uint32 {
	if m != nil {
		return m.MaxVolumeCount
	}
	return 0
}

func (m *Heartbeat) GetMaxFileKey() uint64 {
	if m != nil {
		return m.MaxFileKey
	}
	return 0
}

func (m *Heartbeat) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Heartbeat) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Heartbeat) GetAdminPort() uint32 {
	if m != nil {
		return m.AdminPort
	}
	return 0
}

func (m *Heartbeat) GetVolumes() []*VolumeInformationMessage {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type HeartbeatResponse struct {
	VolumeSizeLimit uint64 `protobuf:"varint,1,opt,name=volumeSizeLimit" json:"volumeSizeLimit,omitempty"`
	SecretKey       string `protobuf:"bytes,2,opt,name=secretKey" json:"secretKey,omitempty"`
	Leader          string `protobuf:"bytes,3,opt,name=leader" json:"leader,omitempty"`
}

func (m *HeartbeatResponse) Reset()                    { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()               {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HeartbeatResponse) GetVolumeSizeLimit() uint64 {
	if m != nil {
		return m.VolumeSizeLimit
	}
	return 0
}

func (m *HeartbeatResponse) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *HeartbeatResponse) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

type VolumeInformationMessage struct {
	Id               uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Size             uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Collection       string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
	FileCount        uint64 `protobuf:"varint,4,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	DeleteCount      uint64 `protobuf:"varint,5,opt,name=delete_count,json=deleteCount" json:"delete_count,omitempty"`
	DeletedByteCount uint64 `protobuf:"varint,6,opt,name=deleted_byte_count,json=deletedByteCount" json:"deleted_byte_count,omitempty"`
	ReadOnly         bool   `protobuf:"varint,7,opt,name=read_only,json=readOnly" json:"read_only,omitempty"`
	ReplicaPlacement uint32 `protobuf:"varint,8,opt,name=replica_placement,json=replicaPlacement" json:"replica_placement,omitempty"`
	Version          uint32 `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	Ttl              uint32 `protobuf:"varint,10,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *VolumeInformationMessage) Reset()                    { *m = VolumeInformationMessage{} }
func (m *VolumeInformationMessage) String() string            { return proto.CompactTextString(m) }
func (*VolumeInformationMessage) ProtoMessage()               {}
func (*VolumeInformationMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VolumeInformationMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VolumeInformationMessage) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VolumeInformationMessage) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *VolumeInformationMessage) GetFileCount() uint64 {
	if m != nil {
		return m.FileCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetDeleteCount() uint64 {
	if m != nil {
		return m.DeleteCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetDeletedByteCount() uint64 {
	if m != nil {
		return m.DeletedByteCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *VolumeInformationMessage) GetReplicaPlacement() uint32 {
	if m != nil {
		return m.ReplicaPlacement
	}
	return 0
}

func (m *VolumeInformationMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeInformationMessage) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func init() {
	proto.RegisterType((*Heartbeat)(nil), "master_pb.Heartbeat")
	proto.RegisterType((*HeartbeatResponse)(nil), "master_pb.HeartbeatResponse")
	proto.RegisterType((*VolumeInformationMessage)(nil), "master_pb.VolumeInformationMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Seaweed service

type SeaweedClient interface {
	SendHeartbeat(ctx context.Context, opts ...grpc.CallOption) (Seaweed_SendHeartbeatClient, error)
}

type seaweedClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedClient(cc *grpc.ClientConn) SeaweedClient {
	return &seaweedClient{cc}
}

func (c *seaweedClient) SendHeartbeat(ctx context.Context, opts ...grpc.CallOption) (Seaweed_SendHeartbeatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Seaweed_serviceDesc.Streams[0], c.cc, "/master_pb.Seaweed/SendHeartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedSendHeartbeatClient{stream}
	return x, nil
}

type Seaweed_SendHeartbeatClient interface {
	Send(*Heartbeat) error
	Recv() (*HeartbeatResponse, error)
	grpc.ClientStream
}

type seaweedSendHeartbeatClient struct {
	grpc.ClientStream
}

func (x *seaweedSendHeartbeatClient) Send(m *Heartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedSendHeartbeatClient) Recv() (*HeartbeatResponse, error) {
	m := new(HeartbeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Seaweed service

type SeaweedServer interface {
	SendHeartbeat(Seaweed_SendHeartbeatServer) error
}

func RegisterSeaweedServer(s *grpc.Server, srv SeaweedServer) {
	s.RegisterService(&_Seaweed_serviceDesc, srv)
}

func _Seaweed_SendHeartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedServer).SendHeartbeat(&seaweedSendHeartbeatServer{stream})
}

type Seaweed_SendHeartbeatServer interface {
	Send(*HeartbeatResponse) error
	Recv() (*Heartbeat, error)
	grpc.ServerStream
}

type seaweedSendHeartbeatServer struct {
	grpc.ServerStream
}

func (x *seaweedSendHeartbeatServer) Send(m *HeartbeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedSendHeartbeatServer) Recv() (*Heartbeat, error) {
	m := new(Heartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Seaweed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master_pb.Seaweed",
	HandlerType: (*SeaweedServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendHeartbeat",
			Handler:       _Seaweed_SendHeartbeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "seaweed.proto",
}

func init() { proto.RegisterFile("seaweed.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x16, 0xda, 0xfa, 0x74, 0x1d, 0x9d, 0x85, 0x90, 0x05, 0x03, 0x4a, 0xb9, 0x89,
	0x04, 0xaa, 0xd0, 0xb8, 0xe6, 0x86, 0x49, 0x88, 0x69, 0x20, 0x26, 0x17, 0xb8, 0x8d, 0xdc, 0xe4,
	0x0c, 0x59, 0x73, 0xfe, 0xc8, 0x76, 0x47, 0xb3, 0x77, 0xe2, 0x2d, 0x78, 0x30, 0xe4, 0x93, 0xa6,
	0x9d, 0x10, 0xdc, 0x1d, 0xff, 0xce, 0xe7, 0xf8, 0xe4, 0xfb, 0x6c, 0x98, 0x38, 0x54, 0x3f, 0x11,
	0x8b, 0x45, 0x63, 0x6b, 0x5f, 0x73, 0x56, 0x2a, 0xe7, 0xd1, 0x66, 0xcd, 0x6a, 0xfe, 0x2b, 0x06,
	0xf6, 0x11, 0x95, 0xf5, 0x2b, 0x54, 0x9e, 0x1f, 0x41, 0xac, 0x1b, 0x11, 0xcd, 0xa2, 0x94, 0xc9,
	0x58, 0x37, 0x9c, 0x43, 0xd2, 0xd4, 0xd6, 0x8b, 0x78, 0x16, 0xa5, 0x13, 0x49, 0x35, 0x7f, 0x0a,
	0xd0, 0xac, 0x57, 0x46, 0xe7, 0xd9, 0xda, 0x1a, 0x71, 0x40, 0x5a, 0xd6, 0x91, 0x6f, 0xd6, 0xf0,
	0x14, 0xa6, 0xa5, 0xda, 0x64, 0x37, 0xb5, 0x59, 0x97, 0x98, 0xe5, 0xf5, 0xba, 0xf2, 0x22, 0xa1,
	0xed, 0x47, 0xa5, 0xda, 0x7c, 0x27, 0x7c, 0x16, 0x28, 0x9f, 0xc1, 0x61, 0x50, 0x5e, 0x69, 0x83,
	0xd9, 0x35, 0xb6, 0xe2, 0xfe, 0x2c, 0x4a, 0x13, 0x09, 0xa5, 0xda, 0x7c, 0xd0, 0x06, 0x2f, 0xb0,
	0xe5, 0xcf, 0x61, 0x5c, 0x28, 0xaf, 0xb2, 0x1c, 0x2b, 0x8f, 0x56, 0x0c, 0xe8, 0x2c, 0x08, 0xe8,
	0x8c, 0x48, 0x98, 0xcf, 0xaa, 0xfc, 0x5a, 0x0c, 0xa9, 0x43, 0x75, 0x98, 0x4f, 0x15, 0xa5, 0xae,
	0x32, 0x9a, 0x7c, 0x44, 0x47, 0x33, 0x22, 0x97, 0x61, 0xfc, 0x77, 0x30, 0xec, 0x66, 0x73, 0x82,
	0xcd, 0x0e, 0xd2, 0xf1, 0xe9, 0xcb, 0xc5, 0xce, 0x8d, 0x45, 0x37, 0xde, 0x79, 0x75, 0x55, 0xdb,
	0x52, 0x79, 0x5d, 0x57, 0x9f, 0xd1, 0x39, 0xf5, 0x03, 0x65, 0xbf, 0x67, 0xee, 0xe0, 0x78, 0x67,
	0x97, 0x44, 0xd7, 0xd4, 0x95, 0x43, 0x9e, 0xc2, 0x83, 0xae, 0xbf, 0xd4, 0xb7, 0xf8, 0x49, 0x97,
	0xda, 0x93, 0x87, 0x89, 0xfc, 0x1b, 0xf3, 0x13, 0x60, 0x0e, 0x73, 0x8b, 0xfe, 0x02, 0x5b, 0x72,
	0x95, 0xc9, 0x3d, 0xe0, 0x8f, 0x60, 0x60, 0x50, 0x15, 0x68, 0xb7, 0xb6, 0x6e, 0x57, 0xf3, 0xdf,
	0x31, 0x88, 0xff, 0x8d, 0x46, 0x99, 0x15, 0x74, 0xde, 0x44, 0xc6, 0xba, 0x08, 0x9e, 0x38, 0x7d,
	0x8b, 0xf4, 0xf5, 0x44, 0x52, 0xcd, 0x9f, 0x01, 0xe4, 0xb5, 0x31, 0x98, 0x87, 0x8d, 0xdb, 0x8f,
	0xdf, 0x21, 0xc1, 0x33, 0x8a, 0x61, 0x1f, 0x57, 0x22, 0x59, 0x20, 0x5d, 0x52, 0x2f, 0xe0, 0xb0,
	0x40, 0x83, 0xbe, 0x17, 0x74, 0x49, 0x8d, 0x3b, 0xd6, 0x49, 0x5e, 0x03, 0xef, 0x96, 0x45, 0xb6,
	0x6a, 0x77, 0xc2, 0x01, 0x09, 0xa7, 0xdb, 0xce, 0xfb, 0xb6, 0x57, 0x3f, 0x01, 0x66, 0x51, 0x15,
	0x59, 0x5d, 0x99, 0x96, 0xc2, 0x1b, 0xc9, 0x51, 0x00, 0x5f, 0x2a, 0xd3, 0xf2, 0x57, 0x70, 0x6c,
	0xb1, 0x31, 0x3a, 0x57, 0x59, 0x63, 0x54, 0x8e, 0x25, 0x56, 0x7d, 0x8e, 0xd3, 0x6d, 0xe3, 0xb2,
	0xe7, 0x5c, 0xc0, 0xf0, 0x06, 0xad, 0x0b, 0xbf, 0xc5, 0x48, 0xd2, 0x2f, 0xf9, 0x14, 0x0e, 0xbc,
	0x37, 0x02, 0x88, 0x86, 0xf2, 0xf4, 0x2b, 0x0c, 0x97, 0xdd, 0x3b, 0xe0, 0xe7, 0x30, 0x59, 0x62,
	0x55, 0xec, 0x6f, 0xfe, 0xc3, 0x3b, 0xb7, 0x60, 0x47, 0x1f, 0x9f, 0xfc, 0x8b, 0xf6, 0xb1, 0xcf,
	0xef, 0xa5, 0xd1, 0x9b, 0x68, 0x35, 0xa0, 0x37, 0xf5, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x14, 0xbb, 0x3a, 0x64, 0x03, 0x00, 0x00,
}
