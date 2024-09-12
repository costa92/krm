package options

import (
	"github.com/spf13/pflag"
	// ensure libs have a chance to globally register their flags
	_ "k8s.io/apiserver/pkg/admission"
)

// AddCustomGlobalFlags explicitly registers flags that internal packages register
func AddCustomGlobalFlags(fs *pflag.FlagSet) {
	//globalflag.Register(fs, "default-not-ready-toleration-seconds")
}
