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

// checks if the RebootServerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RebootServerCommand{}

// RebootServerCommand struct for RebootServerCommand
type RebootServerCommand struct {
	ServerId *int32 `json:"serverId,omitempty"`
	Type NullableString `json:"type,omitempty"`
}

// NewRebootServerCommand instantiates a new RebootServerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootServerCommand() *RebootServerCommand {
	this := RebootServerCommand{}
	return &this
}

// NewRebootServerCommandWithDefaults instantiates a new RebootServerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootServerCommandWithDefaults() *RebootServerCommand {
	this := RebootServerCommand{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *RebootServerCommand) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootServerCommand) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *RebootServerCommand) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *RebootServerCommand) SetServerId(v int32) {
	o.ServerId = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RebootServerCommand) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RebootServerCommand) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *RebootServerCommand) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *RebootServerCommand) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *RebootServerCommand) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *RebootServerCommand) UnsetType() {
	o.Type.Unset()
}

func (o RebootServerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RebootServerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableRebootServerCommand struct {
	value *RebootServerCommand
	isSet bool
}

func (v NullableRebootServerCommand) Get() *RebootServerCommand {
	return v.value
}

func (v *NullableRebootServerCommand) Set(val *RebootServerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootServerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootServerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootServerCommand(val *RebootServerCommand) *NullableRebootServerCommand {
	return &NullableRebootServerCommand{value: val, isSet: true}
}

func (v NullableRebootServerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootServerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


