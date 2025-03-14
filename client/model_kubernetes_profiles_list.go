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

// checks if the KubernetesProfilesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesProfilesList{}

// KubernetesProfilesList struct for KubernetesProfilesList
type KubernetesProfilesList struct {
	Data []KubernetesProfilesListDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _KubernetesProfilesList KubernetesProfilesList

// NewKubernetesProfilesList instantiates a new KubernetesProfilesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesProfilesList(data []KubernetesProfilesListDto, totalCount int32) *KubernetesProfilesList {
	this := KubernetesProfilesList{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewKubernetesProfilesListWithDefaults instantiates a new KubernetesProfilesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesProfilesListWithDefaults() *KubernetesProfilesList {
	this := KubernetesProfilesList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []KubernetesProfilesListDto will be returned
func (o *KubernetesProfilesList) GetData() []KubernetesProfilesListDto {
	if o == nil {
		var ret []KubernetesProfilesListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesList) GetDataOk() ([]KubernetesProfilesListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *KubernetesProfilesList) SetData(v []KubernetesProfilesListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *KubernetesProfilesList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *KubernetesProfilesList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o KubernetesProfilesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesProfilesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *KubernetesProfilesList) UnmarshalJSON(data []byte) (err error) {
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

	varKubernetesProfilesList := _KubernetesProfilesList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesProfilesList)

	if err != nil {
		return err
	}

	*o = KubernetesProfilesList(varKubernetesProfilesList)

	return err
}

type NullableKubernetesProfilesList struct {
	value *KubernetesProfilesList
	isSet bool
}

func (v NullableKubernetesProfilesList) Get() *KubernetesProfilesList {
	return v.value
}

func (v *NullableKubernetesProfilesList) Set(val *KubernetesProfilesList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesProfilesList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesProfilesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesProfilesList(val *KubernetesProfilesList) *NullableKubernetesProfilesList {
	return &NullableKubernetesProfilesList{value: val, isSet: true}
}

func (v NullableKubernetesProfilesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesProfilesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


