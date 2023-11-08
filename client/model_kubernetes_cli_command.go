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

// checks if the KubernetesCliCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCliCommand{}

// KubernetesCliCommand struct for KubernetesCliCommand
type KubernetesCliCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	KubeConfigId NullableInt32 `json:"kubeConfigId,omitempty"`
}

// NewKubernetesCliCommand instantiates a new KubernetesCliCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCliCommand() *KubernetesCliCommand {
	this := KubernetesCliCommand{}
	return &this
}

// NewKubernetesCliCommandWithDefaults instantiates a new KubernetesCliCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCliCommandWithDefaults() *KubernetesCliCommand {
	this := KubernetesCliCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubernetesCliCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCliCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubernetesCliCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubernetesCliCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetKubeConfigId returns the KubeConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesCliCommand) GetKubeConfigId() int32 {
	if o == nil || IsNil(o.KubeConfigId.Get()) {
		var ret int32
		return ret
	}
	return *o.KubeConfigId.Get()
}

// GetKubeConfigIdOk returns a tuple with the KubeConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCliCommand) GetKubeConfigIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubeConfigId.Get(), o.KubeConfigId.IsSet()
}

// HasKubeConfigId returns a boolean if a field has been set.
func (o *KubernetesCliCommand) HasKubeConfigId() bool {
	if o != nil && o.KubeConfigId.IsSet() {
		return true
	}

	return false
}

// SetKubeConfigId gets a reference to the given NullableInt32 and assigns it to the KubeConfigId field.
func (o *KubernetesCliCommand) SetKubeConfigId(v int32) {
	o.KubeConfigId.Set(&v)
}
// SetKubeConfigIdNil sets the value for KubeConfigId to be an explicit nil
func (o *KubernetesCliCommand) SetKubeConfigIdNil() {
	o.KubeConfigId.Set(nil)
}

// UnsetKubeConfigId ensures that no value is present for KubeConfigId, not even an explicit nil
func (o *KubernetesCliCommand) UnsetKubeConfigId() {
	o.KubeConfigId.Unset()
}

func (o KubernetesCliCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCliCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.KubeConfigId.IsSet() {
		toSerialize["kubeConfigId"] = o.KubeConfigId.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesCliCommand struct {
	value *KubernetesCliCommand
	isSet bool
}

func (v NullableKubernetesCliCommand) Get() *KubernetesCliCommand {
	return v.value
}

func (v *NullableKubernetesCliCommand) Set(val *KubernetesCliCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCliCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCliCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCliCommand(val *KubernetesCliCommand) *NullableKubernetesCliCommand {
	return &NullableKubernetesCliCommand{value: val, isSet: true}
}

func (v NullableKubernetesCliCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCliCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


