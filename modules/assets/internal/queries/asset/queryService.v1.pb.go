// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/queries/asset/queryService.v1.proto

package asset

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
	proto.RegisterFile("modules/assets/internal/queries/asset/queryService.v1.proto", fileDescriptor_bc3f2158e0c6eb35)
}

var fileDescriptor_bc3f2158e0c6eb35 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xce, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x84, 0x89, 0x83, 0x79, 0x95, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x22,
	0x10, 0x4d, 0x7a, 0x50, 0xb5, 0x7a, 0x60, 0xae, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa,
	0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31,
	0x44, 0x8f, 0x14, 0x29, 0x16, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0xc0, 0x2d, 0x94, 0xb2,
	0x21, 0x49, 0x73, 0x71, 0x41, 0x7e, 0x5e, 0x31, 0xc2, 0xb9, 0x46, 0x3d, 0x8c, 0x5c, 0xac, 0x81,
	0x20, 0x29, 0xa1, 0x26, 0x46, 0x2e, 0x56, 0x47, 0x90, 0x52, 0x21, 0x25, 0x3d, 0x6c, 0x7e, 0xd0,
	0x0b, 0x44, 0xb2, 0x5e, 0x4a, 0x19, 0xaf, 0x1a, 0x88, 0x2d, 0x4a, 0x06, 0x4d, 0x97, 0x9f, 0x4c,
	0x66, 0xd2, 0x12, 0xd2, 0xd0, 0xcf, 0x4d, 0xcc, 0x2b, 0x01, 0x79, 0x1f, 0xe2, 0xc4, 0x32, 0xc3,
	0xa4, 0xd4, 0x92, 0x44, 0x43, 0xa8, 0xcb, 0xaa, 0xc1, 0x54, 0x7c, 0x66, 0x7c, 0x4a, 0xad, 0xd3,
	0x7f, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x92, 0x48, 0xce, 0xcf,
	0xc5, 0x6a, 0xa9, 0x93, 0x70, 0x20, 0x52, 0x4c, 0x84, 0x19, 0x06, 0x80, 0x3c, 0x16, 0xc0, 0x18,
	0xe5, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf6, 0x98, 0x2f,
	0xc4, 0x11, 0xb0, 0xf0, 0x22, 0x2a, 0xdc, 0x16, 0x31, 0x31, 0x3b, 0x06, 0x3a, 0xae, 0x62, 0x12,
	0x71, 0x84, 0xd8, 0x1c, 0x08, 0xb5, 0x19, 0xcc, 0x3d, 0x05, 0x13, 0x8e, 0x81, 0x0a, 0xc7, 0x80,
	0xb9, 0x8f, 0x98, 0x14, 0xb0, 0x09, 0xc7, 0xb8, 0x07, 0x38, 0xf9, 0xa6, 0x96, 0x24, 0xa6, 0x24,
	0x96, 0x24, 0xbe, 0x62, 0x12, 0x83, 0x28, 0xb1, 0xb2, 0x82, 0xaa, 0xb1, 0xb2, 0x02, 0x0b, 0x24,
	0xb1, 0x81, 0xe3, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x09, 0x68, 0x0d, 0x85, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Asset(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Asset(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/assets.queries.asset.Query/Asset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Asset(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Asset(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Asset not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Asset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Asset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assets.queries.asset.Query/Asset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Asset(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.queries.asset.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Asset",
			Handler:    _Query_Asset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/queries/asset/queryService.v1.proto",
}
