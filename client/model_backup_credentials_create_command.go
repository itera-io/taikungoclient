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

// checks if the BackupCredentialsCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupCredentialsCreateCommand{}

// BackupCredentialsCreateCommand struct for BackupCredentialsCreateCommand
type BackupCredentialsCreateCommand struct {
	S3Name *string `json:"s3Name,omitempty"`
	S3AccessKeyId *string `json:"s3AccessKeyId,omitempty"`
	S3SecretKey *string `json:"s3SecretKey,omitempty"`
	S3Endpoint *string `json:"s3Endpoint,omitempty"`
	S3Region *string `json:"s3Region,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewBackupCredentialsCreateCommand instantiates a new BackupCredentialsCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupCredentialsCreateCommand() *BackupCredentialsCreateCommand {
	this := BackupCredentialsCreateCommand{}
	return &this
}

// NewBackupCredentialsCreateCommandWithDefaults instantiates a new BackupCredentialsCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupCredentialsCreateCommandWithDefaults() *BackupCredentialsCreateCommand {
	this := BackupCredentialsCreateCommand{}
	return &this
}

// GetS3Name returns the S3Name field value if set, zero value otherwise.
func (o *BackupCredentialsCreateCommand) GetS3Name() string {
	if o == nil || IsNil(o.S3Name) {
		var ret string
		return ret
	}
	return *o.S3Name
}

// GetS3NameOk returns a tuple with the S3Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsCreateCommand) GetS3NameOk() (*string, bool) {
	if o == nil || IsNil(o.S3Name) {
		return nil, false
	}
	return o.S3Name, true
}

// HasS3Name returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasS3Name() bool {
	if o != nil && !IsNil(o.S3Name) {
		return true
	}

	return false
}

// SetS3Name gets a reference to the given string and assigns it to the S3Name field.
func (o *BackupCredentialsCreateCommand) SetS3Name(v string) {
	o.S3Name = &v
}

// GetS3AccessKeyId returns the S3AccessKeyId field value if set, zero value otherwise.
func (o *BackupCredentialsCreateCommand) GetS3AccessKeyId() string {
	if o == nil || IsNil(o.S3AccessKeyId) {
		var ret string
		return ret
	}
	return *o.S3AccessKeyId
}

// GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsCreateCommand) GetS3AccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.S3AccessKeyId) {
		return nil, false
	}
	return o.S3AccessKeyId, true
}

// HasS3AccessKeyId returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasS3AccessKeyId() bool {
	if o != nil && !IsNil(o.S3AccessKeyId) {
		return true
	}

	return false
}

// SetS3AccessKeyId gets a reference to the given string and assigns it to the S3AccessKeyId field.
func (o *BackupCredentialsCreateCommand) SetS3AccessKeyId(v string) {
	o.S3AccessKeyId = &v
}

// GetS3SecretKey returns the S3SecretKey field value if set, zero value otherwise.
func (o *BackupCredentialsCreateCommand) GetS3SecretKey() string {
	if o == nil || IsNil(o.S3SecretKey) {
		var ret string
		return ret
	}
	return *o.S3SecretKey
}

// GetS3SecretKeyOk returns a tuple with the S3SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsCreateCommand) GetS3SecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.S3SecretKey) {
		return nil, false
	}
	return o.S3SecretKey, true
}

// HasS3SecretKey returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasS3SecretKey() bool {
	if o != nil && !IsNil(o.S3SecretKey) {
		return true
	}

	return false
}

// SetS3SecretKey gets a reference to the given string and assigns it to the S3SecretKey field.
func (o *BackupCredentialsCreateCommand) SetS3SecretKey(v string) {
	o.S3SecretKey = &v
}

// GetS3Endpoint returns the S3Endpoint field value if set, zero value otherwise.
func (o *BackupCredentialsCreateCommand) GetS3Endpoint() string {
	if o == nil || IsNil(o.S3Endpoint) {
		var ret string
		return ret
	}
	return *o.S3Endpoint
}

// GetS3EndpointOk returns a tuple with the S3Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsCreateCommand) GetS3EndpointOk() (*string, bool) {
	if o == nil || IsNil(o.S3Endpoint) {
		return nil, false
	}
	return o.S3Endpoint, true
}

// HasS3Endpoint returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasS3Endpoint() bool {
	if o != nil && !IsNil(o.S3Endpoint) {
		return true
	}

	return false
}

// SetS3Endpoint gets a reference to the given string and assigns it to the S3Endpoint field.
func (o *BackupCredentialsCreateCommand) SetS3Endpoint(v string) {
	o.S3Endpoint = &v
}

// GetS3Region returns the S3Region field value if set, zero value otherwise.
func (o *BackupCredentialsCreateCommand) GetS3Region() string {
	if o == nil || IsNil(o.S3Region) {
		var ret string
		return ret
	}
	return *o.S3Region
}

// GetS3RegionOk returns a tuple with the S3Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsCreateCommand) GetS3RegionOk() (*string, bool) {
	if o == nil || IsNil(o.S3Region) {
		return nil, false
	}
	return o.S3Region, true
}

// HasS3Region returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasS3Region() bool {
	if o != nil && !IsNil(o.S3Region) {
		return true
	}

	return false
}

// SetS3Region gets a reference to the given string and assigns it to the S3Region field.
func (o *BackupCredentialsCreateCommand) SetS3Region(v string) {
	o.S3Region = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsCreateCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsCreateCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *BackupCredentialsCreateCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *BackupCredentialsCreateCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *BackupCredentialsCreateCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *BackupCredentialsCreateCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o BackupCredentialsCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupCredentialsCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.S3Name) {
		toSerialize["s3Name"] = o.S3Name
	}
	if !IsNil(o.S3AccessKeyId) {
		toSerialize["s3AccessKeyId"] = o.S3AccessKeyId
	}
	if !IsNil(o.S3SecretKey) {
		toSerialize["s3SecretKey"] = o.S3SecretKey
	}
	if !IsNil(o.S3Endpoint) {
		toSerialize["s3Endpoint"] = o.S3Endpoint
	}
	if !IsNil(o.S3Region) {
		toSerialize["s3Region"] = o.S3Region
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableBackupCredentialsCreateCommand struct {
	value *BackupCredentialsCreateCommand
	isSet bool
}

func (v NullableBackupCredentialsCreateCommand) Get() *BackupCredentialsCreateCommand {
	return v.value
}

func (v *NullableBackupCredentialsCreateCommand) Set(val *BackupCredentialsCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupCredentialsCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupCredentialsCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupCredentialsCreateCommand(val *BackupCredentialsCreateCommand) *NullableBackupCredentialsCreateCommand {
	return &NullableBackupCredentialsCreateCommand{value: val, isSet: true}
}

func (v NullableBackupCredentialsCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupCredentialsCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


