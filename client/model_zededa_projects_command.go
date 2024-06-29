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

// checks if the ZededaProjectsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZededaProjectsCommand{}

// ZededaProjectsCommand struct for ZededaProjectsCommand
type ZededaProjectsCommand struct {
	ApiUrl NullableString `json:"apiUrl,omitempty"`
	ApiToken NullableString `json:"apiToken,omitempty"`
}

// NewZededaProjectsCommand instantiates a new ZededaProjectsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZededaProjectsCommand() *ZededaProjectsCommand {
	this := ZededaProjectsCommand{}
	return &this
}

// NewZededaProjectsCommandWithDefaults instantiates a new ZededaProjectsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZededaProjectsCommandWithDefaults() *ZededaProjectsCommand {
	this := ZededaProjectsCommand{}
	return &this
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZededaProjectsCommand) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ApiUrl.Get()
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZededaProjectsCommand) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl.Get(), o.ApiUrl.IsSet()
}

// HasApiUrl returns a boolean if a field has been set.
func (o *ZededaProjectsCommand) HasApiUrl() bool {
	if o != nil && o.ApiUrl.IsSet() {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given NullableString and assigns it to the ApiUrl field.
func (o *ZededaProjectsCommand) SetApiUrl(v string) {
	o.ApiUrl.Set(&v)
}
// SetApiUrlNil sets the value for ApiUrl to be an explicit nil
func (o *ZededaProjectsCommand) SetApiUrlNil() {
	o.ApiUrl.Set(nil)
}

// UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
func (o *ZededaProjectsCommand) UnsetApiUrl() {
	o.ApiUrl.Unset()
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZededaProjectsCommand) GetApiToken() string {
	if o == nil || IsNil(o.ApiToken.Get()) {
		var ret string
		return ret
	}
	return *o.ApiToken.Get()
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZededaProjectsCommand) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiToken.Get(), o.ApiToken.IsSet()
}

// HasApiToken returns a boolean if a field has been set.
func (o *ZededaProjectsCommand) HasApiToken() bool {
	if o != nil && o.ApiToken.IsSet() {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given NullableString and assigns it to the ApiToken field.
func (o *ZededaProjectsCommand) SetApiToken(v string) {
	o.ApiToken.Set(&v)
}
// SetApiTokenNil sets the value for ApiToken to be an explicit nil
func (o *ZededaProjectsCommand) SetApiTokenNil() {
	o.ApiToken.Set(nil)
}

// UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
func (o *ZededaProjectsCommand) UnsetApiToken() {
	o.ApiToken.Unset()
}

func (o ZededaProjectsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZededaProjectsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiUrl.IsSet() {
		toSerialize["apiUrl"] = o.ApiUrl.Get()
	}
	if o.ApiToken.IsSet() {
		toSerialize["apiToken"] = o.ApiToken.Get()
	}
	return toSerialize, nil
}

type NullableZededaProjectsCommand struct {
	value *ZededaProjectsCommand
	isSet bool
}

func (v NullableZededaProjectsCommand) Get() *ZededaProjectsCommand {
	return v.value
}

func (v *NullableZededaProjectsCommand) Set(val *ZededaProjectsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableZededaProjectsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableZededaProjectsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZededaProjectsCommand(val *ZededaProjectsCommand) *NullableZededaProjectsCommand {
	return &NullableZededaProjectsCommand{value: val, isSet: true}
}

func (v NullableZededaProjectsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZededaProjectsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


