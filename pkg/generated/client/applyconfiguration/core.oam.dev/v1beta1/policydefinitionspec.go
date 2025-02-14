/*
Copyright 2023 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	common "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

// PolicyDefinitionSpecApplyConfiguration represents an declarative configuration of the PolicyDefinitionSpec type for use
// with apply.
type PolicyDefinitionSpecApplyConfiguration struct {
	Reference         *common.DefinitionReference `json:"definitionRef,omitempty"`
	Schematic         *common.Schematic           `json:"schematic,omitempty"`
	ManageHealthCheck *bool                       `json:"manageHealthCheck,omitempty"`
	Version           *string                     `json:"version,omitempty"`
}

// PolicyDefinitionSpecApplyConfiguration constructs an declarative configuration of the PolicyDefinitionSpec type for use with
// apply.
func PolicyDefinitionSpec() *PolicyDefinitionSpecApplyConfiguration {
	return &PolicyDefinitionSpecApplyConfiguration{}
}

// WithReference sets the Reference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reference field is set to the value of the last call.
func (b *PolicyDefinitionSpecApplyConfiguration) WithReference(value common.DefinitionReference) *PolicyDefinitionSpecApplyConfiguration {
	b.Reference = &value
	return b
}

// WithSchematic sets the Schematic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Schematic field is set to the value of the last call.
func (b *PolicyDefinitionSpecApplyConfiguration) WithSchematic(value common.Schematic) *PolicyDefinitionSpecApplyConfiguration {
	b.Schematic = &value
	return b
}

// WithManageHealthCheck sets the ManageHealthCheck field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManageHealthCheck field is set to the value of the last call.
func (b *PolicyDefinitionSpecApplyConfiguration) WithManageHealthCheck(value bool) *PolicyDefinitionSpecApplyConfiguration {
	b.ManageHealthCheck = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *PolicyDefinitionSpecApplyConfiguration) WithVersion(value string) *PolicyDefinitionSpecApplyConfiguration {
	b.Version = &value
	return b
}
