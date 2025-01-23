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

// checks if the PartnerEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerEntity{}

// PartnerEntity struct for PartnerEntity
type PartnerEntity struct {
	PartnerId *int32 `json:"partnerId,omitempty"`
	PartnerName *string `json:"partnerName,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
}

// NewPartnerEntity instantiates a new PartnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEntity() *PartnerEntity {
	this := PartnerEntity{}
	return &this
}

// NewPartnerEntityWithDefaults instantiates a new PartnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEntityWithDefaults() *PartnerEntity {
	this := PartnerEntity{}
	return &this
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *PartnerEntity) GetPartnerId() int32 {
	if o == nil || IsNil(o.PartnerId) {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEntity) GetPartnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *PartnerEntity) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *PartnerEntity) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *PartnerEntity) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName) {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEntity) GetPartnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerName) {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *PartnerEntity) HasPartnerName() bool {
	if o != nil && !IsNil(o.PartnerName) {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *PartnerEntity) SetPartnerName(v string) {
	o.PartnerName = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerEntity) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerEntity) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *PartnerEntity) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *PartnerEntity) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *PartnerEntity) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *PartnerEntity) UnsetLogo() {
	o.Logo.Unset()
}

func (o PartnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !IsNil(o.PartnerName) {
		toSerialize["partnerName"] = o.PartnerName
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	return toSerialize, nil
}

type NullablePartnerEntity struct {
	value *PartnerEntity
	isSet bool
}

func (v NullablePartnerEntity) Get() *PartnerEntity {
	return v.value
}

func (v *NullablePartnerEntity) Set(val *PartnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEntity(val *PartnerEntity) *NullablePartnerEntity {
	return &NullablePartnerEntity{value: val, isSet: true}
}

func (v NullablePartnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


