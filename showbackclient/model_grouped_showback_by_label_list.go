/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the GroupedShowbackByLabelList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedShowbackByLabelList{}

// GroupedShowbackByLabelList struct for GroupedShowbackByLabelList
type GroupedShowbackByLabelList struct {
	ByLabelValues []GroupedShowbackSummaryByLabelInfo `json:"byLabelValues,omitempty"`
}

// NewGroupedShowbackByLabelList instantiates a new GroupedShowbackByLabelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedShowbackByLabelList() *GroupedShowbackByLabelList {
	this := GroupedShowbackByLabelList{}
	return &this
}

// NewGroupedShowbackByLabelListWithDefaults instantiates a new GroupedShowbackByLabelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedShowbackByLabelListWithDefaults() *GroupedShowbackByLabelList {
	this := GroupedShowbackByLabelList{}
	return &this
}

// GetByLabelValues returns the ByLabelValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedShowbackByLabelList) GetByLabelValues() []GroupedShowbackSummaryByLabelInfo {
	if o == nil {
		var ret []GroupedShowbackSummaryByLabelInfo
		return ret
	}
	return o.ByLabelValues
}

// GetByLabelValuesOk returns a tuple with the ByLabelValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedShowbackByLabelList) GetByLabelValuesOk() ([]GroupedShowbackSummaryByLabelInfo, bool) {
	if o == nil || IsNil(o.ByLabelValues) {
		return nil, false
	}
	return o.ByLabelValues, true
}

// HasByLabelValues returns a boolean if a field has been set.
func (o *GroupedShowbackByLabelList) HasByLabelValues() bool {
	if o != nil && IsNil(o.ByLabelValues) {
		return true
	}

	return false
}

// SetByLabelValues gets a reference to the given []GroupedShowbackSummaryByLabelInfo and assigns it to the ByLabelValues field.
func (o *GroupedShowbackByLabelList) SetByLabelValues(v []GroupedShowbackSummaryByLabelInfo) {
	o.ByLabelValues = v
}

func (o GroupedShowbackByLabelList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedShowbackByLabelList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ByLabelValues != nil {
		toSerialize["byLabelValues"] = o.ByLabelValues
	}
	return toSerialize, nil
}

type NullableGroupedShowbackByLabelList struct {
	value *GroupedShowbackByLabelList
	isSet bool
}

func (v NullableGroupedShowbackByLabelList) Get() *GroupedShowbackByLabelList {
	return v.value
}

func (v *NullableGroupedShowbackByLabelList) Set(val *GroupedShowbackByLabelList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedShowbackByLabelList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedShowbackByLabelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedShowbackByLabelList(val *GroupedShowbackByLabelList) *NullableGroupedShowbackByLabelList {
	return &NullableGroupedShowbackByLabelList{value: val, isSet: true}
}

func (v NullableGroupedShowbackByLabelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedShowbackByLabelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
