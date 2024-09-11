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

// checks if the StandAloneVmDiskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneVmDiskDto{}

// StandAloneVmDiskDto struct for StandAloneVmDiskDto
type StandAloneVmDiskDto struct {
	Name *string `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`
	VolumeType *string `json:"volumeType,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	LunId NullableInt32 `json:"lunId,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *StandAloneVmDiskDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StandAloneVmDiskDto) SetName(v string) {
	o.Name = &v
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

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *StandAloneVmDiskDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *StandAloneVmDiskDto) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *StandAloneVmDiskDto) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskDto) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *StandAloneVmDiskDto) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *StandAloneVmDiskDto) SetDeviceName(v string) {
	o.DeviceName = &v
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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneVmDiskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
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


