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

// checks if the ProjectAppParamDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAppParamDto{}

// ProjectAppParamDto struct for ProjectAppParamDto
type ProjectAppParamDto struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	IsEditableWhenInstalling *bool `json:"isEditableWhenInstalling,omitempty"`
	IsEditableAfterInstallation *bool `json:"isEditableAfterInstallation,omitempty"`
	IsMandatory *bool `json:"isMandatory,omitempty"`
}

// NewProjectAppParamDto instantiates a new ProjectAppParamDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAppParamDto() *ProjectAppParamDto {
	this := ProjectAppParamDto{}
	return &this
}

// NewProjectAppParamDtoWithDefaults instantiates a new ProjectAppParamDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAppParamDtoWithDefaults() *ProjectAppParamDto {
	this := ProjectAppParamDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProjectAppParamDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectAppParamDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProjectAppParamDto) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProjectAppParamDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectAppParamDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProjectAppParamDto) SetValue(v string) {
	o.Value = &v
}

// GetIsEditableWhenInstalling returns the IsEditableWhenInstalling field value if set, zero value otherwise.
func (o *ProjectAppParamDto) GetIsEditableWhenInstalling() bool {
	if o == nil || IsNil(o.IsEditableWhenInstalling) {
		var ret bool
		return ret
	}
	return *o.IsEditableWhenInstalling
}

// GetIsEditableWhenInstallingOk returns a tuple with the IsEditableWhenInstalling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamDto) GetIsEditableWhenInstallingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEditableWhenInstalling) {
		return nil, false
	}
	return o.IsEditableWhenInstalling, true
}

// HasIsEditableWhenInstalling returns a boolean if a field has been set.
func (o *ProjectAppParamDto) HasIsEditableWhenInstalling() bool {
	if o != nil && !IsNil(o.IsEditableWhenInstalling) {
		return true
	}

	return false
}

// SetIsEditableWhenInstalling gets a reference to the given bool and assigns it to the IsEditableWhenInstalling field.
func (o *ProjectAppParamDto) SetIsEditableWhenInstalling(v bool) {
	o.IsEditableWhenInstalling = &v
}

// GetIsEditableAfterInstallation returns the IsEditableAfterInstallation field value if set, zero value otherwise.
func (o *ProjectAppParamDto) GetIsEditableAfterInstallation() bool {
	if o == nil || IsNil(o.IsEditableAfterInstallation) {
		var ret bool
		return ret
	}
	return *o.IsEditableAfterInstallation
}

// GetIsEditableAfterInstallationOk returns a tuple with the IsEditableAfterInstallation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamDto) GetIsEditableAfterInstallationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEditableAfterInstallation) {
		return nil, false
	}
	return o.IsEditableAfterInstallation, true
}

// HasIsEditableAfterInstallation returns a boolean if a field has been set.
func (o *ProjectAppParamDto) HasIsEditableAfterInstallation() bool {
	if o != nil && !IsNil(o.IsEditableAfterInstallation) {
		return true
	}

	return false
}

// SetIsEditableAfterInstallation gets a reference to the given bool and assigns it to the IsEditableAfterInstallation field.
func (o *ProjectAppParamDto) SetIsEditableAfterInstallation(v bool) {
	o.IsEditableAfterInstallation = &v
}

// GetIsMandatory returns the IsMandatory field value if set, zero value otherwise.
func (o *ProjectAppParamDto) GetIsMandatory() bool {
	if o == nil || IsNil(o.IsMandatory) {
		var ret bool
		return ret
	}
	return *o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppParamDto) GetIsMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMandatory) {
		return nil, false
	}
	return o.IsMandatory, true
}

// HasIsMandatory returns a boolean if a field has been set.
func (o *ProjectAppParamDto) HasIsMandatory() bool {
	if o != nil && !IsNil(o.IsMandatory) {
		return true
	}

	return false
}

// SetIsMandatory gets a reference to the given bool and assigns it to the IsMandatory field.
func (o *ProjectAppParamDto) SetIsMandatory(v bool) {
	o.IsMandatory = &v
}

func (o ProjectAppParamDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAppParamDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsEditableWhenInstalling) {
		toSerialize["isEditableWhenInstalling"] = o.IsEditableWhenInstalling
	}
	if !IsNil(o.IsEditableAfterInstallation) {
		toSerialize["isEditableAfterInstallation"] = o.IsEditableAfterInstallation
	}
	if !IsNil(o.IsMandatory) {
		toSerialize["isMandatory"] = o.IsMandatory
	}
	return toSerialize, nil
}

type NullableProjectAppParamDto struct {
	value *ProjectAppParamDto
	isSet bool
}

func (v NullableProjectAppParamDto) Get() *ProjectAppParamDto {
	return v.value
}

func (v *NullableProjectAppParamDto) Set(val *ProjectAppParamDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAppParamDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAppParamDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAppParamDto(val *ProjectAppParamDto) *NullableProjectAppParamDto {
	return &NullableProjectAppParamDto{value: val, isSet: true}
}

func (v NullableProjectAppParamDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAppParamDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


