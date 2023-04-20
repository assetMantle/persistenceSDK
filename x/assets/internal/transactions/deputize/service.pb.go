// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/transactions/deputize/service.proto

package deputize

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
	proto.RegisterFile("x/assets/internal/transactions/deputize/service.proto", fileDescriptor_34e4a416bdb8b9b6)
}

var fileDescriptor_34e4a416bdb8b9b6 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x13, 0x2a, 0x04, 0xa7, 0x2e, 0x62, 0x28, 0xa7, 0x14, 0x04, 0xa7, 0x3b, 0x54,
	0x5c, 0xb2, 0xa5, 0x14, 0x75, 0x29, 0x14, 0xed, 0x24, 0x5d, 0xbe, 0x26, 0x1f, 0xf1, 0x20, 0xb9,
	0x0b, 0xb9, 0xab, 0x88, 0xa3, 0xfe, 0x00, 0x05, 0xff, 0x41, 0x47, 0x7f, 0x89, 0x38, 0x15, 0x5c,
	0x1c, 0x25, 0x75, 0xf2, 0x57, 0x48, 0x7b, 0x09, 0x04, 0x84, 0xd0, 0x35, 0xef, 0xf3, 0xe6, 0xc9,
	0xfb, 0xc5, 0x3d, 0xbb, 0xe7, 0xa0, 0x35, 0x1a, 0xcd, 0x85, 0x34, 0x98, 0x4b, 0x48, 0xb8, 0xc9,
	0x41, 0x6a, 0x08, 0x8d, 0x50, 0x52, 0xf3, 0x08, 0xb3, 0x99, 0x11, 0x0f, 0xc8, 0x35, 0xe6, 0x77,
	0x22, 0x44, 0x96, 0xe5, 0xca, 0xa8, 0x4e, 0xd7, 0x96, 0x58, 0x9d, 0x65, 0x15, 0xeb, 0x75, 0x63,
	0xa5, 0xe2, 0x04, 0x39, 0x64, 0x82, 0x83, 0x94, 0xca, 0x80, 0x05, 0xd6, 0x5d, 0x6f, 0x63, 0x65,
	0x8a, 0x5a, 0x43, 0x5c, 0x2a, 0xbd, 0x60, 0xd3, 0x5a, 0xed, 0xe9, 0x15, 0xea, 0x4c, 0x49, 0x5d,
	0xbe, 0xe2, 0xe4, 0xd9, 0x71, 0xb7, 0xaf, 0xed, 0x8e, 0xce, 0x93, 0xe3, 0xb6, 0x2f, 0x41, 0x46,
	0x09, 0x76, 0x0e, 0x59, 0xd3, 0x1a, 0x36, 0xb4, 0x9f, 0xe1, 0x1d, 0x37, 0x63, 0xe3, 0xff, 0xda,
	0xde, 0xfe, 0xe3, 0xe7, 0xcf, 0x2b, 0xd9, 0xeb, 0xed, 0xf2, 0x14, 0xa4, 0x59, 0x5d, 0xc4, 0x2e,
	0xa8, 0x4a, 0xfd, 0x39, 0x79, 0x2f, 0xa8, 0xb3, 0x28, 0xa8, 0xf3, 0x5d, 0x50, 0xe7, 0x65, 0x49,
	0x5b, 0x8b, 0x25, 0x6d, 0x7d, 0x2d, 0x69, 0xcb, 0x3d, 0x08, 0x55, 0xda, 0x68, 0xec, 0xef, 0x94,
	0x5b, 0x46, 0xab, 0x71, 0x23, 0xe7, 0xe6, 0x3c, 0x16, 0xe6, 0x76, 0x36, 0x65, 0xa1, 0x4a, 0x79,
	0xb0, 0x2a, 0x0e, 0xad, 0x34, 0x55, 0xd1, 0x2c, 0x41, 0xcd, 0x37, 0x3c, 0xe0, 0x9c, 0x6c, 0x05,
	0xe3, 0xc1, 0x1b, 0xe9, 0x06, 0x56, 0x3e, 0xae, 0xcb, 0x07, 0x25, 0xf4, 0x51, 0xc5, 0x93, 0x7a,
	0x3c, 0xa9, 0xe2, 0x82, 0x1c, 0x35, 0xc5, 0x93, 0x8b, 0x51, 0x7f, 0x88, 0x06, 0x22, 0x30, 0xf0,
	0x4b, 0xa8, 0x45, 0x7d, 0xbf, 0xce, 0xfa, 0x7e, 0x05, 0x4f, 0xdb, 0xeb, 0xbf, 0x77, 0xfa, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x15, 0xce, 0x31, 0x1d, 0xac, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assets.transactions.deputize.Service/Handle", in, out, opts...)
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
		FullMethod: "/assets.transactions.deputize.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.deputize.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/assets/internal/transactions/deputize/service.proto",
}
