// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/classifications/internal/queries/classifications/queryRequest.proto

package classifications

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3441e07b32c471d9, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "classifications.queries.classifications.QueryRequest")
}

func init() {
	proto.RegisterFile("modules/classifications/internal/queries/classifications/queryRequest.proto", fileDescriptor_3441e07b32c471d9)
}

var fileDescriptor_3441e07b32c471d9 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xce, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d,
	0x2d, 0xca, 0xc4, 0xa2, 0x00, 0x24, 0x5e, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x57,
	0x50, 0x94, 0x5f, 0x92, 0x2f, 0xa4, 0x8e, 0xa6, 0x46, 0x0f, 0xaa, 0x57, 0x0f, 0x4d, 0x5c, 0x4a,
	0x2b, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x58, 0x3f, 0x29, 0xb1, 0x38, 0x15, 0x62, 0x90, 0x7e, 0x99,
	0x61, 0x52, 0x6a, 0x49, 0xa2, 0xa1, 0x7e, 0x41, 0x62, 0x7a, 0x66, 0x1e, 0x58, 0x19, 0xc4, 0x50,
	0xa5, 0x30, 0x2e, 0x9e, 0x40, 0x24, 0xab, 0x84, 0xdc, 0xb8, 0xb8, 0x10, 0x6a, 0x24, 0x98, 0x14,
	0x18, 0x35, 0xb8, 0x8d, 0xd4, 0xf4, 0x20, 0x06, 0xea, 0x81, 0x0c, 0x04, 0xdb, 0x5a, 0xa9, 0x07,
	0x35, 0x50, 0x2f, 0x20, 0x31, 0x3d, 0x15, 0xaa, 0x37, 0x08, 0x49, 0xa7, 0xd3, 0x7b, 0xa6, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0xd2, 0x4e, 0xce, 0xcf, 0xd5, 0x23, 0xd2,
	0x2f, 0x4e, 0x82, 0xc8, 0xae, 0x0b, 0x00, 0x39, 0x39, 0x80, 0x31, 0x2a, 0x3c, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8, 0x38, 0xb5, 0xc4, 0x37, 0x31, 0xaf,
	0x24, 0x27, 0x55, 0x1f, 0x16, 0xda, 0xe4, 0x86, 0xfa, 0x22, 0x26, 0x66, 0xe7, 0x40, 0xe7, 0x55,
	0x4c, 0xea, 0xce, 0x68, 0xae, 0x0b, 0x84, 0xba, 0x0e, 0x4d, 0xfc, 0x14, 0x86, 0xca, 0x18, 0xa8,
	0xca, 0x18, 0x34, 0xf1, 0x47, 0x4c, 0xc6, 0x44, 0xaa, 0x8c, 0x71, 0x0f, 0x70, 0xf2, 0x4d, 0x2d,
	0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0xa4, 0x89, 0x26, 0x6b, 0x65, 0x05, 0xd5, 0x66, 0x65,
	0x85, 0x26, 0x93, 0xc4, 0x06, 0x8e, 0x50, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xc0,
	0xb2, 0x55, 0x74, 0x02, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryRequest(uint64(l))
	}
	return n
}

func sovQueryRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRequest(x uint64) (n int) {
	return sovQueryRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRequest
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRequest
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
				return ErrInvalidLengthQueryRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRequest
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
func skipQueryRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRequest
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
					return 0, ErrIntOverflowQueryRequest
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
					return 0, ErrIntOverflowQueryRequest
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
				return 0, ErrInvalidLengthQueryRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRequest = fmt.Errorf("proto: unexpected end of group")
)
