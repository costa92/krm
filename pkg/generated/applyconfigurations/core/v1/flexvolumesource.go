// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FlexVolumeSourceApplyConfiguration represents an declarative configuration of the FlexVolumeSource type for use
// with apply.
type FlexVolumeSourceApplyConfiguration struct {
	Driver    *string                                 `json:"driver,omitempty"`
	FSType    *string                                 `json:"fsType,omitempty"`
	SecretRef *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	ReadOnly  *bool                                   `json:"readOnly,omitempty"`
	Options   map[string]string                       `json:"options,omitempty"`
}

// FlexVolumeSourceApplyConfiguration constructs an declarative configuration of the FlexVolumeSource type for use with
// apply.
func FlexVolumeSource() *FlexVolumeSourceApplyConfiguration {
	return &FlexVolumeSourceApplyConfiguration{}
}

// WithDriver sets the Driver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Driver field is set to the value of the last call.
func (b *FlexVolumeSourceApplyConfiguration) WithDriver(value string) *FlexVolumeSourceApplyConfiguration {
	b.Driver = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *FlexVolumeSourceApplyConfiguration) WithFSType(value string) *FlexVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *FlexVolumeSourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *FlexVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *FlexVolumeSourceApplyConfiguration) WithReadOnly(value bool) *FlexVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithOptions puts the entries into the Options field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Options field,
// overwriting an existing map entries in Options field with the same key.
func (b *FlexVolumeSourceApplyConfiguration) WithOptions(entries map[string]string) *FlexVolumeSourceApplyConfiguration {
	if b.Options == nil && len(entries) > 0 {
		b.Options = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Options[k] = v
	}
	return b
}
