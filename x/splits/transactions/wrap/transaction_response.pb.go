// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/transactions/wrap/transaction_response.proto

package wrap

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
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

type TransactionResponse struct {
	CoinID *base.CoinID `protobuf:"bytes,1,opt,name=coin_i_d,json=coinID,proto3" json:"coin_i_d,omitempty"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_073f1ac95c378fb0, []int{0}
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

func (m *TransactionResponse) GetCoinID() *base.CoinID {
	if m != nil {
		return m.CoinID
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "assetmantle.modules.splits.transactions.wrap.TransactionResponse")
}

func init() {
	proto.RegisterFile("splits/transactions/wrap/transaction_response.proto", fileDescriptor_073f1ac95c378fb0)
}

var fileDescriptor_073f1ac95c378fb0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0x87, 0x9b, 0x80, 0x45, 0xe2, 0x56, 0x41, 0x8a, 0x43, 0x10, 0x5d, 0x1c, 0xca, 0x5d, 0xb1,
	0x8b, 0x1e, 0x38, 0xb4, 0x0a, 0xe2, 0x10, 0x28, 0x6d, 0xa0, 0x20, 0x81, 0x72, 0xbd, 0x3b, 0xec,
	0x41, 0x73, 0x17, 0xf2, 0xbf, 0xa2, 0x8f, 0xe1, 0x33, 0x38, 0xfa, 0x24, 0xe2, 0xd4, 0xd1, 0x51,
	0xd2, 0xcd, 0xdd, 0x5d, 0xae, 0x17, 0xf0, 0x04, 0x1d, 0x3a, 0x05, 0xc2, 0xf7, 0xfb, 0xfe, 0x1f,
	0x49, 0xd4, 0x83, 0x62, 0x21, 0x0d, 0x60, 0x53, 0x52, 0x05, 0x94, 0x19, 0xa9, 0x15, 0xe0, 0x87,
	0x92, 0x16, 0xfe, 0x9b, 0x69, 0x29, 0xa0, 0xd0, 0x0a, 0x04, 0x2a, 0x4a, 0x6d, 0x74, 0xab, 0x43,
	0x01, 0x84, 0xc9, 0xa9, 0x32, 0x0b, 0x81, 0x72, 0xcd, 0x97, 0x0b, 0x01, 0xc8, 0x89, 0x90, 0x2f,
	0x42, 0x56, 0x74, 0x78, 0x20, 0x39, 0xe0, 0x19, 0x05, 0x81, 0x99, 0x96, 0x6a, 0x2a, 0xb9, 0xb3,
	0x1c, 0xa7, 0xd1, 0x7e, 0xfa, 0x03, 0x8f, 0xea, 0x13, 0xad, 0xcb, 0x68, 0xd7, 0x71, 0x53, 0xde,
	0x0e, 0x8e, 0x82, 0xd3, 0xbd, 0xb3, 0x13, 0xe4, 0xdf, 0x03, 0x36, 0x17, 0x39, 0x45, 0x92, 0x03,
	0xb2, 0x52, 0x74, 0xa5, 0xa5, 0xba, 0xbd, 0x1e, 0x35, 0xd9, 0xe6, 0x39, 0xf8, 0x0a, 0x5f, 0xab,
	0x38, 0x58, 0x55, 0x71, 0xf0, 0x51, 0xc5, 0xc1, 0xd3, 0x3a, 0x6e, 0xac, 0xd6, 0x71, 0xe3, 0x7d,
	0x1d, 0x37, 0xa2, 0x2e, 0xd3, 0x39, 0xda, 0x26, 0x7d, 0xd0, 0xfe, 0x23, 0x70, 0x68, 0xe3, 0x87,
	0xc1, 0xdd, 0xc5, 0xbd, 0x34, 0xf3, 0xe5, 0x0c, 0x31, 0x9d, 0xe3, 0xbe, 0x95, 0x26, 0x1b, 0x29,
	0xae, 0xa5, 0xf8, 0x11, 0xff, 0xf7, 0x69, 0x9f, 0xc3, 0x9d, 0x7e, 0x32, 0x4e, 0x27, 0x2f, 0x61,
	0xa7, 0xef, 0xd5, 0x24, 0x75, 0xcd, 0xd8, 0xd5, 0xa4, 0x7e, 0xcd, 0xa4, 0xa4, 0xc5, 0xdb, 0x2f,
	0x3c, 0xab, 0xf1, 0xcc, 0xe1, 0x99, 0x8f, 0x67, 0x16, 0xaf, 0xc2, 0xf3, 0x6d, 0xf0, 0xec, 0x66,
	0x38, 0x48, 0x84, 0xa1, 0x9c, 0x1a, 0xfa, 0x19, 0x76, 0xbd, 0x29, 0x21, 0xf5, 0x96, 0x10, 0x37,
	0x26, 0xc4, 0x5f, 0x13, 0x62, 0xe7, 0xb3, 0xe6, 0xe6, 0xa7, 0xf6, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0xc7, 0xc3, 0x30, 0x51, 0x02, 0x00, 0x00,
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
	if m.CoinID != nil {
		{
			size, err := m.CoinID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransactionResponse(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
	if m.CoinID != nil {
		l = m.CoinID.Size()
		n += 1 + l + sovTransactionResponse(uint64(l))
	}
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionResponse
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
				return ErrInvalidLengthTransactionResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinID == nil {
				m.CoinID = &base.CoinID{}
			}
			if err := m.CoinID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
