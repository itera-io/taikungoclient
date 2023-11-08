/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the ShowbackLabelCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowbackLabelCreateDto{}

// ShowbackLabelCreateDto struct for ShowbackLabelCreateDto
type ShowbackLabelCreateDto struct {
	Label NullableString `json:"label,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewShowbackLabelCreateDto instantiates a new ShowbackLabelCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowbackLabelCreateDto() *ShowbackLabelCreateDto {
	this := ShowbackLabelCreateDto{}
	return &this
}

// NewShowbackLabelCreateDtoWithDefaults instantiates a new ShowbackLabelCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowbackLabelCreateDtoWithDefaults() *ShowbackLabelCreateDto {
	this := ShowbackLabelCreateDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShowbackLabelCreateDto) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShowbackLabelCreateDto) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *ShowbackLabelCreateDto) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *ShowbackLabelCreateDto) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil
func (o *ShowbackLabelCreateDto) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *ShowbackLabelCreateDto) UnsetLabel() {
	o.Label.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShowbackLabelCreateDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShowbackLabelCreateDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ShowbackLabelCreateDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ShowbackLabelCreateDto) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ShowbackLabelCreateDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ShowbackLabelCreateDto) UnsetValue() {
	o.Value.Unset()
}

func (o ShowbackLabelCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowbackLabelCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableShowbackLabelCreateDto struct {
	value *ShowbackLabelCreateDto
	isSet bool
}

func (v NullableShowbackLabelCreateDto) Get() *ShowbackLabelCreateDto {
	return v.value
}

func (v *NullableShowbackLabelCreateDto) Set(val *ShowbackLabelCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableShowbackLabelCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableShowbackLabelCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowbackLabelCreateDto(val *ShowbackLabelCreateDto) *NullableShowbackLabelCreateDto {
	return &NullableShowbackLabelCreateDto{value: val, isSet: true}
}

func (v NullableShowbackLabelCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowbackLabelCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
