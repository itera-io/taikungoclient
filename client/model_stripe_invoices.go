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

// checks if the StripeInvoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeInvoices{}

// StripeInvoices struct for StripeInvoices
type StripeInvoices struct {
	Data []StripeInvoiceListDto `json:"data,omitempty"`
}

// NewStripeInvoices instantiates a new StripeInvoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeInvoices() *StripeInvoices {
	this := StripeInvoices{}
	return &this
}

// NewStripeInvoicesWithDefaults instantiates a new StripeInvoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeInvoicesWithDefaults() *StripeInvoices {
	this := StripeInvoices{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StripeInvoices) GetData() []StripeInvoiceListDto {
	if o == nil {
		var ret []StripeInvoiceListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeInvoices) GetDataOk() ([]StripeInvoiceListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StripeInvoices) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []StripeInvoiceListDto and assigns it to the Data field.
func (o *StripeInvoices) SetData(v []StripeInvoiceListDto) {
	o.Data = v
}

func (o StripeInvoices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeInvoices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStripeInvoices struct {
	value *StripeInvoices
	isSet bool
}

func (v NullableStripeInvoices) Get() *StripeInvoices {
	return v.value
}

func (v *NullableStripeInvoices) Set(val *StripeInvoices) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeInvoices(val *StripeInvoices) *NullableStripeInvoices {
	return &NullableStripeInvoices{value: val, isSet: true}
}

func (v NullableStripeInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


