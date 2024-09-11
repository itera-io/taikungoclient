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

// checks if the VsphereFlavorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsphereFlavorData{}

// VsphereFlavorData struct for VsphereFlavorData
type VsphereFlavorData struct {
	Name NullableString `json:"name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *float64 `json:"ram,omitempty"`
}

// NewVsphereFlavorData instantiates a new VsphereFlavorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereFlavorData() *VsphereFlavorData {
	this := VsphereFlavorData{}
	return &this
}

// NewVsphereFlavorDataWithDefaults instantiates a new VsphereFlavorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereFlavorDataWithDefaults() *VsphereFlavorData {
	this := VsphereFlavorData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereFlavorData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereFlavorData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VsphereFlavorData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VsphereFlavorData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VsphereFlavorData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VsphereFlavorData) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VsphereFlavorData) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereFlavorData) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VsphereFlavorData) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *VsphereFlavorData) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *VsphereFlavorData) GetRam() float64 {
	if o == nil || IsNil(o.Ram) {
		var ret float64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereFlavorData) GetRamOk() (*float64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *VsphereFlavorData) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given float64 and assigns it to the Ram field.
func (o *VsphereFlavorData) SetRam(v float64) {
	o.Ram = &v
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	return toSerialize, nil
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


