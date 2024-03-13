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

// checks if the CatalogAppList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogAppList{}

// CatalogAppList struct for CatalogAppList
type CatalogAppList struct {
	Data []CatalogAppListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewCatalogAppList instantiates a new CatalogAppList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogAppList() *CatalogAppList {
	this := CatalogAppList{}
	return &this
}

// NewCatalogAppListWithDefaults instantiates a new CatalogAppList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogAppListWithDefaults() *CatalogAppList {
	this := CatalogAppList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppList) GetData() []CatalogAppListDto {
	if o == nil {
		var ret []CatalogAppListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppList) GetDataOk() ([]CatalogAppListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CatalogAppList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CatalogAppListDto and assigns it to the Data field.
func (o *CatalogAppList) SetData(v []CatalogAppListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CatalogAppList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CatalogAppList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *CatalogAppList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o CatalogAppList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogAppList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableCatalogAppList struct {
	value *CatalogAppList
	isSet bool
}

func (v NullableCatalogAppList) Get() *CatalogAppList {
	return v.value
}

func (v *NullableCatalogAppList) Set(val *CatalogAppList) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogAppList) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogAppList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogAppList(val *CatalogAppList) *NullableCatalogAppList {
	return &NullableCatalogAppList{value: val, isSet: true}
}

func (v NullableCatalogAppList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogAppList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

