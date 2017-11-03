// Code generated by protoc-gen-gogo.
// source: mixer/adapter/memquota/config/config.proto
// DO NOT EDIT!

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/memquota/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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
	// The set of known quotas.
	Quotas []Params_Quota `protobuf:"bytes,1,rep,name=quotas" json:"quotas"`
	// Minimum number of seconds that deduplication is possible for a given operation.
	MinDeduplicationDuration time.Duration `protobuf:"bytes,2,opt,name=min_deduplication_duration,json=minDeduplicationDuration,stdduration" json:"min_deduplication_duration"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

type Params_Quota struct {
	// The name of the quota
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The upper limit for this quota.
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
	// The amount of time allocated quota remains valid before it is
	// automatically released. This is only meaningful for rate limit
	// quotas, otherwise the value must be zero.
	ValidDuration time.Duration `protobuf:"bytes,3,opt,name=valid_duration,json=validDuration,stdduration" json:"valid_duration"`
	// Overrides associated with this quota.
	// The first matching override is applied.
	Overrides []Params_Override `protobuf:"bytes,4,rep,name=overrides" json:"overrides"`
}

func (m *Params_Quota) Reset()                    { *m = Params_Quota{} }
func (*Params_Quota) ProtoMessage()               {}
func (*Params_Quota) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

func (m *Params_Quota) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params_Quota) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *Params_Quota) GetValidDuration() time.Duration {
	if m != nil {
		return m.ValidDuration
	}
	return 0
}

func (m *Params_Quota) GetOverrides() []Params_Override {
	if m != nil {
		return m.Overrides
	}
	return nil
}

type Params_Override struct {
	// The specific dimensions for which this override applies.
	// String representation of instance dimensions is used to check against configured dimensions.
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The upper limit for this quota.
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
	// The amount of time allocated quota remains valid before it is
	// automatically released. This is only meaningful for rate limit
	// quotas, otherwise the value must be zero.
	ValidDuration time.Duration `protobuf:"bytes,3,opt,name=valid_duration,json=validDuration,stdduration" json:"valid_duration"`
}

func (m *Params_Override) Reset()                    { *m = Params_Override{} }
func (*Params_Override) ProtoMessage()               {}
func (*Params_Override) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

func (m *Params_Override) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Params_Override) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *Params_Override) GetValidDuration() time.Duration {
	if m != nil {
		return m.ValidDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "adapter.memquota.config.Params")
	proto.RegisterType((*Params_Quota)(nil), "adapter.memquota.config.Params.Quota")
	proto.RegisterType((*Params_Override)(nil), "adapter.memquota.config.Params.Override")
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
	if len(m.Quotas) > 0 {
		for _, msg := range m.Quotas {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinDeduplicationDuration)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MinDeduplicationDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Params_Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Quota) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)))
	n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Overrides) > 0 {
		for _, msg := range m.Overrides {
			dAtA[i] = 0x22
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Params_Override) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Override) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)))
	n3, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
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
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinDeduplicationDuration)
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *Params_Quota) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
	if len(m.Overrides) > 0 {
		for _, e := range m.Overrides {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *Params_Override) Size() (n int) {
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
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
		`Quotas:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Quotas), "Params_Quota", "Params_Quota", 1), `&`, ``, 1) + `,`,
		`MinDeduplicationDuration:` + strings.Replace(strings.Replace(this.MinDeduplicationDuration.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Quota) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_Quota{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(this.ValidDuration.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`Overrides:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Overrides), "Params_Override", "Params_Override", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Override) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Params_Override{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(this.ValidDuration.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
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
			m.Quotas = append(m.Quotas, Params_Quota{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeduplicationDuration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MinDeduplicationDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *Params_Quota) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
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
			m.Overrides = append(m.Overrides, Params_Override{})
			if err := m.Overrides[len(m.Overrides)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *Params_Override) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Override: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Override: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
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
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthConfig
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthConfig
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Dimensions[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Dimensions[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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

func init() { proto.RegisterFile("mixer/adapter/memquota/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0x3f, 0x8b, 0x14, 0x31,
	0x18, 0xc6, 0x93, 0xfd, 0xc7, 0xed, 0x7b, 0xf8, 0x87, 0x70, 0xe0, 0x38, 0x60, 0x76, 0x11, 0x84,
	0xc5, 0x22, 0x03, 0x67, 0xb3, 0x1c, 0x58, 0xb8, 0xae, 0x8d, 0x08, 0xea, 0x54, 0x62, 0xb3, 0xe4,
	0x6e, 0x72, 0x43, 0x70, 0x93, 0xac, 0xd9, 0x99, 0x65, 0xaf, 0xb3, 0xb4, 0xb4, 0xb0, 0xb0, 0xb4,
	0xb0, 0xf0, 0xa3, 0x6c, 0x79, 0xa5, 0x58, 0xa8, 0x33, 0x36, 0x96, 0xf7, 0x11, 0x64, 0x92, 0x19,
	0xef, 0x10, 0xc4, 0xad, 0xae, 0x9a, 0x77, 0x92, 0xe7, 0x79, 0xde, 0x1f, 0x4f, 0xe0, 0xae, 0x92,
	0x6b, 0x61, 0x23, 0x9e, 0xf0, 0x45, 0x26, 0x6c, 0xa4, 0x84, 0x7a, 0x9d, 0x9b, 0x8c, 0x47, 0x47,
	0x46, 0x1f, 0xcb, 0xb4, 0xfe, 0xb0, 0x85, 0x35, 0x99, 0x21, 0x37, 0x6a, 0x15, 0x6b, 0x54, 0xcc,
	0x5f, 0x87, 0x34, 0x35, 0x26, 0x9d, 0x8b, 0xc8, 0xc9, 0x0e, 0xf3, 0xe3, 0x28, 0xc9, 0x2d, 0xcf,
	0xa4, 0xd1, 0xde, 0x18, 0xee, 0xa5, 0x26, 0x35, 0x6e, 0x8c, 0xaa, 0xc9, 0x9f, 0xde, 0xfe, 0xd4,
	0x85, 0xde, 0x33, 0x6e, 0xb9, 0x5a, 0x92, 0x87, 0xd0, 0x73, 0x81, 0xcb, 0x00, 0x0f, 0xdb, 0xa3,
	0xdd, 0xfd, 0x3b, 0xec, 0x1f, 0xab, 0x98, 0x37, 0xb0, 0xe7, 0xd5, 0xd9, 0xa4, 0xb3, 0xf9, 0x36,
	0x40, 0x71, 0x6d, 0x25, 0x1c, 0x42, 0x25, 0xf5, 0x2c, 0x11, 0x49, 0xbe, 0x98, 0xcb, 0x23, 0x07,
	0x30, 0x6b, 0x48, 0x82, 0xd6, 0x10, 0x8f, 0x76, 0xf7, 0x6f, 0x32, 0x8f, 0xca, 0x1a, 0x54, 0x36,
	0xad, 0x05, 0x93, 0x9d, 0x2a, 0xec, 0xc3, 0xf7, 0x01, 0x8e, 0x03, 0x25, 0xf5, 0xf4, 0x62, 0x4a,
	0xa3, 0x09, 0xbf, 0x62, 0xe8, 0xba, 0xd5, 0x84, 0x40, 0x47, 0x73, 0x25, 0x02, 0x3c, 0xc4, 0xa3,
	0x7e, 0xec, 0x66, 0x72, 0x0b, 0x40, 0xf1, 0xf5, 0x8c, 0x2b, 0x93, 0xeb, 0xcc, 0x2d, 0x6c, 0xc7,
	0x7d, 0xc5, 0xd7, 0x0f, 0xdc, 0x01, 0x79, 0x0c, 0x57, 0x57, 0x7c, 0x2e, 0x93, 0x73, 0xa6, 0xf6,
	0xf6, 0x4c, 0x57, 0x9c, 0xb5, 0xb9, 0x20, 0x4f, 0xa0, 0x6f, 0x56, 0xc2, 0x5a, 0x99, 0x88, 0x65,
	0xd0, 0x71, 0x9d, 0x8d, 0xfe, 0xd7, 0xd9, 0xd3, 0xda, 0x50, 0xd7, 0x76, 0x1e, 0x70, 0xd0, 0x79,
	0xfb, 0x71, 0x80, 0xc3, 0xf7, 0x2d, 0xd8, 0x69, 0x34, 0xe4, 0x05, 0x40, 0x22, 0x95, 0xd0, 0x4b,
	0x69, 0x74, 0xf3, 0x2a, 0xe3, 0x6d, 0x37, 0xb0, 0xe9, 0x1f, 0xeb, 0x23, 0x9d, 0xd9, 0x93, 0xf8,
	0x42, 0xd6, 0x25, 0xb6, 0x14, 0xde, 0x87, 0x6b, 0x7f, 0x91, 0x90, 0xeb, 0xd0, 0x7e, 0x25, 0x4e,
	0xea, 0x67, 0xab, 0x46, 0xb2, 0x07, 0xdd, 0x15, 0x9f, 0xe7, 0xc2, 0xa1, 0xf4, 0x63, 0xff, 0x73,
	0xd0, 0x1a, 0x63, 0x5f, 0xcb, 0x64, 0xbc, 0x29, 0x28, 0x3a, 0x2d, 0x28, 0xfa, 0x52, 0x50, 0x74,
	0x56, 0x50, 0xf4, 0xa6, 0xa4, 0xf8, 0x73, 0x49, 0xd1, 0xa6, 0xa4, 0xf8, 0xb4, 0xa4, 0xf8, 0x47,
	0x49, 0xf1, 0xaf, 0x92, 0xa2, 0xb3, 0x92, 0xe2, 0x77, 0x3f, 0x29, 0x7a, 0xd9, 0xf3, 0xad, 0x1c,
	0xf6, 0x1c, 0xea, 0xbd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x4b, 0x3f, 0x47, 0x64, 0x03,
	0x00, 0x00,
}
