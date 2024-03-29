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

// checks if the RestoreBackupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreBackupCommand{}

// RestoreBackupCommand struct for RestoreBackupCommand
type RestoreBackupCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	BackupName NullableString `json:"backupName,omitempty"`
	RestoreName NullableString `json:"restoreName,omitempty"`
	IncludeNamespaces []string `json:"includeNamespaces,omitempty"`
	ExcludeNamespaces []string `json:"excludeNamespaces,omitempty"`
}

// NewRestoreBackupCommand instantiates a new RestoreBackupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreBackupCommand() *RestoreBackupCommand {
	this := RestoreBackupCommand{}
	return &this
}

// NewRestoreBackupCommandWithDefaults instantiates a new RestoreBackupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreBackupCommandWithDefaults() *RestoreBackupCommand {
	this := RestoreBackupCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RestoreBackupCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreBackupCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RestoreBackupCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *RestoreBackupCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreBackupCommand) GetBackupName() string {
	if o == nil || IsNil(o.BackupName.Get()) {
		var ret string
		return ret
	}
	return *o.BackupName.Get()
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreBackupCommand) GetBackupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupName.Get(), o.BackupName.IsSet()
}

// HasBackupName returns a boolean if a field has been set.
func (o *RestoreBackupCommand) HasBackupName() bool {
	if o != nil && o.BackupName.IsSet() {
		return true
	}

	return false
}

// SetBackupName gets a reference to the given NullableString and assigns it to the BackupName field.
func (o *RestoreBackupCommand) SetBackupName(v string) {
	o.BackupName.Set(&v)
}
// SetBackupNameNil sets the value for BackupName to be an explicit nil
func (o *RestoreBackupCommand) SetBackupNameNil() {
	o.BackupName.Set(nil)
}

// UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
func (o *RestoreBackupCommand) UnsetBackupName() {
	o.BackupName.Unset()
}

// GetRestoreName returns the RestoreName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreBackupCommand) GetRestoreName() string {
	if o == nil || IsNil(o.RestoreName.Get()) {
		var ret string
		return ret
	}
	return *o.RestoreName.Get()
}

// GetRestoreNameOk returns a tuple with the RestoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreBackupCommand) GetRestoreNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestoreName.Get(), o.RestoreName.IsSet()
}

// HasRestoreName returns a boolean if a field has been set.
func (o *RestoreBackupCommand) HasRestoreName() bool {
	if o != nil && o.RestoreName.IsSet() {
		return true
	}

	return false
}

// SetRestoreName gets a reference to the given NullableString and assigns it to the RestoreName field.
func (o *RestoreBackupCommand) SetRestoreName(v string) {
	o.RestoreName.Set(&v)
}
// SetRestoreNameNil sets the value for RestoreName to be an explicit nil
func (o *RestoreBackupCommand) SetRestoreNameNil() {
	o.RestoreName.Set(nil)
}

// UnsetRestoreName ensures that no value is present for RestoreName, not even an explicit nil
func (o *RestoreBackupCommand) UnsetRestoreName() {
	o.RestoreName.Unset()
}

// GetIncludeNamespaces returns the IncludeNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreBackupCommand) GetIncludeNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeNamespaces
}

// GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreBackupCommand) GetIncludeNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeNamespaces) {
		return nil, false
	}
	return o.IncludeNamespaces, true
}

// HasIncludeNamespaces returns a boolean if a field has been set.
func (o *RestoreBackupCommand) HasIncludeNamespaces() bool {
	if o != nil && !IsNil(o.IncludeNamespaces) {
		return true
	}

	return false
}

// SetIncludeNamespaces gets a reference to the given []string and assigns it to the IncludeNamespaces field.
func (o *RestoreBackupCommand) SetIncludeNamespaces(v []string) {
	o.IncludeNamespaces = v
}

// GetExcludeNamespaces returns the ExcludeNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreBackupCommand) GetExcludeNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeNamespaces
}

// GetExcludeNamespacesOk returns a tuple with the ExcludeNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreBackupCommand) GetExcludeNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeNamespaces) {
		return nil, false
	}
	return o.ExcludeNamespaces, true
}

// HasExcludeNamespaces returns a boolean if a field has been set.
func (o *RestoreBackupCommand) HasExcludeNamespaces() bool {
	if o != nil && !IsNil(o.ExcludeNamespaces) {
		return true
	}

	return false
}

// SetExcludeNamespaces gets a reference to the given []string and assigns it to the ExcludeNamespaces field.
func (o *RestoreBackupCommand) SetExcludeNamespaces(v []string) {
	o.ExcludeNamespaces = v
}

func (o RestoreBackupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreBackupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.BackupName.IsSet() {
		toSerialize["backupName"] = o.BackupName.Get()
	}
	if o.RestoreName.IsSet() {
		toSerialize["restoreName"] = o.RestoreName.Get()
	}
	if o.IncludeNamespaces != nil {
		toSerialize["includeNamespaces"] = o.IncludeNamespaces
	}
	if o.ExcludeNamespaces != nil {
		toSerialize["excludeNamespaces"] = o.ExcludeNamespaces
	}
	return toSerialize, nil
}

type NullableRestoreBackupCommand struct {
	value *RestoreBackupCommand
	isSet bool
}

func (v NullableRestoreBackupCommand) Get() *RestoreBackupCommand {
	return v.value
}

func (v *NullableRestoreBackupCommand) Set(val *RestoreBackupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreBackupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreBackupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreBackupCommand(val *RestoreBackupCommand) *NullableRestoreBackupCommand {
	return &NullableRestoreBackupCommand{value: val, isSet: true}
}

func (v NullableRestoreBackupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreBackupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


