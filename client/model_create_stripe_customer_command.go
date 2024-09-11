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

// checks if the CreateStripeCustomerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStripeCustomerCommand{}

// CreateStripeCustomerCommand struct for CreateStripeCustomerCommand
type CreateStripeCustomerCommand struct {
	LegalName *string `json:"legalName,omitempty"`
	BillingEmail *string `json:"billingEmail,omitempty"`
	Country *string `json:"country,omitempty"`
	Address *string `json:"address,omitempty"`
	City *string `json:"city,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewCreateStripeCustomerCommand instantiates a new CreateStripeCustomerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStripeCustomerCommand() *CreateStripeCustomerCommand {
	this := CreateStripeCustomerCommand{}
	return &this
}

// NewCreateStripeCustomerCommandWithDefaults instantiates a new CreateStripeCustomerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStripeCustomerCommandWithDefaults() *CreateStripeCustomerCommand {
	this := CreateStripeCustomerCommand{}
	return &this
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *CreateStripeCustomerCommand) SetLegalName(v string) {
	o.LegalName = &v
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail) {
		var ret string
		return ret
	}
	return *o.BillingEmail
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetBillingEmailOk() (*string, bool) {
	if o == nil || IsNil(o.BillingEmail) {
		return nil, false
	}
	return o.BillingEmail, true
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasBillingEmail() bool {
	if o != nil && !IsNil(o.BillingEmail) {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.
func (o *CreateStripeCustomerCommand) SetBillingEmail(v string) {
	o.BillingEmail = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CreateStripeCustomerCommand) SetCountry(v string) {
	o.Country = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateStripeCustomerCommand) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CreateStripeCustomerCommand) SetCity(v string) {
	o.City = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CreateStripeCustomerCommand) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStripeCustomerCommand) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CreateStripeCustomerCommand) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o CreateStripeCustomerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStripeCustomerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalName) {
		toSerialize["legalName"] = o.LegalName
	}
	if !IsNil(o.BillingEmail) {
		toSerialize["billingEmail"] = o.BillingEmail
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

type NullableCreateStripeCustomerCommand struct {
	value *CreateStripeCustomerCommand
	isSet bool
}

func (v NullableCreateStripeCustomerCommand) Get() *CreateStripeCustomerCommand {
	return v.value
}

func (v *NullableCreateStripeCustomerCommand) Set(val *CreateStripeCustomerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStripeCustomerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStripeCustomerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStripeCustomerCommand(val *CreateStripeCustomerCommand) *NullableCreateStripeCustomerCommand {
	return &NullableCreateStripeCustomerCommand{value: val, isSet: true}
}

func (v NullableCreateStripeCustomerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStripeCustomerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


