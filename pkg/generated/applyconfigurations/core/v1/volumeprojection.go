// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VolumeProjectionApplyConfiguration represents an declarative configuration of the VolumeProjection type for use
// with apply.
type VolumeProjectionApplyConfiguration struct {
	Secret              *SecretProjectionApplyConfiguration              `json:"secret,omitempty"`
	DownwardAPI         *DownwardAPIProjectionApplyConfiguration         `json:"downwardAPI,omitempty"`
	ConfigMap           *ConfigMapProjectionApplyConfiguration           `json:"configMap,omitempty"`
	ServiceAccountToken *ServiceAccountTokenProjectionApplyConfiguration `json:"serviceAccountToken,omitempty"`
	ClusterTrustBundle  *ClusterTrustBundleProjectionApplyConfiguration  `json:"clusterTrustBundle,omitempty"`
}

// VolumeProjectionApplyConfiguration constructs an declarative configuration of the VolumeProjection type for use with
// apply.
func VolumeProjection() *VolumeProjectionApplyConfiguration {
	return &VolumeProjectionApplyConfiguration{}
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *VolumeProjectionApplyConfiguration) WithSecret(value *SecretProjectionApplyConfiguration) *VolumeProjectionApplyConfiguration {
	b.Secret = value
	return b
}

// WithDownwardAPI sets the DownwardAPI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DownwardAPI field is set to the value of the last call.
func (b *VolumeProjectionApplyConfiguration) WithDownwardAPI(value *DownwardAPIProjectionApplyConfiguration) *VolumeProjectionApplyConfiguration {
	b.DownwardAPI = value
	return b
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *VolumeProjectionApplyConfiguration) WithConfigMap(value *ConfigMapProjectionApplyConfiguration) *VolumeProjectionApplyConfiguration {
	b.ConfigMap = value
	return b
}

// WithServiceAccountToken sets the ServiceAccountToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountToken field is set to the value of the last call.
func (b *VolumeProjectionApplyConfiguration) WithServiceAccountToken(value *ServiceAccountTokenProjectionApplyConfiguration) *VolumeProjectionApplyConfiguration {
	b.ServiceAccountToken = value
	return b
}

// WithClusterTrustBundle sets the ClusterTrustBundle field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterTrustBundle field is set to the value of the last call.
func (b *VolumeProjectionApplyConfiguration) WithClusterTrustBundle(value *ClusterTrustBundleProjectionApplyConfiguration) *VolumeProjectionApplyConfiguration {
	b.ClusterTrustBundle = value
	return b
}
