// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/provision/service.v1.proto

package provision

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
	proto.RegisterFile("modules/identities/internal/transactions/provision/service.v1.proto", fileDescriptor_f1f72d59a3fd5486)
}

var fileDescriptor_f1f72d59a3fd5486 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xbc, 0xd0, 0x17, 0x32, 0x38, 0x74, 0x0c, 0xe5, 0xa0, 0x55, 0x17, 0xc1, 0x3b,
	0xaa, 0xe0, 0x90, 0xcd, 0x3a, 0x68, 0x87, 0x62, 0xd4, 0xe2, 0x20, 0x5d, 0xae, 0xc9, 0x11, 0x0f,
	0x92, 0x7b, 0x42, 0xee, 0x69, 0x3f, 0x80, 0x9f, 0x40, 0x74, 0x77, 0x70, 0xf4, 0x93, 0x14, 0xa7,
	0x82, 0x8b, 0xa3, 0xa4, 0x4e, 0x7e, 0x0a, 0xa9, 0xb5, 0xc9, 0x81, 0x43, 0xb4, 0xdb, 0x71, 0xfc,
	0xfe, 0xcf, 0xfd, 0xfe, 0xcf, 0x39, 0x47, 0x09, 0x84, 0xe3, 0x58, 0x68, 0x26, 0x43, 0xa1, 0x50,
	0xa2, 0x5c, 0x1c, 0x15, 0x8a, 0x4c, 0xf1, 0x98, 0x61, 0xc6, 0x95, 0xe6, 0x01, 0x4a, 0x50, 0x9a,
	0xa5, 0x19, 0x4c, 0xa4, 0x96, 0xa0, 0x98, 0x16, 0xd9, 0x44, 0x06, 0x82, 0x4e, 0x3a, 0x34, 0xcd,
	0x00, 0xa1, 0xd1, 0x2a, 0xc3, 0xd4, 0xcc, 0xd0, 0x22, 0xe3, 0x36, 0x23, 0x80, 0x28, 0x16, 0x8c,
	0xa7, 0x92, 0x71, 0xa5, 0x00, 0x79, 0x41, 0x20, 0xb8, 0xeb, 0x58, 0x24, 0x42, 0x6b, 0x1e, 0x95,
	0x16, 0xee, 0xe9, 0x1a, 0x43, 0x8c, 0xeb, 0x73, 0xa1, 0x53, 0x50, 0xba, 0x1c, 0xb8, 0xf7, 0x60,
	0x39, 0xff, 0x2f, 0x96, 0x5d, 0x1b, 0x77, 0x96, 0x53, 0x3f, 0xe1, 0x2a, 0x8c, 0x45, 0x63, 0x87,
	0x56, 0xd6, 0xa5, 0xfd, 0xa5, 0x9c, 0x7b, 0xf0, 0x0b, 0x76, 0xf0, 0xd3, 0xa1, 0xbd, 0x75, 0xf3,
	0xf2, 0x7e, 0x6f, 0x93, 0x76, 0x93, 0x25, 0x5c, 0x61, 0x2c, 0xcc, 0x4e, 0x45, 0xb2, 0x3b, 0xb5,
	0xa7, 0x39, 0xb1, 0x66, 0x39, 0xb1, 0xde, 0x72, 0x62, 0xdd, 0xce, 0x49, 0x6d, 0x36, 0x27, 0xb5,
	0xd7, 0x39, 0xa9, 0x39, 0xdb, 0x01, 0x24, 0xd5, 0x6f, 0x77, 0x37, 0xbe, 0xfb, 0x5d, 0x76, 0xfc,
	0x45, 0x65, 0xdf, 0xba, 0x3a, 0x8b, 0x24, 0x5e, 0x8f, 0x47, 0x34, 0x80, 0x84, 0x1d, 0x6a, 0x2d,
	0xb0, 0xbf, 0x14, 0x58, 0x2d, 0xf7, 0xef, 0x4b, 0x7e, 0xb4, 0xff, 0xf5, 0x06, 0xfe, 0x93, 0xdd,
	0xea, 0x95, 0x3a, 0x03, 0x53, 0xc7, 0x5f, 0x91, 0xcf, 0x26, 0x33, 0x34, 0x99, 0x61, 0xc1, 0xe4,
	0xf6, 0x6e, 0x25, 0x33, 0x3c, 0xf6, 0xbb, 0x7d, 0x81, 0x3c, 0xe4, 0xc8, 0x3f, 0xec, 0xcd, 0x92,
	0xf7, 0x3c, 0x33, 0xe0, 0x79, 0x45, 0x62, 0x54, 0xff, 0xfa, 0xf2, 0xfd, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x05, 0x1a, 0xe6, 0x82, 0x10, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.provision.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.provision.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.provision.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/provision/service.v1.proto",
}
