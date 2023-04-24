// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/queries/orders/query_response.proto

package orders

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	mappable "github.com/AssetMantle/modules/x/orders/internal/mappable"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type QueryResponse struct {
	List []*mappable.Mappable `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e6807eb54d95bce, []int{0}
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
	proto.RegisterType((*QueryResponse)(nil), "assetmantle.modules.orders.queries.orders.QueryResponse")
}

func init() {
	proto.RegisterFile("x/orders/internal/queries/orders/query_response.proto", fileDescriptor_0e6807eb54d95bce)
}

var fileDescriptor_0e6807eb54d95bce = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xa9, 0xd0, 0xcf, 0x2f,
	0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x86, 0x89, 0x83, 0xb8, 0x95, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9a, 0x89, 0xc5, 0xc5, 0xa9, 0x25, 0xb9,
	0x89, 0x79, 0x25, 0x39, 0xa9, 0x7a, 0xb9, 0xf9, 0x29, 0xa5, 0x39, 0xa9, 0xc5, 0x7a, 0x10, 0xf5,
	0x7a, 0x50, 0xed, 0x50, 0xae, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x97, 0x3e, 0x88, 0x05,
	0x31, 0x40, 0x4a, 0x13, 0xd3, 0xda, 0xdc, 0xc4, 0x82, 0x82, 0xc4, 0xa4, 0x9c, 0x54, 0x38, 0x03,
	0xa2, 0x54, 0xc9, 0x9f, 0x8b, 0x37, 0x10, 0xd9, 0x09, 0x42, 0x16, 0x5c, 0x2c, 0x39, 0x99, 0xc5,
	0x25, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x2a, 0x7a, 0x78, 0xdc, 0xe2, 0x0b, 0x35, 0x2a,
	0x08, 0xac, 0xc3, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0xd3, 0x33, 0xa6, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0xd2, 0x4d, 0xce, 0xcf, 0xd5, 0x23, 0xda, 0x6f, 0x4e, 0x42,
	0x28, 0x0e, 0x0b, 0x00, 0x39, 0x37, 0x80, 0x31, 0xca, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0xdf, 0x11, 0x64, 0x96, 0x2f, 0xd8, 0x2c, 0x7d, 0xa8, 0x59, 0xfa, 0x84,
	0x42, 0x7c, 0x11, 0x13, 0xab, 0xa3, 0xaf, 0x7f, 0xa0, 0xff, 0x2a, 0x26, 0x4d, 0x47, 0x24, 0xa7,
	0xf8, 0x42, 0x9d, 0xe2, 0x0f, 0x71, 0x4a, 0x20, 0xd4, 0x29, 0x10, 0xee, 0x29, 0x14, 0xb5, 0x31,
	0x50, 0xb5, 0x31, 0x10, 0xc9, 0x18, 0xa8, 0x5a, 0x28, 0xf7, 0x11, 0x93, 0x29, 0xd1, 0x6a, 0x63,
	0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0x31, 0xe9, 0x22, 0xe9,
	0xb3, 0xb2, 0x82, 0x6a, 0xb4, 0xb2, 0x82, 0x28, 0xb5, 0xb2, 0x82, 0x6a, 0x85, 0x09, 0x24, 0xb1,
	0x81, 0x23, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x26, 0x1f, 0xa9, 0x69, 0x64, 0x02, 0x00,
	0x00,
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
