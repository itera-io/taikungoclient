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

// ActionType the model 'ActionType'
type ActionType string

// List of ActionType
const (
	ACTIONTYPE_START_ADD_SERVER               ActionType = "StartAddServer"
	ACTIONTYPE_ADDED_SERVER                   ActionType = "AddedServer"
	ACTIONTYPE_START_ADD_VM                   ActionType = "StartAddVm"
	ACTIONTYPE_START_ADD_VM_DISK              ActionType = "StartAddVmDisk"
	ACTIONTYPE_ADDED_VM_DISK                  ActionType = "AddedVmDisk"
	ACTIONTYPE_START_DELETE_VM_DISK           ActionType = "StartDeleteVmDisk"
	ACTIONTYPE_DELETED_VM_DISK                ActionType = "DeletedVmDisk"
	ACTIONTYPE_ADDED_VM                       ActionType = "AddedVm"
	ACTIONTYPE_START_DELETE_VM                ActionType = "StartDeleteVm"
	ACTIONTYPE_DELETED_VM                     ActionType = "DeletedVm"
	ACTIONTYPE_START_DELETE_SERVER            ActionType = "StartDeleteServer"
	ACTIONTYPE_DELETED_SERVER                 ActionType = "DeletedServer"
	ACTIONTYPE_COMMIT_CHANGES                 ActionType = "CommitChanges"
	ACTIONTYPE_CREATED_PROJECT                ActionType = "CreatedProject"
	ACTIONTYPE_START_PURGE_PROJECT            ActionType = "StartPurgeProject"
	ACTIONTYPE_PURGED_PROJECT                 ActionType = "PurgedProject"
	ACTIONTYPE_CAPACITY_NOT_AVAILABLE         ActionType = "CapacityNotAvailable"
	ACTIONTYPE_SPOT_FAILURE                   ActionType = "SpotFailure"
	ACTIONTYPE_CREATE_INFRA_FAILURE           ActionType = "CreateInfraFailure"
	ACTIONTYPE_DESTROY_INFRA_FAILURE          ActionType = "DestroyInfraFailure"
	ACTIONTYPE_LOGIN_FAILED                   ActionType = "LoginFailed"
	ACTIONTYPE_KUBERNETES_ERROR               ActionType = "KubernetesError"
	ACTIONTYPE_START_UPGRADE_PROJECT          ActionType = "StartUpgradeProject"
	ACTIONTYPE_UPGRADED_PROJECT               ActionType = "UpgradedProject"
	ACTIONTYPE_START_UPDATE_CLOUD_CREDENTIALS ActionType = "StartUpdateCloudCredentials"
	ACTIONTYPE_UPDATED_CLOUD_CREDENTIALS      ActionType = "UpdatedCloudCredentials"
	ACTIONTYPE_STARTED_ENABLE_GATEKEEPER      ActionType = "StartedEnableGatekeeper"
	ACTIONTYPE_STARTED_DISABLE_GATEKEEPER     ActionType = "StartedDisableGatekeeper"
	ACTIONTYPE_STARTED_ENABLE_BACKUP          ActionType = "StartedEnableBackup"
	ACTIONTYPE_STARTED_DISABLE_BACKUP         ActionType = "StartedDisableBackup"
	ACTIONTYPE_ENABLED_BACKUP                 ActionType = "EnabledBackup"
	ACTIONTYPE_DISABLED_BACKUP                ActionType = "DisabledBackup"
	ACTIONTYPE_ENABLED_GATEKEEPER             ActionType = "EnabledGatekeeper"
	ACTIONTYPE_DISABLED_GATEKEEPER            ActionType = "DisabledGatekeeper"
	ACTIONTYPE_ADDED_BACKUP_POLICY            ActionType = "AddedBackupPolicy"
	ACTIONTYPE_DELETED_BACKUP_POLICY          ActionType = "DeletedBackupPolicy"
	ACTIONTYPE_START_ENABLE_MONITORING        ActionType = "StartEnableMonitoring"
	ACTIONTYPE_ENABLED_MONITORING             ActionType = "EnabledMonitoring"
	ACTIONTYPE_START_DISABLE_MONITORING       ActionType = "StartDisableMonitoring"
	ACTIONTYPE_DISABLED_MONITORING            ActionType = "DisabledMonitoring"
	ACTIONTYPE_HEALTH_STATUS                  ActionType = "HealthStatus"
	ACTIONTYPE_SERVER_REBOOTED                ActionType = "ServerRebooted"
	ACTIONTYPE_SILENCE_ALERT                  ActionType = "SilenceAlert"
	ACTIONTYPE_ADDED_PROJECT                  ActionType = "AddedProject"
	ACTIONTYPE_ADDED_USER                     ActionType = "AddedUser"
	ACTIONTYPE_SERVER_STOPPED                 ActionType = "ServerStopped"
	ACTIONTYPE_SERVER_STARTED                 ActionType = "ServerStarted"
	ACTIONTYPE_SERVER_SHELVED                 ActionType = "ServerShelved"
	ACTIONTYPE_SERVER_UNSHELVED               ActionType = "ServerUnshelved"
	ACTIONTYPE_SERVER_CONSOLE                 ActionType = "ServerConsole"
	ACTIONTYPE_SERVER_STATUS                  ActionType = "ServerStatus"
	ACTIONTYPE_PATCH_KUBERNETES               ActionType = "PatchKubernetes"
	ACTIONTYPE_INSTALL_APPLICATION            ActionType = "InstallApplication"
	ACTIONTYPE_UNINSTALL_APPLICATION          ActionType = "UninstallApplication"
	ACTIONTYPE_ENABLE_AUTOSCALING             ActionType = "EnableAutoscaling"
	ACTIONTYPE_DISABLE_AUTOSCALING            ActionType = "DisableAutoscaling"
	ACTIONTYPE_BOUND_PROJECT                  ActionType = "BoundProject"
	ACTIONTYPE_LOCK_UNLOCK_PROJECT            ActionType = "LockUnlockProject"
	ACTIONTYPE_ENABLED_AI                     ActionType = "EnabledAi"
	ACTIONTYPE_STARTED_ENABLE_AI              ActionType = "StartedEnableAi"
	ACTIONTYPE_DISABLED_AI                    ActionType = "DisabledAi"
	ACTIONTYPE_STARTED_DISABLE_AI             ActionType = "StartedDisableAi"
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	"StartAddServer",
	"AddedServer",
	"StartAddVm",
	"StartAddVmDisk",
	"AddedVmDisk",
	"StartDeleteVmDisk",
	"DeletedVmDisk",
	"AddedVm",
	"StartDeleteVm",
	"DeletedVm",
	"StartDeleteServer",
	"DeletedServer",
	"CommitChanges",
	"CreatedProject",
	"StartPurgeProject",
	"PurgedProject",
	"CapacityNotAvailable",
	"SpotFailure",
	"CreateInfraFailure",
	"DestroyInfraFailure",
	"LoginFailed",
	"KubernetesError",
	"StartUpgradeProject",
	"UpgradedProject",
	"StartUpdateCloudCredentials",
	"UpdatedCloudCredentials",
	"StartedEnableGatekeeper",
	"StartedDisableGatekeeper",
	"StartedEnableBackup",
	"StartedDisableBackup",
	"EnabledBackup",
	"DisabledBackup",
	"EnabledGatekeeper",
	"DisabledGatekeeper",
	"AddedBackupPolicy",
	"DeletedBackupPolicy",
	"StartEnableMonitoring",
	"EnabledMonitoring",
	"StartDisableMonitoring",
	"DisabledMonitoring",
	"HealthStatus",
	"ServerRebooted",
	"SilenceAlert",
	"AddedProject",
	"AddedUser",
	"ServerStopped",
	"ServerStarted",
	"ServerShelved",
	"ServerUnshelved",
	"ServerConsole",
	"ServerStatus",
	"PatchKubernetes",
	"InstallApplication",
	"UninstallApplication",
	"EnableAutoscaling",
	"DisableAutoscaling",
	"BoundProject",
	"LockUnlockProject",
	"EnabledAi",
	"StartedEnableAi",
	"DisabledAi",
	"StartedDisableAi",
}

func (v *ActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionType(value)
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionType", value)
}

// NewActionTypeFromValue returns a pointer to a valid ActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionTypeFromValue(v string) (*ActionType, error) {
	ev := ActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionType: valid values are %v", v, AllowedActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionType value
func (v ActionType) Ptr() *ActionType {
	return &v
}

type NullableActionType struct {
	value *ActionType
	isSet bool
}

func (v NullableActionType) Get() *ActionType {
	return v.value
}

func (v *NullableActionType) Set(val *ActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionType(val *ActionType) *NullableActionType {
	return &NullableActionType{value: val, isSet: true}
}

func (v NullableActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}