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

// checks if the VerifyEmailCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEmailCommand{}

// VerifyEmailCommand struct for VerifyEmailCommand
type VerifyEmailCommand struct {
	Token NullableString `json:"token,omitempty"`
	Mode *EmailMode `json:"mode,omitempty"`
}

// NewVerifyEmailCommand instantiates a new VerifyEmailCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEmailCommand() *VerifyEmailCommand {
	this := VerifyEmailCommand{}
	return &this
}

// NewVerifyEmailCommandWithDefaults instantiates a new VerifyEmailCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEmailCommandWithDefaults() *VerifyEmailCommand {
	this := VerifyEmailCommand{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyEmailCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyEmailCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyEmailCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *VerifyEmailCommand) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *VerifyEmailCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *VerifyEmailCommand) UnsetToken() {
	o.Token.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VerifyEmailCommand) GetMode() EmailMode {
	if o == nil || IsNil(o.Mode) {
		var ret EmailMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEmailCommand) GetModeOk() (*EmailMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VerifyEmailCommand) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given EmailMode and assigns it to the Mode field.
func (o *VerifyEmailCommand) SetMode(v EmailMode) {
	o.Mode = &v
}

func (o VerifyEmailCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEmailCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableVerifyEmailCommand struct {
	value *VerifyEmailCommand
	isSet bool
}

func (v NullableVerifyEmailCommand) Get() *VerifyEmailCommand {
	return v.value
}

func (v *NullableVerifyEmailCommand) Set(val *VerifyEmailCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEmailCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEmailCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEmailCommand(val *VerifyEmailCommand) *NullableVerifyEmailCommand {
	return &NullableVerifyEmailCommand{value: val, isSet: true}
}

func (v NullableVerifyEmailCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEmailCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


