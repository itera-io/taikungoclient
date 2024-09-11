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

// checks if the KubeConfigInteractiveShellCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubeConfigInteractiveShellCommand{}

// KubeConfigInteractiveShellCommand struct for KubeConfigInteractiveShellCommand
type KubeConfigInteractiveShellCommand struct {
	KubeConfigId *int32 `json:"kubeConfigId,omitempty"`
	Token *string `json:"token,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewKubeConfigInteractiveShellCommand instantiates a new KubeConfigInteractiveShellCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeConfigInteractiveShellCommand() *KubeConfigInteractiveShellCommand {
	this := KubeConfigInteractiveShellCommand{}
	return &this
}

// NewKubeConfigInteractiveShellCommandWithDefaults instantiates a new KubeConfigInteractiveShellCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeConfigInteractiveShellCommandWithDefaults() *KubeConfigInteractiveShellCommand {
	this := KubeConfigInteractiveShellCommand{}
	return &this
}

// GetKubeConfigId returns the KubeConfigId field value if set, zero value otherwise.
func (o *KubeConfigInteractiveShellCommand) GetKubeConfigId() int32 {
	if o == nil || IsNil(o.KubeConfigId) {
		var ret int32
		return ret
	}
	return *o.KubeConfigId
}

// GetKubeConfigIdOk returns a tuple with the KubeConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigInteractiveShellCommand) GetKubeConfigIdOk() (*int32, bool) {
	if o == nil || IsNil(o.KubeConfigId) {
		return nil, false
	}
	return o.KubeConfigId, true
}

// HasKubeConfigId returns a boolean if a field has been set.
func (o *KubeConfigInteractiveShellCommand) HasKubeConfigId() bool {
	if o != nil && !IsNil(o.KubeConfigId) {
		return true
	}

	return false
}

// SetKubeConfigId gets a reference to the given int32 and assigns it to the KubeConfigId field.
func (o *KubeConfigInteractiveShellCommand) SetKubeConfigId(v int32) {
	o.KubeConfigId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KubeConfigInteractiveShellCommand) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigInteractiveShellCommand) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KubeConfigInteractiveShellCommand) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KubeConfigInteractiveShellCommand) SetToken(v string) {
	o.Token = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubeConfigInteractiveShellCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigInteractiveShellCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubeConfigInteractiveShellCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubeConfigInteractiveShellCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o KubeConfigInteractiveShellCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubeConfigInteractiveShellCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KubeConfigId) {
		toSerialize["kubeConfigId"] = o.KubeConfigId
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableKubeConfigInteractiveShellCommand struct {
	value *KubeConfigInteractiveShellCommand
	isSet bool
}

func (v NullableKubeConfigInteractiveShellCommand) Get() *KubeConfigInteractiveShellCommand {
	return v.value
}

func (v *NullableKubeConfigInteractiveShellCommand) Set(val *KubeConfigInteractiveShellCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeConfigInteractiveShellCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeConfigInteractiveShellCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeConfigInteractiveShellCommand(val *KubeConfigInteractiveShellCommand) *NullableKubeConfigInteractiveShellCommand {
	return &NullableKubeConfigInteractiveShellCommand{value: val, isSet: true}
}

func (v NullableKubeConfigInteractiveShellCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeConfigInteractiveShellCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


