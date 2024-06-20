// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/orders/transactions/revoke/message.proto

package revoke

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
	proto "github.com/cosmos/gogoproto/proto"
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
	From             string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID           *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ToID             *base.IdentityID       `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3" json:"to_i_d,omitempty"`
	ClassificationID *base.ClassificationID `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b55db8a80a4eebf, []int{0}
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

func (m *Message) GetToID() *base.IdentityID {
	if m != nil {
		return m.ToID
	}
	return nil
}

func (m *Message) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "AssetMantle.modules.x.orders.transactions.revoke.Message")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/orders/transactions/revoke/message.proto", fileDescriptor_3b55db8a80a4eebf)
}

var fileDescriptor_3b55db8a80a4eebf = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x3f, 0xcb, 0xd3, 0x40,
	0x1c, 0x07, 0xf0, 0x5e, 0x9e, 0x12, 0x35, 0x3a, 0x48, 0xa6, 0xf2, 0x0c, 0xa1, 0xb8, 0xd8, 0xa5,
	0x77, 0x62, 0x75, 0x39, 0xff, 0x40, 0x6b, 0x40, 0x32, 0x84, 0x96, 0xd0, 0x21, 0x68, 0x20, 0x5c,
	0x73, 0xd7, 0xf6, 0xb0, 0xc9, 0x49, 0xee, 0x2a, 0x75, 0xf4, 0x1d, 0xf8, 0x1a, 0x1c, 0x7d, 0x25,
	0xe2, 0xd4, 0xd1, 0x51, 0xd2, 0xcd, 0xd7, 0xe0, 0x20, 0xc9, 0x05, 0xb9, 0x76, 0x28, 0x64, 0x6a,
	0x09, 0xf9, 0x7c, 0xbf, 0xbf, 0xbb, 0xfc, 0x9c, 0xd7, 0x53, 0x29, 0x99, 0x0a, 0x49, 0xa1, 0x76,
	0x0c, 0xe5, 0x82, 0xee, 0x77, 0x4c, 0xa2, 0x03, 0x12, 0x25, 0x65, 0xa5, 0x44, 0xaa, 0x24, 0x85,
	0x24, 0x99, 0xe2, 0xa2, 0x90, 0xa8, 0x64, 0x9f, 0xc4, 0x07, 0x86, 0x72, 0x26, 0x25, 0xd9, 0x30,
	0xf8, 0xb1, 0x14, 0x4a, 0xb8, 0x4f, 0x0c, 0x0f, 0x5b, 0x0f, 0x0f, 0x50, 0x7b, 0x68, 0x7a, 0xa8,
	0xfd, 0xed, 0xc4, 0x6c, 0x94, 0xd9, 0x96, 0xe5, 0x04, 0x71, 0x2a, 0xd1, 0x8a, 0x48, 0x86, 0xb2,
	0x1d, 0x91, 0x92, 0xaf, 0x79, 0x46, 0x6a, 0x93, 0x72, 0xaa, 0x6b, 0x6e, 0xc7, 0xd7, 0x10, 0xa7,
	0xac, 0x50, 0x5c, 0x7d, 0xfe, 0xff, 0xfa, 0xa3, 0xbf, 0xc0, 0xb9, 0x13, 0xea, 0x39, 0x5d, 0xd7,
	0xe9, 0xaf, 0x4b, 0x91, 0x0f, 0xc0, 0x10, 0x8c, 0xee, 0x45, 0xcd, 0x7f, 0x77, 0xea, 0xdc, 0xad,
	0x7f, 0x53, 0x9e, 0xd2, 0x81, 0x35, 0x04, 0xa3, 0xfb, 0x4f, 0x1f, 0x43, 0xf3, 0x20, 0xba, 0x01,
	0x72, 0x2a, 0x61, 0xdd, 0x00, 0x83, 0xb6, 0x21, 0xf0, 0x23, 0xbb, 0x86, 0x81, 0xef, 0xbe, 0x72,
	0x6c, 0x25, 0x9a, 0x80, 0x9b, 0x6e, 0x01, 0x7d, 0x25, 0x02, 0xdf, 0x7d, 0xef, 0xb8, 0x97, 0x67,
	0x4d, 0xe9, 0xa0, 0xdf, 0x44, 0x8d, 0xaf, 0x46, 0xbd, 0x39, 0x63, 0x81, 0x1f, 0x3d, 0xcc, 0x2e,
	0x9e, 0xcc, 0xbe, 0xdc, 0xfc, 0xa8, 0x3c, 0x70, 0xac, 0x3c, 0xf0, 0xbb, 0xf2, 0xc0, 0xd7, 0x93,
	0xd7, 0x3b, 0x9e, 0xbc, 0xde, 0xaf, 0x93, 0xd7, 0x73, 0x9e, 0x65, 0x22, 0x87, 0x5d, 0xbf, 0xd9,
	0xec, 0x41, 0x7b, 0x99, 0x8b, 0xfa, 0x76, 0x17, 0xe0, 0xdd, 0x8b, 0x0d, 0x57, 0xdb, 0xfd, 0x0a,
	0x66, 0x22, 0x47, 0x5d, 0x17, 0xe8, 0x9b, 0x65, 0x4f, 0xc3, 0x78, 0xbe, 0x8c, 0xbe, 0x5b, 0x67,
	0xcb, 0x13, 0xb6, 0x83, 0xc4, 0x70, 0xae, 0x07, 0x59, 0x9a, 0x83, 0x44, 0x8d, 0xfd, 0x79, 0x46,
	0x92, 0x96, 0x24, 0x71, 0xa2, 0x49, 0x62, 0x92, 0x44, 0x93, 0xca, 0x7a, 0xd9, 0x95, 0x24, 0x6f,
	0x17, 0xb3, 0x90, 0x29, 0x42, 0x89, 0x22, 0x7f, 0xac, 0xe7, 0x06, 0xc7, 0xb8, 0xf5, 0x18, 0xc7,
	0x18, 0xeb, 0x04, 0x8c, 0xcd, 0x08, 0x8c, 0x75, 0xc6, 0xca, 0x6e, 0x36, 0x71, 0xf2, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0xc2, 0xab, 0x30, 0x61, 0x03, 0x00, 0x00,
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
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ToID != nil {
		{
			size, err := m.ToID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.ToID != nil {
		l = m.ToID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ToID", wireType)
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
			if m.ToID == nil {
				m.ToID = &base.IdentityID{}
			}
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
