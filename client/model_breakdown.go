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

// checks if the Breakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Breakdown{}

// Breakdown struct for Breakdown
type Breakdown struct {
	Resources []Resource `json:"resources,omitempty"`
	TotalHourlyCost NullableString `json:"totalHourlyCost,omitempty"`
	TotalMonthlyCost NullableString `json:"totalMonthlyCost,omitempty"`
	TotalMonthlyUsageCost NullableString `json:"totalMonthlyUsageCost,omitempty"`
}

// NewBreakdown instantiates a new Breakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakdown() *Breakdown {
	this := Breakdown{}
	return &this
}

// NewBreakdownWithDefaults instantiates a new Breakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakdownWithDefaults() *Breakdown {
	this := Breakdown{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetResources() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetResourcesOk() ([]Resource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Breakdown) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *Breakdown) SetResources(v []Resource) {
	o.Resources = v
}

// GetTotalHourlyCost returns the TotalHourlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetTotalHourlyCost() string {
	if o == nil || IsNil(o.TotalHourlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.TotalHourlyCost.Get()
}

// GetTotalHourlyCostOk returns a tuple with the TotalHourlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetTotalHourlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalHourlyCost.Get(), o.TotalHourlyCost.IsSet()
}

// HasTotalHourlyCost returns a boolean if a field has been set.
func (o *Breakdown) HasTotalHourlyCost() bool {
	if o != nil && o.TotalHourlyCost.IsSet() {
		return true
	}

	return false
}

// SetTotalHourlyCost gets a reference to the given NullableString and assigns it to the TotalHourlyCost field.
func (o *Breakdown) SetTotalHourlyCost(v string) {
	o.TotalHourlyCost.Set(&v)
}
// SetTotalHourlyCostNil sets the value for TotalHourlyCost to be an explicit nil
func (o *Breakdown) SetTotalHourlyCostNil() {
	o.TotalHourlyCost.Set(nil)
}

// UnsetTotalHourlyCost ensures that no value is present for TotalHourlyCost, not even an explicit nil
func (o *Breakdown) UnsetTotalHourlyCost() {
	o.TotalHourlyCost.Unset()
}

// GetTotalMonthlyCost returns the TotalMonthlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetTotalMonthlyCost() string {
	if o == nil || IsNil(o.TotalMonthlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.TotalMonthlyCost.Get()
}

// GetTotalMonthlyCostOk returns a tuple with the TotalMonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetTotalMonthlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalMonthlyCost.Get(), o.TotalMonthlyCost.IsSet()
}

// HasTotalMonthlyCost returns a boolean if a field has been set.
func (o *Breakdown) HasTotalMonthlyCost() bool {
	if o != nil && o.TotalMonthlyCost.IsSet() {
		return true
	}

	return false
}

// SetTotalMonthlyCost gets a reference to the given NullableString and assigns it to the TotalMonthlyCost field.
func (o *Breakdown) SetTotalMonthlyCost(v string) {
	o.TotalMonthlyCost.Set(&v)
}
// SetTotalMonthlyCostNil sets the value for TotalMonthlyCost to be an explicit nil
func (o *Breakdown) SetTotalMonthlyCostNil() {
	o.TotalMonthlyCost.Set(nil)
}

// UnsetTotalMonthlyCost ensures that no value is present for TotalMonthlyCost, not even an explicit nil
func (o *Breakdown) UnsetTotalMonthlyCost() {
	o.TotalMonthlyCost.Unset()
}

// GetTotalMonthlyUsageCost returns the TotalMonthlyUsageCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetTotalMonthlyUsageCost() string {
	if o == nil || IsNil(o.TotalMonthlyUsageCost.Get()) {
		var ret string
		return ret
	}
	return *o.TotalMonthlyUsageCost.Get()
}

// GetTotalMonthlyUsageCostOk returns a tuple with the TotalMonthlyUsageCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetTotalMonthlyUsageCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalMonthlyUsageCost.Get(), o.TotalMonthlyUsageCost.IsSet()
}

// HasTotalMonthlyUsageCost returns a boolean if a field has been set.
func (o *Breakdown) HasTotalMonthlyUsageCost() bool {
	if o != nil && o.TotalMonthlyUsageCost.IsSet() {
		return true
	}

	return false
}

// SetTotalMonthlyUsageCost gets a reference to the given NullableString and assigns it to the TotalMonthlyUsageCost field.
func (o *Breakdown) SetTotalMonthlyUsageCost(v string) {
	o.TotalMonthlyUsageCost.Set(&v)
}
// SetTotalMonthlyUsageCostNil sets the value for TotalMonthlyUsageCost to be an explicit nil
func (o *Breakdown) SetTotalMonthlyUsageCostNil() {
	o.TotalMonthlyUsageCost.Set(nil)
}

// UnsetTotalMonthlyUsageCost ensures that no value is present for TotalMonthlyUsageCost, not even an explicit nil
func (o *Breakdown) UnsetTotalMonthlyUsageCost() {
	o.TotalMonthlyUsageCost.Unset()
}

func (o Breakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Breakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.TotalHourlyCost.IsSet() {
		toSerialize["totalHourlyCost"] = o.TotalHourlyCost.Get()
	}
	if o.TotalMonthlyCost.IsSet() {
		toSerialize["totalMonthlyCost"] = o.TotalMonthlyCost.Get()
	}
	if o.TotalMonthlyUsageCost.IsSet() {
		toSerialize["totalMonthlyUsageCost"] = o.TotalMonthlyUsageCost.Get()
	}
	return toSerialize, nil
}

type NullableBreakdown struct {
	value *Breakdown
	isSet bool
}

func (v NullableBreakdown) Get() *Breakdown {
	return v.value
}

func (v *NullableBreakdown) Set(val *Breakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakdown(val *Breakdown) *NullableBreakdown {
	return &NullableBreakdown{value: val, isSet: true}
}

func (v NullableBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


