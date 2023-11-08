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

// checks if the ToggleDemoModeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToggleDemoModeCommand{}

// ToggleDemoModeCommand struct for ToggleDemoModeCommand
type ToggleDemoModeCommand struct {
	Mode NullableString `json:"mode,omitempty"`
}

// NewToggleDemoModeCommand instantiates a new ToggleDemoModeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToggleDemoModeCommand() *ToggleDemoModeCommand {
	this := ToggleDemoModeCommand{}
	return &this
}

// NewToggleDemoModeCommandWithDefaults instantiates a new ToggleDemoModeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToggleDemoModeCommandWithDefaults() *ToggleDemoModeCommand {
	this := ToggleDemoModeCommand{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToggleDemoModeCommand) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToggleDemoModeCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *ToggleDemoModeCommand) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *ToggleDemoModeCommand) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *ToggleDemoModeCommand) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *ToggleDemoModeCommand) UnsetMode() {
	o.Mode.Unset()
}

func (o ToggleDemoModeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToggleDemoModeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableToggleDemoModeCommand struct {
	value *ToggleDemoModeCommand
	isSet bool
}

func (v NullableToggleDemoModeCommand) Get() *ToggleDemoModeCommand {
	return v.value
}

func (v *NullableToggleDemoModeCommand) Set(val *ToggleDemoModeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableToggleDemoModeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableToggleDemoModeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToggleDemoModeCommand(val *ToggleDemoModeCommand) *NullableToggleDemoModeCommand {
	return &NullableToggleDemoModeCommand{value: val, isSet: true}
}

func (v NullableToggleDemoModeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToggleDemoModeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


