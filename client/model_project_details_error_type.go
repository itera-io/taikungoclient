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

// ProjectDetailsErrorType the model 'ProjectDetailsErrorType'
type ProjectDetailsErrorType string

// List of ProjectDetailsErrorType
const (
	PROJECTDETAILSERRORTYPE_DANGER ProjectDetailsErrorType = "danger"
	PROJECTDETAILSERRORTYPE_INFO ProjectDetailsErrorType = "info"
	PROJECTDETAILSERRORTYPE_WARNING ProjectDetailsErrorType = "warning"
)

// All allowed values of ProjectDetailsErrorType enum
var AllowedProjectDetailsErrorTypeEnumValues = []ProjectDetailsErrorType{
	"danger",
	"info",
	"warning",
}

func (v *ProjectDetailsErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectDetailsErrorType(value)
	for _, existing := range AllowedProjectDetailsErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectDetailsErrorType", value)
}

// NewProjectDetailsErrorTypeFromValue returns a pointer to a valid ProjectDetailsErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectDetailsErrorTypeFromValue(v string) (*ProjectDetailsErrorType, error) {
	ev := ProjectDetailsErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectDetailsErrorType: valid values are %v", v, AllowedProjectDetailsErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectDetailsErrorType) IsValid() bool {
	for _, existing := range AllowedProjectDetailsErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectDetailsErrorType value
func (v ProjectDetailsErrorType) Ptr() *ProjectDetailsErrorType {
	return &v
}

type NullableProjectDetailsErrorType struct {
	value *ProjectDetailsErrorType
	isSet bool
}

func (v NullableProjectDetailsErrorType) Get() *ProjectDetailsErrorType {
	return v.value
}

func (v *NullableProjectDetailsErrorType) Set(val *ProjectDetailsErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetailsErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetailsErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetailsErrorType(val *ProjectDetailsErrorType) *NullableProjectDetailsErrorType {
	return &NullableProjectDetailsErrorType{value: val, isSet: true}
}

func (v NullableProjectDetailsErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetailsErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

