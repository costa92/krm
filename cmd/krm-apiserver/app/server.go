package app

import (
	"fmt"
	"github.com/costa92/krm/cmd/krm-apiserver/app/options"
	"github.com/costa92/krm/pkg/version"
	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
)

const appName = "krm-apiserver"

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

type Option func(*options.ServerRunOptions)

// NewAPIServerCommand creates a new command for running the apiserver.
func NewAPIServerCommand(serverRunOptions ...Option) *cobra.Command {
	s := options.NewServerRunOptions()

	for _, opt := range serverRunOptions {
		opt(s)
	}

	cmd := &cobra.Command{
		Use: appName,
		Long: "The Kubernetes Resource Manager API server is a REST API that provides " +
			"access to the Kubernetes Resource Manager API.",
		// stop printing usage when the command errors
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			// set default options
			completedOptions, err := s.Complete()
			if err != nil {
				return err
			}
			return Run(completedOptions, genericapiserver.SetupSignalHandler())
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	fs := cmd.Flags()
	namedFlagSets := s.Flags()
	version.AddFlags(namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name(), logs.SkipLoggingConfigurationFlags())
	// The custom flag is actually not used. It is just a placeholder. In order to be consistent with
	// the kube-apiserver code, learning onex-apiserver is equivalent to learning kube-apiserver.
	//options.AddCustomGlobalFlags(namedFlagSets.FlagSet("generic"))
	// Add flags for the named flag sets.
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)
	return cmd
}

// Run runs the specified APIServer. This should never exit.
func Run(opts options.CompletedOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	klog.Infof("Version: %+v", version.Get().String())
	return nil
}
