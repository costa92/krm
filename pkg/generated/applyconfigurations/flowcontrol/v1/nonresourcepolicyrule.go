// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NonResourcePolicyRuleApplyConfiguration represents an declarative configuration of the NonResourcePolicyRule type for use
// with apply.
type NonResourcePolicyRuleApplyConfiguration struct {
	Verbs           []string `json:"verbs,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
}

// NonResourcePolicyRuleApplyConfiguration constructs an declarative configuration of the NonResourcePolicyRule type for use with
// apply.
func NonResourcePolicyRule() *NonResourcePolicyRuleApplyConfiguration {
	return &NonResourcePolicyRuleApplyConfiguration{}
}

// WithVerbs adds the given value to the Verbs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Verbs field.
func (b *NonResourcePolicyRuleApplyConfiguration) WithVerbs(values ...string) *NonResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Verbs = append(b.Verbs, values[i])
	}
	return b
}

// WithNonResourceURLs adds the given value to the NonResourceURLs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NonResourceURLs field.
func (b *NonResourcePolicyRuleApplyConfiguration) WithNonResourceURLs(values ...string) *NonResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.NonResourceURLs = append(b.NonResourceURLs, values[i])
	}
	return b
}
