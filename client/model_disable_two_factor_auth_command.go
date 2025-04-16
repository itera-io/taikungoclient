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

// checks if the DisableTwoFactorAuthCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableTwoFactorAuthCommand{}

// DisableTwoFactorAuthCommand struct for DisableTwoFactorAuthCommand
type DisableTwoFactorAuthCommand struct {
	Code NullableString `json:"code,omitempty"`
	IsRecoveryCode *bool `json:"isRecoveryCode,omitempty"`
}

// NewDisableTwoFactorAuthCommand instantiates a new DisableTwoFactorAuthCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableTwoFactorAuthCommand() *DisableTwoFactorAuthCommand {
	this := DisableTwoFactorAuthCommand{}
	return &this
}

// NewDisableTwoFactorAuthCommandWithDefaults instantiates a new DisableTwoFactorAuthCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableTwoFactorAuthCommandWithDefaults() *DisableTwoFactorAuthCommand {
	this := DisableTwoFactorAuthCommand{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisableTwoFactorAuthCommand) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisableTwoFactorAuthCommand) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *DisableTwoFactorAuthCommand) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *DisableTwoFactorAuthCommand) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *DisableTwoFactorAuthCommand) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *DisableTwoFactorAuthCommand) UnsetCode() {
	o.Code.Unset()
}

// GetIsRecoveryCode returns the IsRecoveryCode field value if set, zero value otherwise.
func (o *DisableTwoFactorAuthCommand) GetIsRecoveryCode() bool {
	if o == nil || IsNil(o.IsRecoveryCode) {
		var ret bool
		return ret
	}
	return *o.IsRecoveryCode
}

// GetIsRecoveryCodeOk returns a tuple with the IsRecoveryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableTwoFactorAuthCommand) GetIsRecoveryCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRecoveryCode) {
		return nil, false
	}
	return o.IsRecoveryCode, true
}

// HasIsRecoveryCode returns a boolean if a field has been set.
func (o *DisableTwoFactorAuthCommand) HasIsRecoveryCode() bool {
	if o != nil && !IsNil(o.IsRecoveryCode) {
		return true
	}

	return false
}

// SetIsRecoveryCode gets a reference to the given bool and assigns it to the IsRecoveryCode field.
func (o *DisableTwoFactorAuthCommand) SetIsRecoveryCode(v bool) {
	o.IsRecoveryCode = &v
}

func (o DisableTwoFactorAuthCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableTwoFactorAuthCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if !IsNil(o.IsRecoveryCode) {
		toSerialize["isRecoveryCode"] = o.IsRecoveryCode
	}
	return toSerialize, nil
}

type NullableDisableTwoFactorAuthCommand struct {
	value *DisableTwoFactorAuthCommand
	isSet bool
}

func (v NullableDisableTwoFactorAuthCommand) Get() *DisableTwoFactorAuthCommand {
	return v.value
}

func (v *NullableDisableTwoFactorAuthCommand) Set(val *DisableTwoFactorAuthCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableTwoFactorAuthCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableTwoFactorAuthCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableTwoFactorAuthCommand(val *DisableTwoFactorAuthCommand) *NullableDisableTwoFactorAuthCommand {
	return &NullableDisableTwoFactorAuthCommand{value: val, isSet: true}
}

func (v NullableDisableTwoFactorAuthCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableTwoFactorAuthCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


