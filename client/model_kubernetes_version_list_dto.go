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

// checks if the KubernetesVersionListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesVersionListDto{}

// KubernetesVersionListDto struct for KubernetesVersionListDto
type KubernetesVersionListDto struct {
	Version NullableString `json:"version,omitempty"`
	IsKubevapEnabled *bool `json:"isKubevapEnabled,omitempty"`
}

// NewKubernetesVersionListDto instantiates a new KubernetesVersionListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVersionListDto() *KubernetesVersionListDto {
	this := KubernetesVersionListDto{}
	return &this
}

// NewKubernetesVersionListDtoWithDefaults instantiates a new KubernetesVersionListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVersionListDtoWithDefaults() *KubernetesVersionListDto {
	this := KubernetesVersionListDto{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVersionListDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVersionListDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesVersionListDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *KubernetesVersionListDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *KubernetesVersionListDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *KubernetesVersionListDto) UnsetVersion() {
	o.Version.Unset()
}

// GetIsKubevapEnabled returns the IsKubevapEnabled field value if set, zero value otherwise.
func (o *KubernetesVersionListDto) GetIsKubevapEnabled() bool {
	if o == nil || IsNil(o.IsKubevapEnabled) {
		var ret bool
		return ret
	}
	return *o.IsKubevapEnabled
}

// GetIsKubevapEnabledOk returns a tuple with the IsKubevapEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersionListDto) GetIsKubevapEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsKubevapEnabled) {
		return nil, false
	}
	return o.IsKubevapEnabled, true
}

// HasIsKubevapEnabled returns a boolean if a field has been set.
func (o *KubernetesVersionListDto) HasIsKubevapEnabled() bool {
	if o != nil && !IsNil(o.IsKubevapEnabled) {
		return true
	}

	return false
}

// SetIsKubevapEnabled gets a reference to the given bool and assigns it to the IsKubevapEnabled field.
func (o *KubernetesVersionListDto) SetIsKubevapEnabled(v bool) {
	o.IsKubevapEnabled = &v
}

func (o KubernetesVersionListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesVersionListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.IsKubevapEnabled) {
		toSerialize["isKubevapEnabled"] = o.IsKubevapEnabled
	}
	return toSerialize, nil
}

type NullableKubernetesVersionListDto struct {
	value *KubernetesVersionListDto
	isSet bool
}

func (v NullableKubernetesVersionListDto) Get() *KubernetesVersionListDto {
	return v.value
}

func (v *NullableKubernetesVersionListDto) Set(val *KubernetesVersionListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVersionListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVersionListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVersionListDto(val *KubernetesVersionListDto) *NullableKubernetesVersionListDto {
	return &NullableKubernetesVersionListDto{value: val, isSet: true}
}

func (v NullableKubernetesVersionListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVersionListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


