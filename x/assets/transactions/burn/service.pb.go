// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/burn/service.proto

package burn

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
	proto.RegisterFile("assets/transactions/burn/service.proto", fileDescriptor_54ae5fdb34eafef3)
}

var fileDescriptor_54ae5fdb34eafef3 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4a, 0xf4, 0x40,
	0x10, 0xc7, 0x2f, 0x81, 0xef, 0x3e, 0x08, 0x56, 0x27, 0x36, 0x51, 0x52, 0x5c, 0x61, 0x75, 0xec,
	0x1e, 0x1e, 0x82, 0x6e, 0x97, 0x34, 0xda, 0x04, 0x0e, 0xbd, 0x4a, 0x02, 0x32, 0x97, 0x5b, 0x62,
	0x20, 0xd9, 0x3d, 0xb2, 0x1b, 0xb1, 0xf6, 0x09, 0x04, 0x1f, 0x40, 0xb0, 0xd4, 0x17, 0x11, 0xab,
	0x03, 0x1b, 0x4b, 0xc9, 0xd9, 0xe8, 0x53, 0xc8, 0xde, 0x2e, 0xb8, 0x16, 0x29, 0xd2, 0xce, 0xfe,
	0x7f, 0x33, 0xbf, 0x19, 0xd6, 0xdb, 0x07, 0x21, 0xa8, 0x14, 0x58, 0x56, 0xc0, 0x04, 0xa4, 0x32,
	0xe7, 0x4c, 0xe0, 0x79, 0x5d, 0x31, 0x2c, 0x68, 0x75, 0x9d, 0xa7, 0x14, 0x2d, 0x2b, 0x2e, 0xf9,
	0x60, 0xb4, 0xc9, 0x95, 0xc0, 0x64, 0x41, 0x51, 0xc9, 0x17, 0x75, 0x41, 0x05, 0xd2, 0x2c, 0xb2,
	0x59, 0xa4, 0x58, 0x7f, 0x2f, 0xe3, 0x3c, 0x2b, 0x28, 0x86, 0x65, 0x8e, 0x81, 0x31, 0x2e, 0x41,
	0x3f, 0x6e, 0x7a, 0xf9, 0xed, 0x33, 0x4b, 0x2a, 0x04, 0x64, 0x66, 0xa6, 0x3f, 0x69, 0xcd, 0x59,
	0x95, 0xcb, 0x8a, 0x8a, 0x25, 0x67, 0xc2, 0x40, 0x07, 0xcf, 0x8e, 0xf7, 0xff, 0x5c, 0xab, 0x0f,
	0x1e, 0x1c, 0xaf, 0x7f, 0x0a, 0x6c, 0x51, 0xd0, 0xc1, 0x21, 0xea, 0xb2, 0x00, 0x8a, 0xb5, 0x88,
	0x1f, 0x76, 0xc3, 0x66, 0xbf, 0x95, 0x33, 0xa3, 0x35, 0xdc, 0xbd, 0x7d, 0xfb, 0xbc, 0x77, 0x77,
	0x86, 0xdb, 0x58, 0x77, 0xc1, 0x66, 0x2b, 0x05, 0x44, 0x5f, 0xee, 0x4b, 0x13, 0x38, 0xab, 0x26,
	0x70, 0x3e, 0x9a, 0xc0, 0xb9, 0x5b, 0x07, 0xbd, 0xd5, 0x3a, 0xe8, 0xbd, 0xaf, 0x83, 0x9e, 0x37,
	0x4e, 0x79, 0xd9, 0x69, 0x7a, 0xb4, 0x65, 0xf6, 0x9e, 0xaa, 0x43, 0x4c, 0x9d, 0x8b, 0xe3, 0x2c,
	0x97, 0x57, 0xf5, 0x1c, 0xa5, 0xbc, 0xc4, 0xa1, 0x82, 0x62, 0x2d, 0x60, 0x1a, 0xe1, 0x1b, 0xdc,
	0x76, 0xe0, 0x47, 0xf7, 0x5f, 0x18, 0x87, 0xb3, 0xe8, 0xc9, 0x1d, 0x85, 0x96, 0x41, 0x6c, 0x0c,
	0x42, 0x6d, 0x30, 0xb3, 0x0d, 0xa2, 0xba, 0x62, 0xaf, 0x7f, 0xe2, 0x89, 0x89, 0x27, 0x3a, 0x9e,
	0xd8, 0xf1, 0x44, 0xc5, 0x1b, 0xf7, 0xa8, 0x4b, 0x3c, 0x39, 0x99, 0x46, 0x31, 0x95, 0xb0, 0x00,
	0x09, 0xdf, 0xee, 0xd8, 0x42, 0x09, 0x31, 0x2c, 0x21, 0x1a, 0x26, 0xc4, 0xa6, 0x09, 0x51, 0xf8,
	0xbc, 0xbf, 0xf9, 0x20, 0x93, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0xed, 0x05, 0x66, 0xf3,
	0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.transactions.burn.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.assets.transactions.burn.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.transactions.burn.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/transactions/burn/service.proto",
}
