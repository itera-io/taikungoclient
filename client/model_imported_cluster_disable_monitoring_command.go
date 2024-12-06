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

// checks if the ImportedClusterDisableMonitoringCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedClusterDisableMonitoringCommand{}

// ImportedClusterDisableMonitoringCommand struct for ImportedClusterDisableMonitoringCommand
type ImportedClusterDisableMonitoringCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewImportedClusterDisableMonitoringCommand instantiates a new ImportedClusterDisableMonitoringCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedClusterDisableMonitoringCommand() *ImportedClusterDisableMonitoringCommand {
	this := ImportedClusterDisableMonitoringCommand{}
	return &this
}

// NewImportedClusterDisableMonitoringCommandWithDefaults instantiates a new ImportedClusterDisableMonitoringCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedClusterDisableMonitoringCommandWithDefaults() *ImportedClusterDisableMonitoringCommand {
	this := ImportedClusterDisableMonitoringCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ImportedClusterDisableMonitoringCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportedClusterDisableMonitoringCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ImportedClusterDisableMonitoringCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ImportedClusterDisableMonitoringCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o ImportedClusterDisableMonitoringCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedClusterDisableMonitoringCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableImportedClusterDisableMonitoringCommand struct {
	value *ImportedClusterDisableMonitoringCommand
	isSet bool
}

func (v NullableImportedClusterDisableMonitoringCommand) Get() *ImportedClusterDisableMonitoringCommand {
	return v.value
}

func (v *NullableImportedClusterDisableMonitoringCommand) Set(val *ImportedClusterDisableMonitoringCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedClusterDisableMonitoringCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedClusterDisableMonitoringCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedClusterDisableMonitoringCommand(val *ImportedClusterDisableMonitoringCommand) *NullableImportedClusterDisableMonitoringCommand {
	return &NullableImportedClusterDisableMonitoringCommand{value: val, isSet: true}
}

func (v NullableImportedClusterDisableMonitoringCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedClusterDisableMonitoringCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


