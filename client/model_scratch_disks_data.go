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

// checks if the ScratchDisksData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScratchDisksData{}

// ScratchDisksData struct for ScratchDisksData
type ScratchDisksData struct {
	DiskGb NullableInt32 `json:"diskGb,omitempty"`
}

// NewScratchDisksData instantiates a new ScratchDisksData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScratchDisksData() *ScratchDisksData {
	this := ScratchDisksData{}
	return &this
}

// NewScratchDisksDataWithDefaults instantiates a new ScratchDisksData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScratchDisksDataWithDefaults() *ScratchDisksData {
	this := ScratchDisksData{}
	return &this
}

// GetDiskGb returns the DiskGb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScratchDisksData) GetDiskGb() int32 {
	if o == nil || IsNil(o.DiskGb.Get()) {
		var ret int32
		return ret
	}
	return *o.DiskGb.Get()
}

// GetDiskGbOk returns a tuple with the DiskGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScratchDisksData) GetDiskGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskGb.Get(), o.DiskGb.IsSet()
}

// HasDiskGb returns a boolean if a field has been set.
func (o *ScratchDisksData) HasDiskGb() bool {
	if o != nil && o.DiskGb.IsSet() {
		return true
	}

	return false
}

// SetDiskGb gets a reference to the given NullableInt32 and assigns it to the DiskGb field.
func (o *ScratchDisksData) SetDiskGb(v int32) {
	o.DiskGb.Set(&v)
}
// SetDiskGbNil sets the value for DiskGb to be an explicit nil
func (o *ScratchDisksData) SetDiskGbNil() {
	o.DiskGb.Set(nil)
}

// UnsetDiskGb ensures that no value is present for DiskGb, not even an explicit nil
func (o *ScratchDisksData) UnsetDiskGb() {
	o.DiskGb.Unset()
}

func (o ScratchDisksData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScratchDisksData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskGb.IsSet() {
		toSerialize["diskGb"] = o.DiskGb.Get()
	}
	return toSerialize, nil
}

type NullableScratchDisksData struct {
	value *ScratchDisksData
	isSet bool
}

func (v NullableScratchDisksData) Get() *ScratchDisksData {
	return v.value
}

func (v *NullableScratchDisksData) Set(val *ScratchDisksData) {
	v.value = val
	v.isSet = true
}

func (v NullableScratchDisksData) IsSet() bool {
	return v.isSet
}

func (v *NullableScratchDisksData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScratchDisksData(val *ScratchDisksData) *NullableScratchDisksData {
	return &NullableScratchDisksData{value: val, isSet: true}
}

func (v NullableScratchDisksData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScratchDisksData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


