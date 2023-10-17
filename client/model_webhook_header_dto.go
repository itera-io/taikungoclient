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

// checks if the WebhookHeaderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookHeaderDto{}

// WebhookHeaderDto struct for WebhookHeaderDto
type WebhookHeaderDto struct {
	Id    *int32         `json:"id,omitempty"`
	Key   NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewWebhookHeaderDto instantiates a new WebhookHeaderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookHeaderDto() *WebhookHeaderDto {
	this := WebhookHeaderDto{}
	return &this
}

// NewWebhookHeaderDtoWithDefaults instantiates a new WebhookHeaderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookHeaderDtoWithDefaults() *WebhookHeaderDto {
	this := WebhookHeaderDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookHeaderDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookHeaderDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookHeaderDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WebhookHeaderDto) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookHeaderDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookHeaderDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *WebhookHeaderDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *WebhookHeaderDto) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *WebhookHeaderDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *WebhookHeaderDto) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookHeaderDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookHeaderDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *WebhookHeaderDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *WebhookHeaderDto) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *WebhookHeaderDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *WebhookHeaderDto) UnsetValue() {
	o.Value.Unset()
}

func (o WebhookHeaderDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookHeaderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableWebhookHeaderDto struct {
	value *WebhookHeaderDto
	isSet bool
}

func (v NullableWebhookHeaderDto) Get() *WebhookHeaderDto {
	return v.value
}

func (v *NullableWebhookHeaderDto) Set(val *WebhookHeaderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookHeaderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookHeaderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookHeaderDto(val *WebhookHeaderDto) *NullableWebhookHeaderDto {
	return &NullableWebhookHeaderDto{value: val, isSet: true}
}

func (v NullableWebhookHeaderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookHeaderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
