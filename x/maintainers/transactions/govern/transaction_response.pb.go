// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/maintainers/transactions/govern/transaction_response.proto

package govern

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44c0657d1d3b388, []int{0}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "AssetMantle.modules.x.maintainers.transactions.govern.TransactionResponse")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/maintainers/transactions/govern/transaction_response.proto", fileDescriptor_e44c0657d1d3b388)
}

var fileDescriptor_e44c0657d1d3b388 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x06, 0xe0, 0x24, 0xc2, 0x15, 0x29, 0x15, 0xc1, 0x6a, 0x0b, 0x1f, 0x60, 0xb7, 0x10, 0x0b,
	0xa7, 0x32, 0x69, 0x52, 0x2d, 0x04, 0x49, 0x11, 0x64, 0x41, 0xf6, 0x72, 0xcb, 0x19, 0xb8, 0xec,
	0x1e, 0xbb, 0x7b, 0x72, 0x8f, 0xe1, 0x33, 0x58, 0xea, 0x8b, 0x88, 0xd5, 0x95, 0x96, 0x92, 0x74,
	0x3e, 0x85, 0x78, 0x59, 0x70, 0x02, 0x56, 0x69, 0x07, 0xfe, 0x6f, 0x7e, 0x86, 0x49, 0xcb, 0xcc,
	0x39, 0xe5, 0xb9, 0xd4, 0x7e, 0xa3, 0x58, 0x67, 0x56, 0xbb, 0x8d, 0x72, 0x6c, 0xcf, 0x3a, 0xd9,
	0x6a, 0x2f, 0x5b, 0xad, 0xac, 0x63, 0xde, 0x4a, 0xed, 0x64, 0xe3, 0x5b, 0xa3, 0x1d, 0x5b, 0x9b,
	0x27, 0x65, 0x35, 0x9e, 0x3d, 0x58, 0xe5, 0xb6, 0x46, 0x3b, 0x45, 0xb7, 0xd6, 0x78, 0x73, 0x7a,
	0x8d, 0x44, 0x1a, 0x44, 0xba, 0xa7, 0x48, 0xa4, 0x58, 0xa4, 0xa3, 0x78, 0x79, 0x9e, 0x9e, 0x55,
	0x7f, 0xe3, 0xbb, 0x60, 0xe6, 0x6f, 0x27, 0xef, 0x3d, 0x89, 0x0f, 0x3d, 0x89, 0xbf, 0x7a, 0x12,
	0x3f, 0x0f, 0x24, 0x3a, 0x0c, 0x24, 0xfa, 0x1c, 0x48, 0x94, 0xde, 0x34, 0xa6, 0xa3, 0xb3, 0x96,
	0xe5, 0x17, 0xff, 0xac, 0x2a, 0x7f, 0xdb, 0x97, 0xf1, 0xfd, 0xed, 0xba, 0xf5, 0x8f, 0xbb, 0x25,
	0x6d, 0x4c, 0xc7, 0x66, 0x1d, 0xe7, 0x25, 0x59, 0x64, 0xbc, 0xe6, 0x55, 0xf1, 0x9a, 0x4c, 0x6e,
	0xc1, 0x43, 0xbd, 0x9a, 0x72, 0x54, 0xaf, 0xc2, 0xf5, 0x8a, 0x23, 0xf0, 0x31, 0xc9, 0x89, 0x90,
	0x13, 0xb5, 0x40, 0x39, 0x81, 0x73, 0x62, 0xcc, 0xf5, 0x49, 0x36, 0x2b, 0x27, 0x8a, 0x32, 0xe7,
	0xca, 0xcb, 0x95, 0xf4, 0xf2, 0x3b, 0x01, 0x64, 0x00, 0x04, 0x04, 0xa0, 0x06, 0x40, 0x0c, 0x00,
	0x76, 0x00, 0x46, 0x68, 0xb9, 0x38, 0xbe, 0xc0, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88,
	0x0e, 0xd3, 0xf7, 0x56, 0x02, 0x00, 0x00,
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTransactionResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTransactionResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionResponse(x uint64) (n int) {
	return sovTransactionResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionResponse
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
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionResponse
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
func skipTransactionResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionResponse
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
					return 0, ErrIntOverflowTransactionResponse
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
					return 0, ErrIntOverflowTransactionResponse
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
				return 0, ErrInvalidLengthTransactionResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionResponse = fmt.Errorf("proto: unexpected end of group")
)
