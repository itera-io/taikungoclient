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

// EKubernetesState the model 'EKubernetesState'
type EKubernetesState string

// List of EKubernetesState
const (
	EKUBERNETESSTATE_TERMINATING EKubernetesState = "Terminating"
	EKUBERNETESSTATE_UNKNOWN EKubernetesState = "Unknown"
	EKUBERNETESSTATE_EVICTED EKubernetesState = "Evicted"
	EKUBERNETESSTATE_CRASH_LOOP_BACK_OFF EKubernetesState = "CrashLoopBackOff"
	EKUBERNETESSTATE_IMAGE_PULL_BACK_OFF EKubernetesState = "ImagePullBackOff"
	EKUBERNETESSTATE_OOM_KILLED EKubernetesState = "OOMKilled"
	EKUBERNETESSTATE_COMPLETED EKubernetesState = "Completed"
	EKUBERNETESSTATE_RUNNING EKubernetesState = "Running"
	EKUBERNETESSTATE_INIT_NM EKubernetesState = "InitNM"
	EKUBERNETESSTATE_INIT_ERROR EKubernetesState = "InitError"
)

// All allowed values of EKubernetesState enum
var AllowedEKubernetesStateEnumValues = []EKubernetesState{
	"Terminating",
	"Unknown",
	"Evicted",
	"CrashLoopBackOff",
	"ImagePullBackOff",
	"OOMKilled",
	"Completed",
	"Running",
	"InitNM",
	"InitError",
}

func (v *EKubernetesState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EKubernetesState(value)
	for _, existing := range AllowedEKubernetesStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EKubernetesState", value)
}

// NewEKubernetesStateFromValue returns a pointer to a valid EKubernetesState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEKubernetesStateFromValue(v string) (*EKubernetesState, error) {
	ev := EKubernetesState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EKubernetesState: valid values are %v", v, AllowedEKubernetesStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EKubernetesState) IsValid() bool {
	for _, existing := range AllowedEKubernetesStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EKubernetesState value
func (v EKubernetesState) Ptr() *EKubernetesState {
	return &v
}

type NullableEKubernetesState struct {
	value *EKubernetesState
	isSet bool
}

func (v NullableEKubernetesState) Get() *EKubernetesState {
	return v.value
}

func (v *NullableEKubernetesState) Set(val *EKubernetesState) {
	v.value = val
	v.isSet = true
}

func (v NullableEKubernetesState) IsSet() bool {
	return v.isSet
}

func (v *NullableEKubernetesState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEKubernetesState(val *EKubernetesState) *NullableEKubernetesState {
	return &NullableEKubernetesState{value: val, isSet: true}
}

func (v NullableEKubernetesState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEKubernetesState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

