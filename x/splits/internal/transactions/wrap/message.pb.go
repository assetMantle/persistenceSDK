// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/wrap/message.proto

package wrap

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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
	From   string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID *base.IdentityID                         `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	Coins  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb9dd6ec524ca386, []int{0}
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
	proto.RegisterFile("x/splits/internal/transactions/wrap/message.proto", fileDescriptor_fb9dd6ec524ca386)
}

var fileDescriptor_fb9dd6ec524ca386 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0xcd, 0xa4, 0xfa, 0xd4, 0x3c, 0x41, 0x08, 0x2e, 0x62, 0x95, 0xbc, 0x22, 0x2e, 0xe2, 0xc2,
	0x19, 0xfb, 0xdc, 0x65, 0x67, 0x2c, 0x48, 0x17, 0x85, 0xd2, 0x16, 0x0a, 0x52, 0x28, 0x93, 0x64,
	0x8c, 0x83, 0xc9, 0x4c, 0xc8, 0x9d, 0x6a, 0xfb, 0x0b, 0xae, 0xfc, 0x06, 0x97, 0xf5, 0x47, 0x8a,
	0xab, 0x2e, 0x5d, 0xa9, 0xa4, 0x3b, 0xbf, 0x42, 0x26, 0x13, 0x25, 0x9b, 0x82, 0xab, 0x7b, 0x99,
	0x73, 0xce, 0x3d, 0xf7, 0xcc, 0x75, 0x86, 0x5b, 0x02, 0x65, 0xce, 0x15, 0x10, 0x2e, 0x14, 0xab,
	0x04, 0xcd, 0x89, 0xaa, 0xa8, 0x00, 0x9a, 0x28, 0x2e, 0x05, 0x90, 0x8f, 0x15, 0x2d, 0x49, 0xc1,
	0x00, 0x68, 0xc6, 0x70, 0x59, 0x49, 0x25, 0x5d, 0xcf, 0x08, 0x70, 0x97, 0x87, 0x35, 0xaf, 0xef,
	0x27, 0x12, 0x0a, 0x09, 0x24, 0xa6, 0xc0, 0xc8, 0x87, 0x61, 0xcc, 0x14, 0x1d, 0x92, 0x44, 0x72,
	0x61, 0x94, 0xfd, 0xfb, 0x99, 0xcc, 0x64, 0xd3, 0x12, 0xdd, 0xb5, 0xaf, 0x0f, 0xb7, 0x84, 0xa7,
	0xad, 0x88, 0xa7, 0x4c, 0x28, 0xae, 0x76, 0xe3, 0x91, 0x01, 0x1f, 0x7f, 0x45, 0xce, 0xad, 0x89,
	0xb1, 0x77, 0x5d, 0xe7, 0xc6, 0xdb, 0x4a, 0x16, 0x1e, 0x1a, 0xa0, 0xe0, 0xce, 0xac, 0xe9, 0xdd,
	0xa7, 0xce, 0x6d, 0x5d, 0xd7, 0x7c, 0x9d, 0x7a, 0xf6, 0x00, 0x05, 0x97, 0xd7, 0xf7, 0x30, 0x4f,
	0x01, 0x8f, 0xff, 0x0d, 0x9a, 0x5d, 0x68, 0xc2, 0x78, 0xe4, 0x52, 0xe7, 0xa6, 0xde, 0x05, 0xbc,
	0xde, 0xa0, 0x17, 0x5c, 0x5e, 0x3f, 0xc0, 0x66, 0x5b, 0xac, 0x8d, 0x71, 0xbb, 0x2d, 0x7e, 0x25,
	0xb9, 0x88, 0x9e, 0x1f, 0x7e, 0x5c, 0x59, 0xfb, 0x9f, 0x57, 0x41, 0xc6, 0xd5, 0xbb, 0x4d, 0x8c,
	0x13, 0x59, 0x90, 0x36, 0x9a, 0x29, 0xcf, 0x20, 0x7d, 0x4f, 0xd4, 0xae, 0x64, 0xd0, 0x08, 0x60,
	0x66, 0x26, 0x47, 0x9f, 0xec, 0x43, 0xed, 0xa3, 0x63, 0xed, 0xa3, 0x5f, 0xb5, 0x8f, 0x3e, 0x9f,
	0x7c, 0xeb, 0x78, 0xf2, 0xad, 0xef, 0x27, 0xdf, 0x72, 0x1e, 0x25, 0xb2, 0xc0, 0xe7, 0x7e, 0x2e,
	0xba, 0xdb, 0x66, 0x9c, 0xea, 0xd0, 0x53, 0xf4, 0x26, 0xea, 0x38, 0xbf, 0x04, 0x60, 0x6a, 0x42,
	0x85, 0xca, 0x19, 0x29, 0x64, 0xba, 0xc9, 0x19, 0x90, 0xff, 0xb8, 0xda, 0x17, 0xbb, 0x37, 0x5f,
	0x2c, 0xf7, 0xb6, 0x37, 0x37, 0xa6, 0x8b, 0xae, 0xe9, 0xb2, 0xa2, 0xe5, 0xb7, 0xbf, 0xd0, 0xaa,
	0x0b, 0xad, 0x34, 0x54, 0xdb, 0x4f, 0xce, 0x41, 0xab, 0xd7, 0xd3, 0x68, 0xc2, 0x14, 0x4d, 0xa9,
	0xa2, 0xbf, 0xed, 0xbe, 0xa1, 0x85, 0x61, 0x97, 0x17, 0x86, 0x9a, 0x18, 0x5f, 0x34, 0x17, 0x7c,
	0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x58, 0xa5, 0x07, 0xfc, 0x63, 0x02, 0x00, 0x00,
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
				i = encodeVarintMessage(dAtA, i, uint64(size))
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
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
