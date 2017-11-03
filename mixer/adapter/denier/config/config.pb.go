// Code generated by protoc-gen-gogo.
// source: mixer/adapter/denier/config/config.proto
// DO NOT EDIT!

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/denier/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import google_rpc "github.com/googleapis/googleapis/google/rpc"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Params struct {
	// The error to return when denying a request.
	Status google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status"`
	// The duration for which the denial is valid.
	ValidDuration time.Duration `protobuf:"bytes,2,opt,name=valid_duration,json=validDuration,stdduration" json:"valid_duration"`
	// The number of times the denial may be used.
	ValidUseCount int32 `protobuf:"varint,3,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.denier.config.Params")
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintConfig(dAtA, i, uint64(m.Status.Size()))
	n1, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)))
	n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.ValidUseCount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ValidUseCount))
	}
	return i, nil
}

func encodeFixed64Config(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Config(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = m.Status.Size()
	n += 1 + l + sovConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
	if m.ValidUseCount != 0 {
		n += 1 + sovConfig(uint64(m.ValidUseCount))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "Status", "google_rpc.Status", 1), `&`, ``, 1) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(this.ValidDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`ValidUseCount:` + fmt.Sprintf("%v", this.ValidUseCount) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUseCount", wireType)
			}
			m.ValidUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUseCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/denier/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x6d, 0x7e, 0x22, 0x14, 0x04, 0x48, 0x11, 0x88, 0xd2, 0xe1, 0x6b, 0xc5, 0x80, 0x32,
	0xd9, 0x08, 0x16, 0xe6, 0xc0, 0xc4, 0x84, 0x8a, 0x58, 0x58, 0x2a, 0x37, 0x71, 0xa3, 0x48, 0x6d,
	0x1c, 0x39, 0x36, 0x62, 0xe4, 0x08, 0x8c, 0x1c, 0x81, 0x89, 0x73, 0x64, 0xec, 0xc8, 0x04, 0xc4,
	0x2c, 0x8c, 0x3d, 0x02, 0x8a, 0xed, 0x4c, 0xfe, 0x79, 0x9e, 0xf7, 0xf3, 0x2b, 0x87, 0xf1, 0xb2,
	0x78, 0xe6, 0x92, 0xb2, 0x8c, 0x55, 0x8a, 0x4b, 0x9a, 0xf1, 0xb2, 0xe0, 0x92, 0xa6, 0xa2, 0x9c,
	0x17, 0xb9, 0x5f, 0x48, 0x25, 0x85, 0x12, 0xd1, 0x91, 0x77, 0x88, 0x73, 0x88, 0x83, 0xc3, 0xc3,
	0x5c, 0xe4, 0xc2, 0x1a, 0xb4, 0xdb, 0x39, 0x79, 0x08, 0xb9, 0x10, 0xf9, 0x82, 0x53, 0x7b, 0x9a,
	0xe9, 0x39, 0xcd, 0xb4, 0x64, 0xaa, 0x10, 0xa5, 0xe7, 0xc7, 0x9e, 0xcb, 0x2a, 0xa5, 0xb5, 0x62,
	0x4a, 0xd7, 0x0e, 0x9c, 0x7e, 0xe0, 0x30, 0xb8, 0x63, 0x92, 0x2d, 0xeb, 0xe8, 0x3c, 0x0c, 0x1c,
	0x1a, 0xe0, 0x31, 0x8e, 0x77, 0x2f, 0x22, 0xe2, 0x42, 0x44, 0x56, 0x29, 0xb9, 0xb7, 0x24, 0xd9,
	0x6a, 0xbe, 0x46, 0x68, 0xe2, 0xbd, 0xe8, 0x36, 0xdc, 0x7f, 0x62, 0x8b, 0x22, 0x9b, 0xf6, 0xaf,
	0x0d, 0x36, 0x6c, 0xf2, 0xa4, 0x4f, 0xf6, 0x75, 0xc8, 0x8d, 0x17, 0x92, 0x9d, 0x6e, 0xc0, 0xdb,
	0xf7, 0x08, 0x4f, 0xf6, 0x6c, 0xb4, 0x07, 0xd1, 0x59, 0x78, 0xe0, 0x66, 0xe9, 0x9a, 0x4f, 0x53,
	0xa1, 0x4b, 0x35, 0xd8, 0x1c, 0xe3, 0x78, 0xdb, 0x7b, 0x0f, 0x35, 0xbf, 0xee, 0x2e, 0x93, 0xab,
	0xa6, 0x05, 0xb4, 0x6a, 0x01, 0x7d, 0xb6, 0x80, 0xd6, 0x2d, 0xa0, 0x17, 0x03, 0xf8, 0xdd, 0x00,
	0x6a, 0x0c, 0xe0, 0x95, 0x01, 0xfc, 0x63, 0x00, 0xff, 0x19, 0x40, 0x6b, 0x03, 0xf8, 0xf5, 0x17,
	0xd0, 0x63, 0xe0, 0x7e, 0x6e, 0x16, 0xd8, 0x36, 0x97, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x3a, 0x7e, 0x0f, 0x83, 0x01, 0x00, 0x00,
}
