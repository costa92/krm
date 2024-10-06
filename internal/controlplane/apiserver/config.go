package apiserver

import (
	"fmt"

	initializer "github.com/costa92/krm/internal/controlplane/admission/initializer"
	controlplaneoptions "github.com/costa92/krm/internal/controlplane/apiserver/options"
	"github.com/costa92/krm/pkg/generated/clientset/versioned"
	"github.com/costa92/krm/pkg/generated/informers"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/reconcilers"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storageversion"
	utilpeerproxy "k8s.io/apiserver/pkg/util/peerproxy"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/transport"
	"k8s.io/klog/v2"
	openapicommon "k8s.io/kube-openapi/pkg/common"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
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

// CreatePeerEndpointLeaseReconciler creates a apiserver endpoint lease reconciliation loop
// The peer endpoint leases are used to find network locations of apiservers for peer proxy
func CreatePeerEndpointLeaseReconciler(c genericapiserver.Config, storageFactory serverstorage.StorageFactory) (reconcilers.PeerEndpointLeaseReconciler, error) {
	ttl := controlplane.DefaultEndpointReconcilerTTL
	config, err := storageFactory.NewConfig(api.Resource("apiServerPeerIPInfo"))
	if err != nil {
		return nil, fmt.Errorf("error creating storage factory config: %w", err)
	}
	reconciler, err := reconcilers.NewPeerEndpointLeaseReconciler(config, "/peerserverleases/", ttl)
	return reconciler, err
}

func BuildPeerProxy(versionedInformer kubeinformers.SharedInformerFactory, svm storageversion.Manager,
	proxyClientCertFile string, proxyClientKeyFile string, peerCAFile string, peerAdvertiseAddress reconcilers.PeerAdvertiseAddress,
	apiServerID string, reconciler reconcilers.PeerEndpointLeaseReconciler, serializer runtime.NegotiatedSerializer) (utilpeerproxy.Interface, error) {
	if proxyClientCertFile == "" {
		return nil, fmt.Errorf("error building peer proxy handler, proxy-cert-file not specified")
	}
	if proxyClientKeyFile == "" {
		return nil, fmt.Errorf("error building peer proxy handler, proxy-key-file not specified")
	}
	// create proxy client config
	clientConfig := &transport.Config{
		TLS: transport.TLSConfig{
			Insecure:   false,
			CertFile:   proxyClientCertFile,
			KeyFile:    proxyClientKeyFile,
			CAFile:     peerCAFile,
			ServerName: "kubernetes.default.svc",
		}}

	// build proxy transport
	proxyRoundTripper, transportBuildingError := transport.New(clientConfig)
	if transportBuildingError != nil {
		klog.Error(transportBuildingError.Error())
		return nil, transportBuildingError
	}
	return utilpeerproxy.NewPeerProxyHandler(
		versionedInformer,
		svm,
		proxyRoundTripper,
		apiServerID,
		reconciler,
		serializer,
	), nil
}
