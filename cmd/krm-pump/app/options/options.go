package options

import (
	"github.com/costa92/krm/internal/pump"
	"github.com/costa92/krm/pkg/app"
	genericoptions "github.com/costa92/krm/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/util/feature"
	cliflag "k8s.io/component-base/cli/flag"
	"strings"
)

var _ app.CliOptions = (*Options)(nil)

// Options contains state for master/api server.
type Options struct {
	HealthOptions *genericoptions.HealthOptions `json:"health" mapstructure:"health"`
	KafkaOptions  *genericoptions.KafkaOptions  `json:"kafka" mapstructure:"kafka"`
	MongoOptions  *genericoptions.MongoOptions  `json:"mongo" mapstructure:"mongo"`
	FeatureGates  map[string]bool               `json:"feature-gates"`
}

// NewOptions returns initialized Options.
func NewOptions() *Options {
	o := &Options{
		// RedisOptions: genericoptions.NewRedisOptions(),
		HealthOptions: genericoptions.NewHealthOptions(),
		KafkaOptions:  genericoptions.NewKafkaOptions(),
		MongoOptions:  genericoptions.NewMongoOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.HealthOptions.AddFlags(fss.FlagSet("health"))
	o.KafkaOptions.AddFlags(fss.FlagSet("kafka"))
	o.MongoOptions.AddFlags(fss.FlagSet("mongo"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	feature.DefaultMutableFeatureGate.AddFlag(fs)

	return fss
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	if !strings.HasPrefix(o.MongoOptions.URL, "mongodb://") && !strings.HasPrefix(o.MongoOptions.URL, "mongodb+srv://") {
		// Preserve backwards compatibility for hostnames without a
		// scheme, broken in go 1.8. Remove in Telegraf 2.0
		o.MongoOptions.URL = "mongodb://" + o.MongoOptions.URL
	}

	_ = feature.DefaultMutableFeatureGate.SetFromMap(o.FeatureGates)
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	errs := []error{}

	errs = append(errs, o.HealthOptions.Validate()...)
	errs = append(errs, o.KafkaOptions.Validate()...)
	errs = append(errs, o.MongoOptions.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// ApplyTo fills up onex-pump config with options.
func (o *Options) ApplyTo(c *pump.Config) error {
	c.KafkaOptions = o.KafkaOptions
	c.MongoOptions = o.MongoOptions
	return nil
}

// Config return an onex-pump config object.
func (o *Options) Config() (*pump.Config, error) {
	c := &pump.Config{}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
