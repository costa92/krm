package options

import cliflag "k8s.io/component-base/cli/flag"

type Options struct {
}

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	Options
}

type CompletedOptions struct {
	*completedOptions
}

func NewOptions() *Options {
	o := &Options{}
	return o
}

func (o *Options) AddFlags(fss *cliflag.NamedFlagSets) {

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
