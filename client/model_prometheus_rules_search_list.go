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

// checks if the PrometheusRulesSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusRulesSearchList{}

// PrometheusRulesSearchList struct for PrometheusRulesSearchList
type PrometheusRulesSearchList struct {
	Data []PrometheusRulesSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPrometheusRulesSearchList instantiates a new PrometheusRulesSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusRulesSearchList() *PrometheusRulesSearchList {
	this := PrometheusRulesSearchList{}
	return &this
}

// NewPrometheusRulesSearchListWithDefaults instantiates a new PrometheusRulesSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusRulesSearchListWithDefaults() *PrometheusRulesSearchList {
	this := PrometheusRulesSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PrometheusRulesSearchList) GetData() []PrometheusRulesSearchResponseData {
	if o == nil || IsNil(o.Data) {
		var ret []PrometheusRulesSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesSearchList) GetDataOk() ([]PrometheusRulesSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrometheusRulesSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrometheusRulesSearchResponseData and assigns it to the Data field.
func (o *PrometheusRulesSearchList) SetData(v []PrometheusRulesSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PrometheusRulesSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRulesSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PrometheusRulesSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PrometheusRulesSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PrometheusRulesSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusRulesSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePrometheusRulesSearchList struct {
	value *PrometheusRulesSearchList
	isSet bool
}

func (v NullablePrometheusRulesSearchList) Get() *PrometheusRulesSearchList {
	return v.value
}

func (v *NullablePrometheusRulesSearchList) Set(val *PrometheusRulesSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusRulesSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusRulesSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusRulesSearchList(val *PrometheusRulesSearchList) *NullablePrometheusRulesSearchList {
	return &NullablePrometheusRulesSearchList{value: val, isSet: true}
}

func (v NullablePrometheusRulesSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusRulesSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


