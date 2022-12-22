// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/define/transactionResponse.v1.proto

package define

import (
	fmt "fmt"
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

type TransactionResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// TODO define error object
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1323841a58d8ac03, []int{0}
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
	proto.RegisterType((*TransactionResponse)(nil), "assets.define.burn.TransactionResponse")
}

func init() {
	proto.RegisterFile("modules/assets/internal/transactions/define/transactionResponse.v1.proto", fileDescriptor_1323841a58d8ac03)
}

var fileDescriptor_1323841a58d8ac03 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4a, 0xc4, 0x30,
	0x1c, 0xc7, 0x9b, 0xf8, 0xbf, 0x63, 0x2d, 0x52, 0x6e, 0x88, 0x87, 0xd3, 0x4d, 0x09, 0x87, 0x5b,
	0xb7, 0x96, 0x03, 0x45, 0x39, 0x28, 0x45, 0x1c, 0xa4, 0x4b, 0xda, 0xc6, 0x5a, 0x68, 0x93, 0x23,
	0x49, 0x7d, 0x00, 0x27, 0x47, 0x1f, 0x41, 0x1c, 0x7d, 0x12, 0x71, 0xba, 0xd1, 0x51, 0xda, 0xcd,
	0xa7, 0x90, 0x5e, 0x28, 0x08, 0x75, 0xb9, 0x2d, 0xbf, 0x4f, 0x3e, 0xf9, 0xe5, 0xcb, 0xd7, 0xbe,
	0xac, 0x45, 0xde, 0x54, 0x4c, 0x11, 0xaa, 0x14, 0xd3, 0x8a, 0x94, 0x5c, 0x33, 0xc9, 0x69, 0x45,
	0xb4, 0xa4, 0x5c, 0xd1, 0x4c, 0x97, 0x82, 0x2b, 0x92, 0xb3, 0xfb, 0x92, 0xb3, 0xbf, 0x2c, 0x66,
	0x6a, 0x25, 0xb8, 0x62, 0xf8, 0x71, 0x8e, 0x57, 0x52, 0x68, 0xe1, 0x38, 0x66, 0x03, 0x36, 0x2e,
	0x4e, 0x1b, 0xc9, 0x27, 0x6e, 0x21, 0x0a, 0xb1, 0xb9, 0x26, 0xfd, 0xc9, 0x98, 0x67, 0xd7, 0xf6,
	0xf1, 0xcd, 0x78, 0x93, 0xe3, 0xd9, 0x07, 0xaa, 0xc9, 0x32, 0xa6, 0x94, 0x07, 0xa6, 0x60, 0x76,
	0x18, 0x0f, 0xa3, 0xe3, 0xda, 0x7b, 0x4c, 0x4a, 0x21, 0x3d, 0x38, 0x05, 0xb3, 0xa3, 0xd8, 0x0c,
	0xfe, 0xee, 0xf3, 0xeb, 0xa9, 0x15, 0x3e, 0xc1, 0x8f, 0x16, 0x81, 0x75, 0x8b, 0xc0, 0x77, 0x8b,
	0xc0, 0x4b, 0x87, 0xac, 0x75, 0x87, 0xac, 0xaf, 0x0e, 0x59, 0xf6, 0x49, 0x26, 0x6a, 0x3c, 0x4e,
	0x15, 0x4e, 0xfe, 0xf9, 0xfd, 0x76, 0x1e, 0xf5, 0xd9, 0x22, 0x70, 0x77, 0x55, 0x94, 0xfa, 0xa1,
	0x49, 0x71, 0x26, 0x6a, 0x12, 0xf4, 0x8f, 0x97, 0x94, 0xeb, 0x8a, 0x91, 0xa1, 0xa8, 0x2d, 0x0a,
	0x7b, 0x83, 0x3b, 0xc1, 0x22, 0x7c, 0x87, 0x4e, 0x60, 0x42, 0x2c, 0x4c, 0x88, 0xb0, 0x91, 0xfc,
	0x73, 0x80, 0x89, 0x81, 0x49, 0x0f, 0x5b, 0x88, 0xc6, 0x30, 0xb9, 0x88, 0xc2, 0x25, 0xd3, 0x34,
	0xa7, 0x9a, 0xfe, 0x40, 0xd7, 0x08, 0xbe, 0x6f, 0x0c, 0xdf, 0xef, 0x95, 0x74, 0x7f, 0x53, 0xec,
	0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xbf, 0x89, 0x27, 0xce, 0x01, 0x00, 0x00,
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
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTransactionResponseV1(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionResponseV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionResponseV1(v)
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
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTransactionResponseV1(uint64(l))
	}
	return n
}

func sovTransactionResponseV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionResponseV1(x uint64) (n int) {
	return sovTransactionResponseV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionResponseV1
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionResponseV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionResponseV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransactionResponseV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionResponseV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionResponseV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionResponseV1
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
func skipTransactionResponseV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionResponseV1
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
					return 0, ErrIntOverflowTransactionResponseV1
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
					return 0, ErrIntOverflowTransactionResponseV1
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
				return 0, ErrInvalidLengthTransactionResponseV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionResponseV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionResponseV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionResponseV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionResponseV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionResponseV1 = fmt.Errorf("proto: unexpected end of group")
)
