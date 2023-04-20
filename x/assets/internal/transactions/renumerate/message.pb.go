// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/transactions/renumerate/message.proto

package renumerate

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

type Message struct {
	From    string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID  *base.IdentityID `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	AssetID *base.AssetID    `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5ac258a2202c4d1, []int{0}
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

func (m *Message) GetFromID() *base.IdentityID {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *Message) GetAssetID() *base.AssetID {
	if m != nil {
		return m.AssetID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assets.transactions.renumerate.Message")
}

func init() {
	proto.RegisterFile("x/assets/internal/transactions/renumerate/message.proto", fileDescriptor_d5ac258a2202c4d1)
}

var fileDescriptor_d5ac258a2202c4d1 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xc7, 0x7b, 0xe9, 0x8f, 0xf6, 0xd7, 0xe8, 0x62, 0x10, 0x2c, 0x1d, 0x8e, 0xd2, 0xa9, 0x08,
	0xde, 0x81, 0x0e, 0x42, 0xb6, 0x96, 0x82, 0x66, 0x28, 0x94, 0xd0, 0x49, 0x0a, 0xe1, 0xda, 0x3b,
	0xeb, 0x41, 0x73, 0x27, 0xb9, 0xa7, 0xa0, 0xe0, 0x8b, 0xf0, 0x35, 0x38, 0x3a, 0xf8, 0x3a, 0xc4,
	0xa9, 0xa3, 0xa3, 0xa4, 0x9b, 0xaf, 0x42, 0x92, 0xab, 0x49, 0xa6, 0xe2, 0x94, 0x07, 0x9e, 0xcf,
	0xf7, 0xcf, 0x93, 0x73, 0x2f, 0x1f, 0x28, 0x33, 0x46, 0x80, 0xa1, 0x52, 0x81, 0x48, 0x14, 0x5b,
	0x51, 0x48, 0x98, 0x32, 0x6c, 0x01, 0x52, 0x2b, 0x43, 0x13, 0xa1, 0xd6, 0xb1, 0x48, 0x18, 0x08,
	0x1a, 0x0b, 0x63, 0xd8, 0x52, 0x90, 0xfb, 0x44, 0x83, 0xf6, 0xb0, 0x95, 0x91, 0x2a, 0x4d, 0x4a,
	0xba, 0x73, 0x22, 0xb9, 0xa1, 0x73, 0x66, 0x84, 0xf5, 0x8f, 0x24, 0xb7, 0xc2, 0x4e, 0xa7, 0x58,
	0x48, 0x2e, 0x14, 0x48, 0x78, 0x2c, 0x76, 0xbd, 0x27, 0xb7, 0x39, 0xb6, 0x29, 0x9e, 0xe7, 0xfe,
	0xbb, 0x4d, 0x74, 0xdc, 0x46, 0x5d, 0xd4, 0x6f, 0x85, 0xf9, 0xec, 0x11, 0xf7, 0x7f, 0xf6, 0x8d,
	0x64, 0xc4, 0xdb, 0x4e, 0x17, 0xf5, 0x0f, 0xce, 0x8f, 0x89, 0xe4, 0x86, 0x64, 0x6e, 0x24, 0xd8,
	0xb9, 0x05, 0xa3, 0xb0, 0x91, 0x51, 0xc1, 0xc8, 0x3b, 0x73, 0x5b, 0xbb, 0xf0, 0x88, 0xb7, 0xeb,
	0xb9, 0xe0, 0xa8, 0x14, 0x0c, 0xb2, 0x55, 0x30, 0x0a, 0x9b, 0xcc, 0x0e, 0xc3, 0x37, 0xe7, 0x3d,
	0xc5, 0x68, 0x93, 0x62, 0xf4, 0x95, 0x62, 0xf4, 0xbc, 0xc5, 0xb5, 0xcd, 0x16, 0xd7, 0x3e, 0xb7,
	0xb8, 0xe6, 0xf6, 0x16, 0x3a, 0x26, 0xfb, 0x2f, 0x1e, 0x1e, 0xee, 0xaa, 0x4f, 0xb2, 0x53, 0x26,
	0xe8, 0xe6, 0x7a, 0x29, 0xe1, 0x6e, 0x3d, 0x27, 0x0b, 0x1d, 0xd3, 0x3c, 0x6b, 0xcc, 0x14, 0xac,
	0x04, 0x8d, 0x35, 0x5f, 0xaf, 0x84, 0xa1, 0x7f, 0xfe, 0xf3, 0x2f, 0x4e, 0x7d, 0x30, 0x0d, 0x5f,
	0x1d, 0x3c, 0xb0, 0x05, 0xa6, 0xd5, 0x02, 0x61, 0x81, 0x7d, 0xfc, 0x02, 0xb3, 0x2a, 0x30, 0x2b,
	0x81, 0xd4, 0x39, 0xdd, 0x0f, 0xcc, 0xae, 0x26, 0xc3, 0xb1, 0x00, 0xc6, 0x19, 0xb0, 0x6f, 0xa7,
	0x6b, 0x61, 0xdf, 0xaf, 0xd2, 0xbe, 0x5f, 0xe2, 0xf3, 0x46, 0xfe, 0x6a, 0x17, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x80, 0x6b, 0xfb, 0x5d, 0x45, 0x02, 0x00, 0x00,
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
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID != nil {
		{
			size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.AssetID != nil {
		l = m.AssetID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
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
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
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
			if m.AssetID == nil {
				m.AssetID = &base.AssetID{}
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
