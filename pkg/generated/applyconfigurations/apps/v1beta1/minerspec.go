// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MinerSpecApplyConfiguration represents an declarative configuration of the MinerSpec type for use
// with apply.
type MinerSpecApplyConfiguration struct {
	*ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	DisplayName                   *string           `json:"displayName,omitempty"`
	MinerType                     *string           `json:"minerType,omitempty"`
	ChainName                     *string           `json:"chainName,omitempty"`
	RestartPolicy                 *v1.RestartPolicy `json:"restartPolicy,omitempty"`
	PodDeletionTimeout            *metav1.Duration  `json:"podDeletionTimeout,omitempty"`
}

// MinerSpecApplyConfiguration constructs an declarative configuration of the MinerSpec type for use with
// apply.
func MinerSpec() *MinerSpecApplyConfiguration {
	return &MinerSpecApplyConfiguration{}
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *MinerSpecApplyConfiguration) WithLabels(entries map[string]string) *MinerSpecApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *MinerSpecApplyConfiguration) WithAnnotations(entries map[string]string) *MinerSpecApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

func (b *MinerSpecApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &ObjectMetaApplyConfiguration{}
	}
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *MinerSpecApplyConfiguration) WithDisplayName(value string) *MinerSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithMinerType sets the MinerType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinerType field is set to the value of the last call.
func (b *MinerSpecApplyConfiguration) WithMinerType(value string) *MinerSpecApplyConfiguration {
	b.MinerType = &value
	return b
}

// WithChainName sets the ChainName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChainName field is set to the value of the last call.
func (b *MinerSpecApplyConfiguration) WithChainName(value string) *MinerSpecApplyConfiguration {
	b.ChainName = &value
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *MinerSpecApplyConfiguration) WithRestartPolicy(value v1.RestartPolicy) *MinerSpecApplyConfiguration {
	b.RestartPolicy = &value
	return b
}

// WithPodDeletionTimeout sets the PodDeletionTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodDeletionTimeout field is set to the value of the last call.
func (b *MinerSpecApplyConfiguration) WithPodDeletionTimeout(value metav1.Duration) *MinerSpecApplyConfiguration {
	b.PodDeletionTimeout = &value
	return b
}
