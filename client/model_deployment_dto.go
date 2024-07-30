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

// checks if the DeploymentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDto{}

// DeploymentDto struct for DeploymentDto
type DeploymentDto struct {
	MetadataName NullableString `json:"metadataName,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Age NullableString `json:"age,omitempty"`
}

// NewDeploymentDto instantiates a new DeploymentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDto() *DeploymentDto {
	this := DeploymentDto{}
	return &this
}

// NewDeploymentDtoWithDefaults instantiates a new DeploymentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDtoWithDefaults() *DeploymentDto {
	this := DeploymentDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *DeploymentDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *DeploymentDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}
// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *DeploymentDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *DeploymentDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *DeploymentDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *DeploymentDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *DeploymentDto) UnsetStatus() {
	o.Status.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *DeploymentDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *DeploymentDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *DeploymentDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *DeploymentDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentDto) GetAge() string {
	if o == nil || IsNil(o.Age.Get()) {
		var ret string
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *DeploymentDto) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableString and assigns it to the Age field.
func (o *DeploymentDto) SetAge(v string) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *DeploymentDto) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *DeploymentDto) UnsetAge() {
	o.Age.Unset()
}

func (o DeploymentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	return toSerialize, nil
}

type NullableDeploymentDto struct {
	value *DeploymentDto
	isSet bool
}

func (v NullableDeploymentDto) Get() *DeploymentDto {
	return v.value
}

func (v *NullableDeploymentDto) Set(val *DeploymentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDto(val *DeploymentDto) *NullableDeploymentDto {
	return &NullableDeploymentDto{value: val, isSet: true}
}

func (v NullableDeploymentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


