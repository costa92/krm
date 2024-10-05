// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SleepActionApplyConfiguration represents an declarative configuration of the SleepAction type for use
// with apply.
type SleepActionApplyConfiguration struct {
	Seconds *int64 `json:"seconds,omitempty"`
}

// SleepActionApplyConfiguration constructs an declarative configuration of the SleepAction type for use with
// apply.
func SleepAction() *SleepActionApplyConfiguration {
	return &SleepActionApplyConfiguration{}
}

// WithSeconds sets the Seconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Seconds field is set to the value of the last call.
func (b *SleepActionApplyConfiguration) WithSeconds(value int64) *SleepActionApplyConfiguration {
	b.Seconds = &value
	return b
}
