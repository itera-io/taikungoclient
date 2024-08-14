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

// ProjectStatus the model 'ProjectStatus'
type ProjectStatus string

// List of ProjectStatus
const (
	PROJECTSTATUS_NULL ProjectStatus = "Null"
	PROJECTSTATUS_DELETING ProjectStatus = "Deleting"
	PROJECTSTATUS_DISABLE_AI ProjectStatus = "DisableAi"
	PROJECTSTATUS_DISABLE_GATEKEEPER ProjectStatus = "DisableGatekeeper"
	PROJECTSTATUS_DELETING_BACKUP_POLICIES ProjectStatus = "DeletingBackupPolicies"
	PROJECTSTATUS_DISABLE_BACKUP ProjectStatus = "DisableBackup"
	PROJECTSTATUS_DISABLE_MONITORING ProjectStatus = "DisableMonitoring"
	PROJECTSTATUS_ENABLE_BACKUP ProjectStatus = "EnableBackup"
	PROJECTSTATUS_ENABLE_MONITORING ProjectStatus = "EnableMonitoring"
	PROJECTSTATUS_ENABLE_AI ProjectStatus = "EnableAi"
	PROJECTSTATUS_ENABLE_GATEKEEPER ProjectStatus = "EnableGatekeeper"
	PROJECTSTATUS_FAILURE ProjectStatus = "Failure"
	PROJECTSTATUS_FAILED_UPGRADE ProjectStatus = "FailedUpgrade"
	PROJECTSTATUS_PURGING ProjectStatus = "Purging"
	PROJECTSTATUS_READY ProjectStatus = "Ready"
	PROJECTSTATUS_UPDATING ProjectStatus = "Updating"
	PROJECTSTATUS_UPGRADING ProjectStatus = "Upgrading"
	PROJECTSTATUS_UPDATING_CREDENTIALS ProjectStatus = "UpdatingCredentials"
	PROJECTSTATUS_UNINSTALL_VIRTUAL_CLUSTER ProjectStatus = "UninstallVirtualCluster"
	PROJECTSTATUS_REPAIRING ProjectStatus = "Repairing"
)

// All allowed values of ProjectStatus enum
var AllowedProjectStatusEnumValues = []ProjectStatus{
	"Null",
	"Deleting",
	"DisableAi",
	"DisableGatekeeper",
	"DeletingBackupPolicies",
	"DisableBackup",
	"DisableMonitoring",
	"EnableBackup",
	"EnableMonitoring",
	"EnableAi",
	"EnableGatekeeper",
	"Failure",
	"FailedUpgrade",
	"Purging",
	"Ready",
	"Updating",
	"Upgrading",
	"UpdatingCredentials",
	"UninstallVirtualCluster",
	"Repairing",
}

func (v *ProjectStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectStatus(value)
	for _, existing := range AllowedProjectStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectStatus", value)
}

// NewProjectStatusFromValue returns a pointer to a valid ProjectStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectStatusFromValue(v string) (*ProjectStatus, error) {
	ev := ProjectStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectStatus: valid values are %v", v, AllowedProjectStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectStatus) IsValid() bool {
	for _, existing := range AllowedProjectStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectStatus value
func (v ProjectStatus) Ptr() *ProjectStatus {
	return &v
}

type NullableProjectStatus struct {
	value *ProjectStatus
	isSet bool
}

func (v NullableProjectStatus) Get() *ProjectStatus {
	return v.value
}

func (v *NullableProjectStatus) Set(val *ProjectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStatus(val *ProjectStatus) *NullableProjectStatus {
	return &NullableProjectStatus{value: val, isSet: true}
}

func (v NullableProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

