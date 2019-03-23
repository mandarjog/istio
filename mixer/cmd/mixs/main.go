// Copyright 2016 Istio Authors
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

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"istio.io/istio/mixer/adapter"
	"istio.io/istio/mixer/cmd/mixs/cmd"
	"istio.io/istio/mixer/cmd/shared"
	adptr "istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/pkg/template"
	generatedTmplRepo "istio.io/istio/mixer/template"
)

func supportedTemplates() map[string]template.Info {
	return generatedTmplRepo.SupportedTmplInfo
}

func supportedAdapters() []adptr.InfoFn {
	return adapter.Inventory()
}

func setMutexProfile() {
	sgm := os.Getenv("GOMUTEXPROFILE")
	if sgm == "" {
		return
	}

	gm, err := strconv.ParseInt(sgm, 10, 32)
	if err != nil {
		return
	}
	fmt.Printf("GOMUTEXPROFILE = %v", gm)
	runtime.SetBlockProfileRate(int(gm))
}

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[1:], supportedTemplates(), supportedAdapters(), shared.Printf, shared.Fatalf)
	setMutexProfile()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
