// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/deputize/message.proto

package deputize

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/lists/base"
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
	From                 string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID               *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ToID                 *base.IdentityID       `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3" json:"to_i_d,omitempty"`
	ClassificationID     *base.ClassificationID `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	MaintainedProperties *base1.PropertyList    `protobuf:"bytes,5,opt,name=maintained_properties,json=maintainedProperties,proto3" json:"maintained_properties,omitempty"`
	CanMintAsset         bool                   `protobuf:"varint,6,opt,name=can_mint_asset,json=canMintAsset,proto3" json:"can_mint_asset,omitempty"`
	CanBurnAsset         bool                   `protobuf:"varint,7,opt,name=can_burn_asset,json=canBurnAsset,proto3" json:"can_burn_asset,omitempty"`
	CanRenumerateAsset   bool                   `protobuf:"varint,8,opt,name=can_renumerate_asset,json=canRenumerateAsset,proto3" json:"can_renumerate_asset,omitempty"`
	CanAddMaintainer     bool                   `protobuf:"varint,9,opt,name=can_add_maintainer,json=canAddMaintainer,proto3" json:"can_add_maintainer,omitempty"`
	CanRemoveMaintainer  bool                   `protobuf:"varint,10,opt,name=can_remove_maintainer,json=canRemoveMaintainer,proto3" json:"can_remove_maintainer,omitempty"`
	CanMutateMaintainer  bool                   `protobuf:"varint,11,opt,name=can_mutate_maintainer,json=canMutateMaintainer,proto3" json:"can_mutate_maintainer,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74c3809526902bd, []int{0}
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

func (m *Message) GetMaintainedProperties() *base1.PropertyList {
	if m != nil {
		return m.MaintainedProperties
	}
	return nil
}

func (m *Message) GetCanMintAsset() bool {
	if m != nil {
		return m.CanMintAsset
	}
	return false
}

func (m *Message) GetCanBurnAsset() bool {
	if m != nil {
		return m.CanBurnAsset
	}
	return false
}

func (m *Message) GetCanRenumerateAsset() bool {
	if m != nil {
		return m.CanRenumerateAsset
	}
	return false
}

func (m *Message) GetCanAddMaintainer() bool {
	if m != nil {
		return m.CanAddMaintainer
	}
	return false
}

func (m *Message) GetCanRemoveMaintainer() bool {
	if m != nil {
		return m.CanRemoveMaintainer
	}
	return false
}

func (m *Message) GetCanMutateMaintainer() bool {
	if m != nil {
		return m.CanMutateMaintainer
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "identities.transactions.deputize.Message")
}

func init() {
	proto.RegisterFile("x/identities/internal/transactions/deputize/message.proto", fileDescriptor_e74c3809526902bd)
}

var fileDescriptor_e74c3809526902bd = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0xb3, 0xdb, 0x98, 0xa6, 0xd3, 0xa2, 0x65, 0xda, 0x40, 0xa8, 0xb0, 0xc6, 0x52, 0x21,
	0x42, 0xd9, 0x91, 0x7a, 0x32, 0xb7, 0xa6, 0x01, 0x5d, 0x71, 0x21, 0x2c, 0x3d, 0x49, 0x60, 0x99,
	0xec, 0x4c, 0xeb, 0x40, 0x76, 0x26, 0xcc, 0x7c, 0x2b, 0xad, 0x4f, 0x21, 0x3e, 0x82, 0x47, 0x1f,
	0xc1, 0x27, 0x10, 0x4f, 0x3d, 0x7a, 0x94, 0xe4, 0xe6, 0x53, 0xc8, 0xce, 0x6e, 0x36, 0x6b, 0x15,
	0x8a, 0xa7, 0x5d, 0xe6, 0xfb, 0xfd, 0xbe, 0xf9, 0xfe, 0xf0, 0x0d, 0x7a, 0x71, 0x45, 0x04, 0xe3,
	0x12, 0x04, 0x08, 0x6e, 0x88, 0x90, 0xc0, 0xb5, 0xa4, 0x33, 0x02, 0x9a, 0x4a, 0x43, 0x13, 0x10,
	0x4a, 0x1a, 0xc2, 0xf8, 0x3c, 0x03, 0xf1, 0x81, 0x93, 0x94, 0x1b, 0x43, 0x2f, 0xb9, 0x3f, 0xd7,
	0x0a, 0x14, 0xee, 0xad, 0x45, 0xbf, 0xce, 0xfb, 0x2b, 0xfe, 0xe0, 0x71, 0xde, 0xdc, 0x90, 0x29,
	0x35, 0x9c, 0x24, 0x33, 0x6a, 0x8c, 0xb8, 0x10, 0x09, 0xcd, 0xa1, 0x60, 0x54, 0x34, 0x39, 0x78,
	0x58, 0x43, 0xca, 0x7e, 0xd7, 0x55, 0xf1, 0xd1, 0x15, 0x99, 0x09, 0x03, 0x65, 0x79, 0xae, 0xd5,
	0x9c, 0x6b, 0xb8, 0x7e, 0x23, 0x0c, 0x14, 0xc0, 0xe1, 0xa7, 0x26, 0xda, 0x0c, 0x8b, 0xa1, 0x30,
	0x46, 0xcd, 0x0b, 0xad, 0xd2, 0xae, 0xd3, 0x73, 0xfa, 0x5b, 0x91, 0xfd, 0xc7, 0x4f, 0x51, 0x3b,
	0xff, 0xc6, 0x22, 0x66, 0x5d, 0xb7, 0xe7, 0xf4, 0xb7, 0x4f, 0x1e, 0xf8, 0x82, 0x19, 0x3f, 0xa8,
	0x6e, 0x8a, 0x5a, 0x39, 0x10, 0x8c, 0xf0, 0x13, 0xd4, 0x02, 0x65, 0xc1, 0x8d, 0x7f, 0x83, 0x4d,
	0x50, 0xc1, 0x08, 0x9f, 0x21, 0xfc, 0x67, 0x12, 0xab, 0x34, 0xad, 0xd2, 0xb1, 0xca, 0xd9, 0xad,
	0xa0, 0xd1, 0xee, 0xed, 0xe8, 0xf8, 0x15, 0xea, 0xa4, 0x54, 0x48, 0xa0, 0x42, 0x72, 0x16, 0x97,
	0xb9, 0x04, 0x37, 0xdd, 0x7b, 0xb6, 0xcf, 0x9e, 0x6f, 0x53, 0xfb, 0xe3, 0x5a, 0xe0, 0x68, 0x7f,
	0x6d, 0x8c, 0x2b, 0x01, 0x1f, 0xa1, 0xfb, 0x09, 0x95, 0x71, 0x2a, 0x24, 0xc4, 0xd4, 0x18, 0x0e,
	0xdd, 0x56, 0xcf, 0xe9, 0xb7, 0xa3, 0x9d, 0x84, 0xca, 0x50, 0x48, 0x38, 0xcd, 0xcf, 0x56, 0xd4,
	0x34, 0xd3, 0xb2, 0xa4, 0x36, 0x2b, 0x6a, 0x98, 0x69, 0x59, 0x50, 0xcf, 0xd0, 0x7e, 0x4e, 0x69,
	0x2e, 0xb3, 0x94, 0x6b, 0x0a, 0xbc, 0x64, 0xdb, 0x96, 0xc5, 0x09, 0x95, 0x51, 0x55, 0x2a, 0x8c,
	0x63, 0x94, 0x9f, 0xc6, 0x94, 0xb1, 0xb8, 0x9a, 0x4e, 0x77, 0xb7, 0x2c, 0xbf, 0x9b, 0x50, 0x79,
	0xca, 0x58, 0x58, 0x9d, 0xe3, 0x13, 0xd4, 0x29, 0xfa, 0xa7, 0xea, 0x3d, 0xaf, 0x0b, 0xc8, 0x0a,
	0x7b, 0xf6, 0x82, 0xbc, 0xf6, 0xb7, 0x93, 0x66, 0x90, 0xcf, 0x53, 0x73, 0xb6, 0x2b, 0x27, 0xb4,
	0xb5, 0xb5, 0x33, 0xfc, 0xea, 0x7e, 0x5b, 0x78, 0xce, 0xcd, 0xc2, 0x73, 0x7e, 0x2e, 0x3c, 0xe7,
	0xe3, 0xd2, 0x6b, 0xdc, 0x2c, 0xbd, 0xc6, 0x8f, 0xa5, 0xd7, 0x40, 0x47, 0x89, 0x4a, 0xfd, 0xbb,
	0xd6, 0x76, 0xb8, 0x53, 0xae, 0xd4, 0x38, 0xdf, 0xb1, 0xb1, 0xf3, 0xf6, 0xf5, 0xa5, 0x80, 0x77,
	0xd9, 0xd4, 0x4f, 0x54, 0x4a, 0x6c, 0xf0, 0x90, 0x4a, 0x98, 0x71, 0x92, 0x2a, 0x96, 0xcd, 0xb8,
	0x21, 0xff, 0xf1, 0x84, 0x3e, 0xbb, 0x1b, 0xc1, 0xf9, 0xe8, 0x8b, 0xdb, 0x0b, 0xd6, 0x43, 0x9c,
	0xd7, 0x87, 0x18, 0x95, 0xe0, 0xf7, 0x3a, 0x32, 0xa9, 0x23, 0x93, 0x15, 0xb2, 0x70, 0x8f, 0xef,
	0x42, 0x26, 0x2f, 0xc7, 0xc3, 0x90, 0x03, 0x65, 0x14, 0xe8, 0x2f, 0xf7, 0x70, 0x8d, 0x0f, 0x06,
	0x75, 0x7e, 0x30, 0x58, 0x09, 0xd3, 0x96, 0x7d, 0x58, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0x53, 0xc5, 0xf9, 0x18, 0x04, 0x00, 0x00,
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
	if m.CanMutateMaintainer {
		i--
		if m.CanMutateMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.CanRemoveMaintainer {
		i--
		if m.CanRemoveMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.CanAddMaintainer {
		i--
		if m.CanAddMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.CanRenumerateAsset {
		i--
		if m.CanRenumerateAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CanBurnAsset {
		i--
		if m.CanBurnAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CanMintAsset {
		i--
		if m.CanMintAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MaintainedProperties != nil {
		{
			size, err := m.MaintainedProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
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
	if m.MaintainedProperties != nil {
		l = m.MaintainedProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.CanMintAsset {
		n += 2
	}
	if m.CanBurnAsset {
		n += 2
	}
	if m.CanRenumerateAsset {
		n += 2
	}
	if m.CanAddMaintainer {
		n += 2
	}
	if m.CanRemoveMaintainer {
		n += 2
	}
	if m.CanMutateMaintainer {
		n += 2
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainedProperties", wireType)
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
			if m.MaintainedProperties == nil {
				m.MaintainedProperties = &base1.PropertyList{}
			}
			if err := m.MaintainedProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMintAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanMintAsset = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanBurnAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanBurnAsset = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRenumerateAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanRenumerateAsset = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanAddMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanAddMaintainer = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRemoveMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanRemoveMaintainer = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMutateMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanMutateMaintainer = bool(v != 0)
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
