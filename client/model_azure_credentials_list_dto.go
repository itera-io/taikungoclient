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

// checks if the AzureCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCredentialsListDto{}

// AzureCredentialsListDto struct for AzureCredentialsListDto
type AzureCredentialsListDto struct {
	Id *int32 `json:"id,omitempty"`
	ProjectCount *int32 `json:"projectCount,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Location NullableString `json:"location,omitempty"`
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	AvailabilityZonesCount *int32 `json:"availabilityZonesCount,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	ContinentName NullableString `json:"continentName,omitempty"`
}

// NewAzureCredentialsListDto instantiates a new AzureCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCredentialsListDto() *AzureCredentialsListDto {
	this := AzureCredentialsListDto{}
	return &this
}

// NewAzureCredentialsListDtoWithDefaults instantiates a new AzureCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCredentialsListDtoWithDefaults() *AzureCredentialsListDto {
	this := AzureCredentialsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AzureCredentialsListDto) SetId(v int32) {
	o.Id = &v
}

// GetProjectCount returns the ProjectCount field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetProjectCount() int32 {
	if o == nil || IsNil(o.ProjectCount) {
		var ret int32
		return ret
	}
	return *o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectCount) {
		return nil, false
	}
	return o.ProjectCount, true
}

// HasProjectCount returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasProjectCount() bool {
	if o != nil && !IsNil(o.ProjectCount) {
		return true
	}

	return false
}

// SetProjectCount gets a reference to the given int32 and assigns it to the ProjectCount field.
func (o *AzureCredentialsListDto) SetProjectCount(v int32) {
	o.ProjectCount = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *AzureCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AzureCredentialsListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AzureCredentialsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetName() {
	o.Name.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *AzureCredentialsListDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *AzureCredentialsListDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *AzureCredentialsListDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *AzureCredentialsListDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetLocation() {
	o.Location.Unset()
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetAvailabilityZones() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetAvailabilityZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailabilityZones) {
		return nil, false
	}
	return o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasAvailabilityZones() bool {
	if o != nil && !IsNil(o.AvailabilityZones) {
		return true
	}

	return false
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *AzureCredentialsListDto) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetAvailabilityZonesCount returns the AvailabilityZonesCount field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetAvailabilityZonesCount() int32 {
	if o == nil || IsNil(o.AvailabilityZonesCount) {
		var ret int32
		return ret
	}
	return *o.AvailabilityZonesCount
}

// GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailabilityZonesCount) {
		return nil, false
	}
	return o.AvailabilityZonesCount, true
}

// HasAvailabilityZonesCount returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasAvailabilityZonesCount() bool {
	if o != nil && !IsNil(o.AvailabilityZonesCount) {
		return true
	}

	return false
}

// SetAvailabilityZonesCount gets a reference to the given int32 and assigns it to the AvailabilityZonesCount field.
func (o *AzureCredentialsListDto) SetAvailabilityZonesCount(v int32) {
	o.AvailabilityZonesCount = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *AzureCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *AzureCredentialsListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *AzureCredentialsListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *AzureCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *AzureCredentialsListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *AzureCredentialsListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *AzureCredentialsListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *AzureCredentialsListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *AzureCredentialsListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *AzureCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *AzureCredentialsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *AzureCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *AzureCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *AzureCredentialsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetContinentName returns the ContinentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureCredentialsListDto) GetContinentName() string {
	if o == nil || IsNil(o.ContinentName.Get()) {
		var ret string
		return ret
	}
	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureCredentialsListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// HasContinentName returns a boolean if a field has been set.
func (o *AzureCredentialsListDto) HasContinentName() bool {
	if o != nil && o.ContinentName.IsSet() {
		return true
	}

	return false
}

// SetContinentName gets a reference to the given NullableString and assigns it to the ContinentName field.
func (o *AzureCredentialsListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}
// SetContinentNameNil sets the value for ContinentName to be an explicit nil
func (o *AzureCredentialsListDto) SetContinentNameNil() {
	o.ContinentName.Set(nil)
}

// UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
func (o *AzureCredentialsListDto) UnsetContinentName() {
	o.ContinentName.Unset()
}

func (o AzureCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCredentialsListDto) ToMap() (map[string]interface{}, error) {
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
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.AvailabilityZones != nil {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if !IsNil(o.AvailabilityZonesCount) {
		toSerialize["availabilityZonesCount"] = o.AvailabilityZonesCount
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
	return toSerialize, nil
}

type NullableAzureCredentialsListDto struct {
	value *AzureCredentialsListDto
	isSet bool
}

func (v NullableAzureCredentialsListDto) Get() *AzureCredentialsListDto {
	return v.value
}

func (v *NullableAzureCredentialsListDto) Set(val *AzureCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCredentialsListDto(val *AzureCredentialsListDto) *NullableAzureCredentialsListDto {
	return &NullableAzureCredentialsListDto{value: val, isSet: true}
}

func (v NullableAzureCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


