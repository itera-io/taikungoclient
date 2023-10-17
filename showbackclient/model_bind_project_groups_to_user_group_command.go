/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the BindProjectGroupsToUserGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindProjectGroupsToUserGroupCommand{}

// BindProjectGroupsToUserGroupCommand struct for BindProjectGroupsToUserGroupCommand
type BindProjectGroupsToUserGroupCommand struct {
	ProjectGroups []UpdateUserProjectGroupDto `json:"projectGroups,omitempty"`
	UserGroupId   *int32                      `json:"userGroupId,omitempty"`
	UserGroupName NullableString              `json:"userGroupName,omitempty"`
}

// NewBindProjectGroupsToUserGroupCommand instantiates a new BindProjectGroupsToUserGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindProjectGroupsToUserGroupCommand() *BindProjectGroupsToUserGroupCommand {
	this := BindProjectGroupsToUserGroupCommand{}
	return &this
}

// NewBindProjectGroupsToUserGroupCommandWithDefaults instantiates a new BindProjectGroupsToUserGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindProjectGroupsToUserGroupCommandWithDefaults() *BindProjectGroupsToUserGroupCommand {
	this := BindProjectGroupsToUserGroupCommand{}
	return &this
}

// GetProjectGroups returns the ProjectGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindProjectGroupsToUserGroupCommand) GetProjectGroups() []UpdateUserProjectGroupDto {
	if o == nil {
		var ret []UpdateUserProjectGroupDto
		return ret
	}
	return o.ProjectGroups
}

// GetProjectGroupsOk returns a tuple with the ProjectGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindProjectGroupsToUserGroupCommand) GetProjectGroupsOk() ([]UpdateUserProjectGroupDto, bool) {
	if o == nil || IsNil(o.ProjectGroups) {
		return nil, false
	}
	return o.ProjectGroups, true
}

// HasProjectGroups returns a boolean if a field has been set.
func (o *BindProjectGroupsToUserGroupCommand) HasProjectGroups() bool {
	if o != nil && IsNil(o.ProjectGroups) {
		return true
	}

	return false
}

// SetProjectGroups gets a reference to the given []UpdateUserProjectGroupDto and assigns it to the ProjectGroups field.
func (o *BindProjectGroupsToUserGroupCommand) SetProjectGroups(v []UpdateUserProjectGroupDto) {
	o.ProjectGroups = v
}

// GetUserGroupId returns the UserGroupId field value if set, zero value otherwise.
func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupId() int32 {
	if o == nil || IsNil(o.UserGroupId) {
		var ret int32
		return ret
	}
	return *o.UserGroupId
}

// GetUserGroupIdOk returns a tuple with the UserGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserGroupId) {
		return nil, false
	}
	return o.UserGroupId, true
}

// HasUserGroupId returns a boolean if a field has been set.
func (o *BindProjectGroupsToUserGroupCommand) HasUserGroupId() bool {
	if o != nil && !IsNil(o.UserGroupId) {
		return true
	}

	return false
}

// SetUserGroupId gets a reference to the given int32 and assigns it to the UserGroupId field.
func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupId(v int32) {
	o.UserGroupId = &v
}

// GetUserGroupName returns the UserGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupName() string {
	if o == nil || IsNil(o.UserGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.UserGroupName.Get()
}

// GetUserGroupNameOk returns a tuple with the UserGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindProjectGroupsToUserGroupCommand) GetUserGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGroupName.Get(), o.UserGroupName.IsSet()
}

// HasUserGroupName returns a boolean if a field has been set.
func (o *BindProjectGroupsToUserGroupCommand) HasUserGroupName() bool {
	if o != nil && o.UserGroupName.IsSet() {
		return true
	}

	return false
}

// SetUserGroupName gets a reference to the given NullableString and assigns it to the UserGroupName field.
func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupName(v string) {
	o.UserGroupName.Set(&v)
}

// SetUserGroupNameNil sets the value for UserGroupName to be an explicit nil
func (o *BindProjectGroupsToUserGroupCommand) SetUserGroupNameNil() {
	o.UserGroupName.Set(nil)
}

// UnsetUserGroupName ensures that no value is present for UserGroupName, not even an explicit nil
func (o *BindProjectGroupsToUserGroupCommand) UnsetUserGroupName() {
	o.UserGroupName.Unset()
}

func (o BindProjectGroupsToUserGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindProjectGroupsToUserGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectGroups != nil {
		toSerialize["projectGroups"] = o.ProjectGroups
	}
	if !IsNil(o.UserGroupId) {
		toSerialize["userGroupId"] = o.UserGroupId
	}
	if o.UserGroupName.IsSet() {
		toSerialize["userGroupName"] = o.UserGroupName.Get()
	}
	return toSerialize, nil
}

type NullableBindProjectGroupsToUserGroupCommand struct {
	value *BindProjectGroupsToUserGroupCommand
	isSet bool
}

func (v NullableBindProjectGroupsToUserGroupCommand) Get() *BindProjectGroupsToUserGroupCommand {
	return v.value
}

func (v *NullableBindProjectGroupsToUserGroupCommand) Set(val *BindProjectGroupsToUserGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindProjectGroupsToUserGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindProjectGroupsToUserGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindProjectGroupsToUserGroupCommand(val *BindProjectGroupsToUserGroupCommand) *NullableBindProjectGroupsToUserGroupCommand {
	return &NullableBindProjectGroupsToUserGroupCommand{value: val, isSet: true}
}

func (v NullableBindProjectGroupsToUserGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindProjectGroupsToUserGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
