package apps

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder()

	// AddToScheme applies all stored functions t oa scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name use in this package.
const GroupName = "apps.onex.io"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

// Kind takes an unqualified kind and returns a Group qualified GroupKind.
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Chain{},
		&ChainList{},
		&Miner{},
		&MinerList{},
		&MinerSet{},
		&MinerSetList{},
		&autoscaling.Scale{})
	return nil
}
