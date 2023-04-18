// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/key/key.proto

package key

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

type Key struct {
	IdentityID *base.IdentityID `protobuf:"bytes,1,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_66f66b7657fe10bc, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetIdentityID() *base.IdentityID {
	if m != nil {
		return m.IdentityID
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "identities.Key")
}

func init() {
	proto.RegisterFile("x/identities/internal/key/key.proto", fileDescriptor_66f66b7657fe10bc)
}

var fileDescriptor_66f66b7657fe10bc = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xae, 0xd0, 0xcf, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0xd1, 0xcf, 0x4e, 0xad, 0x04, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x2e, 0x84,
	0x12, 0x29, 0x69, 0x90, 0x86, 0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0x98, 0xce, 0x4a, 0x4f, 0x17,
	0x88, 0x42, 0x25, 0x0b, 0x2e, 0x66, 0xef, 0xd4, 0x4a, 0x21, 0x43, 0x2e, 0x1e, 0x98, 0x54, 0x7c,
	0x66, 0x7c, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xbf, 0x5e, 0x66, 0x4a, 0xb1, 0x9e,
	0x27, 0x5c, 0x4f, 0x10, 0x17, 0x42, 0xbf, 0xd3, 0x16, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0xe0, 0xe2, 0x4b, 0xce, 0xcf, 0xd5, 0x43, 0xb8, 0xc0, 0x89, 0xc3, 0x3b, 0xb5,
	0x32, 0x00, 0x64, 0x5d, 0x00, 0x63, 0x94, 0x45, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x49, 0x4e, 0xaa, 0x7e, 0x6e,
	0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0xb1, 0x3e, 0x4e, 0xdf, 0x2d, 0x62, 0x62, 0xf6, 0x8c, 0x88, 0x58,
	0xc5, 0xc4, 0xe5, 0x09, 0x97, 0x3f, 0x85, 0xcc, 0x79, 0xc4, 0x24, 0x86, 0xe0, 0xc4, 0xb8, 0x07,
	0x38, 0xf9, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0xbe, 0x42, 0x56, 0x95, 0xc4, 0x06, 0xf6,
	0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x67, 0xa9, 0xfd, 0x25, 0x47, 0x01, 0x00, 0x00,
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
		n += 1 + l + sovKey(uint64(l))
	}
	return n
}

func sovKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKey(x uint64) (n int) {
	return sovKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKey
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
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func skipKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
				return 0, ErrInvalidLengthKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKey = fmt.Errorf("proto: unexpected end of group")
)
