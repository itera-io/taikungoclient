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

// checks if the DaemonSets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaemonSets{}

// DaemonSets struct for DaemonSets
type DaemonSets struct {
	Data []DaemonSetDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewDaemonSets instantiates a new DaemonSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaemonSets() *DaemonSets {
	this := DaemonSets{}
	return &this
}

// NewDaemonSetsWithDefaults instantiates a new DaemonSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaemonSetsWithDefaults() *DaemonSets {
	this := DaemonSets{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DaemonSets) GetData() []DaemonSetDto {
	if o == nil {
		var ret []DaemonSetDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DaemonSets) GetDataOk() ([]DaemonSetDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DaemonSets) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DaemonSetDto and assigns it to the Data field.
func (o *DaemonSets) SetData(v []DaemonSetDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DaemonSets) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSets) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DaemonSets) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *DaemonSets) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o DaemonSets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaemonSets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableDaemonSets struct {
	value *DaemonSets
	isSet bool
}

func (v NullableDaemonSets) Get() *DaemonSets {
	return v.value
}

func (v *NullableDaemonSets) Set(val *DaemonSets) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonSets) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonSets(val *DaemonSets) *NullableDaemonSets {
	return &NullableDaemonSets{value: val, isSet: true}
}

func (v NullableDaemonSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


