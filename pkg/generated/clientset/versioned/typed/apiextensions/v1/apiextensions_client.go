package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/kube-aggregator/pkg/apiserver/scheme"
	"net/http"
)

type ApiExtensionsV1Interface interface {
	RESTClient() rest.Interface
	CustomResourceDefinitionsGetter
}

// ApiExtensionsV1Client is used to interact with features provided by the apiextensions.k8s.io group.
type ApiExtensionsV1Client struct {
	restClient rest.Interface
}

// CustomResourceDefinitions returns a CustomResourceDefinitionInterface
func (c *ApiExtensionsV1Client) CustomResourceDefinitions() CustomResourceDefinitionInterface {
	return newCustomResourceDefinitions(c)
}

// New creates a new ApiExtensionsV1Client for the given RESTClient.
func New(c rest.Interface) *ApiExtensionsV1Client {
	return &ApiExtensionsV1Client{c}
}

func NewForConfig(c *rest.Config) (*ApiExtensionsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ApiExtensionsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}

	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ApiExtensionsV1Client{client}, nil
}

// NewForConfigOrDie creates a new ApiExtensionsV1Client for the given config and
func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}

// RESTClient New creates a new ApiExtensionsV1Client for the given RESTClient.
func (c *ApiExtensionsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
