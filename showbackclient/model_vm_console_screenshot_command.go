/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the VmConsoleScreenshotCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmConsoleScreenshotCommand{}

// VmConsoleScreenshotCommand struct for VmConsoleScreenshotCommand
type VmConsoleScreenshotCommand struct {
	ServerId *int32 `json:"serverId,omitempty"`
}

// NewVmConsoleScreenshotCommand instantiates a new VmConsoleScreenshotCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmConsoleScreenshotCommand() *VmConsoleScreenshotCommand {
	this := VmConsoleScreenshotCommand{}
	return &this
}

// NewVmConsoleScreenshotCommandWithDefaults instantiates a new VmConsoleScreenshotCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmConsoleScreenshotCommandWithDefaults() *VmConsoleScreenshotCommand {
	this := VmConsoleScreenshotCommand{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *VmConsoleScreenshotCommand) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmConsoleScreenshotCommand) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *VmConsoleScreenshotCommand) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *VmConsoleScreenshotCommand) SetServerId(v int32) {
	o.ServerId = &v
}

func (o VmConsoleScreenshotCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmConsoleScreenshotCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	return toSerialize, nil
}

type NullableVmConsoleScreenshotCommand struct {
	value *VmConsoleScreenshotCommand
	isSet bool
}

func (v NullableVmConsoleScreenshotCommand) Get() *VmConsoleScreenshotCommand {
	return v.value
}

func (v *NullableVmConsoleScreenshotCommand) Set(val *VmConsoleScreenshotCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVmConsoleScreenshotCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVmConsoleScreenshotCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmConsoleScreenshotCommand(val *VmConsoleScreenshotCommand) *NullableVmConsoleScreenshotCommand {
	return &NullableVmConsoleScreenshotCommand{value: val, isSet: true}
}

func (v NullableVmConsoleScreenshotCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmConsoleScreenshotCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
