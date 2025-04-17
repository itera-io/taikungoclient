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

// checks if the UninstallProjectAppResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UninstallProjectAppResult{}

// UninstallProjectAppResult struct for UninstallProjectAppResult
type UninstallProjectAppResult struct {
	IsJobSkipped *bool `json:"isJobSkipped,omitempty"`
}

// NewUninstallProjectAppResult instantiates a new UninstallProjectAppResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUninstallProjectAppResult() *UninstallProjectAppResult {
	this := UninstallProjectAppResult{}
	return &this
}

// NewUninstallProjectAppResultWithDefaults instantiates a new UninstallProjectAppResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUninstallProjectAppResultWithDefaults() *UninstallProjectAppResult {
	this := UninstallProjectAppResult{}
	return &this
}

// GetIsJobSkipped returns the IsJobSkipped field value if set, zero value otherwise.
func (o *UninstallProjectAppResult) GetIsJobSkipped() bool {
	if o == nil || IsNil(o.IsJobSkipped) {
		var ret bool
		return ret
	}
	return *o.IsJobSkipped
}

// GetIsJobSkippedOk returns a tuple with the IsJobSkipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallProjectAppResult) GetIsJobSkippedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsJobSkipped) {
		return nil, false
	}
	return o.IsJobSkipped, true
}

// HasIsJobSkipped returns a boolean if a field has been set.
func (o *UninstallProjectAppResult) HasIsJobSkipped() bool {
	if o != nil && !IsNil(o.IsJobSkipped) {
		return true
	}

	return false
}

// SetIsJobSkipped gets a reference to the given bool and assigns it to the IsJobSkipped field.
func (o *UninstallProjectAppResult) SetIsJobSkipped(v bool) {
	o.IsJobSkipped = &v
}

func (o UninstallProjectAppResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UninstallProjectAppResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsJobSkipped) {
		toSerialize["isJobSkipped"] = o.IsJobSkipped
	}
	return toSerialize, nil
}

type NullableUninstallProjectAppResult struct {
	value *UninstallProjectAppResult
	isSet bool
}

func (v NullableUninstallProjectAppResult) Get() *UninstallProjectAppResult {
	return v.value
}

func (v *NullableUninstallProjectAppResult) Set(val *UninstallProjectAppResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUninstallProjectAppResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUninstallProjectAppResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUninstallProjectAppResult(val *UninstallProjectAppResult) *NullableUninstallProjectAppResult {
	return &NullableUninstallProjectAppResult{value: val, isSet: true}
}

func (v NullableUninstallProjectAppResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUninstallProjectAppResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


