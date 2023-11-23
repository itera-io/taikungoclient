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

// checks if the PvcSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PvcSearchList{}

// PvcSearchList struct for PvcSearchList
type PvcSearchList struct {
	Data []CommonSearchKubernetesResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPvcSearchList instantiates a new PvcSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPvcSearchList() *PvcSearchList {
	this := PvcSearchList{}
	return &this
}

// NewPvcSearchListWithDefaults instantiates a new PvcSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvcSearchListWithDefaults() *PvcSearchList {
	this := PvcSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PvcSearchList) GetData() []CommonSearchKubernetesResponseData {
	if o == nil {
		var ret []CommonSearchKubernetesResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PvcSearchList) GetDataOk() ([]CommonSearchKubernetesResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PvcSearchList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchKubernetesResponseData and assigns it to the Data field.
func (o *PvcSearchList) SetData(v []CommonSearchKubernetesResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PvcSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PvcSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PvcSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PvcSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PvcSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PvcSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePvcSearchList struct {
	value *PvcSearchList
	isSet bool
}

func (v NullablePvcSearchList) Get() *PvcSearchList {
	return v.value
}

func (v *NullablePvcSearchList) Set(val *PvcSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullablePvcSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullablePvcSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePvcSearchList(val *PvcSearchList) *NullablePvcSearchList {
	return &NullablePvcSearchList{value: val, isSet: true}
}

func (v NullablePvcSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePvcSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


