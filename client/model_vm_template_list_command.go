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

// checks if the VmTemplateListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmTemplateListCommand{}

// VmTemplateListCommand struct for VmTemplateListCommand
type VmTemplateListCommand struct {
	CloudId NullableInt32 `json:"cloudId,omitempty"`
	Url *string `json:"url,omitempty"`
	TokenId *string `json:"tokenId,omitempty"`
	TokenSecret *string `json:"tokenSecret,omitempty"`
}

// NewVmTemplateListCommand instantiates a new VmTemplateListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmTemplateListCommand() *VmTemplateListCommand {
	this := VmTemplateListCommand{}
	return &this
}

// NewVmTemplateListCommandWithDefaults instantiates a new VmTemplateListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmTemplateListCommandWithDefaults() *VmTemplateListCommand {
	this := VmTemplateListCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmTemplateListCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudId.Get()
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmTemplateListCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudId.Get(), o.CloudId.IsSet()
}

// HasCloudId returns a boolean if a field has been set.
func (o *VmTemplateListCommand) HasCloudId() bool {
	if o != nil && o.CloudId.IsSet() {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given NullableInt32 and assigns it to the CloudId field.
func (o *VmTemplateListCommand) SetCloudId(v int32) {
	o.CloudId.Set(&v)
}
// SetCloudIdNil sets the value for CloudId to be an explicit nil
func (o *VmTemplateListCommand) SetCloudIdNil() {
	o.CloudId.Set(nil)
}

// UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
func (o *VmTemplateListCommand) UnsetCloudId() {
	o.CloudId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VmTemplateListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmTemplateListCommand) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VmTemplateListCommand) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VmTemplateListCommand) SetUrl(v string) {
	o.Url = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *VmTemplateListCommand) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmTemplateListCommand) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *VmTemplateListCommand) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *VmTemplateListCommand) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenSecret returns the TokenSecret field value if set, zero value otherwise.
func (o *VmTemplateListCommand) GetTokenSecret() string {
	if o == nil || IsNil(o.TokenSecret) {
		var ret string
		return ret
	}
	return *o.TokenSecret
}

// GetTokenSecretOk returns a tuple with the TokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmTemplateListCommand) GetTokenSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TokenSecret) {
		return nil, false
	}
	return o.TokenSecret, true
}

// HasTokenSecret returns a boolean if a field has been set.
func (o *VmTemplateListCommand) HasTokenSecret() bool {
	if o != nil && !IsNil(o.TokenSecret) {
		return true
	}

	return false
}

// SetTokenSecret gets a reference to the given string and assigns it to the TokenSecret field.
func (o *VmTemplateListCommand) SetTokenSecret(v string) {
	o.TokenSecret = &v
}

func (o VmTemplateListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmTemplateListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudId.IsSet() {
		toSerialize["cloudId"] = o.CloudId.Get()
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.TokenSecret) {
		toSerialize["tokenSecret"] = o.TokenSecret
	}
	return toSerialize, nil
}

type NullableVmTemplateListCommand struct {
	value *VmTemplateListCommand
	isSet bool
}

func (v NullableVmTemplateListCommand) Get() *VmTemplateListCommand {
	return v.value
}

func (v *NullableVmTemplateListCommand) Set(val *VmTemplateListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVmTemplateListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVmTemplateListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmTemplateListCommand(val *VmTemplateListCommand) *NullableVmTemplateListCommand {
	return &NullableVmTemplateListCommand{value: val, isSet: true}
}

func (v NullableVmTemplateListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmTemplateListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


