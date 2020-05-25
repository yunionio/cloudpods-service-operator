package options

import (
	"reflect"
	"testing"
	"yunion.io/x/structarg"
)

func TestParseOptions(t *testing.T) {
	want := OperatroOptions{
		BaseOptions:          structarg.BaseOptions{},
		MetricsAddr:          "127.0.0.1:8080",
		EnableLeaderElection: true,
		SyncPeriod:           15,
		AdminAuthConfig: AdminAuthConfig{
			AdminProject:  "example",
			AuthURL:       "https://10.0.1.4:30000/v3",
			AdminUsername: "goodone",
			AdminPassword: "cj5ezKSEbbZhYh7C",
		},
		AnsiblePlaybookConfig: AnsiblePlaybookConfig{
			IntervalPending: 15,
			IntervalWaiting: 20,
			Dense:           true,
		},
		VirtualMachineConfig: VirtualMachineConfig{
			IntervalPending: 5,
		},
	}
	ParseOptions("--config", "./demo.conf")
	Options.BaseOptions = want.BaseOptions
	if !reflect.DeepEqual(want, Options) {
		t.Fatalf("want: %#v\n get: %#v", want, Options)
	}
}
