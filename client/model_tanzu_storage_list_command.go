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

// checks if the TanzuStorageListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuStorageListCommand{}

// TanzuStorageListCommand struct for TanzuStorageListCommand
type TanzuStorageListCommand struct {
	Username  NullableString `json:"username,omitempty"`
	Url       NullableString `json:"url,omitempty"`
	Password  NullableString `json:"password,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Port      NullableInt32  `json:"port,omitempty"`
}

// NewTanzuStorageListCommand instantiates a new TanzuStorageListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuStorageListCommand() *TanzuStorageListCommand {
	this := TanzuStorageListCommand{}
	return &this
}

// NewTanzuStorageListCommandWithDefaults instantiates a new TanzuStorageListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuStorageListCommandWithDefaults() *TanzuStorageListCommand {
	this := TanzuStorageListCommand{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuStorageListCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuStorageListCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TanzuStorageListCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TanzuStorageListCommand) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TanzuStorageListCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TanzuStorageListCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuStorageListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuStorageListCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *TanzuStorageListCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *TanzuStorageListCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *TanzuStorageListCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *TanzuStorageListCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuStorageListCommand) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuStorageListCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TanzuStorageListCommand) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TanzuStorageListCommand) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TanzuStorageListCommand) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TanzuStorageListCommand) UnsetPassword() {
	o.Password.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuStorageListCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuStorageListCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *TanzuStorageListCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *TanzuStorageListCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *TanzuStorageListCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *TanzuStorageListCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuStorageListCommand) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuStorageListCommand) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TanzuStorageListCommand) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *TanzuStorageListCommand) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *TanzuStorageListCommand) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TanzuStorageListCommand) UnsetPort() {
	o.Port.Unset()
}

func (o TanzuStorageListCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuStorageListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return toSerialize, nil
}

type NullableTanzuStorageListCommand struct {
	value *TanzuStorageListCommand
	isSet bool
}

func (v NullableTanzuStorageListCommand) Get() *TanzuStorageListCommand {
	return v.value
}

func (v *NullableTanzuStorageListCommand) Set(val *TanzuStorageListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuStorageListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuStorageListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuStorageListCommand(val *TanzuStorageListCommand) *NullableTanzuStorageListCommand {
	return &NullableTanzuStorageListCommand{value: val, isSet: true}
}

func (v NullableTanzuStorageListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuStorageListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
