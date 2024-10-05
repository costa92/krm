// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// LocalVolumeSourceApplyConfiguration represents an declarative configuration of the LocalVolumeSource type for use
// with apply.
type LocalVolumeSourceApplyConfiguration struct {
	Path   *string `json:"path,omitempty"`
	FSType *string `json:"fsType,omitempty"`
}

// LocalVolumeSourceApplyConfiguration constructs an declarative configuration of the LocalVolumeSource type for use with
// apply.
func LocalVolumeSource() *LocalVolumeSourceApplyConfiguration {
	return &LocalVolumeSourceApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *LocalVolumeSourceApplyConfiguration) WithPath(value string) *LocalVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *LocalVolumeSourceApplyConfiguration) WithFSType(value string) *LocalVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}
