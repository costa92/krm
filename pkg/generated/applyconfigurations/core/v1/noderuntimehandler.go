// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeRuntimeHandlerApplyConfiguration represents an declarative configuration of the NodeRuntimeHandler type for use
// with apply.
type NodeRuntimeHandlerApplyConfiguration struct {
	Name     *string                                       `json:"name,omitempty"`
	Features *NodeRuntimeHandlerFeaturesApplyConfiguration `json:"features,omitempty"`
}

// NodeRuntimeHandlerApplyConfiguration constructs an declarative configuration of the NodeRuntimeHandler type for use with
// apply.
func NodeRuntimeHandler() *NodeRuntimeHandlerApplyConfiguration {
	return &NodeRuntimeHandlerApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NodeRuntimeHandlerApplyConfiguration) WithName(value string) *NodeRuntimeHandlerApplyConfiguration {
	b.Name = &value
	return b
}

// WithFeatures sets the Features field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Features field is set to the value of the last call.
func (b *NodeRuntimeHandlerApplyConfiguration) WithFeatures(value *NodeRuntimeHandlerFeaturesApplyConfiguration) *NodeRuntimeHandlerApplyConfiguration {
	b.Features = value
	return b
}
