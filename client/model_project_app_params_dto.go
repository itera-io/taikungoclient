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

// checks if the ProjectAppParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAppParamsDto{}

// ProjectAppParamsDto struct for ProjectAppParamsDto
type ProjectAppParamsDto struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewProjectAppParamsDto instantiates a new ProjectAppParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAppParamsDto() *ProjectAppParamsDto {
	this := ProjectAppParamsDto{}
	return &this
}

// NewProjectAppParamsDtoWithDefaults instantiates a new ProjectAppParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAppParamsDtoWithDefaults() *ProjectAppParamsDto {
	this := ProjectAppParamsDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProjectAppParamsDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamsDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectAppParamsDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProjectAppParamsDto) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProjectAppParamsDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamsDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectAppParamsDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProjectAppParamsDto) SetValue(v string) {
	o.Value = &v
}

func (o ProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAppParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableProjectAppParamsDto struct {
	value *ProjectAppParamsDto
	isSet bool
}

func (v NullableProjectAppParamsDto) Get() *ProjectAppParamsDto {
	return v.value
}

func (v *NullableProjectAppParamsDto) Set(val *ProjectAppParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAppParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAppParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAppParamsDto(val *ProjectAppParamsDto) *NullableProjectAppParamsDto {
	return &NullableProjectAppParamsDto{value: val, isSet: true}
}

func (v NullableProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAppParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


