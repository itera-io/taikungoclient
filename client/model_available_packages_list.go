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

// checks if the AvailablePackagesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePackagesList{}

// AvailablePackagesList struct for AvailablePackagesList
type AvailablePackagesList struct {
	Data []AvailablePackagesDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAvailablePackagesList instantiates a new AvailablePackagesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePackagesList() *AvailablePackagesList {
	this := AvailablePackagesList{}
	return &this
}

// NewAvailablePackagesListWithDefaults instantiates a new AvailablePackagesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePackagesListWithDefaults() *AvailablePackagesList {
	this := AvailablePackagesList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AvailablePackagesList) GetData() []AvailablePackagesDto {
	if o == nil || IsNil(o.Data) {
		var ret []AvailablePackagesDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesList) GetDataOk() ([]AvailablePackagesDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AvailablePackagesList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AvailablePackagesDto and assigns it to the Data field.
func (o *AvailablePackagesList) SetData(v []AvailablePackagesDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AvailablePackagesList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AvailablePackagesList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AvailablePackagesList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AvailablePackagesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePackagesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAvailablePackagesList struct {
	value *AvailablePackagesList
	isSet bool
}

func (v NullableAvailablePackagesList) Get() *AvailablePackagesList {
	return v.value
}

func (v *NullableAvailablePackagesList) Set(val *AvailablePackagesList) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePackagesList) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePackagesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePackagesList(val *AvailablePackagesList) *NullableAvailablePackagesList {
	return &NullableAvailablePackagesList{value: val, isSet: true}
}

func (v NullableAvailablePackagesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePackagesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


