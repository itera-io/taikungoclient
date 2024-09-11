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

// checks if the ProjectsSearchList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsSearchList{}

// ProjectsSearchList struct for ProjectsSearchList
type ProjectsSearchList struct {
	Data []CommonSearchResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewProjectsSearchList instantiates a new ProjectsSearchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsSearchList() *ProjectsSearchList {
	this := ProjectsSearchList{}
	return &this
}

// NewProjectsSearchListWithDefaults instantiates a new ProjectsSearchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsSearchListWithDefaults() *ProjectsSearchList {
	this := ProjectsSearchList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProjectsSearchList) GetData() []CommonSearchResponseData {
	if o == nil || IsNil(o.Data) {
		var ret []CommonSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsSearchList) GetDataOk() ([]CommonSearchResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectsSearchList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonSearchResponseData and assigns it to the Data field.
func (o *ProjectsSearchList) SetData(v []CommonSearchResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProjectsSearchList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsSearchList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProjectsSearchList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProjectsSearchList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ProjectsSearchList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsSearchList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProjectsSearchList struct {
	value *ProjectsSearchList
	isSet bool
}

func (v NullableProjectsSearchList) Get() *ProjectsSearchList {
	return v.value
}

func (v *NullableProjectsSearchList) Set(val *ProjectsSearchList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsSearchList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsSearchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsSearchList(val *ProjectsSearchList) *NullableProjectsSearchList {
	return &NullableProjectsSearchList{value: val, isSet: true}
}

func (v NullableProjectsSearchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsSearchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


