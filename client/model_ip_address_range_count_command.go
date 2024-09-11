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

// checks if the IpAddressRangeCountCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAddressRangeCountCommand{}

// IpAddressRangeCountCommand struct for IpAddressRangeCountCommand
type IpAddressRangeCountCommand struct {
	Begin *string `json:"begin,omitempty"`
	End *string `json:"end,omitempty"`
}

// NewIpAddressRangeCountCommand instantiates a new IpAddressRangeCountCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddressRangeCountCommand() *IpAddressRangeCountCommand {
	this := IpAddressRangeCountCommand{}
	return &this
}

// NewIpAddressRangeCountCommandWithDefaults instantiates a new IpAddressRangeCountCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddressRangeCountCommandWithDefaults() *IpAddressRangeCountCommand {
	this := IpAddressRangeCountCommand{}
	return &this
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *IpAddressRangeCountCommand) GetBegin() string {
	if o == nil || IsNil(o.Begin) {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressRangeCountCommand) GetBeginOk() (*string, bool) {
	if o == nil || IsNil(o.Begin) {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *IpAddressRangeCountCommand) HasBegin() bool {
	if o != nil && !IsNil(o.Begin) {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *IpAddressRangeCountCommand) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *IpAddressRangeCountCommand) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddressRangeCountCommand) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *IpAddressRangeCountCommand) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *IpAddressRangeCountCommand) SetEnd(v string) {
	o.End = &v
}

func (o IpAddressRangeCountCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAddressRangeCountCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Begin) {
		toSerialize["begin"] = o.Begin
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableIpAddressRangeCountCommand struct {
	value *IpAddressRangeCountCommand
	isSet bool
}

func (v NullableIpAddressRangeCountCommand) Get() *IpAddressRangeCountCommand {
	return v.value
}

func (v *NullableIpAddressRangeCountCommand) Set(val *IpAddressRangeCountCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddressRangeCountCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddressRangeCountCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddressRangeCountCommand(val *IpAddressRangeCountCommand) *NullableIpAddressRangeCountCommand {
	return &NullableIpAddressRangeCountCommand{value: val, isSet: true}
}

func (v NullableIpAddressRangeCountCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddressRangeCountCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


