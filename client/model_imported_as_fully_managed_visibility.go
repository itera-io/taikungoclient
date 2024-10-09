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

// checks if the ImportedAsFullyManagedVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedAsFullyManagedVisibility{}

// ImportedAsFullyManagedVisibility struct for ImportedAsFullyManagedVisibility
type ImportedAsFullyManagedVisibility struct {
	Lock ButtonStatusDto `json:"lock"`
	Unlock ButtonStatusDto `json:"unlock"`
	AddVCluster ButtonStatusDto `json:"addVCluster"`
	AttachAlertingProfile ButtonStatusDto `json:"attachAlertingProfile"`
	DetachAlertingProfile ButtonStatusDto `json:"detachAlertingProfile"`
}

type _ImportedAsFullyManagedVisibility ImportedAsFullyManagedVisibility

// NewImportedAsFullyManagedVisibility instantiates a new ImportedAsFullyManagedVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedAsFullyManagedVisibility(lock ButtonStatusDto, unlock ButtonStatusDto, addVCluster ButtonStatusDto, attachAlertingProfile ButtonStatusDto, detachAlertingProfile ButtonStatusDto) *ImportedAsFullyManagedVisibility {
	this := ImportedAsFullyManagedVisibility{}
	this.Lock = lock
	this.Unlock = unlock
	this.AddVCluster = addVCluster
	this.AttachAlertingProfile = attachAlertingProfile
	this.DetachAlertingProfile = detachAlertingProfile
	return &this
}

// NewImportedAsFullyManagedVisibilityWithDefaults instantiates a new ImportedAsFullyManagedVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedAsFullyManagedVisibilityWithDefaults() *ImportedAsFullyManagedVisibility {
	this := ImportedAsFullyManagedVisibility{}
	return &this
}

// GetLock returns the Lock field value
func (o *ImportedAsFullyManagedVisibility) GetLock() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.Lock
}

// GetLockOk returns a tuple with the Lock field value
// and a boolean to check if the value has been set.
func (o *ImportedAsFullyManagedVisibility) GetLockOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lock, true
}

// SetLock sets field value
func (o *ImportedAsFullyManagedVisibility) SetLock(v ButtonStatusDto) {
	o.Lock = v
}

// GetUnlock returns the Unlock field value
func (o *ImportedAsFullyManagedVisibility) GetUnlock() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.Unlock
}

// GetUnlockOk returns a tuple with the Unlock field value
// and a boolean to check if the value has been set.
func (o *ImportedAsFullyManagedVisibility) GetUnlockOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unlock, true
}

// SetUnlock sets field value
func (o *ImportedAsFullyManagedVisibility) SetUnlock(v ButtonStatusDto) {
	o.Unlock = v
}

// GetAddVCluster returns the AddVCluster field value
func (o *ImportedAsFullyManagedVisibility) GetAddVCluster() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.AddVCluster
}

// GetAddVClusterOk returns a tuple with the AddVCluster field value
// and a boolean to check if the value has been set.
func (o *ImportedAsFullyManagedVisibility) GetAddVClusterOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddVCluster, true
}

// SetAddVCluster sets field value
func (o *ImportedAsFullyManagedVisibility) SetAddVCluster(v ButtonStatusDto) {
	o.AddVCluster = v
}

// GetAttachAlertingProfile returns the AttachAlertingProfile field value
func (o *ImportedAsFullyManagedVisibility) GetAttachAlertingProfile() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.AttachAlertingProfile
}

// GetAttachAlertingProfileOk returns a tuple with the AttachAlertingProfile field value
// and a boolean to check if the value has been set.
func (o *ImportedAsFullyManagedVisibility) GetAttachAlertingProfileOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachAlertingProfile, true
}

// SetAttachAlertingProfile sets field value
func (o *ImportedAsFullyManagedVisibility) SetAttachAlertingProfile(v ButtonStatusDto) {
	o.AttachAlertingProfile = v
}

// GetDetachAlertingProfile returns the DetachAlertingProfile field value
func (o *ImportedAsFullyManagedVisibility) GetDetachAlertingProfile() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.DetachAlertingProfile
}

// GetDetachAlertingProfileOk returns a tuple with the DetachAlertingProfile field value
// and a boolean to check if the value has been set.
func (o *ImportedAsFullyManagedVisibility) GetDetachAlertingProfileOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetachAlertingProfile, true
}

// SetDetachAlertingProfile sets field value
func (o *ImportedAsFullyManagedVisibility) SetDetachAlertingProfile(v ButtonStatusDto) {
	o.DetachAlertingProfile = v
}

func (o ImportedAsFullyManagedVisibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedAsFullyManagedVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lock"] = o.Lock
	toSerialize["unlock"] = o.Unlock
	toSerialize["addVCluster"] = o.AddVCluster
	toSerialize["attachAlertingProfile"] = o.AttachAlertingProfile
	toSerialize["detachAlertingProfile"] = o.DetachAlertingProfile
	return toSerialize, nil
}

func (o *ImportedAsFullyManagedVisibility) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lock",
		"unlock",
		"addVCluster",
		"attachAlertingProfile",
		"detachAlertingProfile",
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

	varImportedAsFullyManagedVisibility := _ImportedAsFullyManagedVisibility{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportedAsFullyManagedVisibility)

	if err != nil {
		return err
	}

	*o = ImportedAsFullyManagedVisibility(varImportedAsFullyManagedVisibility)

	return err
}

type NullableImportedAsFullyManagedVisibility struct {
	value *ImportedAsFullyManagedVisibility
	isSet bool
}

func (v NullableImportedAsFullyManagedVisibility) Get() *ImportedAsFullyManagedVisibility {
	return v.value
}

func (v *NullableImportedAsFullyManagedVisibility) Set(val *ImportedAsFullyManagedVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedAsFullyManagedVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedAsFullyManagedVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedAsFullyManagedVisibility(val *ImportedAsFullyManagedVisibility) *NullableImportedAsFullyManagedVisibility {
	return &NullableImportedAsFullyManagedVisibility{value: val, isSet: true}
}

func (v NullableImportedAsFullyManagedVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedAsFullyManagedVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


