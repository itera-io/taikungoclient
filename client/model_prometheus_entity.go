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

// checks if the PrometheusEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusEntity{}

// PrometheusEntity struct for PrometheusEntity
type PrometheusEntity struct {
	PrometheusRuleId   *int32         `json:"prometheusRuleId,omitempty"`
	PrometheusRuleName NullableString `json:"prometheusRuleName,omitempty"`
	RuleDiscountRate   *float64       `json:"ruleDiscountRate,omitempty"`
}

// NewPrometheusEntity instantiates a new PrometheusEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusEntity() *PrometheusEntity {
	this := PrometheusEntity{}
	return &this
}

// NewPrometheusEntityWithDefaults instantiates a new PrometheusEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusEntityWithDefaults() *PrometheusEntity {
	this := PrometheusEntity{}
	return &this
}

// GetPrometheusRuleId returns the PrometheusRuleId field value if set, zero value otherwise.
func (o *PrometheusEntity) GetPrometheusRuleId() int32 {
	if o == nil || IsNil(o.PrometheusRuleId) {
		var ret int32
		return ret
	}
	return *o.PrometheusRuleId
}

// GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusEntity) GetPrometheusRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PrometheusRuleId) {
		return nil, false
	}
	return o.PrometheusRuleId, true
}

// HasPrometheusRuleId returns a boolean if a field has been set.
func (o *PrometheusEntity) HasPrometheusRuleId() bool {
	if o != nil && !IsNil(o.PrometheusRuleId) {
		return true
	}

	return false
}

// SetPrometheusRuleId gets a reference to the given int32 and assigns it to the PrometheusRuleId field.
func (o *PrometheusEntity) SetPrometheusRuleId(v int32) {
	o.PrometheusRuleId = &v
}

// GetPrometheusRuleName returns the PrometheusRuleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusEntity) GetPrometheusRuleName() string {
	if o == nil || IsNil(o.PrometheusRuleName.Get()) {
		var ret string
		return ret
	}
	return *o.PrometheusRuleName.Get()
}

// GetPrometheusRuleNameOk returns a tuple with the PrometheusRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusEntity) GetPrometheusRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrometheusRuleName.Get(), o.PrometheusRuleName.IsSet()
}

// HasPrometheusRuleName returns a boolean if a field has been set.
func (o *PrometheusEntity) HasPrometheusRuleName() bool {
	if o != nil && o.PrometheusRuleName.IsSet() {
		return true
	}

	return false
}

// SetPrometheusRuleName gets a reference to the given NullableString and assigns it to the PrometheusRuleName field.
func (o *PrometheusEntity) SetPrometheusRuleName(v string) {
	o.PrometheusRuleName.Set(&v)
}

// SetPrometheusRuleNameNil sets the value for PrometheusRuleName to be an explicit nil
func (o *PrometheusEntity) SetPrometheusRuleNameNil() {
	o.PrometheusRuleName.Set(nil)
}

// UnsetPrometheusRuleName ensures that no value is present for PrometheusRuleName, not even an explicit nil
func (o *PrometheusEntity) UnsetPrometheusRuleName() {
	o.PrometheusRuleName.Unset()
}

// GetRuleDiscountRate returns the RuleDiscountRate field value if set, zero value otherwise.
func (o *PrometheusEntity) GetRuleDiscountRate() float64 {
	if o == nil || IsNil(o.RuleDiscountRate) {
		var ret float64
		return ret
	}
	return *o.RuleDiscountRate
}

// GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusEntity) GetRuleDiscountRateOk() (*float64, bool) {
	if o == nil || IsNil(o.RuleDiscountRate) {
		return nil, false
	}
	return o.RuleDiscountRate, true
}

// HasRuleDiscountRate returns a boolean if a field has been set.
func (o *PrometheusEntity) HasRuleDiscountRate() bool {
	if o != nil && !IsNil(o.RuleDiscountRate) {
		return true
	}

	return false
}

// SetRuleDiscountRate gets a reference to the given float64 and assigns it to the RuleDiscountRate field.
func (o *PrometheusEntity) SetRuleDiscountRate(v float64) {
	o.RuleDiscountRate = &v
}

func (o PrometheusEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrometheusRuleId) {
		toSerialize["prometheusRuleId"] = o.PrometheusRuleId
	}
	if o.PrometheusRuleName.IsSet() {
		toSerialize["prometheusRuleName"] = o.PrometheusRuleName.Get()
	}
	if !IsNil(o.RuleDiscountRate) {
		toSerialize["ruleDiscountRate"] = o.RuleDiscountRate
	}
	return toSerialize, nil
}

type NullablePrometheusEntity struct {
	value *PrometheusEntity
	isSet bool
}

func (v NullablePrometheusEntity) Get() *PrometheusEntity {
	return v.value
}

func (v *NullablePrometheusEntity) Set(val *PrometheusEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusEntity(val *PrometheusEntity) *NullablePrometheusEntity {
	return &NullablePrometheusEntity{value: val, isSet: true}
}

func (v NullablePrometheusEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
