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

// checks if the ProxmoxListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxListDto{}

// ProxmoxListDto struct for ProxmoxListDto
type ProxmoxListDto struct {
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
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	ContinentName NullableString `json:"continentName,omitempty"`
	Hypervisors []CommonDropdownDto `json:"hypervisors,omitempty"`
	TokenId NullableString `json:"tokenId,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Storage NullableString `json:"storage,omitempty"`
	VmTemplateName NullableString `json:"vmTemplateName,omitempty"`
	ProxmoxNetworks []ProxmoxNetworkListDto `json:"proxmoxNetworks,omitempty"`
	SkipTlsFlag *bool `json:"skipTlsFlag,omitempty"`
}

// NewProxmoxListDto instantiates a new ProxmoxListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxListDto() *ProxmoxListDto {
	this := ProxmoxListDto{}
	return &this
}

// NewProxmoxListDtoWithDefaults instantiates a new ProxmoxListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxListDtoWithDefaults() *ProxmoxListDto {
	this := ProxmoxListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProxmoxListDto) SetId(v int32) {
	o.Id = &v
}

// GetProjectCount returns the ProjectCount field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetProjectCount() int32 {
	if o == nil || IsNil(o.ProjectCount) {
		var ret int32
		return ret
	}
	return *o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectCount) {
		return nil, false
	}
	return o.ProjectCount, true
}

// HasProjectCount returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasProjectCount() bool {
	if o != nil && !IsNil(o.ProjectCount) {
		return true
	}

	return false
}

// SetProjectCount gets a reference to the given int32 and assigns it to the ProjectCount field.
func (o *ProxmoxListDto) SetProjectCount(v int32) {
	o.ProjectCount = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *ProxmoxListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProxmoxListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProxmoxListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProxmoxListDto) UnsetName() {
	o.Name.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *ProxmoxListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProxmoxListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProxmoxListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProxmoxListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *ProxmoxListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ProxmoxListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ProxmoxListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *ProxmoxListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *ProxmoxListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *ProxmoxListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *ProxmoxListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *ProxmoxListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *ProxmoxListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ProxmoxListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *ProxmoxListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *ProxmoxListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *ProxmoxListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *ProxmoxListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *ProxmoxListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}
// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *ProxmoxListDto) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *ProxmoxListDto) UnsetContinentName() {
	o.ContinentName.Unset()
}

// GetHypervisors returns the Hypervisors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetHypervisors() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Hypervisors
}

// GetHypervisorsOk returns a tuple with the Hypervisors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetHypervisorsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Hypervisors) {
		return nil, false
	}
	return o.Hypervisors, true
}

// HasHypervisors returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasHypervisors() bool {
	if o != nil && !IsNil(o.Hypervisors) {
		return true
	}

	return false
}

// SetHypervisors gets a reference to the given []CommonDropdownDto and assigns it to the Hypervisors field.
func (o *ProxmoxListDto) SetHypervisors(v []CommonDropdownDto) {
	o.Hypervisors = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *ProxmoxListDto) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *ProxmoxListDto) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *ProxmoxListDto) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ProxmoxListDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ProxmoxListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ProxmoxListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetStorage() string {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret string
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableString and assigns it to the Storage field.
func (o *ProxmoxListDto) SetStorage(v string) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *ProxmoxListDto) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *ProxmoxListDto) UnsetStorage() {
	o.Storage.Unset()
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetVmTemplateName() string {
	if o == nil || IsNil(o.VmTemplateName.Get()) {
		var ret string
		return ret
	}
	return *o.VmTemplateName.Get()
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetVmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTemplateName.Get(), o.VmTemplateName.IsSet()
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasVmTemplateName() bool {
	if o != nil && o.VmTemplateName.IsSet() {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given NullableString and assigns it to the VmTemplateName field.
func (o *ProxmoxListDto) SetVmTemplateName(v string) {
	o.VmTemplateName.Set(&v)
}
// SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil
func (o *ProxmoxListDto) SetVmTemplateNameNil() {
	o.VmTemplateName.Set(nil)
}

// UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
func (o *ProxmoxListDto) UnsetVmTemplateName() {
	o.VmTemplateName.Unset()
}

// GetProxmoxNetworks returns the ProxmoxNetworks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxListDto) GetProxmoxNetworks() []ProxmoxNetworkListDto {
	if o == nil {
		var ret []ProxmoxNetworkListDto
		return ret
	}
	return o.ProxmoxNetworks
}

// GetProxmoxNetworksOk returns a tuple with the ProxmoxNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxListDto) GetProxmoxNetworksOk() ([]ProxmoxNetworkListDto, bool) {
	if o == nil || IsNil(o.ProxmoxNetworks) {
		return nil, false
	}
	return o.ProxmoxNetworks, true
}

// HasProxmoxNetworks returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasProxmoxNetworks() bool {
	if o != nil && !IsNil(o.ProxmoxNetworks) {
		return true
	}

	return false
}

// SetProxmoxNetworks gets a reference to the given []ProxmoxNetworkListDto and assigns it to the ProxmoxNetworks field.
func (o *ProxmoxListDto) SetProxmoxNetworks(v []ProxmoxNetworkListDto) {
	o.ProxmoxNetworks = v
}

// GetSkipTlsFlag returns the SkipTlsFlag field value if set, zero value otherwise.
func (o *ProxmoxListDto) GetSkipTlsFlag() bool {
	if o == nil || IsNil(o.SkipTlsFlag) {
		var ret bool
		return ret
	}
	return *o.SkipTlsFlag
}

// GetSkipTlsFlagOk returns a tuple with the SkipTlsFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxListDto) GetSkipTlsFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipTlsFlag) {
		return nil, false
	}
	return o.SkipTlsFlag, true
}

// HasSkipTlsFlag returns a boolean if a field has been set.
func (o *ProxmoxListDto) HasSkipTlsFlag() bool {
	if o != nil && !IsNil(o.SkipTlsFlag) {
		return true
	}

	return false
}

// SetSkipTlsFlag gets a reference to the given bool and assigns it to the SkipTlsFlag field.
func (o *ProxmoxListDto) SetSkipTlsFlag(v bool) {
	o.SkipTlsFlag = &v
}

func (o ProxmoxListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxListDto) ToMap() (map[string]interface{}, error) {
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
	if o.TokenId.IsSet() {
		toSerialize["tokenId"] = o.TokenId.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.VmTemplateName.IsSet() {
		toSerialize["vmTemplateName"] = o.VmTemplateName.Get()
	}
	if o.ProxmoxNetworks != nil {
		toSerialize["proxmoxNetworks"] = o.ProxmoxNetworks
	}
	if !IsNil(o.SkipTlsFlag) {
		toSerialize["skipTlsFlag"] = o.SkipTlsFlag
	}
	return toSerialize, nil
}

type NullableProxmoxListDto struct {
	value *ProxmoxListDto
	isSet bool
}

func (v NullableProxmoxListDto) Get() *ProxmoxListDto {
	return v.value
}

func (v *NullableProxmoxListDto) Set(val *ProxmoxListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxListDto(val *ProxmoxListDto) *NullableProxmoxListDto {
	return &NullableProxmoxListDto{value: val, isSet: true}
}

func (v NullableProxmoxListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


