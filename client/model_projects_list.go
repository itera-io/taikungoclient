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
	"bytes"
	"fmt"
)

// checks if the ProjectsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsList{}

// ProjectsList struct for ProjectsList
type ProjectsList struct {
	Data []ProjectListDetailDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _ProjectsList ProjectsList

// NewProjectsList instantiates a new ProjectsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsList(data []ProjectListDetailDto, totalCount int32) *ProjectsList {
	this := ProjectsList{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewProjectsListWithDefaults instantiates a new ProjectsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsListWithDefaults() *ProjectsList {
	this := ProjectsList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []ProjectListDetailDto will be returned
func (o *ProjectsList) GetData() []ProjectListDetailDto {
	if o == nil {
		var ret []ProjectListDetailDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsList) GetDataOk() ([]ProjectListDetailDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ProjectsList) SetData(v []ProjectListDetailDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *ProjectsList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ProjectsList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ProjectsList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o ProjectsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ProjectsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"totalCount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectsList := _ProjectsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectsList)

	if err != nil {
		return err
	}

	*o = ProjectsList(varProjectsList)

	return err
}

type NullableProjectsList struct {
	value *ProjectsList
	isSet bool
}

func (v NullableProjectsList) Get() *ProjectsList {
	return v.value
}

func (v *NullableProjectsList) Set(val *ProjectsList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsList(val *ProjectsList) *NullableProjectsList {
	return &NullableProjectsList{value: val, isSet: true}
}

func (v NullableProjectsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


