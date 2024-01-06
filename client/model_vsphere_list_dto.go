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

// checks if the VsphereListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsphereListDto{}

// VsphereListDto struct for VsphereListDto
type VsphereListDto struct {
	Id *int32 `json:"id,omitempty"`
	ProjectCount *int32 `json:"projectCount,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	DrsEnabled *bool `json:"drsEnabled,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	ContinentName NullableString `json:"continentName,omitempty"`
	Hypervisors []CommonStringBasedDropdownDto `json:"hypervisors,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Password NullableString `json:"password,omitempty"`
	DatacenterId NullableString `json:"datacenterId,omitempty"`
	DatacenterName NullableString `json:"datacenterName,omitempty"`
	Datastore NullableString `json:"datastore,omitempty"`
	VmTemplateName NullableString `json:"vmTemplateName,omitempty"`
	VsphereNetworks []VsphereNetworkListDto `json:"vsphereNetworks,omitempty"`
}

// NewVsphereListDto instantiates a new VsphereListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereListDto() *VsphereListDto {
	this := VsphereListDto{}
	return &this
}

// NewVsphereListDtoWithDefaults instantiates a new VsphereListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereListDtoWithDefaults() *VsphereListDto {
	this := VsphereListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VsphereListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VsphereListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VsphereListDto) SetId(v int32) {
	o.Id = &v
}

// GetProjectCount returns the ProjectCount field value if set, zero value otherwise.
func (o *VsphereListDto) GetProjectCount() int32 {
	if o == nil || IsNil(o.ProjectCount) {
		var ret int32
		return ret
	}
	return *o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectCount) {
		return nil, false
	}
	return o.ProjectCount, true
}

// HasProjectCount returns a boolean if a field has been set.
func (o *VsphereListDto) HasProjectCount() bool {
	if o != nil && !IsNil(o.ProjectCount) {
		return true
	}

	return false
}

// SetProjectCount gets a reference to the given int32 and assigns it to the ProjectCount field.
func (o *VsphereListDto) SetProjectCount(v int32) {
	o.ProjectCount = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *VsphereListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *VsphereListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *VsphereListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VsphereListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VsphereListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VsphereListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VsphereListDto) UnsetName() {
	o.Name.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *VsphereListDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *VsphereListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VsphereListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *VsphereListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *VsphereListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *VsphereListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VsphereListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *VsphereListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *VsphereListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *VsphereListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *VsphereListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *VsphereListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *VsphereListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *VsphereListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *VsphereListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *VsphereListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *VsphereListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *VsphereListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *VsphereListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *VsphereListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *VsphereListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetDrsEnabled returns the DrsEnabled field value if set, zero value otherwise.
func (o *VsphereListDto) GetDrsEnabled() bool {
	if o == nil || IsNil(o.DrsEnabled) {
		var ret bool
		return ret
	}
	return *o.DrsEnabled
}

// GetDrsEnabledOk returns a tuple with the DrsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetDrsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DrsEnabled) {
		return nil, false
	}
	return o.DrsEnabled, true
}

// HasDrsEnabled returns a boolean if a field has been set.
func (o *VsphereListDto) HasDrsEnabled() bool {
	if o != nil && !IsNil(o.DrsEnabled) {
		return true
	}

	return false
}

// SetDrsEnabled gets a reference to the given bool and assigns it to the DrsEnabled field.
func (o *VsphereListDto) SetDrsEnabled(v bool) {
	o.DrsEnabled = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *VsphereListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *VsphereListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *VsphereListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *VsphereListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *VsphereListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *VsphereListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *VsphereListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *VsphereListDto) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *VsphereListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}
// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *VsphereListDto) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *VsphereListDto) UnsetContinentName() {
	o.ContinentName.Unset()
}

// GetHypervisors returns the Hypervisors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetHypervisors() []CommonStringBasedDropdownDto {
	if o == nil {
		var ret []CommonStringBasedDropdownDto
		return ret
	}
	return o.Hypervisors
}

// GetHypervisorsOk returns a tuple with the Hypervisors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetHypervisorsOk() ([]CommonStringBasedDropdownDto, bool) {
	if o == nil || IsNil(o.Hypervisors) {
		return nil, false
	}
	return o.Hypervisors, true
}

// HasHypervisors returns a boolean if a field has been set.
func (o *VsphereListDto) HasHypervisors() bool {
	if o != nil && IsNil(o.Hypervisors) {
		return true
	}

	return false
}

// SetHypervisors gets a reference to the given []CommonStringBasedDropdownDto and assigns it to the Hypervisors field.
func (o *VsphereListDto) SetHypervisors(v []CommonStringBasedDropdownDto) {
	o.Hypervisors = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *VsphereListDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *VsphereListDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *VsphereListDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *VsphereListDto) UnsetUsername() {
	o.Username.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *VsphereListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *VsphereListDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *VsphereListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *VsphereListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *VsphereListDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *VsphereListDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *VsphereListDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *VsphereListDto) UnsetPassword() {
	o.Password.Unset()
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetDatacenterId() string {
	if o == nil || IsNil(o.DatacenterId.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterId.Get()
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterId.Get(), o.DatacenterId.IsSet()
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *VsphereListDto) HasDatacenterId() bool {
	if o != nil && o.DatacenterId.IsSet() {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given NullableString and assigns it to the DatacenterId field.
func (o *VsphereListDto) SetDatacenterId(v string) {
	o.DatacenterId.Set(&v)
}
// SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil
func (o *VsphereListDto) SetDatacenterIdNil() {
	o.DatacenterId.Set(nil)
}

// UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
func (o *VsphereListDto) UnsetDatacenterId() {
	o.DatacenterId.Unset()
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetDatacenterName() string {
	if o == nil || IsNil(o.DatacenterName.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterName.Get()
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterName.Get(), o.DatacenterName.IsSet()
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *VsphereListDto) HasDatacenterName() bool {
	if o != nil && o.DatacenterName.IsSet() {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given NullableString and assigns it to the DatacenterName field.
func (o *VsphereListDto) SetDatacenterName(v string) {
	o.DatacenterName.Set(&v)
}
// SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil
func (o *VsphereListDto) SetDatacenterNameNil() {
	o.DatacenterName.Set(nil)
}

// UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
func (o *VsphereListDto) UnsetDatacenterName() {
	o.DatacenterName.Unset()
}

// GetDatastore returns the Datastore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetDatastore() string {
	if o == nil || IsNil(o.Datastore.Get()) {
		var ret string
		return ret
	}
	return *o.Datastore.Get()
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetDatastoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datastore.Get(), o.Datastore.IsSet()
}

// HasDatastore returns a boolean if a field has been set.
func (o *VsphereListDto) HasDatastore() bool {
	if o != nil && o.Datastore.IsSet() {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given NullableString and assigns it to the Datastore field.
func (o *VsphereListDto) SetDatastore(v string) {
	o.Datastore.Set(&v)
}
// SetDatastoreNil sets the value for Datastore to be an explicit nil
func (o *VsphereListDto) SetDatastoreNil() {
	o.Datastore.Set(nil)
}

// UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
func (o *VsphereListDto) UnsetDatastore() {
	o.Datastore.Unset()
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetVmTemplateName() string {
	if o == nil || IsNil(o.VmTemplateName.Get()) {
		var ret string
		return ret
	}
	return *o.VmTemplateName.Get()
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetVmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTemplateName.Get(), o.VmTemplateName.IsSet()
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *VsphereListDto) HasVmTemplateName() bool {
	if o != nil && o.VmTemplateName.IsSet() {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given NullableString and assigns it to the VmTemplateName field.
func (o *VsphereListDto) SetVmTemplateName(v string) {
	o.VmTemplateName.Set(&v)
}
// SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil
func (o *VsphereListDto) SetVmTemplateNameNil() {
	o.VmTemplateName.Set(nil)
}

// UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
func (o *VsphereListDto) UnsetVmTemplateName() {
	o.VmTemplateName.Unset()
}

// GetVsphereNetworks returns the VsphereNetworks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereListDto) GetVsphereNetworks() []VsphereNetworkListDto {
	if o == nil {
		var ret []VsphereNetworkListDto
		return ret
	}
	return o.VsphereNetworks
}

// GetVsphereNetworksOk returns a tuple with the VsphereNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereListDto) GetVsphereNetworksOk() ([]VsphereNetworkListDto, bool) {
	if o == nil || IsNil(o.VsphereNetworks) {
		return nil, false
	}
	return o.VsphereNetworks, true
}

// HasVsphereNetworks returns a boolean if a field has been set.
func (o *VsphereListDto) HasVsphereNetworks() bool {
	if o != nil && IsNil(o.VsphereNetworks) {
		return true
	}

	return false
}

// SetVsphereNetworks gets a reference to the given []VsphereNetworkListDto and assigns it to the VsphereNetworks field.
func (o *VsphereListDto) SetVsphereNetworks(v []VsphereNetworkListDto) {
	o.VsphereNetworks = v
}

func (o VsphereListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsphereListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectCount) {
		toSerialize["projectCount"] = o.ProjectCount
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.DrsEnabled) {
		toSerialize["drsEnabled"] = o.DrsEnabled
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.ContinentName.IsSet() {
		toSerialize["continentName"] = o.ContinentName.Get()
	}
	if o.Hypervisors != nil {
		toSerialize["hypervisors"] = o.Hypervisors
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.DatacenterId.IsSet() {
		toSerialize["datacenterId"] = o.DatacenterId.Get()
	}
	if o.DatacenterName.IsSet() {
		toSerialize["datacenterName"] = o.DatacenterName.Get()
	}
	if o.Datastore.IsSet() {
		toSerialize["datastore"] = o.Datastore.Get()
	}
	if o.VmTemplateName.IsSet() {
		toSerialize["vmTemplateName"] = o.VmTemplateName.Get()
	}
	if o.VsphereNetworks != nil {
		toSerialize["vsphereNetworks"] = o.VsphereNetworks
	}
	return toSerialize, nil
}

type NullableVsphereListDto struct {
	value *VsphereListDto
	isSet bool
}

func (v NullableVsphereListDto) Get() *VsphereListDto {
	return v.value
}

func (v *NullableVsphereListDto) Set(val *VsphereListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVsphereListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVsphereListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsphereListDto(val *VsphereListDto) *NullableVsphereListDto {
	return &NullableVsphereListDto{value: val, isSet: true}
}

func (v NullableVsphereListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsphereListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


