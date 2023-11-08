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

// checks if the UpdateServerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServerCommand{}

// UpdateServerCommand struct for UpdateServerCommand
type UpdateServerCommand struct {
	Servers []ServerUpdateDto `json:"servers,omitempty"`
}

// NewUpdateServerCommand instantiates a new UpdateServerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerCommand() *UpdateServerCommand {
	this := UpdateServerCommand{}
	return &this
}

// NewUpdateServerCommandWithDefaults instantiates a new UpdateServerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerCommandWithDefaults() *UpdateServerCommand {
	this := UpdateServerCommand{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerCommand) GetServers() []ServerUpdateDto {
	if o == nil {
		var ret []ServerUpdateDto
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerCommand) GetServersOk() ([]ServerUpdateDto, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *UpdateServerCommand) HasServers() bool {
	if o != nil && IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerUpdateDto and assigns it to the Servers field.
func (o *UpdateServerCommand) SetServers(v []ServerUpdateDto) {
	o.Servers = v
}

func (o UpdateServerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableUpdateServerCommand struct {
	value *UpdateServerCommand
	isSet bool
}

func (v NullableUpdateServerCommand) Get() *UpdateServerCommand {
	return v.value
}

func (v *NullableUpdateServerCommand) Set(val *UpdateServerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerCommand(val *UpdateServerCommand) *NullableUpdateServerCommand {
	return &NullableUpdateServerCommand{value: val, isSet: true}
}

func (v NullableUpdateServerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


