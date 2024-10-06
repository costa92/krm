package options

import (
	"time"

	"github.com/costa92/krm/internal/pkg/options"
	"github.com/costa92/krm/pkg/generated/informers"
	corev1 "k8s.io/api/core/v1"
	peerreconcilers "k8s.io/apiserver/pkg/reconcilers"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/metrics"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

const defaultEtcdPathPrefix = "/registry/krm.io"

// Options contains state for master/api server.
type Options struct {
	// RecommendedOptions *genericoptions.RecommendedOptions
	GenericServerRunOptions *genericoptions.ServerRunOptions
	RecommendedOptions      *options.RecommendedOptions
	Features                *genericoptions.FeatureOptions
	Metrics                 *metrics.Options
	Logs                    *logs.Options
	Traces                  *genericoptions.TracingOptions
	// CloudOptions            *cloud.CloudOptions
	APIEnablement *genericoptions.APIEnablementOptions
	AlternateDNS  []string

	ProxyClientCertFile string
	ProxyClientKeyFile  string

	// PeerCAFile is the ca bundle used by this kube-apiserver to verify peer apiservers'
	// serving certs when routing a request to the peer in the case the request can not be served
	// locally due to version skew.
	PeerCAFile string

	// PeerAdvertiseAddress is the IP for this kube-apiserver which is used by peer apiservers to route a request
	// to this apiserver. This happens in cases where the peer is not able to serve the request due to
	// version skew.
	PeerAdvertiseAddress peerreconcilers.PeerAdvertiseAddress

	EnableAggregatorRouting             bool
	AggregatorRejectForwardingRedirects bool

	EnableLogsHandler          bool
	EventTTL                   time.Duration
	InternalVersionedInformers informers.SharedInformerFactory
}

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	Options
}

type CompletedOptions struct {
	*completedOptions
}

// NewOptions returns a new Options.
func NewOptions() *Options {
	o := &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		RecommendedOptions: options.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			legacyscheme.Codecs.LegacyCodec(corev1.SchemeGroupVersion), // NOTICE: [Custom API] Set default with corev1.SchemeGroupVersion
		),
		Features:      genericoptions.NewFeatureOptions(),
		Metrics:       metrics.NewOptions(),
		Logs:          logs.NewOptions(),
		Traces:        genericoptions.NewTracingOptions(),
		APIEnablement: genericoptions.NewAPIEnablementOptions(),

		AlternateDNS:            []string{"krm.io"},
		EnableLogsHandler:       true,
		EventTTL:                2 * time.Hour,
		EnableAggregatorRouting: false,
		// CloudOptions: cloud.NewCloudOptions(),
	}
	return o
}

func (o *Options) AddFlags(fss *cliflag.NamedFlagSets) {
	// Add the generic flags.
	o.GenericServerRunOptions.AddUniversalFlags(fss.FlagSet("generic"))
	o.RecommendedOptions.AddFlags(fss.FlagSet("recommended"))
	o.Features.AddFlags(fss.FlagSet("features"))
	o.APIEnablement.AddFlags(fss.FlagSet("API enablement"))
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	logsapi.AddFlags(o.Logs, fss.FlagSet("logs"))
	o.Traces.AddFlags(fss.FlagSet("traces"))
	// o.CloudOptions.AddFlags(fss.FlagSet("cloud"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.StringSliceVar(&o.AlternateDNS, "alternate-dns", o.AlternateDNS, "Specify an alternate DNS to use (e.g. 'onex.io').")
	fs.DurationVar(&o.EventTTL, "event-ttl", o.EventTTL, "Amount of time to retain events.")

	fs.BoolVar(&o.EnableLogsHandler, "enable-logs-handler", o.EnableLogsHandler,
		"If true, install a /logs handler for the apiserver logs.")
	_ = fs.MarkDeprecated("enable-logs-handler", "This flag will be removed in v1.19")

	fs.StringVar(&o.ProxyClientCertFile, "proxy-client-cert-file", o.ProxyClientCertFile, ""+
		"Client certificate used to prove the identity of the aggregator or kube-apiserver "+
		"when it must call out during a request. This includes proxying requests to a user "+
		"api-server and calling out to webhook admission plugins. It is expected that this "+
		"cert includes a signature from the CA in the --requestheader-client-ca-file flag. "+
		"That CA is published in the 'extension-apiserver-authentication' configmap in "+
		"the kube-system namespace. Components receiving calls from kube-aggregator should "+
		"use that CA to perform their half of the mutual TLS verification.")
	fs.StringVar(&o.ProxyClientKeyFile, "proxy-client-key-file", o.ProxyClientKeyFile, ""+
		"Private key for the client certificate used to prove the identity of the aggregator or kube-apiserver "+
		"when it must call out during a request. This includes proxying requests to a user "+
		"api-server and calling out to webhook admission plugins.")

	fs.StringVar(&o.PeerCAFile, "peer-ca-file", o.PeerCAFile,
		"If set and the UnknownVersionInteroperabilityProxy feature gate is enabled, this file will be used to verify serving certificates of peer kube-apiservers. "+
			"This flag is only used in clusters configured with multiple kube-apiservers for high availability.")

	fs.StringVar(&o.PeerAdvertiseAddress.PeerAdvertiseIP, "peer-advertise-ip", o.PeerAdvertiseAddress.PeerAdvertiseIP,
		"If set and the UnknownVersionInteroperabilityProxy feature gate is enabled, this IP will be used by peer kube-apiservers to proxy requests to this kube-apiserver "+
			"when the request cannot be handled by the peer due to version skew between the kube-apiservers. "+
			"This flag is only used in clusters configured with multiple kube-apiservers for high availability. ")

	fs.StringVar(&o.PeerAdvertiseAddress.PeerAdvertisePort, "peer-advertise-port", o.PeerAdvertiseAddress.PeerAdvertisePort,
		"If set and the UnknownVersionInteroperabilityProxy feature gate is enabled, this port will be used by peer kube-apiservers to proxy requests to this kube-apiserver "+
			"when the request cannot be handled by the peer due to version skew between the kube-apiservers. "+
			"This flag is only used in clusters configured with multiple kube-apiservers for high availability. ")

	fs.BoolVar(&o.EnableAggregatorRouting, "enable-aggregator-routing", o.EnableAggregatorRouting,
		"Turns on aggregator routing requests to endpoints IP rather than cluster IP.")

	fs.BoolVar(&o.AggregatorRejectForwardingRedirects, "aggregator-reject-forwarding-redirect", o.AggregatorRejectForwardingRedirects,
		"Aggregator reject forwarding redirect response back to client.")
}

func (o *Options) Complete() (CompletedOptions, error) {
	if o == nil {
		return CompletedOptions{completedOptions: &completedOptions{}}, nil
	}

	completed := completedOptions{Options: *o}

	return CompletedOptions{
		completedOptions: &completed,
	}, nil
}
