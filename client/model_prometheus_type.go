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

// PrometheusType the model 'PrometheusType'
type PrometheusType string

// List of PrometheusType
const (
	PROMETHEUSTYPE_COUNT PrometheusType = "Count"
	PROMETHEUSTYPE_SUM   PrometheusType = "Sum"
)

// All allowed values of PrometheusType enum
var AllowedPrometheusTypeEnumValues = []PrometheusType{
	"Count",
	"Sum",
}

func (v *PrometheusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrometheusType(value)
	for _, existing := range AllowedPrometheusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrometheusType", value)
}

// NewPrometheusTypeFromValue returns a pointer to a valid PrometheusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrometheusTypeFromValue(v string) (*PrometheusType, error) {
	ev := PrometheusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrometheusType: valid values are %v", v, AllowedPrometheusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrometheusType) IsValid() bool {
	for _, existing := range AllowedPrometheusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrometheusType value
func (v PrometheusType) Ptr() *PrometheusType {
	return &v
}

type NullablePrometheusType struct {
	value *PrometheusType
	isSet bool
}

func (v NullablePrometheusType) Get() *PrometheusType {
	return v.value
}

func (v *NullablePrometheusType) Set(val *PrometheusType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusType(val *PrometheusType) *NullablePrometheusType {
	return &NullablePrometheusType{value: val, isSet: true}
}

func (v NullablePrometheusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
