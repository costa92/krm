// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	appsv1beta1 "github.com/costa92/krm/pkg/apis/apps/v1beta1"
	internal "github.com/costa92/krm/pkg/generated/applyconfigurations/internal"
	v1 "github.com/costa92/krm/pkg/generated/applyconfigurations/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// ChainApplyConfiguration represents an declarative configuration of the Chain type for use
// with apply.
type ChainApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ChainSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *ChainStatusApplyConfiguration `json:"status,omitempty"`
}

// Chain constructs an declarative configuration of the Chain type for use with
// apply.
func Chain(name, namespace string) *ChainApplyConfiguration {
	b := &ChainApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Chain")
	b.WithAPIVersion("apps.onex.io/v1beta1")
	return b
}

// ExtractChain extracts the applied configuration owned by fieldManager from
// chain. If no managedFields are found in chain for fieldManager, a
// ChainApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// chain must be a unmodified Chain API object that was retrieved from the Kubernetes API.
// ExtractChain provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractChain(chain *appsv1beta1.Chain, fieldManager string) (*ChainApplyConfiguration, error) {
	return extractChain(chain, fieldManager, "")
}

// ExtractChainStatus is the same as ExtractChain except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractChainStatus(chain *appsv1beta1.Chain, fieldManager string) (*ChainApplyConfiguration, error) {
	return extractChain(chain, fieldManager, "status")
}

func extractChain(chain *appsv1beta1.Chain, fieldManager string, subresource string) (*ChainApplyConfiguration, error) {
	b := &ChainApplyConfiguration{}
	err := managedfields.ExtractInto(chain, internal.Parser().Type("com.github.superproj.onex.pkg.apis.apps.v1beta1.Chain"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(chain.Name)
	b.WithNamespace(chain.Namespace)

	b.WithKind("Chain")
	b.WithAPIVersion("apps.onex.io/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithKind(value string) *ChainApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithAPIVersion(value string) *ChainApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithName(value string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithGenerateName(value string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithNamespace(value string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithUID(value types.UID) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithResourceVersion(value string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithGeneration(value int64) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ChainApplyConfiguration) WithLabels(entries map[string]string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ChainApplyConfiguration) WithAnnotations(entries map[string]string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ChainApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ChainApplyConfiguration) WithFinalizers(values ...string) *ChainApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ChainApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithSpec(value *ChainSpecApplyConfiguration) *ChainApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ChainApplyConfiguration) WithStatus(value *ChainStatusApplyConfiguration) *ChainApplyConfiguration {
	b.Status = value
	return b
}
