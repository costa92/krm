package options

import (
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	known "github.com/costa92/krm/internal/pkg/known/usercenter"
	"github.com/costa92/krm/internal/usercenter"
	"github.com/costa92/krm/pkg/app"
	"github.com/costa92/krm/pkg/log"
	genericoptions "github.com/costa92/krm/pkg/options"
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

	// Mysql options for configuring Mysql related options.
	MysqlOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`

	// Redis options for configuring Redis related options.
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`

	// Metrics options for configuring metric related options.
	Metrics *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`

	// JWT options for configuring JWT related options.
	JWTOptions *genericoptions.JWTOptions `json:"jwt" mapstructure:"jwt"`

	// Kafka options for configuring Kafka related options.
	KafkaOptions *genericoptions.KafkaOptions `json:"kafka" mapstructure:"kafka"`

	// Log options for configuring log related options.
	Log *log.Options `json:"log" mapstructure:"log"`
}

// NewOptions returns initialized Options.
func NewOptions() *Options {
	o := &Options{
		GRPCOptions:  genericoptions.NewGRPCOptions(),
		HTTPOptions:  genericoptions.NewHTTPOptions(),
		TLSOptions:   genericoptions.NewTLSOptions(),
		MysqlOptions: genericoptions.NewMySQLOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
		JWTOptions:   genericoptions.NewJWTOptions(),
		KafkaOptions: genericoptions.NewKafkaOptions(),
		Metrics:      genericoptions.NewMetricsOptions(),
		Log:          log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	o.JWTOptions.AddFlags(fss.FlagSet("jwt"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.MysqlOptions.AddFlags(fss.FlagSet("mysql"))
	o.KafkaOptions.AddFlags(fss.FlagSet("kafka"))
	o.Log.AddFlags(fss.FlagSet("log"))

	return fss
}

// ApplyTo fills up krm-usercenter config with options.
func (o *Options) ApplyTo(c *usercenter.Config) error {
	c.GRPCOptions = o.GRPCOptions
	c.HTTPOptions = o.HTTPOptions
	c.TLSOptions = o.TLSOptions
	c.MySQLOptions = o.MysqlOptions
	c.RedisOptions = o.RedisOptions
	c.JWTOptions = o.JWTOptions
	c.KafkaOptions = o.KafkaOptions
	return nil
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	o.JWTOptions.Expired = known.RefreshTokenExpire
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	var errs []error
	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.MysqlOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.JWTOptions.Validate()...)
	errs = append(errs, o.KafkaOptions.Validate()...)
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
