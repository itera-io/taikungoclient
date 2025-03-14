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

// checks if the VsphereFlavorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsphereFlavorData{}

// VsphereFlavorData struct for VsphereFlavorData
type VsphereFlavorData struct {
	Name NullableString `json:"name"`
	Cpu int32 `json:"cpu"`
	Ram float64 `json:"ram"`
}

type _VsphereFlavorData VsphereFlavorData

// NewVsphereFlavorData instantiates a new VsphereFlavorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereFlavorData(name NullableString, cpu int32, ram float64) *VsphereFlavorData {
	this := VsphereFlavorData{}
	this.Name = name
	this.Cpu = cpu
	this.Ram = ram
	return &this
}

// NewVsphereFlavorDataWithDefaults instantiates a new VsphereFlavorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereFlavorDataWithDefaults() *VsphereFlavorData {
	this := VsphereFlavorData{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VsphereFlavorData) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereFlavorData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *VsphereFlavorData) SetName(v string) {
	o.Name.Set(&v)
}

// GetCpu returns the Cpu field value
func (o *VsphereFlavorData) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *VsphereFlavorData) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *VsphereFlavorData) SetCpu(v int32) {
	o.Cpu = v
}

// GetRam returns the Ram field value
func (o *VsphereFlavorData) GetRam() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *VsphereFlavorData) GetRamOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *VsphereFlavorData) SetRam(v float64) {
	o.Ram = v
}

func (o VsphereFlavorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsphereFlavorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram"] = o.Ram
	return toSerialize, nil
}

func (o *VsphereFlavorData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varVsphereFlavorData := _VsphereFlavorData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVsphereFlavorData)

	if err != nil {
		return err
	}

	*o = VsphereFlavorData(varVsphereFlavorData)

	return err
}

type NullableVsphereFlavorData struct {
	value *VsphereFlavorData
	isSet bool
}

func (v NullableVsphereFlavorData) Get() *VsphereFlavorData {
	return v.value
}

func (v *NullableVsphereFlavorData) Set(val *VsphereFlavorData) {
	v.value = val
	v.isSet = true
}

func (v NullableVsphereFlavorData) IsSet() bool {
	return v.isSet
}

func (v *NullableVsphereFlavorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsphereFlavorData(val *VsphereFlavorData) *NullableVsphereFlavorData {
	return &NullableVsphereFlavorData{value: val, isSet: true}
}

func (v NullableVsphereFlavorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsphereFlavorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


