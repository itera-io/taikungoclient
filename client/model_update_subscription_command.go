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

// checks if the UpdateSubscriptionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSubscriptionCommand{}

// UpdateSubscriptionCommand struct for UpdateSubscriptionCommand
type UpdateSubscriptionCommand struct {
	Id                   *int32         `json:"id,omitempty"`
	Name                 NullableString `json:"name,omitempty"`
	ProjectLimit         *int32         `json:"projectLimit,omitempty"`
	ServerLimit          *int32         `json:"serverLimit,omitempty"`
	UserLimit            *int32         `json:"userLimit,omitempty"`
	CloudCredentialLimit *int32         `json:"cloudCredentialLimit,omitempty"`
	MonthlyPrice         *float64       `json:"monthlyPrice,omitempty"`
	YearlyPrice          *float64       `json:"yearlyPrice,omitempty"`
	TcuPrice             *float64       `json:"tcuPrice,omitempty"`
	TrialDays            *int32         `json:"trialDays,omitempty"`
}

// NewUpdateSubscriptionCommand instantiates a new UpdateSubscriptionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionCommand() *UpdateSubscriptionCommand {
	this := UpdateSubscriptionCommand{}
	return &this
}

// NewUpdateSubscriptionCommandWithDefaults instantiates a new UpdateSubscriptionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionCommandWithDefaults() *UpdateSubscriptionCommand {
	this := UpdateSubscriptionCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateSubscriptionCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSubscriptionCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSubscriptionCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateSubscriptionCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateSubscriptionCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateSubscriptionCommand) UnsetName() {
	o.Name.Unset()
}

// GetProjectLimit returns the ProjectLimit field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetProjectLimit() int32 {
	if o == nil || IsNil(o.ProjectLimit) {
		var ret int32
		return ret
	}
	return *o.ProjectLimit
}

// GetProjectLimitOk returns a tuple with the ProjectLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetProjectLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectLimit) {
		return nil, false
	}
	return o.ProjectLimit, true
}

// HasProjectLimit returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasProjectLimit() bool {
	if o != nil && !IsNil(o.ProjectLimit) {
		return true
	}

	return false
}

// SetProjectLimit gets a reference to the given int32 and assigns it to the ProjectLimit field.
func (o *UpdateSubscriptionCommand) SetProjectLimit(v int32) {
	o.ProjectLimit = &v
}

// GetServerLimit returns the ServerLimit field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetServerLimit() int32 {
	if o == nil || IsNil(o.ServerLimit) {
		var ret int32
		return ret
	}
	return *o.ServerLimit
}

// GetServerLimitOk returns a tuple with the ServerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetServerLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerLimit) {
		return nil, false
	}
	return o.ServerLimit, true
}

// HasServerLimit returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasServerLimit() bool {
	if o != nil && !IsNil(o.ServerLimit) {
		return true
	}

	return false
}

// SetServerLimit gets a reference to the given int32 and assigns it to the ServerLimit field.
func (o *UpdateSubscriptionCommand) SetServerLimit(v int32) {
	o.ServerLimit = &v
}

// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *UpdateSubscriptionCommand) SetUserLimit(v int32) {
	o.UserLimit = &v
}

// GetCloudCredentialLimit returns the CloudCredentialLimit field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetCloudCredentialLimit() int32 {
	if o == nil || IsNil(o.CloudCredentialLimit) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialLimit
}

// GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetCloudCredentialLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudCredentialLimit) {
		return nil, false
	}
	return o.CloudCredentialLimit, true
}

// HasCloudCredentialLimit returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasCloudCredentialLimit() bool {
	if o != nil && !IsNil(o.CloudCredentialLimit) {
		return true
	}

	return false
}

// SetCloudCredentialLimit gets a reference to the given int32 and assigns it to the CloudCredentialLimit field.
func (o *UpdateSubscriptionCommand) SetCloudCredentialLimit(v int32) {
	o.CloudCredentialLimit = &v
}

// GetMonthlyPrice returns the MonthlyPrice field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetMonthlyPrice() float64 {
	if o == nil || IsNil(o.MonthlyPrice) {
		var ret float64
		return ret
	}
	return *o.MonthlyPrice
}

// GetMonthlyPriceOk returns a tuple with the MonthlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetMonthlyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.MonthlyPrice) {
		return nil, false
	}
	return o.MonthlyPrice, true
}

// HasMonthlyPrice returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasMonthlyPrice() bool {
	if o != nil && !IsNil(o.MonthlyPrice) {
		return true
	}

	return false
}

// SetMonthlyPrice gets a reference to the given float64 and assigns it to the MonthlyPrice field.
func (o *UpdateSubscriptionCommand) SetMonthlyPrice(v float64) {
	o.MonthlyPrice = &v
}

// GetYearlyPrice returns the YearlyPrice field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetYearlyPrice() float64 {
	if o == nil || IsNil(o.YearlyPrice) {
		var ret float64
		return ret
	}
	return *o.YearlyPrice
}

// GetYearlyPriceOk returns a tuple with the YearlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetYearlyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.YearlyPrice) {
		return nil, false
	}
	return o.YearlyPrice, true
}

// HasYearlyPrice returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasYearlyPrice() bool {
	if o != nil && !IsNil(o.YearlyPrice) {
		return true
	}

	return false
}

// SetYearlyPrice gets a reference to the given float64 and assigns it to the YearlyPrice field.
func (o *UpdateSubscriptionCommand) SetYearlyPrice(v float64) {
	o.YearlyPrice = &v
}

// GetTcuPrice returns the TcuPrice field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetTcuPrice() float64 {
	if o == nil || IsNil(o.TcuPrice) {
		var ret float64
		return ret
	}
	return *o.TcuPrice
}

// GetTcuPriceOk returns a tuple with the TcuPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetTcuPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.TcuPrice) {
		return nil, false
	}
	return o.TcuPrice, true
}

// HasTcuPrice returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasTcuPrice() bool {
	if o != nil && !IsNil(o.TcuPrice) {
		return true
	}

	return false
}

// SetTcuPrice gets a reference to the given float64 and assigns it to the TcuPrice field.
func (o *UpdateSubscriptionCommand) SetTcuPrice(v float64) {
	o.TcuPrice = &v
}

// GetTrialDays returns the TrialDays field value if set, zero value otherwise.
func (o *UpdateSubscriptionCommand) GetTrialDays() int32 {
	if o == nil || IsNil(o.TrialDays) {
		var ret int32
		return ret
	}
	return *o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionCommand) GetTrialDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialDays) {
		return nil, false
	}
	return o.TrialDays, true
}

// HasTrialDays returns a boolean if a field has been set.
func (o *UpdateSubscriptionCommand) HasTrialDays() bool {
	if o != nil && !IsNil(o.TrialDays) {
		return true
	}

	return false
}

// SetTrialDays gets a reference to the given int32 and assigns it to the TrialDays field.
func (o *UpdateSubscriptionCommand) SetTrialDays(v int32) {
	o.TrialDays = &v
}

func (o UpdateSubscriptionCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSubscriptionCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ProjectLimit) {
		toSerialize["projectLimit"] = o.ProjectLimit
	}
	if !IsNil(o.ServerLimit) {
		toSerialize["serverLimit"] = o.ServerLimit
	}
	if !IsNil(o.UserLimit) {
		toSerialize["userLimit"] = o.UserLimit
	}
	if !IsNil(o.CloudCredentialLimit) {
		toSerialize["cloudCredentialLimit"] = o.CloudCredentialLimit
	}
	if !IsNil(o.MonthlyPrice) {
		toSerialize["monthlyPrice"] = o.MonthlyPrice
	}
	if !IsNil(o.YearlyPrice) {
		toSerialize["yearlyPrice"] = o.YearlyPrice
	}
	if !IsNil(o.TcuPrice) {
		toSerialize["tcuPrice"] = o.TcuPrice
	}
	if !IsNil(o.TrialDays) {
		toSerialize["trialDays"] = o.TrialDays
	}
	return toSerialize, nil
}

type NullableUpdateSubscriptionCommand struct {
	value *UpdateSubscriptionCommand
	isSet bool
}

func (v NullableUpdateSubscriptionCommand) Get() *UpdateSubscriptionCommand {
	return v.value
}

func (v *NullableUpdateSubscriptionCommand) Set(val *UpdateSubscriptionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionCommand(val *UpdateSubscriptionCommand) *NullableUpdateSubscriptionCommand {
	return &NullableUpdateSubscriptionCommand{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
