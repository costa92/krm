package options

import (
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
)

type Options struct {
	Logs *logs.Options
}

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	Options
}

type CompletedOptions struct {
	*completedOptions
}

func NewOptions() *Options {
	o := &Options{
		Logs: logs.NewOptions(),
	}
	return o
}

func (o *Options) AddFlags(fss *cliflag.NamedFlagSets) {
	logsapi.AddFlags(o.Logs, fss.FlagSet("logs"))
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
