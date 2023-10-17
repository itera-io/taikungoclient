/*
Taikun - WebApi

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

// SecurityGroupProtocol the model 'SecurityGroupProtocol'
type SecurityGroupProtocol string

// List of SecurityGroupProtocol
const (
	SECURITYGROUPPROTOCOL_ICMP SecurityGroupProtocol = "ICMP"
	SECURITYGROUPPROTOCOL_TCP  SecurityGroupProtocol = "TCP"
	SECURITYGROUPPROTOCOL_UDP  SecurityGroupProtocol = "UDP"
)

// All allowed values of SecurityGroupProtocol enum
var AllowedSecurityGroupProtocolEnumValues = []SecurityGroupProtocol{
	"ICMP",
	"TCP",
	"UDP",
}

func (v *SecurityGroupProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityGroupProtocol(value)
	for _, existing := range AllowedSecurityGroupProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityGroupProtocol", value)
}

// NewSecurityGroupProtocolFromValue returns a pointer to a valid SecurityGroupProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityGroupProtocolFromValue(v string) (*SecurityGroupProtocol, error) {
	ev := SecurityGroupProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityGroupProtocol: valid values are %v", v, AllowedSecurityGroupProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityGroupProtocol) IsValid() bool {
	for _, existing := range AllowedSecurityGroupProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityGroupProtocol value
func (v SecurityGroupProtocol) Ptr() *SecurityGroupProtocol {
	return &v
}

type NullableSecurityGroupProtocol struct {
	value *SecurityGroupProtocol
	isSet bool
}

func (v NullableSecurityGroupProtocol) Get() *SecurityGroupProtocol {
	return v.value
}

func (v *NullableSecurityGroupProtocol) Set(val *SecurityGroupProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupProtocol(val *SecurityGroupProtocol) *NullableSecurityGroupProtocol {
	return &NullableSecurityGroupProtocol{value: val, isSet: true}
}

func (v NullableSecurityGroupProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
