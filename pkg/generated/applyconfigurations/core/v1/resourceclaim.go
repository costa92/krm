// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ResourceClaimApplyConfiguration represents an declarative configuration of the ResourceClaim type for use
// with apply.
type ResourceClaimApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// ResourceClaimApplyConfiguration constructs an declarative configuration of the ResourceClaim type for use with
// apply.
func ResourceClaim() *ResourceClaimApplyConfiguration {
	return &ResourceClaimApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceClaimApplyConfiguration) WithName(value string) *ResourceClaimApplyConfiguration {
	b.Name = &value
	return b
}
