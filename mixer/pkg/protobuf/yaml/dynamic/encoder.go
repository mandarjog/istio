// Copyright 2018 Istio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dynamic

import (
	"istio.io/istio/mixer/pkg/lang/compiled"
	"fmt"
	"istio.io/istio/mixer/pkg/attribute"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"sort"
	"istio.io/istio/mixer/pkg/protobuf/yaml"
)


type (
	// Encoder transforms yaml that represents protobuf data into []byte
	// The yaml representation may have dynamic content
	Encoder interface {
		Encode(bag attribute.Bag, ba []byte) ([]byte, error)
	}

	messageEncoder struct {
		// skipEncodeLength skip encoding length of the message in the output
		// should be true only for top level message.
		skipEncodeLength bool

		// fields of the message.
		fields []*field
	}

	field struct {
		// proto key  -- EncodeVarInt ((field_number << 3) | wire_type)
		protoKey []byte

		// encodedData is available if the entire field can be encoded
		// at compile time.
		encodedData []byte

		// encoder is needed if encodedData is not available.
		encoder Encoder

		// number fields are sorted by this.
		number int

		// name for debug.
		name string
	}

	DynamicEncoderBuilder struct {
		msgName string
		resolver yaml.Resolver
		data map[interface{}]interface{}
		compiler compiled.Compiler
		skipUnknown bool
	}
)

// NewDynamicEncoderBuilder creates a DynamicEncoderBuilder.
func NewDynamicEncoderBuilder(msgName string, resolver yaml.Resolver, data map[interface{}]interface{},
	compiler compiled.Compiler, skipUnknown bool) *DynamicEncoderBuilder {
	return &DynamicEncoderBuilder{
		msgName: msgName,
		resolver: resolver,
		data: data,
		compiler: compiler,
		skipUnknown: skipUnknown}
}

// Build builds a DynamicEncoder
func (c DynamicEncoderBuilder) Build() (Encoder, error) {
	m := c.resolver.ResolveMessage(c.msgName)
	if m == nil {
		return nil, fmt.Errorf("cannot resolve message '%s'", c.msgName)
	}

	return c.buildMessage(m, c.data, true)
}

func (c DynamicEncoderBuilder) buildMessage(md *descriptor.DescriptorProto, data map[interface{}]interface{}, skipEncodeLength bool) (Encoder, error) {
	var err error
	var ok bool

	me := messageEncoder{
		skipEncodeLength: skipEncodeLength,
	}

	for kk, v := range data {
		var k string

		if k, ok = kk.(string); !ok {
			return nil, fmt.Errorf("error processing message '%s':%v got %T want string", md.GetName(), kk, kk)
		}

		fd := yaml.FindFieldByName(md, k)
		if fd == nil {
			if c.skipUnknown {
				continue
			}
			return nil, fmt.Errorf("field '%s' not found in message '%s'", k, md.GetName())
		}

		repeated := fd.IsRepeated()
		//packed := fd.IsPacked() || fd.IsPacked3()

		if fd.IsScalar() || fd.IsString() {
			if repeated {  // TODO
			}

			fld := makeField(fd)
			if fld.encodedData, err = EncodePrimitive(v, fd.GetType(), fld.encodedData); err != nil {
				return nil, fmt.Errorf("unable to encode: %v. %v", k, err)
			}
			me.fields = append(me.fields, fld)

		} else if *fd.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {  // MESSAGE, or map
			m := c.resolver.ResolveMessage(fd.GetTypeName())
			if m == nil {
				return nil, fmt.Errorf("unable to resolve message '%s'", fd.GetTypeName())
			}

			var ma []interface{}
			if m.GetOptions().GetMapEntry() {  // this is a Map
				ma, err = convertMapToMapentry(v)
				if err != nil {
					return nil, fmt.Errorf("unable to process: %v, %v", fd, err)
				}
			} else if repeated { // map entry is always repeated.
				ma, ok = v.([]interface{})
				if !ok {
					return nil, fmt.Errorf("unable to process: %v, got %T, want: []interface{}", fd, v)
				}
			} else {
				ma = []interface{}{v}
			}

			// now maps, messages and repeated maps all look the same.
			for _, vv := range ma {
				var vq map[interface{}]interface{}
				if vq, ok = vv.(map[interface{}]interface{}); !ok {
					return nil, fmt.Errorf("unable to process: %v, got %T, want: map[string]interface{}", fd, vv)
				}

				var de Encoder
				if de, err = c.buildMessage(m, vq, false); err != nil {
					return nil, fmt.Errorf("unable to process: %v, %v", fd, err)
				}
				fld := makeField(fd)
				fld.encoder = de
				me.fields = append(me.fields, fld)// create new fields ...
			}
		} else {
			return nil, fmt.Errorf("field not supported '%v'", fd)
		}

	}

	// sorting is recommended, but not required.
	sort.Slice(me.fields, func(i, j int) bool {
		return me.fields[i].number < me.fields[j].number
	})

	return me, nil
}


func convertMapToMapentry(data interface{}) ([]interface{}, error){
	md, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("incorrect map type:%T, want:map[interface{}]interface{}", data)
	}
	res := make([]interface{}, 0, len(md))
	for k, v := range md {
		res = append(res, map[interface{}]interface{}{
			"key": k,
			"value": v,
		})
	}
	return res, nil
}

func extendSlice(ba []byte, n int) []byte {
	for k:=0; k<n; k++ {
		ba = append(ba, 0xff)
	}
	return ba
}

// expected length of the varint encoded word
// 2 byte words represent 2 ** 14 = 16K bytes
// If message lenght is more, it involves an array copy
const varLength  = 2


// encode message including length of the message into []byte
func (m messageEncoder) Encode(bag attribute.Bag, ba []byte) ([]byte, error) {
	var err error

	if m.skipEncodeLength {
		return m.encodeNoLength(bag, ba)
	}
	l0 := len(ba)

	// #pragma inline reserve varLength bytes
	ba = extendSlice(ba, varLength)

	l1 := len(ba)

	if ba, err = m.encodeNoLength(bag, ba); err != nil {
		return nil, err
	}

	length := len(ba) - l1

	diff := proto.SizeVarint(uint64(length)) - varLength
	// move data forward because we need more than varLength bytes
	if diff > 0 {
		ba = extendSlice(ba, diff)
		// shift data down. This should rarely occur.
		copy(ba[l1+diff:], ba[l1:])
	}

	// ignore return value. EncodeLength is writing in the middle of the array.
	_ = EncodeVarintZeroExtend(ba[l0:l0], uint64(length), varLength)

	return ba, nil
}


func (m messageEncoder) encodeNoLength(bag attribute.Bag, ba []byte) ([]byte, error) {
	var err error
	for _, f  := range m.fields {
		ba, err = f.Encode(bag, ba)
		if err != nil {
			return nil, err
		}
	}
	return ba, nil
}

type evalEncoder struct {
	//TODO handle google.proto.Value type
	useValueType bool
	etype descriptor.FieldDescriptorProto_Type
	ex compiled.Expression
}

func (e evalEncoder) Encode(bag attribute.Bag, ba []byte) ([]byte, error) {
	v, err := e.ex.Evaluate(bag)
	if err != nil {
		return nil, err
	}

	return EncodePrimitive(v, e.etype, ba)
}

func (f field) Encode(bag attribute.Bag, ba []byte) ([]byte, error) {
	ba = append(ba, f.protoKey...)

	// Varint, 64-bit, 32-bit are directly encoded
	// should take care of its own length
	if f.encodedData != nil {
		return append(ba, f.encodedData...), nil
	}

	// The following call happens when
	// 1. value requires expression evaluation.
	// 2. field is of type map
	// 3. field is of Message
	// In all cases Encode function must correctly set Length.
	return f.encoder.Encode(bag, ba)
}
