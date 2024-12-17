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

// checks if the ProjectCanAddVClusterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCanAddVClusterDto{}

// ProjectCanAddVClusterDto struct for ProjectCanAddVClusterDto
type ProjectCanAddVClusterDto struct {
	AddVCluster ButtonStatusDto `json:"addVCluster"`
}

type _ProjectCanAddVClusterDto ProjectCanAddVClusterDto

// NewProjectCanAddVClusterDto instantiates a new ProjectCanAddVClusterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCanAddVClusterDto(addVCluster ButtonStatusDto) *ProjectCanAddVClusterDto {
	this := ProjectCanAddVClusterDto{}
	this.AddVCluster = addVCluster
	return &this
}

// NewProjectCanAddVClusterDtoWithDefaults instantiates a new ProjectCanAddVClusterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCanAddVClusterDtoWithDefaults() *ProjectCanAddVClusterDto {
	this := ProjectCanAddVClusterDto{}
	return &this
}

// GetAddVCluster returns the AddVCluster field value
func (o *ProjectCanAddVClusterDto) GetAddVCluster() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.AddVCluster
}

// GetAddVClusterOk returns a tuple with the AddVCluster field value
// and a boolean to check if the value has been set.
func (o *ProjectCanAddVClusterDto) GetAddVClusterOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddVCluster, true
}

// SetAddVCluster sets field value
func (o *ProjectCanAddVClusterDto) SetAddVCluster(v ButtonStatusDto) {
	o.AddVCluster = v
}

func (o ProjectCanAddVClusterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCanAddVClusterDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addVCluster"] = o.AddVCluster
	return toSerialize, nil
}

func (o *ProjectCanAddVClusterDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addVCluster",
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

	varProjectCanAddVClusterDto := _ProjectCanAddVClusterDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectCanAddVClusterDto)

	if err != nil {
		return err
	}

	*o = ProjectCanAddVClusterDto(varProjectCanAddVClusterDto)

	return err
}

type NullableProjectCanAddVClusterDto struct {
	value *ProjectCanAddVClusterDto
	isSet bool
}

func (v NullableProjectCanAddVClusterDto) Get() *ProjectCanAddVClusterDto {
	return v.value
}

func (v *NullableProjectCanAddVClusterDto) Set(val *ProjectCanAddVClusterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCanAddVClusterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCanAddVClusterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCanAddVClusterDto(val *ProjectCanAddVClusterDto) *NullableProjectCanAddVClusterDto {
	return &NullableProjectCanAddVClusterDto{value: val, isSet: true}
}

func (v NullableProjectCanAddVClusterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCanAddVClusterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


