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

// checks if the Invoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoices{}

// Invoices struct for Invoices
type Invoices struct {
	Data       []InvoiceListDto `json:"data,omitempty"`
	TotalCount *int32           `json:"totalCount,omitempty"`
}

// NewInvoices instantiates a new Invoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoices() *Invoices {
	this := Invoices{}
	return &this
}

// NewInvoicesWithDefaults instantiates a new Invoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesWithDefaults() *Invoices {
	this := Invoices{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoices) GetData() []InvoiceListDto {
	if o == nil {
		var ret []InvoiceListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoices) GetDataOk() ([]InvoiceListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Invoices) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InvoiceListDto and assigns it to the Data field.
func (o *Invoices) SetData(v []InvoiceListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Invoices) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Invoices) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *Invoices) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o Invoices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableInvoices struct {
	value *Invoices
	isSet bool
}

func (v NullableInvoices) Get() *Invoices {
	return v.value
}

func (v *NullableInvoices) Set(val *Invoices) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoices(val *Invoices) *NullableInvoices {
	return &NullableInvoices{value: val, isSet: true}
}

func (v NullableInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
