// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/metas/queries/parameters/service.proto

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
	proto.RegisterFile("AssetMantle/modules/x/metas/queries/parameters/service.proto", fileDescriptor_9108a8da77f2a61f)
}

var fileDescriptor_9108a8da77f2a61f = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x06, 0xe0, 0x38, 0xd2, 0xed, 0x10, 0xdd, 0xa9, 0x63, 0x07, 0x0f, 0x3c, 0x80, 0x2d, 0xb5,
	0x13, 0x56, 0x19, 0x1a, 0x06, 0x58, 0x22, 0xa5, 0xb0, 0x44, 0xc8, 0x12, 0x72, 0xdb, 0x23, 0xa8,
	0x94, 0xc4, 0xa9, 0xed, 0xa0, 0xf2, 0x04, 0xac, 0xf0, 0x0a, 0x8c, 0x3c, 0x09, 0x62, 0xea, 0xc8,
	0x88, 0x92, 0x8d, 0x9d, 0x1d, 0xc5, 0xb1, 0x48, 0x3b, 0xa6, 0x6b, 0x8e, 0xbe, 0xff, 0xfc, 0x27,
	0x49, 0x30, 0x9d, 0x69, 0x0d, 0x26, 0x12, 0xb9, 0x49, 0x81, 0x66, 0x72, 0x55, 0xa6, 0xa0, 0xe9,
	0x96, 0x66, 0x60, 0x84, 0xa6, 0x9b, 0x12, 0xd4, 0x1a, 0x34, 0x2d, 0x84, 0x12, 0x19, 0x18, 0x50,
	0x9a, 0x6a, 0x50, 0x0f, 0xeb, 0x25, 0x90, 0x42, 0x49, 0x23, 0x87, 0x64, 0x4f, 0x13, 0xa7, 0xc9,
	0x96, 0x58, 0x4d, 0x9c, 0x26, 0x9d, 0x1e, 0x85, 0x3d, 0xb7, 0x35, 0x8f, 0x1e, 0x6f, 0x15, 0x6c,
	0x4a, 0xd0, 0xa6, 0xdd, 0x39, 0x3a, 0x3f, 0x32, 0x43, 0x17, 0x32, 0xd7, 0xae, 0xf8, 0xf8, 0x05,
	0x05, 0xff, 0xe6, 0xcd, 0x60, 0xf8, 0x84, 0x82, 0xc1, 0xa5, 0xc8, 0x57, 0x29, 0x0c, 0xa7, 0x3d,
	0xcf, 0x21, 0x36, 0xe1, 0xaa, 0x6d, 0x37, 0x3a, 0x3b, 0x52, 0xb7, 0xbd, 0x4e, 0xbc, 0xf0, 0xc7,
	0x7f, 0xaf, 0x30, 0xda, 0x55, 0x18, 0x7d, 0x55, 0x18, 0x3d, 0xd7, 0xd8, 0xdb, 0xd5, 0xd8, 0xfb,
	0xac, 0xb1, 0x17, 0x8c, 0x97, 0x32, 0xeb, 0x19, 0x1f, 0xfe, 0xbf, 0x6e, 0x3f, 0x55, 0xdc, 0x1c,
	0x1c, 0xa3, 0x9b, 0xd3, 0xbb, 0xb5, 0xb9, 0x2f, 0x17, 0x64, 0x29, 0x33, 0xda, 0xef, 0x15, 0xbe,
	0xfa, 0x83, 0x59, 0x94, 0x44, 0xf3, 0xf8, 0xcd, 0x3f, 0x28, 0x11, 0xb9, 0x12, 0x09, 0x89, 0x6c,
	0x89, 0xb9, 0x2b, 0x11, 0xff, 0xc9, 0x8f, 0x03, 0xc0, 0x1d, 0xe0, 0x09, 0xb7, 0x80, 0x3b, 0xc0,
	0x3b, 0x50, 0xf9, 0xac, 0x1f, 0xe0, 0x17, 0x71, 0xd8, 0xcc, 0x56, 0xc2, 0x88, 0x6f, 0x7f, 0xb2,
	0x87, 0x19, 0x73, 0x9a, 0xb1, 0x84, 0x31, 0xeb, 0x19, 0x73, 0x01, 0x8c, 0x75, 0x09, 0x8b, 0x81,
	0xfd, 0x25, 0x26, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc4, 0x78, 0xfa, 0x0b, 0x03, 0x00,
	0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.metas.queries.parameters.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.metas.queries.parameters.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.metas.queries.parameters.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/metas/queries/parameters/service.proto",
}
