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

// checks if the BindFlavorToProjectCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindFlavorToProjectCommand{}

// BindFlavorToProjectCommand struct for BindFlavorToProjectCommand
type BindFlavorToProjectCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Flavors []string `json:"flavors,omitempty"`
}

// NewBindFlavorToProjectCommand instantiates a new BindFlavorToProjectCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindFlavorToProjectCommand() *BindFlavorToProjectCommand {
	this := BindFlavorToProjectCommand{}
	return &this
}

// NewBindFlavorToProjectCommandWithDefaults instantiates a new BindFlavorToProjectCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindFlavorToProjectCommandWithDefaults() *BindFlavorToProjectCommand {
	this := BindFlavorToProjectCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BindFlavorToProjectCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindFlavorToProjectCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BindFlavorToProjectCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *BindFlavorToProjectCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetFlavors returns the Flavors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindFlavorToProjectCommand) GetFlavors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Flavors
}

// GetFlavorsOk returns a tuple with the Flavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindFlavorToProjectCommand) GetFlavorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flavors) {
		return nil, false
	}
	return o.Flavors, true
}

// HasFlavors returns a boolean if a field has been set.
func (o *BindFlavorToProjectCommand) HasFlavors() bool {
	if o != nil && !IsNil(o.Flavors) {
		return true
	}

	return false
}

// SetFlavors gets a reference to the given []string and assigns it to the Flavors field.
func (o *BindFlavorToProjectCommand) SetFlavors(v []string) {
	o.Flavors = v
}

func (o BindFlavorToProjectCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindFlavorToProjectCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Flavors != nil {
		toSerialize["flavors"] = o.Flavors
	}
	return toSerialize, nil
}

type NullableBindFlavorToProjectCommand struct {
	value *BindFlavorToProjectCommand
	isSet bool
}

func (v NullableBindFlavorToProjectCommand) Get() *BindFlavorToProjectCommand {
	return v.value
}

func (v *NullableBindFlavorToProjectCommand) Set(val *BindFlavorToProjectCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindFlavorToProjectCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindFlavorToProjectCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindFlavorToProjectCommand(val *BindFlavorToProjectCommand) *NullableBindFlavorToProjectCommand {
	return &NullableBindFlavorToProjectCommand{value: val, isSet: true}
}

func (v NullableBindFlavorToProjectCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindFlavorToProjectCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


