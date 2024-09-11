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

// checks if the DisableUserCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableUserCommand{}

// DisableUserCommand struct for DisableUserCommand
type DisableUserCommand struct {
	Id *string `json:"id,omitempty"`
	Disable *bool `json:"disable,omitempty"`
}

// NewDisableUserCommand instantiates a new DisableUserCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableUserCommand() *DisableUserCommand {
	this := DisableUserCommand{}
	return &this
}

// NewDisableUserCommandWithDefaults instantiates a new DisableUserCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableUserCommandWithDefaults() *DisableUserCommand {
	this := DisableUserCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DisableUserCommand) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableUserCommand) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DisableUserCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DisableUserCommand) SetId(v string) {
	o.Id = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *DisableUserCommand) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableUserCommand) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *DisableUserCommand) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *DisableUserCommand) SetDisable(v bool) {
	o.Disable = &v
}

func (o DisableUserCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableUserCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	return toSerialize, nil
}

type NullableDisableUserCommand struct {
	value *DisableUserCommand
	isSet bool
}

func (v NullableDisableUserCommand) Get() *DisableUserCommand {
	return v.value
}

func (v *NullableDisableUserCommand) Set(val *DisableUserCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableUserCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableUserCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableUserCommand(val *DisableUserCommand) *NullableDisableUserCommand {
	return &NullableDisableUserCommand{value: val, isSet: true}
}

func (v NullableDisableUserCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableUserCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


