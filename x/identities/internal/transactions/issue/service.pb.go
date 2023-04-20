// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/issue/service.proto

package issue

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
	proto.RegisterFile("x/identities/internal/transactions/issue/service.proto", fileDescriptor_9fb83e1028f7d314)
}

var fileDescriptor_9fb83e1028f7d314 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xfc, 0xa1, 0x7f, 0x08, 0x4e, 0x9d, 0x24, 0xd8, 0xa8, 0x1d, 0x04, 0x97, 0x1c,
	0x54, 0x70, 0xc8, 0x66, 0x96, 0x9a, 0xa1, 0x50, 0xb4, 0x93, 0x74, 0x79, 0x9b, 0xbc, 0xc4, 0x83,
	0xe4, 0xae, 0xe4, 0x7d, 0x2b, 0xce, 0x0e, 0xce, 0xa2, 0xdf, 0xc0, 0x49, 0xfc, 0x24, 0xe2, 0x54,
	0x70, 0x71, 0x94, 0xd4, 0xc9, 0x4f, 0x21, 0x4d, 0x0a, 0x3d, 0x11, 0x6a, 0xd7, 0x3c, 0xbf, 0x27,
	0xbf, 0x7b, 0xee, 0x9c, 0xe3, 0x6b, 0x21, 0x13, 0x54, 0x2c, 0x59, 0x22, 0x09, 0xa9, 0x18, 0x0b,
	0x05, 0x99, 0xe0, 0x02, 0x14, 0x41, 0xcc, 0x52, 0x2b, 0x12, 0x92, 0x68, 0x8a, 0x82, 0xb0, 0xb8,
	0x92, 0x31, 0xfa, 0x93, 0x42, 0xb3, 0x6e, 0xb5, 0x57, 0x2d, 0xdf, 0x84, 0xfd, 0x0a, 0x76, 0x77,
	0x52, 0xad, 0xd3, 0x0c, 0x05, 0x4c, 0xa4, 0x00, 0xa5, 0x34, 0x43, 0x9d, 0x56, 0x65, 0x77, 0x73,
	0x69, 0x8e, 0x44, 0x90, 0x2e, 0xa5, 0x6e, 0xb8, 0x71, 0xcf, 0xf8, 0x74, 0x86, 0x34, 0xd1, 0x8a,
	0x96, 0xff, 0xe8, 0xde, 0x5b, 0xce, 0xff, 0xf3, 0x7a, 0x4a, 0xeb, 0xd6, 0x72, 0x9a, 0xa7, 0xa0,
	0x92, 0x0c, 0x5b, 0x07, 0xfe, 0xda, 0x41, 0x7e, 0xbf, 0x3e, 0x88, 0xdb, 0xfd, 0x83, 0x1b, 0xfe,
	0x16, 0x77, 0xf6, 0x6e, 0xde, 0x3e, 0x1f, 0x6c, 0xb7, 0xb3, 0x2d, 0x72, 0x50, 0x9c, 0xe1, 0x8f,
	0x11, 0x8b, 0x56, 0xf8, 0x64, 0xbf, 0x94, 0x9e, 0x35, 0x2b, 0x3d, 0xeb, 0xa3, 0xf4, 0xac, 0xbb,
	0xb9, 0xd7, 0x98, 0xcd, 0xbd, 0xc6, 0xfb, 0xdc, 0x6b, 0x38, 0xfb, 0xb1, 0xce, 0xd7, 0x3b, 0xc3,
	0xad, 0xe5, 0x9e, 0xc1, 0x62, 0xe0, 0xc0, 0xba, 0xe8, 0xa5, 0x92, 0x2f, 0xa7, 0x63, 0x3f, 0xd6,
	0xb9, 0x38, 0x21, 0x42, 0xee, 0xd7, 0xda, 0x5c, 0x27, 0xd3, 0x0c, 0x49, 0x6c, 0x7a, 0x8b, 0x8f,
	0xf6, 0xbf, 0x68, 0x18, 0x3d, 0xdb, 0xed, 0x68, 0xa5, 0x1f, 0x9a, 0xfa, 0x68, 0x41, 0xbd, 0x9a,
	0xf9, 0xc8, 0xcc, 0x47, 0x55, 0x5e, 0xda, 0x87, 0x6b, 0xf3, 0x51, 0x6f, 0x10, 0xf6, 0x91, 0x21,
	0x01, 0x86, 0x2f, 0x7b, 0x77, 0xc5, 0x06, 0x81, 0x09, 0x07, 0x41, 0x45, 0x8f, 0x9b, 0xd5, 0x33,
	0x1e, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x78, 0x3e, 0x99, 0xc1, 0xb9, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.issue.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.issue.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.issue.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/transactions/issue/service.proto",
}
