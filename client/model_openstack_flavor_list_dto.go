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

// checks if the OpenstackFlavorListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackFlavorListDto{}

// OpenstackFlavorListDto struct for OpenstackFlavorListDto
type OpenstackFlavorListDto struct {
	Ram         *int64         `json:"ram,omitempty"`
	Cpu         *int64         `json:"cpu,omitempty"`
	Name        NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewOpenstackFlavorListDto instantiates a new OpenstackFlavorListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackFlavorListDto() *OpenstackFlavorListDto {
	this := OpenstackFlavorListDto{}
	return &this
}

// NewOpenstackFlavorListDtoWithDefaults instantiates a new OpenstackFlavorListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackFlavorListDtoWithDefaults() *OpenstackFlavorListDto {
	this := OpenstackFlavorListDto{}
	return &this
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *OpenstackFlavorListDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackFlavorListDto) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *OpenstackFlavorListDto) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *OpenstackFlavorListDto) SetRam(v int64) {
	o.Ram = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *OpenstackFlavorListDto) GetCpu() int64 {
	if o == nil || IsNil(o.Cpu) {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackFlavorListDto) GetCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *OpenstackFlavorListDto) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *OpenstackFlavorListDto) SetCpu(v int64) {
	o.Cpu = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackFlavorListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackFlavorListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpenstackFlavorListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpenstackFlavorListDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *OpenstackFlavorListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpenstackFlavorListDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackFlavorListDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackFlavorListDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenstackFlavorListDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OpenstackFlavorListDto) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OpenstackFlavorListDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OpenstackFlavorListDto) UnsetDescription() {
	o.Description.Unset()
}

func (o OpenstackFlavorListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackFlavorListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableOpenstackFlavorListDto struct {
	value *OpenstackFlavorListDto
	isSet bool
}

func (v NullableOpenstackFlavorListDto) Get() *OpenstackFlavorListDto {
	return v.value
}

func (v *NullableOpenstackFlavorListDto) Set(val *OpenstackFlavorListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackFlavorListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackFlavorListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackFlavorListDto(val *OpenstackFlavorListDto) *NullableOpenstackFlavorListDto {
	return &NullableOpenstackFlavorListDto{value: val, isSet: true}
}

func (v NullableOpenstackFlavorListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackFlavorListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
