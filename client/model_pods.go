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

// checks if the Pods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pods{}

// Pods struct for Pods
type Pods struct {
	Data []PodListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPods instantiates a new Pods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPods() *Pods {
	this := Pods{}
	return &this
}

// NewPodsWithDefaults instantiates a new Pods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodsWithDefaults() *Pods {
	this := Pods{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pods) GetData() []PodListDto {
	if o == nil {
		var ret []PodListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pods) GetDataOk() ([]PodListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Pods) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PodListDto and assigns it to the Data field.
func (o *Pods) SetData(v []PodListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Pods) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pods) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Pods) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *Pods) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o Pods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePods struct {
	value *Pods
	isSet bool
}

func (v NullablePods) Get() *Pods {
	return v.value
}

func (v *NullablePods) Set(val *Pods) {
	v.value = val
	v.isSet = true
}

func (v NullablePods) IsSet() bool {
	return v.isSet
}

func (v *NullablePods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePods(val *Pods) *NullablePods {
	return &NullablePods{value: val, isSet: true}
}

func (v NullablePods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


