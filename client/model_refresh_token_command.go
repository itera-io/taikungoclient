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

// checks if the RefreshTokenCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshTokenCommand{}

// RefreshTokenCommand struct for RefreshTokenCommand
type RefreshTokenCommand struct {
	Token *string `json:"token,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
}

// NewRefreshTokenCommand instantiates a new RefreshTokenCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenCommand() *RefreshTokenCommand {
	this := RefreshTokenCommand{}
	return &this
}

// NewRefreshTokenCommandWithDefaults instantiates a new RefreshTokenCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenCommandWithDefaults() *RefreshTokenCommand {
	this := RefreshTokenCommand{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RefreshTokenCommand) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenCommand) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RefreshTokenCommand) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RefreshTokenCommand) SetToken(v string) {
	o.Token = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *RefreshTokenCommand) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenCommand) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *RefreshTokenCommand) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *RefreshTokenCommand) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o RefreshTokenCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshTokenCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableRefreshTokenCommand struct {
	value *RefreshTokenCommand
	isSet bool
}

func (v NullableRefreshTokenCommand) Get() *RefreshTokenCommand {
	return v.value
}

func (v *NullableRefreshTokenCommand) Set(val *RefreshTokenCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenCommand(val *RefreshTokenCommand) *NullableRefreshTokenCommand {
	return &NullableRefreshTokenCommand{value: val, isSet: true}
}

func (v NullableRefreshTokenCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


