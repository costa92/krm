package options

import (
	"github.com/costa92/krm/internal/usercenter"
	"github.com/costa92/krm/pkg/app"
	"github.com/costa92/krm/pkg/log"
	genericoptions "github.com/costa92/krm/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"
)

const (
	// UserAgent is the userAgent name when starting onex-gateway server.
	UserAgent = "krm-usercenter"
)

var _ app.CliOptions = (*Options)(nil)

// Options contains state for master/api server.
type Options struct {
	// GenericOptions *genericoptions.Options       `json:"server"   mapstructure:"server"`
	// gRPC options for configuring gRPC related options.
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// HTTP options for configuring HTTP related options.
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
	// TLS options for configuring TLS related options.
	TLSOptions *genericoptions.TLSOptions `json:"tls" mapstructure:"tls"`
	// Metrics options for configuring metric related options.
	Metrics *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`
	// Log options for configuring log related options.
	Log *log.Options `json:"log" mapstructure:"log"`
}

// NewOptions returns initialized Options.
func NewOptions() *Options {
	o := &Options{
		GRPCOptions: genericoptions.NewGRPCOptions(),
		HTTPOptions: genericoptions.NewHTTPOptions(),
		TLSOptions:  genericoptions.NewTLSOptions(),
		Metrics:     genericoptions.NewMetricsOptions(),
		Log:         log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	o.Log.AddFlags(fss.FlagSet("log"))

	return fss
}

// ApplyTo fills up krm-usercenter config with options.
func (o *Options) ApplyTo(c *usercenter.Config) error {
	c.GRPCOptions = o.GRPCOptions
	c.HTTPOptions = o.HTTPOptions
	c.TLSOptions = o.TLSOptions

	return nil
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	var errs []error
	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.Metrics.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return utilerrors.NewAggregate(errs)
}

// Config return an krm-usercenter config object.
func (o *Options) Config() (*usercenter.Config, error) {
	c := &usercenter.Config{}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
