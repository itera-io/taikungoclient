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

// checks if the OrganizationDropdownDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDropdownDto{}

// OrganizationDropdownDto struct for OrganizationDropdownDto
type OrganizationDropdownDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	IsInfra *bool `json:"isInfra,omitempty"`
	IsBound *bool `json:"isBound,omitempty"`
	DiscountRate *float64 `json:"discountRate,omitempty"`
}

// NewOrganizationDropdownDto instantiates a new OrganizationDropdownDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDropdownDto() *OrganizationDropdownDto {
	this := OrganizationDropdownDto{}
	return &this
}

// NewOrganizationDropdownDtoWithDefaults instantiates a new OrganizationDropdownDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDropdownDtoWithDefaults() *OrganizationDropdownDto {
	this := OrganizationDropdownDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationDropdownDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDropdownDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationDropdownDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganizationDropdownDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationDropdownDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDropdownDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationDropdownDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationDropdownDto) SetName(v string) {
	o.Name = &v
}

// GetIsInfra returns the IsInfra field value if set, zero value otherwise.
func (o *OrganizationDropdownDto) GetIsInfra() bool {
	if o == nil || IsNil(o.IsInfra) {
		var ret bool
		return ret
	}
	return *o.IsInfra
}

// GetIsInfraOk returns a tuple with the IsInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDropdownDto) GetIsInfraOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInfra) {
		return nil, false
	}
	return o.IsInfra, true
}

// HasIsInfra returns a boolean if a field has been set.
func (o *OrganizationDropdownDto) HasIsInfra() bool {
	if o != nil && !IsNil(o.IsInfra) {
		return true
	}

	return false
}

// SetIsInfra gets a reference to the given bool and assigns it to the IsInfra field.
func (o *OrganizationDropdownDto) SetIsInfra(v bool) {
	o.IsInfra = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *OrganizationDropdownDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDropdownDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *OrganizationDropdownDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *OrganizationDropdownDto) SetIsBound(v bool) {
	o.IsBound = &v
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise.
func (o *OrganizationDropdownDto) GetDiscountRate() float64 {
	if o == nil || IsNil(o.DiscountRate) {
		var ret float64
		return ret
	}
	return *o.DiscountRate
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDropdownDto) GetDiscountRateOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountRate) {
		return nil, false
	}
	return o.DiscountRate, true
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *OrganizationDropdownDto) HasDiscountRate() bool {
	if o != nil && !IsNil(o.DiscountRate) {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given float64 and assigns it to the DiscountRate field.
func (o *OrganizationDropdownDto) SetDiscountRate(v float64) {
	o.DiscountRate = &v
}

func (o OrganizationDropdownDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDropdownDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsInfra) {
		toSerialize["isInfra"] = o.IsInfra
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	if !IsNil(o.DiscountRate) {
		toSerialize["discountRate"] = o.DiscountRate
	}
	return toSerialize, nil
}

type NullableOrganizationDropdownDto struct {
	value *OrganizationDropdownDto
	isSet bool
}

func (v NullableOrganizationDropdownDto) Get() *OrganizationDropdownDto {
	return v.value
}

func (v *NullableOrganizationDropdownDto) Set(val *OrganizationDropdownDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDropdownDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDropdownDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDropdownDto(val *OrganizationDropdownDto) *NullableOrganizationDropdownDto {
	return &NullableOrganizationDropdownDto{value: val, isSet: true}
}

func (v NullableOrganizationDropdownDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDropdownDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


