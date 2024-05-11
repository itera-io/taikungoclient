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

// checks if the DeleteVirtualClusterCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVirtualClusterCommand{}

// DeleteVirtualClusterCommand struct for DeleteVirtualClusterCommand
type DeleteVirtualClusterCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewDeleteVirtualClusterCommand instantiates a new DeleteVirtualClusterCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVirtualClusterCommand() *DeleteVirtualClusterCommand {
	this := DeleteVirtualClusterCommand{}
	return &this
}

// NewDeleteVirtualClusterCommandWithDefaults instantiates a new DeleteVirtualClusterCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVirtualClusterCommandWithDefaults() *DeleteVirtualClusterCommand {
	this := DeleteVirtualClusterCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteVirtualClusterCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVirtualClusterCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteVirtualClusterCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteVirtualClusterCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o DeleteVirtualClusterCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVirtualClusterCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableDeleteVirtualClusterCommand struct {
	value *DeleteVirtualClusterCommand
	isSet bool
}

func (v NullableDeleteVirtualClusterCommand) Get() *DeleteVirtualClusterCommand {
	return v.value
}

func (v *NullableDeleteVirtualClusterCommand) Set(val *DeleteVirtualClusterCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVirtualClusterCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVirtualClusterCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVirtualClusterCommand(val *DeleteVirtualClusterCommand) *NullableDeleteVirtualClusterCommand {
	return &NullableDeleteVirtualClusterCommand{value: val, isSet: true}
}

func (v NullableDeleteVirtualClusterCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVirtualClusterCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


