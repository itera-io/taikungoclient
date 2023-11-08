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

// checks if the ProjectAppListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAppListDto{}

// ProjectAppListDto struct for ProjectAppListDto
type ProjectAppListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Version NullableString `json:"version,omitempty"`
	CatalogId *int32 `json:"catalogId,omitempty"`
	CatalogName NullableString `json:"catalogName,omitempty"`
	CatalogAppName NullableString `json:"catalogAppName,omitempty"`
	CatalogAppId *int32 `json:"catalogAppId,omitempty"`
	AppRepoName NullableString `json:"appRepoName,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
	AutoSync *bool `json:"autoSync,omitempty"`
	Created NullableString `json:"created,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
}

// NewProjectAppListDto instantiates a new ProjectAppListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAppListDto() *ProjectAppListDto {
	this := ProjectAppListDto{}
	return &this
}

// NewProjectAppListDtoWithDefaults instantiates a new ProjectAppListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAppListDtoWithDefaults() *ProjectAppListDto {
	this := ProjectAppListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectAppListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectAppListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectAppListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectAppListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectAppListDto) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *ProjectAppListDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *ProjectAppListDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *ProjectAppListDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ProjectAppListDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ProjectAppListDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ProjectAppListDto) UnsetStatus() {
	o.Status.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ProjectAppListDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ProjectAppListDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ProjectAppListDto) UnsetVersion() {
	o.Version.Unset()
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *ProjectAppListDto) GetCatalogId() int32 {
	if o == nil || IsNil(o.CatalogId) {
		var ret int32
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppListDto) GetCatalogIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.
func (o *ProjectAppListDto) SetCatalogId(v int32) {
	o.CatalogId = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetCatalogName() string {
	if o == nil || IsNil(o.CatalogName.Get()) {
		var ret string
		return ret
	}
	return *o.CatalogName.Get()
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetCatalogNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogName.Get(), o.CatalogName.IsSet()
}

// HasCatalogName returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCatalogName() bool {
	if o != nil && o.CatalogName.IsSet() {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given NullableString and assigns it to the CatalogName field.
func (o *ProjectAppListDto) SetCatalogName(v string) {
	o.CatalogName.Set(&v)
}
// SetCatalogNameNil sets the value for CatalogName to be an explicit nil
func (o *ProjectAppListDto) SetCatalogNameNil() {
	o.CatalogName.Set(nil)
}

// UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
func (o *ProjectAppListDto) UnsetCatalogName() {
	o.CatalogName.Unset()
}

// GetCatalogAppName returns the CatalogAppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetCatalogAppName() string {
	if o == nil || IsNil(o.CatalogAppName.Get()) {
		var ret string
		return ret
	}
	return *o.CatalogAppName.Get()
}

// GetCatalogAppNameOk returns a tuple with the CatalogAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetCatalogAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogAppName.Get(), o.CatalogAppName.IsSet()
}

// HasCatalogAppName returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCatalogAppName() bool {
	if o != nil && o.CatalogAppName.IsSet() {
		return true
	}

	return false
}

// SetCatalogAppName gets a reference to the given NullableString and assigns it to the CatalogAppName field.
func (o *ProjectAppListDto) SetCatalogAppName(v string) {
	o.CatalogAppName.Set(&v)
}
// SetCatalogAppNameNil sets the value for CatalogAppName to be an explicit nil
func (o *ProjectAppListDto) SetCatalogAppNameNil() {
	o.CatalogAppName.Set(nil)
}

// UnsetCatalogAppName ensures that no value is present for CatalogAppName, not even an explicit nil
func (o *ProjectAppListDto) UnsetCatalogAppName() {
	o.CatalogAppName.Unset()
}

// GetCatalogAppId returns the CatalogAppId field value if set, zero value otherwise.
func (o *ProjectAppListDto) GetCatalogAppId() int32 {
	if o == nil || IsNil(o.CatalogAppId) {
		var ret int32
		return ret
	}
	return *o.CatalogAppId
}

// GetCatalogAppIdOk returns a tuple with the CatalogAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppListDto) GetCatalogAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogAppId) {
		return nil, false
	}
	return o.CatalogAppId, true
}

// HasCatalogAppId returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCatalogAppId() bool {
	if o != nil && !IsNil(o.CatalogAppId) {
		return true
	}

	return false
}

// SetCatalogAppId gets a reference to the given int32 and assigns it to the CatalogAppId field.
func (o *ProjectAppListDto) SetCatalogAppId(v int32) {
	o.CatalogAppId = &v
}

// GetAppRepoName returns the AppRepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetAppRepoName() string {
	if o == nil || IsNil(o.AppRepoName.Get()) {
		var ret string
		return ret
	}
	return *o.AppRepoName.Get()
}

// GetAppRepoNameOk returns a tuple with the AppRepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetAppRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppRepoName.Get(), o.AppRepoName.IsSet()
}

// HasAppRepoName returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasAppRepoName() bool {
	if o != nil && o.AppRepoName.IsSet() {
		return true
	}

	return false
}

// SetAppRepoName gets a reference to the given NullableString and assigns it to the AppRepoName field.
func (o *ProjectAppListDto) SetAppRepoName(v string) {
	o.AppRepoName.Set(&v)
}
// SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil
func (o *ProjectAppListDto) SetAppRepoNameNil() {
	o.AppRepoName.Set(nil)
}

// UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
func (o *ProjectAppListDto) UnsetAppRepoName() {
	o.AppRepoName.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *ProjectAppListDto) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *ProjectAppListDto) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *ProjectAppListDto) UnsetLogo() {
	o.Logo.Unset()
}

// GetAutoSync returns the AutoSync field value if set, zero value otherwise.
func (o *ProjectAppListDto) GetAutoSync() bool {
	if o == nil || IsNil(o.AutoSync) {
		var ret bool
		return ret
	}
	return *o.AutoSync
}

// GetAutoSyncOk returns a tuple with the AutoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppListDto) GetAutoSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSync) {
		return nil, false
	}
	return o.AutoSync, true
}

// HasAutoSync returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasAutoSync() bool {
	if o != nil && !IsNil(o.AutoSync) {
		return true
	}

	return false
}

// SetAutoSync gets a reference to the given bool and assigns it to the AutoSync field.
func (o *ProjectAppListDto) SetAutoSync(v bool) {
	o.AutoSync = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetCreated() string {
	if o == nil || IsNil(o.Created.Get()) {
		var ret string
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableString and assigns it to the Created field.
func (o *ProjectAppListDto) SetCreated(v string) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *ProjectAppListDto) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *ProjectAppListDto) UnsetCreated() {
	o.Created.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProjectAppListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProjectAppListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProjectAppListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *ProjectAppListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *ProjectAppListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *ProjectAppListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *ProjectAppListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *ProjectAppListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *ProjectAppListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectAppListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ProjectAppListDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectAppListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectAppListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *ProjectAppListDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *ProjectAppListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *ProjectAppListDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *ProjectAppListDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

func (o ProjectAppListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAppListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	if o.CatalogName.IsSet() {
		toSerialize["catalogName"] = o.CatalogName.Get()
	}
	if o.CatalogAppName.IsSet() {
		toSerialize["catalogAppName"] = o.CatalogAppName.Get()
	}
	if !IsNil(o.CatalogAppId) {
		toSerialize["catalogAppId"] = o.CatalogAppId
	}
	if o.AppRepoName.IsSet() {
		toSerialize["appRepoName"] = o.AppRepoName.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if !IsNil(o.AutoSync) {
		toSerialize["autoSync"] = o.AutoSync
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
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
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	return toSerialize, nil
}

type NullableProjectAppListDto struct {
	value *ProjectAppListDto
	isSet bool
}

func (v NullableProjectAppListDto) Get() *ProjectAppListDto {
	return v.value
}

func (v *NullableProjectAppListDto) Set(val *ProjectAppListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAppListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAppListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAppListDto(val *ProjectAppListDto) *NullableProjectAppListDto {
	return &NullableProjectAppListDto{value: val, isSet: true}
}

func (v NullableProjectAppListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAppListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


