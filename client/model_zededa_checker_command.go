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

// checks if the ZededaCheckerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZededaCheckerCommand{}

// ZededaCheckerCommand struct for ZededaCheckerCommand
type ZededaCheckerCommand struct {
	Url *string `json:"url,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewZededaCheckerCommand instantiates a new ZededaCheckerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZededaCheckerCommand() *ZededaCheckerCommand {
	this := ZededaCheckerCommand{}
	return &this
}

// NewZededaCheckerCommandWithDefaults instantiates a new ZededaCheckerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZededaCheckerCommandWithDefaults() *ZededaCheckerCommand {
	this := ZededaCheckerCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ZededaCheckerCommand) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZededaCheckerCommand) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ZededaCheckerCommand) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ZededaCheckerCommand) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ZededaCheckerCommand) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZededaCheckerCommand) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ZededaCheckerCommand) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ZededaCheckerCommand) SetToken(v string) {
	o.Token = &v
}

func (o ZededaCheckerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZededaCheckerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableZededaCheckerCommand struct {
	value *ZededaCheckerCommand
	isSet bool
}

func (v NullableZededaCheckerCommand) Get() *ZededaCheckerCommand {
	return v.value
}

func (v *NullableZededaCheckerCommand) Set(val *ZededaCheckerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableZededaCheckerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableZededaCheckerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZededaCheckerCommand(val *ZededaCheckerCommand) *NullableZededaCheckerCommand {
	return &NullableZededaCheckerCommand{value: val, isSet: true}
}

func (v NullableZededaCheckerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZededaCheckerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


