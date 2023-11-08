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

// checks if the PrometheusDashboardListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusDashboardListDto{}

// PrometheusDashboardListDto struct for PrometheusDashboardListDto
type PrometheusDashboardListDto struct {
	CategoryName NullableString           `json:"categoryName,omitempty"`
	Data         []PrometheusDashboardDto `json:"data,omitempty"`
}

// NewPrometheusDashboardListDto instantiates a new PrometheusDashboardListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusDashboardListDto() *PrometheusDashboardListDto {
	this := PrometheusDashboardListDto{}
	return &this
}

// NewPrometheusDashboardListDtoWithDefaults instantiates a new PrometheusDashboardListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusDashboardListDtoWithDefaults() *PrometheusDashboardListDto {
	this := PrometheusDashboardListDto{}
	return &this
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardListDto) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryName.Get()
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardListDto) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryName.Get(), o.CategoryName.IsSet()
}

// HasCategoryName returns a boolean if a field has been set.
func (o *PrometheusDashboardListDto) HasCategoryName() bool {
	if o != nil && o.CategoryName.IsSet() {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given NullableString and assigns it to the CategoryName field.
func (o *PrometheusDashboardListDto) SetCategoryName(v string) {
	o.CategoryName.Set(&v)
}

// SetCategoryNameNil sets the value for CategoryName to be an explicit nil
func (o *PrometheusDashboardListDto) SetCategoryNameNil() {
	o.CategoryName.Set(nil)
}

// UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
func (o *PrometheusDashboardListDto) UnsetCategoryName() {
	o.CategoryName.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardListDto) GetData() []PrometheusDashboardDto {
	if o == nil {
		var ret []PrometheusDashboardDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardListDto) GetDataOk() ([]PrometheusDashboardDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrometheusDashboardListDto) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrometheusDashboardDto and assigns it to the Data field.
func (o *PrometheusDashboardListDto) SetData(v []PrometheusDashboardDto) {
	o.Data = v
}

func (o PrometheusDashboardListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusDashboardListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryName.IsSet() {
		toSerialize["categoryName"] = o.CategoryName.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePrometheusDashboardListDto struct {
	value *PrometheusDashboardListDto
	isSet bool
}

func (v NullablePrometheusDashboardListDto) Get() *PrometheusDashboardListDto {
	return v.value
}

func (v *NullablePrometheusDashboardListDto) Set(val *PrometheusDashboardListDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusDashboardListDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusDashboardListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusDashboardListDto(val *PrometheusDashboardListDto) *NullablePrometheusDashboardListDto {
	return &NullablePrometheusDashboardListDto{value: val, isSet: true}
}

func (v NullablePrometheusDashboardListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusDashboardListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
