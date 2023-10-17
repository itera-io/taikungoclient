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

// checks if the DuplicateNameCheckerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicateNameCheckerCommand{}

// DuplicateNameCheckerCommand struct for DuplicateNameCheckerCommand
type DuplicateNameCheckerCommand struct {
	OrganizationId NullableInt32  `json:"organizationId,omitempty"`
	Type           NullableString `json:"type,omitempty"`
	Name           NullableString `json:"name,omitempty"`
}

// NewDuplicateNameCheckerCommand instantiates a new DuplicateNameCheckerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicateNameCheckerCommand() *DuplicateNameCheckerCommand {
	this := DuplicateNameCheckerCommand{}
	return &this
}

// NewDuplicateNameCheckerCommandWithDefaults instantiates a new DuplicateNameCheckerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicateNameCheckerCommandWithDefaults() *DuplicateNameCheckerCommand {
	this := DuplicateNameCheckerCommand{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateNameCheckerCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateNameCheckerCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *DuplicateNameCheckerCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *DuplicateNameCheckerCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *DuplicateNameCheckerCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *DuplicateNameCheckerCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateNameCheckerCommand) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateNameCheckerCommand) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *DuplicateNameCheckerCommand) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *DuplicateNameCheckerCommand) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *DuplicateNameCheckerCommand) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *DuplicateNameCheckerCommand) UnsetType() {
	o.Type.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateNameCheckerCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateNameCheckerCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DuplicateNameCheckerCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DuplicateNameCheckerCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DuplicateNameCheckerCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DuplicateNameCheckerCommand) UnsetName() {
	o.Name.Unset()
}

func (o DuplicateNameCheckerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicateNameCheckerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableDuplicateNameCheckerCommand struct {
	value *DuplicateNameCheckerCommand
	isSet bool
}

func (v NullableDuplicateNameCheckerCommand) Get() *DuplicateNameCheckerCommand {
	return v.value
}

func (v *NullableDuplicateNameCheckerCommand) Set(val *DuplicateNameCheckerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicateNameCheckerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicateNameCheckerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicateNameCheckerCommand(val *DuplicateNameCheckerCommand) *NullableDuplicateNameCheckerCommand {
	return &NullableDuplicateNameCheckerCommand{value: val, isSet: true}
}

func (v NullableDuplicateNameCheckerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicateNameCheckerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
