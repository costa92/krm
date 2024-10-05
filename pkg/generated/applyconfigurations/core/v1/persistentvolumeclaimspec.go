// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "github.com/costa92/krm/pkg/generated/applyconfigurations/meta/v1"
	v1 "k8s.io/api/core/v1"
)

// PersistentVolumeClaimSpecApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimSpec type for use
// with apply.
type PersistentVolumeClaimSpecApplyConfiguration struct {
	AccessModes               []v1.PersistentVolumeAccessMode               `json:"accessModes,omitempty"`
	Selector                  *metav1.LabelSelectorApplyConfiguration       `json:"selector,omitempty"`
	Resources                 *VolumeResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
	VolumeName                *string                                       `json:"volumeName,omitempty"`
	StorageClassName          *string                                       `json:"storageClassName,omitempty"`
	VolumeMode                *v1.PersistentVolumeMode                      `json:"volumeMode,omitempty"`
	DataSource                *TypedLocalObjectReferenceApplyConfiguration  `json:"dataSource,omitempty"`
	DataSourceRef             *TypedObjectReferenceApplyConfiguration       `json:"dataSourceRef,omitempty"`
	VolumeAttributesClassName *string                                       `json:"volumeAttributesClassName,omitempty"`
}

// PersistentVolumeClaimSpecApplyConfiguration constructs an declarative configuration of the PersistentVolumeClaimSpec type for use with
// apply.
func PersistentVolumeClaimSpec() *PersistentVolumeClaimSpecApplyConfiguration {
	return &PersistentVolumeClaimSpecApplyConfiguration{}
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithAccessModes(values ...v1.PersistentVolumeAccessMode) *PersistentVolumeClaimSpecApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithSelector(value *metav1.LabelSelectorApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithResources(value *VolumeResourceRequirementsApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.Resources = value
	return b
}

// WithVolumeName sets the VolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeName field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithVolumeName(value string) *PersistentVolumeClaimSpecApplyConfiguration {
	b.VolumeName = &value
	return b
}

// WithStorageClassName sets the StorageClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClassName field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithStorageClassName(value string) *PersistentVolumeClaimSpecApplyConfiguration {
	b.StorageClassName = &value
	return b
}

// WithVolumeMode sets the VolumeMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeMode field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithVolumeMode(value v1.PersistentVolumeMode) *PersistentVolumeClaimSpecApplyConfiguration {
	b.VolumeMode = &value
	return b
}

// WithDataSource sets the DataSource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataSource field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithDataSource(value *TypedLocalObjectReferenceApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.DataSource = value
	return b
}

// WithDataSourceRef sets the DataSourceRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataSourceRef field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithDataSourceRef(value *TypedObjectReferenceApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.DataSourceRef = value
	return b
}

// WithVolumeAttributesClassName sets the VolumeAttributesClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeAttributesClassName field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithVolumeAttributesClassName(value string) *PersistentVolumeClaimSpecApplyConfiguration {
	b.VolumeAttributesClassName = &value
	return b
}
