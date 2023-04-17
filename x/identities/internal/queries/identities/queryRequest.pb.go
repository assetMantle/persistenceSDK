// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/queries/identities/queryRequest.proto

package identities

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
	PageRequest *query.PageRequest `protobuf:"bytes,2,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae3877e46cf633b6, []int{0}
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

func (m *QueryRequest) GetPageRequest() *query.PageRequest {
	if m != nil {
		return m.PageRequest
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "identities.queries.identities.QueryRequest")
}

func init() {
	proto.RegisterFile("x/identities/internal/queries/identities/queryRequest.proto", fileDescriptor_ae3877e46cf633b6)
}

var fileDescriptor_ae3877e46cf633b6 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x1b, 0x23, 0x31, 0xa4, 0x5d, 0xe8, 0x84, 0x90, 0x6a, 0x7e, 0x06, 0x04, 0x0c, 0xb6,
	0x0a, 0x5b, 0x98, 0xe8, 0x52, 0x65, 0xa8, 0x94, 0x76, 0x03, 0x55, 0x42, 0x4e, 0x7b, 0x15, 0x2c,
	0x35, 0x76, 0x1a, 0xdf, 0x20, 0xfa, 0x16, 0x3c, 0x03, 0x23, 0x12, 0xef, 0x81, 0x98, 0x3a, 0x32,
	0xa2, 0x74, 0xe3, 0x29, 0x50, 0x13, 0x4b, 0xce, 0x54, 0x31, 0xde, 0xf3, 0xf9, 0x9c, 0x7b, 0x6c,
	0xfb, 0xb7, 0x2f, 0x5c, 0xce, 0x41, 0xa1, 0x44, 0x09, 0x86, 0x4b, 0x85, 0x90, 0x2b, 0xb1, 0xe0,
	0xcb, 0x02, 0xf2, 0x4a, 0x70, 0x6c, 0x2b, 0xad, 0x26, 0xb0, 0x2c, 0xc0, 0x20, 0xcb, 0x72, 0x8d,
	0xba, 0xdb, 0x73, 0x98, 0x59, 0x07, 0x73, 0xd2, 0xd1, 0xd5, 0x4c, 0x9b, 0x54, 0x1b, 0x1e, 0x0b,
	0x03, 0xb5, 0x9d, 0x3f, 0xf7, 0x63, 0x40, 0xd1, 0xe7, 0x99, 0x48, 0xa4, 0x12, 0x28, 0xb5, 0xaa,
	0xa3, 0xce, 0xee, 0xfd, 0xce, 0xb8, 0xb1, 0xa0, 0x1b, 0xfa, 0x9d, 0x4c, 0x24, 0xf0, 0x98, 0xd7,
	0xf3, 0x21, 0x39, 0xf1, 0x2e, 0xda, 0xd7, 0xe7, 0xac, 0x8e, 0x64, 0xdb, 0xc8, 0x6a, 0xe5, 0x8a,
	0xd9, 0x48, 0x16, 0x89, 0x04, 0xac, 0x7b, 0xd2, 0xce, 0xdc, 0x30, 0xf8, 0x20, 0x9f, 0x25, 0xf5,
	0xd6, 0x25, 0xf5, 0x7e, 0x4a, 0xea, 0xbd, 0x6e, 0x68, 0x6b, 0xbd, 0xa1, 0xad, 0xef, 0x0d, 0x6d,
	0xf9, 0xa7, 0x33, 0x9d, 0xb2, 0x9d, 0x97, 0x18, 0x1c, 0x34, 0x6b, 0x45, 0xdb, 0xae, 0x91, 0xf7,
	0x30, 0x4c, 0x24, 0x3e, 0x15, 0x31, 0x9b, 0xe9, 0x94, 0xdf, 0x19, 0x03, 0x38, 0x12, 0x0a, 0x17,
	0xc0, 0x53, 0x3d, 0x2f, 0x16, 0x60, 0xf8, 0x7f, 0x1f, 0xf5, 0x8d, 0xec, 0x85, 0xe3, 0xf0, 0x9d,
	0xf4, 0x42, 0xd7, 0x61, 0x6c, 0x3b, 0x38, 0xe9, 0xab, 0xc9, 0xa7, 0x96, 0x4f, 0x9d, 0x54, 0x92,
	0xcb, 0x9d, 0x7c, 0x3a, 0x8c, 0x06, 0x23, 0x40, 0x31, 0x17, 0x28, 0x7e, 0xc9, 0xb1, 0x03, 0x41,
	0x60, 0x0f, 0x07, 0x81, 0x13, 0xe3, 0xfd, 0xea, 0x47, 0x6e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x2a, 0x31, 0xd8, 0xbb, 0x1b, 0x02, 0x00, 0x00,
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
	if m.PageRequest != nil {
		{
			size, err := m.PageRequest.MarshalToSizedBuffer(dAtA[:i])
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
	if m.PageRequest != nil {
		l = m.PageRequest.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field PageRequest", wireType)
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
			if m.PageRequest == nil {
				m.PageRequest = &query.PageRequest{}
			}
			if err := m.PageRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
