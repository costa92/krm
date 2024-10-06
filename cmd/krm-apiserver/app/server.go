package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/costa92/krm/cmd/krm-apiserver/app/options"
	"github.com/costa92/krm/internal/controlplane"
	controlplaneapiserver "github.com/costa92/krm/internal/controlplane/apiserver"
	"github.com/costa92/krm/pkg/version"
	"github.com/spf13/cobra"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	extensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	genericapifilters "k8s.io/apiserver/pkg/endpoints/filters"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/apiserver/pkg/util/notfoundhandler"
	"k8s.io/apiserver/pkg/util/webhook"
	kubeinformers "k8s.io/client-go/informers"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	aggregatorscheme "k8s.io/kube-aggregator/pkg/apiserver/scheme"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/features"
)

const appName = "krm-apiserver"

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

type Option func(*options.ServerRunOptions)

// NewAPIServerCommand creates a new command for running the apiserver.
func NewAPIServerCommand(serverRunOptions ...Option) *cobra.Command {
	s := options.NewServerRunOptions()

	for _, opt := range serverRunOptions {
		opt(s)
	}

	cmd := &cobra.Command{
		Use: appName,
		Long: "The Kubernetes Resource Manager API server is a REST API that provides " +
			"access to the Kubernetes Resource Manager API.",
		// stop printing usage when the command errors
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			version.PrintAndExitIfRequested(appName)
			fs := cmd.Flags()

			// Activate logging as soon as possible, after that
			// show flags with the final logging configuration.
			if err := logsapi.ValidateAndApply(s.Logs, utilfeature.DefaultFeatureGate); err != nil {
				return err
			}
			cliflag.PrintFlags(fs)

			// set default options
			completedOptions, err := s.Complete()
			if err != nil {
				return err
			}
			// validate options
			if errs := completedOptions.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}
			// add feature enablement metrics
			utilfeature.DefaultMutableFeatureGate.AddMetrics()
			// run the server
			return Run(cmd.Context(), completedOptions, genericapiserver.SetupSignalHandler())
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	fs := cmd.Flags()
	namedFlagSets := s.Flags()
	version.AddFlags(namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name(), logs.SkipLoggingConfigurationFlags())
	// The custom flag is actually not used. It is just a placeholder. In order to be consistent with
	// the kube-apiserver code, learning onex-apiserver is equivalent to learning kube-apiserver.
	options.AddCustomGlobalFlags(namedFlagSets.FlagSet("generic"))
	// Add flags for the named flag sets.
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)
	return cmd
}

// Run runs the specified APIServer. This should never exit.
func Run(ctx context.Context, opts options.CompletedOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	klog.Infof("Version: %+v", version.Get().String())
	config, err := NewConfig(opts)
	if err != nil {
		return err
	}
	// Complete the configuration
	// 完成配置
	completed, err := config.Complete()
	if err != nil {
		return err
	}
	server, err := CreateServerChain(completed)
	if err != nil {
		return err
	}
	// Start the server
	prepared, err := server.PrepareRun()
	if err != nil {
		return err
	}
	// Run the server
	return prepared.Run(stopCh)
}

func CreateServerChain(config CompletedConfig) (*aggregatorapiserver.APIAggregator, error) {
	notFoundHandler := notfoundhandler.New(config.ControlPlane.GenericConfig.Serializer, genericapifilters.NoMuxAndDiscoveryIncompleteKey)
	apiExtensionsServer, err := config.ApiExtensions.New(genericapiserver.NewEmptyDelegateWithCustomHandler(notFoundHandler))
	if err != nil {
		return nil, err
	}

	crdAPIEnabled := config.ApiExtensions.GenericConfig.MergedResourceConfig.ResourceEnabled(apiextensionsv1.SchemeGroupVersion.WithResource("customresourcedefinitions"))

	// Create the API aggregator server
	krmAPIServer, err := config.ControlPlane.New(apiExtensionsServer.GenericAPIServer)
	if err != nil {
		return nil, err
	}

	// aggregator comes last in the chain
	aggregatorServer, err := createAggregatorServer(config.Aggregator, krmAPIServer.GenericAPIServer, apiExtensionsServer.Informers, crdAPIEnabled)

	return aggregatorServer, nil
}

func CreateProxyTransport() *http.Transport {
	var proxyDialerFn utilnet.DialFunc
	// Proxying to pods and services is IP-based... don't expect to be able to verify the hostname
	proxyTLSClientConfig := &tls.Config{InsecureSkipVerify: true}
	proxyTransport := utilnet.SetTransportDefaults(&http.Transport{
		DialContext:     proxyDialerFn,
		TLSClientConfig: proxyTLSClientConfig,
	})
	return proxyTransport
}

// CreateKrmAPIServerConfig creates the configuration for the KRM API server.
func CreateKrmAPIServerConfig(opts options.CompletedOptions) (
	*controlplane.Config,
	aggregatorapiserver.ServiceResolver,
	error,
) {
	proxyTransport := CreateProxyTransport()

	genericConfig, _, kubeSharedInformers, storageFactory, err := controlplaneapiserver.BuildGenericConfig(
		opts.CompletedOptions,
		[]*runtime.Scheme{legacyscheme.Scheme, extensionsapiserver.Scheme, aggregatorscheme.Scheme},
		opts.GetOpenAPIDefinitions,
	)
	fmt.Println("CreateKrmAPIServerConfig storageFactory:", storageFactory)
	if err != nil {
		return nil, nil, err
	}

	opts.Metrics.Apply()

	config := &controlplane.Config{
		GenericConfig: genericConfig,
		ExtraConfig: controlplane.ExtraConfig{
			APIResourceConfigSource: storageFactory.APIResourceConfigSource,
			StorageFactory:          storageFactory,
			EventTTL:                opts.EventTTL,
			EnableLogsSupport:       opts.EnableLogsHandler,
			ProxyTransport:          proxyTransport,
			//ExternalGroupResources: opts.ExternalGroupResources,
			ExternalRESTStorageProviders: opts.ExternalRESTStorageProviders,
			MasterCount:                  opts.MasterCount,
			//VersionedInformers:           opts.SharedInformerFactory,
			// Here we will use the config file of "onex" to create a client-go informers.
			KubeVersionedInformers:     kubeSharedInformers,
			InternalVersionedInformers: opts.InternalVersionedInformers,
			ExternalVersionedInformers: opts.ExternalVersionedInformers,
			ExternalPostStartHooks:     opts.ExternalPostStartHooks,
		},
	}

	if utilfeature.DefaultFeatureGate.Enabled(features.UnknownVersionInteroperabilityProxy) {
		config.ExtraConfig.PeerEndpointLeaseReconciler, err = controlplaneapiserver.CreatePeerEndpointLeaseReconciler(genericConfig.Config, storageFactory)
		if err != nil {
			return nil, nil, err
		}
		// build peer proxy config only if peer ca file exists
		if opts.PeerCAFile != "" {
			config.ExtraConfig.PeerProxy, err = controlplaneapiserver.BuildPeerProxy(
				kubeSharedInformers,
				genericConfig.StorageVersionManager,
				opts.ProxyClientCertFile,
				opts.ProxyClientKeyFile,
				opts.PeerCAFile,
				opts.PeerAdvertiseAddress,
				genericConfig.APIServerID,
				config.ExtraConfig.PeerEndpointLeaseReconciler,
				config.GenericConfig.Serializer,
			)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	/* UPDATEME: when add authentication features.
	clientCAProvider, err := opts.Authentication.ClientCert.GetClientCAContentProvider()
	if err != nil {
		return nil, nil, nil, err
	}
	config.ExtraConfig.ClusterAuthenticationInfo.ClientCA = clientCAProvider

	requestHeaderConfig, err := opts.Authentication.RequestHeader.ToAuthenticationRequestHeaderConfig()
	if err != nil {
		return nil, nil, nil, err
	}
	if requestHeaderConfig != nil {
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderCA = requestHeaderConfig.CAContentProvider
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderAllowedNames = requestHeaderConfig.AllowedClientNames
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderExtraHeaderPrefixes = requestHeaderConfig.ExtraHeaderPrefixes
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderGroupHeaders = requestHeaderConfig.GroupHeaders
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderUsernameHeaders = requestHeaderConfig.UsernameHeaders
	}
	*/

	serviceResolver := buildServiceResolver(opts.EnableAggregatorRouting, genericConfig.LoopbackClientConfig.Host, kubeSharedInformers)

	return config, serviceResolver, nil
}

var testServiceResolver webhook.ServiceResolver

func buildServiceResolver(enabledAggregatorRouting bool, hostname string, informer kubeinformers.SharedInformerFactory) webhook.ServiceResolver {
	if testServiceResolver != nil {
		return testServiceResolver
	}

	var serviceResolver webhook.ServiceResolver
	if enabledAggregatorRouting {
		serviceResolver = aggregatorapiserver.NewEndpointServiceResolver(
			informer.Core().V1().Services().Lister(),
			informer.Core().V1().Endpoints().Lister(),
		)
	} else {
		serviceResolver = aggregatorapiserver.NewClusterIPServiceResolver(
			informer.Core().V1().Services().Lister(),
		)
	}

	// resolve kubernetes.default.svc locally
	if localHost, err := url.Parse(hostname); err == nil {
		serviceResolver = aggregatorapiserver.NewLoopbackServiceResolver(serviceResolver, localHost)
	}
	return serviceResolver
}
