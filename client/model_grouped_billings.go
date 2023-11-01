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

// checks if the GroupedBillings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedBillings{}

// GroupedBillings struct for GroupedBillings
type GroupedBillings struct {
	StartDate NullableString `json:"startDate,omitempty"`
	Tcu       *int64         `json:"tcu,omitempty"`
}

// NewGroupedBillings instantiates a new GroupedBillings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedBillings() *GroupedBillings {
	this := GroupedBillings{}
	return &this
}

// NewGroupedBillingsWithDefaults instantiates a new GroupedBillings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedBillingsWithDefaults() *GroupedBillings {
	this := GroupedBillings{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedBillings) GetStartDate() string {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedBillings) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *GroupedBillings) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *GroupedBillings) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *GroupedBillings) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *GroupedBillings) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetTcu returns the Tcu field value if set, zero value otherwise.
func (o *GroupedBillings) GetTcu() int64 {
	if o == nil || IsNil(o.Tcu) {
		var ret int64
		return ret
	}
	return *o.Tcu
}

// GetTcuOk returns a tuple with the Tcu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedBillings) GetTcuOk() (*int64, bool) {
	if o == nil || IsNil(o.Tcu) {
		return nil, false
	}
	return o.Tcu, true
}

// HasTcu returns a boolean if a field has been set.
func (o *GroupedBillings) HasTcu() bool {
	if o != nil && !IsNil(o.Tcu) {
		return true
	}

	return false
}

// SetTcu gets a reference to the given int64 and assigns it to the Tcu field.
func (o *GroupedBillings) SetTcu(v int64) {
	o.Tcu = &v
}

func (o GroupedBillings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedBillings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if !IsNil(o.Tcu) {
		toSerialize["tcu"] = o.Tcu
	}
	return toSerialize, nil
}

type NullableGroupedBillings struct {
	value *GroupedBillings
	isSet bool
}

func (v NullableGroupedBillings) Get() *GroupedBillings {
	return v.value
}

func (v *NullableGroupedBillings) Set(val *GroupedBillings) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedBillings) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedBillings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedBillings(val *GroupedBillings) *NullableGroupedBillings {
	return &NullableGroupedBillings{value: val, isSet: true}
}

func (v NullableGroupedBillings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedBillings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}