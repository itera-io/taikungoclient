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

// checks if the ListAllRestores type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllRestores{}

// ListAllRestores struct for ListAllRestores
type ListAllRestores struct {
	Data       []CRestoreDto `json:"data,omitempty"`
	TotalCount *int32        `json:"totalCount,omitempty"`
}

// NewListAllRestores instantiates a new ListAllRestores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllRestores() *ListAllRestores {
	this := ListAllRestores{}
	return &this
}

// NewListAllRestoresWithDefaults instantiates a new ListAllRestores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllRestoresWithDefaults() *ListAllRestores {
	this := ListAllRestores{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAllRestores) GetData() []CRestoreDto {
	if o == nil {
		var ret []CRestoreDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAllRestores) GetDataOk() ([]CRestoreDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAllRestores) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CRestoreDto and assigns it to the Data field.
func (o *ListAllRestores) SetData(v []CRestoreDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListAllRestores) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllRestores) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListAllRestores) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListAllRestores) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ListAllRestores) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllRestores) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableListAllRestores struct {
	value *ListAllRestores
	isSet bool
}

func (v NullableListAllRestores) Get() *ListAllRestores {
	return v.value
}

func (v *NullableListAllRestores) Set(val *ListAllRestores) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllRestores) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllRestores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllRestores(val *ListAllRestores) *NullableListAllRestores {
	return &NullableListAllRestores{value: val, isSet: true}
}

func (v NullableListAllRestores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllRestores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
