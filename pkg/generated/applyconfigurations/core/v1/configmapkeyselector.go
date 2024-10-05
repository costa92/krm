// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConfigMapKeySelectorApplyConfiguration represents an declarative configuration of the ConfigMapKeySelector type for use
// with apply.
type ConfigMapKeySelectorApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Key                                    *string `json:"key,omitempty"`
	Optional                               *bool   `json:"optional,omitempty"`
}

// ConfigMapKeySelectorApplyConfiguration constructs an declarative configuration of the ConfigMapKeySelector type for use with
// apply.
func ConfigMapKeySelector() *ConfigMapKeySelectorApplyConfiguration {
	return &ConfigMapKeySelectorApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConfigMapKeySelectorApplyConfiguration) WithName(value string) *ConfigMapKeySelectorApplyConfiguration {
	b.Name = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *ConfigMapKeySelectorApplyConfiguration) WithKey(value string) *ConfigMapKeySelectorApplyConfiguration {
	b.Key = &value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *ConfigMapKeySelectorApplyConfiguration) WithOptional(value bool) *ConfigMapKeySelectorApplyConfiguration {
	b.Optional = &value
	return b
}
