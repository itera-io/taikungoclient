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

// checks if the PartnerImageSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerImageSettingsDto{}

// PartnerImageSettingsDto struct for PartnerImageSettingsDto
type PartnerImageSettingsDto struct {
	Expanded NullableString `json:"expanded"`
	Collapsed NullableString `json:"collapsed"`
}

type _PartnerImageSettingsDto PartnerImageSettingsDto

// NewPartnerImageSettingsDto instantiates a new PartnerImageSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerImageSettingsDto(expanded NullableString, collapsed NullableString) *PartnerImageSettingsDto {
	this := PartnerImageSettingsDto{}
	this.Expanded = expanded
	this.Collapsed = collapsed
	return &this
}

// NewPartnerImageSettingsDtoWithDefaults instantiates a new PartnerImageSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerImageSettingsDtoWithDefaults() *PartnerImageSettingsDto {
	this := PartnerImageSettingsDto{}
	return &this
}

// GetExpanded returns the Expanded field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartnerImageSettingsDto) GetExpanded() string {
	if o == nil || o.Expanded.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expanded.Get()
}

// GetExpandedOk returns a tuple with the Expanded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerImageSettingsDto) GetExpandedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expanded.Get(), o.Expanded.IsSet()
}

// SetExpanded sets field value
func (o *PartnerImageSettingsDto) SetExpanded(v string) {
	o.Expanded.Set(&v)
}

// GetCollapsed returns the Collapsed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartnerImageSettingsDto) GetCollapsed() string {
	if o == nil || o.Collapsed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Collapsed.Get()
}

// GetCollapsedOk returns a tuple with the Collapsed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerImageSettingsDto) GetCollapsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collapsed.Get(), o.Collapsed.IsSet()
}

// SetCollapsed sets field value
func (o *PartnerImageSettingsDto) SetCollapsed(v string) {
	o.Collapsed.Set(&v)
}

func (o PartnerImageSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerImageSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expanded"] = o.Expanded.Get()
	toSerialize["collapsed"] = o.Collapsed.Get()
	return toSerialize, nil
}

func (o *PartnerImageSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expanded",
		"collapsed",
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

	varPartnerImageSettingsDto := _PartnerImageSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartnerImageSettingsDto)

	if err != nil {
		return err
	}

	*o = PartnerImageSettingsDto(varPartnerImageSettingsDto)

	return err
}

type NullablePartnerImageSettingsDto struct {
	value *PartnerImageSettingsDto
	isSet bool
}

func (v NullablePartnerImageSettingsDto) Get() *PartnerImageSettingsDto {
	return v.value
}

func (v *NullablePartnerImageSettingsDto) Set(val *PartnerImageSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerImageSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerImageSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerImageSettingsDto(val *PartnerImageSettingsDto) *NullablePartnerImageSettingsDto {
	return &NullablePartnerImageSettingsDto{value: val, isSet: true}
}

func (v NullablePartnerImageSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerImageSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


