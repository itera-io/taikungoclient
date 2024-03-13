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

// EInstanceStatus the model 'EInstanceStatus'
type EInstanceStatus string

// List of EInstanceStatus
const (
	EINSTANCESTATUS_NONE EInstanceStatus = "None"
	EINSTANCESTATUS_INSTALLING EInstanceStatus = "Installing"
	EINSTANCESTATUS_UNINSTALLING EInstanceStatus = "Uninstalling"
	EINSTANCESTATUS_READY EInstanceStatus = "Ready"
	EINSTANCESTATUS_FAILED EInstanceStatus = "Failed"
	EINSTANCESTATUS_NOT_READY EInstanceStatus = "NotReady"
	EINSTANCESTATUS_FAILURE EInstanceStatus = "Failure"
)

// All allowed values of EInstanceStatus enum
var AllowedEInstanceStatusEnumValues = []EInstanceStatus{
	"None",
	"Installing",
	"Uninstalling",
	"Ready",
	"Failed",
	"NotReady",
	"Failure",
}

func (v *EInstanceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EInstanceStatus(value)
	for _, existing := range AllowedEInstanceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EInstanceStatus", value)
}

// NewEInstanceStatusFromValue returns a pointer to a valid EInstanceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEInstanceStatusFromValue(v string) (*EInstanceStatus, error) {
	ev := EInstanceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EInstanceStatus: valid values are %v", v, AllowedEInstanceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EInstanceStatus) IsValid() bool {
	for _, existing := range AllowedEInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EInstanceStatus value
func (v EInstanceStatus) Ptr() *EInstanceStatus {
	return &v
}

type NullableEInstanceStatus struct {
	value *EInstanceStatus
	isSet bool
}

func (v NullableEInstanceStatus) Get() *EInstanceStatus {
	return v.value
}

func (v *NullableEInstanceStatus) Set(val *EInstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEInstanceStatus(val *EInstanceStatus) *NullableEInstanceStatus {
	return &NullableEInstanceStatus{value: val, isSet: true}
}

func (v NullableEInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
