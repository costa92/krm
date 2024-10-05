// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SELinuxOptionsApplyConfiguration represents an declarative configuration of the SELinuxOptions type for use
// with apply.
type SELinuxOptionsApplyConfiguration struct {
	User  *string `json:"user,omitempty"`
	Role  *string `json:"role,omitempty"`
	Type  *string `json:"type,omitempty"`
	Level *string `json:"level,omitempty"`
}

// SELinuxOptionsApplyConfiguration constructs an declarative configuration of the SELinuxOptions type for use with
// apply.
func SELinuxOptions() *SELinuxOptionsApplyConfiguration {
	return &SELinuxOptionsApplyConfiguration{}
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *SELinuxOptionsApplyConfiguration) WithUser(value string) *SELinuxOptionsApplyConfiguration {
	b.User = &value
	return b
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *SELinuxOptionsApplyConfiguration) WithRole(value string) *SELinuxOptionsApplyConfiguration {
	b.Role = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *SELinuxOptionsApplyConfiguration) WithType(value string) *SELinuxOptionsApplyConfiguration {
	b.Type = &value
	return b
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *SELinuxOptionsApplyConfiguration) WithLevel(value string) *SELinuxOptionsApplyConfiguration {
	b.Level = &value
	return b
}
