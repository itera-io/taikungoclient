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

// checks if the DuplicateProjectAppCheckerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicateProjectAppCheckerCommand{}

// DuplicateProjectAppCheckerCommand struct for DuplicateProjectAppCheckerCommand
type DuplicateProjectAppCheckerCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewDuplicateProjectAppCheckerCommand instantiates a new DuplicateProjectAppCheckerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicateProjectAppCheckerCommand() *DuplicateProjectAppCheckerCommand {
	this := DuplicateProjectAppCheckerCommand{}
	return &this
}

// NewDuplicateProjectAppCheckerCommandWithDefaults instantiates a new DuplicateProjectAppCheckerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicateProjectAppCheckerCommandWithDefaults() *DuplicateProjectAppCheckerCommand {
	this := DuplicateProjectAppCheckerCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DuplicateProjectAppCheckerCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicateProjectAppCheckerCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DuplicateProjectAppCheckerCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DuplicateProjectAppCheckerCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateProjectAppCheckerCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateProjectAppCheckerCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DuplicateProjectAppCheckerCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DuplicateProjectAppCheckerCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DuplicateProjectAppCheckerCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DuplicateProjectAppCheckerCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DuplicateProjectAppCheckerCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DuplicateProjectAppCheckerCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *DuplicateProjectAppCheckerCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *DuplicateProjectAppCheckerCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *DuplicateProjectAppCheckerCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *DuplicateProjectAppCheckerCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o DuplicateProjectAppCheckerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicateProjectAppCheckerCommand) ToMap() (map[string]interface{}, error) {
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

type NullableDuplicateProjectAppCheckerCommand struct {
	value *DuplicateProjectAppCheckerCommand
	isSet bool
}

func (v NullableDuplicateProjectAppCheckerCommand) Get() *DuplicateProjectAppCheckerCommand {
	return v.value
}

func (v *NullableDuplicateProjectAppCheckerCommand) Set(val *DuplicateProjectAppCheckerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicateProjectAppCheckerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicateProjectAppCheckerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicateProjectAppCheckerCommand(val *DuplicateProjectAppCheckerCommand) *NullableDuplicateProjectAppCheckerCommand {
	return &NullableDuplicateProjectAppCheckerCommand{value: val, isSet: true}
}

func (v NullableDuplicateProjectAppCheckerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicateProjectAppCheckerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


