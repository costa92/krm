// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// PersistentVolumeClaimStatusApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimStatus type for use
// with apply.
type PersistentVolumeClaimStatusApplyConfiguration struct {
	Phase                            *v1.PersistentVolumeClaimPhase                     `json:"phase,omitempty"`
	AccessModes                      []v1.PersistentVolumeAccessMode                    `json:"accessModes,omitempty"`
	Capacity                         *v1.ResourceList                                   `json:"capacity,omitempty"`
	Conditions                       []PersistentVolumeClaimConditionApplyConfiguration `json:"conditions,omitempty"`
	AllocatedResources               *v1.ResourceList                                   `json:"allocatedResources,omitempty"`
	AllocatedResourceStatuses        map[v1.ResourceName]v1.ClaimResourceStatus         `json:"allocatedResourceStatuses,omitempty"`
	CurrentVolumeAttributesClassName *string                                            `json:"currentVolumeAttributesClassName,omitempty"`
	ModifyVolumeStatus               *ModifyVolumeStatusApplyConfiguration              `json:"modifyVolumeStatus,omitempty"`
}

// PersistentVolumeClaimStatusApplyConfiguration constructs an declarative configuration of the PersistentVolumeClaimStatus type for use with
// apply.
func PersistentVolumeClaimStatus() *PersistentVolumeClaimStatusApplyConfiguration {
	return &PersistentVolumeClaimStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithPhase(value v1.PersistentVolumeClaimPhase) *PersistentVolumeClaimStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithAccessModes(values ...v1.PersistentVolumeAccessMode) *PersistentVolumeClaimStatusApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithCapacity(value v1.ResourceList) *PersistentVolumeClaimStatusApplyConfiguration {
	b.Capacity = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithConditions(values ...*PersistentVolumeClaimConditionApplyConfiguration) *PersistentVolumeClaimStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithAllocatedResources sets the AllocatedResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocatedResources field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithAllocatedResources(value v1.ResourceList) *PersistentVolumeClaimStatusApplyConfiguration {
	b.AllocatedResources = &value
	return b
}

// WithAllocatedResourceStatuses puts the entries into the AllocatedResourceStatuses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the AllocatedResourceStatuses field,
// overwriting an existing map entries in AllocatedResourceStatuses field with the same key.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithAllocatedResourceStatuses(entries map[v1.ResourceName]v1.ClaimResourceStatus) *PersistentVolumeClaimStatusApplyConfiguration {
	if b.AllocatedResourceStatuses == nil && len(entries) > 0 {
		b.AllocatedResourceStatuses = make(map[v1.ResourceName]v1.ClaimResourceStatus, len(entries))
	}
	for k, v := range entries {
		b.AllocatedResourceStatuses[k] = v
	}
	return b
}

// WithCurrentVolumeAttributesClassName sets the CurrentVolumeAttributesClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentVolumeAttributesClassName field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithCurrentVolumeAttributesClassName(value string) *PersistentVolumeClaimStatusApplyConfiguration {
	b.CurrentVolumeAttributesClassName = &value
	return b
}

// WithModifyVolumeStatus sets the ModifyVolumeStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ModifyVolumeStatus field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithModifyVolumeStatus(value *ModifyVolumeStatusApplyConfiguration) *PersistentVolumeClaimStatusApplyConfiguration {
	b.ModifyVolumeStatus = value
	return b
}
