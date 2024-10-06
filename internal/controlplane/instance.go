package controlplane

import (
	"net/http"
	"reflect"
	"time"

	"github.com/costa92/krm/pkg/apiserver/storage"
	"github.com/costa92/krm/pkg/generated/informers"
	peerreconcilers "k8s.io/apiserver/pkg/reconcilers"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	utilpeerproxy "k8s.io/apiserver/pkg/util/peerproxy"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/kubernetes/pkg/routes"
)

type ExternalSharedInformerFactory interface {
	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})
	// WaitForCacheSync blocks until all started informers' caches were synced
	// or the stop channel gets closed.
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

// ExtraConfig defines extra configuration for the onex-apiserver.
type ExtraConfig struct {
	// Place you custom config here.
	APIResourceConfigSource serverstorage.APIResourceConfigSource
	StorageFactory          serverstorage.StorageFactory
	EventTTL                time.Duration
	EnableLogsSupport       bool
	ProxyTransport          *http.Transport

	// PeerProxy, if not nil, sets proxy transport between kube-apiserver peers for requests
	// that can not be served locally
	PeerProxy utilpeerproxy.Interface
	// PeerEndpointLeaseReconciler updates the peer endpoint leases
	PeerEndpointLeaseReconciler peerreconcilers.PeerEndpointLeaseReconciler

	// For external resources and rest storage providers.
	ExternalRESTStorageProviders []storage.RESTStorageProvider
	//ExternalGroupResources       []schema.GroupResource

	// Number of masters running; all masters must be started with the
	// same value for this field. (Numbers > 1 currently untested.)
	MasterCount int

	KubeVersionedInformers     kubeinformers.SharedInformerFactory
	InternalVersionedInformers informers.SharedInformerFactory
	ExternalVersionedInformers ExternalSharedInformerFactory
	ExternalPostStartHooks     map[string]genericapiserver.PostStartHookFunc
}
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   ExtraConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package.
type CompletedConfig struct {
	*completedConfig
}

// Instance is a control plane instance.
type Instance struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

// Complete returns a CompletedConfig.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{&completedConfig{
		GenericConfig: c.GenericConfig.Complete(),
		ExtraConfig:   c.ExtraConfig,
	}}
}

func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*Instance, error) {
	s, err := c.GenericConfig.New("krm-apiserver", delegationTarget)
	if err != nil {
		return nil, err

	}

	if c.ExtraConfig.EnableLogsSupport {
		// Add log support.
		routes.Logs{}.Install(s.Handler.GoRestfulContainer)
	}

	m := &Instance{
		GenericAPIServer: s,
	}
	return m, nil
}
