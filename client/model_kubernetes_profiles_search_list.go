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

// checks if the KubernetesProfilesSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesProfilesSearchList{}

// KubernetesProfilesSearchList struct for KubernetesProfilesSearchList
type KubernetesProfilesSearchList struct {
	Data []CommonSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewKubernetesProfilesSearchList instantiates a new KubernetesProfilesSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesProfilesSearchList() *KubernetesProfilesSearchList {
	this := KubernetesProfilesSearchList{}
	return &this
}

// NewKubernetesProfilesSearchListWithDefaults instantiates a new KubernetesProfilesSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesProfilesSearchListWithDefaults() *KubernetesProfilesSearchList {
	this := KubernetesProfilesSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesSearchList) GetData() []CommonSearchResponseData {
	if o == nil {
		var ret []CommonSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesSearchList) GetDataOk() ([]CommonSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KubernetesProfilesSearchList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchResponseData and assigns it to the Data field.
func (o *KubernetesProfilesSearchList) SetData(v []CommonSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *KubernetesProfilesSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *KubernetesProfilesSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *KubernetesProfilesSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o KubernetesProfilesSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesProfilesSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableKubernetesProfilesSearchList struct {
	value *KubernetesProfilesSearchList
	isSet bool
}

func (v NullableKubernetesProfilesSearchList) Get() *KubernetesProfilesSearchList {
	return v.value
}

func (v *NullableKubernetesProfilesSearchList) Set(val *KubernetesProfilesSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesProfilesSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesProfilesSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesProfilesSearchList(val *KubernetesProfilesSearchList) *NullableKubernetesProfilesSearchList {
	return &NullableKubernetesProfilesSearchList{value: val, isSet: true}
}

func (v NullableKubernetesProfilesSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesProfilesSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


