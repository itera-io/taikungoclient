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

// checks if the BindImageToProjectCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindImageToProjectCommand{}

// BindImageToProjectCommand struct for BindImageToProjectCommand
type BindImageToProjectCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Images []string `json:"images,omitempty"`
}

// NewBindImageToProjectCommand instantiates a new BindImageToProjectCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindImageToProjectCommand() *BindImageToProjectCommand {
	this := BindImageToProjectCommand{}
	return &this
}

// NewBindImageToProjectCommandWithDefaults instantiates a new BindImageToProjectCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindImageToProjectCommandWithDefaults() *BindImageToProjectCommand {
	this := BindImageToProjectCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BindImageToProjectCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindImageToProjectCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BindImageToProjectCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *BindImageToProjectCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindImageToProjectCommand) GetImages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindImageToProjectCommand) GetImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *BindImageToProjectCommand) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []string and assigns it to the Images field.
func (o *BindImageToProjectCommand) SetImages(v []string) {
	o.Images = v
}

func (o BindImageToProjectCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindImageToProjectCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableBindImageToProjectCommand struct {
	value *BindImageToProjectCommand
	isSet bool
}

func (v NullableBindImageToProjectCommand) Get() *BindImageToProjectCommand {
	return v.value
}

func (v *NullableBindImageToProjectCommand) Set(val *BindImageToProjectCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindImageToProjectCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindImageToProjectCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindImageToProjectCommand(val *BindImageToProjectCommand) *NullableBindImageToProjectCommand {
	return &NullableBindImageToProjectCommand{value: val, isSet: true}
}

func (v NullableBindImageToProjectCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindImageToProjectCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


