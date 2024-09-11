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

// checks if the SshKeyCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshKeyCommand{}

// SshKeyCommand struct for SshKeyCommand
type SshKeyCommand struct {
	SshPublicKey *string `json:"sshPublicKey,omitempty"`
}

// NewSshKeyCommand instantiates a new SshKeyCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyCommand() *SshKeyCommand {
	this := SshKeyCommand{}
	return &this
}

// NewSshKeyCommandWithDefaults instantiates a new SshKeyCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyCommandWithDefaults() *SshKeyCommand {
	this := SshKeyCommand{}
	return &this
}

// GetSshPublicKey returns the SshPublicKey field value if set, zero value otherwise.
func (o *SshKeyCommand) GetSshPublicKey() string {
	if o == nil || IsNil(o.SshPublicKey) {
		var ret string
		return ret
	}
	return *o.SshPublicKey
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyCommand) GetSshPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SshPublicKey) {
		return nil, false
	}
	return o.SshPublicKey, true
}

// HasSshPublicKey returns a boolean if a field has been set.
func (o *SshKeyCommand) HasSshPublicKey() bool {
	if o != nil && !IsNil(o.SshPublicKey) {
		return true
	}

	return false
}

// SetSshPublicKey gets a reference to the given string and assigns it to the SshPublicKey field.
func (o *SshKeyCommand) SetSshPublicKey(v string) {
	o.SshPublicKey = &v
}

func (o SshKeyCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SshPublicKey) {
		toSerialize["sshPublicKey"] = o.SshPublicKey
	}
	return toSerialize, nil
}

type NullableSshKeyCommand struct {
	value *SshKeyCommand
	isSet bool
}

func (v NullableSshKeyCommand) Get() *SshKeyCommand {
	return v.value
}

func (v *NullableSshKeyCommand) Set(val *SshKeyCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyCommand(val *SshKeyCommand) *NullableSshKeyCommand {
	return &NullableSshKeyCommand{value: val, isSet: true}
}

func (v NullableSshKeyCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


