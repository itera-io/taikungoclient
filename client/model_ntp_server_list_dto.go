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
	"bytes"
	"fmt"
)

// checks if the NtpServerListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtpServerListDto{}

// NtpServerListDto struct for NtpServerListDto
type NtpServerListDto struct {
	Id int32 `json:"id"`
	Address NullableString `json:"address"`
}

type _NtpServerListDto NtpServerListDto

// NewNtpServerListDto instantiates a new NtpServerListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpServerListDto(id int32, address NullableString) *NtpServerListDto {
	this := NtpServerListDto{}
	this.Id = id
	this.Address = address
	return &this
}

// NewNtpServerListDtoWithDefaults instantiates a new NtpServerListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpServerListDtoWithDefaults() *NtpServerListDto {
	this := NtpServerListDto{}
	return &this
}

// GetId returns the Id field value
func (o *NtpServerListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NtpServerListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NtpServerListDto) SetId(v int32) {
	o.Id = v
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NtpServerListDto) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpServerListDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *NtpServerListDto) SetAddress(v string) {
	o.Address.Set(&v)
}

func (o NtpServerListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtpServerListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["address"] = o.Address.Get()
	return toSerialize, nil
}

func (o *NtpServerListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNtpServerListDto := _NtpServerListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNtpServerListDto)

	if err != nil {
		return err
	}

	*o = NtpServerListDto(varNtpServerListDto)

	return err
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


