package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func init() {
	// AddToSchemes is a list of functions added to Scheme
	Install(legacyscheme.Scheme)
}

func Install(scheme *runtime.Scheme) {
}
