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

// checks if the PodDisruptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodDisruptionDto{}

// PodDisruptionDto struct for PodDisruptionDto
type PodDisruptionDto struct {
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	MinAvailable interface{} `json:"minAvailable,omitempty"`
	MaxAvailable interface{} `json:"maxAvailable,omitempty"`
	AllowedDisruptions interface{} `json:"allowedDisruptions,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
}

// NewPodDisruptionDto instantiates a new PodDisruptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodDisruptionDto() *PodDisruptionDto {
	this := PodDisruptionDto{}
	return &this
}

// NewPodDisruptionDtoWithDefaults instantiates a new PodDisruptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodDisruptionDtoWithDefaults() *PodDisruptionDto {
	this := PodDisruptionDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PodDisruptionDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PodDisruptionDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PodDisruptionDto) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PodDisruptionDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PodDisruptionDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PodDisruptionDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetMinAvailable returns the MinAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetMinAvailable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinAvailable
}

// GetMinAvailableOk returns a tuple with the MinAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetMinAvailableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinAvailable) {
		return nil, false
	}
	return &o.MinAvailable, true
}

// HasMinAvailable returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasMinAvailable() bool {
	if o != nil && !IsNil(o.MinAvailable) {
		return true
	}

	return false
}

// SetMinAvailable gets a reference to the given interface{} and assigns it to the MinAvailable field.
func (o *PodDisruptionDto) SetMinAvailable(v interface{}) {
	o.MinAvailable = v
}

// GetMaxAvailable returns the MaxAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetMaxAvailable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxAvailable
}

// GetMaxAvailableOk returns a tuple with the MaxAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetMaxAvailableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxAvailable) {
		return nil, false
	}
	return &o.MaxAvailable, true
}

// HasMaxAvailable returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasMaxAvailable() bool {
	if o != nil && !IsNil(o.MaxAvailable) {
		return true
	}

	return false
}

// SetMaxAvailable gets a reference to the given interface{} and assigns it to the MaxAvailable field.
func (o *PodDisruptionDto) SetMaxAvailable(v interface{}) {
	o.MaxAvailable = v
}

// GetAllowedDisruptions returns the AllowedDisruptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetAllowedDisruptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowedDisruptions
}

// GetAllowedDisruptionsOk returns a tuple with the AllowedDisruptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetAllowedDisruptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedDisruptions) {
		return nil, false
	}
	return &o.AllowedDisruptions, true
}

// HasAllowedDisruptions returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasAllowedDisruptions() bool {
	if o != nil && !IsNil(o.AllowedDisruptions) {
		return true
	}

	return false
}

// SetAllowedDisruptions gets a reference to the given interface{} and assigns it to the AllowedDisruptions field.
func (o *PodDisruptionDto) SetAllowedDisruptions(v interface{}) {
	o.AllowedDisruptions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDisruptionDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PodDisruptionDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *PodDisruptionDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PodDisruptionDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PodDisruptionDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

func (o PodDisruptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodDisruptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.MinAvailable != nil {
		toSerialize["minAvailable"] = o.MinAvailable
	}
	if o.MaxAvailable != nil {
		toSerialize["maxAvailable"] = o.MaxAvailable
	}
	if o.AllowedDisruptions != nil {
		toSerialize["allowedDisruptions"] = o.AllowedDisruptions
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	return toSerialize, nil
}

type NullablePodDisruptionDto struct {
	value *PodDisruptionDto
	isSet bool
}

func (v NullablePodDisruptionDto) Get() *PodDisruptionDto {
	return v.value
}

func (v *NullablePodDisruptionDto) Set(val *PodDisruptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodDisruptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodDisruptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodDisruptionDto(val *PodDisruptionDto) *NullablePodDisruptionDto {
	return &NullablePodDisruptionDto{value: val, isSet: true}
}

func (v NullablePodDisruptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodDisruptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


