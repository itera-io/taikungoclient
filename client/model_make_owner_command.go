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

// checks if the MakeOwnerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MakeOwnerCommand{}

// MakeOwnerCommand struct for MakeOwnerCommand
type MakeOwnerCommand struct {
	UserId NullableString `json:"userId,omitempty"`
}

// NewMakeOwnerCommand instantiates a new MakeOwnerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMakeOwnerCommand() *MakeOwnerCommand {
	this := MakeOwnerCommand{}
	return &this
}

// NewMakeOwnerCommandWithDefaults instantiates a new MakeOwnerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMakeOwnerCommandWithDefaults() *MakeOwnerCommand {
	this := MakeOwnerCommand{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MakeOwnerCommand) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MakeOwnerCommand) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MakeOwnerCommand) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MakeOwnerCommand) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MakeOwnerCommand) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MakeOwnerCommand) UnsetUserId() {
	o.UserId.Unset()
}

func (o MakeOwnerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MakeOwnerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return toSerialize, nil
}

type NullableMakeOwnerCommand struct {
	value *MakeOwnerCommand
	isSet bool
}

func (v NullableMakeOwnerCommand) Get() *MakeOwnerCommand {
	return v.value
}

func (v *NullableMakeOwnerCommand) Set(val *MakeOwnerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMakeOwnerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMakeOwnerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMakeOwnerCommand(val *MakeOwnerCommand) *NullableMakeOwnerCommand {
	return &NullableMakeOwnerCommand{value: val, isSet: true}
}

func (v NullableMakeOwnerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMakeOwnerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
