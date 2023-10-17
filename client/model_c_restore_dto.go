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
	"time"
)

// checks if the CRestoreDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CRestoreDto{}

// CRestoreDto struct for CRestoreDto
type CRestoreDto struct {
	MetadataName       NullableString `json:"metadataName,omitempty"`
	BackupName         NullableString `json:"backupName,omitempty"`
	ScheduleName       NullableString `json:"scheduleName,omitempty"`
	Namespace          NullableString `json:"namespace,omitempty"`
	ExcludeNamespaces  []string       `json:"excludeNamespaces,omitempty"`
	IncludeNamespaces  []string       `json:"includeNamespaces,omitempty"`
	CompletionDateTime *time.Time     `json:"completionDateTime,omitempty"`
	StartTimeStamp     *time.Time     `json:"startTimeStamp,omitempty"`
	CreatedAt          NullableTime   `json:"createdAt,omitempty"`
	Warnings           *int64         `json:"warnings,omitempty"`
	Phase              NullableString `json:"phase,omitempty"`
}

// NewCRestoreDto instantiates a new CRestoreDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCRestoreDto() *CRestoreDto {
	this := CRestoreDto{}
	return &this
}

// NewCRestoreDtoWithDefaults instantiates a new CRestoreDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCRestoreDtoWithDefaults() *CRestoreDto {
	this := CRestoreDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *CRestoreDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *CRestoreDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}

// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *CRestoreDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *CRestoreDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetBackupName returns the BackupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetBackupName() string {
	if o == nil || IsNil(o.BackupName.Get()) {
		var ret string
		return ret
	}
	return *o.BackupName.Get()
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetBackupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupName.Get(), o.BackupName.IsSet()
}

// HasBackupName returns a boolean if a field has been set.
func (o *CRestoreDto) HasBackupName() bool {
	if o != nil && o.BackupName.IsSet() {
		return true
	}

	return false
}

// SetBackupName gets a reference to the given NullableString and assigns it to the BackupName field.
func (o *CRestoreDto) SetBackupName(v string) {
	o.BackupName.Set(&v)
}

// SetBackupNameNil sets the value for BackupName to be an explicit nil
func (o *CRestoreDto) SetBackupNameNil() {
	o.BackupName.Set(nil)
}

// UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
func (o *CRestoreDto) UnsetBackupName() {
	o.BackupName.Unset()
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetScheduleName() string {
	if o == nil || IsNil(o.ScheduleName.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduleName.Get()
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetScheduleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleName.Get(), o.ScheduleName.IsSet()
}

// HasScheduleName returns a boolean if a field has been set.
func (o *CRestoreDto) HasScheduleName() bool {
	if o != nil && o.ScheduleName.IsSet() {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given NullableString and assigns it to the ScheduleName field.
func (o *CRestoreDto) SetScheduleName(v string) {
	o.ScheduleName.Set(&v)
}

// SetScheduleNameNil sets the value for ScheduleName to be an explicit nil
func (o *CRestoreDto) SetScheduleNameNil() {
	o.ScheduleName.Set(nil)
}

// UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
func (o *CRestoreDto) UnsetScheduleName() {
	o.ScheduleName.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *CRestoreDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *CRestoreDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *CRestoreDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *CRestoreDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetExcludeNamespaces returns the ExcludeNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetExcludeNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeNamespaces
}

// GetExcludeNamespacesOk returns a tuple with the ExcludeNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetExcludeNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeNamespaces) {
		return nil, false
	}
	return o.ExcludeNamespaces, true
}

// HasExcludeNamespaces returns a boolean if a field has been set.
func (o *CRestoreDto) HasExcludeNamespaces() bool {
	if o != nil && IsNil(o.ExcludeNamespaces) {
		return true
	}

	return false
}

// SetExcludeNamespaces gets a reference to the given []string and assigns it to the ExcludeNamespaces field.
func (o *CRestoreDto) SetExcludeNamespaces(v []string) {
	o.ExcludeNamespaces = v
}

// GetIncludeNamespaces returns the IncludeNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetIncludeNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeNamespaces
}

// GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetIncludeNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeNamespaces) {
		return nil, false
	}
	return o.IncludeNamespaces, true
}

// HasIncludeNamespaces returns a boolean if a field has been set.
func (o *CRestoreDto) HasIncludeNamespaces() bool {
	if o != nil && IsNil(o.IncludeNamespaces) {
		return true
	}

	return false
}

// SetIncludeNamespaces gets a reference to the given []string and assigns it to the IncludeNamespaces field.
func (o *CRestoreDto) SetIncludeNamespaces(v []string) {
	o.IncludeNamespaces = v
}

// GetCompletionDateTime returns the CompletionDateTime field value if set, zero value otherwise.
func (o *CRestoreDto) GetCompletionDateTime() time.Time {
	if o == nil || IsNil(o.CompletionDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDateTime
}

// GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CRestoreDto) GetCompletionDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionDateTime) {
		return nil, false
	}
	return o.CompletionDateTime, true
}

// HasCompletionDateTime returns a boolean if a field has been set.
func (o *CRestoreDto) HasCompletionDateTime() bool {
	if o != nil && !IsNil(o.CompletionDateTime) {
		return true
	}

	return false
}

// SetCompletionDateTime gets a reference to the given time.Time and assigns it to the CompletionDateTime field.
func (o *CRestoreDto) SetCompletionDateTime(v time.Time) {
	o.CompletionDateTime = &v
}

// GetStartTimeStamp returns the StartTimeStamp field value if set, zero value otherwise.
func (o *CRestoreDto) GetStartTimeStamp() time.Time {
	if o == nil || IsNil(o.StartTimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.StartTimeStamp
}

// GetStartTimeStampOk returns a tuple with the StartTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CRestoreDto) GetStartTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTimeStamp) {
		return nil, false
	}
	return o.StartTimeStamp, true
}

// HasStartTimeStamp returns a boolean if a field has been set.
func (o *CRestoreDto) HasStartTimeStamp() bool {
	if o != nil && !IsNil(o.StartTimeStamp) {
		return true
	}

	return false
}

// SetStartTimeStamp gets a reference to the given time.Time and assigns it to the StartTimeStamp field.
func (o *CRestoreDto) SetStartTimeStamp(v time.Time) {
	o.StartTimeStamp = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CRestoreDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *CRestoreDto) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CRestoreDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CRestoreDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CRestoreDto) GetWarnings() int64 {
	if o == nil || IsNil(o.Warnings) {
		var ret int64
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CRestoreDto) GetWarningsOk() (*int64, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CRestoreDto) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given int64 and assigns it to the Warnings field.
func (o *CRestoreDto) SetWarnings(v int64) {
	o.Warnings = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CRestoreDto) GetPhase() string {
	if o == nil || IsNil(o.Phase.Get()) {
		var ret string
		return ret
	}
	return *o.Phase.Get()
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CRestoreDto) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase.Get(), o.Phase.IsSet()
}

// HasPhase returns a boolean if a field has been set.
func (o *CRestoreDto) HasPhase() bool {
	if o != nil && o.Phase.IsSet() {
		return true
	}

	return false
}

// SetPhase gets a reference to the given NullableString and assigns it to the Phase field.
func (o *CRestoreDto) SetPhase(v string) {
	o.Phase.Set(&v)
}

// SetPhaseNil sets the value for Phase to be an explicit nil
func (o *CRestoreDto) SetPhaseNil() {
	o.Phase.Set(nil)
}

// UnsetPhase ensures that no value is present for Phase, not even an explicit nil
func (o *CRestoreDto) UnsetPhase() {
	o.Phase.Unset()
}

func (o CRestoreDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CRestoreDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.BackupName.IsSet() {
		toSerialize["backupName"] = o.BackupName.Get()
	}
	if o.ScheduleName.IsSet() {
		toSerialize["scheduleName"] = o.ScheduleName.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.ExcludeNamespaces != nil {
		toSerialize["excludeNamespaces"] = o.ExcludeNamespaces
	}
	if o.IncludeNamespaces != nil {
		toSerialize["includeNamespaces"] = o.IncludeNamespaces
	}
	if !IsNil(o.CompletionDateTime) {
		toSerialize["completionDateTime"] = o.CompletionDateTime
	}
	if !IsNil(o.StartTimeStamp) {
		toSerialize["startTimeStamp"] = o.StartTimeStamp
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Phase.IsSet() {
		toSerialize["phase"] = o.Phase.Get()
	}
	return toSerialize, nil
}

type NullableCRestoreDto struct {
	value *CRestoreDto
	isSet bool
}

func (v NullableCRestoreDto) Get() *CRestoreDto {
	return v.value
}

func (v *NullableCRestoreDto) Set(val *CRestoreDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCRestoreDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCRestoreDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCRestoreDto(val *CRestoreDto) *NullableCRestoreDto {
	return &NullableCRestoreDto{value: val, isSet: true}
}

func (v NullableCRestoreDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCRestoreDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
