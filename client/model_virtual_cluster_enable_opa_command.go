/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the VirtualClusterEnableOpaCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualClusterEnableOpaCommand{}

// VirtualClusterEnableOpaCommand struct for VirtualClusterEnableOpaCommand
type VirtualClusterEnableOpaCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	OpaProfileId *int32 `json:"opaProfileId,omitempty"`
}

// NewVirtualClusterEnableOpaCommand instantiates a new VirtualClusterEnableOpaCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualClusterEnableOpaCommand() *VirtualClusterEnableOpaCommand {
	this := VirtualClusterEnableOpaCommand{}
	return &this
}

// NewVirtualClusterEnableOpaCommandWithDefaults instantiates a new VirtualClusterEnableOpaCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualClusterEnableOpaCommandWithDefaults() *VirtualClusterEnableOpaCommand {
	this := VirtualClusterEnableOpaCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *VirtualClusterEnableOpaCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualClusterEnableOpaCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *VirtualClusterEnableOpaCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *VirtualClusterEnableOpaCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetOpaProfileId returns the OpaProfileId field value if set, zero value otherwise.
func (o *VirtualClusterEnableOpaCommand) GetOpaProfileId() int32 {
	if o == nil || IsNil(o.OpaProfileId) {
		var ret int32
		return ret
	}
	return *o.OpaProfileId
}

// GetOpaProfileIdOk returns a tuple with the OpaProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualClusterEnableOpaCommand) GetOpaProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OpaProfileId) {
		return nil, false
	}
	return o.OpaProfileId, true
}

// HasOpaProfileId returns a boolean if a field has been set.
func (o *VirtualClusterEnableOpaCommand) HasOpaProfileId() bool {
	if o != nil && !IsNil(o.OpaProfileId) {
		return true
	}

	return false
}

// SetOpaProfileId gets a reference to the given int32 and assigns it to the OpaProfileId field.
func (o *VirtualClusterEnableOpaCommand) SetOpaProfileId(v int32) {
	o.OpaProfileId = &v
}

func (o VirtualClusterEnableOpaCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualClusterEnableOpaCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.OpaProfileId) {
		toSerialize["opaProfileId"] = o.OpaProfileId
	}
	return toSerialize, nil
}

type NullableVirtualClusterEnableOpaCommand struct {
	value *VirtualClusterEnableOpaCommand
	isSet bool
}

func (v NullableVirtualClusterEnableOpaCommand) Get() *VirtualClusterEnableOpaCommand {
	return v.value
}

func (v *NullableVirtualClusterEnableOpaCommand) Set(val *VirtualClusterEnableOpaCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualClusterEnableOpaCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualClusterEnableOpaCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualClusterEnableOpaCommand(val *VirtualClusterEnableOpaCommand) *NullableVirtualClusterEnableOpaCommand {
	return &NullableVirtualClusterEnableOpaCommand{value: val, isSet: true}
}

func (v NullableVirtualClusterEnableOpaCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualClusterEnableOpaCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


