package apiserver

import (
	controlplaneoptions "github.com/costa92/krm/internal/controlplane/apiserver/options"
	"github.com/costa92/krm/pkg/generated/clientset/versioned"
	"github.com/costa92/krm/pkg/generated/informers"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	kubeinformers "k8s.io/client-go/informers"
	openapicommon "k8s.io/kube-openapi/pkg/common"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/controlplane"
)

func BuildGenericConfig(
	s controlplaneoptions.CompletedOptions,
	schema []*runtime.Scheme,
	getOpenAPIDefinitions func(ref openapicommon.ReferenceCallback) map[string]openapicommon.OpenAPIDefinition,
) (genericConfig *genericapiserver.RecommendedConfig,
	versionedInformers informers.SharedInformerFactory,
	kubeSharedInformers kubeinformers.SharedInformerFactory,
	storageFactory *serverstorage.DefaultStorageFactory,
	lastErr error) {
	// This function is a placeholder for the actual implementation.

	genericConfig = genericapiserver.NewRecommendedConfig(legacyscheme.Codecs)
	genericConfig.MergedResourceConfig = controlplane.DefaultAPIResourceConfigSource()

	if lastErr = s.GenericServerRunOptions.ApplyTo(&genericConfig.Config); lastErr != nil {
		return
	}

	s.RecommendedOptions.ExternalAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := versioned.NewForConfig(c.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
		s.InternalVersionedInformers = informerFactory
		return []admission.PluginInitializer{initializer.New(informerFactory, client)}, nil
	}
	return
}
