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

// checks if the UserDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDetails{}

// UserDetails struct for UserDetails
type UserDetails struct {
	Data                     *UserForListDto `json:"data,omitempty"`
	IsMaintenanceModeEnabled *bool           `json:"isMaintenanceModeEnabled,omitempty"`
	DemoOrganization         NullableInt32   `json:"demoOrganization,omitempty"`
	TrialDays                NullableInt32   `json:"trialDays,omitempty"`
}

// NewUserDetails instantiates a new UserDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDetails() *UserDetails {
	this := UserDetails{}
	return &this
}

// NewUserDetailsWithDefaults instantiates a new UserDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDetailsWithDefaults() *UserDetails {
	this := UserDetails{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserDetails) GetData() UserForListDto {
	if o == nil || IsNil(o.Data) {
		var ret UserForListDto
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetails) GetDataOk() (*UserForListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserDetails) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserForListDto and assigns it to the Data field.
func (o *UserDetails) SetData(v UserForListDto) {
	o.Data = &v
}

// GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field value if set, zero value otherwise.
func (o *UserDetails) GetIsMaintenanceModeEnabled() bool {
	if o == nil || IsNil(o.IsMaintenanceModeEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMaintenanceModeEnabled
}

// GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetails) GetIsMaintenanceModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaintenanceModeEnabled) {
		return nil, false
	}
	return o.IsMaintenanceModeEnabled, true
}

// HasIsMaintenanceModeEnabled returns a boolean if a field has been set.
func (o *UserDetails) HasIsMaintenanceModeEnabled() bool {
	if o != nil && !IsNil(o.IsMaintenanceModeEnabled) {
		return true
	}

	return false
}

// SetIsMaintenanceModeEnabled gets a reference to the given bool and assigns it to the IsMaintenanceModeEnabled field.
func (o *UserDetails) SetIsMaintenanceModeEnabled(v bool) {
	o.IsMaintenanceModeEnabled = &v
}

// GetDemoOrganization returns the DemoOrganization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetails) GetDemoOrganization() int32 {
	if o == nil || IsNil(o.DemoOrganization.Get()) {
		var ret int32
		return ret
	}
	return *o.DemoOrganization.Get()
}

// GetDemoOrganizationOk returns a tuple with the DemoOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetails) GetDemoOrganizationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DemoOrganization.Get(), o.DemoOrganization.IsSet()
}

// HasDemoOrganization returns a boolean if a field has been set.
func (o *UserDetails) HasDemoOrganization() bool {
	if o != nil && o.DemoOrganization.IsSet() {
		return true
	}

	return false
}

// SetDemoOrganization gets a reference to the given NullableInt32 and assigns it to the DemoOrganization field.
func (o *UserDetails) SetDemoOrganization(v int32) {
	o.DemoOrganization.Set(&v)
}

// SetDemoOrganizationNil sets the value for DemoOrganization to be an explicit nil
func (o *UserDetails) SetDemoOrganizationNil() {
	o.DemoOrganization.Set(nil)
}

// UnsetDemoOrganization ensures that no value is present for DemoOrganization, not even an explicit nil
func (o *UserDetails) UnsetDemoOrganization() {
	o.DemoOrganization.Unset()
}

// GetTrialDays returns the TrialDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDetails) GetTrialDays() int32 {
	if o == nil || IsNil(o.TrialDays.Get()) {
		var ret int32
		return ret
	}
	return *o.TrialDays.Get()
}

// GetTrialDaysOk returns a tuple with the TrialDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDetails) GetTrialDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialDays.Get(), o.TrialDays.IsSet()
}

// HasTrialDays returns a boolean if a field has been set.
func (o *UserDetails) HasTrialDays() bool {
	if o != nil && o.TrialDays.IsSet() {
		return true
	}

	return false
}

// SetTrialDays gets a reference to the given NullableInt32 and assigns it to the TrialDays field.
func (o *UserDetails) SetTrialDays(v int32) {
	o.TrialDays.Set(&v)
}

// SetTrialDaysNil sets the value for TrialDays to be an explicit nil
func (o *UserDetails) SetTrialDaysNil() {
	o.TrialDays.Set(nil)
}

// UnsetTrialDays ensures that no value is present for TrialDays, not even an explicit nil
func (o *UserDetails) UnsetTrialDays() {
	o.TrialDays.Unset()
}

func (o UserDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.IsMaintenanceModeEnabled) {
		toSerialize["isMaintenanceModeEnabled"] = o.IsMaintenanceModeEnabled
	}
	if o.DemoOrganization.IsSet() {
		toSerialize["demoOrganization"] = o.DemoOrganization.Get()
	}
	if o.TrialDays.IsSet() {
		toSerialize["trialDays"] = o.TrialDays.Get()
	}
	return toSerialize, nil
}

type NullableUserDetails struct {
	value *UserDetails
	isSet bool
}

func (v NullableUserDetails) Get() *UserDetails {
	return v.value
}

func (v *NullableUserDetails) Set(val *UserDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDetails(val *UserDetails) *NullableUserDetails {
	return &NullableUserDetails{value: val, isSet: true}
}

func (v NullableUserDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
