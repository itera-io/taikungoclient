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

// checks if the PrometheusMetricsAutocompleteCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusMetricsAutocompleteCommand{}

// PrometheusMetricsAutocompleteCommand struct for PrometheusMetricsAutocompleteCommand
type PrometheusMetricsAutocompleteCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewPrometheusMetricsAutocompleteCommand instantiates a new PrometheusMetricsAutocompleteCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusMetricsAutocompleteCommand() *PrometheusMetricsAutocompleteCommand {
	this := PrometheusMetricsAutocompleteCommand{}
	return &this
}

// NewPrometheusMetricsAutocompleteCommandWithDefaults instantiates a new PrometheusMetricsAutocompleteCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusMetricsAutocompleteCommandWithDefaults() *PrometheusMetricsAutocompleteCommand {
	this := PrometheusMetricsAutocompleteCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PrometheusMetricsAutocompleteCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMetricsAutocompleteCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PrometheusMetricsAutocompleteCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PrometheusMetricsAutocompleteCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o PrometheusMetricsAutocompleteCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusMetricsAutocompleteCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullablePrometheusMetricsAutocompleteCommand struct {
	value *PrometheusMetricsAutocompleteCommand
	isSet bool
}

func (v NullablePrometheusMetricsAutocompleteCommand) Get() *PrometheusMetricsAutocompleteCommand {
	return v.value
}

func (v *NullablePrometheusMetricsAutocompleteCommand) Set(val *PrometheusMetricsAutocompleteCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusMetricsAutocompleteCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusMetricsAutocompleteCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusMetricsAutocompleteCommand(val *PrometheusMetricsAutocompleteCommand) *NullablePrometheusMetricsAutocompleteCommand {
	return &NullablePrometheusMetricsAutocompleteCommand{value: val, isSet: true}
}

func (v NullablePrometheusMetricsAutocompleteCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusMetricsAutocompleteCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


