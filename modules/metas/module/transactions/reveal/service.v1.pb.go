// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/metas/internal/transactions/reveal/service.v1.proto

package reveal

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
	proto.RegisterFile("modules/metas/internal/transactions/reveal/service.v1.proto", fileDescriptor_abbde032fd385abe)
}

var fileDescriptor_abbde032fd385abe = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0xb3, 0xfb, 0x41, 0x8a, 0xfd, 0xc0, 0x22, 0x88, 0x90, 0x35, 0x6e, 0x11, 0xb1, 0x9d,
	0x21, 0xda, 0xad, 0x36, 0xa6, 0x51, 0x8b, 0x85, 0x10, 0x83, 0x85, 0xa4, 0xb9, 0xd9, 0x5c, 0xd6,
	0x85, 0xf9, 0x13, 0x66, 0x26, 0x29, 0x2d, 0x7c, 0x02, 0xc1, 0xc6, 0xda, 0xd2, 0x27, 0x11, 0xab,
	0x80, 0x8d, 0xa5, 0x6c, 0xac, 0x7c, 0x0a, 0xc9, 0xcc, 0x4a, 0xb6, 0x89, 0x98, 0xfa, 0xfc, 0x0e,
	0xbf, 0xb9, 0x67, 0x82, 0x63, 0x2e, 0xc7, 0x53, 0x86, 0x9a, 0x72, 0x34, 0xa0, 0x69, 0x2e, 0x0c,
	0x2a, 0x01, 0x8c, 0x1a, 0x05, 0x42, 0x43, 0x6a, 0x72, 0x29, 0x34, 0x55, 0x38, 0x43, 0x60, 0x54,
	0xa3, 0x9a, 0xe5, 0x29, 0x92, 0x59, 0x87, 0x4c, 0x94, 0x34, 0xb2, 0xd1, 0xb4, 0x25, 0x52, 0x65,
	0x89, 0x63, 0xc3, 0x56, 0x26, 0x65, 0xc6, 0x90, 0xc2, 0x24, 0xa7, 0x20, 0x84, 0x34, 0xe0, 0x62,
	0x5b, 0x0c, 0x37, 0xb1, 0x72, 0xd4, 0x1a, 0xb2, 0x95, 0x35, 0x3c, 0xd9, 0xa0, 0xac, 0x50, 0x4f,
	0xa4, 0xd0, 0xab, 0xf6, 0xe1, 0x6d, 0xf0, 0x7f, 0xb0, 0x02, 0x1b, 0x32, 0xa8, 0x9f, 0x83, 0x18,
	0x33, 0x6c, 0xb4, 0xc9, 0xda, 0x6b, 0x48, 0xe2, 0xde, 0x10, 0xee, 0xff, 0xc2, 0xf4, 0x4b, 0x55,
	0xbb, 0x75, 0xf7, 0xf6, 0xf9, 0xe0, 0xef, 0xb4, 0xb7, 0x29, 0x07, 0x61, 0x18, 0x96, 0xef, 0x74,
	0x58, 0xf7, 0xd1, 0x7f, 0x29, 0x22, 0x6f, 0x5e, 0x44, 0xde, 0x47, 0x11, 0x79, 0xf7, 0x8b, 0xa8,
	0x36, 0x5f, 0x44, 0xb5, 0xf7, 0x45, 0x54, 0x0b, 0xf6, 0x52, 0xc9, 0xd7, 0x0b, 0xba, 0x5b, 0x97,
	0x6e, 0xff, 0xab, 0x4e, 0x6f, 0x79, 0x49, 0xcf, 0xbb, 0xbe, 0xc8, 0x72, 0x73, 0x33, 0x1d, 0x91,
	0x54, 0x72, 0x7a, 0xaa, 0x35, 0x9a, 0xa4, 0x14, 0xfe, 0x0c, 0xf4, 0xe7, 0xa1, 0x9e, 0xfc, 0x7f,
	0xc9, 0xa0, 0xff, 0xec, 0x37, 0x13, 0xab, 0x1f, 0x54, 0xf5, 0x7d, 0x4b, 0xbc, 0x96, 0xd9, 0xb0,
	0x9a, 0x0d, 0x5d, 0x56, 0xf8, 0x07, 0x6b, 0xb3, 0xe1, 0x59, 0xaf, 0xbb, 0x0c, 0xc7, 0x60, 0xe0,
	0xcb, 0xdf, 0xb5, 0x5c, 0x1c, 0x57, 0xc1, 0x38, 0x76, 0xe4, 0xa8, 0x6e, 0x7f, 0xe8, 0xe8, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0xf7, 0x91, 0x47, 0x94, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type transactionClient struct {
	cc grpc1.ClientConn
}

func NewTransactionClient(cc grpc1.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/metas.transactions.reveal.Transaction/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	Handle(context.Context, *Message) (*Response, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) Handle(ctx context.Context, req *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterTransactionServer(s grpc1.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metas.transactions.reveal.Transaction/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metas.transactions.reveal.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Transaction_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/metas/internal/transactions/reveal/service.v1.proto",
}
