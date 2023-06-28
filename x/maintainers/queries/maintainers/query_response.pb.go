// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: maintainers/queries/maintainers/query_response.proto

package maintainers

import (
	fmt "fmt"
	mappable "github.com/AssetMantle/modules/x/maintainers/mappable"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryResponse struct {
	List         []*mappable.Mappable `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	PageResponse *query.PageResponse  `protobuf:"bytes,2,opt,name=page_response,json=pageResponse,proto3" json:"page_response,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55014fe8b1c5a67, []int{0}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryResponse)(nil), "assetmantle.modules.maintainers.queries.maintainers.QueryResponse")
}

func init() {
	proto.RegisterFile("maintainers/queries/maintainers/query_response.proto", fileDescriptor_e55014fe8b1c5a67)
}

var fileDescriptor_e55014fe8b1c5a67 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0xe3, 0x40,
	0x10, 0x85, 0x6d, 0x27, 0x77, 0x85, 0x73, 0x69, 0xac, 0x2b, 0xa2, 0x14, 0xbe, 0xe8, 0xae, 0xb8,
	0x88, 0x62, 0x57, 0x49, 0x90, 0x22, 0xb9, 0x81, 0xa4, 0x80, 0x86, 0x95, 0x9c, 0x94, 0xc8, 0x12,
	0x1a, 0x27, 0x2b, 0x63, 0xc9, 0xf6, 0x1a, 0xef, 0x06, 0x91, 0x7f, 0x40, 0x49, 0x4b, 0x87, 0x28,
	0xe1, 0x8f, 0x20, 0xaa, 0x94, 0x94, 0xc8, 0xe9, 0xf8, 0x15, 0xc8, 0xf6, 0x12, 0x6d, 0x68, 0x22,
	0xa5, 0x1b, 0x3f, 0xcd, 0xfb, 0xfc, 0xde, 0x68, 0xcd, 0xc3, 0x18, 0xc2, 0x44, 0x40, 0x98, 0xd0,
	0x8c, 0xe3, 0xab, 0x05, 0xcd, 0x42, 0xca, 0xf1, 0x77, 0x6d, 0x79, 0x91, 0x51, 0x9e, 0xb2, 0x84,
	0x53, 0x94, 0x66, 0x4c, 0x30, 0x6b, 0x00, 0x9c, 0x53, 0x11, 0x43, 0x22, 0x22, 0x8a, 0x62, 0x36,
	0x5f, 0x44, 0x94, 0x23, 0xc5, 0x85, 0x24, 0x49, 0xd5, 0xda, 0x07, 0x33, 0xc6, 0x63, 0xc6, 0xb1,
	0x0f, 0x9c, 0x56, 0x58, 0x7c, 0xdd, 0xf3, 0xa9, 0x80, 0x1e, 0x4e, 0x21, 0x08, 0x13, 0x10, 0x21,
	0x4b, 0xaa, 0x1f, 0xb4, 0x7f, 0x07, 0x2c, 0x60, 0xe5, 0x88, 0x8b, 0x49, 0xaa, 0xff, 0xd4, 0x60,
	0x31, 0xa4, 0x29, 0xf8, 0x11, 0xdd, 0x0c, 0xd5, 0xd2, 0xdf, 0x67, 0xdd, 0x6c, 0x4e, 0x0a, 0xfa,
	0x54, 0x66, 0xb6, 0x4e, 0xcc, 0x7a, 0x14, 0x72, 0xd1, 0xd2, 0x3b, 0xb5, 0x6e, 0xa3, 0xdf, 0x47,
	0xbb, 0xc2, 0x6f, 0x80, 0x44, 0x0e, 0xd3, 0xd2, 0x6f, 0x9d, 0x99, 0xcd, 0x14, 0x02, 0xba, 0x39,
	0x46, 0xcb, 0xe8, 0xe8, 0xdd, 0x46, 0xff, 0x3f, 0xaa, 0x8a, 0xa1, 0xa2, 0x58, 0xd9, 0x7c, 0x89,
	0x64, 0x31, 0xe4, 0x42, 0x40, 0xbf, 0x72, 0x4c, 0x7f, 0xa5, 0xca, 0x97, 0x53, 0xbf, 0x7d, 0xf8,
	0xa3, 0x8d, 0xef, 0x6b, 0x2f, 0xb9, 0xad, 0xaf, 0x72, 0x5b, 0x7f, 0xcf, 0x6d, 0xfd, 0x6e, 0x6d,
	0x6b, 0xab, 0xb5, 0xad, 0xbd, 0xad, 0x6d, 0xcd, 0x1c, 0xce, 0x58, 0x8c, 0xf6, 0x38, 0xf4, 0xd8,
	0xda, 0xaa, 0xef, 0x16, 0x57, 0x71, 0xf5, 0xf3, 0xe3, 0x20, 0x14, 0x97, 0x0b, 0x1f, 0xcd, 0x58,
	0x8c, 0x47, 0x05, 0x95, 0x94, 0x54, 0x2c, 0xa9, 0xf8, 0x06, 0xef, 0x78, 0x0a, 0x8f, 0xc6, 0x8f,
	0x11, 0x21, 0x13, 0xf2, 0x64, 0x0c, 0x46, 0x4a, 0x2a, 0x22, 0x53, 0x11, 0x25, 0xd5, 0x44, 0xa6,
	0x52, 0xb4, 0xd7, 0x2d, 0x97, 0x27, 0x5d, 0x9e, 0xb2, 0xe1, 0x49, 0x97, 0xaa, 0xe5, 0xc6, 0xd1,
	0x1e, 0x2e, 0xef, 0xd4, 0x1d, 0x13, 0x2a, 0x60, 0x0e, 0x02, 0x3e, 0x8c, 0xa1, 0x42, 0x70, 0x1c,
	0x89, 0x70, 0x1c, 0x65, 0xdf, 0x71, 0x24, 0x64, 0x4b, 0xf5, 0x7f, 0x96, 0x0f, 0x6a, 0xf0, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0xc9, 0x9e, 0x09, 0x24, 0x03, 0x00, 0x00,
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PageResponse != nil {
		{
			size, err := m.PageResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryResponse(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryResponse(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovQueryResponse(uint64(l))
		}
	}
	if m.PageResponse != nil {
		l = m.PageResponse.Size()
		n += 1 + l + sovQueryResponse(uint64(l))
	}
	return n
}

func sovQueryResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryResponse(x uint64) (n int) {
	return sovQueryResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResponse
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
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
				return ErrInvalidLengthQueryResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &mappable.Mappable{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
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
				return ErrInvalidLengthQueryResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PageResponse == nil {
				m.PageResponse = &query.PageResponse{}
			}
			if err := m.PageResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryResponse
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
func skipQueryResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryResponse
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
					return 0, ErrIntOverflowQueryResponse
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
					return 0, ErrIntOverflowQueryResponse
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
				return 0, ErrInvalidLengthQueryResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryResponse = fmt.Errorf("proto: unexpected end of group")
)
