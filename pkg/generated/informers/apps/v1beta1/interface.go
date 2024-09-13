package v1beta1

import (
	internalinterfaces "github.com/costa92/krm/pkg/generated/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Chains returns a ChainInformer.
	Chains() ChainInformer
	// ChargeRequests returns a ChargeRequestInformer.
	ChargeRequests() ChargeRequestInformer
	// Miners returns a MinerInformer.
	Miners() MinerInformer
	// MinerSets returns a MinerSetInformer.
	MinerSets() MinerSetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Chains returns a ChainInformer.
func (v *version) Chains() ChainInformer {
	return &chainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChargeRequests returns a ChargeRequestInformer.
func (v *version) ChargeRequests() ChargeRequestInformer {
	return &chargeRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Miners returns a MinerInformer.
func (v *version) Miners() MinerInformer {
	return &minerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MinerSets returns a MinerSetInformer.
func (v *version) MinerSets() MinerSetInformer {
	return &minerSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
