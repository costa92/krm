// Copyright 2024 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeleteOptionsApplyConfiguration represents an declarative configuration of the DeleteOptions type for use
// with apply.
type DeleteOptionsApplyConfiguration struct {
	TypeMetaApplyConfiguration `json:",inline"`
	GracePeriodSeconds         *int64                      `json:"gracePeriodSeconds,omitempty"`
	Preconditions              *metav1.Preconditions       `json:"preconditions,omitempty"`
	OrphanDependents           *bool                       `json:"orphanDependents,omitempty"`
	PropagationPolicy          *metav1.DeletionPropagation `json:"propagationPolicy,omitempty"`
	DryRun                     []string                    `json:"dryRun,omitempty"`
}

// DeleteOptionsApplyConfiguration constructs an declarative configuration of the DeleteOptions type for use with
// apply.
func DeleteOptions() *DeleteOptionsApplyConfiguration {
	b := &DeleteOptionsApplyConfiguration{}
	b.WithKind("DeleteOptions")
	b.WithAPIVersion("meta.k8s.io/v1")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithKind(value string) *DeleteOptionsApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithAPIVersion(value string) *DeleteOptionsApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithGracePeriodSeconds sets the GracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GracePeriodSeconds field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithGracePeriodSeconds(value int64) *DeleteOptionsApplyConfiguration {
	b.GracePeriodSeconds = &value
	return b
}

// WithPreconditions sets the Preconditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Preconditions field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithPreconditions(value metav1.Preconditions) *DeleteOptionsApplyConfiguration {
	b.Preconditions = &value
	return b
}

// WithOrphanDependents sets the OrphanDependents field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OrphanDependents field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithOrphanDependents(value bool) *DeleteOptionsApplyConfiguration {
	b.OrphanDependents = &value
	return b
}

// WithPropagationPolicy sets the PropagationPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PropagationPolicy field is set to the value of the last call.
func (b *DeleteOptionsApplyConfiguration) WithPropagationPolicy(value metav1.DeletionPropagation) *DeleteOptionsApplyConfiguration {
	b.PropagationPolicy = &value
	return b
}

// WithDryRun adds the given value to the DryRun field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DryRun field.
func (b *DeleteOptionsApplyConfiguration) WithDryRun(values ...string) *DeleteOptionsApplyConfiguration {
	for i := range values {
		b.DryRun = append(b.DryRun, values[i])
	}
	return b
}