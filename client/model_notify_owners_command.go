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

// checks if the NotifyOwnersCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyOwnersCommand{}

// NotifyOwnersCommand struct for NotifyOwnersCommand
type NotifyOwnersCommand struct {
	OrganizationId *int32 `json:"organizationId,omitempty"`
}

// NewNotifyOwnersCommand instantiates a new NotifyOwnersCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyOwnersCommand() *NotifyOwnersCommand {
	this := NotifyOwnersCommand{}
	return &this
}

// NewNotifyOwnersCommandWithDefaults instantiates a new NotifyOwnersCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyOwnersCommandWithDefaults() *NotifyOwnersCommand {
	this := NotifyOwnersCommand{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *NotifyOwnersCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyOwnersCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *NotifyOwnersCommand) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *NotifyOwnersCommand) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

func (o NotifyOwnersCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyOwnersCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	return toSerialize, nil
}

type NullableNotifyOwnersCommand struct {
	value *NotifyOwnersCommand
	isSet bool
}

func (v NullableNotifyOwnersCommand) Get() *NotifyOwnersCommand {
	return v.value
}

func (v *NullableNotifyOwnersCommand) Set(val *NotifyOwnersCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyOwnersCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyOwnersCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyOwnersCommand(val *NotifyOwnersCommand) *NullableNotifyOwnersCommand {
	return &NullableNotifyOwnersCommand{value: val, isSet: true}
}

func (v NullableNotifyOwnersCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyOwnersCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
