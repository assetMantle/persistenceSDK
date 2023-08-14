// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/unwrap/service.proto

package unwrap

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("assets/transactions/unwrap/service.proto", fileDescriptor_f4511852f8b26e2c)
}

var fileDescriptor_f4511852f8b26e2c = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0x07, 0xf0, 0x26, 0xf0, 0xf6, 0x85, 0xf0, 0x4e, 0x85, 0x17, 0x21, 0x68, 0x86, 0x4e, 0x4e,
	0x77, 0x50, 0x11, 0xe1, 0x9c, 0x52, 0x07, 0x5d, 0x02, 0x45, 0xdb, 0x45, 0x02, 0x72, 0x4d, 0x8f,
	0x18, 0x48, 0xee, 0x42, 0x9e, 0xab, 0x3a, 0xfb, 0x09, 0x04, 0xbf, 0x81, 0xe0, 0xe2, 0xe6, 0xb7,
	0x10, 0xa7, 0x82, 0x8b, 0xa3, 0xa4, 0x4e, 0xee, 0xee, 0xd2, 0x3c, 0x07, 0x9e, 0x83, 0x85, 0xac,
	0x4f, 0xee, 0xf7, 0xbf, 0xff, 0x73, 0xc4, 0xdb, 0xe6, 0x00, 0x42, 0x03, 0xd5, 0x15, 0x97, 0xc0,
	0x13, 0x9d, 0x29, 0x09, 0x74, 0x2e, 0x2f, 0x2b, 0x5e, 0x52, 0x10, 0xd5, 0x45, 0x96, 0x08, 0x52,
	0x56, 0x4a, 0xab, 0x1e, 0x69, 0x4e, 0x16, 0x5c, 0xea, 0x5c, 0x90, 0x42, 0xcd, 0xe6, 0xb9, 0x00,
	0x9c, 0x01, 0xb1, 0x35, 0x41, 0xed, 0x6f, 0xa6, 0x4a, 0xa5, 0xb9, 0xa0, 0xbc, 0xcc, 0x28, 0x97,
	0x52, 0x69, 0x8e, 0x9f, 0x9b, 0x34, 0x7f, 0xdd, 0xbd, 0x85, 0x00, 0xe0, 0xa9, 0xb9, 0xd7, 0xdf,
	0x5d, 0x73, 0xd2, 0x9a, 0x9d, 0x55, 0x02, 0x4a, 0x25, 0xc1, 0xb0, 0xc1, 0xa3, 0xe3, 0xfd, 0x3d,
	0xc1, 0x05, 0x7a, 0xf7, 0x8e, 0xd7, 0x3d, 0xe2, 0x72, 0x96, 0x8b, 0xde, 0x5e, 0xcb, 0x35, 0x48,
	0x84, 0x65, 0xfc, 0x83, 0xb6, 0x70, 0xfc, 0x3d, 0x3b, 0x36, 0xd5, 0xfa, 0x5b, 0xd7, 0x2f, 0xef,
	0xb7, 0xee, 0x46, 0xff, 0x3f, 0xc5, 0x1c, 0x6a, 0x76, 0x43, 0x32, 0xfc, 0x74, 0x9f, 0xea, 0xc0,
	0x59, 0xd4, 0x81, 0xf3, 0x56, 0x07, 0xce, 0xcd, 0x32, 0xe8, 0x2c, 0x96, 0x41, 0xe7, 0x75, 0x19,
	0x74, 0xbc, 0x41, 0xa2, 0x8a, 0x96, 0x0d, 0x86, 0xff, 0xcc, 0xfe, 0xa3, 0xd5, 0x83, 0x8c, 0x9c,
	0xd3, 0xfd, 0x34, 0xd3, 0xe7, 0xf3, 0x29, 0x49, 0x54, 0x41, 0xc3, 0x15, 0x8b, 0xb0, 0x84, 0x89,
	0xa2, 0x57, 0xf4, 0xf7, 0xa7, 0xbe, 0x73, 0xff, 0x84, 0x51, 0x38, 0x9e, 0x3c, 0xb8, 0x24, 0xb4,
	0x5a, 0x44, 0xa6, 0x45, 0x88, 0x2d, 0xc6, 0x76, 0x8b, 0x49, 0x03, 0x9f, 0x7f, 0x80, 0xd8, 0x80,
	0x18, 0x41, 0x6c, 0x83, 0x18, 0x41, 0xed, 0xb2, 0x76, 0x20, 0x3e, 0x1c, 0x0d, 0x23, 0xa1, 0xf9,
	0x8c, 0x6b, 0xfe, 0xe1, 0x0e, 0x2c, 0xcc, 0x98, 0xd1, 0x0c, 0x23, 0x81, 0x31, 0xdb, 0x33, 0x86,
	0x01, 0xd3, 0x6e, 0xf3, 0xcb, 0xec, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x67, 0x6b, 0xff,
	0x0d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.transactions.unwrap.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.assets.transactions.unwrap.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.transactions.unwrap.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/transactions/unwrap/service.proto",
}
