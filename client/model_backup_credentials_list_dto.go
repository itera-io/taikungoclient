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

// checks if the BackupCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupCredentialsListDto{}

// BackupCredentialsListDto struct for BackupCredentialsListDto
type BackupCredentialsListDto struct {
	Id               *int32              `json:"id,omitempty"`
	S3Name           NullableString      `json:"s3Name,omitempty"`
	S3AccessKeyId    NullableString      `json:"s3AccessKeyId,omitempty"`
	S3Endpoint       NullableString      `json:"s3Endpoint,omitempty"`
	S3Region         NullableString      `json:"s3Region,omitempty"`
	OrganizationId   NullableInt32       `json:"organizationId,omitempty"`
	OrganizationName NullableString      `json:"organizationName,omitempty"`
	Projects         []CommonDropdownDto `json:"projects,omitempty"`
	IsLocked         *bool               `json:"isLocked,omitempty"`
	CreatedBy        NullableString      `json:"createdBy,omitempty"`
	LastModified     NullableString      `json:"lastModified,omitempty"`
	LastModifiedBy   NullableString      `json:"lastModifiedBy,omitempty"`
	CreatedAt        NullableString      `json:"createdAt,omitempty"`
	IsDefault        *bool               `json:"isDefault,omitempty"`
	IsInfra          *bool               `json:"isInfra,omitempty"`
}

// NewBackupCredentialsListDto instantiates a new BackupCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupCredentialsListDto() *BackupCredentialsListDto {
	this := BackupCredentialsListDto{}
	return &this
}

// NewBackupCredentialsListDtoWithDefaults instantiates a new BackupCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupCredentialsListDtoWithDefaults() *BackupCredentialsListDto {
	this := BackupCredentialsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupCredentialsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BackupCredentialsListDto) SetId(v int32) {
	o.Id = &v
}

// GetS3Name returns the S3Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetS3Name() string {
	if o == nil || IsNil(o.S3Name.Get()) {
		var ret string
		return ret
	}
	return *o.S3Name.Get()
}

// GetS3NameOk returns a tuple with the S3Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetS3NameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Name.Get(), o.S3Name.IsSet()
}

// HasS3Name returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasS3Name() bool {
	if o != nil && o.S3Name.IsSet() {
		return true
	}

	return false
}

// SetS3Name gets a reference to the given NullableString and assigns it to the S3Name field.
func (o *BackupCredentialsListDto) SetS3Name(v string) {
	o.S3Name.Set(&v)
}

// SetS3NameNil sets the value for S3Name to be an explicit nil
func (o *BackupCredentialsListDto) SetS3NameNil() {
	o.S3Name.Set(nil)
}

// UnsetS3Name ensures that no value is present for S3Name, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetS3Name() {
	o.S3Name.Unset()
}

// GetS3AccessKeyId returns the S3AccessKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetS3AccessKeyId() string {
	if o == nil || IsNil(o.S3AccessKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.S3AccessKeyId.Get()
}

// GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetS3AccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3AccessKeyId.Get(), o.S3AccessKeyId.IsSet()
}

// HasS3AccessKeyId returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasS3AccessKeyId() bool {
	if o != nil && o.S3AccessKeyId.IsSet() {
		return true
	}

	return false
}

// SetS3AccessKeyId gets a reference to the given NullableString and assigns it to the S3AccessKeyId field.
func (o *BackupCredentialsListDto) SetS3AccessKeyId(v string) {
	o.S3AccessKeyId.Set(&v)
}

// SetS3AccessKeyIdNil sets the value for S3AccessKeyId to be an explicit nil
func (o *BackupCredentialsListDto) SetS3AccessKeyIdNil() {
	o.S3AccessKeyId.Set(nil)
}

// UnsetS3AccessKeyId ensures that no value is present for S3AccessKeyId, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetS3AccessKeyId() {
	o.S3AccessKeyId.Unset()
}

// GetS3Endpoint returns the S3Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetS3Endpoint() string {
	if o == nil || IsNil(o.S3Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.S3Endpoint.Get()
}

// GetS3EndpointOk returns a tuple with the S3Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetS3EndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Endpoint.Get(), o.S3Endpoint.IsSet()
}

// HasS3Endpoint returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasS3Endpoint() bool {
	if o != nil && o.S3Endpoint.IsSet() {
		return true
	}

	return false
}

// SetS3Endpoint gets a reference to the given NullableString and assigns it to the S3Endpoint field.
func (o *BackupCredentialsListDto) SetS3Endpoint(v string) {
	o.S3Endpoint.Set(&v)
}

// SetS3EndpointNil sets the value for S3Endpoint to be an explicit nil
func (o *BackupCredentialsListDto) SetS3EndpointNil() {
	o.S3Endpoint.Set(nil)
}

// UnsetS3Endpoint ensures that no value is present for S3Endpoint, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetS3Endpoint() {
	o.S3Endpoint.Unset()
}

// GetS3Region returns the S3Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetS3Region() string {
	if o == nil || IsNil(o.S3Region.Get()) {
		var ret string
		return ret
	}
	return *o.S3Region.Get()
}

// GetS3RegionOk returns a tuple with the S3Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetS3RegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Region.Get(), o.S3Region.IsSet()
}

// HasS3Region returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasS3Region() bool {
	if o != nil && o.S3Region.IsSet() {
		return true
	}

	return false
}

// SetS3Region gets a reference to the given NullableString and assigns it to the S3Region field.
func (o *BackupCredentialsListDto) SetS3Region(v string) {
	o.S3Region.Set(&v)
}

// SetS3RegionNil sets the value for S3Region to be an explicit nil
func (o *BackupCredentialsListDto) SetS3RegionNil() {
	o.S3Region.Set(nil)
}

// UnsetS3Region ensures that no value is present for S3Region, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetS3Region() {
	o.S3Region.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *BackupCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *BackupCredentialsListDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *BackupCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *BackupCredentialsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *BackupCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *BackupCredentialsListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *BackupCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *BackupCredentialsListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *BackupCredentialsListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *BackupCredentialsListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *BackupCredentialsListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *BackupCredentialsListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *BackupCredentialsListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentialsListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *BackupCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *BackupCredentialsListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *BackupCredentialsListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *BackupCredentialsListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *BackupCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsInfra returns the IsInfra field value if set, zero value otherwise.
func (o *BackupCredentialsListDto) GetIsInfra() bool {
	if o == nil || IsNil(o.IsInfra) {
		var ret bool
		return ret
	}
	return *o.IsInfra
}

// GetIsInfraOk returns a tuple with the IsInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentialsListDto) GetIsInfraOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInfra) {
		return nil, false
	}
	return o.IsInfra, true
}

// HasIsInfra returns a boolean if a field has been set.
func (o *BackupCredentialsListDto) HasIsInfra() bool {
	if o != nil && !IsNil(o.IsInfra) {
		return true
	}

	return false
}

// SetIsInfra gets a reference to the given bool and assigns it to the IsInfra field.
func (o *BackupCredentialsListDto) SetIsInfra(v bool) {
	o.IsInfra = &v
}

func (o BackupCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.S3Name.IsSet() {
		toSerialize["s3Name"] = o.S3Name.Get()
	}
	if o.S3AccessKeyId.IsSet() {
		toSerialize["s3AccessKeyId"] = o.S3AccessKeyId.Get()
	}
	if o.S3Endpoint.IsSet() {
		toSerialize["s3Endpoint"] = o.S3Endpoint.Get()
	}
	if o.S3Region.IsSet() {
		toSerialize["s3Region"] = o.S3Region.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.IsInfra) {
		toSerialize["isInfra"] = o.IsInfra
	}
	return toSerialize, nil
}

type NullableBackupCredentialsListDto struct {
	value *BackupCredentialsListDto
	isSet bool
}

func (v NullableBackupCredentialsListDto) Get() *BackupCredentialsListDto {
	return v.value
}

func (v *NullableBackupCredentialsListDto) Set(val *BackupCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupCredentialsListDto(val *BackupCredentialsListDto) *NullableBackupCredentialsListDto {
	return &NullableBackupCredentialsListDto{value: val, isSet: true}
}

func (v NullableBackupCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
