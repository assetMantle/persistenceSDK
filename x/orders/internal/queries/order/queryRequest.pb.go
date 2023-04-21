// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/queries/order/queryRequest.proto

package order

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	proto "github.com/gogo/protobuf/proto"
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

type QueryRequest struct {
	OrderID *base.OrderID `protobuf:"bytes,1,opt,name=order_i_d,json=orderID,proto3" json:"order_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f59362c74fc312, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetOrderID() *base.OrderID {
	if m != nil {
		return m.OrderID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "orders.queries.order.QueryRequest")
}

func init() {
	proto.RegisterFile("x/orders/internal/queries/order/queryRequest.proto", fileDescriptor_09f59362c74fc312)
}

var fileDescriptor_09f59362c74fc312 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xaa, 0xd0, 0xcf, 0x2f,
	0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x86, 0x88, 0x83, 0x79, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x22, 0x10, 0x1d, 0x7a, 0x50, 0x85, 0x7a, 0x60,
	0xae, 0x94, 0x78, 0x66, 0x4a, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0x44, 0x63, 0x7c, 0x66, 0x0a,
	0x44, 0xb9, 0x52, 0x00, 0x17, 0x4f, 0x20, 0x92, 0x21, 0x42, 0x0e, 0x5c, 0x9c, 0x50, 0x15, 0xf1,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x2a, 0x7a, 0x89, 0xc5, 0xc5, 0xa9, 0x25, 0xb9,
	0x89, 0x79, 0x25, 0x39, 0xa9, 0x7a, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0x7a, 0x99, 0x29, 0xc5,
	0x7a, 0x20, 0xf3, 0xf4, 0xfc, 0x41, 0xaa, 0x3d, 0x5d, 0x82, 0xd8, 0xf3, 0x21, 0x0c, 0xa7, 0xef,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0x25, 0x91, 0x9c, 0x9f, 0xab,
	0x87, 0xcd, 0x7d, 0x4e, 0x82, 0xc8, 0x8e, 0x08, 0x00, 0xb9, 0x2c, 0x80, 0x31, 0xca, 0x2e, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0x11, 0xe4, 0x04, 0x5f, 0xb0, 0x13,
	0xf4, 0x73, 0xf3, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0xf5, 0x09, 0x84, 0xce, 0x22, 0x26, 0x66, 0xff,
	0x40, 0xff, 0x55, 0x4c, 0x22, 0xfe, 0x10, 0x1b, 0x03, 0xa1, 0x36, 0x82, 0xb9, 0xa7, 0x60, 0xc2,
	0x31, 0x50, 0xe1, 0x18, 0x30, 0xf7, 0x11, 0x93, 0x02, 0x36, 0xe1, 0x18, 0xf7, 0x00, 0x27, 0xdf,
	0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x4c, 0x62, 0x10, 0x25, 0x56, 0x56, 0x50, 0x35,
	0x56, 0x56, 0x60, 0x81, 0x24, 0x36, 0x70, 0x90, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xca,
	0x4e, 0xd4, 0xc3, 0xb7, 0x01, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderID != nil {
		{
			size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderID != nil {
		l = m.OrderID.Size()
		n += 1 + l + sovQueryRequest(uint64(l))
	}
	return n
}

func sovQueryRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRequest(x uint64) (n int) {
	return sovQueryRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRequest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OrderID == nil {
				m.OrderID = &base.OrderID{}
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQueryRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRequest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQueryRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRequest = fmt.Errorf("proto: unexpected end of group")
)
