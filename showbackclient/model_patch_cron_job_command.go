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

// checks if the PatchCronJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchCronJobCommand{}

// PatchCronJobCommand struct for PatchCronJobCommand
type PatchCronJobCommand struct {
	ProjectId *int32         `json:"projectId,omitempty"`
	Yaml      NullableString `json:"yaml,omitempty"`
	Name      NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewPatchCronJobCommand instantiates a new PatchCronJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchCronJobCommand() *PatchCronJobCommand {
	this := PatchCronJobCommand{}
	return &this
}

// NewPatchCronJobCommandWithDefaults instantiates a new PatchCronJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchCronJobCommandWithDefaults() *PatchCronJobCommand {
	this := PatchCronJobCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PatchCronJobCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchCronJobCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PatchCronJobCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PatchCronJobCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCronJobCommand) GetYaml() string {
	if o == nil || IsNil(o.Yaml.Get()) {
		var ret string
		return ret
	}
	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCronJobCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// HasYaml returns a boolean if a field has been set.
func (o *PatchCronJobCommand) HasYaml() bool {
	if o != nil && o.Yaml.IsSet() {
		return true
	}

	return false
}

// SetYaml gets a reference to the given NullableString and assigns it to the Yaml field.
func (o *PatchCronJobCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}

// SetYamlNil sets the value for Yaml to be an explicit nil
func (o *PatchCronJobCommand) SetYamlNil() {
	o.Yaml.Set(nil)
}

// UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
func (o *PatchCronJobCommand) UnsetYaml() {
	o.Yaml.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCronJobCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCronJobCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchCronJobCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchCronJobCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchCronJobCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchCronJobCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCronJobCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCronJobCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PatchCronJobCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PatchCronJobCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PatchCronJobCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PatchCronJobCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o PatchCronJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCronJobCommand) ToMap() (map[string]interface{}, error) {
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

type NullablePatchCronJobCommand struct {
	value *PatchCronJobCommand
	isSet bool
}

func (v NullablePatchCronJobCommand) Get() *PatchCronJobCommand {
	return v.value
}

func (v *NullablePatchCronJobCommand) Set(val *PatchCronJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCronJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCronJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCronJobCommand(val *PatchCronJobCommand) *NullablePatchCronJobCommand {
	return &NullablePatchCronJobCommand{value: val, isSet: true}
}

func (v NullablePatchCronJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCronJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
