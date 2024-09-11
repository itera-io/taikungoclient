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

// checks if the EditArticleCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditArticleCommand{}

// EditArticleCommand struct for EditArticleCommand
type EditArticleCommand struct {
	MessageId *string `json:"messageId,omitempty"`
	Body *string `json:"body,omitempty"`
}

// NewEditArticleCommand instantiates a new EditArticleCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditArticleCommand() *EditArticleCommand {
	this := EditArticleCommand{}
	return &this
}

// NewEditArticleCommandWithDefaults instantiates a new EditArticleCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditArticleCommandWithDefaults() *EditArticleCommand {
	this := EditArticleCommand{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EditArticleCommand) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditArticleCommand) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EditArticleCommand) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *EditArticleCommand) SetMessageId(v string) {
	o.MessageId = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *EditArticleCommand) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditArticleCommand) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *EditArticleCommand) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *EditArticleCommand) SetBody(v string) {
	o.Body = &v
}

func (o EditArticleCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditArticleCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableEditArticleCommand struct {
	value *EditArticleCommand
	isSet bool
}

func (v NullableEditArticleCommand) Get() *EditArticleCommand {
	return v.value
}

func (v *NullableEditArticleCommand) Set(val *EditArticleCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditArticleCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditArticleCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditArticleCommand(val *EditArticleCommand) *NullableEditArticleCommand {
	return &NullableEditArticleCommand{value: val, isSet: true}
}

func (v NullableEditArticleCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditArticleCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


