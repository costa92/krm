package app

import (
	"github.com/costa92/krm/cmd/krm-apiserver/app/options"
	"github.com/costa92/krm/internal/controlplane"
	apiextensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
)

type Config struct {
	Options options.CompletedOptions

	Aggregator    *aggregatorapiserver.Config
	ControlPlane  *controlplane.Config
	ApiExtensions *apiextensionsapiserver.Config
	// ExtraConfig is a placeholder for additional configuration.
	ExtraConfig
}

// ExtraConfig is a placeholder for additional configuration.
type ExtraConfig struct{}

type completedConfig struct {
	Options options.CompletedOptions

	Aggregator    aggregatorapiserver.CompletedConfig
	ControlPlane  controlplane.CompletedConfig
	ApiExtensions apiextensionsapiserver.CompletedConfig

	ExtraConfig
}

// CompletedConfig is a wrapper for the completed configuration.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete sets the default ServerRunOptions.
func (c *Config) Complete() (CompletedConfig, error) {
	return CompletedConfig{&completedConfig{
		Options: c.Options,

		Aggregator:    c.Aggregator.Complete(),
		ControlPlane:  c.ControlPlane.Complete(),
		ApiExtensions: c.ApiExtensions.Complete(),

		ExtraConfig: c.ExtraConfig,
	}}, nil
}

// NewConfig creates all the resources for running krm-apiserver, but runs none of them.
func NewConfig(opts options.CompletedOptions) (*Config, error) {
	c := &Config{
		Options: opts,
	}
	return c, nil
}
