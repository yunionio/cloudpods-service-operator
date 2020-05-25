package options

import (
	"fmt"
	"log"
	"os"

	"yunion.io/x/structarg"
)

type OperatroOptions struct {
	structarg.BaseOptions

	MetricsAddr          string `help:"The address the metric endpoint binds to." default:":8080"`
	EnableLeaderElection bool   `help:"Enable leader election for controller manager. Enable intensive information collection during the reconcile process." default:"false"`
	EnableWebhooks       bool   `help:"Enable webhooks for controller manager." default:"false"`
	SyncPeriod           int    `help:"The interval between two adjacent local cache refreshes. unit:m" default:"10"`

	AdminAuthConfig
	AnsiblePlaybookConfig
	VirtualMachineConfig
}

type AdminAuthConfig struct {
	Region        string `help:"Region name or ID"`
	AuthURL       string `help:"Keystone auth URL"`
	AdminUsername string `help:"Admin username"`
	AdminPassword string `help:"Admin password"`
	AdminDomain   string `help:"Admin domain"`
	AdminProject  string `help:"Admin project"`
}

type AnsiblePlaybookConfig struct {
	IntervalPending int  `json:"ap_interval_pending" help:"Reconcile interval when the state of the ansibleplaybook is pending." default:"15"`
	IntervalWaiting int  `json:"ap_interval_waiting" help:"Reconcile interval when the state of the ansibleplaybook is waiting." default:"15"`
	Dense           bool `json:"ap_dense" help:"Enable intensive information collection during the reconcile process."  default:"false"`
}

type VirtualMachineConfig struct {
	IntervalPending int `json:"vm_interval_pending" help:"Reconcile interval when the state of the virtualmachine is pending." default:"5"`
}

var Options OperatroOptions

func ParseOptions(args ...string) {
	parser, err := structarg.NewArgumentParser(&Options, "", "", "")
	if err != nil {
		log.Fatalf("Unable to define argument parser: %s.", err.Error())
	}

	if len(args) == 0 {
		args = os.Args[1:]
	}

	err = parser.ParseArgs2(args, false, false)
	if err != nil {
		log.Fatalf("Unable to parse args: %s.", err.Error())
	}

	if Options.Help {
		fmt.Println(parser.HelpString())
		os.Exit(0)
	}

	if len(Options.Config) == 0 {
		defaultConfig := "/etc/yunion/oso.conf"
		Options.Config = defaultConfig
	}

	log.Printf("Use configuration file '%s'.", Options.Config)

	err = parser.ParseFile(Options.Config)
	if err != nil {
		log.Fatalf("Unable to parse configuration file: %s.", err.Error())
	}

	parser.SetDefault()
}
