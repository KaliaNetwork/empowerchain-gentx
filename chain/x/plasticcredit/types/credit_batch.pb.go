// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/plasticcredit/credit_batch.proto

package types

import (
	fmt "fmt"
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

type CreditBatch struct {
	BatchDenom     string        `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	ProjectId      uint64        `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TotalAmount    *CreditAmount `protobuf:"bytes,3,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CreditData     []*ProvenData `protobuf:"bytes,4,rep,name=credit_data,json=creditData,proto3" json:"credit_data,omitempty"`
	AdditionalData []*ProvenData `protobuf:"bytes,5,rep,name=additional_data,json=additionalData,proto3" json:"additional_data,omitempty"`
}

func (m *CreditBatch) Reset()         { *m = CreditBatch{} }
func (m *CreditBatch) String() string { return proto.CompactTextString(m) }
func (*CreditBatch) ProtoMessage()    {}
func (*CreditBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_013616f774bc0fd4, []int{0}
}
func (m *CreditBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreditBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreditBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreditBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditBatch.Merge(m, src)
}
func (m *CreditBatch) XXX_Size() int {
	return m.Size()
}
func (m *CreditBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CreditBatch proto.InternalMessageInfo

func (m *CreditBatch) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *CreditBatch) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *CreditBatch) GetTotalAmount() *CreditAmount {
	if m != nil {
		return m.TotalAmount
	}
	return nil
}

func (m *CreditBatch) GetCreditData() []*ProvenData {
	if m != nil {
		return m.CreditData
	}
	return nil
}

func (m *CreditBatch) GetAdditionalData() []*ProvenData {
	if m != nil {
		return m.AdditionalData
	}
	return nil
}

func init() {
	proto.RegisterType((*CreditBatch)(nil), "empowerchain.plasticcredit.CreditBatch")
}

func init() {
	proto.RegisterFile("empowerchain/plasticcredit/credit_batch.proto", fileDescriptor_013616f774bc0fd4)
}

var fileDescriptor_013616f774bc0fd4 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcd, 0x4a, 0x03, 0x31,
	0x18, 0x6c, 0xda, 0x2a, 0x34, 0x11, 0x85, 0x9c, 0x96, 0x82, 0x71, 0xf1, 0x20, 0x7b, 0xd0, 0x14,
	0xea, 0xcd, 0x9b, 0xb5, 0x20, 0xe2, 0x41, 0x59, 0x3c, 0x79, 0x59, 0xd2, 0x24, 0xd8, 0x48, 0x77,
	0x13, 0xb6, 0x5f, 0xfd, 0x79, 0x0b, 0x1f, 0xc6, 0x87, 0xf0, 0xd8, 0xa3, 0x47, 0x69, 0x5f, 0x44,
	0x36, 0x29, 0xe8, 0x0a, 0x56, 0x3c, 0x25, 0xdf, 0x64, 0xbe, 0xc9, 0x0c, 0x83, 0x8f, 0x74, 0xee,
	0xec, 0xa3, 0x2e, 0xe5, 0x58, 0x98, 0xa2, 0xe7, 0x26, 0x62, 0x0a, 0x46, 0xca, 0x52, 0x2b, 0x03,
	0xbd, 0x70, 0x64, 0x23, 0x01, 0x72, 0xcc, 0x5d, 0x69, 0xc1, 0xd2, 0xee, 0x77, 0x3a, 0xaf, 0xd1,
	0xbb, 0xfc, 0x6f, 0x29, 0x91, 0xdb, 0x59, 0x01, 0x41, 0xab, 0x7b, 0xb8, 0x86, 0xef, 0x4a, 0xfb,
	0xa0, 0x8b, 0x4c, 0x09, 0x10, 0x81, 0xbd, 0xff, 0xda, 0xc4, 0xe4, 0xcc, 0x3f, 0x0e, 0x2a, 0x3f,
	0x74, 0x0f, 0x13, 0x6f, 0x2c, 0x53, 0xba, 0xb0, 0x79, 0x84, 0x62, 0x94, 0x74, 0x52, 0xec, 0xa1,
	0x61, 0x85, 0xd0, 0x5d, 0x8c, 0x5d, 0x69, 0xef, 0xb5, 0x84, 0xcc, 0xa8, 0xa8, 0x19, 0xa3, 0xa4,
	0x9d, 0x76, 0x56, 0xc8, 0x85, 0xa2, 0x97, 0x78, 0x0b, 0x2c, 0x88, 0xc9, 0xca, 0x53, 0xd4, 0x8a,
	0x51, 0x42, 0xfa, 0x09, 0xff, 0x3d, 0x20, 0x0f, 0xdf, 0x9f, 0x7a, 0x7e, 0x4a, 0xfc, 0x76, 0x18,
	0xe8, 0x39, 0x26, 0xab, 0x84, 0x95, 0xe3, 0xa8, 0x1d, 0xb7, 0x12, 0xd2, 0x3f, 0x58, 0xa7, 0x75,
	0xed, 0x03, 0x0e, 0x05, 0x88, 0x14, 0x07, 0xa8, 0xba, 0xd3, 0x2b, 0xbc, 0x23, 0x94, 0x32, 0x60,
	0x6c, 0x21, 0x26, 0x41, 0x6c, 0xe3, 0x5f, 0x62, 0xdb, 0x5f, 0xeb, 0xd5, 0x3c, 0xb8, 0x79, 0x5b,
	0x30, 0x34, 0x5f, 0x30, 0xf4, 0xb1, 0x60, 0xe8, 0x65, 0xc9, 0x1a, 0xf3, 0x25, 0x6b, 0xbc, 0x2f,
	0x59, 0xe3, 0xf6, 0xe4, 0xce, 0xc0, 0x78, 0x36, 0xe2, 0xd2, 0xe6, 0xbd, 0x5a, 0x13, 0xb5, 0xe1,
	0xe9, 0x47, 0x31, 0xf0, 0xec, 0xf4, 0x74, 0xb4, 0xe9, 0x3b, 0x39, 0xfe, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x6d, 0xf0, 0xc7, 0x3e, 0x02, 0x00, 0x00,
}

func (m *CreditBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreditBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreditBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdditionalData) > 0 {
		for iNdEx := len(m.AdditionalData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdditionalData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCreditBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.CreditData) > 0 {
		for iNdEx := len(m.CreditData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCreditBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.TotalAmount != nil {
		{
			size, err := m.TotalAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCreditBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ProjectId != 0 {
		i = encodeVarintCreditBatch(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintCreditBatch(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCreditBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovCreditBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreditBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovCreditBatch(uint64(l))
	}
	if m.ProjectId != 0 {
		n += 1 + sovCreditBatch(uint64(m.ProjectId))
	}
	if m.TotalAmount != nil {
		l = m.TotalAmount.Size()
		n += 1 + l + sovCreditBatch(uint64(l))
	}
	if len(m.CreditData) > 0 {
		for _, e := range m.CreditData {
			l = e.Size()
			n += 1 + l + sovCreditBatch(uint64(l))
		}
	}
	if len(m.AdditionalData) > 0 {
		for _, e := range m.AdditionalData {
			l = e.Size()
			n += 1 + l + sovCreditBatch(uint64(l))
		}
	}
	return n
}

func sovCreditBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCreditBatch(x uint64) (n int) {
	return sovCreditBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreditBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCreditBatch
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
			return fmt.Errorf("proto: CreditBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreditBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TotalAmount == nil {
				m.TotalAmount = &CreditAmount{}
			}
			if err := m.TotalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditData = append(m.CreditData, &ProvenData{})
			if err := m.CreditData[len(m.CreditData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalData = append(m.AdditionalData, &ProvenData{})
			if err := m.AdditionalData[len(m.AdditionalData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCreditBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCreditBatch
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
func skipCreditBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCreditBatch
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
					return 0, ErrIntOverflowCreditBatch
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
					return 0, ErrIntOverflowCreditBatch
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
				return 0, ErrInvalidLengthCreditBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCreditBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCreditBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCreditBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCreditBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCreditBatch = fmt.Errorf("proto: unexpected end of group")
)
