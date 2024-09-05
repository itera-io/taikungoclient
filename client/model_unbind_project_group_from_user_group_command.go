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

// checks if the UnbindProjectGroupFromUserGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnbindProjectGroupFromUserGroupCommand{}

// UnbindProjectGroupFromUserGroupCommand struct for UnbindProjectGroupFromUserGroupCommand
type UnbindProjectGroupFromUserGroupCommand struct {
	UserGroupId *int32 `json:"userGroupId,omitempty"`
	ProjectGroupIds []int32 `json:"projectGroupIds,omitempty"`
}

// NewUnbindProjectGroupFromUserGroupCommand instantiates a new UnbindProjectGroupFromUserGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnbindProjectGroupFromUserGroupCommand() *UnbindProjectGroupFromUserGroupCommand {
	this := UnbindProjectGroupFromUserGroupCommand{}
	return &this
}

// NewUnbindProjectGroupFromUserGroupCommandWithDefaults instantiates a new UnbindProjectGroupFromUserGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbindProjectGroupFromUserGroupCommandWithDefaults() *UnbindProjectGroupFromUserGroupCommand {
	this := UnbindProjectGroupFromUserGroupCommand{}
	return &this
}

// GetUserGroupId returns the UserGroupId field value if set, zero value otherwise.
func (o *UnbindProjectGroupFromUserGroupCommand) GetUserGroupId() int32 {
	if o == nil || IsNil(o.UserGroupId) {
		var ret int32
		return ret
	}
	return *o.UserGroupId
}

// GetUserGroupIdOk returns a tuple with the UserGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindProjectGroupFromUserGroupCommand) GetUserGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserGroupId) {
		return nil, false
	}
	return o.UserGroupId, true
}

// HasUserGroupId returns a boolean if a field has been set.
func (o *UnbindProjectGroupFromUserGroupCommand) HasUserGroupId() bool {
	if o != nil && !IsNil(o.UserGroupId) {
		return true
	}

	return false
}

// SetUserGroupId gets a reference to the given int32 and assigns it to the UserGroupId field.
func (o *UnbindProjectGroupFromUserGroupCommand) SetUserGroupId(v int32) {
	o.UserGroupId = &v
}

// GetProjectGroupIds returns the ProjectGroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnbindProjectGroupFromUserGroupCommand) GetProjectGroupIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ProjectGroupIds
}

// GetProjectGroupIdsOk returns a tuple with the ProjectGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnbindProjectGroupFromUserGroupCommand) GetProjectGroupIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ProjectGroupIds) {
		return nil, false
	}
	return o.ProjectGroupIds, true
}

// HasProjectGroupIds returns a boolean if a field has been set.
func (o *UnbindProjectGroupFromUserGroupCommand) HasProjectGroupIds() bool {
	if o != nil && !IsNil(o.ProjectGroupIds) {
		return true
	}

	return false
}

// SetProjectGroupIds gets a reference to the given []int32 and assigns it to the ProjectGroupIds field.
func (o *UnbindProjectGroupFromUserGroupCommand) SetProjectGroupIds(v []int32) {
	o.ProjectGroupIds = v
}

func (o UnbindProjectGroupFromUserGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnbindProjectGroupFromUserGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserGroupId) {
		toSerialize["userGroupId"] = o.UserGroupId
	}
	if o.ProjectGroupIds != nil {
		toSerialize["projectGroupIds"] = o.ProjectGroupIds
	}
	return toSerialize, nil
}

type NullableUnbindProjectGroupFromUserGroupCommand struct {
	value *UnbindProjectGroupFromUserGroupCommand
	isSet bool
}

func (v NullableUnbindProjectGroupFromUserGroupCommand) Get() *UnbindProjectGroupFromUserGroupCommand {
	return v.value
}

func (v *NullableUnbindProjectGroupFromUserGroupCommand) Set(val *UnbindProjectGroupFromUserGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUnbindProjectGroupFromUserGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUnbindProjectGroupFromUserGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnbindProjectGroupFromUserGroupCommand(val *UnbindProjectGroupFromUserGroupCommand) *NullableUnbindProjectGroupFromUserGroupCommand {
	return &NullableUnbindProjectGroupFromUserGroupCommand{value: val, isSet: true}
}

func (v NullableUnbindProjectGroupFromUserGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnbindProjectGroupFromUserGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


