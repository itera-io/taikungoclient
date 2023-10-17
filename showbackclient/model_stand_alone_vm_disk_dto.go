/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the StandAloneVmDiskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneVmDiskDto{}

// StandAloneVmDiskDto struct for StandAloneVmDiskDto
type StandAloneVmDiskDto struct {
	Name       NullableString `json:"name,omitempty"`
	Size       *int64         `json:"size,omitempty"`
	VolumeType NullableString `json:"volumeType,omitempty"`
	DeviceName NullableString `json:"deviceName,omitempty"`
	LunId      NullableInt32  `json:"lunId,omitempty"`
}

// NewStandAloneVmDiskDto instantiates a new StandAloneVmDiskDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneVmDiskDto() *StandAloneVmDiskDto {
	this := StandAloneVmDiskDto{}
	return &this
}

// NewStandAloneVmDiskDtoWithDefaults instantiates a new StandAloneVmDiskDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneVmDiskDtoWithDefaults() *StandAloneVmDiskDto {
	this := StandAloneVmDiskDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneVmDiskDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneVmDiskDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandAloneVmDiskDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *StandAloneVmDiskDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandAloneVmDiskDto) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StandAloneVmDiskDto) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskDto) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StandAloneVmDiskDto) SetSize(v int64) {
	o.Size = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneVmDiskDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType.Get()) {
		var ret string
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneVmDiskDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullableString and assigns it to the VolumeType field.
func (o *StandAloneVmDiskDto) SetVolumeType(v string) {
	o.VolumeType.Set(&v)
}

// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *StandAloneVmDiskDto) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *StandAloneVmDiskDto) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneVmDiskDto) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneVmDiskDto) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *StandAloneVmDiskDto) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}

// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *StandAloneVmDiskDto) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *StandAloneVmDiskDto) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetLunId returns the LunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneVmDiskDto) GetLunId() int32 {
	if o == nil || IsNil(o.LunId.Get()) {
		var ret int32
		return ret
	}
	return *o.LunId.Get()
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneVmDiskDto) GetLunIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LunId.Get(), o.LunId.IsSet()
}

// HasLunId returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasLunId() bool {
	if o != nil && o.LunId.IsSet() {
		return true
	}

	return false
}

// SetLunId gets a reference to the given NullableInt32 and assigns it to the LunId field.
func (o *StandAloneVmDiskDto) SetLunId(v int32) {
	o.LunId.Set(&v)
}

// SetLunIdNil sets the value for LunId to be an explicit nil
func (o *StandAloneVmDiskDto) SetLunIdNil() {
	o.LunId.Set(nil)
}

// UnsetLunId ensures that no value is present for LunId, not even an explicit nil
func (o *StandAloneVmDiskDto) UnsetLunId() {
	o.LunId.Unset()
}

func (o StandAloneVmDiskDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneVmDiskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if o.VolumeType.IsSet() {
		toSerialize["volumeType"] = o.VolumeType.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["deviceName"] = o.DeviceName.Get()
	}
	if o.LunId.IsSet() {
		toSerialize["lunId"] = o.LunId.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneVmDiskDto struct {
	value *StandAloneVmDiskDto
	isSet bool
}

func (v NullableStandAloneVmDiskDto) Get() *StandAloneVmDiskDto {
	return v.value
}

func (v *NullableStandAloneVmDiskDto) Set(val *StandAloneVmDiskDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneVmDiskDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneVmDiskDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneVmDiskDto(val *StandAloneVmDiskDto) *NullableStandAloneVmDiskDto {
	return &NullableStandAloneVmDiskDto{value: val, isSet: true}
}

func (v NullableStandAloneVmDiskDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneVmDiskDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
