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

// checks if the VirtualClusterDisableBackupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualClusterDisableBackupCommand{}

// VirtualClusterDisableBackupCommand struct for VirtualClusterDisableBackupCommand
type VirtualClusterDisableBackupCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewVirtualClusterDisableBackupCommand instantiates a new VirtualClusterDisableBackupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualClusterDisableBackupCommand() *VirtualClusterDisableBackupCommand {
	this := VirtualClusterDisableBackupCommand{}
	return &this
}

// NewVirtualClusterDisableBackupCommandWithDefaults instantiates a new VirtualClusterDisableBackupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualClusterDisableBackupCommandWithDefaults() *VirtualClusterDisableBackupCommand {
	this := VirtualClusterDisableBackupCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *VirtualClusterDisableBackupCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualClusterDisableBackupCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *VirtualClusterDisableBackupCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *VirtualClusterDisableBackupCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o VirtualClusterDisableBackupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualClusterDisableBackupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableVirtualClusterDisableBackupCommand struct {
	value *VirtualClusterDisableBackupCommand
	isSet bool
}

func (v NullableVirtualClusterDisableBackupCommand) Get() *VirtualClusterDisableBackupCommand {
	return v.value
}

func (v *NullableVirtualClusterDisableBackupCommand) Set(val *VirtualClusterDisableBackupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualClusterDisableBackupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualClusterDisableBackupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualClusterDisableBackupCommand(val *VirtualClusterDisableBackupCommand) *NullableVirtualClusterDisableBackupCommand {
	return &NullableVirtualClusterDisableBackupCommand{value: val, isSet: true}
}

func (v NullableVirtualClusterDisableBackupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualClusterDisableBackupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


