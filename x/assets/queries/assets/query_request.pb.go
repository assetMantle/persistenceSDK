// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/queries/assets/query_request.proto

package assets

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
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae204ce9ce05a19, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.assets.queries.assets.QueryRequest")
}

func init() {
	proto.RegisterFile("assets/queries/assets/query_request.proto", fileDescriptor_5ae204ce9ce05a19)
}

var fileDescriptor_5ae204ce9ce05a19 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x2f, 0x05, 0x1d, 0xaa, 0x8b, 0x37, 0x89, 0x43, 0x10, 0x07, 0xf1, 0x84, 0x4b, 0x38,
	0x45, 0x87, 0x6c, 0xe9, 0xa0, 0x53, 0xa1, 0xbd, 0xc1, 0x41, 0x0a, 0x92, 0xf6, 0x42, 0x2d, 0x5c,
	0x9b, 0xbb, 0x26, 0x15, 0xef, 0x2d, 0x7c, 0x06, 0x47, 0x9f, 0x44, 0x9c, 0x6e, 0x74, 0x70, 0x90,
	0x76, 0xf3, 0x29, 0xa4, 0xcd, 0x5f, 0xac, 0xe0, 0x70, 0xe3, 0xd7, 0xfe, 0xbe, 0x5f, 0xbe, 0x36,
	0xee, 0x48, 0x68, 0x2d, 0x8d, 0xa6, 0xcb, 0x4a, 0x96, 0x99, 0xd4, 0xb4, 0x17, 0x57, 0x77, 0xa5,
	0x5c, 0x56, 0x52, 0x1b, 0xb2, 0x28, 0x95, 0x51, 0x43, 0x8b, 0xe6, 0xa2, 0x30, 0x73, 0x49, 0x72,
	0x35, 0xab, 0xe6, 0x52, 0x13, 0xcb, 0x13, 0xa8, 0x43, 0x3c, 0x38, 0x4d, 0x94, 0xce, 0x95, 0xa6,
	0xb1, 0xd0, 0xd2, 0xba, 0xe8, 0xc3, 0x24, 0x96, 0x46, 0x4c, 0xe8, 0x42, 0xa4, 0x59, 0x21, 0x4c,
	0xa6, 0x0a, 0xab, 0x3d, 0xba, 0x71, 0x77, 0xc3, 0x96, 0x98, 0xda, 0xc3, 0x86, 0x57, 0xae, 0xfb,
	0xcb, 0xec, 0xa3, 0x43, 0x74, 0xb2, 0x73, 0x76, 0x4c, 0xac, 0x90, 0xb4, 0xc2, 0xee, 0xb0, 0x15,
	0x01, 0x21, 0x09, 0x44, 0x2a, 0xa1, 0x3b, 0xed, 0x35, 0xbd, 0x0f, 0xe7, 0xb5, 0xc6, 0x68, 0x5d,
	0x63, 0xf4, 0x59, 0x63, 0xf4, 0xd4, 0xe0, 0xc1, 0xba, 0xc1, 0x83, 0xf7, 0x06, 0x0f, 0xdc, 0x71,
	0xa2, 0x72, 0xb2, 0xf1, 0xd7, 0x78, 0x7b, 0xfd, 0x7d, 0x41, 0x3b, 0x3a, 0x40, 0xb7, 0x97, 0x69,
	0x66, 0xee, 0xab, 0x98, 0x24, 0x2a, 0xa7, 0xbc, 0xe5, 0xfc, 0x4e, 0x45, 0x41, 0x45, 0x1f, 0xe9,
	0xbf, 0x7f, 0xf6, 0xd9, 0xd9, 0xe2, 0x3e, 0x0f, 0xf9, 0x8b, 0x33, 0xe2, 0xbd, 0x01, 0x3e, 0x0c,
	0xe0, 0x76, 0x40, 0x08, 0x03, 0x6c, 0x7c, 0xfb, 0xc3, 0x46, 0xc0, 0x46, 0xf6, 0x65, 0x04, 0x2c,
	0xc4, 0xda, 0xb9, 0xd8, 0x98, 0x8d, 0xae, 0x03, 0xcf, 0x97, 0x46, 0xcc, 0x84, 0x11, 0x5f, 0xce,
	0xb8, 0xd7, 0x63, 0x0c, 0x8a, 0x8c, 0x59, 0x94, 0x31, 0xa8, 0xfe, 0x3c, 0x88, 0xb7, 0xbb, 0xdb,
	0x3b, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x69, 0x02, 0xc5, 0xaf, 0x41, 0x02, 0x00, 0x00,
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
		dAtA[i] = 0xa
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
		case 1:
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
