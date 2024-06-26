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

// checks if the InstanceAppListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAppListDto{}

// InstanceAppListDto struct for InstanceAppListDto
type InstanceAppListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Status *EInstanceStatus `json:"status,omitempty"`
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
	Logs NullableString `json:"logs,omitempty"`
}

// NewInstanceAppListDto instantiates a new InstanceAppListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAppListDto() *InstanceAppListDto {
	this := InstanceAppListDto{}
	return &this
}

// NewInstanceAppListDtoWithDefaults instantiates a new InstanceAppListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAppListDtoWithDefaults() *InstanceAppListDto {
	this := InstanceAppListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InstanceAppListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InstanceAppListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InstanceAppListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InstanceAppListDto) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *InstanceAppListDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *InstanceAppListDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *InstanceAppListDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetStatus() EInstanceStatus {
	if o == nil || IsNil(o.Status) {
		var ret EInstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetStatusOk() (*EInstanceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EInstanceStatus and assigns it to the Status field.
func (o *InstanceAppListDto) SetStatus(v EInstanceStatus) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *InstanceAppListDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *InstanceAppListDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *InstanceAppListDto) UnsetVersion() {
	o.Version.Unset()
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetCatalogId() int32 {
	if o == nil || IsNil(o.CatalogId) {
		var ret int32
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetCatalogIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.
func (o *InstanceAppListDto) SetCatalogId(v int32) {
	o.CatalogId = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetCatalogName() string {
	if o == nil || IsNil(o.CatalogName.Get()) {
		var ret string
		return ret
	}
	return *o.CatalogName.Get()
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetCatalogNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogName.Get(), o.CatalogName.IsSet()
}

// HasCatalogName returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCatalogName() bool {
	if o != nil && o.CatalogName.IsSet() {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given NullableString and assigns it to the CatalogName field.
func (o *InstanceAppListDto) SetCatalogName(v string) {
	o.CatalogName.Set(&v)
}
// SetCatalogNameNil sets the value for CatalogName to be an explicit nil
func (o *InstanceAppListDto) SetCatalogNameNil() {
	o.CatalogName.Set(nil)
}

// UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
func (o *InstanceAppListDto) UnsetCatalogName() {
	o.CatalogName.Unset()
}

// GetCatalogAppName returns the CatalogAppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetCatalogAppName() string {
	if o == nil || IsNil(o.CatalogAppName.Get()) {
		var ret string
		return ret
	}
	return *o.CatalogAppName.Get()
}

// GetCatalogAppNameOk returns a tuple with the CatalogAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetCatalogAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogAppName.Get(), o.CatalogAppName.IsSet()
}

// HasCatalogAppName returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCatalogAppName() bool {
	if o != nil && o.CatalogAppName.IsSet() {
		return true
	}

	return false
}

// SetCatalogAppName gets a reference to the given NullableString and assigns it to the CatalogAppName field.
func (o *InstanceAppListDto) SetCatalogAppName(v string) {
	o.CatalogAppName.Set(&v)
}
// SetCatalogAppNameNil sets the value for CatalogAppName to be an explicit nil
func (o *InstanceAppListDto) SetCatalogAppNameNil() {
	o.CatalogAppName.Set(nil)
}

// UnsetCatalogAppName ensures that no value is present for CatalogAppName, not even an explicit nil
func (o *InstanceAppListDto) UnsetCatalogAppName() {
	o.CatalogAppName.Unset()
}

// GetCatalogAppId returns the CatalogAppId field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetCatalogAppId() int32 {
	if o == nil || IsNil(o.CatalogAppId) {
		var ret int32
		return ret
	}
	return *o.CatalogAppId
}

// GetCatalogAppIdOk returns a tuple with the CatalogAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetCatalogAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogAppId) {
		return nil, false
	}
	return o.CatalogAppId, true
}

// HasCatalogAppId returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCatalogAppId() bool {
	if o != nil && !IsNil(o.CatalogAppId) {
		return true
	}

	return false
}

// SetCatalogAppId gets a reference to the given int32 and assigns it to the CatalogAppId field.
func (o *InstanceAppListDto) SetCatalogAppId(v int32) {
	o.CatalogAppId = &v
}

// GetAppRepoName returns the AppRepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetAppRepoName() string {
	if o == nil || IsNil(o.AppRepoName.Get()) {
		var ret string
		return ret
	}
	return *o.AppRepoName.Get()
}

// GetAppRepoNameOk returns a tuple with the AppRepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetAppRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppRepoName.Get(), o.AppRepoName.IsSet()
}

// HasAppRepoName returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasAppRepoName() bool {
	if o != nil && o.AppRepoName.IsSet() {
		return true
	}

	return false
}

// SetAppRepoName gets a reference to the given NullableString and assigns it to the AppRepoName field.
func (o *InstanceAppListDto) SetAppRepoName(v string) {
	o.AppRepoName.Set(&v)
}
// SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil
func (o *InstanceAppListDto) SetAppRepoNameNil() {
	o.AppRepoName.Set(nil)
}

// UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
func (o *InstanceAppListDto) UnsetAppRepoName() {
	o.AppRepoName.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *InstanceAppListDto) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *InstanceAppListDto) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *InstanceAppListDto) UnsetLogo() {
	o.Logo.Unset()
}

// GetAutoSync returns the AutoSync field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetAutoSync() bool {
	if o == nil || IsNil(o.AutoSync) {
		var ret bool
		return ret
	}
	return *o.AutoSync
}

// GetAutoSyncOk returns a tuple with the AutoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetAutoSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSync) {
		return nil, false
	}
	return o.AutoSync, true
}

// HasAutoSync returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasAutoSync() bool {
	if o != nil && !IsNil(o.AutoSync) {
		return true
	}

	return false
}

// SetAutoSync gets a reference to the given bool and assigns it to the AutoSync field.
func (o *InstanceAppListDto) SetAutoSync(v bool) {
	o.AutoSync = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetCreated() string {
	if o == nil || IsNil(o.Created.Get()) {
		var ret string
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableString and assigns it to the Created field.
func (o *InstanceAppListDto) SetCreated(v string) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *InstanceAppListDto) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *InstanceAppListDto) UnsetCreated() {
	o.Created.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *InstanceAppListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *InstanceAppListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *InstanceAppListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *InstanceAppListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *InstanceAppListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *InstanceAppListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *InstanceAppListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *InstanceAppListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *InstanceAppListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InstanceAppListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceAppListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *InstanceAppListDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *InstanceAppListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *InstanceAppListDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *InstanceAppListDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetLogs returns the Logs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAppListDto) GetLogs() string {
	if o == nil || IsNil(o.Logs.Get()) {
		var ret string
		return ret
	}
	return *o.Logs.Get()
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAppListDto) GetLogsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs.Get(), o.Logs.IsSet()
}

// HasLogs returns a boolean if a field has been set.
func (o *InstanceAppListDto) HasLogs() bool {
	if o != nil && o.Logs.IsSet() {
		return true
	}

	return false
}

// SetLogs gets a reference to the given NullableString and assigns it to the Logs field.
func (o *InstanceAppListDto) SetLogs(v string) {
	o.Logs.Set(&v)
}
// SetLogsNil sets the value for Logs to be an explicit nil
func (o *InstanceAppListDto) SetLogsNil() {
	o.Logs.Set(nil)
}

// UnsetLogs ensures that no value is present for Logs, not even an explicit nil
func (o *InstanceAppListDto) UnsetLogs() {
	o.Logs.Unset()
}

func (o InstanceAppListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAppListDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	if o.Logs.IsSet() {
		toSerialize["logs"] = o.Logs.Get()
	}
	return toSerialize, nil
}

type NullableInstanceAppListDto struct {
	value *InstanceAppListDto
	isSet bool
}

func (v NullableInstanceAppListDto) Get() *InstanceAppListDto {
	return v.value
}

func (v *NullableInstanceAppListDto) Set(val *InstanceAppListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAppListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAppListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAppListDto(val *InstanceAppListDto) *NullableInstanceAppListDto {
	return &NullableInstanceAppListDto{value: val, isSet: true}
}

func (v NullableInstanceAppListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAppListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


