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
	"bytes"
	"fmt"
)

// checks if the UserGroupDetailsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupDetailsListDto{}

// UserGroupDetailsListDto struct for UserGroupDetailsListDto
type UserGroupDetailsListDto struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	OrganizationId int32 `json:"organizationId"`
	OrganizationName NullableString `json:"organizationName"`
	Users []UserListDto `json:"users"`
	ProjectGroups []ProjectGroupEntityListDto `json:"projectGroups"`
}

type _UserGroupDetailsListDto UserGroupDetailsListDto

// NewUserGroupDetailsListDto instantiates a new UserGroupDetailsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupDetailsListDto(id int32, name NullableString, organizationId int32, organizationName NullableString, users []UserListDto, projectGroups []ProjectGroupEntityListDto) *UserGroupDetailsListDto {
	this := UserGroupDetailsListDto{}
	this.Id = id
	this.Name = name
	this.OrganizationId = organizationId
	this.OrganizationName = organizationName
	this.Users = users
	this.ProjectGroups = projectGroups
	return &this
}

// NewUserGroupDetailsListDtoWithDefaults instantiates a new UserGroupDetailsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupDetailsListDtoWithDefaults() *UserGroupDetailsListDto {
	this := UserGroupDetailsListDto{}
	return &this
}

// GetId returns the Id field value
func (o *UserGroupDetailsListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserGroupDetailsListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserGroupDetailsListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserGroupDetailsListDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *UserGroupDetailsListDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetOrganizationId returns the OrganizationId field value
func (o *UserGroupDetailsListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *UserGroupDetailsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *UserGroupDetailsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationName returns the OrganizationName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserGroupDetailsListDto) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// SetOrganizationName sets field value
func (o *UserGroupDetailsListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// GetUsers returns the Users field value
// If the value is explicit nil, the zero value for []UserListDto will be returned
func (o *UserGroupDetailsListDto) GetUsers() []UserListDto {
	if o == nil {
		var ret []UserListDto
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetUsersOk() ([]UserListDto, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UserGroupDetailsListDto) SetUsers(v []UserListDto) {
	o.Users = v
}

// GetProjectGroups returns the ProjectGroups field value
// If the value is explicit nil, the zero value for []ProjectGroupEntityListDto will be returned
func (o *UserGroupDetailsListDto) GetProjectGroups() []ProjectGroupEntityListDto {
	if o == nil {
		var ret []ProjectGroupEntityListDto
		return ret
	}

	return o.ProjectGroups
}

// GetProjectGroupsOk returns a tuple with the ProjectGroups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupDetailsListDto) GetProjectGroupsOk() ([]ProjectGroupEntityListDto, bool) {
	if o == nil || IsNil(o.ProjectGroups) {
		return nil, false
	}
	return o.ProjectGroups, true
}

// SetProjectGroups sets field value
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
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationName"] = o.OrganizationName.Get()
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.ProjectGroups != nil {
		toSerialize["projectGroups"] = o.ProjectGroups
	}
	return toSerialize, nil
}

func (o *UserGroupDetailsListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"organizationId",
		"organizationName",
		"users",
		"projectGroups",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserGroupDetailsListDto := _UserGroupDetailsListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroupDetailsListDto)

	if err != nil {
		return err
	}

	*o = UserGroupDetailsListDto(varUserGroupDetailsListDto)

	return err
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


