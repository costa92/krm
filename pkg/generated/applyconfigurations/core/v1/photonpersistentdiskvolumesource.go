// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PhotonPersistentDiskVolumeSourceApplyConfiguration represents an declarative configuration of the PhotonPersistentDiskVolumeSource type for use
// with apply.
type PhotonPersistentDiskVolumeSourceApplyConfiguration struct {
	PdID   *string `json:"pdID,omitempty"`
	FSType *string `json:"fsType,omitempty"`
}

// PhotonPersistentDiskVolumeSourceApplyConfiguration constructs an declarative configuration of the PhotonPersistentDiskVolumeSource type for use with
// apply.
func PhotonPersistentDiskVolumeSource() *PhotonPersistentDiskVolumeSourceApplyConfiguration {
	return &PhotonPersistentDiskVolumeSourceApplyConfiguration{}
}

// WithPdID sets the PdID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PdID field is set to the value of the last call.
func (b *PhotonPersistentDiskVolumeSourceApplyConfiguration) WithPdID(value string) *PhotonPersistentDiskVolumeSourceApplyConfiguration {
	b.PdID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *PhotonPersistentDiskVolumeSourceApplyConfiguration) WithFSType(value string) *PhotonPersistentDiskVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}
