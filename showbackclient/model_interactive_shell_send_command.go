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

// checks if the InteractiveShellSendCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractiveShellSendCommand{}

// InteractiveShellSendCommand struct for InteractiveShellSendCommand
type InteractiveShellSendCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Token     NullableString `json:"token,omitempty"`
}

// NewInteractiveShellSendCommand instantiates a new InteractiveShellSendCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractiveShellSendCommand() *InteractiveShellSendCommand {
	this := InteractiveShellSendCommand{}
	return &this
}

// NewInteractiveShellSendCommandWithDefaults instantiates a new InteractiveShellSendCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractiveShellSendCommandWithDefaults() *InteractiveShellSendCommand {
	this := InteractiveShellSendCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InteractiveShellSendCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InteractiveShellSendCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InteractiveShellSendCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *InteractiveShellSendCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractiveShellSendCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractiveShellSendCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *InteractiveShellSendCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *InteractiveShellSendCommand) SetToken(v string) {
	o.Token.Set(&v)
}

// SetTokenNil sets the value for Token to be an explicit nil
func (o *InteractiveShellSendCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *InteractiveShellSendCommand) UnsetToken() {
	o.Token.Unset()
}

func (o InteractiveShellSendCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractiveShellSendCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableInteractiveShellSendCommand struct {
	value *InteractiveShellSendCommand
	isSet bool
}

func (v NullableInteractiveShellSendCommand) Get() *InteractiveShellSendCommand {
	return v.value
}

func (v *NullableInteractiveShellSendCommand) Set(val *InteractiveShellSendCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractiveShellSendCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractiveShellSendCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractiveShellSendCommand(val *InteractiveShellSendCommand) *NullableInteractiveShellSendCommand {
	return &NullableInteractiveShellSendCommand{value: val, isSet: true}
}

func (v NullableInteractiveShellSendCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractiveShellSendCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
