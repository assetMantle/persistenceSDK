// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/mint/service.proto

package mint

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
	proto.RegisterFile("assets/transactions/mint/service.proto", fileDescriptor_c3e3f73367fe7997)
}

var fileDescriptor_c3e3f73367fe7997 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0x80, 0x9b, 0x7c, 0x7c, 0x1d, 0x82, 0x53, 0xc5, 0x25, 0x4a, 0x86, 0x0e, 0x4e, 0xe5, 0xae,
	0x58, 0x04, 0xbd, 0x2d, 0x5d, 0x74, 0x39, 0x28, 0xda, 0x49, 0x02, 0x72, 0x4d, 0x8f, 0x18, 0x48,
	0xee, 0x4a, 0xde, 0xab, 0x38, 0xfb, 0x0b, 0x04, 0x7f, 0x80, 0xe0, 0x22, 0xf8, 0x4b, 0xc4, 0xa9,
	0xe0, 0xe2, 0x28, 0xa9, 0x8b, 0xfe, 0x0a, 0xb9, 0xde, 0x81, 0xe7, 0x90, 0x21, 0xeb, 0xe5, 0x79,
	0xde, 0xf7, 0xc9, 0x25, 0xc1, 0x3e, 0x03, 0xe0, 0x0a, 0xb0, 0xaa, 0x98, 0x00, 0x96, 0xaa, 0x5c,
	0x0a, 0xc0, 0x65, 0x2e, 0x14, 0x06, 0x5e, 0x5d, 0xe7, 0x29, 0x47, 0x8b, 0x4a, 0x2a, 0xd9, 0x1b,
	0x6c, 0xb8, 0x92, 0x09, 0x55, 0x70, 0x54, 0xca, 0xf9, 0xb2, 0xe0, 0x80, 0x8c, 0x8b, 0x5c, 0x17,
	0x69, 0x37, 0xdc, 0xcb, 0xa4, 0xcc, 0x0a, 0x8e, 0xd9, 0x22, 0xc7, 0x4c, 0x08, 0xa9, 0x98, 0x79,
	0xb8, 0x99, 0x15, 0x36, 0xef, 0x2c, 0x39, 0x00, 0xcb, 0xec, 0xce, 0x70, 0xd4, 0xc8, 0x39, 0x27,
	0x97, 0x15, 0x87, 0x85, 0x14, 0x60, 0xa5, 0x83, 0x27, 0x2f, 0xf8, 0x47, 0x21, 0xeb, 0x3d, 0x78,
	0x41, 0xf7, 0x94, 0x89, 0x79, 0xc1, 0x7b, 0x87, 0xa8, 0x4d, 0x3c, 0xa2, 0x26, 0x22, 0x8c, 0xdb,
	0x69, 0xd3, 0xdf, 0x93, 0x33, 0x9b, 0xd4, 0xdf, 0xbd, 0x7d, 0xfb, 0xbc, 0xf7, 0x77, 0xfa, 0xdb,
	0xd8, 0x4c, 0xc1, 0xf6, 0x8d, 0xb4, 0x30, 0xfe, 0xf2, 0x5f, 0xea, 0xc8, 0x5b, 0xd5, 0x91, 0xf7,
	0x51, 0x47, 0xde, 0xdd, 0x3a, 0xea, 0xac, 0xd6, 0x51, 0xe7, 0x7d, 0x1d, 0x75, 0x82, 0x61, 0x2a,
	0xcb, 0x56, 0xdb, 0xc7, 0x5b, 0xe7, 0xe6, 0x73, 0x4d, 0xf4, 0x25, 0x4c, 0xbc, 0x8b, 0xe3, 0x2c,
	0x57, 0x57, 0xcb, 0x19, 0x4a, 0x65, 0x89, 0x63, 0x2d, 0x51, 0x13, 0x60, 0x07, 0xe1, 0x1b, 0xdc,
	0x74, 0xb9, 0x8f, 0xfe, 0xff, 0x98, 0xc6, 0x53, 0xfa, 0xec, 0x0f, 0x62, 0xa7, 0x80, 0xda, 0x82,
	0xd8, 0x14, 0x4c, 0xdd, 0x02, 0x9a, 0x0b, 0xf5, 0xfa, 0x07, 0x4f, 0x2c, 0x9e, 0x18, 0x3c, 0x71,
	0xf1, 0x44, 0xe3, 0xb5, 0x7f, 0xd4, 0x06, 0x4f, 0x4e, 0x26, 0x63, 0xca, 0x15, 0x9b, 0x33, 0xc5,
	0xbe, 0xfd, 0xa1, 0xa3, 0x12, 0x62, 0x5d, 0x42, 0x8c, 0x4c, 0x88, 0x6b, 0x13, 0xa2, 0xf5, 0x59,
	0x77, 0xf3, 0x73, 0x8c, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0x55, 0x24, 0x0b, 0xef, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.transactions.mint.Msg/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.assets.transactions.mint.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.transactions.mint.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/transactions/mint/service.proto",
}
