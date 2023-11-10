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

// checks if the PurgeWholeProjectCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurgeWholeProjectCommand{}

// PurgeWholeProjectCommand struct for PurgeWholeProjectCommand
type PurgeWholeProjectCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewPurgeWholeProjectCommand instantiates a new PurgeWholeProjectCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeWholeProjectCommand() *PurgeWholeProjectCommand {
	this := PurgeWholeProjectCommand{}
	return &this
}

// NewPurgeWholeProjectCommandWithDefaults instantiates a new PurgeWholeProjectCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeWholeProjectCommandWithDefaults() *PurgeWholeProjectCommand {
	this := PurgeWholeProjectCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PurgeWholeProjectCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeWholeProjectCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PurgeWholeProjectCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PurgeWholeProjectCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o PurgeWholeProjectCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurgeWholeProjectCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullablePurgeWholeProjectCommand struct {
	value *PurgeWholeProjectCommand
	isSet bool
}

func (v NullablePurgeWholeProjectCommand) Get() *PurgeWholeProjectCommand {
	return v.value
}

func (v *NullablePurgeWholeProjectCommand) Set(val *PurgeWholeProjectCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeWholeProjectCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeWholeProjectCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeWholeProjectCommand(val *PurgeWholeProjectCommand) *NullablePurgeWholeProjectCommand {
	return &NullablePurgeWholeProjectCommand{value: val, isSet: true}
}

func (v NullablePurgeWholeProjectCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeWholeProjectCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


