// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/deputize/message.proto

package deputize

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/lists/base"
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
	From                 string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID               *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ToID                 *base.IdentityID       `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3" json:"to_i_d,omitempty"`
	ClassificationID     *base.ClassificationID `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	MaintainedProperties *base1.PropertyList    `protobuf:"bytes,5,opt,name=maintained_properties,json=maintainedProperties,proto3" json:"maintained_properties,omitempty"`
	CanMakeOrder         bool                   `protobuf:"varint,6,opt,name=can_make_order,json=canMakeOrder,proto3" json:"can_make_order,omitempty"`
	CanModifyOrder       bool                   `protobuf:"varint,7,opt,name=can_modify_order,json=canModifyOrder,proto3" json:"can_modify_order,omitempty"`
	CanCancelOrder       bool                   `protobuf:"varint,8,opt,name=can_cancel_order,json=canCancelOrder,proto3" json:"can_cancel_order,omitempty"`
	CanAddMaintainer     bool                   `protobuf:"varint,9,opt,name=can_add_maintainer,json=canAddMaintainer,proto3" json:"can_add_maintainer,omitempty"`
	CanRemoveMaintainer  bool                   `protobuf:"varint,10,opt,name=can_remove_maintainer,json=canRemoveMaintainer,proto3" json:"can_remove_maintainer,omitempty"`
	CanMutateMaintainer  bool                   `protobuf:"varint,11,opt,name=can_mutate_maintainer,json=canMutateMaintainer,proto3" json:"can_mutate_maintainer,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_677c594995283e8c, []int{0}
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

func (m *Message) GetCanMakeOrder() bool {
	if m != nil {
		return m.CanMakeOrder
	}
	return false
}

func (m *Message) GetCanModifyOrder() bool {
	if m != nil {
		return m.CanModifyOrder
	}
	return false
}

func (m *Message) GetCanCancelOrder() bool {
	if m != nil {
		return m.CanCancelOrder
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
	proto.RegisterType((*Message)(nil), "assetmantle.modules.orders.transactions.deputize.Message")
}

func init() {
	proto.RegisterFile("orders/transactions/deputize/message.proto", fileDescriptor_677c594995283e8c)
}

var fileDescriptor_677c594995283e8c = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0xdb, 0x3e,
	0x18, 0xc7, 0xe3, 0x34, 0x4d, 0x53, 0xb5, 0xfc, 0x28, 0xfa, 0xad, 0x60, 0x72, 0x30, 0x61, 0x0c,
	0x16, 0xf6, 0xc7, 0x2e, 0x5d, 0x4f, 0x66, 0x3d, 0x24, 0x35, 0x8c, 0xc0, 0x4c, 0x82, 0xe9, 0x69,
	0x0b, 0x78, 0x8a, 0xa4, 0xb4, 0xa2, 0xb1, 0x15, 0x2c, 0x65, 0x2c, 0xbb, 0xed, 0x1d, 0xec, 0x35,
	0xec, 0xb8, 0x57, 0x32, 0x76, 0xea, 0x71, 0xc7, 0x91, 0xdc, 0x76, 0xd8, 0x6b, 0x18, 0x92, 0x15,
	0xc7, 0xe9, 0x46, 0x21, 0xa7, 0x84, 0xe7, 0xf9, 0x7e, 0x3e, 0x7a, 0x24, 0x4b, 0xe0, 0x09, 0xcf,
	0x08, 0xcd, 0x84, 0x27, 0x33, 0x94, 0x0a, 0x84, 0x25, 0xe3, 0xa9, 0xf0, 0x08, 0x9d, 0xce, 0x24,
	0xfb, 0x48, 0xbd, 0x84, 0x0a, 0x81, 0xae, 0xa8, 0x3b, 0xcd, 0xb8, 0xe4, 0xf0, 0x04, 0x09, 0x41,
	0x65, 0x82, 0x52, 0x39, 0xa1, 0x6e, 0xc2, 0xc9, 0x6c, 0x42, 0x85, 0x9b, 0xf3, 0x6e, 0x99, 0x77,
	0x57, 0x7c, 0xb3, 0xc5, 0x88, 0xf0, 0x46, 0x48, 0x50, 0x0f, 0x4f, 0x90, 0x10, 0x6c, 0xcc, 0x30,
	0x52, 0x91, 0x98, 0x91, 0xdc, 0xd9, 0x6c, 0x16, 0x09, 0x46, 0x68, 0x2a, 0x99, 0x9c, 0xaf, 0x7b,
	0xce, 0x84, 0x09, 0x69, 0xba, 0xd3, 0x8c, 0x4f, 0x69, 0x26, 0xe7, 0xb1, 0xaa, 0xe5, 0xfd, 0x87,
	0xbf, 0x6b, 0x60, 0x2f, 0xcc, 0x27, 0x84, 0x10, 0xd4, 0xc6, 0x19, 0x4f, 0x6c, 0xab, 0x65, 0xb5,
	0xf7, 0x23, 0xfd, 0x1f, 0x76, 0x40, 0x43, 0xfd, 0xc6, 0x2c, 0x26, 0x76, 0xb5, 0x65, 0xb5, 0x0f,
	0x4e, 0x1f, 0xbb, 0xe5, 0x2d, 0x08, 0x7c, 0x4d, 0x13, 0xe4, 0x32, 0x22, 0x5c, 0xb5, 0x86, 0xdb,
	0x33, 0x13, 0xf4, 0x82, 0xa8, 0xae, 0xc0, 0x5e, 0x00, 0xcf, 0x41, 0x5d, 0x72, 0x2d, 0xd8, 0xd9,
	0x4e, 0x50, 0x93, 0xbc, 0x17, 0xc0, 0xb7, 0x00, 0xde, 0xdd, 0x78, 0x4c, 0xec, 0x9a, 0x56, 0x3d,
	0xbf, 0x57, 0x75, 0xb1, 0x81, 0xf5, 0x82, 0xe8, 0x08, 0xdf, 0xa9, 0xc0, 0x77, 0xe0, 0x38, 0x41,
	0x2c, 0x95, 0x88, 0xa5, 0x94, 0xc4, 0xe6, 0x80, 0x18, 0x15, 0xf6, 0xae, 0xf6, 0x3f, 0xfd, 0x97,
	0x5f, 0x9f, 0x68, 0xbe, 0xc2, 0xc0, 0x9c, 0xe8, 0x6b, 0x26, 0x64, 0xf4, 0x60, 0x6d, 0x1a, 0x14,
	0x22, 0xf8, 0x08, 0xfc, 0x87, 0x51, 0x1a, 0x27, 0xe8, 0x86, 0xc6, 0xfa, 0x3b, 0xdb, 0xf5, 0x96,
	0xd5, 0x6e, 0x44, 0x87, 0x18, 0xa5, 0x21, 0xba, 0xa1, 0x7d, 0x55, 0x83, 0x6d, 0x70, 0xa4, 0x53,
	0x9c, 0xb0, 0xf1, 0xdc, 0xe4, 0xf6, 0x74, 0x4e, 0xd1, 0xa1, 0x2e, 0x6f, 0x24, 0x31, 0x4a, 0x31,
	0x9d, 0x98, 0x64, 0xa3, 0x48, 0x5e, 0xe8, 0x72, 0x9e, 0x7c, 0x06, 0xa0, 0x4a, 0x22, 0x42, 0xe2,
	0x62, 0xb2, 0xcc, 0xde, 0xd7, 0x59, 0xe5, 0xe8, 0x10, 0x12, 0x16, 0x75, 0x78, 0x0a, 0x8e, 0x55,
	0x3a, 0xa3, 0x09, 0x7f, 0x4f, 0xcb, 0x00, 0xd0, 0xc0, 0xff, 0x18, 0xa5, 0x91, 0xee, 0xfd, 0xcd,
	0x24, 0x33, 0x89, 0xe4, 0x06, 0x73, 0x50, 0x30, 0xa1, 0xee, 0xad, 0x99, 0xee, 0xa7, 0x9d, 0x6f,
	0x0b, 0xc7, 0xba, 0x5d, 0x38, 0xd6, 0xcf, 0x85, 0x63, 0x7d, 0x5e, 0x3a, 0x95, 0xdb, 0xa5, 0x53,
	0xf9, 0xb1, 0x74, 0x2a, 0xe0, 0x0c, 0xf3, 0xc4, 0xdd, 0xf6, 0x7d, 0x74, 0x0f, 0xcd, 0xf5, 0x1d,
	0xa8, 0xfb, 0x3c, 0xb0, 0xde, 0x9c, 0x5f, 0x31, 0x79, 0x3d, 0x1b, 0xb9, 0x98, 0x27, 0x5e, 0x47,
	0xc9, 0x42, 0x2d, 0xf3, 0x8c, 0xcc, 0xfb, 0xe0, 0xdd, 0xf7, 0x5c, 0xbf, 0x54, 0x77, 0x3b, 0x61,
	0xff, 0x32, 0xf8, 0x5a, 0x3d, 0xe9, 0x94, 0x26, 0x09, 0xcd, 0x24, 0xfd, 0x7c, 0x92, 0xcb, 0xf2,
	0x24, 0x81, 0x41, 0xbf, 0x6f, 0x20, 0x43, 0x83, 0x0c, 0x73, 0x64, 0x58, 0x46, 0x86, 0x2b, 0x64,
	0x51, 0x7d, 0xb9, 0x2d, 0x32, 0x7c, 0x35, 0xe8, 0x86, 0x54, 0x22, 0x82, 0x24, 0xfa, 0x55, 0x3d,
	0x2b, 0xe1, 0xbe, 0x6f, 0x78, 0xdf, 0xcf, 0x05, 0xbe, 0x5f, 0x36, 0xf8, 0xfe, 0x4a, 0x31, 0xaa,
	0xeb, 0xb7, 0xff, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa4, 0x4e, 0xf9, 0xb9, 0x04,
	0x00, 0x00,
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
	if m.CanCancelOrder {
		i--
		if m.CanCancelOrder {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CanModifyOrder {
		i--
		if m.CanModifyOrder {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CanMakeOrder {
		i--
		if m.CanMakeOrder {
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
	if m.CanMakeOrder {
		n += 2
	}
	if m.CanModifyOrder {
		n += 2
	}
	if m.CanCancelOrder {
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
				return fmt.Errorf("proto: wrong wireType = %d for field CanMakeOrder", wireType)
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
			m.CanMakeOrder = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanModifyOrder", wireType)
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
			m.CanModifyOrder = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanCancelOrder", wireType)
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
			m.CanCancelOrder = bool(v != 0)
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
