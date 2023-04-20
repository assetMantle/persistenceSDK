// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/quash/service.proto

package quash

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
	proto.RegisterFile("x/identities/internal/transactions/quash/service.proto", fileDescriptor_8f8481f42edf41f5)
}

var fileDescriptor_8f8481f42edf41f5 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x15, 0x82, 0x53, 0x27, 0x09, 0x36, 0x6a, 0x07, 0xc1, 0x25, 0x07, 0x15,
	0x1c, 0xb2, 0x99, 0xa5, 0x76, 0x28, 0xb4, 0xda, 0x49, 0xba, 0x7c, 0x4d, 0x3e, 0xd2, 0x83, 0xe4,
	0xae, 0xe6, 0xbe, 0x8a, 0xb3, 0x83, 0xb3, 0xe8, 0x3f, 0x70, 0x12, 0x7f, 0x89, 0x38, 0x15, 0x5c,
	0x1c, 0x25, 0x75, 0xf2, 0x57, 0x48, 0x2f, 0x85, 0x1e, 0x08, 0x35, 0x6b, 0xde, 0xe7, 0xcd, 0x73,
	0xef, 0x9d, 0x73, 0x7a, 0xcb, 0x78, 0x8c, 0x82, 0x38, 0x71, 0x54, 0x8c, 0x0b, 0xc2, 0x5c, 0x40,
	0xca, 0x28, 0x07, 0xa1, 0x20, 0x22, 0x2e, 0x85, 0x62, 0xd7, 0x33, 0x50, 0x13, 0xa6, 0x30, 0xbf,
	0xe1, 0x11, 0xfa, 0xd3, 0x5c, 0x92, 0x6c, 0x34, 0xd7, 0x2d, 0xdf, 0x84, 0x7d, 0x0d, 0xbb, 0x7b,
	0x89, 0x94, 0x49, 0x8a, 0x0c, 0xa6, 0x9c, 0x81, 0x10, 0x92, 0xa0, 0x4c, 0x75, 0xd9, 0xad, 0x2e,
	0xcd, 0x50, 0x29, 0x48, 0x56, 0x52, 0x37, 0xac, 0xdc, 0x33, 0x3e, 0x5d, 0xa0, 0x9a, 0x4a, 0xa1,
	0x56, 0xff, 0x68, 0x3f, 0x5a, 0xce, 0xf6, 0x65, 0x39, 0xa5, 0x71, 0x6f, 0x39, 0xf5, 0x73, 0x10,
	0x71, 0x8a, 0x8d, 0x23, 0x7f, 0xe3, 0x20, 0xbf, 0x57, 0x1e, 0xc4, 0x6d, 0xff, 0xc3, 0x0d, 0xff,
	0x8a, 0x5b, 0x07, 0x77, 0x1f, 0xdf, 0x4f, 0xb6, 0xdb, 0xda, 0x65, 0x19, 0x08, 0x4a, 0xd1, 0x1c,
	0xa1, 0x5b, 0xe1, 0x8b, 0xfd, 0x56, 0x78, 0xd6, 0xbc, 0xf0, 0xac, 0xaf, 0xc2, 0xb3, 0x1e, 0x16,
	0x5e, 0x6d, 0xbe, 0xf0, 0x6a, 0x9f, 0x0b, 0xaf, 0xe6, 0x1c, 0x46, 0x32, 0xdb, 0xec, 0x0c, 0x77,
	0x56, 0x7b, 0xfa, 0xcb, 0x81, 0x7d, 0xeb, 0xaa, 0x93, 0x70, 0x9a, 0xcc, 0xc6, 0x7e, 0x24, 0x33,
	0x76, 0xa6, 0x14, 0x52, 0xaf, 0xd4, 0x66, 0x32, 0x9e, 0xa5, 0xa8, 0x58, 0xd5, 0x5b, 0x7c, 0xb6,
	0xb7, 0xba, 0xc3, 0xc1, 0xab, 0xdd, 0xec, 0xae, 0xf5, 0x43, 0x53, 0x3f, 0x58, 0x52, 0xef, 0x66,
	0x3e, 0x32, 0xf3, 0x91, 0xce, 0x0b, 0xfb, 0x78, 0x63, 0x3e, 0xea, 0xf4, 0xc3, 0x1e, 0x12, 0xc4,
	0x40, 0xf0, 0x63, 0xef, 0xaf, 0xd9, 0x20, 0x30, 0xe1, 0x20, 0xd0, 0xf4, 0xb8, 0xae, 0x9f, 0xf1,
	0xe4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x95, 0x2a, 0xb7, 0xb9, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.quash.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.quash.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.quash.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/transactions/quash/service.proto",
}
