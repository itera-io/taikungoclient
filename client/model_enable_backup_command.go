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

// checks if the EnableBackupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableBackupCommand{}

// EnableBackupCommand struct for EnableBackupCommand
type EnableBackupCommand struct {
	ProjectId      *int32 `json:"projectId,omitempty"`
	S3CredentialId *int32 `json:"s3CredentialId,omitempty"`
}

// NewEnableBackupCommand instantiates a new EnableBackupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableBackupCommand() *EnableBackupCommand {
	this := EnableBackupCommand{}
	return &this
}

// NewEnableBackupCommandWithDefaults instantiates a new EnableBackupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableBackupCommandWithDefaults() *EnableBackupCommand {
	this := EnableBackupCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *EnableBackupCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableBackupCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *EnableBackupCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *EnableBackupCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetS3CredentialId returns the S3CredentialId field value if set, zero value otherwise.
func (o *EnableBackupCommand) GetS3CredentialId() int32 {
	if o == nil || IsNil(o.S3CredentialId) {
		var ret int32
		return ret
	}
	return *o.S3CredentialId
}

// GetS3CredentialIdOk returns a tuple with the S3CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableBackupCommand) GetS3CredentialIdOk() (*int32, bool) {
	if o == nil || IsNil(o.S3CredentialId) {
		return nil, false
	}
	return o.S3CredentialId, true
}

// HasS3CredentialId returns a boolean if a field has been set.
func (o *EnableBackupCommand) HasS3CredentialId() bool {
	if o != nil && !IsNil(o.S3CredentialId) {
		return true
	}

	return false
}

// SetS3CredentialId gets a reference to the given int32 and assigns it to the S3CredentialId field.
func (o *EnableBackupCommand) SetS3CredentialId(v int32) {
	o.S3CredentialId = &v
}

func (o EnableBackupCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableBackupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.S3CredentialId) {
		toSerialize["s3CredentialId"] = o.S3CredentialId
	}
	return toSerialize, nil
}

type NullableEnableBackupCommand struct {
	value *EnableBackupCommand
	isSet bool
}

func (v NullableEnableBackupCommand) Get() *EnableBackupCommand {
	return v.value
}

func (v *NullableEnableBackupCommand) Set(val *EnableBackupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableBackupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableBackupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableBackupCommand(val *EnableBackupCommand) *NullableEnableBackupCommand {
	return &NullableEnableBackupCommand{value: val, isSet: true}
}

func (v NullableEnableBackupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableBackupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
