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
	"bytes"
	"fmt"
)

// checks if the WebhookHeaderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookHeaderDto{}

// WebhookHeaderDto struct for WebhookHeaderDto
type WebhookHeaderDto struct {
	Id int32 `json:"id"`
	Key NullableString `json:"key"`
	Value NullableString `json:"value"`
}

type _WebhookHeaderDto WebhookHeaderDto

// NewWebhookHeaderDto instantiates a new WebhookHeaderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookHeaderDto(id int32, key NullableString, value NullableString) *WebhookHeaderDto {
	this := WebhookHeaderDto{}
	this.Id = id
	this.Key = key
	this.Value = value
	return &this
}

// NewWebhookHeaderDtoWithDefaults instantiates a new WebhookHeaderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookHeaderDtoWithDefaults() *WebhookHeaderDto {
	this := WebhookHeaderDto{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookHeaderDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookHeaderDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookHeaderDto) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookHeaderDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookHeaderDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *WebhookHeaderDto) SetKey(v string) {
	o.Key.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookHeaderDto) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookHeaderDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *WebhookHeaderDto) SetValue(v string) {
	o.Value.Set(&v)
}

func (o WebhookHeaderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookHeaderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key.Get()
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *WebhookHeaderDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookHeaderDto := _WebhookHeaderDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookHeaderDto)

	if err != nil {
		return err
	}

	*o = WebhookHeaderDto(varWebhookHeaderDto)

	return err
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


