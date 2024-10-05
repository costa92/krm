// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ContainerStateWaitingApplyConfiguration represents an declarative configuration of the ContainerStateWaiting type for use
// with apply.
type ContainerStateWaitingApplyConfiguration struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ContainerStateWaitingApplyConfiguration constructs an declarative configuration of the ContainerStateWaiting type for use with
// apply.
func ContainerStateWaiting() *ContainerStateWaitingApplyConfiguration {
	return &ContainerStateWaitingApplyConfiguration{}
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ContainerStateWaitingApplyConfiguration) WithReason(value string) *ContainerStateWaitingApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ContainerStateWaitingApplyConfiguration) WithMessage(value string) *ContainerStateWaitingApplyConfiguration {
	b.Message = &value
	return b
}
