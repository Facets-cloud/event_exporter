/*
Copyright 2020 CaiCloud, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

type CustomFilter struct {
	InvolvedObjectKind      string
	InvolvedObjectName      string
	InvolvedObjectNamespace string
	EventTypes              []string
}

// Since o.flag.Var expects pflag.Value we need to cast it to pflag.Value
// So below three methods does the job for casting
func (this *CustomFilter) String() string {
	b, _ := json.Marshal(*this)
	return string(b)
}

func (this *CustomFilter) Set(s string) error {
	return json.Unmarshal([]byte(s), this)
}

func (this *CustomFilter) Type() string {
	return "CustomFilter"
}

type Options struct {
	KubeMasterURL  string
	KubeConfigPath string
	EventType      []string
	CustomFilter   CustomFilter
	Port           int
	Version        bool
	flag           *pflag.FlagSet
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) AddFlags() {
	o.flag = pflag.NewFlagSet("", pflag.ExitOnError)
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	o.flag.AddGoFlagSet(klogFlags)

	o.flag.StringVar(&o.KubeMasterURL, "kubeMasterURL", "", "The URL of kubernetes apiserver to use as a master")
	o.flag.StringVar(&o.KubeConfigPath, "kubeConfigPath", "", "The path of kubernetes configuration file")
	o.flag.StringArrayVar(&o.EventType, "eventType", []string{"Warning"}, "List of allowed event types. Default to warning type.")
	o.flag.IntVar(&o.Port, "port", 9102, "Port to expose event metrics on")
	o.flag.BoolVar(&o.Version, "version", false, "event exporter version information")
	o.flag.Var(&o.CustomFilter, "customFilter", "Custom filters to select events. Eg: --customFilter={ \"InvolvedObjectKind\": \"Pod\", \"InvolvedObjectName\": \"cluster-autoscaler\", \"InvolvedObjectNamespace\": \"default\", \"EventTypes\": [\"Warning\", \"Normal\"]}")

	o.flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		o.flag.PrintDefaults()
	}

}

func (o *Options) Parse() error {
	return o.flag.Parse(os.Args)
}

func (o *Options) Usage() {
	o.flag.Usage()
}
