// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/queries/splits/service.proto

package splits

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
	proto.RegisterFile("splits/queries/splits/service.proto", fileDescriptor_9c0a8ea65c05e2f4)
}

var fileDescriptor_9c0a8ea65c05e2f4 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4a, 0xc3, 0x40,
	0x1c, 0xc7, 0x9b, 0x80, 0x15, 0x82, 0x53, 0xc6, 0xa2, 0x11, 0x14, 0x87, 0x0a, 0xbd, 0xa3, 0x8a,
	0x7f, 0xc8, 0xd6, 0x2e, 0xba, 0x04, 0x5a, 0xb3, 0x49, 0x40, 0xae, 0xed, 0x8f, 0x7a, 0x90, 0xe4,
	0xd2, 0xdc, 0xa5, 0xe8, 0xea, 0x13, 0x08, 0xbe, 0x41, 0x47, 0x07, 0x9f, 0x43, 0x9c, 0x0a, 0x0e,
	0x3a, 0x4a, 0xea, 0xe4, 0x53, 0x48, 0xbc, 0x5f, 0xb5, 0x42, 0x91, 0x38, 0x7e, 0x93, 0xcf, 0xf7,
	0xcf, 0x1d, 0x67, 0x6d, 0xcb, 0x24, 0xe4, 0x4a, 0xd2, 0x51, 0x06, 0x29, 0x07, 0x49, 0x51, 0x4a,
	0x48, 0xc7, 0xbc, 0x0f, 0x24, 0x49, 0x85, 0x12, 0x76, 0x9d, 0x49, 0x09, 0x2a, 0x62, 0xb1, 0x0a,
	0x81, 0x44, 0x62, 0x90, 0x85, 0x20, 0x89, 0x26, 0x09, 0x1a, 0x51, 0xd6, 0xd6, 0x87, 0x42, 0x0c,
	0x43, 0xa0, 0x2c, 0xe1, 0x94, 0xc5, 0xb1, 0x50, 0x4c, 0x71, 0x11, 0x4b, 0x1d, 0x54, 0xab, 0x2f,
	0x6f, 0x2b, 0xe4, 0xf5, 0x45, 0x0a, 0xa3, 0x0c, 0xa4, 0x42, 0x74, 0xf7, 0x6f, 0x54, 0x26, 0x22,
	0x96, 0xb8, 0x6f, 0xef, 0xc1, 0xb0, 0x56, 0x7d, 0xbd, 0xd8, 0x9e, 0x18, 0x56, 0xf5, 0x94, 0xc5,
	0x83, 0x10, 0xec, 0x23, 0x52, 0x7a, 0x37, 0xe9, 0x16, 0xb9, 0x67, 0x7a, 0x41, 0xed, 0xf8, 0xff,
	0x46, 0xbd, 0x67, 0x6b, 0xe7, 0xe6, 0xf9, 0xfd, 0xce, 0xdc, 0xb4, 0x37, 0xa8, 0x36, 0xcf, 0x67,
	0x8f, 0x9b, 0x3d, 0x50, 0xac, 0x89, 0xb2, 0xfd, 0x62, 0x3e, 0xe6, 0x8e, 0x31, 0xcd, 0x1d, 0xe3,
	0x2d, 0x77, 0x8c, 0xdb, 0x99, 0x53, 0x99, 0xce, 0x9c, 0xca, 0xeb, 0xcc, 0xa9, 0x58, 0x8d, 0xbe,
	0x88, 0xca, 0xd7, 0xb7, 0xd7, 0xf0, 0xdc, 0x9d, 0xe2, 0x22, 0x3a, 0xc6, 0xf9, 0xe1, 0x90, 0xab,
	0xcb, 0xac, 0x47, 0xfa, 0x22, 0xa2, 0xad, 0x22, 0xc5, 0xd3, 0x3b, 0x30, 0x85, 0x5e, 0xd1, 0xa5,
	0xf7, 0x3a, 0x31, 0x57, 0x5a, 0x9e, 0xdf, 0xf5, 0xef, 0xcd, 0x7a, 0x6b, 0xa1, 0xdb, 0xc3, 0x6e,
	0xff, 0xe7, 0xac, 0xfc, 0x5b, 0x3e, 0xfd, 0x62, 0x03, 0x64, 0x03, 0xfd, 0x33, 0x40, 0x16, 0x65,
	0x6e, 0x1e, 0x94, 0x66, 0x83, 0x93, 0x4e, 0xdb, 0x03, 0xc5, 0x06, 0x4c, 0xb1, 0x0f, 0xb3, 0xb1,
	0xe0, 0x73, 0x5d, 0x34, 0xba, 0xae, 0x46, 0x5d, 0x17, 0xad, 0xf3, 0x0f, 0xbd, 0xea, 0xd7, 0x8b,
	0xd8, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x18, 0x5c, 0x83, 0xd8, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.splits.queries.splits.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.splits.queries.splits.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.splits.queries.splits.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "splits/queries/splits/service.proto",
}
