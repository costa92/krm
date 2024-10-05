// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// StorageOSPersistentVolumeSourceApplyConfiguration represents an declarative configuration of the StorageOSPersistentVolumeSource type for use
// with apply.
type StorageOSPersistentVolumeSourceApplyConfiguration struct {
	VolumeName      *string                            `json:"volumeName,omitempty"`
	VolumeNamespace *string                            `json:"volumeNamespace,omitempty"`
	FSType          *string                            `json:"fsType,omitempty"`
	ReadOnly        *bool                              `json:"readOnly,omitempty"`
	SecretRef       *ObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// StorageOSPersistentVolumeSourceApplyConfiguration constructs an declarative configuration of the StorageOSPersistentVolumeSource type for use with
// apply.
func StorageOSPersistentVolumeSource() *StorageOSPersistentVolumeSourceApplyConfiguration {
	return &StorageOSPersistentVolumeSourceApplyConfiguration{}
}

// WithVolumeName sets the VolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeName field is set to the value of the last call.
func (b *StorageOSPersistentVolumeSourceApplyConfiguration) WithVolumeName(value string) *StorageOSPersistentVolumeSourceApplyConfiguration {
	b.VolumeName = &value
	return b
}

// WithVolumeNamespace sets the VolumeNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeNamespace field is set to the value of the last call.
func (b *StorageOSPersistentVolumeSourceApplyConfiguration) WithVolumeNamespace(value string) *StorageOSPersistentVolumeSourceApplyConfiguration {
	b.VolumeNamespace = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *StorageOSPersistentVolumeSourceApplyConfiguration) WithFSType(value string) *StorageOSPersistentVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *StorageOSPersistentVolumeSourceApplyConfiguration) WithReadOnly(value bool) *StorageOSPersistentVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *StorageOSPersistentVolumeSourceApplyConfiguration) WithSecretRef(value *ObjectReferenceApplyConfiguration) *StorageOSPersistentVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}
