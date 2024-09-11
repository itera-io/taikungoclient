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

// checks if the StandAloneMetaDataDtoForVm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneMetaDataDtoForVm{}

// StandAloneMetaDataDtoForVm struct for StandAloneMetaDataDtoForVm
type StandAloneMetaDataDtoForVm struct {
	Id *int32 `json:"id,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewStandAloneMetaDataDtoForVm instantiates a new StandAloneMetaDataDtoForVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneMetaDataDtoForVm() *StandAloneMetaDataDtoForVm {
	this := StandAloneMetaDataDtoForVm{}
	return &this
}

// NewStandAloneMetaDataDtoForVmWithDefaults instantiates a new StandAloneMetaDataDtoForVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneMetaDataDtoForVmWithDefaults() *StandAloneMetaDataDtoForVm {
	this := StandAloneMetaDataDtoForVm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneMetaDataDtoForVm) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneMetaDataDtoForVm) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneMetaDataDtoForVm) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneMetaDataDtoForVm) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneMetaDataDtoForVm) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneMetaDataDtoForVm) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *StandAloneMetaDataDtoForVm) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *StandAloneMetaDataDtoForVm) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *StandAloneMetaDataDtoForVm) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *StandAloneMetaDataDtoForVm) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneMetaDataDtoForVm) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneMetaDataDtoForVm) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *StandAloneMetaDataDtoForVm) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *StandAloneMetaDataDtoForVm) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *StandAloneMetaDataDtoForVm) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *StandAloneMetaDataDtoForVm) UnsetValue() {
	o.Value.Unset()
}

func (o StandAloneMetaDataDtoForVm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneMetaDataDtoForVm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneMetaDataDtoForVm struct {
	value *StandAloneMetaDataDtoForVm
	isSet bool
}

func (v NullableStandAloneMetaDataDtoForVm) Get() *StandAloneMetaDataDtoForVm {
	return v.value
}

func (v *NullableStandAloneMetaDataDtoForVm) Set(val *StandAloneMetaDataDtoForVm) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneMetaDataDtoForVm) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneMetaDataDtoForVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneMetaDataDtoForVm(val *StandAloneMetaDataDtoForVm) *NullableStandAloneMetaDataDtoForVm {
	return &NullableStandAloneMetaDataDtoForVm{value: val, isSet: true}
}

func (v NullableStandAloneMetaDataDtoForVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneMetaDataDtoForVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


