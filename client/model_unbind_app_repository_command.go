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

// checks if the UnbindAppRepositoryCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnbindAppRepositoryCommand{}

// UnbindAppRepositoryCommand struct for UnbindAppRepositoryCommand
type UnbindAppRepositoryCommand struct {
	Ids []string `json:"ids,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewUnbindAppRepositoryCommand instantiates a new UnbindAppRepositoryCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnbindAppRepositoryCommand() *UnbindAppRepositoryCommand {
	this := UnbindAppRepositoryCommand{}
	return &this
}

// NewUnbindAppRepositoryCommandWithDefaults instantiates a new UnbindAppRepositoryCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbindAppRepositoryCommandWithDefaults() *UnbindAppRepositoryCommand {
	this := UnbindAppRepositoryCommand{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnbindAppRepositoryCommand) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnbindAppRepositoryCommand) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *UnbindAppRepositoryCommand) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *UnbindAppRepositoryCommand) SetIds(v []string) {
	o.Ids = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnbindAppRepositoryCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnbindAppRepositoryCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UnbindAppRepositoryCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *UnbindAppRepositoryCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *UnbindAppRepositoryCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *UnbindAppRepositoryCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o UnbindAppRepositoryCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnbindAppRepositoryCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableUnbindAppRepositoryCommand struct {
	value *UnbindAppRepositoryCommand
	isSet bool
}

func (v NullableUnbindAppRepositoryCommand) Get() *UnbindAppRepositoryCommand {
	return v.value
}

func (v *NullableUnbindAppRepositoryCommand) Set(val *UnbindAppRepositoryCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUnbindAppRepositoryCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUnbindAppRepositoryCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnbindAppRepositoryCommand(val *UnbindAppRepositoryCommand) *NullableUnbindAppRepositoryCommand {
	return &NullableUnbindAppRepositoryCommand{value: val, isSet: true}
}

func (v NullableUnbindAppRepositoryCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnbindAppRepositoryCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


