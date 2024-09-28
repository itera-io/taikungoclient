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

// EImportClusterType the model 'EImportClusterType'
type EImportClusterType string

// List of EImportClusterType
const (
	EIMPORTCLUSTERTYPE_CLOUD_CREDENTIAL EImportClusterType = "CloudCredential"
	EIMPORTCLUSTERTYPE_READ_ONLY EImportClusterType = "ReadOnly"
	EIMPORTCLUSTERTYPE_FULLY_MANAGED EImportClusterType = "FullyManaged"
)

// All allowed values of EImportClusterType enum
var AllowedEImportClusterTypeEnumValues = []EImportClusterType{
	"CloudCredential",
	"ReadOnly",
	"FullyManaged",
}

func (v *EImportClusterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EImportClusterType(value)
	for _, existing := range AllowedEImportClusterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EImportClusterType", value)
}

// NewEImportClusterTypeFromValue returns a pointer to a valid EImportClusterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEImportClusterTypeFromValue(v string) (*EImportClusterType, error) {
	ev := EImportClusterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EImportClusterType: valid values are %v", v, AllowedEImportClusterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EImportClusterType) IsValid() bool {
	for _, existing := range AllowedEImportClusterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EImportClusterType value
func (v EImportClusterType) Ptr() *EImportClusterType {
	return &v
}

type NullableEImportClusterType struct {
	value *EImportClusterType
	isSet bool
}

func (v NullableEImportClusterType) Get() *EImportClusterType {
	return v.value
}

func (v *NullableEImportClusterType) Set(val *EImportClusterType) {
	v.value = val
	v.isSet = true
}

func (v NullableEImportClusterType) IsSet() bool {
	return v.isSet
}

func (v *NullableEImportClusterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEImportClusterType(val *EImportClusterType) *NullableEImportClusterType {
	return &NullableEImportClusterType{value: val, isSet: true}
}

func (v NullableEImportClusterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEImportClusterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

