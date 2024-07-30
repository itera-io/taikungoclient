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

// checks if the AdminProjectsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminProjectsList{}

// AdminProjectsList struct for AdminProjectsList
type AdminProjectsList struct {
	Data []AdminProjectsResponseData `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAdminProjectsList instantiates a new AdminProjectsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminProjectsList() *AdminProjectsList {
	this := AdminProjectsList{}
	return &this
}

// NewAdminProjectsListWithDefaults instantiates a new AdminProjectsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminProjectsListWithDefaults() *AdminProjectsList {
	this := AdminProjectsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminProjectsList) GetData() []AdminProjectsResponseData {
	if o == nil {
		var ret []AdminProjectsResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminProjectsList) GetDataOk() ([]AdminProjectsResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdminProjectsList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdminProjectsResponseData and assigns it to the Data field.
func (o *AdminProjectsList) SetData(v []AdminProjectsResponseData) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AdminProjectsList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminProjectsList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AdminProjectsList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AdminProjectsList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AdminProjectsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminProjectsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAdminProjectsList struct {
	value *AdminProjectsList
	isSet bool
}

func (v NullableAdminProjectsList) Get() *AdminProjectsList {
	return v.value
}

func (v *NullableAdminProjectsList) Set(val *AdminProjectsList) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminProjectsList) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminProjectsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminProjectsList(val *AdminProjectsList) *NullableAdminProjectsList {
	return &NullableAdminProjectsList{value: val, isSet: true}
}

func (v NullableAdminProjectsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminProjectsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


