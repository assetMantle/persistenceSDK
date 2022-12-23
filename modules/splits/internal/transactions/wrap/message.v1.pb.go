// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/transactions/wrap/message.v1.proto

package wrap

import (
	fmt "fmt"
	v1beta1 "github.com/AssetMantle/modules/cosmos/base/v1beta1"
	base "github.com/AssetMantle/modules/schema/ids/base"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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
	From   string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID *base.IdentityID                         `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	Coins  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c6420e4543415e9, []int{0}
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

func (m *Message) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "splits.transactions.wrap.Message")
}

func init() {
	proto.RegisterFile("modules/splits/internal/transactions/wrap/message.v1.proto", fileDescriptor_7c6420e4543415e9)
}

var fileDescriptor_7c6420e4543415e9 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0xa4, 0xba, 0x6a, 0x16, 0x14, 0x82, 0x87, 0x58, 0x24, 0x5b, 0xd4, 0x43, 0x3c, 0x38,
	0x63, 0xd6, 0x5b, 0x6e, 0xc6, 0x05, 0xed, 0xa1, 0x50, 0xba, 0x8b, 0x0b, 0x52, 0x58, 0x26, 0xc9,
	0x98, 0x1d, 0x4c, 0x66, 0x42, 0xde, 0x6c, 0xa5, 0xff, 0x42, 0xfc, 0x09, 0x1e, 0xeb, 0x1f, 0x29,
	0x9e, 0x7a, 0xf4, 0xa4, 0x92, 0xde, 0xfc, 0x15, 0x32, 0x99, 0xb4, 0xe4, 0x52, 0xf0, 0xf4, 0x1e,
	0xf3, 0x7d, 0xdf, 0xfb, 0xbe, 0x37, 0xcf, 0x89, 0x4a, 0x99, 0xdd, 0x14, 0x0c, 0x08, 0x54, 0x05,
	0x57, 0x40, 0xb8, 0x50, 0xac, 0x16, 0xb4, 0x20, 0xaa, 0xa6, 0x02, 0x68, 0xaa, 0xb8, 0x14, 0x40,
	0x3e, 0xd7, 0xb4, 0x22, 0x25, 0x03, 0xa0, 0x39, 0xc3, 0x8b, 0x10, 0x57, 0xb5, 0x54, 0xd2, 0xf5,
	0x8c, 0x06, 0xf7, 0xa9, 0x58, 0x53, 0x87, 0x7e, 0x2a, 0xa1, 0x94, 0x40, 0x12, 0x0a, 0x8c, 0x2c,
	0xc2, 0x84, 0x29, 0x1a, 0x92, 0x54, 0x72, 0x61, 0x94, 0xc3, 0x87, 0xb9, 0xcc, 0x65, 0xdb, 0x12,
	0xdd, 0x75, 0xaf, 0x4f, 0x21, 0xbd, 0x66, 0x25, 0x25, 0x3c, 0xeb, 0x94, 0x3c, 0x63, 0x42, 0x71,
	0xb5, 0x1c, 0x9f, 0xed, 0x4d, 0x9f, 0x7c, 0x47, 0xce, 0x9d, 0x89, 0x49, 0xe2, 0xba, 0xce, 0xad,
	0x8f, 0xb5, 0x2c, 0x3d, 0x34, 0x42, 0xc1, 0xbd, 0x59, 0xdb, 0xbb, 0xcf, 0x9d, 0xbb, 0xba, 0x5e,
	0xf1, 0xab, 0xcc, 0xb3, 0x47, 0x28, 0x38, 0x3e, 0x7d, 0x80, 0x79, 0x06, 0x78, 0xbc, 0x9f, 0x35,
	0x3b, 0xd2, 0x84, 0xf1, 0x99, 0x4b, 0x9d, 0xdb, 0x3a, 0x13, 0x78, 0x83, 0xd1, 0x20, 0x38, 0x3e,
	0x7d, 0x84, 0x4d, 0x6a, 0xac, 0xbd, 0x71, 0x97, 0x1a, 0xbf, 0x91, 0x5c, 0xc4, 0x2f, 0xd7, 0xbf,
	0x4e, 0xac, 0xd5, 0xef, 0x93, 0x20, 0xe7, 0xea, 0xfa, 0x26, 0xc1, 0xa9, 0x2c, 0x49, 0xb7, 0xa2,
	0x29, 0x2f, 0x20, 0xfb, 0x44, 0xd4, 0xb2, 0x62, 0xd0, 0x0a, 0x60, 0x66, 0x26, 0xc7, 0x5f, 0xed,
	0x75, 0xe3, 0xa3, 0x4d, 0xe3, 0xa3, 0x3f, 0x8d, 0x8f, 0xbe, 0x6c, 0x7d, 0x6b, 0xb3, 0xf5, 0xad,
	0x9f, 0x5b, 0xdf, 0x72, 0x1e, 0xa7, 0xb2, 0xc4, 0x87, 0x7e, 0x30, 0xbe, 0xdf, 0xed, 0xf8, 0x3e,
	0x9c, 0xea, 0xb5, 0xa7, 0xe8, 0xc3, 0xbb, 0x9e, 0xf7, 0x6b, 0x00, 0xa6, 0x26, 0x54, 0xa8, 0x82,
	0x91, 0xdd, 0x01, 0xff, 0xfb, 0x90, 0xdf, 0xec, 0xc1, 0xf9, 0xc5, 0xe5, 0xca, 0xf6, 0xce, 0x8d,
	0xf9, 0x45, 0xdf, 0xfc, 0xb2, 0xa6, 0xd5, 0x8f, 0x1d, 0x34, 0xef, 0x43, 0x73, 0x0d, 0x35, 0xf6,
	0xb3, 0x43, 0xd0, 0xfc, 0xed, 0x34, 0x9e, 0x30, 0x45, 0x33, 0xaa, 0xe8, 0x5f, 0x7b, 0x68, 0x68,
	0x51, 0xd4, 0xe7, 0x45, 0x91, 0x26, 0x26, 0x47, 0xed, 0x25, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x00, 0x9f, 0xad, 0xa7, 0x7c, 0x02, 0x00, 0x00,
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
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessageV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.FromID != nil {
		{
			size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
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
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovMessageV1(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
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
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			m.Coins = append(m.Coins, v1beta1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
