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
	LegalName    NullableString `json:"legalName,omitempty"`
	BillingEmail NullableString `json:"billingEmail,omitempty"`
	Country      NullableString `json:"country,omitempty"`
	Address      NullableString `json:"address,omitempty"`
	City         NullableString `json:"city,omitempty"`
	VatNumber    NullableString `json:"vatNumber,omitempty"`
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

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetLegalName() string {
	if o == nil || IsNil(o.LegalName.Get()) {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *CreateStripeCustomerCommand) SetLegalName(v string) {
	o.LegalName.Set(&v)
}

// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *CreateStripeCustomerCommand) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasBillingEmail() bool {
	if o != nil && o.BillingEmail.IsSet() {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given NullableString and assigns it to the BillingEmail field.
func (o *CreateStripeCustomerCommand) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}

// SetBillingEmailNil sets the value for BillingEmail to be an explicit nil
func (o *CreateStripeCustomerCommand) SetBillingEmailNil() {
	o.BillingEmail.Set(nil)
}

// UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetBillingEmail() {
	o.BillingEmail.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *CreateStripeCustomerCommand) SetCountry(v string) {
	o.Country.Set(&v)
}

// SetCountryNil sets the value for Country to be an explicit nil
func (o *CreateStripeCustomerCommand) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetCountry() {
	o.Country.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *CreateStripeCustomerCommand) SetAddress(v string) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *CreateStripeCustomerCommand) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetAddress() {
	o.Address.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *CreateStripeCustomerCommand) SetCity(v string) {
	o.City.Set(&v)
}

// SetCityNil sets the value for City to be an explicit nil
func (o *CreateStripeCustomerCommand) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetCity() {
	o.City.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateStripeCustomerCommand) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateStripeCustomerCommand) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CreateStripeCustomerCommand) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *CreateStripeCustomerCommand) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}

// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *CreateStripeCustomerCommand) SetVatNumberNil() {
	o.VatNumber.Set(nil)
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *CreateStripeCustomerCommand) UnsetVatNumber() {
	o.VatNumber.Unset()
}

func (o CreateStripeCustomerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStripeCustomerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LegalName.IsSet() {
		toSerialize["legalName"] = o.LegalName.Get()
	}
	if o.BillingEmail.IsSet() {
		toSerialize["billingEmail"] = o.BillingEmail.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vatNumber"] = o.VatNumber.Get()
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
