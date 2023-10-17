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

// checks if the HelmStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmStatus{}

// HelmStatus struct for HelmStatus
type HelmStatus struct {
	Conditions         []Condition    `json:"conditions,omitempty"`
	Failures           *int64         `json:"failures,omitempty"`
	HelmChart          NullableString `json:"helmChart,omitempty"`
	ObservedGeneration *int64         `json:"observedGeneration,omitempty"`
}

// NewHelmStatus instantiates a new HelmStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmStatus() *HelmStatus {
	this := HelmStatus{}
	return &this
}

// NewHelmStatusWithDefaults instantiates a new HelmStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmStatusWithDefaults() *HelmStatus {
	this := HelmStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmStatus) GetConditions() []Condition {
	if o == nil {
		var ret []Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmStatus) GetConditionsOk() ([]Condition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *HelmStatus) HasConditions() bool {
	if o != nil && IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []Condition and assigns it to the Conditions field.
func (o *HelmStatus) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetFailures returns the Failures field value if set, zero value otherwise.
func (o *HelmStatus) GetFailures() int64 {
	if o == nil || IsNil(o.Failures) {
		var ret int64
		return ret
	}
	return *o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmStatus) GetFailuresOk() (*int64, bool) {
	if o == nil || IsNil(o.Failures) {
		return nil, false
	}
	return o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *HelmStatus) HasFailures() bool {
	if o != nil && !IsNil(o.Failures) {
		return true
	}

	return false
}

// SetFailures gets a reference to the given int64 and assigns it to the Failures field.
func (o *HelmStatus) SetFailures(v int64) {
	o.Failures = &v
}

// GetHelmChart returns the HelmChart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmStatus) GetHelmChart() string {
	if o == nil || IsNil(o.HelmChart.Get()) {
		var ret string
		return ret
	}
	return *o.HelmChart.Get()
}

// GetHelmChartOk returns a tuple with the HelmChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmStatus) GetHelmChartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HelmChart.Get(), o.HelmChart.IsSet()
}

// HasHelmChart returns a boolean if a field has been set.
func (o *HelmStatus) HasHelmChart() bool {
	if o != nil && o.HelmChart.IsSet() {
		return true
	}

	return false
}

// SetHelmChart gets a reference to the given NullableString and assigns it to the HelmChart field.
func (o *HelmStatus) SetHelmChart(v string) {
	o.HelmChart.Set(&v)
}

// SetHelmChartNil sets the value for HelmChart to be an explicit nil
func (o *HelmStatus) SetHelmChartNil() {
	o.HelmChart.Set(nil)
}

// UnsetHelmChart ensures that no value is present for HelmChart, not even an explicit nil
func (o *HelmStatus) UnsetHelmChart() {
	o.HelmChart.Unset()
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *HelmStatus) GetObservedGeneration() int64 {
	if o == nil || IsNil(o.ObservedGeneration) {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmStatus) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.ObservedGeneration) {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *HelmStatus) HasObservedGeneration() bool {
	if o != nil && !IsNil(o.ObservedGeneration) {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *HelmStatus) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

func (o HelmStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Failures) {
		toSerialize["failures"] = o.Failures
	}
	if o.HelmChart.IsSet() {
		toSerialize["helmChart"] = o.HelmChart.Get()
	}
	if !IsNil(o.ObservedGeneration) {
		toSerialize["observedGeneration"] = o.ObservedGeneration
	}
	return toSerialize, nil
}

type NullableHelmStatus struct {
	value *HelmStatus
	isSet bool
}

func (v NullableHelmStatus) Get() *HelmStatus {
	return v.value
}

func (v *NullableHelmStatus) Set(val *HelmStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmStatus(val *HelmStatus) *NullableHelmStatus {
	return &NullableHelmStatus{value: val, isSet: true}
}

func (v NullableHelmStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
