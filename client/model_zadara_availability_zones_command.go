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

// checks if the ZadaraAvailabilityZonesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZadaraAvailabilityZonesCommand{}

// ZadaraAvailabilityZonesCommand struct for ZadaraAvailabilityZonesCommand
type ZadaraAvailabilityZonesCommand struct {
	Url *string `json:"url,omitempty"`
	ZadaraAccessKeyId *string `json:"zadaraAccessKeyId,omitempty"`
	ZadaraSecretAccessKey *string `json:"zadaraSecretAccessKey,omitempty"`
	CloudId NullableInt32 `json:"cloudId,omitempty"`
}

// NewZadaraAvailabilityZonesCommand instantiates a new ZadaraAvailabilityZonesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZadaraAvailabilityZonesCommand() *ZadaraAvailabilityZonesCommand {
	this := ZadaraAvailabilityZonesCommand{}
	return &this
}

// NewZadaraAvailabilityZonesCommandWithDefaults instantiates a new ZadaraAvailabilityZonesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZadaraAvailabilityZonesCommandWithDefaults() *ZadaraAvailabilityZonesCommand {
	this := ZadaraAvailabilityZonesCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ZadaraAvailabilityZonesCommand) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZadaraAvailabilityZonesCommand) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ZadaraAvailabilityZonesCommand) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ZadaraAvailabilityZonesCommand) SetUrl(v string) {
	o.Url = &v
}

// GetZadaraAccessKeyId returns the ZadaraAccessKeyId field value if set, zero value otherwise.
func (o *ZadaraAvailabilityZonesCommand) GetZadaraAccessKeyId() string {
	if o == nil || IsNil(o.ZadaraAccessKeyId) {
		var ret string
		return ret
	}
	return *o.ZadaraAccessKeyId
}

// GetZadaraAccessKeyIdOk returns a tuple with the ZadaraAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZadaraAvailabilityZonesCommand) GetZadaraAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZadaraAccessKeyId) {
		return nil, false
	}
	return o.ZadaraAccessKeyId, true
}

// HasZadaraAccessKeyId returns a boolean if a field has been set.
func (o *ZadaraAvailabilityZonesCommand) HasZadaraAccessKeyId() bool {
	if o != nil && !IsNil(o.ZadaraAccessKeyId) {
		return true
	}

	return false
}

// SetZadaraAccessKeyId gets a reference to the given string and assigns it to the ZadaraAccessKeyId field.
func (o *ZadaraAvailabilityZonesCommand) SetZadaraAccessKeyId(v string) {
	o.ZadaraAccessKeyId = &v
}

// GetZadaraSecretAccessKey returns the ZadaraSecretAccessKey field value if set, zero value otherwise.
func (o *ZadaraAvailabilityZonesCommand) GetZadaraSecretAccessKey() string {
	if o == nil || IsNil(o.ZadaraSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.ZadaraSecretAccessKey
}

// GetZadaraSecretAccessKeyOk returns a tuple with the ZadaraSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZadaraAvailabilityZonesCommand) GetZadaraSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ZadaraSecretAccessKey) {
		return nil, false
	}
	return o.ZadaraSecretAccessKey, true
}

// HasZadaraSecretAccessKey returns a boolean if a field has been set.
func (o *ZadaraAvailabilityZonesCommand) HasZadaraSecretAccessKey() bool {
	if o != nil && !IsNil(o.ZadaraSecretAccessKey) {
		return true
	}

	return false
}

// SetZadaraSecretAccessKey gets a reference to the given string and assigns it to the ZadaraSecretAccessKey field.
func (o *ZadaraAvailabilityZonesCommand) SetZadaraSecretAccessKey(v string) {
	o.ZadaraSecretAccessKey = &v
}

// GetCloudId returns the CloudId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZadaraAvailabilityZonesCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudId.Get()
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZadaraAvailabilityZonesCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudId.Get(), o.CloudId.IsSet()
}

// HasCloudId returns a boolean if a field has been set.
func (o *ZadaraAvailabilityZonesCommand) HasCloudId() bool {
	if o != nil && o.CloudId.IsSet() {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given NullableInt32 and assigns it to the CloudId field.
func (o *ZadaraAvailabilityZonesCommand) SetCloudId(v int32) {
	o.CloudId.Set(&v)
}
// SetCloudIdNil sets the value for CloudId to be an explicit nil
func (o *ZadaraAvailabilityZonesCommand) SetCloudIdNil() {
	o.CloudId.Set(nil)
}

// UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
func (o *ZadaraAvailabilityZonesCommand) UnsetCloudId() {
	o.CloudId.Unset()
}

func (o ZadaraAvailabilityZonesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZadaraAvailabilityZonesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ZadaraAccessKeyId) {
		toSerialize["zadaraAccessKeyId"] = o.ZadaraAccessKeyId
	}
	if !IsNil(o.ZadaraSecretAccessKey) {
		toSerialize["zadaraSecretAccessKey"] = o.ZadaraSecretAccessKey
	}
	if o.CloudId.IsSet() {
		toSerialize["cloudId"] = o.CloudId.Get()
	}
	return toSerialize, nil
}

type NullableZadaraAvailabilityZonesCommand struct {
	value *ZadaraAvailabilityZonesCommand
	isSet bool
}

func (v NullableZadaraAvailabilityZonesCommand) Get() *ZadaraAvailabilityZonesCommand {
	return v.value
}

func (v *NullableZadaraAvailabilityZonesCommand) Set(val *ZadaraAvailabilityZonesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableZadaraAvailabilityZonesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableZadaraAvailabilityZonesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZadaraAvailabilityZonesCommand(val *ZadaraAvailabilityZonesCommand) *NullableZadaraAvailabilityZonesCommand {
	return &NullableZadaraAvailabilityZonesCommand{value: val, isSet: true}
}

func (v NullableZadaraAvailabilityZonesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZadaraAvailabilityZonesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


