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

// checks if the PrometheusRulesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusRulesList{}

// PrometheusRulesList struct for PrometheusRulesList
type PrometheusRulesList struct {
	Data []PrometheusRuleListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPrometheusRulesList instantiates a new PrometheusRulesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusRulesList() *PrometheusRulesList {
	this := PrometheusRulesList{}
	return &this
}

// NewPrometheusRulesListWithDefaults instantiates a new PrometheusRulesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusRulesListWithDefaults() *PrometheusRulesList {
	this := PrometheusRulesList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PrometheusRulesList) GetData() []PrometheusRuleListDto {
	if o == nil || IsNil(o.Data) {
		var ret []PrometheusRuleListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesList) GetDataOk() ([]PrometheusRuleListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrometheusRulesList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrometheusRuleListDto and assigns it to the Data field.
func (o *PrometheusRulesList) SetData(v []PrometheusRuleListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PrometheusRulesList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PrometheusRulesList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PrometheusRulesList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PrometheusRulesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusRulesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePrometheusRulesList struct {
	value *PrometheusRulesList
	isSet bool
}

func (v NullablePrometheusRulesList) Get() *PrometheusRulesList {
	return v.value
}

func (v *NullablePrometheusRulesList) Set(val *PrometheusRulesList) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusRulesList) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusRulesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusRulesList(val *PrometheusRulesList) *NullablePrometheusRulesList {
	return &NullablePrometheusRulesList{value: val, isSet: true}
}

func (v NullablePrometheusRulesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusRulesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


