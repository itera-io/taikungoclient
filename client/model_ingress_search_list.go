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

// checks if the IngressSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngressSearchList{}

// IngressSearchList struct for IngressSearchList
type IngressSearchList struct {
	Data []CommonSearchKubernetesResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewIngressSearchList instantiates a new IngressSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngressSearchList() *IngressSearchList {
	this := IngressSearchList{}
	return &this
}

// NewIngressSearchListWithDefaults instantiates a new IngressSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngressSearchListWithDefaults() *IngressSearchList {
	this := IngressSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngressSearchList) GetData() []CommonSearchKubernetesResponseData {
	if o == nil {
		var ret []CommonSearchKubernetesResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngressSearchList) GetDataOk() ([]CommonSearchKubernetesResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IngressSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchKubernetesResponseData and assigns it to the Data field.
func (o *IngressSearchList) SetData(v []CommonSearchKubernetesResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *IngressSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngressSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *IngressSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *IngressSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o IngressSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngressSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableIngressSearchList struct {
	value *IngressSearchList
	isSet bool
}

func (v NullableIngressSearchList) Get() *IngressSearchList {
	return v.value
}

func (v *NullableIngressSearchList) Set(val *IngressSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableIngressSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableIngressSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngressSearchList(val *IngressSearchList) *NullableIngressSearchList {
	return &NullableIngressSearchList{value: val, isSet: true}
}

func (v NullableIngressSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngressSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


