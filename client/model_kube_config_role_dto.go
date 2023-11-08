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

// checks if the KubeConfigRoleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubeConfigRoleDto{}

// KubeConfigRoleDto struct for KubeConfigRoleDto
type KubeConfigRoleDto struct {
	Id   *int32         `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewKubeConfigRoleDto instantiates a new KubeConfigRoleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeConfigRoleDto() *KubeConfigRoleDto {
	this := KubeConfigRoleDto{}
	return &this
}

// NewKubeConfigRoleDtoWithDefaults instantiates a new KubeConfigRoleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeConfigRoleDtoWithDefaults() *KubeConfigRoleDto {
	this := KubeConfigRoleDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubeConfigRoleDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubeConfigRoleDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubeConfigRoleDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubeConfigRoleDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubeConfigRoleDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubeConfigRoleDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KubeConfigRoleDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KubeConfigRoleDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *KubeConfigRoleDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KubeConfigRoleDto) UnsetName() {
	o.Name.Unset()
}

func (o KubeConfigRoleDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubeConfigRoleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableKubeConfigRoleDto struct {
	value *KubeConfigRoleDto
	isSet bool
}

func (v NullableKubeConfigRoleDto) Get() *KubeConfigRoleDto {
	return v.value
}

func (v *NullableKubeConfigRoleDto) Set(val *KubeConfigRoleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeConfigRoleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeConfigRoleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeConfigRoleDto(val *KubeConfigRoleDto) *NullableKubeConfigRoleDto {
	return &NullableKubeConfigRoleDto{value: val, isSet: true}
}

func (v NullableKubeConfigRoleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeConfigRoleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
