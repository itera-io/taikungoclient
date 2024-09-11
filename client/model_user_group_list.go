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

// checks if the UserGroupList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupList{}

// UserGroupList struct for UserGroupList
type UserGroupList struct {
	Data []UserGroupDetailsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewUserGroupList instantiates a new UserGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupList() *UserGroupList {
	this := UserGroupList{}
	return &this
}

// NewUserGroupListWithDefaults instantiates a new UserGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupListWithDefaults() *UserGroupList {
	this := UserGroupList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserGroupList) GetData() []UserGroupDetailsListDto {
	if o == nil || IsNil(o.Data) {
		var ret []UserGroupDetailsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupList) GetDataOk() ([]UserGroupDetailsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserGroupList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []UserGroupDetailsListDto and assigns it to the Data field.
func (o *UserGroupList) SetData(v []UserGroupDetailsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *UserGroupList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *UserGroupList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *UserGroupList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o UserGroupList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableUserGroupList struct {
	value *UserGroupList
	isSet bool
}

func (v NullableUserGroupList) Get() *UserGroupList {
	return v.value
}

func (v *NullableUserGroupList) Set(val *UserGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupList(val *UserGroupList) *NullableUserGroupList {
	return &NullableUserGroupList{value: val, isSet: true}
}

func (v NullableUserGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


