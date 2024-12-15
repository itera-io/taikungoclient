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

// checks if the GenericKubernetesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericKubernetesList{}

// GenericKubernetesList struct for GenericKubernetesList
type GenericKubernetesList struct {
	Data []GenericKubernetesListDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _GenericKubernetesList GenericKubernetesList

// NewGenericKubernetesList instantiates a new GenericKubernetesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericKubernetesList(data []GenericKubernetesListDto, totalCount int32) *GenericKubernetesList {
	this := GenericKubernetesList{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewGenericKubernetesListWithDefaults instantiates a new GenericKubernetesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericKubernetesListWithDefaults() *GenericKubernetesList {
	this := GenericKubernetesList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []GenericKubernetesListDto will be returned
func (o *GenericKubernetesList) GetData() []GenericKubernetesListDto {
	if o == nil {
		var ret []GenericKubernetesListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericKubernetesList) GetDataOk() ([]GenericKubernetesListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GenericKubernetesList) SetData(v []GenericKubernetesListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *GenericKubernetesList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *GenericKubernetesList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *GenericKubernetesList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o GenericKubernetesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericKubernetesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *GenericKubernetesList) UnmarshalJSON(data []byte) (err error) {
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

	varGenericKubernetesList := _GenericKubernetesList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenericKubernetesList)

	if err != nil {
		return err
	}

	*o = GenericKubernetesList(varGenericKubernetesList)

	return err
}

type NullableGenericKubernetesList struct {
	value *GenericKubernetesList
	isSet bool
}

func (v NullableGenericKubernetesList) Get() *GenericKubernetesList {
	return v.value
}

func (v *NullableGenericKubernetesList) Set(val *GenericKubernetesList) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericKubernetesList) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericKubernetesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericKubernetesList(val *GenericKubernetesList) *NullableGenericKubernetesList {
	return &NullableGenericKubernetesList{value: val, isSet: true}
}

func (v NullableGenericKubernetesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericKubernetesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

