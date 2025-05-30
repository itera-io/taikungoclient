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
	"bytes"
	"fmt"
)

// checks if the StripeSubscriptionItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeSubscriptionItemDto{}

// StripeSubscriptionItemDto struct for StripeSubscriptionItemDto
type StripeSubscriptionItemDto struct {
	SubscriptionItemId NullableString `json:"subscriptionItemId"`
	PriceId NullableString `json:"priceId"`
	ProductId NullableString `json:"productId"`
}

type _StripeSubscriptionItemDto StripeSubscriptionItemDto

// NewStripeSubscriptionItemDto instantiates a new StripeSubscriptionItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeSubscriptionItemDto(subscriptionItemId NullableString, priceId NullableString, productId NullableString) *StripeSubscriptionItemDto {
	this := StripeSubscriptionItemDto{}
	this.SubscriptionItemId = subscriptionItemId
	this.PriceId = priceId
	this.ProductId = productId
	return &this
}

// NewStripeSubscriptionItemDtoWithDefaults instantiates a new StripeSubscriptionItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeSubscriptionItemDtoWithDefaults() *StripeSubscriptionItemDto {
	this := StripeSubscriptionItemDto{}
	return &this
}

// GetSubscriptionItemId returns the SubscriptionItemId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StripeSubscriptionItemDto) GetSubscriptionItemId() string {
	if o == nil || o.SubscriptionItemId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionItemId.Get()
}

// GetSubscriptionItemIdOk returns a tuple with the SubscriptionItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeSubscriptionItemDto) GetSubscriptionItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionItemId.Get(), o.SubscriptionItemId.IsSet()
}

// SetSubscriptionItemId sets field value
func (o *StripeSubscriptionItemDto) SetSubscriptionItemId(v string) {
	o.SubscriptionItemId.Set(&v)
}

// GetPriceId returns the PriceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StripeSubscriptionItemDto) GetPriceId() string {
	if o == nil || o.PriceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeSubscriptionItemDto) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// SetPriceId sets field value
func (o *StripeSubscriptionItemDto) SetPriceId(v string) {
	o.PriceId.Set(&v)
}

// GetProductId returns the ProductId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StripeSubscriptionItemDto) GetProductId() string {
	if o == nil || o.ProductId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeSubscriptionItemDto) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// SetProductId sets field value
func (o *StripeSubscriptionItemDto) SetProductId(v string) {
	o.ProductId.Set(&v)
}

func (o StripeSubscriptionItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeSubscriptionItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionItemId"] = o.SubscriptionItemId.Get()
	toSerialize["priceId"] = o.PriceId.Get()
	toSerialize["productId"] = o.ProductId.Get()
	return toSerialize, nil
}

func (o *StripeSubscriptionItemDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionItemId",
		"priceId",
		"productId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStripeSubscriptionItemDto := _StripeSubscriptionItemDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStripeSubscriptionItemDto)

	if err != nil {
		return err
	}

	*o = StripeSubscriptionItemDto(varStripeSubscriptionItemDto)

	return err
}

type NullableStripeSubscriptionItemDto struct {
	value *StripeSubscriptionItemDto
	isSet bool
}

func (v NullableStripeSubscriptionItemDto) Get() *StripeSubscriptionItemDto {
	return v.value
}

func (v *NullableStripeSubscriptionItemDto) Set(val *StripeSubscriptionItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeSubscriptionItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeSubscriptionItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeSubscriptionItemDto(val *StripeSubscriptionItemDto) *NullableStripeSubscriptionItemDto {
	return &NullableStripeSubscriptionItemDto{value: val, isSet: true}
}

func (v NullableStripeSubscriptionItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeSubscriptionItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


