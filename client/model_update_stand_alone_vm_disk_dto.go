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

// checks if the UpdateStandAloneVmDiskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStandAloneVmDiskDto{}

// UpdateStandAloneVmDiskDto struct for UpdateStandAloneVmDiskDto
type UpdateStandAloneVmDiskDto struct {
	Name NullableString `json:"name,omitempty"`
	DeviceName NullableString `json:"deviceName,omitempty"`
	Id NullableString `json:"id,omitempty"`
}

// NewUpdateStandAloneVmDiskDto instantiates a new UpdateStandAloneVmDiskDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStandAloneVmDiskDto() *UpdateStandAloneVmDiskDto {
	this := UpdateStandAloneVmDiskDto{}
	return &this
}

// NewUpdateStandAloneVmDiskDtoWithDefaults instantiates a new UpdateStandAloneVmDiskDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStandAloneVmDiskDtoWithDefaults() *UpdateStandAloneVmDiskDto {
	this := UpdateStandAloneVmDiskDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmDiskDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmDiskDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateStandAloneVmDiskDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateStandAloneVmDiskDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateStandAloneVmDiskDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateStandAloneVmDiskDto) UnsetName() {
	o.Name.Unset()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmDiskDto) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmDiskDto) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *UpdateStandAloneVmDiskDto) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *UpdateStandAloneVmDiskDto) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}
// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *UpdateStandAloneVmDiskDto) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *UpdateStandAloneVmDiskDto) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmDiskDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmDiskDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpdateStandAloneVmDiskDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *UpdateStandAloneVmDiskDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UpdateStandAloneVmDiskDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpdateStandAloneVmDiskDto) UnsetId() {
	o.Id.Unset()
}

func (o UpdateStandAloneVmDiskDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStandAloneVmDiskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["deviceName"] = o.DeviceName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableUpdateStandAloneVmDiskDto struct {
	value *UpdateStandAloneVmDiskDto
	isSet bool
}

func (v NullableUpdateStandAloneVmDiskDto) Get() *UpdateStandAloneVmDiskDto {
	return v.value
}

func (v *NullableUpdateStandAloneVmDiskDto) Set(val *UpdateStandAloneVmDiskDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStandAloneVmDiskDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStandAloneVmDiskDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStandAloneVmDiskDto(val *UpdateStandAloneVmDiskDto) *NullableUpdateStandAloneVmDiskDto {
	return &NullableUpdateStandAloneVmDiskDto{value: val, isSet: true}
}

func (v NullableUpdateStandAloneVmDiskDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStandAloneVmDiskDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


