package versioned

import (
	apiextensionsv1 "github.com/costa92/krm/pkg/generated/clientset/versioned/typed/apiextensions/v1"
	discovery "k8s.io/client-go/discovery"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApiextensionsV1() apiextensionsv1.ApiExtensionsV1Interface
}

type Clientset struct {
	*discovery.DiscoveryClient
	apiextensionsV1 *apiextensionsv1.ApiExtensionsV1Client
}
