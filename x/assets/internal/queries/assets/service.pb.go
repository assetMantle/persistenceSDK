// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/queries/assets/service.proto

package assets

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
	proto.RegisterFile("x/assets/internal/queries/assets/service.proto", fileDescriptor_9a172b94948eebbe)
}

var fileDescriptor_9a172b94948eebbe = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xc0, 0xbf, 0x7f, 0x08, 0x4e, 0x19, 0x8b, 0x46, 0x10, 0x1c, 0x3a, 0xf4, 0x8e,
	0xaa, 0x45, 0xc9, 0x22, 0xe9, 0xa2, 0x4b, 0xa0, 0xd5, 0x4d, 0x02, 0x72, 0x6d, 0x5f, 0xea, 0x41,
	0x72, 0xd7, 0xe6, 0x2e, 0x45, 0x57, 0x3f, 0x81, 0xe0, 0x37, 0xe8, 0xe8, 0xe0, 0xe7, 0x10, 0xa7,
	0x82, 0x8b, 0x63, 0x49, 0x9d, 0xfc, 0x14, 0x92, 0xde, 0x5b, 0xad, 0x53, 0xe3, 0xf8, 0xbc, 0x79,
	0x7e, 0xef, 0x3d, 0xcf, 0x4b, 0x1c, 0x72, 0x4b, 0x99, 0x52, 0xa0, 0x15, 0xe5, 0x42, 0x43, 0x2a,
	0x58, 0x4c, 0xc7, 0x19, 0xa4, 0x1c, 0xd4, 0x6a, 0xae, 0x20, 0x9d, 0xf0, 0x3e, 0x90, 0x51, 0x2a,
	0xb5, 0x74, 0xeb, 0xcb, 0x69, 0xc2, 0x84, 0x8e, 0x81, 0x24, 0x72, 0x90, 0xc5, 0xa0, 0x88, 0x71,
	0x12, 0x04, 0x51, 0xd6, 0xb6, 0x87, 0x52, 0x0e, 0x63, 0xa0, 0x6c, 0xc4, 0x29, 0x13, 0x42, 0x6a,
	0xa6, 0xb9, 0x14, 0xca, 0x2c, 0xaa, 0x1d, 0x6d, 0x7c, 0xb8, 0x90, 0x77, 0xd7, 0x29, 0x8c, 0x33,
	0x50, 0x1a, 0xa9, 0x56, 0x69, 0x4a, 0x8d, 0xa4, 0x50, 0x98, 0xfa, 0xe0, 0xd9, 0x72, 0xfe, 0x5f,
	0x9a, 0x1e, 0xee, 0xd4, 0x72, 0xaa, 0xe7, 0x4c, 0x0c, 0x62, 0x70, 0x8f, 0x49, 0xe9, 0x36, 0xa4,
	0x5b, 0xec, 0xbd, 0x30, 0x61, 0x6a, 0x27, 0x7f, 0x07, 0x4d, 0x9e, 0xbd, 0xfd, 0xfb, 0xb7, 0x8f,
	0x47, 0x7b, 0xd7, 0xdd, 0xa1, 0x06, 0x5e, 0xc5, 0x9e, 0x34, 0x7b, 0xa0, 0x59, 0x13, 0x65, 0x7b,
	0x6e, 0xbf, 0xe4, 0x9e, 0x35, 0xcb, 0x3d, 0x6b, 0x9e, 0x7b, 0xd6, 0xc3, 0xc2, 0xab, 0xcc, 0x16,
	0x5e, 0xe5, 0x7d, 0xe1, 0x55, 0x9c, 0x46, 0x5f, 0x26, 0xe5, 0x9f, 0x6f, 0x6f, 0x61, 0xef, 0x4e,
	0x71, 0x88, 0x8e, 0x75, 0x75, 0x3a, 0xe4, 0xfa, 0x26, 0xeb, 0x91, 0xbe, 0x4c, 0x68, 0x50, 0x58,
	0x42, 0x93, 0x03, 0xb7, 0xd0, 0x4d, 0x07, 0x9e, 0xda, 0xff, 0x82, 0x30, 0xe8, 0x06, 0x4f, 0x76,
	0x3d, 0x58, 0x0b, 0x11, 0x62, 0x88, 0xe0, 0xa7, 0x34, 0xff, 0x96, 0xaf, 0xbf, 0xbc, 0x11, 0x7a,
	0x23, 0xf3, 0x31, 0x42, 0x2f, 0xca, 0xdc, 0x6e, 0x95, 0xf6, 0x46, 0x67, 0x9d, 0x76, 0x08, 0x9a,
	0x0d, 0x98, 0x66, 0x9f, 0x76, 0x63, 0x8d, 0xf3, 0x7d, 0x04, 0x7d, 0xdf, 0x58, 0x7d, 0x1f, 0xd1,
	0xd5, 0xa0, 0x57, 0x5d, 0xfe, 0x1a, 0x87, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x4d, 0x9d,
	0xdf, 0x02, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.queries.assets.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.assets.queries.assets.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.queries.assets.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/assets/internal/queries/assets/service.proto",
}
