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

// checks if the HypervisorListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorListCommand{}

// HypervisorListCommand struct for HypervisorListCommand
type HypervisorListCommand struct {
	Url *string `json:"url,omitempty"`
	TokenId *string `json:"tokenId,omitempty"`
	TokenSecret *string `json:"tokenSecret,omitempty"`
	CloudId NullableInt32 `json:"cloudId,omitempty"`
}

// NewHypervisorListCommand instantiates a new HypervisorListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorListCommand() *HypervisorListCommand {
	this := HypervisorListCommand{}
	return &this
}

// NewHypervisorListCommandWithDefaults instantiates a new HypervisorListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorListCommandWithDefaults() *HypervisorListCommand {
	this := HypervisorListCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HypervisorListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorListCommand) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HypervisorListCommand) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HypervisorListCommand) SetUrl(v string) {
	o.Url = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *HypervisorListCommand) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorListCommand) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *HypervisorListCommand) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *HypervisorListCommand) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenSecret returns the TokenSecret field value if set, zero value otherwise.
func (o *HypervisorListCommand) GetTokenSecret() string {
	if o == nil || IsNil(o.TokenSecret) {
		var ret string
		return ret
	}
	return *o.TokenSecret
}

// GetTokenSecretOk returns a tuple with the TokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorListCommand) GetTokenSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TokenSecret) {
		return nil, false
	}
	return o.TokenSecret, true
}

// HasTokenSecret returns a boolean if a field has been set.
func (o *HypervisorListCommand) HasTokenSecret() bool {
	if o != nil && !IsNil(o.TokenSecret) {
		return true
	}

	return false
}

// SetTokenSecret gets a reference to the given string and assigns it to the TokenSecret field.
func (o *HypervisorListCommand) SetTokenSecret(v string) {
	o.TokenSecret = &v
}

// GetCloudId returns the CloudId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorListCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudId.Get()
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorListCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudId.Get(), o.CloudId.IsSet()
}

// HasCloudId returns a boolean if a field has been set.
func (o *HypervisorListCommand) HasCloudId() bool {
	if o != nil && o.CloudId.IsSet() {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given NullableInt32 and assigns it to the CloudId field.
func (o *HypervisorListCommand) SetCloudId(v int32) {
	o.CloudId.Set(&v)
}
// SetCloudIdNil sets the value for CloudId to be an explicit nil
func (o *HypervisorListCommand) SetCloudIdNil() {
	o.CloudId.Set(nil)
}

// UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
func (o *HypervisorListCommand) UnsetCloudId() {
	o.CloudId.Unset()
}

func (o HypervisorListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.TokenSecret) {
		toSerialize["tokenSecret"] = o.TokenSecret
	}
	if o.CloudId.IsSet() {
		toSerialize["cloudId"] = o.CloudId.Get()
	}
	return toSerialize, nil
}

type NullableHypervisorListCommand struct {
	value *HypervisorListCommand
	isSet bool
}

func (v NullableHypervisorListCommand) Get() *HypervisorListCommand {
	return v.value
}

func (v *NullableHypervisorListCommand) Set(val *HypervisorListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorListCommand(val *HypervisorListCommand) *NullableHypervisorListCommand {
	return &NullableHypervisorListCommand{value: val, isSet: true}
}

func (v NullableHypervisorListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


