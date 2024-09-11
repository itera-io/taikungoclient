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

// checks if the ProjectAppList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAppList{}

// ProjectAppList struct for ProjectAppList
type ProjectAppList struct {
	Data []InstanceAppListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewProjectAppList instantiates a new ProjectAppList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAppList() *ProjectAppList {
	this := ProjectAppList{}
	return &this
}

// NewProjectAppListWithDefaults instantiates a new ProjectAppList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAppListWithDefaults() *ProjectAppList {
	this := ProjectAppList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProjectAppList) GetData() []InstanceAppListDto {
	if o == nil || IsNil(o.Data) {
		var ret []InstanceAppListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppList) GetDataOk() ([]InstanceAppListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectAppList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InstanceAppListDto and assigns it to the Data field.
func (o *ProjectAppList) SetData(v []InstanceAppListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProjectAppList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAppList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProjectAppList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProjectAppList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ProjectAppList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAppList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProjectAppList struct {
	value *ProjectAppList
	isSet bool
}

func (v NullableProjectAppList) Get() *ProjectAppList {
	return v.value
}

func (v *NullableProjectAppList) Set(val *ProjectAppList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAppList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAppList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAppList(val *ProjectAppList) *NullableProjectAppList {
	return &NullableProjectAppList{value: val, isSet: true}
}

func (v NullableProjectAppList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAppList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


