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

// checks if the StsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsDto{}

// StsDto struct for StsDto
type StsDto struct {
	MetadataName string `json:"metadataName"`
	Status string `json:"status"`
	Namespace string `json:"namespace"`
	Age NullableString `json:"age"`
}

type _StsDto StsDto

// NewStsDto instantiates a new StsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsDto(metadataName string, status string, namespace string, age NullableString) *StsDto {
	this := StsDto{}
	this.MetadataName = metadataName
	this.Status = status
	this.Namespace = namespace
	this.Age = age
	return &this
}

// NewStsDtoWithDefaults instantiates a new StsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsDtoWithDefaults() *StsDto {
	this := StsDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value
func (o *StsDto) GetMetadataName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value
// and a boolean to check if the value has been set.
func (o *StsDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataName, true
}

// SetMetadataName sets field value
func (o *StsDto) SetMetadataName(v string) {
	o.MetadataName = v
}

// GetStatus returns the Status field value
func (o *StsDto) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StsDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StsDto) SetStatus(v string) {
	o.Status = v
}

// GetNamespace returns the Namespace field value
func (o *StsDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *StsDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *StsDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetAge returns the Age field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StsDto) GetAge() string {
	if o == nil || o.Age.Get() == nil {
		var ret string
		return ret
	}

	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StsDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// SetAge sets field value
func (o *StsDto) SetAge(v string) {
	o.Age.Set(&v)
}

func (o StsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataName"] = o.MetadataName
	toSerialize["status"] = o.Status
	toSerialize["namespace"] = o.Namespace
	toSerialize["age"] = o.Age.Get()
	return toSerialize, nil
}

func (o *StsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataName",
		"status",
		"namespace",
		"age",
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

	varStsDto := _StsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStsDto)

	if err != nil {
		return err
	}

	*o = StsDto(varStsDto)

	return err
}

type NullableStsDto struct {
	value *StsDto
	isSet bool
}

func (v NullableStsDto) Get() *StsDto {
	return v.value
}

func (v *NullableStsDto) Set(val *StsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsDto(val *StsDto) *NullableStsDto {
	return &NullableStsDto{value: val, isSet: true}
}

func (v NullableStsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


