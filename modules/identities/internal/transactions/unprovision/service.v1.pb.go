// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/unprovision/service.v1.proto

package unprovision

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
	proto.RegisterFile("modules/identities/internal/transactions/unprovision/service.v1.proto", fileDescriptor_22f5fba16661947f)
}

var fileDescriptor_22f5fba16661947f = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x9b, 0x08, 0x15, 0x32, 0x38, 0x74, 0x0c, 0x72, 0x48, 0x8b, 0xba, 0xc8, 0x1d, 0xd5,
	0x45, 0xb2, 0x59, 0x10, 0xed, 0x50, 0xa8, 0xb6, 0x75, 0x90, 0x2e, 0xd7, 0xe4, 0x88, 0x07, 0xc9,
	0x7d, 0x21, 0xf7, 0xb5, 0x0f, 0xe0, 0x13, 0x08, 0x2e, 0xce, 0xe2, 0xe4, 0x93, 0x88, 0x53, 0x41,
	0x07, 0x47, 0x49, 0x9d, 0x7c, 0x0a, 0x69, 0xa3, 0xdc, 0x81, 0x4b, 0xe8, 0x76, 0x70, 0xff, 0xdf,
	0x77, 0xbf, 0xff, 0xc7, 0x79, 0xa7, 0x29, 0x44, 0xd3, 0x44, 0x68, 0x26, 0x23, 0xa1, 0x50, 0xa2,
	0x5c, 0x1e, 0x15, 0x8a, 0x5c, 0xf1, 0x84, 0x61, 0xce, 0x95, 0xe6, 0x21, 0x4a, 0x50, 0x9a, 0x4d,
	0x55, 0x96, 0xc3, 0x4c, 0x6a, 0x09, 0x8a, 0x69, 0x91, 0xcf, 0x64, 0x28, 0xe8, 0xac, 0x4d, 0xb3,
	0x1c, 0x10, 0x1a, 0x2d, 0x83, 0x53, 0x9b, 0xa2, 0x16, 0xe5, 0x6f, 0xc7, 0x00, 0x71, 0x22, 0x18,
	0xcf, 0x24, 0xe3, 0x4a, 0x01, 0xf2, 0x32, 0xb3, 0x1a, 0xe1, 0xaf, 0x67, 0x92, 0x0a, 0xad, 0x79,
	0x6c, 0x4c, 0xfc, 0x8b, 0xb5, 0xc6, 0x58, 0x17, 0x97, 0x42, 0x67, 0xa0, 0xb4, 0x19, 0x79, 0xf8,
	0xe4, 0x78, 0x9b, 0x83, 0xb2, 0x71, 0xe3, 0xc1, 0xf1, 0xea, 0xe7, 0x5c, 0x45, 0x89, 0x68, 0x1c,
	0xd0, 0x0a, 0xa5, 0x69, 0xaf, 0x14, 0xf4, 0x8f, 0x2b, 0xa5, 0x87, 0xff, 0x3d, 0x9a, 0x7b, 0xb7,
	0x6f, 0x5f, 0xf7, 0xee, 0x4e, 0x93, 0xb0, 0x94, 0x2b, 0x4c, 0x84, 0xdd, 0xcc, 0x62, 0x3b, 0xef,
	0xee, 0x4b, 0x41, 0x9c, 0x79, 0x41, 0x9c, 0xcf, 0x82, 0x38, 0x77, 0x0b, 0x52, 0x9b, 0x2f, 0x48,
	0xed, 0x63, 0x41, 0x6a, 0xde, 0x7e, 0x08, 0x69, 0x95, 0xf7, 0x3b, 0x5b, 0xbf, 0x3d, 0xaf, 0xda,
	0xfd, 0x65, 0xf5, 0xbe, 0x73, 0x3d, 0x88, 0x25, 0xde, 0x4c, 0x27, 0x34, 0x84, 0x94, 0x9d, 0x68,
	0x2d, 0xb0, 0x57, 0x4a, 0xfc, 0xad, 0x79, 0x9d, 0x75, 0x3f, 0xba, 0x1b, 0xdd, 0xe1, 0xe8, 0xd9,
	0x6d, 0x75, 0x8d, 0xd2, 0xd0, 0x56, 0x1a, 0x99, 0xec, 0xab, 0x9d, 0x1a, 0xdb, 0xa9, 0xb1, 0x95,
	0x2a, 0x5c, 0x56, 0x21, 0x35, 0x3e, 0xeb, 0x77, 0x7a, 0x02, 0x79, 0xc4, 0x91, 0x7f, 0xbb, 0xbb,
	0x86, 0x08, 0x02, 0x1b, 0x09, 0x02, 0x8b, 0x99, 0xd4, 0x57, 0x9f, 0xe0, 0xe8, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0x56, 0xbd, 0x4b, 0x2a, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.unprovision.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.unprovision.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.unprovision.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/unprovision/service.v1.proto",
}
