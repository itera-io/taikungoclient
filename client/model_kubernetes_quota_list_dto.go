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

// checks if the KubernetesQuotaListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesQuotaListDto{}

// KubernetesQuotaListDto struct for KubernetesQuotaListDto
type KubernetesQuotaListDto struct {
	SumOfCpu float64 `json:"sumOfCpu"`
	SumOfRamInGb float64 `json:"sumOfRamInGb"`
	SumOfCpuUsage float64 `json:"sumOfCpuUsage"`
	SumOfRamUsage float64 `json:"sumOfRamUsage"`
	PodsCapacity int32 `json:"podsCapacity"`
	PodsTotalCount int32 `json:"podsTotalCount"`
}

type _KubernetesQuotaListDto KubernetesQuotaListDto

// NewKubernetesQuotaListDto instantiates a new KubernetesQuotaListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesQuotaListDto(sumOfCpu float64, sumOfRamInGb float64, sumOfCpuUsage float64, sumOfRamUsage float64, podsCapacity int32, podsTotalCount int32) *KubernetesQuotaListDto {
	this := KubernetesQuotaListDto{}
	this.SumOfCpu = sumOfCpu
	this.SumOfRamInGb = sumOfRamInGb
	this.SumOfCpuUsage = sumOfCpuUsage
	this.SumOfRamUsage = sumOfRamUsage
	this.PodsCapacity = podsCapacity
	this.PodsTotalCount = podsTotalCount
	return &this
}

// NewKubernetesQuotaListDtoWithDefaults instantiates a new KubernetesQuotaListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesQuotaListDtoWithDefaults() *KubernetesQuotaListDto {
	this := KubernetesQuotaListDto{}
	return &this
}

// GetSumOfCpu returns the SumOfCpu field value
func (o *KubernetesQuotaListDto) GetSumOfCpu() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumOfCpu
}

// GetSumOfCpuOk returns a tuple with the SumOfCpu field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfCpuOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumOfCpu, true
}

// SetSumOfCpu sets field value
func (o *KubernetesQuotaListDto) SetSumOfCpu(v float64) {
	o.SumOfCpu = v
}

// GetSumOfRamInGb returns the SumOfRamInGb field value
func (o *KubernetesQuotaListDto) GetSumOfRamInGb() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumOfRamInGb
}

// GetSumOfRamInGbOk returns a tuple with the SumOfRamInGb field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfRamInGbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumOfRamInGb, true
}

// SetSumOfRamInGb sets field value
func (o *KubernetesQuotaListDto) SetSumOfRamInGb(v float64) {
	o.SumOfRamInGb = v
}

// GetSumOfCpuUsage returns the SumOfCpuUsage field value
func (o *KubernetesQuotaListDto) GetSumOfCpuUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumOfCpuUsage
}

// GetSumOfCpuUsageOk returns a tuple with the SumOfCpuUsage field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfCpuUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumOfCpuUsage, true
}

// SetSumOfCpuUsage sets field value
func (o *KubernetesQuotaListDto) SetSumOfCpuUsage(v float64) {
	o.SumOfCpuUsage = v
}

// GetSumOfRamUsage returns the SumOfRamUsage field value
func (o *KubernetesQuotaListDto) GetSumOfRamUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SumOfRamUsage
}

// GetSumOfRamUsageOk returns a tuple with the SumOfRamUsage field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetSumOfRamUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumOfRamUsage, true
}

// SetSumOfRamUsage sets field value
func (o *KubernetesQuotaListDto) SetSumOfRamUsage(v float64) {
	o.SumOfRamUsage = v
}

// GetPodsCapacity returns the PodsCapacity field value
func (o *KubernetesQuotaListDto) GetPodsCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PodsCapacity
}

// GetPodsCapacityOk returns a tuple with the PodsCapacity field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetPodsCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodsCapacity, true
}

// SetPodsCapacity sets field value
func (o *KubernetesQuotaListDto) SetPodsCapacity(v int32) {
	o.PodsCapacity = v
}

// GetPodsTotalCount returns the PodsTotalCount field value
func (o *KubernetesQuotaListDto) GetPodsTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PodsTotalCount
}

// GetPodsTotalCountOk returns a tuple with the PodsTotalCount field value
// and a boolean to check if the value has been set.
func (o *KubernetesQuotaListDto) GetPodsTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodsTotalCount, true
}

// SetPodsTotalCount sets field value
func (o *KubernetesQuotaListDto) SetPodsTotalCount(v int32) {
	o.PodsTotalCount = v
}

func (o KubernetesQuotaListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesQuotaListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sumOfCpu"] = o.SumOfCpu
	toSerialize["sumOfRamInGb"] = o.SumOfRamInGb
	toSerialize["sumOfCpuUsage"] = o.SumOfCpuUsage
	toSerialize["sumOfRamUsage"] = o.SumOfRamUsage
	toSerialize["podsCapacity"] = o.PodsCapacity
	toSerialize["podsTotalCount"] = o.PodsTotalCount
	return toSerialize, nil
}

func (o *KubernetesQuotaListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sumOfCpu",
		"sumOfRamInGb",
		"sumOfCpuUsage",
		"sumOfRamUsage",
		"podsCapacity",
		"podsTotalCount",
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

	varKubernetesQuotaListDto := _KubernetesQuotaListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesQuotaListDto)

	if err != nil {
		return err
	}

	*o = KubernetesQuotaListDto(varKubernetesQuotaListDto)

	return err
}

type NullableKubernetesQuotaListDto struct {
	value *KubernetesQuotaListDto
	isSet bool
}

func (v NullableKubernetesQuotaListDto) Get() *KubernetesQuotaListDto {
	return v.value
}

func (v *NullableKubernetesQuotaListDto) Set(val *KubernetesQuotaListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesQuotaListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesQuotaListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesQuotaListDto(val *KubernetesQuotaListDto) *NullableKubernetesQuotaListDto {
	return &NullableKubernetesQuotaListDto{value: val, isSet: true}
}

func (v NullableKubernetesQuotaListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesQuotaListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


