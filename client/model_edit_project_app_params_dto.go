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

// checks if the EditProjectAppParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProjectAppParamsDto{}

// EditProjectAppParamsDto struct for EditProjectAppParamsDto
type EditProjectAppParamsDto struct {
	Key   NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewEditProjectAppParamsDto instantiates a new EditProjectAppParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProjectAppParamsDto() *EditProjectAppParamsDto {
	this := EditProjectAppParamsDto{}
	return &this
}

// NewEditProjectAppParamsDtoWithDefaults instantiates a new EditProjectAppParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProjectAppParamsDtoWithDefaults() *EditProjectAppParamsDto {
	this := EditProjectAppParamsDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProjectAppParamsDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProjectAppParamsDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *EditProjectAppParamsDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *EditProjectAppParamsDto) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *EditProjectAppParamsDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *EditProjectAppParamsDto) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProjectAppParamsDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProjectAppParamsDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EditProjectAppParamsDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EditProjectAppParamsDto) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *EditProjectAppParamsDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EditProjectAppParamsDto) UnsetValue() {
	o.Value.Unset()
}

func (o EditProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProjectAppParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableEditProjectAppParamsDto struct {
	value *EditProjectAppParamsDto
	isSet bool
}

func (v NullableEditProjectAppParamsDto) Get() *EditProjectAppParamsDto {
	return v.value
}

func (v *NullableEditProjectAppParamsDto) Set(val *EditProjectAppParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProjectAppParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProjectAppParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProjectAppParamsDto(val *EditProjectAppParamsDto) *NullableEditProjectAppParamsDto {
	return &NullableEditProjectAppParamsDto{value: val, isSet: true}
}

func (v NullableEditProjectAppParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProjectAppParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
