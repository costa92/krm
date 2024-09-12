package options

import (
	controlplaneoptions "github.com/costa92/krm/internal/controlplane/apiserver/options"
)

// completedOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
// zh: completedOptions 是一个私有包装器，它强制在调用 Run 之前调用 Complete()。
type completedOptions struct {
	// zh: controlplane.CompletedOptions 是一个结构体，包含了完成选项。
	controlplaneoptions.CompletedOptions

	// zh: Extra 是一个结构体。
	Extra
}

// CompletedOptions contains the completed options for running the server.
// zh: CompletedOptions 包含运行服务器的完成选项。
type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	// zh: 嵌入一个私有指针，不能在此包之外实例化。
	*completedOptions
}

// Complete set default ServerRunOptions.
// Should be called after krm-apiserver flags parsed.
func (o *ServerRunOptions) Complete() (CompletedOptions, error) {
	if o == nil {
		return CompletedOptions{completedOptions: &completedOptions{}}, nil
	}

	controllable, err := o.Options.Complete()
	if err != nil {
		return CompletedOptions{}, err
	}

	// zh: completed 是一个结构体，包含了完成选项和额外选项。
	completed := completedOptions{
		CompletedOptions: controllable,
		Extra:            o.Extra,
	}

	return CompletedOptions{&completed}, nil
}
