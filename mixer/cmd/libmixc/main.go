package main

// go build -o libmixc.so -buildmode=c-shared main.go

import "C"

import (
	"context"
	"github.com/gogo/protobuf/proto"
	mixerpb "istio.io/api/mixer/v1"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/server"
	"istio.io/istio/pkg/log"
	"sync"
)

var (
	srv     *server.Server
	srvLock = &sync.RWMutex{}
)

//export InitModule
func InitModule() {
	var err error

	srvLock.Lock()
	defer srvLock.Unlock()
	if srv != nil {
		return
	}

	args := server.DefaultArgs()
	args.ConfigStoreURL = "fs:///Users/mjog/.kube/config"
	args.APIWorkerPoolSize = 10
	args.AdapterWorkerPoolSize = 10
	if srv, err = server.New(args); err != nil {
		log.Errorf("Unable to start server: %v", err)
	}
}

func getServer() *server.Server {
	srvLock.RLock()
	s := srv
	srvLock.RUnlock()

	return s
}

//strcpy copy to gostring
func strcpy(str string) string {
	return string(([]byte(str))[0:])
}

//export Report
func Report(attrString string) bool {

	log.Infof("Got: %d, %v", len(attrString), attrString)
	var attrs mixerpb.Attributes

	if err := proto.Unmarshal([]byte(attrString), &attrs); err != nil {
		log.Errorf("Unable to unmarshal attributes: %v", err)
		return false
	}

	var compressed *mixerpb.CompressedAttributes
	var err error
	if compressed, err = compressAttributes(&attrs); err != nil {
		log.Errorf("Unable to compress attributes: %v", err)
		return false
	}

	request := mixerpb.ReportRequest{
		Attributes: []mixerpb.CompressedAttributes{*compressed},
	}

	s := getServer()

	if s == nil {
		log.Error("unable to get report server")
		return false
	}

	_, err = s.GrpcServer.Report(context.Background(), &request)

	if err != nil {
		log.Errorf("unable to send report: %v", err)
		return false
	}

	return true
}

// compressAttributes return compressed version of attributes.
func compressAttributes(attr *mixerpb.Attributes) (*mixerpb.CompressedAttributes, error) {
	b := attribute.GetMutableBag(nil)
	for k, av := range attr.Attributes {
		switch v := av.Value.(type) {
		case *mixerpb.Attributes_AttributeValue_StringValue:
			b.Set(k, v.StringValue)
		case *mixerpb.Attributes_AttributeValue_Int64Value:
			b.Set(k, v.Int64Value)
		case *mixerpb.Attributes_AttributeValue_DoubleValue:
			b.Set(k, v.DoubleValue)
		case *mixerpb.Attributes_AttributeValue_BoolValue:
			b.Set(k, v.BoolValue)
		case *mixerpb.Attributes_AttributeValue_BytesValue:
			b.Set(k, v.BytesValue)
		case *mixerpb.Attributes_AttributeValue_TimestampValue:
			b.Set(k, v.TimestampValue)
		case *mixerpb.Attributes_AttributeValue_DurationValue:
			b.Set(k, v.DurationValue)
		case *mixerpb.Attributes_AttributeValue_StringMapValue:
			b.Set(k, v.StringMapValue.Entries)
		}
	}
	var comp mixerpb.CompressedAttributes
	b.ToProto(&comp, nil, 0)
	return &comp, nil
}

// empty main required by .so
func main() {}
