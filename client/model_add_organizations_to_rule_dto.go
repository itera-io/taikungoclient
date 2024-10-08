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

// checks if the AddOrganizationsToRuleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrganizationsToRuleDto{}

// AddOrganizationsToRuleDto struct for AddOrganizationsToRuleDto
type AddOrganizationsToRuleDto struct {
	OrganizationId *int32 `json:"organizationId,omitempty"`
	RuleDiscountRate NullableFloat64 `json:"ruleDiscountRate,omitempty"`
}

// NewAddOrganizationsToRuleDto instantiates a new AddOrganizationsToRuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrganizationsToRuleDto() *AddOrganizationsToRuleDto {
	this := AddOrganizationsToRuleDto{}
	return &this
}

// NewAddOrganizationsToRuleDtoWithDefaults instantiates a new AddOrganizationsToRuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrganizationsToRuleDtoWithDefaults() *AddOrganizationsToRuleDto {
	this := AddOrganizationsToRuleDto{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *AddOrganizationsToRuleDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOrganizationsToRuleDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AddOrganizationsToRuleDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *AddOrganizationsToRuleDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetRuleDiscountRate returns the RuleDiscountRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddOrganizationsToRuleDto) GetRuleDiscountRate() float64 {
	if o == nil || IsNil(o.RuleDiscountRate.Get()) {
		var ret float64
		return ret
	}
	return *o.RuleDiscountRate.Get()
}

// GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddOrganizationsToRuleDto) GetRuleDiscountRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleDiscountRate.Get(), o.RuleDiscountRate.IsSet()
}

// HasRuleDiscountRate returns a boolean if a field has been set.
func (o *AddOrganizationsToRuleDto) HasRuleDiscountRate() bool {
	if o != nil && o.RuleDiscountRate.IsSet() {
		return true
	}

	return false
}

// SetRuleDiscountRate gets a reference to the given NullableFloat64 and assigns it to the RuleDiscountRate field.
func (o *AddOrganizationsToRuleDto) SetRuleDiscountRate(v float64) {
	o.RuleDiscountRate.Set(&v)
}
// SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil
func (o *AddOrganizationsToRuleDto) SetRuleDiscountRateNil() {
	o.RuleDiscountRate.Set(nil)
}

// UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil
func (o *AddOrganizationsToRuleDto) UnsetRuleDiscountRate() {
	o.RuleDiscountRate.Unset()
}

func (o AddOrganizationsToRuleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrganizationsToRuleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.RuleDiscountRate.IsSet() {
		toSerialize["ruleDiscountRate"] = o.RuleDiscountRate.Get()
	}
	return toSerialize, nil
}

type NullableAddOrganizationsToRuleDto struct {
	value *AddOrganizationsToRuleDto
	isSet bool
}

func (v NullableAddOrganizationsToRuleDto) Get() *AddOrganizationsToRuleDto {
	return v.value
}

func (v *NullableAddOrganizationsToRuleDto) Set(val *AddOrganizationsToRuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrganizationsToRuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrganizationsToRuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrganizationsToRuleDto(val *AddOrganizationsToRuleDto) *NullableAddOrganizationsToRuleDto {
	return &NullableAddOrganizationsToRuleDto{value: val, isSet: true}
}

func (v NullableAddOrganizationsToRuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrganizationsToRuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


