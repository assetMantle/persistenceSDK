// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/nub/message.proto

package nub

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
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
	From  string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	NubID *base.StringID `protobuf:"bytes,2,opt,name=nub_i_d,json=nubID,proto3" json:"nub_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f05bc956ff1d883, []int{0}
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

func (m *Message) GetNubID() *base.StringID {
	if m != nil {
		return m.NubID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.identities.transactions.nub.Message")
}

func init() {
	proto.RegisterFile("identities/transactions/nub/message.proto", fileDescriptor_2f05bc956ff1d883)
}

var fileDescriptor_2f05bc956ff1d883 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd1, 0xc1, 0x4a, 0xfb, 0x30,
	0x00, 0xc7, 0xf1, 0xa5, 0xfc, 0xf7, 0x1f, 0x56, 0x4f, 0x3d, 0x0d, 0x0f, 0x65, 0x08, 0xc2, 0xbc,
	0x24, 0xe0, 0x6e, 0x51, 0x0f, 0x1b, 0x03, 0xe9, 0xa1, 0x65, 0xcc, 0x9d, 0xa4, 0x30, 0x92, 0x26,
	0x6e, 0x81, 0x25, 0x91, 0x26, 0x01, 0x1f, 0xc3, 0x67, 0xf0, 0xe8, 0x93, 0x88, 0xa7, 0x1d, 0x3d,
	0x4a, 0x77, 0xf3, 0x15, 0xbc, 0xc8, 0xda, 0x42, 0xeb, 0x45, 0xd8, 0xad, 0x97, 0xcf, 0xb7, 0xbf,
	0x24, 0xfe, 0x85, 0x60, 0x5c, 0x59, 0x61, 0x05, 0x37, 0xc8, 0xe6, 0x44, 0x19, 0x92, 0x59, 0xa1,
	0x95, 0x41, 0xca, 0x51, 0x24, 0xb9, 0x31, 0x64, 0xc5, 0xe1, 0x63, 0xae, 0xad, 0x0e, 0x10, 0x31,
	0x86, 0x5b, 0x49, 0x94, 0xdd, 0x70, 0x28, 0x35, 0x73, 0x1b, 0x6e, 0x60, 0xc3, 0x61, 0x9b, 0x43,
	0xe5, 0xe8, 0x69, 0x5f, 0x30, 0x83, 0x28, 0x31, 0x1c, 0x19, 0x9b, 0x0b, 0xb5, 0x5a, 0x0a, 0x56,
	0xa5, 0xce, 0x52, 0xbf, 0x17, 0x57, 0xed, 0x20, 0xf0, 0xff, 0x3d, 0xe4, 0x5a, 0xf6, 0xc1, 0x00,
	0x0c, 0x8f, 0xe6, 0xe5, 0x77, 0x70, 0xe3, 0xf7, 0x94, 0xa3, 0x4b, 0xb1, 0x64, 0x7d, 0x6f, 0x00,
	0x86, 0xc7, 0x97, 0xe7, 0xb0, 0xfd, 0x6f, 0x93, 0xad, 0xb9, 0x24, 0x50, 0x30, 0x03, 0xf7, 0x75,
	0x78, 0x57, 0xd6, 0xa3, 0xe9, 0xbc, 0xab, 0x1c, 0x8d, 0xa6, 0x93, 0x6f, 0xef, 0xad, 0x08, 0xc1,
	0xb6, 0x08, 0xc1, 0x67, 0x11, 0x82, 0xe7, 0x5d, 0xd8, 0xd9, 0xee, 0xc2, 0xce, 0xc7, 0x2e, 0xec,
	0xf8, 0xa3, 0x4c, 0x4b, 0x78, 0xe0, 0x39, 0x26, 0x27, 0xf5, 0xd6, 0xd9, 0x7e, 0xfb, 0x0c, 0xdc,
	0x5f, 0xaf, 0x84, 0x5d, 0x3b, 0x0a, 0x33, 0x2d, 0xd1, 0x78, 0xdf, 0x8a, 0xcb, 0x16, 0xaa, 0x5b,
	0xe8, 0x09, 0xfd, 0x71, 0xa9, 0x2f, 0x5e, 0x77, 0x1c, 0x47, 0x8b, 0xe4, 0xd5, 0xab, 0x6c, 0xbd,
	0x23, 0xae, 0x77, 0x44, 0xcd, 0x8e, 0x45, 0x7b, 0x47, 0xe2, 0xe8, 0xfb, 0x2f, 0x91, 0xd6, 0x22,
	0x6d, 0x44, 0xda, 0x16, 0x69, 0xe2, 0x68, 0xe1, 0x5d, 0x1d, 0x28, 0xd2, 0xdb, 0xd9, 0x24, 0xe6,
	0x96, 0x30, 0x62, 0xc9, 0x97, 0x37, 0x6a, 0x69, 0x8c, 0x6b, 0x8e, 0x71, 0xe3, 0x31, 0x6e, 0x07,
	0x30, 0x4e, 0x1c, 0xa5, 0xff, 0xcb, 0x27, 0x1e, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x1a,
	0x8c, 0xb3, 0x5a, 0x02, 0x00, 0x00,
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
	if m.NubID != nil {
		{
			size, err := m.NubID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
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
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.NubID != nil {
		l = m.NubID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NubID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NubID == nil {
				m.NubID = &base.StringID{}
			}
			if err := m.NubID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
