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

// checks if the PartnersList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnersList{}

// PartnersList struct for PartnersList
type PartnersList struct {
	Data []PartnerDetailsDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPartnersList instantiates a new PartnersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnersList() *PartnersList {
	this := PartnersList{}
	return &this
}

// NewPartnersListWithDefaults instantiates a new PartnersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnersListWithDefaults() *PartnersList {
	this := PartnersList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PartnersList) GetData() []PartnerDetailsDto {
	if o == nil || IsNil(o.Data) {
		var ret []PartnerDetailsDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersList) GetDataOk() ([]PartnerDetailsDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PartnersList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PartnerDetailsDto and assigns it to the Data field.
func (o *PartnersList) SetData(v []PartnerDetailsDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PartnersList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PartnersList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PartnersList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PartnersList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnersList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePartnersList struct {
	value *PartnersList
	isSet bool
}

func (v NullablePartnersList) Get() *PartnersList {
	return v.value
}

func (v *NullablePartnersList) Set(val *PartnersList) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnersList) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnersList(val *PartnersList) *NullablePartnersList {
	return &NullablePartnersList{value: val, isSet: true}
}

func (v NullablePartnersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


