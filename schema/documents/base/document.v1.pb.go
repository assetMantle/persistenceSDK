// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/documents/base/document.v1.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
	base1 "github.com/AssetMantle/modules/schema/qualified/base"
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

type Document struct {
	ClassificationID *base.ClassificationID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	Immutables       *base1.Immutables      `protobuf:"bytes,2,opt,name=immutables,proto3" json:"immutables,omitempty"`
	Mutables         *base1.Mutables        `protobuf:"bytes,3,opt,name=mutables,proto3" json:"mutables,omitempty"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db520417af385cc, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Document.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return m.Size()
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Document)(nil), "documents.Document")
}

func init() {
	proto.RegisterFile("schema/documents/base/document.v1.proto", fileDescriptor_6db520417af385cc)
}

var fileDescriptor_6db520417af385cc = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x4f, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x4a, 0x2c,
	0x4e, 0x85, 0x73, 0xf5, 0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xe1, 0x2a,
	0xa4, 0x34, 0xa1, 0x7a, 0x32, 0x53, 0xa0, 0xaa, 0x93, 0x73, 0x12, 0x8b, 0x8b, 0x33, 0xd3, 0x32,
	0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x3c, 0x5d, 0xe0, 0xba, 0xe0, 0x4a, 0x0b, 0x4b, 0x13, 0x73,
	0x32, 0xd3, 0x32, 0x53, 0x53, 0x20, 0x1a, 0x32, 0x73, 0x73, 0x4b, 0x4b, 0x12, 0x93, 0x72, 0x52,
	0x8b, 0x11, 0x4a, 0xd5, 0xb1, 0x2b, 0xc5, 0x54, 0x28, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea,
	0x83, 0x58, 0x10, 0x51, 0xa5, 0xbd, 0x8c, 0x5c, 0x1c, 0x2e, 0x50, 0x27, 0x0a, 0x39, 0x73, 0x09,
	0xa1, 0xba, 0x29, 0x3e, 0x33, 0x3e, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x54, 0x2f,
	0x33, 0xa5, 0x58, 0xcf, 0x19, 0xcd, 0xc9, 0x41, 0x02, 0xe8, 0x9e, 0x10, 0x32, 0xe5, 0xe2, 0x42,
	0xb8, 0x53, 0x82, 0x09, 0xaa, 0x19, 0xee, 0x3c, 0x3d, 0x4f, 0xb8, 0x64, 0x10, 0x92, 0x42, 0x21,
	0x7d, 0x2e, 0x0e, 0xb8, 0x26, 0x66, 0xb0, 0x26, 0x61, 0x24, 0x4d, 0xbe, 0x30, 0x2d, 0x70, 0x45,
	0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0x9c, 0x36, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x03, 0x17, 0x6f, 0x72, 0x7e, 0xae, 0x1e, 0x3c, 0xf8, 0x9d, 0xf8, 0x61, 0xde, 0x0c,
	0x33, 0x0c, 0x00, 0x79, 0x3d, 0x80, 0x31, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8, 0x38, 0xb5, 0xc4, 0x37, 0x31, 0xaf, 0x24, 0x27, 0x55, 0x3f,
	0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x58, 0x1f, 0x6b, 0x24, 0x2f, 0x62, 0x62, 0x76, 0x89, 0x88,
	0x58, 0xc5, 0xc4, 0x09, 0x33, 0xaf, 0xf8, 0x14, 0x12, 0xfb, 0x11, 0x93, 0x28, 0x9c, 0x1d, 0xe3,
	0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58, 0x92, 0xf8, 0x0a, 0x49, 0x4d, 0x12, 0x1b,
	0x38, 0xec, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xff, 0x29, 0x4c, 0x46, 0x02, 0x00,
	0x00,
}

func (m *Document) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Document) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Document) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mutables != nil {
		{
			size, err := m.Mutables.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocumentV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Immutables != nil {
		{
			size, err := m.Immutables.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocumentV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDocumentV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocumentV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocumentV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Document) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovDocumentV1(uint64(l))
	}
	if m.Immutables != nil {
		l = m.Immutables.Size()
		n += 1 + l + sovDocumentV1(uint64(l))
	}
	if m.Mutables != nil {
		l = m.Mutables.Size()
		n += 1 + l + sovDocumentV1(uint64(l))
	}
	return n
}

func sovDocumentV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocumentV1(x uint64) (n int) {
	return sovDocumentV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Document) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocumentV1
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
			return fmt.Errorf("proto: Document: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Document: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentV1
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
				return ErrInvalidLengthDocumentV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentV1
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Immutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentV1
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
				return ErrInvalidLengthDocumentV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Immutables == nil {
				m.Immutables = &base1.Immutables{}
			}
			if err := m.Immutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentV1
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
				return ErrInvalidLengthDocumentV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mutables == nil {
				m.Mutables = &base1.Mutables{}
			}
			if err := m.Mutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocumentV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocumentV1
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
func skipDocumentV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocumentV1
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
					return 0, ErrIntOverflowDocumentV1
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
					return 0, ErrIntOverflowDocumentV1
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
				return 0, ErrInvalidLengthDocumentV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocumentV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocumentV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocumentV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocumentV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocumentV1 = fmt.Errorf("proto: unexpected end of group")
)
