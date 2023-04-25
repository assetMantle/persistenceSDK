// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/maintainers/internal/queries/maintainers/query_request.proto

package maintainers

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryRequest struct {
	PageRequest *query.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4ec9840fdba3e5d, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.maintainers.queries.maintainers.QueryRequest")
}

func init() {
	proto.RegisterFile("x/maintainers/internal/queries/maintainers/query_request.proto", fileDescriptor_f4ec9840fdba3e5d)
}

var fileDescriptor_f4ec9840fdba3e5d = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x88, 0x0e, 0x69, 0x17, 0x3b, 0x89, 0x43, 0x10, 0x07, 0x11, 0x87, 0x3b, 0x6a,
	0x87, 0x42, 0x06, 0xa5, 0x5d, 0xa4, 0xc3, 0x41, 0xda, 0x4d, 0x09, 0xc8, 0x97, 0xf6, 0x23, 0x1e,
	0x34, 0x77, 0x69, 0xee, 0x22, 0xf6, 0x5f, 0xf8, 0x17, 0x74, 0xf4, 0x97, 0x88, 0x53, 0x47, 0x47,
	0x49, 0x37, 0x7f, 0x85, 0xa4, 0x39, 0xf1, 0xb2, 0x49, 0xc7, 0x7b, 0xef, 0x7b, 0x9f, 0xf7, 0xfd,
	0xb8, 0xf3, 0xae, 0x9e, 0x68, 0x0a, 0x5c, 0x68, 0xe0, 0x02, 0x73, 0x45, 0xb9, 0xd0, 0x98, 0x0b,
	0x58, 0xd0, 0x65, 0x81, 0x39, 0x47, 0xd5, 0xb8, 0xac, 0xb4, 0xd5, 0x7d, 0x8e, 0xcb, 0x02, 0x95,
	0x26, 0x59, 0x2e, 0xb5, 0xec, 0xf6, 0x41, 0x29, 0xd4, 0x29, 0x08, 0xbd, 0x40, 0x92, 0xca, 0x79,
	0xb1, 0x40, 0x45, 0x2c, 0x13, 0x31, 0x20, 0x5b, 0x3b, 0xbe, 0x98, 0x49, 0x95, 0x4a, 0x45, 0x63,
	0x50, 0x58, 0x53, 0xe9, 0x63, 0x2f, 0x46, 0x0d, 0x3d, 0x9a, 0x41, 0xc2, 0x05, 0x68, 0x2e, 0x45,
	0x1d, 0x70, 0x7a, 0xeb, 0x75, 0x26, 0xd5, 0xc4, 0xb4, 0x8e, 0xed, 0x8e, 0xbd, 0x4e, 0x06, 0x09,
	0xfe, 0xd6, 0x38, 0x72, 0x4e, 0x9c, 0xf3, 0xf6, 0xe5, 0x19, 0xa9, 0x91, 0xa4, 0x42, 0x6e, 0x33,
	0x57, 0xc4, 0x20, 0x49, 0x08, 0x09, 0x1a, 0xf7, 0xb4, 0x9d, 0xfd, 0x1d, 0x46, 0x2f, 0x7b, 0xef,
	0xa5, 0xef, 0xac, 0x4b, 0xdf, 0xf9, 0x2a, 0x7d, 0xe7, 0x79, 0xe3, 0xb7, 0xd6, 0x1b, 0xbf, 0xf5,
	0xb9, 0xf1, 0x5b, 0xde, 0x60, 0x26, 0x53, 0xb2, 0xc3, 0x6a, 0xa3, 0x43, 0xbb, 0x6c, 0x58, 0x6d,
	0x10, 0x3a, 0x77, 0xe3, 0x84, 0xeb, 0x87, 0x22, 0x26, 0x33, 0x99, 0xd2, 0x61, 0x05, 0x65, 0x5b,
	0x28, 0x35, 0x50, 0xfa, 0xff, 0x37, 0x78, 0x75, 0xf7, 0x87, 0x8c, 0x4d, 0xd8, 0x9b, 0xdb, 0x1f,
	0x5a, 0xed, 0x98, 0x69, 0xc7, 0xac, 0x76, 0x13, 0xd3, 0xce, 0xd2, 0x3e, 0x1a, 0xae, 0xc8, 0xb8,
	0x22, 0x6b, 0x22, 0x32, 0x2e, 0x5b, 0x2b, 0xdd, 0xeb, 0x1d, 0x5c, 0xd1, 0x4d, 0x38, 0x62, 0xa8,
	0x61, 0x0e, 0x1a, 0xbe, 0xdd, 0x81, 0x45, 0x08, 0x02, 0x83, 0x08, 0x02, 0x6b, 0x3e, 0x08, 0x0c,
	0xa4, 0xa1, 0xc6, 0x07, 0xdb, 0x5f, 0xd0, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x14, 0x42, 0xb0,
	0xd5, 0xa8, 0x02, 0x00, 0x00,
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
		case 1:
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
