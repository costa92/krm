package usercenter

import (
	"github.com/costa92/krm/internal/pkg/bootstrap"
	"github.com/costa92/krm/internal/usercenter/server"
	"github.com/costa92/krm/pkg/log"
	genericoptions "github.com/costa92/krm/pkg/options"
	"github.com/costa92/krm/pkg/version"
	"github.com/go-kratos/kratos/v2"
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
	GRPCOptions  *genericoptions.GRPCOptions
	HTTPOptions  *genericoptions.HTTPOptions
	TLSOptions   *genericoptions.TLSOptions
	JWTOptions   *genericoptions.JWTOptions
	RedisOptions *genericoptions.RedisOptions
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
	app, cleanup, err := wireApp(appInfo, conf, c.JWTOptions, c.RedisOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}

// Server represents the server.
type Server struct {
	app *kratos.App
}

// Run is a method of the Server struct that starts the server.
func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-stopCh

	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
		return err
	}

	return nil
}
