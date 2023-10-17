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

// checks if the AdminAddBalanceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminAddBalanceCommand{}

// AdminAddBalanceCommand struct for AdminAddBalanceCommand
type AdminAddBalanceCommand struct {
	CustomerId NullableString `json:"customerId,omitempty"`
	Balance    *int64         `json:"balance,omitempty"`
}

// NewAdminAddBalanceCommand instantiates a new AdminAddBalanceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminAddBalanceCommand() *AdminAddBalanceCommand {
	this := AdminAddBalanceCommand{}
	return &this
}

// NewAdminAddBalanceCommandWithDefaults instantiates a new AdminAddBalanceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminAddBalanceCommandWithDefaults() *AdminAddBalanceCommand {
	this := AdminAddBalanceCommand{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminAddBalanceCommand) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminAddBalanceCommand) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AdminAddBalanceCommand) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *AdminAddBalanceCommand) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *AdminAddBalanceCommand) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *AdminAddBalanceCommand) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AdminAddBalanceCommand) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAddBalanceCommand) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AdminAddBalanceCommand) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *AdminAddBalanceCommand) SetBalance(v int64) {
	o.Balance = &v
}

func (o AdminAddBalanceCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminAddBalanceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableAdminAddBalanceCommand struct {
	value *AdminAddBalanceCommand
	isSet bool
}

func (v NullableAdminAddBalanceCommand) Get() *AdminAddBalanceCommand {
	return v.value
}

func (v *NullableAdminAddBalanceCommand) Set(val *AdminAddBalanceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminAddBalanceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminAddBalanceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminAddBalanceCommand(val *AdminAddBalanceCommand) *NullableAdminAddBalanceCommand {
	return &NullableAdminAddBalanceCommand{value: val, isSet: true}
}

func (v NullableAdminAddBalanceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminAddBalanceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
