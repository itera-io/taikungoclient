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

// checks if the KubeConfigForUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubeConfigForUserDto{}

// KubeConfigForUserDto struct for KubeConfigForUserDto
type KubeConfigForUserDto struct {
	Id                     *int32         `json:"id,omitempty"`
	UserId                 NullableString `json:"userId,omitempty"`
	DisplayName            NullableString `json:"displayName,omitempty"`
	ProjectId              *int32         `json:"projectId,omitempty"`
	OrganizationId         *int32         `json:"organizationId,omitempty"`
	ProjectName            NullableString `json:"projectName,omitempty"`
	IsAccessibleForAll     *bool          `json:"isAccessibleForAll,omitempty"`
	IsAccessibleForManager *bool          `json:"isAccessibleForManager,omitempty"`
	KubeConfigRoleName     NullableString `json:"kubeConfigRoleName,omitempty"`
	CreatedBy              NullableString `json:"createdBy,omitempty"`
	CreatedAt              NullableString `json:"createdAt,omitempty"`
	Namespace              NullableString `json:"namespace,omitempty"`
	ExpirationDate         NullableString `json:"expirationDate,omitempty"`
	CanDownload            *bool          `json:"canDownload,omitempty"`
	CanAccessTerminal      *bool          `json:"canAccessTerminal,omitempty"`
	CanDelete              *bool          `json:"canDelete,omitempty"`
}

// NewKubeConfigForUserDto instantiates a new KubeConfigForUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeConfigForUserDto() *KubeConfigForUserDto {
	this := KubeConfigForUserDto{}
	return &this
}

// NewKubeConfigForUserDtoWithDefaults instantiates a new KubeConfigForUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeConfigForUserDtoWithDefaults() *KubeConfigForUserDto {
	this := KubeConfigForUserDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubeConfigForUserDto) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *KubeConfigForUserDto) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *KubeConfigForUserDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *KubeConfigForUserDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *KubeConfigForUserDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubeConfigForUserDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *KubeConfigForUserDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *KubeConfigForUserDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *KubeConfigForUserDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetIsAccessibleForAll returns the IsAccessibleForAll field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetIsAccessibleForAll() bool {
	if o == nil || IsNil(o.IsAccessibleForAll) {
		var ret bool
		return ret
	}
	return *o.IsAccessibleForAll
}

// GetIsAccessibleForAllOk returns a tuple with the IsAccessibleForAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetIsAccessibleForAllOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccessibleForAll) {
		return nil, false
	}
	return o.IsAccessibleForAll, true
}

// HasIsAccessibleForAll returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasIsAccessibleForAll() bool {
	if o != nil && !IsNil(o.IsAccessibleForAll) {
		return true
	}

	return false
}

// SetIsAccessibleForAll gets a reference to the given bool and assigns it to the IsAccessibleForAll field.
func (o *KubeConfigForUserDto) SetIsAccessibleForAll(v bool) {
	o.IsAccessibleForAll = &v
}

// GetIsAccessibleForManager returns the IsAccessibleForManager field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetIsAccessibleForManager() bool {
	if o == nil || IsNil(o.IsAccessibleForManager) {
		var ret bool
		return ret
	}
	return *o.IsAccessibleForManager
}

// GetIsAccessibleForManagerOk returns a tuple with the IsAccessibleForManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetIsAccessibleForManagerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccessibleForManager) {
		return nil, false
	}
	return o.IsAccessibleForManager, true
}

// HasIsAccessibleForManager returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasIsAccessibleForManager() bool {
	if o != nil && !IsNil(o.IsAccessibleForManager) {
		return true
	}

	return false
}

// SetIsAccessibleForManager gets a reference to the given bool and assigns it to the IsAccessibleForManager field.
func (o *KubeConfigForUserDto) SetIsAccessibleForManager(v bool) {
	o.IsAccessibleForManager = &v
}

// GetKubeConfigRoleName returns the KubeConfigRoleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetKubeConfigRoleName() string {
	if o == nil || IsNil(o.KubeConfigRoleName.Get()) {
		var ret string
		return ret
	}
	return *o.KubeConfigRoleName.Get()
}

// GetKubeConfigRoleNameOk returns a tuple with the KubeConfigRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetKubeConfigRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubeConfigRoleName.Get(), o.KubeConfigRoleName.IsSet()
}

// HasKubeConfigRoleName returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasKubeConfigRoleName() bool {
	if o != nil && o.KubeConfigRoleName.IsSet() {
		return true
	}

	return false
}

// SetKubeConfigRoleName gets a reference to the given NullableString and assigns it to the KubeConfigRoleName field.
func (o *KubeConfigForUserDto) SetKubeConfigRoleName(v string) {
	o.KubeConfigRoleName.Set(&v)
}

// SetKubeConfigRoleNameNil sets the value for KubeConfigRoleName to be an explicit nil
func (o *KubeConfigForUserDto) SetKubeConfigRoleNameNil() {
	o.KubeConfigRoleName.Set(nil)
}

// UnsetKubeConfigRoleName ensures that no value is present for KubeConfigRoleName, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetKubeConfigRoleName() {
	o.KubeConfigRoleName.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *KubeConfigForUserDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *KubeConfigForUserDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *KubeConfigForUserDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *KubeConfigForUserDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *KubeConfigForUserDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *KubeConfigForUserDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigForUserDto) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret string
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigForUserDto) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableString and assigns it to the ExpirationDate field.
func (o *KubeConfigForUserDto) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}

// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *KubeConfigForUserDto) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *KubeConfigForUserDto) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetCanDownload returns the CanDownload field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetCanDownload() bool {
	if o == nil || IsNil(o.CanDownload) {
		var ret bool
		return ret
	}
	return *o.CanDownload
}

// GetCanDownloadOk returns a tuple with the CanDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetCanDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDownload) {
		return nil, false
	}
	return o.CanDownload, true
}

// HasCanDownload returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasCanDownload() bool {
	if o != nil && !IsNil(o.CanDownload) {
		return true
	}

	return false
}

// SetCanDownload gets a reference to the given bool and assigns it to the CanDownload field.
func (o *KubeConfigForUserDto) SetCanDownload(v bool) {
	o.CanDownload = &v
}

// GetCanAccessTerminal returns the CanAccessTerminal field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetCanAccessTerminal() bool {
	if o == nil || IsNil(o.CanAccessTerminal) {
		var ret bool
		return ret
	}
	return *o.CanAccessTerminal
}

// GetCanAccessTerminalOk returns a tuple with the CanAccessTerminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetCanAccessTerminalOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAccessTerminal) {
		return nil, false
	}
	return o.CanAccessTerminal, true
}

// HasCanAccessTerminal returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasCanAccessTerminal() bool {
	if o != nil && !IsNil(o.CanAccessTerminal) {
		return true
	}

	return false
}

// SetCanAccessTerminal gets a reference to the given bool and assigns it to the CanAccessTerminal field.
func (o *KubeConfigForUserDto) SetCanAccessTerminal(v bool) {
	o.CanAccessTerminal = &v
}

// GetCanDelete returns the CanDelete field value if set, zero value otherwise.
func (o *KubeConfigForUserDto) GetCanDelete() bool {
	if o == nil || IsNil(o.CanDelete) {
		var ret bool
		return ret
	}
	return *o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigForUserDto) GetCanDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDelete) {
		return nil, false
	}
	return o.CanDelete, true
}

// HasCanDelete returns a boolean if a field has been set.
func (o *KubeConfigForUserDto) HasCanDelete() bool {
	if o != nil && !IsNil(o.CanDelete) {
		return true
	}

	return false
}

// SetCanDelete gets a reference to the given bool and assigns it to the CanDelete field.
func (o *KubeConfigForUserDto) SetCanDelete(v bool) {
	o.CanDelete = &v
}

func (o KubeConfigForUserDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubeConfigForUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.IsAccessibleForAll) {
		toSerialize["isAccessibleForAll"] = o.IsAccessibleForAll
	}
	if !IsNil(o.IsAccessibleForManager) {
		toSerialize["isAccessibleForManager"] = o.IsAccessibleForManager
	}
	if o.KubeConfigRoleName.IsSet() {
		toSerialize["kubeConfigRoleName"] = o.KubeConfigRoleName.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expirationDate"] = o.ExpirationDate.Get()
	}
	if !IsNil(o.CanDownload) {
		toSerialize["canDownload"] = o.CanDownload
	}
	if !IsNil(o.CanAccessTerminal) {
		toSerialize["canAccessTerminal"] = o.CanAccessTerminal
	}
	if !IsNil(o.CanDelete) {
		toSerialize["canDelete"] = o.CanDelete
	}
	return toSerialize, nil
}

type NullableKubeConfigForUserDto struct {
	value *KubeConfigForUserDto
	isSet bool
}

func (v NullableKubeConfigForUserDto) Get() *KubeConfigForUserDto {
	return v.value
}

func (v *NullableKubeConfigForUserDto) Set(val *KubeConfigForUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeConfigForUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeConfigForUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeConfigForUserDto(val *KubeConfigForUserDto) *NullableKubeConfigForUserDto {
	return &NullableKubeConfigForUserDto{value: val, isSet: true}
}

func (v NullableKubeConfigForUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeConfigForUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
