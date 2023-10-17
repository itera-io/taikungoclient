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

// checks if the NtpServerListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtpServerListDto{}

// NtpServerListDto struct for NtpServerListDto
type NtpServerListDto struct {
	Id      *int32         `json:"id,omitempty"`
	Address NullableString `json:"address,omitempty"`
}

// NewNtpServerListDto instantiates a new NtpServerListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpServerListDto() *NtpServerListDto {
	this := NtpServerListDto{}
	return &this
}

// NewNtpServerListDtoWithDefaults instantiates a new NtpServerListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpServerListDtoWithDefaults() *NtpServerListDto {
	this := NtpServerListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NtpServerListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpServerListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NtpServerListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NtpServerListDto) SetId(v int32) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NtpServerListDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpServerListDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *NtpServerListDto) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *NtpServerListDto) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *NtpServerListDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *NtpServerListDto) UnsetAddress() {
	o.Address.Unset()
}

func (o NtpServerListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtpServerListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return toSerialize, nil
}

type NullableNtpServerListDto struct {
	value *NtpServerListDto
	isSet bool
}

func (v NullableNtpServerListDto) Get() *NtpServerListDto {
	return v.value
}

func (v *NullableNtpServerListDto) Set(val *NtpServerListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpServerListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpServerListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpServerListDto(val *NtpServerListDto) *NullableNtpServerListDto {
	return &NullableNtpServerListDto{value: val, isSet: true}
}

func (v NullableNtpServerListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpServerListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
