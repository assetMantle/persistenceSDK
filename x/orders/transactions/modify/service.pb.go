// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/orders/transactions/modify/service.proto

package modify

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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
	proto.RegisterFile("AssetMantle/modules/x/orders/transactions/modify/service.proto", fileDescriptor_8ff4398d92b669f7)
}

var fileDescriptor_8ff4398d92b669f7 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0xbf, 0x4b, 0xf3, 0x40,
	0x18, 0x07, 0xf0, 0x24, 0xe5, 0xed, 0x10, 0xde, 0xa9, 0xcb, 0x0b, 0x19, 0x32, 0xbc, 0xa3, 0xc3,
	0x9d, 0x3f, 0x07, 0x4f, 0x11, 0x5a, 0x10, 0x05, 0x39, 0x5a, 0xb4, 0x43, 0x90, 0x80, 0x5c, 0x93,
	0x6b, 0x0c, 0x34, 0xb9, 0x92, 0xe7, 0x5a, 0xda, 0x4d, 0x77, 0x07, 0x77, 0x37, 0x47, 0x27, 0xff,
	0x0c, 0x71, 0xea, 0xe8, 0x28, 0xe9, 0x20, 0xf8, 0x57, 0x48, 0x7b, 0x07, 0x5e, 0xc7, 0x9b, 0x02,
	0x97, 0x7c, 0xbe, 0xf9, 0x3e, 0xc9, 0xe3, 0x9f, 0xb4, 0x01, 0xb8, 0xa4, 0xac, 0x94, 0x23, 0x8e,
	0x0b, 0x91, 0x4e, 0x46, 0x1c, 0xf0, 0x0c, 0x8b, 0x2a, 0xe5, 0x15, 0x60, 0x59, 0xb1, 0x12, 0x58,
	0x22, 0x73, 0x51, 0xc2, 0xea, 0x6e, 0x3e, 0x9c, 0x63, 0xe0, 0xd5, 0x34, 0x4f, 0x38, 0x1a, 0x57,
	0x42, 0x8a, 0xd6, 0xb6, 0xe1, 0x91, 0xf6, 0x68, 0x86, 0x94, 0x47, 0xa6, 0x47, 0xca, 0x07, 0xff,
	0x12, 0x01, 0x85, 0x00, 0x5c, 0x40, 0x86, 0xa7, 0x3b, 0xab, 0x8b, 0x8a, 0x0a, 0xec, 0xab, 0x14,
	0x1c, 0x80, 0x65, 0xba, 0x4a, 0x70, 0x61, 0xed, 0x8d, 0xb3, 0x9b, 0x8a, 0xc3, 0x58, 0x94, 0xa0,
	0xc3, 0x76, 0x9f, 0x5c, 0xbf, 0x41, 0x21, 0x6b, 0x3d, 0xb8, 0x7e, 0xf3, 0x9c, 0x95, 0xe9, 0x88,
	0xb7, 0x0e, 0x91, 0xed, 0xac, 0x88, 0xaa, 0x82, 0xc1, 0xa9, 0x3d, 0xed, 0xff, 0x9e, 0x5d, 0xea,
	0x6a, 0xff, 0x9d, 0xe0, 0xcf, 0xdd, 0xd7, 0xeb, 0x96, 0xdb, 0xb9, 0x6f, 0xbc, 0xd5, 0xa1, 0xbb,
	0xa8, 0x43, 0xf7, 0xb3, 0x0e, 0xdd, 0xc7, 0x65, 0xe8, 0x2c, 0x96, 0xa1, 0xf3, 0xb1, 0x0c, 0x1d,
	0x7f, 0x3f, 0x11, 0x85, 0xf5, 0xdb, 0x3a, 0x7f, 0xaf, 0xd4, 0x5f, 0xed, 0xad, 0x86, 0xef, 0xb9,
	0xd7, 0x47, 0x59, 0x2e, 0x6f, 0x27, 0x03, 0x94, 0x88, 0x02, 0xdb, 0x7e, 0xd6, 0x67, 0xaf, 0xd9,
	0xa6, 0x51, 0xb7, 0x4f, 0x5f, 0xbc, 0x8d, 0xed, 0xa0, 0xba, 0x48, 0x84, 0xba, 0xaa, 0x48, 0xdf,
	0x2c, 0x42, 0xd7, 0xf6, 0x7d, 0x83, 0xc4, 0x9a, 0xc4, 0x51, 0xac, 0x48, 0x6c, 0x92, 0x58, 0x91,
	0xda, 0x3b, 0xb6, 0x25, 0xf1, 0x59, 0xaf, 0x43, 0xb9, 0x64, 0x29, 0x93, 0xec, 0xdb, 0x3b, 0x30,
	0x38, 0x21, 0xda, 0x13, 0x12, 0x11, 0xa2, 0x12, 0x08, 0x31, 0x23, 0xd6, 0x4f, 0xe4, 0xc3, 0xf9,
	0xa0, 0xb9, 0x5e, 0x94, 0xbd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x98, 0x4d, 0x5b, 0x42,
	0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.orders.transactions.modify.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.orders.transactions.modify.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.orders.transactions.modify.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/orders/transactions/modify/service.proto",
}
