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

// checks if the NetworkListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkListCommand{}

// NetworkListCommand struct for NetworkListCommand
type NetworkListCommand struct {
	Url NullableString `json:"url,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	DatacenterId NullableString `json:"datacenterId,omitempty"`
}

// NewNetworkListCommand instantiates a new NetworkListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkListCommand() *NetworkListCommand {
	this := NetworkListCommand{}
	return &this
}

// NewNetworkListCommandWithDefaults instantiates a new NetworkListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkListCommandWithDefaults() *NetworkListCommand {
	this := NetworkListCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkListCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworkListCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *NetworkListCommand) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *NetworkListCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *NetworkListCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkListCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkListCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *NetworkListCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *NetworkListCommand) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *NetworkListCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *NetworkListCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkListCommand) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkListCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworkListCommand) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *NetworkListCommand) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *NetworkListCommand) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *NetworkListCommand) UnsetPassword() {
	o.Password.Unset()
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkListCommand) GetDatacenterId() string {
	if o == nil || IsNil(o.DatacenterId.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterId.Get()
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkListCommand) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterId.Get(), o.DatacenterId.IsSet()
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *NetworkListCommand) HasDatacenterId() bool {
	if o != nil && o.DatacenterId.IsSet() {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given NullableString and assigns it to the DatacenterId field.
func (o *NetworkListCommand) SetDatacenterId(v string) {
	o.DatacenterId.Set(&v)
}
// SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil
func (o *NetworkListCommand) SetDatacenterIdNil() {
	o.DatacenterId.Set(nil)
}

// UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
func (o *NetworkListCommand) UnsetDatacenterId() {
	o.DatacenterId.Unset()
}

func (o NetworkListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.DatacenterId.IsSet() {
		toSerialize["datacenterId"] = o.DatacenterId.Get()
	}
	return toSerialize, nil
}

type NullableNetworkListCommand struct {
	value *NetworkListCommand
	isSet bool
}

func (v NullableNetworkListCommand) Get() *NetworkListCommand {
	return v.value
}

func (v *NullableNetworkListCommand) Set(val *NetworkListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkListCommand(val *NetworkListCommand) *NullableNetworkListCommand {
	return &NullableNetworkListCommand{value: val, isSet: true}
}

func (v NullableNetworkListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


