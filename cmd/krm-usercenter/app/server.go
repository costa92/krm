package app

import (
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/costa92/krm/cmd/krm-usercenter/app/options"
	"github.com/costa92/krm/internal/usercenter"
	"github.com/costa92/krm/pkg/app"
)

// Define the description of the command.
const commandDesc = `The usercenter server is used to manage users, keys, fees, etc.`

// NewApp creates and returns a new App object with default parameters.
func NewApp(name string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(name, "Launch a onex usercenter server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

// Returns the function to run the application.
func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}

		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(c *usercenter.Config, stopCh <-chan struct{}) error {
	server, err := c.Complete().New(stopCh)
	if err != nil {
		return err
	}

	return server.Run(stopCh)
}
