package app

import (
	"github.com/costa92/krm/cmd/krm-apiserver/app/options"
	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	logsapi "k8s.io/component-base/logs/api/v1"
)

const appName = "krm-apiserver"

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

type Option func(*options.ServerRunOptions)

// NewAPIServerCommand creates a new command for running the apiserver.
func NewAPIServerCommand(serverRunOptions ...Option) *cobra.Command {
	return nil
}
