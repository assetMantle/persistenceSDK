// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/revoke/service.v1.proto

package revoke

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
	proto.RegisterFile("modules/identities/internal/transactions/revoke/service.v1.proto", fileDescriptor_af3caf0160ded92d)
}

var fileDescriptor_af3caf0160ded92d = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xc8, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x04, 0x31, 0xf3, 0x4a,
	0x52, 0x8b, 0xf2, 0x12, 0x73, 0xf4, 0x4b, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0xf5, 0x8b, 0x52, 0xcb, 0xf2, 0xb3, 0x53, 0xf5, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0xf5, 0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x10, 0x3a, 0xf5, 0x90,
	0x35, 0xe8, 0x41, 0x34, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64,
	0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x42, 0xa4, 0xc1, 0xba, 0xa5, 0x48, 0xb6, 0x3f, 0x37,
	0xb5, 0xb8, 0x38, 0x31, 0x1d, 0x61, 0xbf, 0x94, 0x0f, 0xa9, 0x26, 0x20, 0x89, 0x05, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0x23, 0x4c, 0x33, 0x9a, 0xca, 0xc8, 0xc5, 0x1e, 0x0c, 0xf1, 0xa2, 0x50,
	0x27, 0x23, 0x17, 0x9b, 0x47, 0x62, 0x5e, 0x4a, 0x4e, 0xaa, 0x90, 0xba, 0x1e, 0x7e, 0x5f, 0xea,
	0xf9, 0x42, 0x9c, 0x25, 0x65, 0x4c, 0x48, 0x61, 0x08, 0xa6, 0xed, 0x4a, 0x8a, 0x4d, 0x97, 0x9f,
	0x4c, 0x66, 0x92, 0x56, 0x92, 0xd4, 0xcf, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0x45, 0xf6, 0x0a, 0x44,
	0x9b, 0xd3, 0x36, 0xa6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x52, 0x4a,
	0xce, 0xcf, 0x25, 0x60, 0xab, 0x13, 0x1f, 0xd4, 0x4f, 0x61, 0x86, 0x01, 0x20, 0x6f, 0x06, 0x30,
	0x46, 0xf9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16, 0x17,
	0xa7, 0x96, 0xf8, 0x42, 0xac, 0x86, 0x85, 0x26, 0x89, 0xa1, 0xba, 0x88, 0x89, 0xd9, 0x33, 0x24,
	0x68, 0x15, 0x93, 0x9c, 0x27, 0xc2, 0x21, 0x21, 0xc8, 0x0e, 0x09, 0x02, 0x2b, 0x3b, 0x85, 0xac,
	0x20, 0x06, 0x59, 0x41, 0x0c, 0x44, 0xc1, 0x23, 0x26, 0x2d, 0xfc, 0x0a, 0x62, 0xdc, 0x03, 0x9c,
	0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0x31, 0x29, 0x20, 0x14, 0x5b, 0x59, 0x21,
	0xab, 0xb6, 0xb2, 0x82, 0x28, 0x4f, 0x62, 0x03, 0xc7, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x60, 0xce, 0x02, 0x42, 0xe9, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.revoke.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.revoke.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.revoke.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/revoke/service.v1.proto",
}
