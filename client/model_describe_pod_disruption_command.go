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

// checks if the DescribePodDisruptionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribePodDisruptionCommand{}

// DescribePodDisruptionCommand struct for DescribePodDisruptionCommand
type DescribePodDisruptionCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Name      NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewDescribePodDisruptionCommand instantiates a new DescribePodDisruptionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePodDisruptionCommand() *DescribePodDisruptionCommand {
	this := DescribePodDisruptionCommand{}
	return &this
}

// NewDescribePodDisruptionCommandWithDefaults instantiates a new DescribePodDisruptionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePodDisruptionCommandWithDefaults() *DescribePodDisruptionCommand {
	this := DescribePodDisruptionCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DescribePodDisruptionCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePodDisruptionCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DescribePodDisruptionCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DescribePodDisruptionCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribePodDisruptionCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribePodDisruptionCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DescribePodDisruptionCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DescribePodDisruptionCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DescribePodDisruptionCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DescribePodDisruptionCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribePodDisruptionCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribePodDisruptionCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *DescribePodDisruptionCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *DescribePodDisruptionCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *DescribePodDisruptionCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *DescribePodDisruptionCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o DescribePodDisruptionCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePodDisruptionCommand) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableDescribePodDisruptionCommand struct {
	value *DescribePodDisruptionCommand
	isSet bool
}

func (v NullableDescribePodDisruptionCommand) Get() *DescribePodDisruptionCommand {
	return v.value
}

func (v *NullableDescribePodDisruptionCommand) Set(val *DescribePodDisruptionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePodDisruptionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePodDisruptionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePodDisruptionCommand(val *DescribePodDisruptionCommand) *NullableDescribePodDisruptionCommand {
	return &NullableDescribePodDisruptionCommand{value: val, isSet: true}
}

func (v NullableDescribePodDisruptionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePodDisruptionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
