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

// checks if the AlertingWebhookDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingWebhookDto{}

// AlertingWebhookDto struct for AlertingWebhookDto
type AlertingWebhookDto struct {
	Id *int32 `json:"id,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Headers []WebhookHeaderDto `json:"headers,omitempty"`
}

// NewAlertingWebhookDto instantiates a new AlertingWebhookDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingWebhookDto() *AlertingWebhookDto {
	this := AlertingWebhookDto{}
	return &this
}

// NewAlertingWebhookDtoWithDefaults instantiates a new AlertingWebhookDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingWebhookDtoWithDefaults() *AlertingWebhookDto {
	this := AlertingWebhookDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertingWebhookDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingWebhookDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertingWebhookDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertingWebhookDto) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingWebhookDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingWebhookDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AlertingWebhookDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AlertingWebhookDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AlertingWebhookDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AlertingWebhookDto) UnsetUrl() {
	o.Url.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingWebhookDto) GetHeaders() []WebhookHeaderDto {
	if o == nil {
		var ret []WebhookHeaderDto
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingWebhookDto) GetHeadersOk() ([]WebhookHeaderDto, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AlertingWebhookDto) HasHeaders() bool {
	if o != nil && IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []WebhookHeaderDto and assigns it to the Headers field.
func (o *AlertingWebhookDto) SetHeaders(v []WebhookHeaderDto) {
	o.Headers = v
}

func (o AlertingWebhookDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingWebhookDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableAlertingWebhookDto struct {
	value *AlertingWebhookDto
	isSet bool
}

func (v NullableAlertingWebhookDto) Get() *AlertingWebhookDto {
	return v.value
}

func (v *NullableAlertingWebhookDto) Set(val *AlertingWebhookDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingWebhookDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingWebhookDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingWebhookDto(val *AlertingWebhookDto) *NullableAlertingWebhookDto {
	return &NullableAlertingWebhookDto{value: val, isSet: true}
}

func (v NullableAlertingWebhookDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingWebhookDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


