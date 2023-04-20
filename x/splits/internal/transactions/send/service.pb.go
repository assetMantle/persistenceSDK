// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/send/service.proto

package send

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
	proto.RegisterFile("x/splits/internal/transactions/send/service.proto", fileDescriptor_60f861d8bd22a550)
}

var fileDescriptor_60f861d8bd22a550 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x15, 0x82, 0x53, 0x45, 0x28, 0xb1, 0x04, 0x2c, 0xae, 0xe6, 0x50, 0xb7,
	0x80, 0x83, 0x59, 0x74, 0x29, 0x14, 0xd3, 0x49, 0xba, 0x7c, 0x4d, 0x3e, 0xe2, 0xc1, 0xe5, 0x2e,
	0xe4, 0xbb, 0xaa, 0x73, 0x47, 0x27, 0xc1, 0x7f, 0xe0, 0xe8, 0x2f, 0x11, 0xa7, 0x82, 0x8b, 0xa3,
	0xa4, 0x4e, 0xfe, 0x0a, 0x49, 0xaf, 0x85, 0x80, 0x04, 0xb2, 0xe6, 0x7d, 0x5e, 0x9e, 0xf7, 0xcb,
	0x39, 0xa7, 0x8f, 0x8c, 0x72, 0xc1, 0x35, 0x31, 0x2e, 0x35, 0x16, 0x12, 0x04, 0xd3, 0x05, 0x48,
	0x82, 0x58, 0x73, 0x25, 0x89, 0x11, 0xca, 0x84, 0x11, 0x16, 0xf7, 0x3c, 0x46, 0x3f, 0x2f, 0x94,
	0x56, 0xbd, 0xbe, 0x29, 0xf8, 0x75, 0xce, 0xaf, 0x38, 0x77, 0x90, 0x2a, 0x95, 0x0a, 0x64, 0x90,
	0x73, 0x06, 0x52, 0x2a, 0x0d, 0x26, 0x5c, 0xf7, 0xdc, 0x56, 0xaa, 0x0c, 0x89, 0x20, 0xdd, 0xa8,
	0xdc, 0x8b, 0x36, 0x95, 0xda, 0x97, 0x1b, 0xa4, 0x5c, 0x49, 0xda, 0xd4, 0xcf, 0x16, 0x96, 0xb3,
	0x1b, 0x99, 0xed, 0xbd, 0x07, 0xa7, 0x7b, 0x0d, 0x32, 0x11, 0xd8, 0x3b, 0xf2, 0x9b, 0x0e, 0xf0,
	0x47, 0xc6, 0xee, 0x9e, 0x34, 0x23, 0x93, 0xff, 0xb6, 0xe1, 0xe1, 0xe2, 0xf3, 0xe7, 0xc5, 0x3e,
	0x18, 0xee, 0xb3, 0x0c, 0xa4, 0x16, 0xb8, 0x1d, 0x5d, 0x15, 0xc2, 0x27, 0xfb, 0xbd, 0xf4, 0xac,
	0x65, 0xe9, 0x59, 0xdf, 0xa5, 0x67, 0x3d, 0xaf, 0xbc, 0xce, 0x72, 0xe5, 0x75, 0xbe, 0x56, 0x5e,
	0xc7, 0x19, 0xc4, 0x2a, 0x6b, 0x34, 0x85, 0x7b, 0x9b, 0xe9, 0xe3, 0xea, 0x96, 0xb1, 0x75, 0x1b,
	0xa6, 0x5c, 0xdf, 0xcd, 0x67, 0x7e, 0xac, 0x32, 0x76, 0x49, 0x84, 0x7a, 0x64, 0x64, 0x99, 0x4a,
	0xe6, 0x02, 0x89, 0xb5, 0xf8, 0x57, 0xaf, 0xf6, 0x4e, 0x34, 0x89, 0xde, 0xec, 0x7e, 0x64, 0xa4,
	0x93, 0xba, 0x34, 0x42, 0x99, 0x7c, 0x6c, 0xa3, 0x69, 0x3d, 0x9a, 0x56, 0x51, 0x69, 0x1f, 0x37,
	0x45, 0xd3, 0xab, 0x71, 0x38, 0x42, 0x0d, 0x09, 0x68, 0xf8, 0xb5, 0x5d, 0x83, 0x05, 0x41, 0x9d,
	0x0b, 0x82, 0x0a, 0x9c, 0x75, 0xd7, 0x0f, 0x73, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x70, 0x9c,
	0xb9, 0x62, 0x77, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/splits.transactions.send.Service/Handle", in, out, opts...)
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
		FullMethod: "/splits.transactions.send.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "splits.transactions.send.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/splits/internal/transactions/send/service.proto",
}
