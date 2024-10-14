// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/define/service.proto

package define

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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/define/service.proto", fileDescriptor_0c544b9b622019f8)
}

var fileDescriptor_0c544b9b622019f8 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0xc0, 0x93, 0x14, 0x3b, 0x04, 0xa7, 0x2e, 0x42, 0x86, 0x0c, 0x8e, 0x0e, 0x77, 0xf8, 0x67,
	0xd0, 0x03, 0xd1, 0x86, 0x82, 0x76, 0x08, 0x16, 0xed, 0x10, 0x24, 0x20, 0xd7, 0xe4, 0x8c, 0x81,
	0xe6, 0xae, 0xe4, 0x5d, 0x4b, 0x47, 0x3f, 0x82, 0x83, 0x8b, 0xab, 0x6e, 0x4e, 0x7e, 0x0c, 0x71,
	0xea, 0xe8, 0x28, 0xe9, 0x20, 0xf8, 0x29, 0xa4, 0xcd, 0x41, 0xae, 0xeb, 0x4d, 0x81, 0x07, 0xbf,
	0x5f, 0x7e, 0x79, 0x79, 0x6e, 0xd0, 0x05, 0x60, 0x32, 0xa4, 0x5c, 0x8e, 0x19, 0x2e, 0x44, 0x3a,
	0x1d, 0x33, 0xc0, 0x73, 0x9c, 0xa7, 0x8c, 0xcb, 0x5c, 0xe6, 0x0c, 0xb0, 0x2c, 0x29, 0x07, 0x9a,
	0xc8, 0x5c, 0x70, 0xc0, 0x29, 0xbb, 0xcf, 0x39, 0xc3, 0xc0, 0xca, 0x59, 0x9e, 0x30, 0x34, 0x29,
	0x85, 0x14, 0x9d, 0x23, 0xcd, 0x81, 0x94, 0x03, 0xcd, 0x51, 0xe3, 0x40, 0xba, 0x03, 0xd5, 0x0e,
	0x6f, 0x27, 0x11, 0x50, 0x08, 0xc0, 0x05, 0x64, 0x78, 0xb6, 0xbf, 0x7a, 0xd4, 0x3a, 0xcf, 0x2c,
	0xa9, 0x60, 0x00, 0x34, 0x53, 0x49, 0xde, 0x95, 0x91, 0x43, 0x9b, 0xdd, 0x95, 0x0c, 0x26, 0x82,
	0x83, 0x12, 0x1e, 0xbc, 0xd9, 0x6e, 0x2b, 0x84, 0xac, 0xf3, 0x6c, 0xbb, 0xed, 0x4b, 0xca, 0xd3,
	0x31, 0xeb, 0x9c, 0x22, 0x93, 0xef, 0x46, 0x61, 0x1d, 0xea, 0xf5, 0xcd, 0xf0, 0x61, 0x33, 0xbb,
	0x56, 0x89, 0xbb, 0x96, 0xb7, 0xf5, 0xf8, 0xfb, 0xb1, 0x67, 0x07, 0x2f, 0xad, 0xcf, 0xca, 0xb7,
	0x17, 0x95, 0x6f, 0xff, 0x54, 0xbe, 0xfd, 0xb4, 0xf4, 0xad, 0xc5, 0xd2, 0xb7, 0xbe, 0x97, 0xbe,
	0xe5, 0x1e, 0x27, 0xa2, 0x30, 0x7a, 0x63, 0xb0, 0x7d, 0x53, 0xff, 0xed, 0xc1, 0x6a, 0x11, 0x03,
	0xfb, 0xf6, 0x2c, 0xcb, 0xe5, 0xc3, 0x74, 0x84, 0x12, 0x51, 0x60, 0x93, 0x35, 0xbf, 0x3a, 0xed,
	0x6e, 0x18, 0xf5, 0x87, 0xbd, 0x77, 0x67, 0xe3, 0x72, 0x42, 0x15, 0x14, 0xa1, 0x7e, 0x13, 0x34,
	0xd4, 0x83, 0x7a, 0x6b, 0xfe, 0x6b, 0x03, 0x8b, 0x15, 0x16, 0x47, 0x71, 0x83, 0xc5, 0x3a, 0x16,
	0xd7, 0x58, 0xe5, 0x9c, 0x9b, 0x60, 0xf1, 0xc5, 0x20, 0x08, 0x99, 0xa4, 0x29, 0x95, 0xf4, 0xcf,
	0x39, 0xd1, 0x14, 0x84, 0x28, 0x07, 0x21, 0x11, 0x21, 0x8d, 0x85, 0x10, 0x5d, 0x43, 0x48, 0xed,
	0x19, 0xb5, 0xd7, 0x87, 0x74, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x61, 0x79, 0xb1, 0x72,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.define.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.define.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.define.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/define/service.proto",
}
