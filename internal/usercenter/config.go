package usercenter

import (
	"github.com/costa92/krm/internal/pkg/bootstrap"
	"github.com/costa92/krm/internal/usercenter/server"
	genericoptions "github.com/costa92/krm/pkg/options"
	"github.com/costa92/krm/pkg/version"
	"os"
)

var (
	// Name is the name of the compiled software.
	Name = "krm-usercenter"

	// ID contains the host name and any error encountered during the retrieval.
	ID, _ = os.Hostname()
)

// Config represents the configuration of the service.
type Config struct {
	GRPCOptions *genericoptions.GRPCOptions
	HTTPOptions *genericoptions.HTTPOptions
	TLSOptions  *genericoptions.TLSOptions
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

// completedConfig holds the configuration after it has been completed.
type completedConfig struct {
	*Config
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New(stopCh <-chan struct{}) (*Server, error) {
	appInfo := bootstrap.NewAppInfo(ID, Name, version.Get().String())

	conf := &server.Config{
		HTTP: *c.HTTPOptions,
		GRPC: *c.GRPCOptions,
		TLS:  *c.TLSOptions,
	}

	// Initialize Kratos application with the provided configurations.
	app, cleanup, err := wireApp(
		appInfo,
		conf)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}
