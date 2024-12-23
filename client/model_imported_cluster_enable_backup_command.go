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
)

// checks if the ImportedClusterEnableBackupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedClusterEnableBackupCommand{}

// ImportedClusterEnableBackupCommand struct for ImportedClusterEnableBackupCommand
type ImportedClusterEnableBackupCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	BackupCredentialId *int32 `json:"backupCredentialId,omitempty"`
}

// NewImportedClusterEnableBackupCommand instantiates a new ImportedClusterEnableBackupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedClusterEnableBackupCommand() *ImportedClusterEnableBackupCommand {
	this := ImportedClusterEnableBackupCommand{}
	return &this
}

// NewImportedClusterEnableBackupCommandWithDefaults instantiates a new ImportedClusterEnableBackupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedClusterEnableBackupCommandWithDefaults() *ImportedClusterEnableBackupCommand {
	this := ImportedClusterEnableBackupCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ImportedClusterEnableBackupCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportedClusterEnableBackupCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ImportedClusterEnableBackupCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ImportedClusterEnableBackupCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetBackupCredentialId returns the BackupCredentialId field value if set, zero value otherwise.
func (o *ImportedClusterEnableBackupCommand) GetBackupCredentialId() int32 {
	if o == nil || IsNil(o.BackupCredentialId) {
		var ret int32
		return ret
	}
	return *o.BackupCredentialId
}

// GetBackupCredentialIdOk returns a tuple with the BackupCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportedClusterEnableBackupCommand) GetBackupCredentialIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BackupCredentialId) {
		return nil, false
	}
	return o.BackupCredentialId, true
}

// HasBackupCredentialId returns a boolean if a field has been set.
func (o *ImportedClusterEnableBackupCommand) HasBackupCredentialId() bool {
	if o != nil && !IsNil(o.BackupCredentialId) {
		return true
	}

	return false
}

// SetBackupCredentialId gets a reference to the given int32 and assigns it to the BackupCredentialId field.
func (o *ImportedClusterEnableBackupCommand) SetBackupCredentialId(v int32) {
	o.BackupCredentialId = &v
}

func (o ImportedClusterEnableBackupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedClusterEnableBackupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.BackupCredentialId) {
		toSerialize["backupCredentialId"] = o.BackupCredentialId
	}
	return toSerialize, nil
}

type NullableImportedClusterEnableBackupCommand struct {
	value *ImportedClusterEnableBackupCommand
	isSet bool
}

func (v NullableImportedClusterEnableBackupCommand) Get() *ImportedClusterEnableBackupCommand {
	return v.value
}

func (v *NullableImportedClusterEnableBackupCommand) Set(val *ImportedClusterEnableBackupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedClusterEnableBackupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedClusterEnableBackupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedClusterEnableBackupCommand(val *ImportedClusterEnableBackupCommand) *NullableImportedClusterEnableBackupCommand {
	return &NullableImportedClusterEnableBackupCommand{value: val, isSet: true}
}

func (v NullableImportedClusterEnableBackupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedClusterEnableBackupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


