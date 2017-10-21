// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vv.proto

/*
Package vv is a generated protocol buffer package.

It is generated from these files:
	proto/vv.proto

It has these top-level messages:
	PingRequest
	PingReply
	OpenFileRequest
	OpenFileReply
	ProtoHeader
*/
package vv

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

type ProtoHeader_ServerKind int32

const (
	ProtoHeader_UNKNOWN_SERVER_KIND ProtoHeader_ServerKind = 0
	ProtoHeader_SERVER              ProtoHeader_ServerKind = 1
	ProtoHeader_CLIENT              ProtoHeader_ServerKind = 2
)

var ProtoHeader_ServerKind_name = map[int32]string{
	0: "UNKNOWN_SERVER_KIND",
	1: "SERVER",
	2: "CLIENT",
}
var ProtoHeader_ServerKind_value = map[string]int32{
	"UNKNOWN_SERVER_KIND": 0,
	"SERVER":              1,
	"CLIENT":              2,
}

func (x ProtoHeader_ServerKind) String() string {
	return proto.EnumName(ProtoHeader_ServerKind_name, int32(x))
}
func (ProtoHeader_ServerKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type ProtoHeader_ConnKind int32

const (
	ProtoHeader_UNKNOWN_CONN_KIND ProtoHeader_ConnKind = 0
	ProtoHeader_LISTEN            ProtoHeader_ConnKind = 1
	ProtoHeader_DIAL              ProtoHeader_ConnKind = 2
)

var ProtoHeader_ConnKind_name = map[int32]string{
	0: "UNKNOWN_CONN_KIND",
	1: "LISTEN",
	2: "DIAL",
}
var ProtoHeader_ConnKind_value = map[string]int32{
	"UNKNOWN_CONN_KIND": 0,
	"LISTEN":            1,
	"DIAL":              2,
}

func (x ProtoHeader_ConnKind) String() string {
	return proto.EnumName(ProtoHeader_ConnKind_name, int32(x))
}
func (ProtoHeader_ConnKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

type PingRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PingReply struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type OpenFileRequest struct {
	FileName string `protobuf:"bytes,1,opt,name=fileName" json:"fileName,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Dir      string `protobuf:"bytes,3,opt,name=dir" json:"dir,omitempty"`
}

func (m *OpenFileRequest) Reset()                    { *m = OpenFileRequest{} }
func (m *OpenFileRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenFileRequest) ProtoMessage()               {}
func (*OpenFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OpenFileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *OpenFileRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *OpenFileRequest) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

type OpenFileReply struct {
	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	IsBsdiff bool   `protobuf:"varint,2,opt,name=isBsdiff" json:"isBsdiff,omitempty"`
	Crc      []byte `protobuf:"bytes,3,opt,name=crc,proto3" json:"crc,omitempty"`
}

func (m *OpenFileReply) Reset()                    { *m = OpenFileReply{} }
func (m *OpenFileReply) String() string            { return proto.CompactTextString(m) }
func (*OpenFileReply) ProtoMessage()               {}
func (*OpenFileReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OpenFileReply) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *OpenFileReply) GetIsBsdiff() bool {
	if m != nil {
		return m.IsBsdiff
	}
	return false
}

func (m *OpenFileReply) GetCrc() []byte {
	if m != nil {
		return m.Crc
	}
	return nil
}

// 握手的header
type ProtoHeader struct {
	Version    string                 `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Token      string                 `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	ServerKind ProtoHeader_ServerKind `protobuf:"varint,3,opt,name=serverKind,enum=vv.ProtoHeader_ServerKind" json:"serverKind,omitempty"`
	ConnKind   ProtoHeader_ConnKind   `protobuf:"varint,4,opt,name=connKind,enum=vv.ProtoHeader_ConnKind" json:"connKind,omitempty"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ProtoHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ProtoHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ProtoHeader) GetServerKind() ProtoHeader_ServerKind {
	if m != nil {
		return m.ServerKind
	}
	return ProtoHeader_UNKNOWN_SERVER_KIND
}

func (m *ProtoHeader) GetConnKind() ProtoHeader_ConnKind {
	if m != nil {
		return m.ConnKind
	}
	return ProtoHeader_UNKNOWN_CONN_KIND
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "vv.PingRequest")
	proto.RegisterType((*PingReply)(nil), "vv.PingReply")
	proto.RegisterType((*OpenFileRequest)(nil), "vv.OpenFileRequest")
	proto.RegisterType((*OpenFileReply)(nil), "vv.OpenFileReply")
	proto.RegisterType((*ProtoHeader)(nil), "vv.ProtoHeader")
	proto.RegisterEnum("vv.ProtoHeader_ServerKind", ProtoHeader_ServerKind_name, ProtoHeader_ServerKind_value)
	proto.RegisterEnum("vv.ProtoHeader_ConnKind", ProtoHeader_ConnKind_name, ProtoHeader_ConnKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VvServer service

type VvServerClient interface {
	// 握手消息
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	// 打开文件，返回文件修改流
	// TODO 打开大文件
	OpenFile(ctx context.Context, in *OpenFileRequest, opts ...grpc.CallOption) (VvServer_OpenFileClient, error)
}

type vvServerClient struct {
	cc *grpc.ClientConn
}

func NewVvServerClient(cc *grpc.ClientConn) VvServerClient {
	return &vvServerClient{cc}
}

func (c *vvServerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/vv.VvServer/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vvServerClient) OpenFile(ctx context.Context, in *OpenFileRequest, opts ...grpc.CallOption) (VvServer_OpenFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VvServer_serviceDesc.Streams[0], c.cc, "/vv.VvServer/OpenFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &vvServerOpenFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VvServer_OpenFileClient interface {
	Recv() (*OpenFileReply, error)
	grpc.ClientStream
}

type vvServerOpenFileClient struct {
	grpc.ClientStream
}

func (x *vvServerOpenFileClient) Recv() (*OpenFileReply, error) {
	m := new(OpenFileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for VvServer service

type VvServerServer interface {
	// 握手消息
	Ping(context.Context, *PingRequest) (*PingReply, error)
	// 打开文件，返回文件修改流
	// TODO 打开大文件
	OpenFile(*OpenFileRequest, VvServer_OpenFileServer) error
}

func RegisterVvServerServer(s *grpc.Server, srv VvServerServer) {
	s.RegisterService(&_VvServer_serviceDesc, srv)
}

func _VvServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VvServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vv.VvServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VvServerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VvServer_OpenFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OpenFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VvServerServer).OpenFile(m, &vvServerOpenFileServer{stream})
}

type VvServer_OpenFileServer interface {
	Send(*OpenFileReply) error
	grpc.ServerStream
}

type vvServerOpenFileServer struct {
	grpc.ServerStream
}

func (x *vvServerOpenFileServer) Send(m *OpenFileReply) error {
	return x.ServerStream.SendMsg(m)
}

var _VvServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vv.VvServer",
	HandlerType: (*VvServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _VvServer_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenFile",
			Handler:       _VvServer_OpenFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/vv.proto",
}

func init() { proto.RegisterFile("proto/vv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xe7, 0xac, 0x6c, 0xee, 0xdd, 0xbf, 0xec, 0x0e, 0x44, 0x94, 0x17, 0x46, 0x24, 0xa4,
	0x3e, 0x05, 0x34, 0x10, 0x48, 0x48, 0x3c, 0xd0, 0xae, 0x88, 0xaa, 0x95, 0x5b, 0xb9, 0xa3, 0x13,
	0x4f, 0x53, 0x97, 0xba, 0x9d, 0xb5, 0xcc, 0x29, 0x49, 0x6a, 0xb4, 0x4f, 0xc4, 0xd7, 0x44, 0x76,
	0x96, 0x3f, 0x9b, 0xf6, 0x76, 0xce, 0xcd, 0x39, 0x3f, 0x5b, 0xce, 0x85, 0xc3, 0x75, 0x9a, 0xe4,
	0xc9, 0x7b, 0xad, 0x43, 0x2b, 0xd0, 0xd1, 0x3a, 0x78, 0x0b, 0x7b, 0x13, 0xa9, 0x56, 0x5c, 0xfc,
	0xd9, 0x88, 0x2c, 0x47, 0x84, 0x96, 0x9a, 0xdf, 0x09, 0x8f, 0x9c, 0x92, 0x4e, 0x9b, 0x5b, 0x1d,
	0xbc, 0x81, 0x76, 0x11, 0x59, 0xc7, 0xf7, 0xcf, 0x06, 0x7e, 0xc3, 0xd1, 0x78, 0x2d, 0xd4, 0x0f,
	0x19, 0x8b, 0x92, 0xe3, 0x03, 0x5d, 0xca, 0x58, 0xb0, 0x3a, 0x5a, 0x79, 0xf4, 0x60, 0x37, 0x4a,
	0x54, 0x2e, 0x54, 0xee, 0x39, 0xa7, 0xa4, 0xb3, 0xcf, 0x4b, 0x8b, 0x2e, 0x6c, 0x2f, 0x64, 0xea,
	0x6d, 0xdb, 0x82, 0x91, 0xc1, 0x25, 0x1c, 0xd4, 0x68, 0x73, 0x7e, 0xa3, 0x4c, 0x1e, 0x97, 0x7d,
	0xa0, 0x32, 0xeb, 0x66, 0x0b, 0xb9, 0x5c, 0x5a, 0x2e, 0xe5, 0x95, 0x37, 0xe0, 0x28, 0x8d, 0x2c,
	0x78, 0x9f, 0x1b, 0x19, 0xfc, 0x73, 0x60, 0x6f, 0x62, 0x5e, 0xe1, 0xa7, 0x98, 0x2f, 0x44, 0x6a,
	0xb8, 0x5a, 0xa4, 0x99, 0x4c, 0xd4, 0xc3, 0x7d, 0x4b, 0x8b, 0x2f, 0xe1, 0x45, 0x9e, 0xdc, 0x0a,
	0x65, 0xa1, 0x6d, 0x5e, 0x18, 0xfc, 0x0a, 0x90, 0x89, 0x54, 0x8b, 0x74, 0x28, 0xd5, 0xc2, 0x82,
	0x0f, 0xcf, 0xfc, 0x50, 0xeb, 0xb0, 0x01, 0x0d, 0xa7, 0x55, 0x82, 0x37, 0xd2, 0xf8, 0x09, 0x68,
	0x94, 0x28, 0x65, 0x9b, 0x2d, 0xdb, 0xf4, 0x9e, 0x36, 0x7b, 0x0f, 0xdf, 0x79, 0x95, 0x0c, 0xbe,
	0x01, 0xd4, 0x3c, 0x7c, 0x0d, 0x27, 0xbf, 0xd8, 0x90, 0x8d, 0x2f, 0xd9, 0xd5, 0xb4, 0xcf, 0x67,
	0x7d, 0x7e, 0x35, 0x1c, 0xb0, 0x73, 0x77, 0x0b, 0x01, 0x76, 0x8a, 0x81, 0x4b, 0x8c, 0xee, 0x8d,
	0x06, 0x7d, 0x76, 0xe1, 0x3a, 0xc1, 0x17, 0xa0, 0x25, 0x14, 0x5f, 0xc1, 0x71, 0x59, 0xee, 0x8d,
	0x19, 0x6b, 0x54, 0x47, 0x83, 0xe9, 0x45, 0x9f, 0xb9, 0x04, 0x29, 0xb4, 0xce, 0x07, 0xdf, 0x47,
	0xae, 0x73, 0x16, 0x03, 0x9d, 0xe9, 0xe2, 0x64, 0xec, 0x40, 0xcb, 0xac, 0x02, 0x1e, 0xd9, 0xfb,
	0xd6, 0x7b, 0xe3, 0x1f, 0xd4, 0x83, 0x75, 0x7c, 0x1f, 0x6c, 0xe1, 0x67, 0xa0, 0xe5, 0x8f, 0xc3,
	0x13, 0xf3, 0xf1, 0xc9, 0x86, 0xf8, 0xc7, 0x8f, 0x87, 0xb6, 0xf5, 0x81, 0x74, 0xdf, 0x81, 0x17,
	0x25, 0x77, 0xe1, 0x4a, 0xe6, 0x37, 0x9b, 0xeb, 0xf0, 0xef, 0x5c, 0xad, 0x6e, 0x45, 0x74, 0xb3,
	0x51, 0xa1, 0xd6, 0xdd, 0xdd, 0x99, 0xb6, 0x6f, 0x34, 0x21, 0xd7, 0x3b, 0x76, 0x83, 0x3f, 0xfe,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa6, 0xc2, 0xfc, 0xd3, 0x02, 0x00, 0x00,
}