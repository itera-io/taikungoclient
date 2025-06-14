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

// checks if the TwoFactorAuthVerifyCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoFactorAuthVerifyCommand{}

// TwoFactorAuthVerifyCommand struct for TwoFactorAuthVerifyCommand
type TwoFactorAuthVerifyCommand struct {
	VerificationCode NullableString `json:"verificationCode,omitempty"`
}

// NewTwoFactorAuthVerifyCommand instantiates a new TwoFactorAuthVerifyCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoFactorAuthVerifyCommand() *TwoFactorAuthVerifyCommand {
	this := TwoFactorAuthVerifyCommand{}
	return &this
}

// NewTwoFactorAuthVerifyCommandWithDefaults instantiates a new TwoFactorAuthVerifyCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoFactorAuthVerifyCommandWithDefaults() *TwoFactorAuthVerifyCommand {
	this := TwoFactorAuthVerifyCommand{}
	return &this
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoFactorAuthVerifyCommand) GetVerificationCode() string {
	if o == nil || IsNil(o.VerificationCode.Get()) {
		var ret string
		return ret
	}
	return *o.VerificationCode.Get()
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoFactorAuthVerifyCommand) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationCode.Get(), o.VerificationCode.IsSet()
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *TwoFactorAuthVerifyCommand) HasVerificationCode() bool {
	if o != nil && o.VerificationCode.IsSet() {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given NullableString and assigns it to the VerificationCode field.
func (o *TwoFactorAuthVerifyCommand) SetVerificationCode(v string) {
	o.VerificationCode.Set(&v)
}
// SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil
func (o *TwoFactorAuthVerifyCommand) SetVerificationCodeNil() {
	o.VerificationCode.Set(nil)
}

// UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil
func (o *TwoFactorAuthVerifyCommand) UnsetVerificationCode() {
	o.VerificationCode.Unset()
}

func (o TwoFactorAuthVerifyCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoFactorAuthVerifyCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VerificationCode.IsSet() {
		toSerialize["verificationCode"] = o.VerificationCode.Get()
	}
	return toSerialize, nil
}

type NullableTwoFactorAuthVerifyCommand struct {
	value *TwoFactorAuthVerifyCommand
	isSet bool
}

func (v NullableTwoFactorAuthVerifyCommand) Get() *TwoFactorAuthVerifyCommand {
	return v.value
}

func (v *NullableTwoFactorAuthVerifyCommand) Set(val *TwoFactorAuthVerifyCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoFactorAuthVerifyCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoFactorAuthVerifyCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoFactorAuthVerifyCommand(val *TwoFactorAuthVerifyCommand) *NullableTwoFactorAuthVerifyCommand {
	return &NullableTwoFactorAuthVerifyCommand{value: val, isSet: true}
}

func (v NullableTwoFactorAuthVerifyCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoFactorAuthVerifyCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


