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

// checks if the StandAloneVmDiskForDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneVmDiskForDetailsDto{}

// StandAloneVmDiskForDetailsDto struct for StandAloneVmDiskForDetailsDto
type StandAloneVmDiskForDetailsDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CurrentSize *int64 `json:"currentSize,omitempty"`
	TargetSize *int64 `json:"targetSize,omitempty"`
	VolumeType *string `json:"volumeType,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	LunId *string `json:"lunId,omitempty"`
	Status *StandAloneVmDiskStatus `json:"status,omitempty"`
}

// NewStandAloneVmDiskForDetailsDto instantiates a new StandAloneVmDiskForDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneVmDiskForDetailsDto() *StandAloneVmDiskForDetailsDto {
	this := StandAloneVmDiskForDetailsDto{}
	return &this
}

// NewStandAloneVmDiskForDetailsDtoWithDefaults instantiates a new StandAloneVmDiskForDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneVmDiskForDetailsDtoWithDefaults() *StandAloneVmDiskForDetailsDto {
	this := StandAloneVmDiskForDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneVmDiskForDetailsDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StandAloneVmDiskForDetailsDto) SetName(v string) {
	o.Name = &v
}

// GetCurrentSize returns the CurrentSize field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetCurrentSize() int64 {
	if o == nil || IsNil(o.CurrentSize) {
		var ret int64
		return ret
	}
	return *o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetCurrentSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentSize) {
		return nil, false
	}
	return o.CurrentSize, true
}

// HasCurrentSize returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasCurrentSize() bool {
	if o != nil && !IsNil(o.CurrentSize) {
		return true
	}

	return false
}

// SetCurrentSize gets a reference to the given int64 and assigns it to the CurrentSize field.
func (o *StandAloneVmDiskForDetailsDto) SetCurrentSize(v int64) {
	o.CurrentSize = &v
}

// GetTargetSize returns the TargetSize field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetTargetSize() int64 {
	if o == nil || IsNil(o.TargetSize) {
		var ret int64
		return ret
	}
	return *o.TargetSize
}

// GetTargetSizeOk returns a tuple with the TargetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetTargetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetSize) {
		return nil, false
	}
	return o.TargetSize, true
}

// HasTargetSize returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasTargetSize() bool {
	if o != nil && !IsNil(o.TargetSize) {
		return true
	}

	return false
}

// SetTargetSize gets a reference to the given int64 and assigns it to the TargetSize field.
func (o *StandAloneVmDiskForDetailsDto) SetTargetSize(v int64) {
	o.TargetSize = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *StandAloneVmDiskForDetailsDto) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *StandAloneVmDiskForDetailsDto) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetLunId() string {
	if o == nil || IsNil(o.LunId) {
		var ret string
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetLunIdOk() (*string, bool) {
	if o == nil || IsNil(o.LunId) {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasLunId() bool {
	if o != nil && !IsNil(o.LunId) {
		return true
	}

	return false
}

// SetLunId gets a reference to the given string and assigns it to the LunId field.
func (o *StandAloneVmDiskForDetailsDto) SetLunId(v string) {
	o.LunId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StandAloneVmDiskForDetailsDto) GetStatus() StandAloneVmDiskStatus {
	if o == nil || IsNil(o.Status) {
		var ret StandAloneVmDiskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneVmDiskForDetailsDto) GetStatusOk() (*StandAloneVmDiskStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StandAloneVmDiskForDetailsDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StandAloneVmDiskStatus and assigns it to the Status field.
func (o *StandAloneVmDiskForDetailsDto) SetStatus(v StandAloneVmDiskStatus) {
	o.Status = &v
}

func (o StandAloneVmDiskForDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneVmDiskForDetailsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CurrentSize) {
		toSerialize["currentSize"] = o.CurrentSize
	}
	if !IsNil(o.TargetSize) {
		toSerialize["targetSize"] = o.TargetSize
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.LunId) {
		toSerialize["lunId"] = o.LunId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableStandAloneVmDiskForDetailsDto struct {
	value *StandAloneVmDiskForDetailsDto
	isSet bool
}

func (v NullableStandAloneVmDiskForDetailsDto) Get() *StandAloneVmDiskForDetailsDto {
	return v.value
}

func (v *NullableStandAloneVmDiskForDetailsDto) Set(val *StandAloneVmDiskForDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneVmDiskForDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneVmDiskForDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneVmDiskForDetailsDto(val *StandAloneVmDiskForDetailsDto) *NullableStandAloneVmDiskForDetailsDto {
	return &NullableStandAloneVmDiskForDetailsDto{value: val, isSet: true}
}

func (v NullableStandAloneVmDiskForDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneVmDiskForDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


