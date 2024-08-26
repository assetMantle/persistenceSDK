// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/queries/parameters/service.proto

package parameters

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	proto.RegisterFile("AssetMantle/modules/x/assets/queries/parameters/service.proto", fileDescriptor_d720f17bfda39b7b)
}

var fileDescriptor_d720f17bfda39b7b = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0x07, 0xf0, 0x5c, 0xc0, 0x0e, 0xc1, 0xa9, 0x63, 0x87, 0x1b, 0x7c, 0x80, 0x3b, 0xb0, 0x4e,
	0x27, 0x15, 0x5a, 0x05, 0x5d, 0x02, 0xa9, 0x2e, 0x41, 0x0e, 0xe4, 0xda, 0x7e, 0x68, 0x20, 0xc9,
	0xa5, 0x77, 0x17, 0xa9, 0x8f, 0xe0, 0x26, 0xf8, 0x06, 0x8e, 0x3e, 0x89, 0x38, 0x75, 0x74, 0x94,
	0x64, 0xf3, 0x15, 0x5c, 0x24, 0xbd, 0xc3, 0xb4, 0x63, 0xba, 0x5e, 0xf8, 0xfd, 0xef, 0xff, 0xe5,
	0xbb, 0x60, 0x34, 0xd6, 0x1a, 0x4c, 0x28, 0x72, 0x93, 0x02, 0xcd, 0xe4, 0xa2, 0x4c, 0x41, 0xd3,
	0x15, 0x15, 0xcd, 0xa9, 0xa6, 0xcb, 0x12, 0x54, 0x02, 0x9a, 0x16, 0x42, 0x89, 0x0c, 0x0c, 0x28,
	0x4d, 0x35, 0xa8, 0xc7, 0x64, 0x0e, 0xa4, 0x50, 0xd2, 0xc8, 0x3e, 0xdd, 0xe2, 0xc4, 0x71, 0xb2,
	0x22, 0x96, 0x13, 0xc7, 0x49, 0xcb, 0x07, 0xe7, 0x5d, 0xef, 0x6b, 0x8e, 0x9e, 0xee, 0x14, 0x2c,
	0x4b, 0xd0, 0xc6, 0xde, 0x3a, 0xb8, 0xd8, 0x37, 0x44, 0x17, 0x32, 0xd7, 0xae, 0xfb, 0xf1, 0x2b,
	0x0a, 0x0e, 0xa6, 0xcd, 0x87, 0xfe, 0x33, 0x0a, 0x7a, 0x57, 0x22, 0x5f, 0xa4, 0xd0, 0x1f, 0x91,
	0x8e, 0x13, 0x91, 0x4d, 0xc4, 0xb5, 0xed, 0x37, 0x38, 0xdb, 0x97, 0xdb, 0x66, 0x47, 0xde, 0xe4,
	0xd7, 0xff, 0xa8, 0x30, 0x5a, 0x57, 0x18, 0x7d, 0x57, 0x18, 0xbd, 0xd4, 0xd8, 0x5b, 0xd7, 0xd8,
	0xfb, 0xaa, 0xb1, 0x17, 0x0c, 0xe7, 0x32, 0xeb, 0x9a, 0x3f, 0x39, 0xbc, 0xb1, 0x0b, 0x8b, 0x9a,
	0x99, 0x23, 0x74, 0xcb, 0xee, 0x13, 0xf3, 0x50, 0xce, 0xc8, 0x5c, 0x66, 0xb4, 0xe3, 0x6f, 0x7c,
	0xf3, 0x7b, 0xe3, 0x30, 0x1e, 0x4f, 0xa3, 0x77, 0x7f, 0x67, 0xef, 0xa1, 0xab, 0x11, 0xdb, 0x72,
	0x76, 0xa6, 0xa6, 0x46, 0xf4, 0x4f, 0x3f, 0x77, 0x04, 0x77, 0x82, 0xc7, 0xdc, 0x0a, 0xee, 0x04,
	0x6f, 0x45, 0xe5, 0x9f, 0x76, 0x14, 0xfc, 0x32, 0x9a, 0x84, 0x60, 0xc4, 0x42, 0x18, 0xf1, 0xe3,
	0x9f, 0x6c, 0x69, 0xc6, 0x1c, 0x67, 0x2c, 0x66, 0xcc, 0x06, 0x30, 0xe6, 0x12, 0x18, 0x6b, 0x23,
	0x66, 0xbd, 0xcd, 0xd3, 0x18, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x4b, 0x56, 0x33, 0x17,
	0x03, 0x00, 0x00,
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
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.assets.queries.parameters.Query/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AssetMantle.modules.x.assets.queries.parameters.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.assets.queries.parameters.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/assets/queries/parameters/service.proto",
}
