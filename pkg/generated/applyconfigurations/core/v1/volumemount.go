// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// VolumeMountApplyConfiguration represents an declarative configuration of the VolumeMount type for use
// with apply.
type VolumeMountApplyConfiguration struct {
	Name              *string                   `json:"name,omitempty"`
	ReadOnly          *bool                     `json:"readOnly,omitempty"`
	RecursiveReadOnly *v1.RecursiveReadOnlyMode `json:"recursiveReadOnly,omitempty"`
	MountPath         *string                   `json:"mountPath,omitempty"`
	SubPath           *string                   `json:"subPath,omitempty"`
	MountPropagation  *v1.MountPropagationMode  `json:"mountPropagation,omitempty"`
	SubPathExpr       *string                   `json:"subPathExpr,omitempty"`
}

// VolumeMountApplyConfiguration constructs an declarative configuration of the VolumeMount type for use with
// apply.
func VolumeMount() *VolumeMountApplyConfiguration {
	return &VolumeMountApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithName(value string) *VolumeMountApplyConfiguration {
	b.Name = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithReadOnly(value bool) *VolumeMountApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithRecursiveReadOnly sets the RecursiveReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RecursiveReadOnly field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithRecursiveReadOnly(value v1.RecursiveReadOnlyMode) *VolumeMountApplyConfiguration {
	b.RecursiveReadOnly = &value
	return b
}

// WithMountPath sets the MountPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountPath field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithMountPath(value string) *VolumeMountApplyConfiguration {
	b.MountPath = &value
	return b
}

// WithSubPath sets the SubPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubPath field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithSubPath(value string) *VolumeMountApplyConfiguration {
	b.SubPath = &value
	return b
}

// WithMountPropagation sets the MountPropagation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountPropagation field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithMountPropagation(value v1.MountPropagationMode) *VolumeMountApplyConfiguration {
	b.MountPropagation = &value
	return b
}

// WithSubPathExpr sets the SubPathExpr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubPathExpr field is set to the value of the last call.
func (b *VolumeMountApplyConfiguration) WithSubPathExpr(value string) *VolumeMountApplyConfiguration {
	b.SubPathExpr = &value
	return b
}
