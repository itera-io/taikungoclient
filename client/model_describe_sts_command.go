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

// checks if the DescribeStsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeStsCommand{}

// DescribeStsCommand struct for DescribeStsCommand
type DescribeStsCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewDescribeStsCommand instantiates a new DescribeStsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeStsCommand() *DescribeStsCommand {
	this := DescribeStsCommand{}
	return &this
}

// NewDescribeStsCommandWithDefaults instantiates a new DescribeStsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeStsCommandWithDefaults() *DescribeStsCommand {
	this := DescribeStsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DescribeStsCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeStsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DescribeStsCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DescribeStsCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeStsCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeStsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DescribeStsCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DescribeStsCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DescribeStsCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DescribeStsCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeStsCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeStsCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *DescribeStsCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *DescribeStsCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *DescribeStsCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *DescribeStsCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o DescribeStsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeStsCommand) ToMap() (map[string]interface{}, error) {
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

type NullableDescribeStsCommand struct {
	value *DescribeStsCommand
	isSet bool
}

func (v NullableDescribeStsCommand) Get() *DescribeStsCommand {
	return v.value
}

func (v *NullableDescribeStsCommand) Set(val *DescribeStsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeStsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeStsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeStsCommand(val *DescribeStsCommand) *NullableDescribeStsCommand {
	return &NullableDescribeStsCommand{value: val, isSet: true}
}

func (v NullableDescribeStsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeStsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


