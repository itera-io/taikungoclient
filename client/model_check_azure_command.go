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

// checks if the CheckAzureCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAzureCommand{}

// CheckAzureCommand struct for CheckAzureCommand
type CheckAzureCommand struct {
	AzureClientId NullableString `json:"azureClientId,omitempty"`
	AzureClientSecret NullableString `json:"azureClientSecret,omitempty"`
	AzureTenantId NullableString `json:"azureTenantId,omitempty"`
}

// NewCheckAzureCommand instantiates a new CheckAzureCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAzureCommand() *CheckAzureCommand {
	this := CheckAzureCommand{}
	return &this
}

// NewCheckAzureCommandWithDefaults instantiates a new CheckAzureCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAzureCommandWithDefaults() *CheckAzureCommand {
	this := CheckAzureCommand{}
	return &this
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckAzureCommand) GetAzureClientId() string {
	if o == nil || IsNil(o.AzureClientId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientId.Get()
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckAzureCommand) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientId.Get(), o.AzureClientId.IsSet()
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *CheckAzureCommand) HasAzureClientId() bool {
	if o != nil && o.AzureClientId.IsSet() {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given NullableString and assigns it to the AzureClientId field.
func (o *CheckAzureCommand) SetAzureClientId(v string) {
	o.AzureClientId.Set(&v)
}
// SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil
func (o *CheckAzureCommand) SetAzureClientIdNil() {
	o.AzureClientId.Set(nil)
}

// UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
func (o *CheckAzureCommand) UnsetAzureClientId() {
	o.AzureClientId.Unset()
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckAzureCommand) GetAzureClientSecret() string {
	if o == nil || IsNil(o.AzureClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientSecret.Get()
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckAzureCommand) GetAzureClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientSecret.Get(), o.AzureClientSecret.IsSet()
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *CheckAzureCommand) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret.IsSet() {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given NullableString and assigns it to the AzureClientSecret field.
func (o *CheckAzureCommand) SetAzureClientSecret(v string) {
	o.AzureClientSecret.Set(&v)
}
// SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil
func (o *CheckAzureCommand) SetAzureClientSecretNil() {
	o.AzureClientSecret.Set(nil)
}

// UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
func (o *CheckAzureCommand) UnsetAzureClientSecret() {
	o.AzureClientSecret.Unset()
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckAzureCommand) GetAzureTenantId() string {
	if o == nil || IsNil(o.AzureTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureTenantId.Get()
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckAzureCommand) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureTenantId.Get(), o.AzureTenantId.IsSet()
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *CheckAzureCommand) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId.IsSet() {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given NullableString and assigns it to the AzureTenantId field.
func (o *CheckAzureCommand) SetAzureTenantId(v string) {
	o.AzureTenantId.Set(&v)
}
// SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil
func (o *CheckAzureCommand) SetAzureTenantIdNil() {
	o.AzureTenantId.Set(nil)
}

// UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
func (o *CheckAzureCommand) UnsetAzureTenantId() {
	o.AzureTenantId.Unset()
}

func (o CheckAzureCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAzureCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureClientId.IsSet() {
		toSerialize["azureClientId"] = o.AzureClientId.Get()
	}
	if o.AzureClientSecret.IsSet() {
		toSerialize["azureClientSecret"] = o.AzureClientSecret.Get()
	}
	if o.AzureTenantId.IsSet() {
		toSerialize["azureTenantId"] = o.AzureTenantId.Get()
	}
	return toSerialize, nil
}

type NullableCheckAzureCommand struct {
	value *CheckAzureCommand
	isSet bool
}

func (v NullableCheckAzureCommand) Get() *CheckAzureCommand {
	return v.value
}

func (v *NullableCheckAzureCommand) Set(val *CheckAzureCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAzureCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAzureCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAzureCommand(val *CheckAzureCommand) *NullableCheckAzureCommand {
	return &NullableCheckAzureCommand{value: val, isSet: true}
}

func (v NullableCheckAzureCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAzureCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


