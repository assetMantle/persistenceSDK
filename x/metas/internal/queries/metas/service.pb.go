// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/metas/internal/queries/metas/service.proto

package metas

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
	proto.RegisterFile("x/metas/internal/queries/metas/service.proto", fileDescriptor_0152c3ca8673c48f)
}

var fileDescriptor_0152c3ca8673c48f = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x08, 0x15, 0x82, 0x53, 0x8a, 0x08, 0x55, 0x4f, 0xac, 0xab, 0xdc, 0xd1, 0xba,
	0x05, 0x1c, 0xec, 0xa2, 0x4b, 0xa0, 0xd5, 0x4d, 0xba, 0x5c, 0xdb, 0x47, 0x3d, 0x48, 0xee, 0xd2,
	0xdc, 0x4b, 0xd1, 0x49, 0xf0, 0x13, 0x08, 0x7e, 0x03, 0x47, 0x3f, 0x89, 0x38, 0x15, 0x5c, 0x1c,
	0x25, 0x71, 0x12, 0x3f, 0x84, 0xe4, 0x72, 0x88, 0x43, 0xa8, 0x6e, 0xf7, 0xff, 0xf1, 0x1e, 0xbf,
	0xff, 0x3d, 0xef, 0xf0, 0x9a, 0xc5, 0x80, 0x5c, 0x33, 0x21, 0x11, 0x52, 0xc9, 0x23, 0x36, 0xcf,
	0x20, 0x15, 0xa0, 0x2d, 0xd6, 0x90, 0x2e, 0xc4, 0x04, 0x68, 0x92, 0x2a, 0x54, 0x7e, 0xcb, 0x40,
	0x6a, 0x47, 0xa8, 0x49, 0xed, 0x9d, 0x99, 0x52, 0xb3, 0x08, 0x18, 0x4f, 0x04, 0xe3, 0x52, 0x2a,
	0xe4, 0x28, 0x94, 0xd4, 0xd5, 0x4a, 0xbb, 0xfb, 0x87, 0xa0, 0x4c, 0x37, 0xe7, 0x30, 0xcf, 0x40,
	0xa3, 0x5d, 0xe9, 0xfd, 0x6f, 0x45, 0x27, 0x4a, 0x6a, 0xdb, 0xac, 0x77, 0xeb, 0xad, 0x5f, 0x54,
	0x55, 0x7d, 0xf4, 0x9a, 0x67, 0x5c, 0x4e, 0x23, 0xf0, 0xf7, 0x69, 0x4d, 0x5f, 0x3a, 0xfc, 0x65,
	0x6c, 0x77, 0x56, 0x8d, 0x54, 0x86, 0xce, 0xc1, 0xdd, 0xeb, 0xc7, 0x83, 0xbb, 0xeb, 0x6f, 0xb3,
	0x98, 0x4b, 0x8c, 0xc0, 0xd6, 0x58, 0x74, 0xc7, 0x80, 0xbc, 0x5b, 0xa5, 0xfe, 0x97, 0xf3, 0x9c,
	0x13, 0x67, 0x99, 0x13, 0xe7, 0x3d, 0x27, 0xce, 0x7d, 0x41, 0x1a, 0xcb, 0x82, 0x34, 0xde, 0x0a,
	0xd2, 0xf0, 0xb6, 0x26, 0x2a, 0xae, 0xd3, 0xf4, 0x37, 0x6c, 0xe5, 0x41, 0xf9, 0x85, 0x81, 0x73,
	0x79, 0x3c, 0x13, 0x78, 0x95, 0x8d, 0xe9, 0x44, 0xc5, 0xec, 0x44, 0x6b, 0xc0, 0xd0, 0xfa, 0xd4,
	0x34, 0x8b, 0x40, 0xb3, 0xd5, 0x77, 0x79, 0x74, 0xd7, 0xc2, 0x61, 0xf8, 0xe4, 0xb6, 0xc2, 0x9f,
	0x3f, 0x94, 0x2a, 0x93, 0x5e, 0x2c, 0x1d, 0x59, 0x3a, 0x32, 0x29, 0x77, 0xf7, 0x6a, 0xe8, 0xe8,
	0x74, 0xd0, 0x2f, 0x1f, 0x53, 0x8e, 0xfc, 0xd3, 0xdd, 0x34, 0x2c, 0x08, 0xec, 0x48, 0x10, 0x98,
	0x3c, 0x6e, 0x9a, 0xb3, 0x1f, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x90, 0x19, 0x6a, 0x4d, 0x40,
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
	err := c.cc.Invoke(ctx, "/metas.queries.metas.Service/Handle", in, out, opts...)
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
		FullMethod: "/metas.queries.metas.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metas.queries.metas.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/metas/internal/queries/metas/service.proto",
}
