// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/queries/orders/service.proto

package orders

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
	proto.RegisterFile("modules/orders/internal/queries/orders/service.proto", fileDescriptor_75467b513dcfdf73)
}

var fileDescriptor_75467b513dcfdf73 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x44, 0x88, 0x17, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x42, 0x44, 0xf5, 0xa0, 0x8a, 0xf4,
	0x20, 0x5c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x26, 0x29, 0x4b, 0x22, 0xad,
	0x02, 0x71, 0x2b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa0, 0x5a, 0xad, 0x48, 0xd3, 0x5a,
	0x5c, 0x90, 0x9f, 0x57, 0x0c, 0x75, 0xab, 0x51, 0x2b, 0x23, 0x17, 0x7b, 0x30, 0xc4, 0xf5, 0x42,
	0x55, 0x5c, 0x6c, 0x1e, 0x89, 0x79, 0x29, 0x39, 0xa9, 0x42, 0xca, 0x7a, 0x58, 0xbd, 0xa0, 0x17,
	0x88, 0x64, 0xb9, 0x94, 0x0a, 0x7e, 0x45, 0x10, 0x6b, 0x94, 0x54, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0x24, 0x2f, 0x24, 0xab, 0x9f, 0x9b, 0x98, 0x57, 0x92, 0x93, 0x0a, 0x73, 0x4c, 0x99, 0x61, 0x52,
	0x6a, 0x49, 0xa2, 0x21, 0x94, 0xeb, 0xf4, 0x8f, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xb8, 0x24, 0x93, 0xf3, 0x73, 0xb1, 0x5b, 0xe5, 0xc4, 0x03, 0x75, 0x7a, 0x00, 0xc8,
	0x2f, 0x01, 0x8c, 0x51, 0xae, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x8e, 0xc5, 0xc5, 0xa9, 0x25, 0xbe, 0x10, 0x3b, 0x61, 0x01, 0x44, 0x5c, 0x40, 0x2d, 0x62, 0x62,
	0xf6, 0x0f, 0xf4, 0x5f, 0xc5, 0x24, 0xea, 0x8f, 0xf0, 0x10, 0xc8, 0x52, 0x08, 0xf7, 0x14, 0x4c,
	0x3c, 0x06, 0x2a, 0x1e, 0x03, 0xe1, 0x3e, 0x62, 0x52, 0xc4, 0x2a, 0x1e, 0xe3, 0x1e, 0xe0, 0xe4,
	0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58, 0x92, 0xf8, 0x8a, 0x49, 0x1c, 0x22, 0x68, 0x65, 0x05, 0x55,
	0x64, 0x65, 0x05, 0x11, 0x48, 0x62, 0x03, 0xc7, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0xc7, 0x88, 0xdc, 0x73, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/orders.queries.orders.Service/Handle", in, out, opts...)
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
		FullMethod: "/orders.queries.orders.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.queries.orders.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/orders/internal/queries/orders/service.proto",
}
