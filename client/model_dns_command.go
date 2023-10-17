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

// checks if the DnsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsCommand{}

// DnsCommand struct for DnsCommand
type DnsCommand struct {
	Address NullableString `json:"address,omitempty"`
}

// NewDnsCommand instantiates a new DnsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsCommand() *DnsCommand {
	this := DnsCommand{}
	return &this
}

// NewDnsCommandWithDefaults instantiates a new DnsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsCommandWithDefaults() *DnsCommand {
	this := DnsCommand{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsCommand) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *DnsCommand) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *DnsCommand) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *DnsCommand) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *DnsCommand) UnsetAddress() {
	o.Address.Unset()
}

func (o DnsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return toSerialize, nil
}

type NullableDnsCommand struct {
	value *DnsCommand
	isSet bool
}

func (v NullableDnsCommand) Get() *DnsCommand {
	return v.value
}

func (v *NullableDnsCommand) Set(val *DnsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsCommand(val *DnsCommand) *NullableDnsCommand {
	return &NullableDnsCommand{value: val, isSet: true}
}

func (v NullableDnsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
