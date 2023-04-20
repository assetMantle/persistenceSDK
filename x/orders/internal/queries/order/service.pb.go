// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/queries/order/service.proto

package order

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
	proto.RegisterFile("x/orders/internal/queries/order/service.proto", fileDescriptor_3f87d2e2ac30d428)
}

var fileDescriptor_3f87d2e2ac30d428 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0x9b, 0x13, 0x2a, 0x1c, 0x4e, 0x47, 0x11, 0x29, 0x25, 0x48, 0x75, 0x35, 0xa1, 0xed,
	0x76, 0x83, 0x60, 0x17, 0x5d, 0x4a, 0x5b, 0xdd, 0xa4, 0x4b, 0xda, 0xfb, 0x53, 0x03, 0x77, 0x49,
	0x9b, 0xe4, 0xaa, 0xae, 0xfd, 0x04, 0x82, 0xdf, 0xc0, 0xd1, 0x4f, 0x22, 0x4e, 0x05, 0x17, 0x47,
	0xb9, 0x73, 0x72, 0xf4, 0x13, 0x48, 0x2f, 0x29, 0x38, 0x1c, 0xd6, 0xf1, 0x3d, 0xde, 0xe3, 0xf7,
	0x92, 0xbf, 0x7f, 0x72, 0x47, 0xa5, 0x8a, 0x40, 0x69, 0xca, 0x85, 0x01, 0x25, 0x58, 0x4c, 0xe7,
	0x29, 0x28, 0x0e, 0xda, 0xfa, 0x54, 0x83, 0x5a, 0xf0, 0x09, 0x90, 0x99, 0x92, 0x46, 0x06, 0x35,
	0x1b, 0x26, 0x2e, 0x43, 0x0a, 0x59, 0x6f, 0x4c, 0xa5, 0x9c, 0xc6, 0x40, 0xd9, 0x8c, 0x53, 0x26,
	0x84, 0x34, 0xcc, 0x70, 0x29, 0xb4, 0xed, 0xd4, 0xdb, 0xdb, 0x10, 0x6b, 0x75, 0x7f, 0x09, 0xf3,
	0x14, 0xb4, 0x71, 0x9d, 0xce, 0x3f, 0x3b, 0x7a, 0x26, 0x85, 0x76, 0xe3, 0xda, 0x4b, 0xe4, 0xef,
	0x5e, 0xd9, 0xb9, 0xc1, 0xad, 0x5f, 0xbd, 0x60, 0x22, 0x8a, 0x21, 0x68, 0x92, 0xb2, 0xcd, 0x64,
	0xf8, 0x0b, 0x5a, 0x3f, 0xfa, 0x33, 0x63, 0x21, 0xcd, 0xe3, 0xe5, 0xdb, 0xe7, 0xa3, 0x87, 0x83,
	0x06, 0x4d, 0x98, 0x30, 0x31, 0x6c, 0x16, 0x2e, 0x5a, 0x63, 0x30, 0xac, 0x65, 0x65, 0xf7, 0x1b,
	0xbd, 0x64, 0x18, 0xad, 0x32, 0x8c, 0x3e, 0x32, 0x8c, 0x1e, 0x72, 0x5c, 0x59, 0xe5, 0xb8, 0xf2,
	0x9e, 0xe3, 0x8a, 0x7f, 0x30, 0x91, 0x49, 0x29, 0xa8, 0xbb, 0xe7, 0x66, 0x0f, 0xd6, 0xef, 0x18,
	0xa0, 0xeb, 0xd3, 0x29, 0x37, 0x37, 0xe9, 0x98, 0x4c, 0x64, 0x42, 0xcf, 0xb4, 0x06, 0xd3, 0xb3,
	0xc4, 0x44, 0x46, 0x69, 0x0c, 0x9a, 0x6e, 0xf9, 0x9d, 0x27, 0x6f, 0xa7, 0x3f, 0xec, 0x3f, 0x7b,
	0xb5, 0xbe, 0x85, 0x0d, 0x1d, 0xac, 0x90, 0xaf, 0x1b, 0x7b, 0xe4, 0xec, 0x51, 0x21, 0x33, 0xef,
	0xb0, 0xcc, 0x1e, 0x9d, 0x0f, 0xba, 0x3d, 0x30, 0x2c, 0x62, 0x86, 0x7d, 0x79, 0xfb, 0x36, 0x12,
	0x86, 0x2e, 0x13, 0x86, 0x85, 0x31, 0xae, 0x16, 0x07, 0xe8, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0xdd, 0x70, 0xae, 0x4e, 0x02, 0x00, 0x00,
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
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/orders.queries.order.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.queries.order.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.queries.order.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/orders/internal/queries/order/service.proto",
}
