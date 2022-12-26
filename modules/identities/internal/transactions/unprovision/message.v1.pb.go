// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/unprovision/message.v1.proto

package unprovision

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Message struct {
	From       string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To         string           `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	IdentityID *base.IdentityID `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a739e92374b8e834, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetIdentityID() *base.IdentityID {
	if m != nil {
		return m.IdentityID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "identities.transactions.unprovision.Message")
}

func init() {
	proto.RegisterFile("modules/identities/internal/transactions/unprovision/message.v1.proto", fileDescriptor_a739e92374b8e834)
}

var fileDescriptor_a739e92374b8e834 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x9d, 0x31, 0x8a, 0xa6, 0x30, 0xd8, 0x93, 0x74, 0x58, 0x24, 0x89, 0x3c, 0xcd, 0x60,
	0xdd, 0xf6, 0x96, 0x14, 0xb1, 0x07, 0x41, 0x4c, 0x3b, 0x84, 0x60, 0xe3, 0xee, 0xa4, 0x03, 0xee,
	0x8c, 0xec, 0x7b, 0x0a, 0x7d, 0x8b, 0x3e, 0x43, 0xc7, 0x3e, 0x49, 0x74, 0xf2, 0xd0, 0xa1, 0x63,
	0xac, 0xb7, 0x3e, 0x45, 0x68, 0xe9, 0xce, 0x51, 0xba, 0x3d, 0x96, 0xdf, 0xef, 0xff, 0xfe, 0xfb,
	0x86, 0x5d, 0x27, 0x36, 0x9e, 0x8e, 0x15, 0x08, 0x1d, 0x2b, 0x83, 0x1a, 0xf5, 0x72, 0x34, 0xa8,
	0x52, 0x23, 0xc7, 0x02, 0x53, 0x69, 0x40, 0x46, 0xa8, 0xad, 0x01, 0x31, 0x35, 0x93, 0xd4, 0xce,
	0x34, 0x68, 0x6b, 0x44, 0xa2, 0x00, 0xe4, 0x50, 0xf1, 0x59, 0x9d, 0x4f, 0x52, 0x8b, 0xd6, 0xab,
	0xe6, 0x3a, 0x77, 0x2d, 0xee, 0x58, 0xc7, 0x55, 0x88, 0x46, 0x2a, 0x91, 0x42, 0xc7, 0x20, 0x06,
	0x12, 0xd4, 0x7a, 0xe7, 0x53, 0x78, 0xb5, 0x49, 0x3a, 0x79, 0x60, 0x7b, 0xcd, 0xdf, 0x74, 0xcf,
	0x63, 0x3b, 0x8f, 0xa9, 0x4d, 0xca, 0xa4, 0x42, 0x6a, 0xfb, 0xed, 0xd5, 0xec, 0x95, 0x18, 0x45,
	0x5b, 0xa6, 0xab, 0x2f, 0x14, 0xad, 0x57, 0x67, 0x87, 0xeb, 0x94, 0xbe, 0xee, 0xc7, 0xe5, 0x62,
	0x85, 0xd4, 0x0e, 0xce, 0x8f, 0xb8, 0x8e, 0x81, 0x87, 0x9b, 0xf8, 0x36, 0xcb, 0x57, 0x35, 0x3e,
	0xe8, 0x5b, 0xe6, 0x93, 0x79, 0xe6, 0x93, 0xaf, 0xcc, 0x27, 0xcf, 0x0b, 0xbf, 0x30, 0x5f, 0xf8,
	0x85, 0xcf, 0x85, 0x5f, 0x60, 0x67, 0x91, 0x4d, 0xf8, 0x16, 0xbf, 0xd2, 0x28, 0xfd, 0x75, 0xbc,
	0xab, 0xb7, 0x96, 0xad, 0x5b, 0xe4, 0xfe, 0x76, 0xa8, 0x71, 0x34, 0x1d, 0xf0, 0xc8, 0x26, 0xe2,
	0x12, 0x40, 0x61, 0x53, 0x1a, 0x1c, 0x2b, 0xb1, 0xbe, 0xef, 0x7f, 0xee, 0xfc, 0x42, 0x8b, 0x61,
	0xa7, 0xfb, 0x4a, 0xab, 0x61, 0x5e, 0xa9, 0xe3, 0x56, 0xea, 0xe6, 0xec, 0xbb, 0x4b, 0xf5, 0x5c,
	0xaa, 0xe7, 0x50, 0x19, 0x15, 0x5b, 0x50, 0xbd, 0x9b, 0x56, 0xa3, 0xa9, 0x50, 0xc6, 0x12, 0xe5,
	0x37, 0x3d, 0xcd, 0x8d, 0x20, 0x70, 0x95, 0x20, 0x70, 0x9c, 0xc1, 0xee, 0xea, 0xfd, 0x2e, 0x7e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x22, 0xaa, 0xea, 0x52, 0x02, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IdentityID != nil {
		{
			size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintMessageV1(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessageV1(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessageV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessageV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessageV1(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	return n
}

func sovMessageV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessageV1(x uint64) (n int) {
	return sovMessageV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageV1
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageV1
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
				return ErrInvalidLengthMessageV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IdentityID == nil {
				m.IdentityID = &base.IdentityID{}
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageV1
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
func skipMessageV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageV1
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
					return 0, ErrIntOverflowMessageV1
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
					return 0, ErrIntOverflowMessageV1
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
				return 0, ErrInvalidLengthMessageV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessageV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessageV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessageV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessageV1 = fmt.Errorf("proto: unexpected end of group")
)
