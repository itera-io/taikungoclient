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

// checks if the EditProjectAppExtraValuesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditProjectAppExtraValuesCommand{}

// EditProjectAppExtraValuesCommand struct for EditProjectAppExtraValuesCommand
type EditProjectAppExtraValuesCommand struct {
	ProjectAppId *int32 `json:"projectAppId,omitempty"`
	ExtraValues NullableString `json:"extraValues,omitempty"`
}

// NewEditProjectAppExtraValuesCommand instantiates a new EditProjectAppExtraValuesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditProjectAppExtraValuesCommand() *EditProjectAppExtraValuesCommand {
	this := EditProjectAppExtraValuesCommand{}
	return &this
}

// NewEditProjectAppExtraValuesCommandWithDefaults instantiates a new EditProjectAppExtraValuesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditProjectAppExtraValuesCommandWithDefaults() *EditProjectAppExtraValuesCommand {
	this := EditProjectAppExtraValuesCommand{}
	return &this
}

// GetProjectAppId returns the ProjectAppId field value if set, zero value otherwise.
func (o *EditProjectAppExtraValuesCommand) GetProjectAppId() int32 {
	if o == nil || IsNil(o.ProjectAppId) {
		var ret int32
		return ret
	}
	return *o.ProjectAppId
}

// GetProjectAppIdOk returns a tuple with the ProjectAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditProjectAppExtraValuesCommand) GetProjectAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectAppId) {
		return nil, false
	}
	return o.ProjectAppId, true
}

// HasProjectAppId returns a boolean if a field has been set.
func (o *EditProjectAppExtraValuesCommand) HasProjectAppId() bool {
	if o != nil && !IsNil(o.ProjectAppId) {
		return true
	}

	return false
}

// SetProjectAppId gets a reference to the given int32 and assigns it to the ProjectAppId field.
func (o *EditProjectAppExtraValuesCommand) SetProjectAppId(v int32) {
	o.ProjectAppId = &v
}

// GetExtraValues returns the ExtraValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditProjectAppExtraValuesCommand) GetExtraValues() string {
	if o == nil || IsNil(o.ExtraValues.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValues.Get()
}

// GetExtraValuesOk returns a tuple with the ExtraValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditProjectAppExtraValuesCommand) GetExtraValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValues.Get(), o.ExtraValues.IsSet()
}

// HasExtraValues returns a boolean if a field has been set.
func (o *EditProjectAppExtraValuesCommand) HasExtraValues() bool {
	if o != nil && o.ExtraValues.IsSet() {
		return true
	}

	return false
}

// SetExtraValues gets a reference to the given NullableString and assigns it to the ExtraValues field.
func (o *EditProjectAppExtraValuesCommand) SetExtraValues(v string) {
	o.ExtraValues.Set(&v)
}
// SetExtraValuesNil sets the value for ExtraValues to be an explicit nil
func (o *EditProjectAppExtraValuesCommand) SetExtraValuesNil() {
	o.ExtraValues.Set(nil)
}

// UnsetExtraValues ensures that no value is present for ExtraValues, not even an explicit nil
func (o *EditProjectAppExtraValuesCommand) UnsetExtraValues() {
	o.ExtraValues.Unset()
}

func (o EditProjectAppExtraValuesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditProjectAppExtraValuesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectAppId) {
		toSerialize["projectAppId"] = o.ProjectAppId
	}
	if o.ExtraValues.IsSet() {
		toSerialize["extraValues"] = o.ExtraValues.Get()
	}
	return toSerialize, nil
}

type NullableEditProjectAppExtraValuesCommand struct {
	value *EditProjectAppExtraValuesCommand
	isSet bool
}

func (v NullableEditProjectAppExtraValuesCommand) Get() *EditProjectAppExtraValuesCommand {
	return v.value
}

func (v *NullableEditProjectAppExtraValuesCommand) Set(val *EditProjectAppExtraValuesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditProjectAppExtraValuesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditProjectAppExtraValuesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditProjectAppExtraValuesCommand(val *EditProjectAppExtraValuesCommand) *NullableEditProjectAppExtraValuesCommand {
	return &NullableEditProjectAppExtraValuesCommand{value: val, isSet: true}
}

func (v NullableEditProjectAppExtraValuesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditProjectAppExtraValuesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


