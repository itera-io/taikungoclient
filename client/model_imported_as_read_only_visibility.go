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

// checks if the ImportedAsReadOnlyVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedAsReadOnlyVisibility{}

// ImportedAsReadOnlyVisibility struct for ImportedAsReadOnlyVisibility
type ImportedAsReadOnlyVisibility struct {
	Lock ButtonStatusDto `json:"lock"`
	Unlock ButtonStatusDto `json:"unlock"`
}

type _ImportedAsReadOnlyVisibility ImportedAsReadOnlyVisibility

// NewImportedAsReadOnlyVisibility instantiates a new ImportedAsReadOnlyVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedAsReadOnlyVisibility(lock ButtonStatusDto, unlock ButtonStatusDto) *ImportedAsReadOnlyVisibility {
	this := ImportedAsReadOnlyVisibility{}
	this.Lock = lock
	this.Unlock = unlock
	return &this
}

// NewImportedAsReadOnlyVisibilityWithDefaults instantiates a new ImportedAsReadOnlyVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedAsReadOnlyVisibilityWithDefaults() *ImportedAsReadOnlyVisibility {
	this := ImportedAsReadOnlyVisibility{}
	return &this
}

// GetLock returns the Lock field value
func (o *ImportedAsReadOnlyVisibility) GetLock() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.Lock
}

// GetLockOk returns a tuple with the Lock field value
// and a boolean to check if the value has been set.
func (o *ImportedAsReadOnlyVisibility) GetLockOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lock, true
}

// SetLock sets field value
func (o *ImportedAsReadOnlyVisibility) SetLock(v ButtonStatusDto) {
	o.Lock = v
}

// GetUnlock returns the Unlock field value
func (o *ImportedAsReadOnlyVisibility) GetUnlock() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.Unlock
}

// GetUnlockOk returns a tuple with the Unlock field value
// and a boolean to check if the value has been set.
func (o *ImportedAsReadOnlyVisibility) GetUnlockOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unlock, true
}

// SetUnlock sets field value
func (o *ImportedAsReadOnlyVisibility) SetUnlock(v ButtonStatusDto) {
	o.Unlock = v
}

func (o ImportedAsReadOnlyVisibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedAsReadOnlyVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lock"] = o.Lock
	toSerialize["unlock"] = o.Unlock
	return toSerialize, nil
}

func (o *ImportedAsReadOnlyVisibility) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lock",
		"unlock",
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

	varImportedAsReadOnlyVisibility := _ImportedAsReadOnlyVisibility{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportedAsReadOnlyVisibility)

	if err != nil {
		return err
	}

	*o = ImportedAsReadOnlyVisibility(varImportedAsReadOnlyVisibility)

	return err
}

type NullableImportedAsReadOnlyVisibility struct {
	value *ImportedAsReadOnlyVisibility
	isSet bool
}

func (v NullableImportedAsReadOnlyVisibility) Get() *ImportedAsReadOnlyVisibility {
	return v.value
}

func (v *NullableImportedAsReadOnlyVisibility) Set(val *ImportedAsReadOnlyVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedAsReadOnlyVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedAsReadOnlyVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedAsReadOnlyVisibility(val *ImportedAsReadOnlyVisibility) *NullableImportedAsReadOnlyVisibility {
	return &NullableImportedAsReadOnlyVisibility{value: val, isSet: true}
}

func (v NullableImportedAsReadOnlyVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedAsReadOnlyVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


