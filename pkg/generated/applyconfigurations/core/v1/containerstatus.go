// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// ContainerStatusApplyConfiguration represents an declarative configuration of the ContainerStatus type for use
// with apply.
type ContainerStatusApplyConfiguration struct {
	Name                 *string                                 `json:"name,omitempty"`
	State                *ContainerStateApplyConfiguration       `json:"state,omitempty"`
	LastTerminationState *ContainerStateApplyConfiguration       `json:"lastState,omitempty"`
	Ready                *bool                                   `json:"ready,omitempty"`
	RestartCount         *int32                                  `json:"restartCount,omitempty"`
	Image                *string                                 `json:"image,omitempty"`
	ImageID              *string                                 `json:"imageID,omitempty"`
	ContainerID          *string                                 `json:"containerID,omitempty"`
	Started              *bool                                   `json:"started,omitempty"`
	AllocatedResources   *corev1.ResourceList                    `json:"allocatedResources,omitempty"`
	Resources            *ResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
	VolumeMounts         []VolumeMountStatusApplyConfiguration   `json:"volumeMounts,omitempty"`
}

// ContainerStatusApplyConfiguration constructs an declarative configuration of the ContainerStatus type for use with
// apply.
func ContainerStatus() *ContainerStatusApplyConfiguration {
	return &ContainerStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithName(value string) *ContainerStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithState(value *ContainerStateApplyConfiguration) *ContainerStatusApplyConfiguration {
	b.State = value
	return b
}

// WithLastTerminationState sets the LastTerminationState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTerminationState field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithLastTerminationState(value *ContainerStateApplyConfiguration) *ContainerStatusApplyConfiguration {
	b.LastTerminationState = value
	return b
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithReady(value bool) *ContainerStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithRestartCount sets the RestartCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartCount field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithRestartCount(value int32) *ContainerStatusApplyConfiguration {
	b.RestartCount = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithImage(value string) *ContainerStatusApplyConfiguration {
	b.Image = &value
	return b
}

// WithImageID sets the ImageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageID field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithImageID(value string) *ContainerStatusApplyConfiguration {
	b.ImageID = &value
	return b
}

// WithContainerID sets the ContainerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerID field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithContainerID(value string) *ContainerStatusApplyConfiguration {
	b.ContainerID = &value
	return b
}

// WithStarted sets the Started field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Started field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithStarted(value bool) *ContainerStatusApplyConfiguration {
	b.Started = &value
	return b
}

// WithAllocatedResources sets the AllocatedResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocatedResources field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithAllocatedResources(value corev1.ResourceList) *ContainerStatusApplyConfiguration {
	b.AllocatedResources = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ContainerStatusApplyConfiguration) WithResources(value *ResourceRequirementsApplyConfiguration) *ContainerStatusApplyConfiguration {
	b.Resources = value
	return b
}

// WithVolumeMounts adds the given value to the VolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeMounts field.
func (b *ContainerStatusApplyConfiguration) WithVolumeMounts(values ...*VolumeMountStatusApplyConfiguration) *ContainerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumeMounts")
		}
		b.VolumeMounts = append(b.VolumeMounts, *values[i])
	}
	return b
}
