package v1

import (
	"context"
	scheme "github.com/costa92/krm/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// CustomResourceDefinitionsGetter has a method to return a CustomResourceDefinitionInterface.
type CustomResourceDefinitionsGetter interface {
	CustomResourceDefinitions() CustomResourceDefinitionInterface
}

// CustomResourceDefinitionInterface has methods to work with CustomResourceDefinition resources.
type CustomResourceDefinitionInterface interface {
	Create(ctx context.Context, customResourceDefinition *v1.CustomResourceDefinition, opts metav1.CreateOptions) (*v1.CustomResourceDefinition, error)
}

type customResourceDefinitions struct {
	client rest.Interface
}

func newCustomResourceDefinitions(c *ApiExtensionsV1Client) *customResourceDefinitions {
	return &customResourceDefinitions{
		client: c.RESTClient(),
	}
}

func (c *customResourceDefinitions) Create(ctx context.Context, customResourceDefinition *v1.CustomResourceDefinition, opts metav1.CreateOptions) (*v1.CustomResourceDefinition, error) {
	result := &v1.CustomResourceDefinition{}
	err := c.client.Post().
		Resource("customresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customResourceDefinition).
		Do(ctx).
		Into(result)
	return result, err
}
