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

// checks if the PublicImageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicImageList{}

// PublicImageList struct for PublicImageList
type PublicImageList struct {
	Data []CommonStringBasedDropdownDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPublicImageList instantiates a new PublicImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicImageList() *PublicImageList {
	this := PublicImageList{}
	return &this
}

// NewPublicImageListWithDefaults instantiates a new PublicImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicImageListWithDefaults() *PublicImageList {
	this := PublicImageList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicImageList) GetData() []CommonStringBasedDropdownDto {
	if o == nil {
		var ret []CommonStringBasedDropdownDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicImageList) GetDataOk() ([]CommonStringBasedDropdownDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PublicImageList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonStringBasedDropdownDto and assigns it to the Data field.
func (o *PublicImageList) SetData(v []CommonStringBasedDropdownDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PublicImageList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImageList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PublicImageList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PublicImageList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PublicImageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicImageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePublicImageList struct {
	value *PublicImageList
	isSet bool
}

func (v NullablePublicImageList) Get() *PublicImageList {
	return v.value
}

func (v *NullablePublicImageList) Set(val *PublicImageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicImageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicImageList(val *PublicImageList) *NullablePublicImageList {
	return &NullablePublicImageList{value: val, isSet: true}
}

func (v NullablePublicImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

