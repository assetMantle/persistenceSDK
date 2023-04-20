// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/transactions/burn/service.proto

package burn

import (
	context "context"
	fmt "fmt"
	math "math"

	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	proto.RegisterFile("x/assets/internal/transactions/burn/service.proto", fileDescriptor_5612498c9ebe80bc)
}

var fileDescriptor_5612498c9ebe80bc = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0x7b, 0xa1, 0x2f, 0x04, 0xa7, 0x8a, 0x50, 0x62, 0x39, 0xb0, 0xb8, 0x7a, 0x87,
	0xba, 0x05, 0x1c, 0x7a, 0x8b, 0x2e, 0x85, 0xa2, 0x9d, 0xa4, 0xcb, 0xd3, 0xf4, 0x88, 0x81, 0xe4,
	0x2e, 0xdc, 0x5d, 0xd4, 0xb9, 0xa3, 0x93, 0xe0, 0x37, 0x70, 0xf4, 0x93, 0x88, 0x53, 0xc1, 0xc5,
	0x51, 0x12, 0x27, 0x3f, 0x85, 0x5c, 0x2f, 0x42, 0x40, 0x02, 0x59, 0xef, 0xff, 0xff, 0xf1, 0x7b,
	0x9e, 0xe7, 0xfc, 0xe3, 0x7b, 0x0a, 0x5a, 0x73, 0xa3, 0x69, 0x22, 0x0c, 0x57, 0x02, 0x52, 0x6a,
	0x14, 0x08, 0x0d, 0x91, 0x49, 0xa4, 0xd0, 0x74, 0x59, 0x28, 0x41, 0x35, 0x57, 0xb7, 0x49, 0xc4,
	0x49, 0xae, 0xa4, 0x91, 0x83, 0xa1, 0x03, 0x48, 0xb3, 0x47, 0x6c, 0x2f, 0x18, 0xc5, 0x52, 0xc6,
	0x29, 0xa7, 0x90, 0x27, 0x14, 0x84, 0x90, 0x06, 0x5c, 0xb8, 0xe5, 0x82, 0x4e, 0xaa, 0x8c, 0x6b,
	0x0d, 0x71, 0xad, 0x0a, 0xce, 0xba, 0x20, 0x8d, 0x97, 0x4b, 0xae, 0x73, 0x29, 0x74, 0x8d, 0x9f,
	0xac, 0x3d, 0xff, 0xff, 0x95, 0x9b, 0x7d, 0x70, 0xe7, 0xf7, 0x2f, 0x40, 0xac, 0x52, 0x3e, 0x38,
	0x20, 0x6d, 0x0b, 0x90, 0xa9, 0xb3, 0x07, 0x47, 0xed, 0x95, 0xf9, 0x5f, 0xdb, 0x78, 0x7f, 0xfd,
	0xfe, 0xf5, 0x84, 0xf6, 0xc6, 0xbb, 0x34, 0x03, 0x61, 0xec, 0x01, 0xdc, 0xd0, 0x16, 0x60, 0x0f,
	0xe8, 0xb5, 0xc4, 0xde, 0xa6, 0xc4, 0xde, 0x67, 0x89, 0xbd, 0xc7, 0x0a, 0xf7, 0x36, 0x15, 0xee,
	0x7d, 0x54, 0xb8, 0xe7, 0x8f, 0x22, 0x99, 0xb5, 0x9a, 0xd8, 0x4e, 0x3d, 0xfa, 0xcc, 0xee, 0x32,
	0xf3, 0xae, 0x59, 0x9c, 0x98, 0x9b, 0x62, 0x49, 0x22, 0x99, 0xd1, 0x89, 0x85, 0xa6, 0x4e, 0x96,
	0xc9, 0x55, 0x91, 0x72, 0x4d, 0x3b, 0xdc, 0xea, 0x19, 0xfd, 0x9b, 0xcc, 0xd9, 0x0b, 0x1a, 0x4e,
	0x9c, 0x74, 0xde, 0x94, 0xb2, 0x42, 0x89, 0xb7, 0xdf, 0x68, 0xd1, 0x8c, 0x16, 0x36, 0x2a, 0xd1,
	0x61, 0x5b, 0xb4, 0x38, 0x9f, 0xb1, 0x29, 0x37, 0xb0, 0x02, 0x03, 0xdf, 0x28, 0x70, 0xb5, 0x30,
	0x6c, 0xf6, 0xc2, 0xd0, 0x16, 0x97, 0xfd, 0xed, 0xc7, 0x9c, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x61, 0xba, 0x54, 0x77, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assets.transactions.burn.Service/Handle", in, out, opts...)
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
		FullMethod: "/assets.transactions.burn.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.burn.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/assets/internal/transactions/burn/service.proto",
}
