// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_http.proto

/*
Package proto2 is a generated protocol buffer package.

It is generated from these files:
	hello_http.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package proto2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "./google/api"

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "proto2.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "proto2.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloHttp service

type HelloHttpClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloHttpClient struct {
	cc *grpc.ClientConn
}

func NewHelloHttpClient(cc *grpc.ClientConn) HelloHttpClient {
	return &helloHttpClient{cc}
}

func (c *helloHttpClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/proto2.HelloHttp/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloHttp service

type HelloHttpServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterHelloHttpServer(s *grpc.Server, srv HelloHttpServer) {
	s.RegisterService(&_HelloHttp_serviceDesc, srv)
}

func _HelloHttp_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloHttpServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto2.HelloHttp/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloHttpServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto2.HelloHttp",
	HandlerType: (*HelloHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloHttp_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_http.proto",
}

func init() { proto.RegisterFile("hello_http.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53,
	0x46, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79,
	0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x55, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e,
	0x20, 0x9d, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x1a, 0x17, 0x17, 0x54, 0x4d,
	0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x11, 0x8c,
	0x6b, 0x14, 0xcd, 0xc5, 0x09, 0x56, 0xe7, 0x51, 0x52, 0x52, 0x20, 0xe4, 0xc7, 0xc5, 0x11, 0x9c,
	0x58, 0x09, 0xe6, 0x0b, 0x89, 0x40, 0x2c, 0x33, 0xd2, 0x43, 0xb6, 0x4a, 0x4a, 0x08, 0x4d, 0xb4,
	0x20, 0xa7, 0x52, 0x49, 0xa2, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x42, 0x4a, 0xbc, 0xfa, 0xa9, 0x15,
	0x89, 0xb9, 0x05, 0x39, 0xa9, 0xfa, 0xa9, 0xc9, 0x19, 0xf9, 0x56, 0x8c, 0x5a, 0x49, 0x10, 0xef,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0xfe, 0xe5, 0x1c, 0xe9, 0x00, 0x00, 0x00,
}
