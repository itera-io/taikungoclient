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

// checks if the KubernetesCronJobsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCronJobsList{}

// KubernetesCronJobsList struct for KubernetesCronJobsList
type KubernetesCronJobsList struct {
	Data []KubernetesCronJobDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _KubernetesCronJobsList KubernetesCronJobsList

// NewKubernetesCronJobsList instantiates a new KubernetesCronJobsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCronJobsList(data []KubernetesCronJobDto, totalCount int32) *KubernetesCronJobsList {
	this := KubernetesCronJobsList{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewKubernetesCronJobsListWithDefaults instantiates a new KubernetesCronJobsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCronJobsListWithDefaults() *KubernetesCronJobsList {
	this := KubernetesCronJobsList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []KubernetesCronJobDto will be returned
func (o *KubernetesCronJobsList) GetData() []KubernetesCronJobDto {
	if o == nil {
		var ret []KubernetesCronJobDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCronJobsList) GetDataOk() ([]KubernetesCronJobDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *KubernetesCronJobsList) SetData(v []KubernetesCronJobDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *KubernetesCronJobsList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *KubernetesCronJobsList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *KubernetesCronJobsList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o KubernetesCronJobsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCronJobsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *KubernetesCronJobsList) UnmarshalJSON(data []byte) (err error) {
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

	varKubernetesCronJobsList := _KubernetesCronJobsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesCronJobsList)

	if err != nil {
		return err
	}

	*o = KubernetesCronJobsList(varKubernetesCronJobsList)

	return err
}

type NullableKubernetesCronJobsList struct {
	value *KubernetesCronJobsList
	isSet bool
}

func (v NullableKubernetesCronJobsList) Get() *KubernetesCronJobsList {
	return v.value
}

func (v *NullableKubernetesCronJobsList) Set(val *KubernetesCronJobsList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCronJobsList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCronJobsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCronJobsList(val *KubernetesCronJobsList) *NullableKubernetesCronJobsList {
	return &NullableKubernetesCronJobsList{value: val, isSet: true}
}

func (v NullableKubernetesCronJobsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCronJobsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


