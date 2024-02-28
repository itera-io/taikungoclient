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

// checks if the CatalogListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogListDto{}

// CatalogListDto struct for CatalogListDto
type CatalogListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	PackageIds []string `json:"packageIds,omitempty"`
	BoundProjects []ProjectCatalogDto `json:"boundProjects,omitempty"`
	BoundApplications []AvailablePackagesDto `json:"boundApplications,omitempty"`
}

// NewCatalogListDto instantiates a new CatalogListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogListDto() *CatalogListDto {
	this := CatalogListDto{}
	return &this
}

// NewCatalogListDtoWithDefaults instantiates a new CatalogListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogListDtoWithDefaults() *CatalogListDto {
	this := CatalogListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CatalogListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CatalogListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CatalogListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CatalogListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CatalogListDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogListDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogListDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogListDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CatalogListDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CatalogListDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CatalogListDto) UnsetDescription() {
	o.Description.Unset()
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *CatalogListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *CatalogListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *CatalogListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CatalogListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CatalogListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CatalogListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CatalogListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CatalogListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *CatalogListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetPackageIds returns the PackageIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogListDto) GetPackageIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PackageIds
}

// GetPackageIdsOk returns a tuple with the PackageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogListDto) GetPackageIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PackageIds) {
		return nil, false
	}
	return o.PackageIds, true
}

// HasPackageIds returns a boolean if a field has been set.
func (o *CatalogListDto) HasPackageIds() bool {
	if o != nil && IsNil(o.PackageIds) {
		return true
	}

	return false
}

// SetPackageIds gets a reference to the given []string and assigns it to the PackageIds field.
func (o *CatalogListDto) SetPackageIds(v []string) {
	o.PackageIds = v
}

// GetBoundProjects returns the BoundProjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogListDto) GetBoundProjects() []ProjectCatalogDto {
	if o == nil {
		var ret []ProjectCatalogDto
		return ret
	}
	return o.BoundProjects
}

// GetBoundProjectsOk returns a tuple with the BoundProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogListDto) GetBoundProjectsOk() ([]ProjectCatalogDto, bool) {
	if o == nil || IsNil(o.BoundProjects) {
		return nil, false
	}
	return o.BoundProjects, true
}

// HasBoundProjects returns a boolean if a field has been set.
func (o *CatalogListDto) HasBoundProjects() bool {
	if o != nil && IsNil(o.BoundProjects) {
		return true
	}

	return false
}

// SetBoundProjects gets a reference to the given []ProjectCatalogDto and assigns it to the BoundProjects field.
func (o *CatalogListDto) SetBoundProjects(v []ProjectCatalogDto) {
	o.BoundProjects = v
}

// GetBoundApplications returns the BoundApplications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogListDto) GetBoundApplications() []AvailablePackagesDto {
	if o == nil {
		var ret []AvailablePackagesDto
		return ret
	}
	return o.BoundApplications
}

// GetBoundApplicationsOk returns a tuple with the BoundApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogListDto) GetBoundApplicationsOk() ([]AvailablePackagesDto, bool) {
	if o == nil || IsNil(o.BoundApplications) {
		return nil, false
	}
	return o.BoundApplications, true
}

// HasBoundApplications returns a boolean if a field has been set.
func (o *CatalogListDto) HasBoundApplications() bool {
	if o != nil && IsNil(o.BoundApplications) {
		return true
	}

	return false
}

// SetBoundApplications gets a reference to the given []AvailablePackagesDto and assigns it to the BoundApplications field.
func (o *CatalogListDto) SetBoundApplications(v []AvailablePackagesDto) {
	o.BoundApplications = v
}

func (o CatalogListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.PackageIds != nil {
		toSerialize["packageIds"] = o.PackageIds
	}
	if o.BoundProjects != nil {
		toSerialize["boundProjects"] = o.BoundProjects
	}
	if o.BoundApplications != nil {
		toSerialize["boundApplications"] = o.BoundApplications
	}
	return toSerialize, nil
}

type NullableCatalogListDto struct {
	value *CatalogListDto
	isSet bool
}

func (v NullableCatalogListDto) Get() *CatalogListDto {
	return v.value
}

func (v *NullableCatalogListDto) Set(val *CatalogListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogListDto(val *CatalogListDto) *NullableCatalogListDto {
	return &NullableCatalogListDto{value: val, isSet: true}
}

func (v NullableCatalogListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


