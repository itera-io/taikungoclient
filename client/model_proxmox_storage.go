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

// ProxmoxStorage the model 'ProxmoxStorage'
type ProxmoxStorage string

// List of ProxmoxStorage
const (
	PROXMOXSTORAGE_NFS ProxmoxStorage = "NFS"
	PROXMOXSTORAGE_OPEN_EBS ProxmoxStorage = "OpenEBS"
	PROXMOXSTORAGE_LONGHORN ProxmoxStorage = "Longhorn"
	PROXMOXSTORAGE_LOCAL_PATH ProxmoxStorage = "LocalPath"
)

// All allowed values of ProxmoxStorage enum
var AllowedProxmoxStorageEnumValues = []ProxmoxStorage{
	"NFS",
	"OpenEBS",
	"Longhorn",
	"LocalPath",
}

func (v *ProxmoxStorage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxmoxStorage(value)
	for _, existing := range AllowedProxmoxStorageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxmoxStorage", value)
}

// NewProxmoxStorageFromValue returns a pointer to a valid ProxmoxStorage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxmoxStorageFromValue(v string) (*ProxmoxStorage, error) {
	ev := ProxmoxStorage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxmoxStorage: valid values are %v", v, AllowedProxmoxStorageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxmoxStorage) IsValid() bool {
	for _, existing := range AllowedProxmoxStorageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxmoxStorage value
func (v ProxmoxStorage) Ptr() *ProxmoxStorage {
	return &v
}

type NullableProxmoxStorage struct {
	value *ProxmoxStorage
	isSet bool
}

func (v NullableProxmoxStorage) Get() *ProxmoxStorage {
	return v.value
}

func (v *NullableProxmoxStorage) Set(val *ProxmoxStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxStorage(val *ProxmoxStorage) *NullableProxmoxStorage {
	return &NullableProxmoxStorage{value: val, isSet: true}
}

func (v NullableProxmoxStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

