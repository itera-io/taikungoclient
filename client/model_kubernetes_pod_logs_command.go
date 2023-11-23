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

// checks if the KubernetesPodLogsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesPodLogsCommand{}

// KubernetesPodLogsCommand struct for KubernetesPodLogsCommand
type KubernetesPodLogsCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Container NullableString `json:"container,omitempty"`
}

// NewKubernetesPodLogsCommand instantiates a new KubernetesPodLogsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPodLogsCommand() *KubernetesPodLogsCommand {
	this := KubernetesPodLogsCommand{}
	return &this
}

// NewKubernetesPodLogsCommandWithDefaults instantiates a new KubernetesPodLogsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodLogsCommandWithDefaults() *KubernetesPodLogsCommand {
	this := KubernetesPodLogsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubernetesPodLogsCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodLogsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubernetesPodLogsCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubernetesPodLogsCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesPodLogsCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesPodLogsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesPodLogsCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KubernetesPodLogsCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KubernetesPodLogsCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KubernetesPodLogsCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesPodLogsCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesPodLogsCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubernetesPodLogsCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *KubernetesPodLogsCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *KubernetesPodLogsCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *KubernetesPodLogsCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesPodLogsCommand) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesPodLogsCommand) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *KubernetesPodLogsCommand) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *KubernetesPodLogsCommand) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *KubernetesPodLogsCommand) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *KubernetesPodLogsCommand) UnsetContainer() {
	o.Container.Unset()
}

func (o KubernetesPodLogsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesPodLogsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Container.IsSet() {
		toSerialize["container"] = o.Container.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesPodLogsCommand struct {
	value *KubernetesPodLogsCommand
	isSet bool
}

func (v NullableKubernetesPodLogsCommand) Get() *KubernetesPodLogsCommand {
	return v.value
}

func (v *NullableKubernetesPodLogsCommand) Set(val *KubernetesPodLogsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPodLogsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPodLogsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPodLogsCommand(val *KubernetesPodLogsCommand) *NullableKubernetesPodLogsCommand {
	return &NullableKubernetesPodLogsCommand{value: val, isSet: true}
}

func (v NullableKubernetesPodLogsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPodLogsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


