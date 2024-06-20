// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/define/transaction_response.proto

package define

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
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
	ClassificationID *base.ClassificationID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b76490aa47e57f, []int{0}
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

func (m *TransactionResponse) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "AssetMantle.modules.x.identities.transactions.define.TransactionResponse")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/define/transaction_response.proto", fileDescriptor_c5b76490aa47e57f)
}

var fileDescriptor_c5b76490aa47e57f = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x1d, 0xe2, 0x22, 0x75, 0x29, 0x0e, 0x41, 0x9c, 0x5c, 0xfc, 0x0e, 0xac,
	0x83, 0xde, 0xa2, 0xad, 0x05, 0xe9, 0x10, 0x2c, 0xa5, 0x43, 0xd0, 0x40, 0xb9, 0xde, 0x5d, 0xed,
	0x41, 0x93, 0x2b, 0xf9, 0xae, 0xd0, 0x9f, 0xe1, 0x6f, 0x10, 0x5c, 0xfc, 0x25, 0xe2, 0xd4, 0xd1,
	0x51, 0xd2, 0xcd, 0x5f, 0x21, 0x6d, 0x02, 0xb9, 0x16, 0xa7, 0xac, 0x07, 0xcf, 0xf3, 0xbe, 0x2f,
	0xdf, 0x79, 0x8f, 0x6d, 0x44, 0x69, 0x02, 0x96, 0x98, 0x99, 0x24, 0xb1, 0x16, 0x8b, 0x99, 0x44,
	0xb2, 0x24, 0x4a, 0xc8, 0xc4, 0x28, 0xa3, 0x24, 0x12, 0x93, 0xb2, 0x04, 0x19, 0x37, 0x4a, 0x27,
	0x48, 0x84, 0x9c, 0xa8, 0x44, 0xda, 0x6f, 0xa3, 0x54, 0xe2, 0x5c, 0x27, 0x28, 0x61, 0x9e, 0x6a,
	0xa3, 0x1b, 0x57, 0x96, 0x10, 0x0a, 0x21, 0x2c, 0xa1, 0x14, 0x82, 0x2d, 0x84, 0x5c, 0x78, 0xd2,
	0xb2, 0x6b, 0x20, 0x9f, 0xca, 0x98, 0x11, 0x25, 0x90, 0x8c, 0x19, 0x4a, 0xc2, 0x67, 0x0c, 0x51,
	0x4d, 0x14, 0x67, 0xdb, 0x40, 0x25, 0xf2, 0xa8, 0xb3, 0xd4, 0x3b, 0x1e, 0x96, 0xae, 0x41, 0xd1,
	0xa3, 0xf1, 0xec, 0x35, 0xf6, 0x89, 0x91, 0x68, 0x3a, 0xa7, 0xce, 0xf9, 0xe1, 0xe5, 0x05, 0xd8,
	0xf5, 0xf2, 0x20, 0x50, 0x02, 0x61, 0x13, 0x04, 0xf7, 0x3b, 0x58, 0xaf, 0x3b, 0x38, 0xe2, 0x7b,
	0x2f, 0x9d, 0xf7, 0x83, 0xcf, 0xcc, 0x77, 0x56, 0x99, 0xef, 0xfc, 0x64, 0xbe, 0xf3, 0xba, 0xf6,
	0x6b, 0xab, 0xb5, 0x5f, 0xfb, 0x5e, 0xfb, 0x35, 0xef, 0x9a, 0xeb, 0x18, 0xaa, 0xac, 0xef, 0x34,
	0xff, 0x99, 0xd1, 0xdf, 0x4c, 0xec, 0x3b, 0x4f, 0xb7, 0x2f, 0xca, 0x4c, 0x17, 0x63, 0xe0, 0x3a,
	0x26, 0x55, 0x6e, 0xf5, 0xe6, 0xd6, 0xdb, 0x41, 0xd8, 0x1b, 0x76, 0x3f, 0xdc, 0x9d, 0xd3, 0x04,
	0x45, 0xb9, 0x10, 0x7a, 0x65, 0xb9, 0xa1, 0x5d, 0xae, 0xbb, 0xe5, 0xbf, 0x76, 0xb0, 0xa8, 0xc0,
	0xa2, 0x30, 0x2a, 0xb1, 0xc8, 0xc6, 0xa2, 0x1c, 0xcb, 0xdc, 0xbb, 0x2a, 0x58, 0xf4, 0xd0, 0xef,
	0x04, 0xd2, 0x30, 0xc1, 0x0c, 0xfb, 0x75, 0x6f, 0x2c, 0x05, 0xa5, 0x85, 0x83, 0xd2, 0x90, 0xd2,
	0xd2, 0x42, 0xa9, 0xad, 0xa1, 0x34, 0xf7, 0x8c, 0xeb, 0xdb, 0x2f, 0xd2, 0xfa, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x58, 0xe5, 0xbe, 0x5d, 0xe0, 0x02, 0x00, 0x00,
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
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
