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

// checks if the CreateDnsServerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDnsServerCommand{}

// CreateDnsServerCommand struct for CreateDnsServerCommand
type CreateDnsServerCommand struct {
	Address         NullableString `json:"address,omitempty"`
	AccessProfileId *int32         `json:"accessProfileId,omitempty"`
}

// NewCreateDnsServerCommand instantiates a new CreateDnsServerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDnsServerCommand() *CreateDnsServerCommand {
	this := CreateDnsServerCommand{}
	return &this
}

// NewCreateDnsServerCommandWithDefaults instantiates a new CreateDnsServerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDnsServerCommandWithDefaults() *CreateDnsServerCommand {
	this := CreateDnsServerCommand{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDnsServerCommand) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDnsServerCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateDnsServerCommand) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *CreateDnsServerCommand) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *CreateDnsServerCommand) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *CreateDnsServerCommand) UnsetAddress() {
	o.Address.Unset()
}

// GetAccessProfileId returns the AccessProfileId field value if set, zero value otherwise.
func (o *CreateDnsServerCommand) GetAccessProfileId() int32 {
	if o == nil || IsNil(o.AccessProfileId) {
		var ret int32
		return ret
	}
	return *o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDnsServerCommand) GetAccessProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessProfileId) {
		return nil, false
	}
	return o.AccessProfileId, true
}

// HasAccessProfileId returns a boolean if a field has been set.
func (o *CreateDnsServerCommand) HasAccessProfileId() bool {
	if o != nil && !IsNil(o.AccessProfileId) {
		return true
	}

	return false
}

// SetAccessProfileId gets a reference to the given int32 and assigns it to the AccessProfileId field.
func (o *CreateDnsServerCommand) SetAccessProfileId(v int32) {
	o.AccessProfileId = &v
}

func (o CreateDnsServerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDnsServerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.AccessProfileId) {
		toSerialize["accessProfileId"] = o.AccessProfileId
	}
	return toSerialize, nil
}

type NullableCreateDnsServerCommand struct {
	value *CreateDnsServerCommand
	isSet bool
}

func (v NullableCreateDnsServerCommand) Get() *CreateDnsServerCommand {
	return v.value
}

func (v *NullableCreateDnsServerCommand) Set(val *CreateDnsServerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDnsServerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDnsServerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDnsServerCommand(val *CreateDnsServerCommand) *NullableCreateDnsServerCommand {
	return &NullableCreateDnsServerCommand{value: val, isSet: true}
}

func (v NullableCreateDnsServerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDnsServerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
