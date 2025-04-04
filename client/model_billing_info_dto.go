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

// checks if the BillingInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfoDto{}

// BillingInfoDto struct for BillingInfoDto
type BillingInfoDto struct {
	Country NullableString `json:"country"`
	VatNumber NullableString `json:"vatNumber"`
	LegalName NullableString `json:"legalName"`
	City NullableString `json:"city"`
	BillingEmail NullableString `json:"billingEmail"`
	Address NullableString `json:"address"`
}

type _BillingInfoDto BillingInfoDto

// NewBillingInfoDto instantiates a new BillingInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoDto(country NullableString, vatNumber NullableString, legalName NullableString, city NullableString, billingEmail NullableString, address NullableString) *BillingInfoDto {
	this := BillingInfoDto{}
	this.Country = country
	this.VatNumber = vatNumber
	this.LegalName = legalName
	this.City = city
	this.BillingEmail = billingEmail
	this.Address = address
	return &this
}

// NewBillingInfoDtoWithDefaults instantiates a new BillingInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoDtoWithDefaults() *BillingInfoDto {
	this := BillingInfoDto{}
	return &this
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *BillingInfoDto) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetVatNumber returns the VatNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetVatNumber() string {
	if o == nil || o.VatNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// SetVatNumber sets field value
func (o *BillingInfoDto) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}

// GetLegalName returns the LegalName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetLegalName() string {
	if o == nil || o.LegalName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// SetLegalName sets field value
func (o *BillingInfoDto) SetLegalName(v string) {
	o.LegalName.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *BillingInfoDto) SetCity(v string) {
	o.City.Set(&v)
}

// GetBillingEmail returns the BillingEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetBillingEmail() string {
	if o == nil || o.BillingEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// SetBillingEmail sets field value
func (o *BillingInfoDto) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingInfoDto) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *BillingInfoDto) SetAddress(v string) {
	o.Address.Set(&v)
}

func (o BillingInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country.Get()
	toSerialize["vatNumber"] = o.VatNumber.Get()
	toSerialize["legalName"] = o.LegalName.Get()
	toSerialize["city"] = o.City.Get()
	toSerialize["billingEmail"] = o.BillingEmail.Get()
	toSerialize["address"] = o.Address.Get()
	return toSerialize, nil
}

func (o *BillingInfoDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
		"vatNumber",
		"legalName",
		"city",
		"billingEmail",
		"address",
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

	varBillingInfoDto := _BillingInfoDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingInfoDto)

	if err != nil {
		return err
	}

	*o = BillingInfoDto(varBillingInfoDto)

	return err
}

type NullableBillingInfoDto struct {
	value *BillingInfoDto
	isSet bool
}

func (v NullableBillingInfoDto) Get() *BillingInfoDto {
	return v.value
}

func (v *NullableBillingInfoDto) Set(val *BillingInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoDto(val *BillingInfoDto) *NullableBillingInfoDto {
	return &NullableBillingInfoDto{value: val, isSet: true}
}

func (v NullableBillingInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


