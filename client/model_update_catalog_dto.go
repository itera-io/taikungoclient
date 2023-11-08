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

// checks if the UpdateCatalogDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCatalogDto{}

// UpdateCatalogDto struct for UpdateCatalogDto
type UpdateCatalogDto struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	IsBound   *bool  `json:"isBound,omitempty"`
}

// NewUpdateCatalogDto instantiates a new UpdateCatalogDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCatalogDto() *UpdateCatalogDto {
	this := UpdateCatalogDto{}
	return &this
}

// NewUpdateCatalogDtoWithDefaults instantiates a new UpdateCatalogDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCatalogDtoWithDefaults() *UpdateCatalogDto {
	this := UpdateCatalogDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *UpdateCatalogDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCatalogDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *UpdateCatalogDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *UpdateCatalogDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *UpdateCatalogDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCatalogDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *UpdateCatalogDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *UpdateCatalogDto) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o UpdateCatalogDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCatalogDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableUpdateCatalogDto struct {
	value *UpdateCatalogDto
	isSet bool
}

func (v NullableUpdateCatalogDto) Get() *UpdateCatalogDto {
	return v.value
}

func (v *NullableUpdateCatalogDto) Set(val *UpdateCatalogDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCatalogDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCatalogDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCatalogDto(val *UpdateCatalogDto) *NullableUpdateCatalogDto {
	return &NullableUpdateCatalogDto{value: val, isSet: true}
}

func (v NullableUpdateCatalogDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCatalogDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
