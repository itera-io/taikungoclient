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

// checks if the PatchPdbCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPdbCommand{}

// PatchPdbCommand struct for PatchPdbCommand
type PatchPdbCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Yaml      NullableString `json:"yaml,omitempty"`
	Name      NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewPatchPdbCommand instantiates a new PatchPdbCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPdbCommand() *PatchPdbCommand {
	this := PatchPdbCommand{}
	return &this
}

// NewPatchPdbCommandWithDefaults instantiates a new PatchPdbCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPdbCommandWithDefaults() *PatchPdbCommand {
	this := PatchPdbCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PatchPdbCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPdbCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PatchPdbCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PatchPdbCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchPdbCommand) GetYaml() string {
	if o == nil || IsNil(o.Yaml.Get()) {
		var ret string
		return ret
	}
	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchPdbCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// HasYaml returns a boolean if a field has been set.
func (o *PatchPdbCommand) HasYaml() bool {
	if o != nil && o.Yaml.IsSet() {
		return true
	}

	return false
}

// SetYaml gets a reference to the given NullableString and assigns it to the Yaml field.
func (o *PatchPdbCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}

// SetYamlNil sets the value for Yaml to be an explicit nil
func (o *PatchPdbCommand) SetYamlNil() {
	o.Yaml.Set(nil)
}

// UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
func (o *PatchPdbCommand) UnsetYaml() {
	o.Yaml.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchPdbCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchPdbCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchPdbCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchPdbCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchPdbCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchPdbCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchPdbCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchPdbCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PatchPdbCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PatchPdbCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PatchPdbCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PatchPdbCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o PatchPdbCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPdbCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Yaml.IsSet() {
		toSerialize["yaml"] = o.Yaml.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	return toSerialize, nil
}

type NullablePatchPdbCommand struct {
	value *PatchPdbCommand
	isSet bool
}

func (v NullablePatchPdbCommand) Get() *PatchPdbCommand {
	return v.value
}

func (v *NullablePatchPdbCommand) Set(val *PatchPdbCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPdbCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPdbCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPdbCommand(val *PatchPdbCommand) *NullablePatchPdbCommand {
	return &NullablePatchPdbCommand{value: val, isSet: true}
}

func (v NullablePatchPdbCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPdbCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
