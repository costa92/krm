package options

import (
	"net"

	"github.com/costa92/krm/internal/controlplane"
	controlplaneoptions "github.com/costa92/krm/internal/controlplane/apiserver/options"
	"github.com/costa92/krm/pkg/apiserver/storage"
	genericapiserver "k8s.io/apiserver/pkg/server"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/kube-openapi/pkg/common"
)

// ServerRunOptions contains the options for running the server.
type ServerRunOptions struct {
	*controlplaneoptions.Options

	Extra
}

type Extra struct {
	MasterCount int
	// In the future, perhaps an "onexlet" will be added, similar to the "kubelet".
	// OnexletConfig onexletclient.OnexletClientConfig
	APIServerServiceIP     net.IP
	EndpointReconcilerType string

	// For external resources
	ExternalRESTStorageProviders []storage.RESTStorageProvider
	ExternalVersionedInformers   controlplane.ExternalSharedInformerFactory
	ExternalPostStartHooks       map[string]genericapiserver.PostStartHookFunc
	GetOpenAPIDefinitions        common.GetOpenAPIDefinitions
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
