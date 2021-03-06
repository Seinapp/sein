// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protodef/utils.proto

package utils

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_195e9e978a66cfee, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_195e9e978a66cfee, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PingRequest)(nil), "utils.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "utils.PingResponse")
}

func init() { proto.RegisterFile("protodef/utils.proto", fileDescriptor_195e9e978a66cfee) }

var fileDescriptor_195e9e978a66cfee = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x3b, 0xaf, 0xc2, 0x30,
	0x0c, 0x46, 0xef, 0x70, 0xcb, 0x10, 0x1e, 0x43, 0x60, 0x62, 0xec, 0x4e, 0x2d, 0x60, 0x63, 0x04,
	0xb1, 0x23, 0x10, 0x0b, 0x5b, 0x1f, 0x1f, 0x6d, 0x44, 0x49, 0x42, 0xec, 0xf2, 0xfb, 0x51, 0x9b,
	0x05, 0xc6, 0x63, 0xfb, 0x58, 0x47, 0x2d, 0x7c, 0x70, 0xe2, 0x2a, 0xdc, 0xa9, 0x13, 0xd3, 0x72,
	0x36, 0xa0, 0x4e, 0x06, 0x48, 0xa7, 0x6a, 0x7c, 0x32, 0xb6, 0x3e, 0xe3, 0xd5, 0x81, 0x25, 0x9d,
	0xa9, 0x49, 0x44, 0xf6, 0xce, 0x32, 0x36, 0x3b, 0x95, 0x5c, 0xfb, 0x3b, 0xbd, 0x56, 0xff, 0xfd,
	0x42, 0xeb, 0x2c, 0x3e, 0xf9, 0x92, 0x96, 0xf3, 0x9f, 0x59, 0x34, 0xd3, 0xbf, 0xfd, 0xf1, 0x76,
	0xa8, 0x8d, 0x34, 0x5d, 0x91, 0x95, 0xee, 0x49, 0x17, 0x18, 0x9b, 0x7b, 0x4f, 0x0c, 0x63, 0x89,
	0x25, 0x2f, 0x5a, 0x50, 0xed, 0xa8, 0x02, 0x3f, 0xc4, 0xf9, 0x55, 0x83, 0xd6, 0x23, 0x10, 0x87,
	0x92, 0x18, 0xe1, 0x6d, 0x4a, 0x70, 0xcc, 0x2d, 0x46, 0x43, 0xef, 0xf6, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0x52, 0x49, 0x46, 0xc7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UtilsClient is the client API for Utils service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type utilsClient struct {
	cc *grpc.ClientConn
}

func NewUtilsClient(cc *grpc.ClientConn) UtilsClient {
	return &utilsClient{cc}
}

func (c *utilsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/utils.Utils/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilsServer is the server API for Utils service.
type UtilsServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterUtilsServer(s *grpc.Server, srv UtilsServer) {
	s.RegisterService(&_Utils_serviceDesc, srv)
}

func _Utils_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utils.Utils/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Utils_serviceDesc = grpc.ServiceDesc{
	ServiceName: "utils.Utils",
	HandlerType: (*UtilsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Utils_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protodef/utils.proto",
}
