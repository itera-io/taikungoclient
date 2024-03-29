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

// checks if the InfraProductDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraProductDto{}

// InfraProductDto struct for InfraProductDto
type InfraProductDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PriceId NullableString `json:"priceId,omitempty"`
	YearlyPriceId NullableString `json:"yearlyPriceId,omitempty"`
	ProductId NullableString `json:"productId,omitempty"`
}

// NewInfraProductDto instantiates a new InfraProductDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraProductDto() *InfraProductDto {
	this := InfraProductDto{}
	return &this
}

// NewInfraProductDtoWithDefaults instantiates a new InfraProductDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraProductDtoWithDefaults() *InfraProductDto {
	this := InfraProductDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InfraProductDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraProductDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InfraProductDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InfraProductDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfraProductDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfraProductDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InfraProductDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InfraProductDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InfraProductDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InfraProductDto) UnsetName() {
	o.Name.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InfraProductDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraProductDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InfraProductDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *InfraProductDto) SetPrice(v float64) {
	o.Price = &v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfraProductDto) GetPriceId() string {
	if o == nil || IsNil(o.PriceId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfraProductDto) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// HasPriceId returns a boolean if a field has been set.
func (o *InfraProductDto) HasPriceId() bool {
	if o != nil && o.PriceId.IsSet() {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given NullableString and assigns it to the PriceId field.
func (o *InfraProductDto) SetPriceId(v string) {
	o.PriceId.Set(&v)
}
// SetPriceIdNil sets the value for PriceId to be an explicit nil
func (o *InfraProductDto) SetPriceIdNil() {
	o.PriceId.Set(nil)
}

// UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
func (o *InfraProductDto) UnsetPriceId() {
	o.PriceId.Unset()
}

// GetYearlyPriceId returns the YearlyPriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfraProductDto) GetYearlyPriceId() string {
	if o == nil || IsNil(o.YearlyPriceId.Get()) {
		var ret string
		return ret
	}
	return *o.YearlyPriceId.Get()
}

// GetYearlyPriceIdOk returns a tuple with the YearlyPriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfraProductDto) GetYearlyPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YearlyPriceId.Get(), o.YearlyPriceId.IsSet()
}

// HasYearlyPriceId returns a boolean if a field has been set.
func (o *InfraProductDto) HasYearlyPriceId() bool {
	if o != nil && o.YearlyPriceId.IsSet() {
		return true
	}

	return false
}

// SetYearlyPriceId gets a reference to the given NullableString and assigns it to the YearlyPriceId field.
func (o *InfraProductDto) SetYearlyPriceId(v string) {
	o.YearlyPriceId.Set(&v)
}
// SetYearlyPriceIdNil sets the value for YearlyPriceId to be an explicit nil
func (o *InfraProductDto) SetYearlyPriceIdNil() {
	o.YearlyPriceId.Set(nil)
}

// UnsetYearlyPriceId ensures that no value is present for YearlyPriceId, not even an explicit nil
func (o *InfraProductDto) UnsetYearlyPriceId() {
	o.YearlyPriceId.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfraProductDto) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfraProductDto) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *InfraProductDto) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *InfraProductDto) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *InfraProductDto) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *InfraProductDto) UnsetProductId() {
	o.ProductId.Unset()
}

func (o InfraProductDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraProductDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.PriceId.IsSet() {
		toSerialize["priceId"] = o.PriceId.Get()
	}
	if o.YearlyPriceId.IsSet() {
		toSerialize["yearlyPriceId"] = o.YearlyPriceId.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["productId"] = o.ProductId.Get()
	}
	return toSerialize, nil
}

type NullableInfraProductDto struct {
	value *InfraProductDto
	isSet bool
}

func (v NullableInfraProductDto) Get() *InfraProductDto {
	return v.value
}

func (v *NullableInfraProductDto) Set(val *InfraProductDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraProductDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraProductDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraProductDto(val *InfraProductDto) *NullableInfraProductDto {
	return &NullableInfraProductDto{value: val, isSet: true}
}

func (v NullableInfraProductDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraProductDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


