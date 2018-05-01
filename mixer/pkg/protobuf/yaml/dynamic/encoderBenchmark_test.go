package dynamic

import (
	"github.com/ghodss/yaml"
	"istio.io/istio/mixer/pkg/protobuf/yaml/testdata/all"
	"bytes"
	"encoding/json"
	"github.com/gogo/protobuf/jsonpb"
	protoyaml "istio.io/istio/mixer/pkg/protobuf/yaml"
	"testing"
)

func setupBench (input string, res protoyaml.Resolver, compiler Compiler)(Encoder, *foo.Simple, error) {
	data := map[string]interface{}{}
	var err error
	var ba []byte

	if ba, err = yaml.YAMLToJSON([]byte(input)); err != nil {
		return nil, nil, err
	}

	if err = json.Unmarshal(ba, &data); err != nil {
		return nil, nil, err
	}

	// build encoder
	db := NewEncoderBuilder(res, compiler, false)
	enc, err := db.Build(".foo.Simple", data)

	if err != nil {
		return nil, nil, err
	}


	ret := foo.Simple{}

	um := jsonpb.Unmarshaler{AllowUnknownFields: false}
	if err = um.Unmarshal(bytes.NewReader(ba), &ret); err != nil {
		return nil, nil, err
	}

	return enc, &ret, nil
}

func sbench() (Encoder, *foo.Simple, error){
	fds, err := protoyaml.GetFileDescSet("../testdata/all/types.descriptor")
	if err != nil {
		return nil, nil, err
	}
	//compiler := compiled.NewBuilder(statdardVocabulary())

	res := protoyaml.NewResolver(fds)

	return setupBench(everything, res, nil)
}

func BenchmarkGogo(b *testing.B) {

	b.StopTimer()
	_, f, err := sbench()
	if err != nil {
		b.Fatalf("got error: %v", err)
	}
	ba := make([]byte, 1000)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err = f.MarshalTo(ba)
		if err != nil {
			b.Fatalf("failed")
		}
	}
}

func BenchmarkDynamic_default(b *testing.B) {
	inlineFieldEncode = false
	noInterfaceFieldEncode = false
	benchmarkDynamic(b)
}

func BenchmarkDynamic_inline(b *testing.B) {
	inlineFieldEncode = true
	noInterfaceFieldEncode = false
	benchmarkDynamic(b)
}

func BenchmarkDynamic_noInterface(b *testing.B) {
	inlineFieldEncode = false
	noInterfaceFieldEncode = true
	benchmarkDynamic(b)
}

func BenchmarkDynamic_inline_noInterface(b *testing.B) {
	inlineFieldEncode = true
	noInterfaceFieldEncode = true
	benchmarkDynamic(b)
}


func benchmarkDynamic(b *testing.B) {
	b.StopTimer()
	enc, _, err := sbench()
	if err != nil {
		b.Fatalf("got error: %v", err)
	}
	ba := make([]byte, 0, 1000)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ba = ba[0:0]
		ba, err = enc.Encode(nil, ba)
		if err != nil {
			b.Fatalf("failed")
		}
	}
}

func Test_bench(t *testing.T)  {
	enc, _, err := sbench()
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	ba := make([]byte, 0, 1000)
	ba = ba[0:0]
	ba, err = enc.Encode(nil, ba)
	if err != nil {
		t.Fatalf("failed")
	}
}