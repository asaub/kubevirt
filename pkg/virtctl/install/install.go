package install

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"kubevirt.io/client-go/kubecli"
	"kubevirt.io/kubevirt/pkg/util/cluster"
	"kubevirt.io/kubevirt/pkg/virtctl/templates"
)

func NewInstallCommand(clientConfig clientcmd.ClientConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install kubevirt to destination cluster.",
		Example: usage(),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := command{clientConfig: clientConfig}
			return c.run()
		},
	}

	cmd.SetUsageTemplate(templates.UsageTemplate())
	return cmd
}

func usage() string {
	usage := "  # Install kubevirt in the kubernetes cluster:\n"
	usage += "  {{ProgramName}} permitted-devices"
	return usage
}

type command struct {
	clientConfig clientcmd.ClientConfig
}

func (c *command) run() error {
	glog.Infoln("Hello World")

	virtCli, err := kubecli.GetKubevirtClientFromClientConfig(c.clientConfig)
	if err != nil {
		return err
	}

	// Check if cluster is OS

	isOS, err := cluster.IsOnOpenShift(virtCli)
	if err != nil {
		return err
	}
	if isOS {
		// This command would be destructive on Openshift.
		fmt.Println("Detected an openshift cluster.  This command doesn't work with openshift currently.")
		return nil
	}

	// Compile resources

	// Install resources
	/* virtCli. ... */

	// Tell users they may want to install more things

	return nil
}
