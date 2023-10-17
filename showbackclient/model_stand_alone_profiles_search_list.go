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

// checks if the StandAloneProfilesSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfilesSearchList{}

// StandAloneProfilesSearchList struct for StandAloneProfilesSearchList
type StandAloneProfilesSearchList struct {
	Data       []CommonSearchResponseData `json:"data,omitempty"`
	TotalCount *int32                     `json:"totalCount,omitempty"`
}

// NewStandAloneProfilesSearchList instantiates a new StandAloneProfilesSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfilesSearchList() *StandAloneProfilesSearchList {
	this := StandAloneProfilesSearchList{}
	return &this
}

// NewStandAloneProfilesSearchListWithDefaults instantiates a new StandAloneProfilesSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfilesSearchListWithDefaults() *StandAloneProfilesSearchList {
	this := StandAloneProfilesSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfilesSearchList) GetData() []CommonSearchResponseData {
	if o == nil {
		var ret []CommonSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfilesSearchList) GetDataOk() ([]CommonSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StandAloneProfilesSearchList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchResponseData and assigns it to the Data field.
func (o *StandAloneProfilesSearchList) SetData(v []CommonSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *StandAloneProfilesSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *StandAloneProfilesSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *StandAloneProfilesSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o StandAloneProfilesSearchList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfilesSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableStandAloneProfilesSearchList struct {
	value *StandAloneProfilesSearchList
	isSet bool
}

func (v NullableStandAloneProfilesSearchList) Get() *StandAloneProfilesSearchList {
	return v.value
}

func (v *NullableStandAloneProfilesSearchList) Set(val *StandAloneProfilesSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfilesSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfilesSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfilesSearchList(val *StandAloneProfilesSearchList) *NullableStandAloneProfilesSearchList {
	return &NullableStandAloneProfilesSearchList{value: val, isSet: true}
}

func (v NullableStandAloneProfilesSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfilesSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
