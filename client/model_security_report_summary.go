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

// checks if the SecurityReportSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityReportSummary{}

// SecurityReportSummary struct for SecurityReportSummary
type SecurityReportSummary struct {
	Low NullableInt64 `json:"low,omitempty"`
	High NullableInt64 `json:"high,omitempty"`
	Medium NullableInt64 `json:"medium,omitempty"`
	Unknown NullableInt64 `json:"unknown,omitempty"`
	Critical NullableInt64 `json:"critical,omitempty"`
}

// NewSecurityReportSummary instantiates a new SecurityReportSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityReportSummary() *SecurityReportSummary {
	this := SecurityReportSummary{}
	return &this
}

// NewSecurityReportSummaryWithDefaults instantiates a new SecurityReportSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityReportSummaryWithDefaults() *SecurityReportSummary {
	this := SecurityReportSummary{}
	return &this
}

// GetLow returns the Low field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityReportSummary) GetLow() int64 {
	if o == nil || IsNil(o.Low.Get()) {
		var ret int64
		return ret
	}
	return *o.Low.Get()
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityReportSummary) GetLowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Low.Get(), o.Low.IsSet()
}

// HasLow returns a boolean if a field has been set.
func (o *SecurityReportSummary) HasLow() bool {
	if o != nil && o.Low.IsSet() {
		return true
	}

	return false
}

// SetLow gets a reference to the given NullableInt64 and assigns it to the Low field.
func (o *SecurityReportSummary) SetLow(v int64) {
	o.Low.Set(&v)
}
// SetLowNil sets the value for Low to be an explicit nil
func (o *SecurityReportSummary) SetLowNil() {
	o.Low.Set(nil)
}

// UnsetLow ensures that no value is present for Low, not even an explicit nil
func (o *SecurityReportSummary) UnsetLow() {
	o.Low.Unset()
}

// GetHigh returns the High field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityReportSummary) GetHigh() int64 {
	if o == nil || IsNil(o.High.Get()) {
		var ret int64
		return ret
	}
	return *o.High.Get()
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityReportSummary) GetHighOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.High.Get(), o.High.IsSet()
}

// HasHigh returns a boolean if a field has been set.
func (o *SecurityReportSummary) HasHigh() bool {
	if o != nil && o.High.IsSet() {
		return true
	}

	return false
}

// SetHigh gets a reference to the given NullableInt64 and assigns it to the High field.
func (o *SecurityReportSummary) SetHigh(v int64) {
	o.High.Set(&v)
}
// SetHighNil sets the value for High to be an explicit nil
func (o *SecurityReportSummary) SetHighNil() {
	o.High.Set(nil)
}

// UnsetHigh ensures that no value is present for High, not even an explicit nil
func (o *SecurityReportSummary) UnsetHigh() {
	o.High.Unset()
}

// GetMedium returns the Medium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityReportSummary) GetMedium() int64 {
	if o == nil || IsNil(o.Medium.Get()) {
		var ret int64
		return ret
	}
	return *o.Medium.Get()
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityReportSummary) GetMediumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medium.Get(), o.Medium.IsSet()
}

// HasMedium returns a boolean if a field has been set.
func (o *SecurityReportSummary) HasMedium() bool {
	if o != nil && o.Medium.IsSet() {
		return true
	}

	return false
}

// SetMedium gets a reference to the given NullableInt64 and assigns it to the Medium field.
func (o *SecurityReportSummary) SetMedium(v int64) {
	o.Medium.Set(&v)
}
// SetMediumNil sets the value for Medium to be an explicit nil
func (o *SecurityReportSummary) SetMediumNil() {
	o.Medium.Set(nil)
}

// UnsetMedium ensures that no value is present for Medium, not even an explicit nil
func (o *SecurityReportSummary) UnsetMedium() {
	o.Medium.Unset()
}

// GetUnknown returns the Unknown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityReportSummary) GetUnknown() int64 {
	if o == nil || IsNil(o.Unknown.Get()) {
		var ret int64
		return ret
	}
	return *o.Unknown.Get()
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityReportSummary) GetUnknownOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unknown.Get(), o.Unknown.IsSet()
}

// HasUnknown returns a boolean if a field has been set.
func (o *SecurityReportSummary) HasUnknown() bool {
	if o != nil && o.Unknown.IsSet() {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given NullableInt64 and assigns it to the Unknown field.
func (o *SecurityReportSummary) SetUnknown(v int64) {
	o.Unknown.Set(&v)
}
// SetUnknownNil sets the value for Unknown to be an explicit nil
func (o *SecurityReportSummary) SetUnknownNil() {
	o.Unknown.Set(nil)
}

// UnsetUnknown ensures that no value is present for Unknown, not even an explicit nil
func (o *SecurityReportSummary) UnsetUnknown() {
	o.Unknown.Unset()
}

// GetCritical returns the Critical field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityReportSummary) GetCritical() int64 {
	if o == nil || IsNil(o.Critical.Get()) {
		var ret int64
		return ret
	}
	return *o.Critical.Get()
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityReportSummary) GetCriticalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Critical.Get(), o.Critical.IsSet()
}

// HasCritical returns a boolean if a field has been set.
func (o *SecurityReportSummary) HasCritical() bool {
	if o != nil && o.Critical.IsSet() {
		return true
	}

	return false
}

// SetCritical gets a reference to the given NullableInt64 and assigns it to the Critical field.
func (o *SecurityReportSummary) SetCritical(v int64) {
	o.Critical.Set(&v)
}
// SetCriticalNil sets the value for Critical to be an explicit nil
func (o *SecurityReportSummary) SetCriticalNil() {
	o.Critical.Set(nil)
}

// UnsetCritical ensures that no value is present for Critical, not even an explicit nil
func (o *SecurityReportSummary) UnsetCritical() {
	o.Critical.Unset()
}

func (o SecurityReportSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityReportSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Low.IsSet() {
		toSerialize["low"] = o.Low.Get()
	}
	if o.High.IsSet() {
		toSerialize["high"] = o.High.Get()
	}
	if o.Medium.IsSet() {
		toSerialize["medium"] = o.Medium.Get()
	}
	if o.Unknown.IsSet() {
		toSerialize["unknown"] = o.Unknown.Get()
	}
	if o.Critical.IsSet() {
		toSerialize["critical"] = o.Critical.Get()
	}
	return toSerialize, nil
}

type NullableSecurityReportSummary struct {
	value *SecurityReportSummary
	isSet bool
}

func (v NullableSecurityReportSummary) Get() *SecurityReportSummary {
	return v.value
}

func (v *NullableSecurityReportSummary) Set(val *SecurityReportSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityReportSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityReportSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityReportSummary(val *SecurityReportSummary) *NullableSecurityReportSummary {
	return &NullableSecurityReportSummary{value: val, isSet: true}
}

func (v NullableSecurityReportSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityReportSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


