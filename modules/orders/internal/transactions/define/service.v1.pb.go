// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/define/service.v1.proto

package define

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
	proto.RegisterFile("modules/orders/internal/transactions/define/service.v1.proto", fileDescriptor_de30ef8cb9579292)
}

var fileDescriptor_de30ef8cb9579292 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4b, 0xfb, 0x40,
	0x1c, 0xc5, 0x9b, 0xfc, 0xa0, 0x3f, 0xc8, 0xe0, 0x50, 0x10, 0xe1, 0xa8, 0x37, 0x54, 0x70, 0xbc,
	0xa3, 0xba, 0x05, 0x17, 0x4b, 0xc1, 0x22, 0x94, 0x16, 0x2d, 0x0e, 0xd2, 0xe5, 0xdb, 0xe4, 0x6b,
	0x3c, 0x48, 0xee, 0xca, 0xdd, 0xb5, 0x9b, 0x8b, 0x93, 0xa3, 0xe0, 0xee, 0xe0, 0xe8, 0x5f, 0x22,
	0x4e, 0x05, 0x17, 0x47, 0x49, 0x9d, 0xfc, 0x2b, 0xc4, 0x9e, 0xa5, 0x01, 0x69, 0x21, 0x6b, 0xde,
	0xe7, 0xe5, 0xbd, 0xf7, 0xbd, 0xe0, 0x28, 0x53, 0xf1, 0x24, 0x45, 0xc3, 0x95, 0x8e, 0x51, 0x1b,
	0x2e, 0xa4, 0x45, 0x2d, 0x21, 0xe5, 0x56, 0x83, 0x34, 0x10, 0x59, 0xa1, 0xa4, 0xe1, 0x31, 0x5e,
	0x09, 0x89, 0xdc, 0xa0, 0x9e, 0x8a, 0x08, 0xd9, 0xb4, 0xc9, 0xc6, 0x5a, 0x59, 0x55, 0x23, 0xce,
	0xc5, 0x8a, 0x30, 0x73, 0x30, 0xa9, 0x27, 0x4a, 0x25, 0x29, 0x72, 0x18, 0x0b, 0x0e, 0x52, 0x2a,
	0x0b, 0x4e, 0x5e, 0x38, 0x49, 0xa9, 0xdc, 0x0c, 0x8d, 0x81, 0x64, 0x95, 0x4b, 0x3a, 0x65, 0xdc,
	0x85, 0x6f, 0x67, 0x68, 0xc6, 0x4a, 0x9a, 0xd5, 0x9f, 0x0e, 0xee, 0xbc, 0xe0, 0xff, 0xb9, 0x9b,
	0x55, 0xbb, 0x09, 0xaa, 0x1d, 0x90, 0x71, 0x8a, 0xb5, 0x3d, 0xb6, 0x7e, 0x18, 0xeb, 0xba, 0x36,
	0x84, 0x6f, 0x82, 0x06, 0x7f, 0x43, 0x1b, 0xbb, 0xb7, 0x6f, 0x9f, 0x0f, 0xfe, 0x4e, 0x63, 0x9b,
	0x67, 0x20, 0x6d, 0x8a, 0xcb, 0xf6, 0xce, 0xd2, 0x7a, 0xf4, 0x5f, 0x72, 0xea, 0xcd, 0x72, 0xea,
	0x7d, 0xe4, 0xd4, 0xbb, 0x9f, 0xd3, 0xca, 0x6c, 0x4e, 0x2b, 0xef, 0x73, 0x5a, 0x09, 0x68, 0xa4,
	0xb2, 0x0d, 0x69, 0xad, 0xad, 0xdf, 0x09, 0x17, 0xcd, 0xfe, 0xcf, 0xaa, 0xbe, 0x77, 0x79, 0x9a,
	0x08, 0x7b, 0x3d, 0x19, 0xb1, 0x48, 0x65, 0xfc, 0xd8, 0x18, 0xb4, 0x5d, 0x17, 0xb9, 0x3c, 0x5c,
	0x89, 0x03, 0x3e, 0xf9, 0xff, 0x7a, 0x83, 0xf6, 0xb3, 0x4f, 0x7a, 0xae, 0xc0, 0xa0, 0x58, 0xa0,
	0xbd, 0x40, 0x5e, 0x97, 0xe2, 0xb0, 0x28, 0x0e, 0x9d, 0x98, 0xfb, 0xfb, 0xeb, 0xc5, 0xe1, 0x49,
	0xbf, 0xd5, 0x45, 0x0b, 0x31, 0x58, 0xf8, 0xf2, 0xeb, 0x0e, 0x0c, 0xc3, 0x22, 0x19, 0x86, 0x0e,
	0x1d, 0x55, 0x17, 0x4f, 0x76, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xbb, 0x5a, 0xb9, 0xb4,
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
	err := c.cc.Invoke(ctx, "/orders.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/orders.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/orders/internal/transactions/define/service.v1.proto",
}
