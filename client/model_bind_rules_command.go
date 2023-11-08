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

// checks if the BindRulesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindRulesCommand{}

// BindRulesCommand struct for BindRulesCommand
type BindRulesCommand struct {
	Rules []BindRulesToOrganizationDto `json:"rules,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
}

// NewBindRulesCommand instantiates a new BindRulesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindRulesCommand() *BindRulesCommand {
	this := BindRulesCommand{}
	return &this
}

// NewBindRulesCommandWithDefaults instantiates a new BindRulesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindRulesCommandWithDefaults() *BindRulesCommand {
	this := BindRulesCommand{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindRulesCommand) GetRules() []BindRulesToOrganizationDto {
	if o == nil {
		var ret []BindRulesToOrganizationDto
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindRulesCommand) GetRulesOk() ([]BindRulesToOrganizationDto, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *BindRulesCommand) HasRules() bool {
	if o != nil && IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []BindRulesToOrganizationDto and assigns it to the Rules field.
func (o *BindRulesCommand) SetRules(v []BindRulesToOrganizationDto) {
	o.Rules = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *BindRulesCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindRulesCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *BindRulesCommand) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *BindRulesCommand) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

func (o BindRulesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindRulesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	return toSerialize, nil
}

type NullableBindRulesCommand struct {
	value *BindRulesCommand
	isSet bool
}

func (v NullableBindRulesCommand) Get() *BindRulesCommand {
	return v.value
}

func (v *NullableBindRulesCommand) Set(val *BindRulesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindRulesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindRulesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindRulesCommand(val *BindRulesCommand) *NullableBindRulesCommand {
	return &NullableBindRulesCommand{value: val, isSet: true}
}

func (v NullableBindRulesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindRulesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


