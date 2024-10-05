// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ISCSIVolumeSourceApplyConfiguration represents an declarative configuration of the ISCSIVolumeSource type for use
// with apply.
type ISCSIVolumeSourceApplyConfiguration struct {
	TargetPortal      *string                                 `json:"targetPortal,omitempty"`
	IQN               *string                                 `json:"iqn,omitempty"`
	Lun               *int32                                  `json:"lun,omitempty"`
	ISCSIInterface    *string                                 `json:"iscsiInterface,omitempty"`
	FSType            *string                                 `json:"fsType,omitempty"`
	ReadOnly          *bool                                   `json:"readOnly,omitempty"`
	Portals           []string                                `json:"portals,omitempty"`
	DiscoveryCHAPAuth *bool                                   `json:"chapAuthDiscovery,omitempty"`
	SessionCHAPAuth   *bool                                   `json:"chapAuthSession,omitempty"`
	SecretRef         *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	InitiatorName     *string                                 `json:"initiatorName,omitempty"`
}

// ISCSIVolumeSourceApplyConfiguration constructs an declarative configuration of the ISCSIVolumeSource type for use with
// apply.
func ISCSIVolumeSource() *ISCSIVolumeSourceApplyConfiguration {
	return &ISCSIVolumeSourceApplyConfiguration{}
}

// WithTargetPortal sets the TargetPortal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetPortal field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithTargetPortal(value string) *ISCSIVolumeSourceApplyConfiguration {
	b.TargetPortal = &value
	return b
}

// WithIQN sets the IQN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IQN field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithIQN(value string) *ISCSIVolumeSourceApplyConfiguration {
	b.IQN = &value
	return b
}

// WithLun sets the Lun field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Lun field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithLun(value int32) *ISCSIVolumeSourceApplyConfiguration {
	b.Lun = &value
	return b
}

// WithISCSIInterface sets the ISCSIInterface field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ISCSIInterface field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithISCSIInterface(value string) *ISCSIVolumeSourceApplyConfiguration {
	b.ISCSIInterface = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithFSType(value string) *ISCSIVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithReadOnly(value bool) *ISCSIVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithPortals adds the given value to the Portals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Portals field.
func (b *ISCSIVolumeSourceApplyConfiguration) WithPortals(values ...string) *ISCSIVolumeSourceApplyConfiguration {
	for i := range values {
		b.Portals = append(b.Portals, values[i])
	}
	return b
}

// WithDiscoveryCHAPAuth sets the DiscoveryCHAPAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiscoveryCHAPAuth field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithDiscoveryCHAPAuth(value bool) *ISCSIVolumeSourceApplyConfiguration {
	b.DiscoveryCHAPAuth = &value
	return b
}

// WithSessionCHAPAuth sets the SessionCHAPAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SessionCHAPAuth field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithSessionCHAPAuth(value bool) *ISCSIVolumeSourceApplyConfiguration {
	b.SessionCHAPAuth = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *ISCSIVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithInitiatorName sets the InitiatorName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InitiatorName field is set to the value of the last call.
func (b *ISCSIVolumeSourceApplyConfiguration) WithInitiatorName(value string) *ISCSIVolumeSourceApplyConfiguration {
	b.InitiatorName = &value
	return b
}
