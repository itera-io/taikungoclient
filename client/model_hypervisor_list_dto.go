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

// checks if the HypervisorListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorListDto{}

// HypervisorListDto struct for HypervisorListDto
type HypervisorListDto struct {
	Host NullableString `json:"host,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewHypervisorListDto instantiates a new HypervisorListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorListDto() *HypervisorListDto {
	this := HypervisorListDto{}
	return &this
}

// NewHypervisorListDtoWithDefaults instantiates a new HypervisorListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorListDtoWithDefaults() *HypervisorListDto {
	this := HypervisorListDto{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorListDto) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorListDto) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *HypervisorListDto) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *HypervisorListDto) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *HypervisorListDto) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *HypervisorListDto) UnsetHost() {
	o.Host.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HypervisorListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HypervisorListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HypervisorListDto) UnsetName() {
	o.Name.Unset()
}

func (o HypervisorListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableHypervisorListDto struct {
	value *HypervisorListDto
	isSet bool
}

func (v NullableHypervisorListDto) Get() *HypervisorListDto {
	return v.value
}

func (v *NullableHypervisorListDto) Set(val *HypervisorListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorListDto(val *HypervisorListDto) *NullableHypervisorListDto {
	return &NullableHypervisorListDto{value: val, isSet: true}
}

func (v NullableHypervisorListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


