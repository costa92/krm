package options

import (
	controlplaneoptions "github.com/costa92/krm/internal/controlplane/apiserver/options"
	genericapiserver "k8s.io/apiserver/pkg/server"
	cliflag "k8s.io/component-base/cli/flag"
)

// ServerRunOptions contains the options for running the server.
type ServerRunOptions struct {
	*controlplaneoptions.Options

	Extra
}

type Extra struct {
	MasterCount int

	// For external resources
	ExternalPostStartHooks map[string]genericapiserver.PostStartHookFunc
}

// NewServerRunOptions returns a new ServerRunOptions.
func NewServerRunOptions() *ServerRunOptions {
	o := &ServerRunOptions{
		Options: controlplaneoptions.NewOptions(),
		Extra: Extra{
			MasterCount:            1,
			ExternalPostStartHooks: make(map[string]genericapiserver.PostStartHookFunc),
		},
	}
	return o
}

// Flags returns the flags for the ServerRunOptions.
func (o ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {
	o.Options.AddFlags(&fss)

	fs := fss.FlagSet("misc")
	// Add flags for the misc flagset.
	fs.IntVar(&o.MasterCount, "apiserver-count", o.MasterCount,
		"The number of apiservers running in the cluster, must be a positive number. (In use when --endpoint-reconciler-type=master-count is enabled.)")
	_ = fs.MarkDeprecated("apiserver-count", "apiserver-count is deprecated and will be removed in a future version.")
	return fss
}
