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

// checks if the CreateUserGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserGroupCommand{}

// CreateUserGroupCommand struct for CreateUserGroupCommand
type CreateUserGroupCommand struct {
	Name NullableString `json:"name,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	UserIds []string `json:"userIds,omitempty"`
}

// NewCreateUserGroupCommand instantiates a new CreateUserGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserGroupCommand() *CreateUserGroupCommand {
	this := CreateUserGroupCommand{}
	return &this
}

// NewCreateUserGroupCommandWithDefaults instantiates a new CreateUserGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserGroupCommandWithDefaults() *CreateUserGroupCommand {
	this := CreateUserGroupCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserGroupCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserGroupCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateUserGroupCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateUserGroupCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateUserGroupCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateUserGroupCommand) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserGroupCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserGroupCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateUserGroupCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateUserGroupCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateUserGroupCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateUserGroupCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserGroupCommand) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserGroupCommand) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *CreateUserGroupCommand) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *CreateUserGroupCommand) SetUserIds(v []string) {
	o.UserIds = v
}

func (o CreateUserGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

type NullableCreateUserGroupCommand struct {
	value *CreateUserGroupCommand
	isSet bool
}

func (v NullableCreateUserGroupCommand) Get() *CreateUserGroupCommand {
	return v.value
}

func (v *NullableCreateUserGroupCommand) Set(val *CreateUserGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserGroupCommand(val *CreateUserGroupCommand) *NullableCreateUserGroupCommand {
	return &NullableCreateUserGroupCommand{value: val, isSet: true}
}

func (v NullableCreateUserGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


