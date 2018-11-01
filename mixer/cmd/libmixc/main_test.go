package main

import (
	"os"
	"testing"
)

func TestReport(t *testing.T) {
	os.Setenv("CONFIG_STORE_URL", "k8s:///Users/mjog/.kube/config")
	InitModule()
}




