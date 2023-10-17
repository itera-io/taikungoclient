/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the ChatCompletionsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionsCommand{}

// ChatCompletionsCommand struct for ChatCompletionsCommand
type ChatCompletionsCommand struct {
	Messages interface{} `json:"messages,omitempty"`
}

// NewChatCompletionsCommand instantiates a new ChatCompletionsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionsCommand() *ChatCompletionsCommand {
	this := ChatCompletionsCommand{}
	return &this
}

// NewChatCompletionsCommandWithDefaults instantiates a new ChatCompletionsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionsCommandWithDefaults() *ChatCompletionsCommand {
	this := ChatCompletionsCommand{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionsCommand) GetMessages() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionsCommand) GetMessagesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ChatCompletionsCommand) HasMessages() bool {
	if o != nil && IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given interface{} and assigns it to the Messages field.
func (o *ChatCompletionsCommand) SetMessages(v interface{}) {
	o.Messages = v
}

func (o ChatCompletionsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableChatCompletionsCommand struct {
	value *ChatCompletionsCommand
	isSet bool
}

func (v NullableChatCompletionsCommand) Get() *ChatCompletionsCommand {
	return v.value
}

func (v *NullableChatCompletionsCommand) Set(val *ChatCompletionsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionsCommand(val *ChatCompletionsCommand) *NullableChatCompletionsCommand {
	return &NullableChatCompletionsCommand{value: val, isSet: true}
}

func (v NullableChatCompletionsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
