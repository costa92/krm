// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// PodReadinessGateApplyConfiguration represents an declarative configuration of the PodReadinessGate type for use
// with apply.
type PodReadinessGateApplyConfiguration struct {
	ConditionType *v1.PodConditionType `json:"conditionType,omitempty"`
}

// PodReadinessGateApplyConfiguration constructs an declarative configuration of the PodReadinessGate type for use with
// apply.
func PodReadinessGate() *PodReadinessGateApplyConfiguration {
	return &PodReadinessGateApplyConfiguration{}
}

// WithConditionType sets the ConditionType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConditionType field is set to the value of the last call.
func (b *PodReadinessGateApplyConfiguration) WithConditionType(value v1.PodConditionType) *PodReadinessGateApplyConfiguration {
	b.ConditionType = &value
	return b
}
