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

// checks if the HelmReleaseSourceRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmReleaseSourceRef{}

// HelmReleaseSourceRef struct for HelmReleaseSourceRef
type HelmReleaseSourceRef struct {
	Kind NullableString `json:"kind,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewHelmReleaseSourceRef instantiates a new HelmReleaseSourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmReleaseSourceRef() *HelmReleaseSourceRef {
	this := HelmReleaseSourceRef{}
	return &this
}

// NewHelmReleaseSourceRefWithDefaults instantiates a new HelmReleaseSourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmReleaseSourceRefWithDefaults() *HelmReleaseSourceRef {
	this := HelmReleaseSourceRef{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmReleaseSourceRef) GetKind() string {
	if o == nil || IsNil(o.Kind.Get()) {
		var ret string
		return ret
	}
	return *o.Kind.Get()
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmReleaseSourceRef) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kind.Get(), o.Kind.IsSet()
}

// HasKind returns a boolean if a field has been set.
func (o *HelmReleaseSourceRef) HasKind() bool {
	if o != nil && o.Kind.IsSet() {
		return true
	}

	return false
}

// SetKind gets a reference to the given NullableString and assigns it to the Kind field.
func (o *HelmReleaseSourceRef) SetKind(v string) {
	o.Kind.Set(&v)
}
// SetKindNil sets the value for Kind to be an explicit nil
func (o *HelmReleaseSourceRef) SetKindNil() {
	o.Kind.Set(nil)
}

// UnsetKind ensures that no value is present for Kind, not even an explicit nil
func (o *HelmReleaseSourceRef) UnsetKind() {
	o.Kind.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmReleaseSourceRef) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmReleaseSourceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HelmReleaseSourceRef) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HelmReleaseSourceRef) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HelmReleaseSourceRef) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HelmReleaseSourceRef) UnsetName() {
	o.Name.Unset()
}

func (o HelmReleaseSourceRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmReleaseSourceRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind.IsSet() {
		toSerialize["kind"] = o.Kind.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableHelmReleaseSourceRef struct {
	value *HelmReleaseSourceRef
	isSet bool
}

func (v NullableHelmReleaseSourceRef) Get() *HelmReleaseSourceRef {
	return v.value
}

func (v *NullableHelmReleaseSourceRef) Set(val *HelmReleaseSourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmReleaseSourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmReleaseSourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmReleaseSourceRef(val *HelmReleaseSourceRef) *NullableHelmReleaseSourceRef {
	return &NullableHelmReleaseSourceRef{value: val, isSet: true}
}

func (v NullableHelmReleaseSourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmReleaseSourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


