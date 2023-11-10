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
	"fmt"
)

// EmailMode the model 'EmailMode'
type EmailMode string

// List of EmailMode
const (
	EMAILMODE_CONFIRM EmailMode = "Confirm"
	EMAILMODE_CHANGE EmailMode = "Change"
)

// All allowed values of EmailMode enum
var AllowedEmailModeEnumValues = []EmailMode{
	"Confirm",
	"Change",
}

func (v *EmailMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailMode(value)
	for _, existing := range AllowedEmailModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailMode", value)
}

// NewEmailModeFromValue returns a pointer to a valid EmailMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailModeFromValue(v string) (*EmailMode, error) {
	ev := EmailMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailMode: valid values are %v", v, AllowedEmailModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailMode) IsValid() bool {
	for _, existing := range AllowedEmailModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailMode value
func (v EmailMode) Ptr() *EmailMode {
	return &v
}

type NullableEmailMode struct {
	value *EmailMode
	isSet bool
}

func (v NullableEmailMode) Get() *EmailMode {
	return v.value
}

func (v *NullableEmailMode) Set(val *EmailMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMode(val *EmailMode) *NullableEmailMode {
	return &NullableEmailMode{value: val, isSet: true}
}

func (v NullableEmailMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

