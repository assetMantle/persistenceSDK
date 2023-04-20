// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/mappable/mappable.proto

package mappable

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/documents/base"
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

type Mappable struct {
	Order *base.Document `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *Mappable) Reset()         { *m = Mappable{} }
func (m *Mappable) String() string { return proto.CompactTextString(m) }
func (*Mappable) ProtoMessage()    {}
func (*Mappable) Descriptor() ([]byte, []int) {
	return fileDescriptor_e23c70c28f92eecc, []int{0}
}
func (m *Mappable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mappable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mappable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mappable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mappable.Merge(m, src)
}
func (m *Mappable) XXX_Size() int {
	return m.Size()
}
func (m *Mappable) XXX_DiscardUnknown() {
	xxx_messageInfo_Mappable.DiscardUnknown(m)
}

var xxx_messageInfo_Mappable proto.InternalMessageInfo

func (m *Mappable) GetOrder() *base.Document {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterType((*Mappable)(nil), "orders.Mappable")
}

func init() {
	proto.RegisterFile("x/orders/internal/mappable/mappable.proto", fileDescriptor_e23c70c28f92eecc)
}

var fileDescriptor_e23c70c28f92eecc = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xac, 0xd0, 0xcf, 0x2f,
	0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d,
	0x2c, 0x28, 0x48, 0x4c, 0xca, 0x49, 0x85, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8,
	0x20, 0x0a, 0xa5, 0x64, 0x53, 0xf2, 0x93, 0x4b, 0x73, 0x53, 0xf3, 0x4a, 0x8a, 0xf5, 0x93, 0x12,
	0x8b, 0x53, 0xf5, 0x61, 0x5c, 0x88, 0x32, 0x25, 0x2b, 0x2e, 0x0e, 0x5f, 0xa8, 0x46, 0x21, 0x3d,
	0x2e, 0x56, 0xb0, 0x26, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x09, 0x3d, 0xb8, 0x56, 0x3d,
	0x90, 0x56, 0x3d, 0x17, 0x28, 0x37, 0x08, 0xa2, 0xcc, 0x69, 0x19, 0xe3, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x71, 0x25, 0xe7, 0xe7, 0xea, 0x41, 0x5c, 0xe0, 0xc4, 0x0b,
	0xb3, 0x20, 0x00, 0x64, 0x63, 0x00, 0x63, 0x94, 0x65, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x09, 0xc8, 0x07, 0xf9,
	0x29, 0xa5, 0x39, 0xa9, 0xc5, 0xfa, 0xb8, 0x3d, 0xb9, 0x88, 0x89, 0xd9, 0x3f, 0x22, 0x62, 0x15,
	0x13, 0x9b, 0x3f, 0x58, 0xc1, 0x29, 0x18, 0xe3, 0x11, 0x93, 0x10, 0x84, 0x11, 0xe3, 0x1e, 0xe0,
	0xe4, 0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58, 0x92, 0xf8, 0x0a, 0x26, 0x9b, 0xc4, 0x06, 0xf6, 0xab,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x07, 0xcb, 0x37, 0xec, 0x3f, 0x01, 0x00, 0x00,
}

func (m *Mappable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mappable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mappable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMappable(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMappable(dAtA []byte, offset int, v uint64) int {
	offset -= sovMappable(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mappable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovMappable(uint64(l))
	}
	return n
}

func sovMappable(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMappable(x uint64) (n int) {
	return sovMappable(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mappable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMappable
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
			return fmt.Errorf("proto: Mappable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mappable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMappable
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
				return ErrInvalidLengthMappable
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMappable
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &base.Document{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMappable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMappable
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
func skipMappable(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMappable
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
					return 0, ErrIntOverflowMappable
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
					return 0, ErrIntOverflowMappable
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
				return 0, ErrInvalidLengthMappable
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMappable
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMappable
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMappable        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMappable          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMappable = fmt.Errorf("proto: unexpected end of group")
)
