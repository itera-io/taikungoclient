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

// checks if the ServersForBillingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServersForBillingDto{}

// ServersForBillingDto struct for ServersForBillingDto
type ServersForBillingDto struct {
	Cpu int32 `json:"cpu"`
	Ram int64 `json:"ram"`
}

type _ServersForBillingDto ServersForBillingDto

// NewServersForBillingDto instantiates a new ServersForBillingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServersForBillingDto(cpu int32, ram int64) *ServersForBillingDto {
	this := ServersForBillingDto{}
	this.Cpu = cpu
	this.Ram = ram
	return &this
}

// NewServersForBillingDtoWithDefaults instantiates a new ServersForBillingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServersForBillingDtoWithDefaults() *ServersForBillingDto {
	this := ServersForBillingDto{}
	return &this
}

// GetCpu returns the Cpu field value
func (o *ServersForBillingDto) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ServersForBillingDto) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *ServersForBillingDto) SetCpu(v int32) {
	o.Cpu = v
}

// GetRam returns the Ram field value
func (o *ServersForBillingDto) GetRam() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *ServersForBillingDto) GetRamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *ServersForBillingDto) SetRam(v int64) {
	o.Ram = v
}

func (o ServersForBillingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServersForBillingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram"] = o.Ram
	return toSerialize, nil
}

func (o *ServersForBillingDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cpu",
		"ram",
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

	varServersForBillingDto := _ServersForBillingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServersForBillingDto)

	if err != nil {
		return err
	}

	*o = ServersForBillingDto(varServersForBillingDto)

	return err
}

type NullableServersForBillingDto struct {
	value *ServersForBillingDto
	isSet bool
}

func (v NullableServersForBillingDto) Get() *ServersForBillingDto {
	return v.value
}

func (v *NullableServersForBillingDto) Set(val *ServersForBillingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServersForBillingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServersForBillingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServersForBillingDto(val *ServersForBillingDto) *NullableServersForBillingDto {
	return &NullableServersForBillingDto{value: val, isSet: true}
}

func (v NullableServersForBillingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServersForBillingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


