/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
	"fmt"
)

// EShowbackType the model 'EShowbackType'
type EShowbackType string

// List of EShowbackType
const (
	ESHOWBACKTYPE_GENERAL  EShowbackType = "General"
	ESHOWBACKTYPE_EXTERNAL EShowbackType = "External"
)

// All allowed values of EShowbackType enum
var AllowedEShowbackTypeEnumValues = []EShowbackType{
	"General",
	"External",
}

func (v *EShowbackType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EShowbackType(value)
	for _, existing := range AllowedEShowbackTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EShowbackType", value)
}

// NewEShowbackTypeFromValue returns a pointer to a valid EShowbackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEShowbackTypeFromValue(v string) (*EShowbackType, error) {
	ev := EShowbackType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EShowbackType: valid values are %v", v, AllowedEShowbackTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EShowbackType) IsValid() bool {
	for _, existing := range AllowedEShowbackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EShowbackType value
func (v EShowbackType) Ptr() *EShowbackType {
	return &v
}

type NullableEShowbackType struct {
	value *EShowbackType
	isSet bool
}

func (v NullableEShowbackType) Get() *EShowbackType {
	return v.value
}

func (v *NullableEShowbackType) Set(val *EShowbackType) {
	v.value = val
	v.isSet = true
}

func (v NullableEShowbackType) IsSet() bool {
	return v.isSet
}

func (v *NullableEShowbackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEShowbackType(val *EShowbackType) *NullableEShowbackType {
	return &NullableEShowbackType{value: val, isSet: true}
}

func (v NullableEShowbackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEShowbackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
