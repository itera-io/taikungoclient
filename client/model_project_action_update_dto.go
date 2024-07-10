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

// checks if the ProjectActionUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectActionUpdateDto{}

// ProjectActionUpdateDto struct for ProjectActionUpdateDto
type ProjectActionUpdateDto struct {
	Operation NullableString `json:"operation,omitempty"`
	JobUrl NullableString `json:"jobUrl,omitempty"`
	EstimatedTime NullableString `json:"estimatedTime,omitempty"`
}

// NewProjectActionUpdateDto instantiates a new ProjectActionUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectActionUpdateDto() *ProjectActionUpdateDto {
	this := ProjectActionUpdateDto{}
	return &this
}

// NewProjectActionUpdateDtoWithDefaults instantiates a new ProjectActionUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectActionUpdateDtoWithDefaults() *ProjectActionUpdateDto {
	this := ProjectActionUpdateDto{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectActionUpdateDto) GetOperation() string {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectActionUpdateDto) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *ProjectActionUpdateDto) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *ProjectActionUpdateDto) SetOperation(v string) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *ProjectActionUpdateDto) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *ProjectActionUpdateDto) UnsetOperation() {
	o.Operation.Unset()
}

// GetJobUrl returns the JobUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectActionUpdateDto) GetJobUrl() string {
	if o == nil || IsNil(o.JobUrl.Get()) {
		var ret string
		return ret
	}
	return *o.JobUrl.Get()
}

// GetJobUrlOk returns a tuple with the JobUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectActionUpdateDto) GetJobUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobUrl.Get(), o.JobUrl.IsSet()
}

// HasJobUrl returns a boolean if a field has been set.
func (o *ProjectActionUpdateDto) HasJobUrl() bool {
	if o != nil && o.JobUrl.IsSet() {
		return true
	}

	return false
}

// SetJobUrl gets a reference to the given NullableString and assigns it to the JobUrl field.
func (o *ProjectActionUpdateDto) SetJobUrl(v string) {
	o.JobUrl.Set(&v)
}
// SetJobUrlNil sets the value for JobUrl to be an explicit nil
func (o *ProjectActionUpdateDto) SetJobUrlNil() {
	o.JobUrl.Set(nil)
}

// UnsetJobUrl ensures that no value is present for JobUrl, not even an explicit nil
func (o *ProjectActionUpdateDto) UnsetJobUrl() {
	o.JobUrl.Unset()
}

// GetEstimatedTime returns the EstimatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectActionUpdateDto) GetEstimatedTime() string {
	if o == nil || IsNil(o.EstimatedTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedTime.Get()
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectActionUpdateDto) GetEstimatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedTime.Get(), o.EstimatedTime.IsSet()
}

// HasEstimatedTime returns a boolean if a field has been set.
func (o *ProjectActionUpdateDto) HasEstimatedTime() bool {
	if o != nil && o.EstimatedTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedTime gets a reference to the given NullableString and assigns it to the EstimatedTime field.
func (o *ProjectActionUpdateDto) SetEstimatedTime(v string) {
	o.EstimatedTime.Set(&v)
}
// SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil
func (o *ProjectActionUpdateDto) SetEstimatedTimeNil() {
	o.EstimatedTime.Set(nil)
}

// UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
func (o *ProjectActionUpdateDto) UnsetEstimatedTime() {
	o.EstimatedTime.Unset()
}

func (o ProjectActionUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectActionUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	if o.JobUrl.IsSet() {
		toSerialize["jobUrl"] = o.JobUrl.Get()
	}
	if o.EstimatedTime.IsSet() {
		toSerialize["estimatedTime"] = o.EstimatedTime.Get()
	}
	return toSerialize, nil
}

type NullableProjectActionUpdateDto struct {
	value *ProjectActionUpdateDto
	isSet bool
}

func (v NullableProjectActionUpdateDto) Get() *ProjectActionUpdateDto {
	return v.value
}

func (v *NullableProjectActionUpdateDto) Set(val *ProjectActionUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectActionUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectActionUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectActionUpdateDto(val *ProjectActionUpdateDto) *NullableProjectActionUpdateDto {
	return &NullableProjectActionUpdateDto{value: val, isSet: true}
}

func (v NullableProjectActionUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectActionUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

