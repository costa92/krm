package options

import (
	"github.com/spf13/pflag"

	// ensure libs have a chance to globally register their flags
	_ "k8s.io/apiserver/pkg/admission"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers"
)

// AddCustomGlobalFlags explicitly registers flags that internal packages register
func AddCustomGlobalFlags(fs *pflag.FlagSet) {

}
