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

// checks if the NodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeDto{}

// NodeDto struct for NodeDto
type NodeDto struct {
	MetadataName interface{} `json:"metadataName"`
	KubeletReady interface{} `json:"kubeletReady"`
	KubeletSufficient interface{} `json:"kubeletSufficient"`
	KubeletDiskPressure interface{} `json:"kubeletDiskPressure"`
	KubeletMemory interface{} `json:"kubeletMemory"`
}

type _NodeDto NodeDto

// NewNodeDto instantiates a new NodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeDto(metadataName interface{}, kubeletReady interface{}, kubeletSufficient interface{}, kubeletDiskPressure interface{}, kubeletMemory interface{}) *NodeDto {
	this := NodeDto{}
	this.MetadataName = metadataName
	this.KubeletReady = kubeletReady
	this.KubeletSufficient = kubeletSufficient
	this.KubeletDiskPressure = kubeletDiskPressure
	this.KubeletMemory = kubeletMemory
	return &this
}

// NewNodeDtoWithDefaults instantiates a new NodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDtoWithDefaults() *NodeDto {
	this := NodeDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NodeDto) GetMetadataName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetMetadataNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MetadataName) {
		return nil, false
	}
	return &o.MetadataName, true
}

// SetMetadataName sets field value
func (o *NodeDto) SetMetadataName(v interface{}) {
	o.MetadataName = v
}

// GetKubeletReady returns the KubeletReady field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NodeDto) GetKubeletReady() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.KubeletReady
}

// GetKubeletReadyOk returns a tuple with the KubeletReady field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletReadyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletReady) {
		return nil, false
	}
	return &o.KubeletReady, true
}

// SetKubeletReady sets field value
func (o *NodeDto) SetKubeletReady(v interface{}) {
	o.KubeletReady = v
}

// GetKubeletSufficient returns the KubeletSufficient field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NodeDto) GetKubeletSufficient() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.KubeletSufficient
}

// GetKubeletSufficientOk returns a tuple with the KubeletSufficient field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletSufficientOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletSufficient) {
		return nil, false
	}
	return &o.KubeletSufficient, true
}

// SetKubeletSufficient sets field value
func (o *NodeDto) SetKubeletSufficient(v interface{}) {
	o.KubeletSufficient = v
}

// GetKubeletDiskPressure returns the KubeletDiskPressure field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NodeDto) GetKubeletDiskPressure() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.KubeletDiskPressure
}

// GetKubeletDiskPressureOk returns a tuple with the KubeletDiskPressure field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletDiskPressureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletDiskPressure) {
		return nil, false
	}
	return &o.KubeletDiskPressure, true
}

// SetKubeletDiskPressure sets field value
func (o *NodeDto) SetKubeletDiskPressure(v interface{}) {
	o.KubeletDiskPressure = v
}

// GetKubeletMemory returns the KubeletMemory field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NodeDto) GetKubeletMemory() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.KubeletMemory
}

// GetKubeletMemoryOk returns a tuple with the KubeletMemory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeDto) GetKubeletMemoryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KubeletMemory) {
		return nil, false
	}
	return &o.KubeletMemory, true
}

// SetKubeletMemory sets field value
func (o *NodeDto) SetKubeletMemory(v interface{}) {
	o.KubeletMemory = v
}

func (o NodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName != nil {
		toSerialize["metadataName"] = o.MetadataName
	}
	if o.KubeletReady != nil {
		toSerialize["kubeletReady"] = o.KubeletReady
	}
	if o.KubeletSufficient != nil {
		toSerialize["kubeletSufficient"] = o.KubeletSufficient
	}
	if o.KubeletDiskPressure != nil {
		toSerialize["kubeletDiskPressure"] = o.KubeletDiskPressure
	}
	if o.KubeletMemory != nil {
		toSerialize["kubeletMemory"] = o.KubeletMemory
	}
	return toSerialize, nil
}

func (o *NodeDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataName",
		"kubeletReady",
		"kubeletSufficient",
		"kubeletDiskPressure",
		"kubeletMemory",
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

	varNodeDto := _NodeDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeDto)

	if err != nil {
		return err
	}

	*o = NodeDto(varNodeDto)

	return err
}

type NullableNodeDto struct {
	value *NodeDto
	isSet bool
}

func (v NullableNodeDto) Get() *NodeDto {
	return v.value
}

func (v *NullableNodeDto) Set(val *NodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeDto(val *NodeDto) *NullableNodeDto {
	return &NullableNodeDto{value: val, isSet: true}
}

func (v NullableNodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


