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

// checks if the AlertingEmailDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingEmailDto{}

// AlertingEmailDto struct for AlertingEmailDto
type AlertingEmailDto struct {
	Email NullableString `json:"email"`
}

type _AlertingEmailDto AlertingEmailDto

// NewAlertingEmailDto instantiates a new AlertingEmailDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingEmailDto(email NullableString) *AlertingEmailDto {
	this := AlertingEmailDto{}
	this.Email = email
	return &this
}

// NewAlertingEmailDtoWithDefaults instantiates a new AlertingEmailDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingEmailDtoWithDefaults() *AlertingEmailDto {
	this := AlertingEmailDto{}
	return &this
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AlertingEmailDto) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingEmailDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *AlertingEmailDto) SetEmail(v string) {
	o.Email.Set(&v)
}

func (o AlertingEmailDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingEmailDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email.Get()
	return toSerialize, nil
}

func (o *AlertingEmailDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varAlertingEmailDto := _AlertingEmailDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertingEmailDto)

	if err != nil {
		return err
	}

	*o = AlertingEmailDto(varAlertingEmailDto)

	return err
}

type NullableAlertingEmailDto struct {
	value *AlertingEmailDto
	isSet bool
}

func (v NullableAlertingEmailDto) Get() *AlertingEmailDto {
	return v.value
}

func (v *NullableAlertingEmailDto) Set(val *AlertingEmailDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingEmailDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingEmailDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingEmailDto(val *AlertingEmailDto) *NullableAlertingEmailDto {
	return &NullableAlertingEmailDto{value: val, isSet: true}
}

func (v NullableAlertingEmailDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingEmailDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


