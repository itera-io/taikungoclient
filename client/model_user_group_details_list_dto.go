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

// checks if the UserGroupDetailsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupDetailsListDto{}

// UserGroupDetailsListDto struct for UserGroupDetailsListDto
type UserGroupDetailsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Users []UserListDto `json:"users,omitempty"`
	ProjectGroups []ProjectGroupEntityListDto `json:"projectGroups,omitempty"`
}

// NewUserGroupDetailsListDto instantiates a new UserGroupDetailsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupDetailsListDto() *UserGroupDetailsListDto {
	this := UserGroupDetailsListDto{}
	return &this
}

// NewUserGroupDetailsListDtoWithDefaults instantiates a new UserGroupDetailsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupDetailsListDtoWithDefaults() *UserGroupDetailsListDto {
	this := UserGroupDetailsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserGroupDetailsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupDetailsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserGroupDetailsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroupDetailsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserGroupDetailsListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserGroupDetailsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserGroupDetailsListDto) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *UserGroupDetailsListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupDetailsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *UserGroupDetailsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroupDetailsListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *UserGroupDetailsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *UserGroupDetailsListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *UserGroupDetailsListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroupDetailsListDto) GetUsers() []UserListDto {
	if o == nil {
		var ret []UserListDto
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetUsersOk() ([]UserListDto, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserListDto and assigns it to the Users field.
func (o *UserGroupDetailsListDto) SetUsers(v []UserListDto) {
	o.Users = v
}

// GetProjectGroups returns the ProjectGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroupDetailsListDto) GetProjectGroups() []ProjectGroupEntityListDto {
	if o == nil {
		var ret []ProjectGroupEntityListDto
		return ret
	}
	return o.ProjectGroups
}

// GetProjectGroupsOk returns a tuple with the ProjectGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetProjectGroupsOk() ([]ProjectGroupEntityListDto, bool) {
	if o == nil || IsNil(o.ProjectGroups) {
		return nil, false
	}
	return o.ProjectGroups, true
}

// HasProjectGroups returns a boolean if a field has been set.
func (o *UserGroupDetailsListDto) HasProjectGroups() bool {
	if o != nil && !IsNil(o.ProjectGroups) {
		return true
	}

	return false
}

// SetProjectGroups gets a reference to the given []ProjectGroupEntityListDto and assigns it to the ProjectGroups field.
func (o *UserGroupDetailsListDto) SetProjectGroups(v []ProjectGroupEntityListDto) {
	o.ProjectGroups = v
}

func (o UserGroupDetailsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupDetailsListDto) ToMap() (map[string]interface{}, error) {
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
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.ProjectGroups != nil {
		toSerialize["projectGroups"] = o.ProjectGroups
	}
	return toSerialize, nil
}

type NullableUserGroupDetailsListDto struct {
	value *UserGroupDetailsListDto
	isSet bool
}

func (v NullableUserGroupDetailsListDto) Get() *UserGroupDetailsListDto {
	return v.value
}

func (v *NullableUserGroupDetailsListDto) Set(val *UserGroupDetailsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupDetailsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupDetailsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupDetailsListDto(val *UserGroupDetailsListDto) *NullableUserGroupDetailsListDto {
	return &NullableUserGroupDetailsListDto{value: val, isSet: true}
}

func (v NullableUserGroupDetailsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupDetailsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


