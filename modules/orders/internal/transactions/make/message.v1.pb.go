// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/make/message.v1.proto

package make

import (
	fmt "fmt"
	_ "github.com/AssetMantle/modules/schema/data/base"
	base "github.com/AssetMantle/modules/schema/ids/base"
	base3 "github.com/AssetMantle/modules/schema/lists/base"
	base1 "github.com/AssetMantle/modules/schema/types/base"
	_ "github.com/cosmos/cosmos-proto"
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
	From                    string                                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                  *base.IdentityID                       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ClassificationID        *base.ClassificationID                 `protobuf:"bytes,3,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	TakerID                 *base.IdentityID                       `protobuf:"bytes,4,opt,name=taker_i_d,json=takerID,proto3" json:"taker_i_d,omitempty"`
	MakerOwnableID          *base.AnyOwnableID                     `protobuf:"bytes,5,opt,name=maker_ownable_i_d,json=makerOwnableID,proto3" json:"maker_ownable_i_d,omitempty"`
	TakerOwnableID          *base.AnyOwnableID                     `protobuf:"bytes,6,opt,name=taker_ownable_i_d,json=takerOwnableID,proto3" json:"taker_ownable_i_d,omitempty"`
	ExpiresIn               *base1.Height                          `protobuf:"bytes,7,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	MakerOwnableSplit       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=maker_ownable_split,json=makerOwnableSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maker_ownable_split"`
	TakerOwnableSplit       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=taker_ownable_split,json=takerOwnableSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"taker_ownable_split"`
	ImmutableMetaProperties *base3.PropertyList                    `protobuf:"bytes,10,opt,name=immutable_meta_properties,json=immutableMetaProperties,proto3" json:"immutable_meta_properties,omitempty"`
	ImmutableProperties     *base3.PropertyList                    `protobuf:"bytes,11,opt,name=immutable_properties,json=immutableProperties,proto3" json:"immutable_properties,omitempty"`
	MutableMetaProperties   *base3.PropertyList                    `protobuf:"bytes,12,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties       *base3.PropertyList                    `protobuf:"bytes,13,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_73dd7719843f5c27, []int{0}
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

func (m *Message) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func (m *Message) GetTakerID() *base.IdentityID {
	if m != nil {
		return m.TakerID
	}
	return nil
}

func (m *Message) GetMakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.MakerOwnableID
	}
	return nil
}

func (m *Message) GetTakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.TakerOwnableID
	}
	return nil
}

func (m *Message) GetExpiresIn() *base1.Height {
	if m != nil {
		return m.ExpiresIn
	}
	return nil
}

func (m *Message) GetImmutableMetaProperties() *base3.PropertyList {
	if m != nil {
		return m.ImmutableMetaProperties
	}
	return nil
}

func (m *Message) GetImmutableProperties() *base3.PropertyList {
	if m != nil {
		return m.ImmutableProperties
	}
	return nil
}

func (m *Message) GetMutableMetaProperties() *base3.PropertyList {
	if m != nil {
		return m.MutableMetaProperties
	}
	return nil
}

func (m *Message) GetMutableProperties() *base3.PropertyList {
	if m != nil {
		return m.MutableProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "orders.transactions.make.Message")
}

func init() {
	proto.RegisterFile("modules/orders/internal/transactions/make/message.v1.proto", fileDescriptor_73dd7719843f5c27)
}

var fileDescriptor_73dd7719843f5c27 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xb4, 0x34, 0xcd, 0x96, 0x16, 0xe2, 0xb6, 0xaa, 0x1b, 0xa1, 0xb4, 0xfc, 0x6f,
	0x05, 0x5d, 0xab, 0x70, 0x8b, 0x7a, 0x69, 0x1a, 0x41, 0x23, 0x88, 0x12, 0x99, 0x8a, 0x03, 0xaa,
	0x14, 0x6d, 0xec, 0x6d, 0xb2, 0x8a, 0xed, 0x8d, 0xbc, 0x53, 0x20, 0xe2, 0x05, 0x38, 0x22, 0x1e,
	0x81, 0x23, 0x67, 0x1e, 0xa2, 0xe2, 0x54, 0x71, 0x42, 0x1c, 0x2a, 0x94, 0xde, 0x78, 0x0a, 0xb4,
	0x6b, 0xd7, 0x75, 0x93, 0x58, 0xe2, 0xc2, 0xc9, 0xb6, 0xe6, 0xf7, 0xcd, 0x37, 0x9e, 0x99, 0x5d,
	0x54, 0xf6, 0xb8, 0x73, 0xec, 0x52, 0x61, 0xf2, 0xc0, 0xa1, 0x81, 0x30, 0x99, 0x0f, 0x34, 0xf0,
	0x89, 0x6b, 0x42, 0x40, 0x7c, 0x41, 0x6c, 0x60, 0xdc, 0x17, 0xa6, 0x47, 0x7a, 0xd4, 0xf4, 0xa8,
	0x10, 0xa4, 0x43, 0xf1, 0xdb, 0x6d, 0xdc, 0x0f, 0x38, 0x70, 0xdd, 0x08, 0x35, 0x38, 0x89, 0x62,
	0x89, 0x16, 0x97, 0x3a, 0xbc, 0xc3, 0x15, 0x64, 0xca, 0xb7, 0x90, 0x2f, 0xae, 0xda, 0x5c, 0x78,
	0x5c, 0xb4, 0xc2, 0x40, 0xf8, 0x11, 0x85, 0x36, 0x85, 0xdd, 0xa5, 0x1e, 0x31, 0x99, 0x23, 0xcc,
	0x36, 0x11, 0xd4, 0xb4, 0x5d, 0x22, 0x04, 0x3b, 0x62, 0x36, 0x91, 0x69, 0x6b, 0xd5, 0xd8, 0xb5,
	0x78, 0x77, 0x14, 0x65, 0x0e, 0xf5, 0x81, 0xc1, 0x20, 0x09, 0xdd, 0x1f, 0x85, 0x88, 0x3f, 0x68,
	0xbc, 0xf3, 0x49, 0xdb, 0xa5, 0x49, 0xec, 0x61, 0x84, 0xb9, 0x4c, 0x40, 0x04, 0xf6, 0x03, 0xde,
	0xa7, 0x01, 0x0c, 0x5e, 0x32, 0x01, 0x97, 0xe0, 0xed, 0x08, 0x84, 0x41, 0x9f, 0x46, 0x60, 0x97,
	0xb2, 0x4e, 0x77, 0x02, 0xe2, 0x10, 0x20, 0x21, 0xe1, 0x50, 0xbb, 0x4a, 0x80, 0xc4, 0xc8, 0x9d,
	0x8f, 0x39, 0x94, 0xab, 0x87, 0x5d, 0xd4, 0x75, 0x34, 0x7d, 0x14, 0x70, 0xcf, 0xd0, 0xd6, 0xb5,
	0x8d, 0xbc, 0xa5, 0xde, 0xf5, 0x4d, 0x34, 0x2b, 0x9f, 0x2d, 0xd6, 0x72, 0x8c, 0xec, 0xba, 0xb6,
	0x31, 0xf7, 0xe4, 0x06, 0x66, 0x8e, 0xc0, 0xb5, 0xf8, 0x0f, 0xad, 0x19, 0x09, 0xd4, 0xaa, 0xfa,
	0x1e, 0xd2, 0xaf, 0xb6, 0x48, 0x89, 0xa6, 0x94, 0x68, 0x59, 0x89, 0xf6, 0x46, 0x3a, 0x68, 0xdd,
	0x1c, 0xed, 0xa9, 0xfe, 0x08, 0xe5, 0x81, 0xf4, 0x68, 0xa0, 0xb4, 0xd3, 0x93, 0x0d, 0x73, 0x8a,
	0xa8, 0x55, 0xf5, 0x1d, 0x54, 0xf0, 0x14, 0xcc, 0xc3, 0x3e, 0x2a, 0xd1, 0x35, 0x25, 0x2a, 0x28,
	0xd1, 0x6e, 0xa2, 0xc5, 0xd6, 0x82, 0x62, 0xe3, 0x6f, 0xa9, 0x86, 0x31, 0xf5, 0x4c, 0xaa, 0x1a,
	0xae, 0xaa, 0x1f, 0x23, 0x44, 0xdf, 0xf7, 0x59, 0x40, 0x45, 0x8b, 0xf9, 0x46, 0x4e, 0xc9, 0xe6,
	0xb1, 0x1a, 0x06, 0xde, 0x57, 0x73, 0xb0, 0xf2, 0x11, 0x50, 0xf3, 0xf5, 0x0f, 0x68, 0xf1, 0x6a,
	0xa5, 0xa2, 0xef, 0x32, 0x30, 0x66, 0x23, 0x99, 0x1c, 0x10, 0xae, 0x86, 0xb3, 0xa9, 0xec, 0x9c,
	0x9c, 0xad, 0x65, 0x7e, 0x9d, 0xad, 0x3d, 0xe8, 0x30, 0xe8, 0x1e, 0xb7, 0xb1, 0xcd, 0xbd, 0x68,
	0x33, 0xa3, 0xc7, 0x96, 0x70, 0x7a, 0xe1, 0xd8, 0xa5, 0xe6, 0xc7, 0xb7, 0x2d, 0x14, 0x2d, 0x6e,
	0x95, 0xda, 0x56, 0x21, 0xf9, 0x97, 0xaf, 0xa4, 0x8b, 0x34, 0x87, 0x09, 0xe6, 0xf9, 0xff, 0x60,
	0x0e, 0x63, 0xe6, 0x0d, 0xb4, 0xca, 0x3c, 0xef, 0x18, 0x94, 0xb1, 0x47, 0x81, 0xb4, 0xa2, 0x75,
	0x66, 0x54, 0x18, 0x48, 0x95, 0xb0, 0x88, 0xd5, 0xb2, 0xe3, 0x66, 0x62, 0xcf, 0xad, 0x95, 0x58,
	0x55, 0xa7, 0x40, 0x9a, 0xb1, 0x46, 0x7f, 0x86, 0x96, 0x2e, 0x13, 0x26, 0x72, 0xcd, 0xa5, 0xe7,
	0x5a, 0x8c, 0x05, 0x89, 0x3c, 0x2f, 0xd0, 0x4a, 0x5a, 0x59, 0xd7, 0xd3, 0x53, 0x2d, 0x4f, 0x2e,
	0xaa, 0x82, 0xf4, 0x09, 0x25, 0xcd, 0xa7, 0xe7, 0x29, 0x8c, 0x15, 0x54, 0xf9, 0x9c, 0x3d, 0x19,
	0x96, 0xb4, 0xd3, 0x61, 0x49, 0xfb, 0x3d, 0x2c, 0x69, 0x9f, 0xce, 0x4b, 0x99, 0xd3, 0xf3, 0x52,
	0xe6, 0xe7, 0x79, 0x29, 0x83, 0x6e, 0xd9, 0xdc, 0xc3, 0x69, 0x57, 0x5b, 0x65, 0x21, 0x3a, 0xc0,
	0xaf, 0xb7, 0x9b, 0xf2, 0x4c, 0x37, 0xb5, 0x37, 0xfb, 0x89, 0xe9, 0xed, 0x0a, 0x41, 0xa1, 0x4e,
	0x7c, 0x70, 0xa9, 0x79, 0x71, 0xb3, 0xfe, 0xf3, 0x0d, 0xfb, 0x25, 0x3b, 0xd5, 0x38, 0xa8, 0x7f,
	0xcd, 0x1a, 0x8d, 0xd0, 0xfc, 0x20, 0x69, 0x5e, 0x27, 0x3d, 0xfa, 0xfd, 0x22, 0x74, 0x98, 0x0c,
	0x1d, 0xca, 0xd0, 0x30, 0x7b, 0x2f, 0x2d, 0x74, 0xf8, 0xbc, 0x59, 0x91, 0x1d, 0x94, 0xab, 0xf7,
	0x27, 0x5b, 0x0c, 0xb1, 0x72, 0x39, 0xc9, 0x95, 0xcb, 0x12, 0x6c, 0xcf, 0xa8, 0x6b, 0xea, 0xe9,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x92, 0xe9, 0xae, 0x15, 0x06, 0x00, 0x00,
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
	if m.MutableProperties != nil {
		{
			size, err := m.MutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.MutableMetaProperties != nil {
		{
			size, err := m.MutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.ImmutableProperties != nil {
		{
			size, err := m.ImmutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.ImmutableMetaProperties != nil {
		{
			size, err := m.ImmutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	{
		size := m.TakerOwnableSplit.Size()
		i -= size
		if _, err := m.TakerOwnableSplit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessageV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.MakerOwnableSplit.Size()
		i -= size
		if _, err := m.MakerOwnableSplit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessageV1(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.ExpiresIn != nil {
		{
			size, err := m.ExpiresIn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.TakerOwnableID != nil {
		{
			size, err := m.TakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.MakerOwnableID != nil {
		{
			size, err := m.MakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.TakerID != nil {
		{
			size, err := m.TakerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
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
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.TakerID != nil {
		l = m.TakerID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.MakerOwnableID != nil {
		l = m.MakerOwnableID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.TakerOwnableID != nil {
		l = m.TakerOwnableID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.ExpiresIn != nil {
		l = m.ExpiresIn.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	l = m.MakerOwnableSplit.Size()
	n += 1 + l + sovMessageV1(uint64(l))
	l = m.TakerOwnableSplit.Size()
	n += 1 + l + sovMessageV1(uint64(l))
	if m.ImmutableMetaProperties != nil {
		l = m.ImmutableMetaProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.ImmutableProperties != nil {
		l = m.ImmutableProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.MutableMetaProperties != nil {
		l = m.MutableMetaProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.MutableProperties != nil {
		l = m.MutableProperties.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerID", wireType)
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
			if m.TakerID == nil {
				m.TakerID = &base.IdentityID{}
			}
			if err := m.TakerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
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
			if m.MakerOwnableID == nil {
				m.MakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.MakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
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
			if m.TakerOwnableID == nil {
				m.TakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.TakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresIn", wireType)
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
			if m.ExpiresIn == nil {
				m.ExpiresIn = &base1.Height{}
			}
			if err := m.ExpiresIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableSplit", wireType)
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
			if err := m.MakerOwnableSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
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
			if err := m.TakerOwnableSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
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
			if m.ImmutableMetaProperties == nil {
				m.ImmutableMetaProperties = &base3.PropertyList{}
			}
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
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
			if m.ImmutableProperties == nil {
				m.ImmutableProperties = &base3.PropertyList{}
			}
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			if m.MutableMetaProperties == nil {
				m.MutableMetaProperties = &base3.PropertyList{}
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			if m.MutableProperties == nil {
				m.MutableProperties = &base3.PropertyList{}
			}
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
