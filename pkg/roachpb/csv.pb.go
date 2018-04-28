// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/csv.proto

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// CSVOptions describe the format of csv data (delimiter, comment, etc).
type CSVOptions struct {
	// comma is an delimiter used by the CSV file; defaults to a comma.
	Comma int32 `protobuf:"varint,1,opt,name=comma" json:"comma"`
	// comment is an comment rune; zero value means comments not enabled.
	Comment int32 `protobuf:"varint,2,opt,name=comment" json:"comment"`
	// null_encoding, if not nil, is the string which identifies a NULL. Can be the empty string.
	NullEncoding *string `protobuf:"bytes,3,opt,name=null_encoding,json=nullEncoding" json:"null_encoding,omitempty"`
	// skip the first N lines of the input (e.g. to ignore column headers).
	Skip uint32 `protobuf:"varint,4,opt,name=skip" json:"skip"`
}

func (m *CSVOptions) Reset()                    { *m = CSVOptions{} }
func (m *CSVOptions) String() string            { return proto.CompactTextString(m) }
func (*CSVOptions) ProtoMessage()               {}
func (*CSVOptions) Descriptor() ([]byte, []int) { return fileDescriptorCsv, []int{0} }

func init() {
	proto.RegisterType((*CSVOptions)(nil), "cockroach.roachpb.CSVOptions")
}
func (m *CSVOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSVOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintCsv(dAtA, i, uint64(m.Comma))
	dAtA[i] = 0x10
	i++
	i = encodeVarintCsv(dAtA, i, uint64(m.Comment))
	if m.NullEncoding != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCsv(dAtA, i, uint64(len(*m.NullEncoding)))
		i += copy(dAtA[i:], *m.NullEncoding)
	}
	dAtA[i] = 0x20
	i++
	i = encodeVarintCsv(dAtA, i, uint64(m.Skip))
	return i, nil
}

func encodeVarintCsv(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CSVOptions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovCsv(uint64(m.Comma))
	n += 1 + sovCsv(uint64(m.Comment))
	if m.NullEncoding != nil {
		l = len(*m.NullEncoding)
		n += 1 + l + sovCsv(uint64(l))
	}
	n += 1 + sovCsv(uint64(m.Skip))
	return n
}

func sovCsv(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCsv(x uint64) (n int) {
	return sovCsv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CSVOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsv
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CSVOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSVOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comma", wireType)
			}
			m.Comma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comma |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			m.Comment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comment |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCsv
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.NullEncoding = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skip", wireType)
			}
			m.Skip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Skip |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCsv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCsv
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
func skipCsv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsv
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
					return 0, ErrIntOverflowCsv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCsv
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCsv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCsv
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCsv(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCsv = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsv   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("roachpb/csv.proto", fileDescriptorCsv) }

var fileDescriptorCsv = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xca, 0x4f, 0x4c,
	0xce, 0x28, 0x48, 0xd2, 0x4f, 0x2e, 0x2e, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c,
	0xce, 0x4f, 0xce, 0x06, 0x0b, 0xeb, 0x41, 0x25, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2,
	0xfa, 0x20, 0x16, 0x44, 0xa1, 0x52, 0x2f, 0x23, 0x17, 0x97, 0x73, 0x70, 0x98, 0x7f, 0x41, 0x49,
	0x66, 0x7e, 0x5e, 0xb1, 0x90, 0x14, 0x17, 0x6b, 0x72, 0x7e, 0x6e, 0x6e, 0xa2, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xab, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x10, 0x21, 0x21, 0x39, 0x2e, 0x76,
	0x10, 0x23, 0x35, 0xaf, 0x44, 0x82, 0x09, 0x49, 0x16, 0x26, 0x28, 0xa4, 0xc9, 0xc5, 0x9b, 0x57,
	0x9a, 0x93, 0x13, 0x9f, 0x9a, 0x97, 0x9c, 0x9f, 0x92, 0x99, 0x97, 0x2e, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x09, 0x56, 0xc5, 0x18, 0xc4, 0x03, 0x92, 0x72, 0x85, 0xca, 0x08, 0x49, 0x70, 0xb1, 0x14,
	0x67, 0x67, 0x16, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x42, 0xcd, 0x01, 0x8b, 0x38, 0x29, 0x9e,
	0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x37, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x14, 0x3b, 0xd4, 0x23, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0xd6, 0x09, 0x86, 0xef, 0x00, 0x00, 0x00,
}