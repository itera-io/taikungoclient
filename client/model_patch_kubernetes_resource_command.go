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
	"bytes"
	"fmt"
)

// checks if the PatchKubernetesResourceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchKubernetesResourceCommand{}

// PatchKubernetesResourceCommand struct for PatchKubernetesResourceCommand
type PatchKubernetesResourceCommand struct {
	ProjectId int32 `json:"projectId"`
	Yaml NullableString `json:"yaml"`
	Name NullableString `json:"name"`
	Namespace NullableString `json:"namespace,omitempty"`
	Kind EKubernetesResource `json:"kind"`
}

type _PatchKubernetesResourceCommand PatchKubernetesResourceCommand

// NewPatchKubernetesResourceCommand instantiates a new PatchKubernetesResourceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchKubernetesResourceCommand(projectId int32, yaml NullableString, name NullableString, kind EKubernetesResource) *PatchKubernetesResourceCommand {
	this := PatchKubernetesResourceCommand{}
	this.ProjectId = projectId
	this.Yaml = yaml
	this.Name = name
	this.Kind = kind
	return &this
}

// NewPatchKubernetesResourceCommandWithDefaults instantiates a new PatchKubernetesResourceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchKubernetesResourceCommandWithDefaults() *PatchKubernetesResourceCommand {
	this := PatchKubernetesResourceCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchKubernetesResourceCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchKubernetesResourceCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchKubernetesResourceCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetYaml returns the Yaml field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PatchKubernetesResourceCommand) GetYaml() string {
	if o == nil || o.Yaml.Get() == nil {
		var ret string
		return ret
	}

	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchKubernetesResourceCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// SetYaml sets field value
func (o *PatchKubernetesResourceCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PatchKubernetesResourceCommand) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchKubernetesResourceCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PatchKubernetesResourceCommand) SetName(v string) {
	o.Name.Set(&v)
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchKubernetesResourceCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchKubernetesResourceCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PatchKubernetesResourceCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PatchKubernetesResourceCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PatchKubernetesResourceCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PatchKubernetesResourceCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetKind returns the Kind field value
func (o *PatchKubernetesResourceCommand) GetKind() EKubernetesResource {
	if o == nil {
		var ret EKubernetesResource
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PatchKubernetesResourceCommand) GetKindOk() (*EKubernetesResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PatchKubernetesResourceCommand) SetKind(v EKubernetesResource) {
	o.Kind = v
}

func (o PatchKubernetesResourceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchKubernetesResourceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["yaml"] = o.Yaml.Get()
	toSerialize["name"] = o.Name.Get()
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	toSerialize["kind"] = o.Kind
	return toSerialize, nil
}

func (o *PatchKubernetesResourceCommand) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
		"yaml",
		"name",
		"kind",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPatchKubernetesResourceCommand := _PatchKubernetesResourceCommand{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchKubernetesResourceCommand)

	if err != nil {
		return err
	}

	*o = PatchKubernetesResourceCommand(varPatchKubernetesResourceCommand)

	return err
}

type NullablePatchKubernetesResourceCommand struct {
	value *PatchKubernetesResourceCommand
	isSet bool
}

func (v NullablePatchKubernetesResourceCommand) Get() *PatchKubernetesResourceCommand {
	return v.value
}

func (v *NullablePatchKubernetesResourceCommand) Set(val *PatchKubernetesResourceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchKubernetesResourceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchKubernetesResourceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchKubernetesResourceCommand(val *PatchKubernetesResourceCommand) *NullablePatchKubernetesResourceCommand {
	return &NullablePatchKubernetesResourceCommand{value: val, isSet: true}
}

func (v NullablePatchKubernetesResourceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchKubernetesResourceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

