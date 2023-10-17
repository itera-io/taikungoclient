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

// checks if the BecomePartnerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BecomePartnerCommand{}

// BecomePartnerCommand struct for BecomePartnerCommand
type BecomePartnerCommand struct {
	FullName NullableString `json:"fullName,omitempty"`
	Email    NullableString `json:"email,omitempty"`
}

// NewBecomePartnerCommand instantiates a new BecomePartnerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBecomePartnerCommand() *BecomePartnerCommand {
	this := BecomePartnerCommand{}
	return &this
}

// NewBecomePartnerCommandWithDefaults instantiates a new BecomePartnerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBecomePartnerCommandWithDefaults() *BecomePartnerCommand {
	this := BecomePartnerCommand{}
	return &this
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BecomePartnerCommand) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BecomePartnerCommand) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *BecomePartnerCommand) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *BecomePartnerCommand) SetFullName(v string) {
	o.FullName.Set(&v)
}

// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *BecomePartnerCommand) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *BecomePartnerCommand) UnsetFullName() {
	o.FullName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BecomePartnerCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BecomePartnerCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *BecomePartnerCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *BecomePartnerCommand) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *BecomePartnerCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *BecomePartnerCommand) UnsetEmail() {
	o.Email.Unset()
}

func (o BecomePartnerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BecomePartnerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableBecomePartnerCommand struct {
	value *BecomePartnerCommand
	isSet bool
}

func (v NullableBecomePartnerCommand) Get() *BecomePartnerCommand {
	return v.value
}

func (v *NullableBecomePartnerCommand) Set(val *BecomePartnerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBecomePartnerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBecomePartnerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBecomePartnerCommand(val *BecomePartnerCommand) *NullableBecomePartnerCommand {
	return &NullableBecomePartnerCommand{value: val, isSet: true}
}

func (v NullableBecomePartnerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBecomePartnerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
