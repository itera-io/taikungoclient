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

// checks if the ProjectGroupDetailsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectGroupDetailsListDto{}

// ProjectGroupDetailsListDto struct for ProjectGroupDetailsListDto
type ProjectGroupDetailsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Projects []ProjectListDto `json:"projects,omitempty"`
	UserGroups []UserGroupEntityListDto `json:"userGroups,omitempty"`
}

// NewProjectGroupDetailsListDto instantiates a new ProjectGroupDetailsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectGroupDetailsListDto() *ProjectGroupDetailsListDto {
	this := ProjectGroupDetailsListDto{}
	return &this
}

// NewProjectGroupDetailsListDtoWithDefaults instantiates a new ProjectGroupDetailsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectGroupDetailsListDtoWithDefaults() *ProjectGroupDetailsListDto {
	this := ProjectGroupDetailsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectGroupDetailsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectGroupDetailsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectGroupDetailsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectGroupDetailsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectGroupDetailsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectGroupDetailsListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectGroupDetailsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectGroupDetailsListDto) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ProjectGroupDetailsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectGroupDetailsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *ProjectGroupDetailsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectGroupDetailsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectGroupDetailsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *ProjectGroupDetailsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *ProjectGroupDetailsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *ProjectGroupDetailsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectGroupDetailsListDto) GetProjects() []ProjectListDto {
	if o == nil {
		var ret []ProjectListDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectGroupDetailsListDto) GetProjectsOk() ([]ProjectListDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectListDto and assigns it to the Projects field.
func (o *ProjectGroupDetailsListDto) SetProjects(v []ProjectListDto) {
	o.Projects = v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectGroupDetailsListDto) GetUserGroups() []UserGroupEntityListDto {
	if o == nil {
		var ret []UserGroupEntityListDto
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectGroupDetailsListDto) GetUserGroupsOk() ([]UserGroupEntityListDto, bool) {
	if o == nil || IsNil(o.UserGroups) {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *ProjectGroupDetailsListDto) HasUserGroups() bool {
	if o != nil && !IsNil(o.UserGroups) {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []UserGroupEntityListDto and assigns it to the UserGroups field.
func (o *ProjectGroupDetailsListDto) SetUserGroups(v []UserGroupEntityListDto) {
	o.UserGroups = v
}

func (o ProjectGroupDetailsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectGroupDetailsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.UserGroups != nil {
		toSerialize["userGroups"] = o.UserGroups
	}
	return toSerialize, nil
}

type NullableProjectGroupDetailsListDto struct {
	value *ProjectGroupDetailsListDto
	isSet bool
}

func (v NullableProjectGroupDetailsListDto) Get() *ProjectGroupDetailsListDto {
	return v.value
}

func (v *NullableProjectGroupDetailsListDto) Set(val *ProjectGroupDetailsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectGroupDetailsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectGroupDetailsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectGroupDetailsListDto(val *ProjectGroupDetailsListDto) *NullableProjectGroupDetailsListDto {
	return &NullableProjectGroupDetailsListDto{value: val, isSet: true}
}

func (v NullableProjectGroupDetailsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectGroupDetailsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


