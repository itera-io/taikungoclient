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

// checks if the BindAppRepositoryCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindAppRepositoryCommand{}

// BindAppRepositoryCommand struct for BindAppRepositoryCommand
type BindAppRepositoryCommand struct {
	FilteringElements []FilteringElementDto `json:"filteringElements,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
}

// NewBindAppRepositoryCommand instantiates a new BindAppRepositoryCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindAppRepositoryCommand() *BindAppRepositoryCommand {
	this := BindAppRepositoryCommand{}
	return &this
}

// NewBindAppRepositoryCommandWithDefaults instantiates a new BindAppRepositoryCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindAppRepositoryCommandWithDefaults() *BindAppRepositoryCommand {
	this := BindAppRepositoryCommand{}
	return &this
}

// GetFilteringElements returns the FilteringElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindAppRepositoryCommand) GetFilteringElements() []FilteringElementDto {
	if o == nil {
		var ret []FilteringElementDto
		return ret
	}
	return o.FilteringElements
}

// GetFilteringElementsOk returns a tuple with the FilteringElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindAppRepositoryCommand) GetFilteringElementsOk() ([]FilteringElementDto, bool) {
	if o == nil || IsNil(o.FilteringElements) {
		return nil, false
	}
	return o.FilteringElements, true
}

// HasFilteringElements returns a boolean if a field has been set.
func (o *BindAppRepositoryCommand) HasFilteringElements() bool {
	if o != nil && !IsNil(o.FilteringElements) {
		return true
	}

	return false
}

// SetFilteringElements gets a reference to the given []FilteringElementDto and assigns it to the FilteringElements field.
func (o *BindAppRepositoryCommand) SetFilteringElements(v []FilteringElementDto) {
	o.FilteringElements = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindAppRepositoryCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindAppRepositoryCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *BindAppRepositoryCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *BindAppRepositoryCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *BindAppRepositoryCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *BindAppRepositoryCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o BindAppRepositoryCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindAppRepositoryCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilteringElements != nil {
		toSerialize["filteringElements"] = o.FilteringElements
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableBindAppRepositoryCommand struct {
	value *BindAppRepositoryCommand
	isSet bool
}

func (v NullableBindAppRepositoryCommand) Get() *BindAppRepositoryCommand {
	return v.value
}

func (v *NullableBindAppRepositoryCommand) Set(val *BindAppRepositoryCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindAppRepositoryCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindAppRepositoryCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindAppRepositoryCommand(val *BindAppRepositoryCommand) *NullableBindAppRepositoryCommand {
	return &NullableBindAppRepositoryCommand{value: val, isSet: true}
}

func (v NullableBindAppRepositoryCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindAppRepositoryCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


