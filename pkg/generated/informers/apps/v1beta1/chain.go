package v1beta1

import (
	versioned "github.com/costa92/krm/pkg/generated/clientset/versioned"
	"github.com/costa92/krm/pkg/generated/informers/internalinterfaces"
	"github.com/costa92/krm/pkg/generated/listers/apps/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// ChainInformer provides access to a shared informer and lister for
// Chains.
type ChainInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ChainLister
}

type chainInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChainInformer constructs a new informer for Chain type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChainInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChainInformer(client, namespace, resyncPeriod, indexers, nil)
}
