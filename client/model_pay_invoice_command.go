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

// checks if the PayInvoiceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayInvoiceCommand{}

// PayInvoiceCommand struct for PayInvoiceCommand
type PayInvoiceCommand struct {
	InvoiceId *string `json:"invoiceId,omitempty"`
}

// NewPayInvoiceCommand instantiates a new PayInvoiceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayInvoiceCommand() *PayInvoiceCommand {
	this := PayInvoiceCommand{}
	return &this
}

// NewPayInvoiceCommandWithDefaults instantiates a new PayInvoiceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayInvoiceCommandWithDefaults() *PayInvoiceCommand {
	this := PayInvoiceCommand{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *PayInvoiceCommand) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayInvoiceCommand) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PayInvoiceCommand) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *PayInvoiceCommand) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

func (o PayInvoiceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayInvoiceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	return toSerialize, nil
}

type NullablePayInvoiceCommand struct {
	value *PayInvoiceCommand
	isSet bool
}

func (v NullablePayInvoiceCommand) Get() *PayInvoiceCommand {
	return v.value
}

func (v *NullablePayInvoiceCommand) Set(val *PayInvoiceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePayInvoiceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePayInvoiceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayInvoiceCommand(val *PayInvoiceCommand) *NullablePayInvoiceCommand {
	return &NullablePayInvoiceCommand{value: val, isSet: true}
}

func (v NullablePayInvoiceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayInvoiceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


