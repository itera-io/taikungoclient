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

// checks if the ConfirmEmailCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmEmailCommand{}

// ConfirmEmailCommand struct for ConfirmEmailCommand
type ConfirmEmailCommand struct {
	NewEmail *string `json:"newEmail,omitempty"`
	Mode *EmailMode `json:"mode,omitempty"`
}

// NewConfirmEmailCommand instantiates a new ConfirmEmailCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmEmailCommand() *ConfirmEmailCommand {
	this := ConfirmEmailCommand{}
	return &this
}

// NewConfirmEmailCommandWithDefaults instantiates a new ConfirmEmailCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmEmailCommandWithDefaults() *ConfirmEmailCommand {
	this := ConfirmEmailCommand{}
	return &this
}

// GetNewEmail returns the NewEmail field value if set, zero value otherwise.
func (o *ConfirmEmailCommand) GetNewEmail() string {
	if o == nil || IsNil(o.NewEmail) {
		var ret string
		return ret
	}
	return *o.NewEmail
}

// GetNewEmailOk returns a tuple with the NewEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmEmailCommand) GetNewEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NewEmail) {
		return nil, false
	}
	return o.NewEmail, true
}

// HasNewEmail returns a boolean if a field has been set.
func (o *ConfirmEmailCommand) HasNewEmail() bool {
	if o != nil && !IsNil(o.NewEmail) {
		return true
	}

	return false
}

// SetNewEmail gets a reference to the given string and assigns it to the NewEmail field.
func (o *ConfirmEmailCommand) SetNewEmail(v string) {
	o.NewEmail = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ConfirmEmailCommand) GetMode() EmailMode {
	if o == nil || IsNil(o.Mode) {
		var ret EmailMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmEmailCommand) GetModeOk() (*EmailMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ConfirmEmailCommand) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given EmailMode and assigns it to the Mode field.
func (o *ConfirmEmailCommand) SetMode(v EmailMode) {
	o.Mode = &v
}

func (o ConfirmEmailCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmEmailCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewEmail) {
		toSerialize["newEmail"] = o.NewEmail
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableConfirmEmailCommand struct {
	value *ConfirmEmailCommand
	isSet bool
}

func (v NullableConfirmEmailCommand) Get() *ConfirmEmailCommand {
	return v.value
}

func (v *NullableConfirmEmailCommand) Set(val *ConfirmEmailCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmEmailCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmEmailCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmEmailCommand(val *ConfirmEmailCommand) *NullableConfirmEmailCommand {
	return &NullableConfirmEmailCommand{value: val, isSet: true}
}

func (v NullableConfirmEmailCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmEmailCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


