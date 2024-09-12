package controlplane

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/kubernetes/pkg/routes"
)

type ExtraConfig struct {
	EnableLogsSupport bool
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   ExtraConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package.
type CompletedConfig struct {
	*completedConfig
}

// Instance is a control plane instance.
type Instance struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

// Complete returns a CompletedConfig.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{&completedConfig{
		GenericConfig: c.GenericConfig.Complete(),
		ExtraConfig:   c.ExtraConfig,
	}}
}

func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*Instance, error) {
	s, err := c.GenericConfig.New("krm-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}

	if c.ExtraConfig.EnableLogsSupport {
		// Add log support.
		routes.Logs{}.Install(s.Handler.GoRestfulContainer)
	}

	m := &Instance{
		GenericAPIServer: s,
	}
	return m, nil
}
