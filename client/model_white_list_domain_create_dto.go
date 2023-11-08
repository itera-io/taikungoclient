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

// checks if the WhiteListDomainCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteListDomainCreateDto{}

// WhiteListDomainCreateDto struct for WhiteListDomainCreateDto
type WhiteListDomainCreateDto struct {
	Name NullableString `json:"name,omitempty"`
}

// NewWhiteListDomainCreateDto instantiates a new WhiteListDomainCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteListDomainCreateDto() *WhiteListDomainCreateDto {
	this := WhiteListDomainCreateDto{}
	return &this
}

// NewWhiteListDomainCreateDtoWithDefaults instantiates a new WhiteListDomainCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteListDomainCreateDtoWithDefaults() *WhiteListDomainCreateDto {
	this := WhiteListDomainCreateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteListDomainCreateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteListDomainCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WhiteListDomainCreateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WhiteListDomainCreateDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *WhiteListDomainCreateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WhiteListDomainCreateDto) UnsetName() {
	o.Name.Unset()
}

func (o WhiteListDomainCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteListDomainCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableWhiteListDomainCreateDto struct {
	value *WhiteListDomainCreateDto
	isSet bool
}

func (v NullableWhiteListDomainCreateDto) Get() *WhiteListDomainCreateDto {
	return v.value
}

func (v *NullableWhiteListDomainCreateDto) Set(val *WhiteListDomainCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteListDomainCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteListDomainCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteListDomainCreateDto(val *WhiteListDomainCreateDto) *NullableWhiteListDomainCreateDto {
	return &NullableWhiteListDomainCreateDto{value: val, isSet: true}
}

func (v NullableWhiteListDomainCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteListDomainCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
