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

// checks if the HelmReleaseSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmReleaseSpec{}

// HelmReleaseSpec struct for HelmReleaseSpec
type HelmReleaseSpec struct {
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	Chart *HelmReleaseChart `json:"chart,omitempty"`
}

// NewHelmReleaseSpec instantiates a new HelmReleaseSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmReleaseSpec() *HelmReleaseSpec {
	this := HelmReleaseSpec{}
	return &this
}

// NewHelmReleaseSpecWithDefaults instantiates a new HelmReleaseSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmReleaseSpecWithDefaults() *HelmReleaseSpec {
	this := HelmReleaseSpec{}
	return &this
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise.
func (o *HelmReleaseSpec) GetTargetNamespace() string {
	if o == nil || IsNil(o.TargetNamespace) {
		var ret string
		return ret
	}
	return *o.TargetNamespace
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmReleaseSpec) GetTargetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNamespace) {
		return nil, false
	}
	return o.TargetNamespace, true
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *HelmReleaseSpec) HasTargetNamespace() bool {
	if o != nil && !IsNil(o.TargetNamespace) {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given string and assigns it to the TargetNamespace field.
func (o *HelmReleaseSpec) SetTargetNamespace(v string) {
	o.TargetNamespace = &v
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *HelmReleaseSpec) GetChart() HelmReleaseChart {
	if o == nil || IsNil(o.Chart) {
		var ret HelmReleaseChart
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmReleaseSpec) GetChartOk() (*HelmReleaseChart, bool) {
	if o == nil || IsNil(o.Chart) {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *HelmReleaseSpec) HasChart() bool {
	if o != nil && !IsNil(o.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given HelmReleaseChart and assigns it to the Chart field.
func (o *HelmReleaseSpec) SetChart(v HelmReleaseChart) {
	o.Chart = &v
}

func (o HelmReleaseSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmReleaseSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetNamespace) {
		toSerialize["targetNamespace"] = o.TargetNamespace
	}
	if !IsNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	return toSerialize, nil
}

type NullableHelmReleaseSpec struct {
	value *HelmReleaseSpec
	isSet bool
}

func (v NullableHelmReleaseSpec) Get() *HelmReleaseSpec {
	return v.value
}

func (v *NullableHelmReleaseSpec) Set(val *HelmReleaseSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmReleaseSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmReleaseSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmReleaseSpec(val *HelmReleaseSpec) *NullableHelmReleaseSpec {
	return &NullableHelmReleaseSpec{value: val, isSet: true}
}

func (v NullableHelmReleaseSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmReleaseSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


