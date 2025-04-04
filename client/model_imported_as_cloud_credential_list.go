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

// checks if the ImportedAsCloudCredentialList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedAsCloudCredentialList{}

// ImportedAsCloudCredentialList struct for ImportedAsCloudCredentialList
type ImportedAsCloudCredentialList struct {
	Visibility ImportedAsCloudCredentialVisibility `json:"visibility"`
	Data ImportedClusterDetailsDto `json:"data"`
}

type _ImportedAsCloudCredentialList ImportedAsCloudCredentialList

// NewImportedAsCloudCredentialList instantiates a new ImportedAsCloudCredentialList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedAsCloudCredentialList(visibility ImportedAsCloudCredentialVisibility, data ImportedClusterDetailsDto) *ImportedAsCloudCredentialList {
	this := ImportedAsCloudCredentialList{}
	this.Visibility = visibility
	this.Data = data
	return &this
}

// NewImportedAsCloudCredentialListWithDefaults instantiates a new ImportedAsCloudCredentialList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedAsCloudCredentialListWithDefaults() *ImportedAsCloudCredentialList {
	this := ImportedAsCloudCredentialList{}
	return &this
}

// GetVisibility returns the Visibility field value
func (o *ImportedAsCloudCredentialList) GetVisibility() ImportedAsCloudCredentialVisibility {
	if o == nil {
		var ret ImportedAsCloudCredentialVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ImportedAsCloudCredentialList) GetVisibilityOk() (*ImportedAsCloudCredentialVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ImportedAsCloudCredentialList) SetVisibility(v ImportedAsCloudCredentialVisibility) {
	o.Visibility = v
}

// GetData returns the Data field value
func (o *ImportedAsCloudCredentialList) GetData() ImportedClusterDetailsDto {
	if o == nil {
		var ret ImportedClusterDetailsDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImportedAsCloudCredentialList) GetDataOk() (*ImportedClusterDetailsDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImportedAsCloudCredentialList) SetData(v ImportedClusterDetailsDto) {
	o.Data = v
}

func (o ImportedAsCloudCredentialList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedAsCloudCredentialList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visibility"] = o.Visibility
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ImportedAsCloudCredentialList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"visibility",
		"data",
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

	varImportedAsCloudCredentialList := _ImportedAsCloudCredentialList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportedAsCloudCredentialList)

	if err != nil {
		return err
	}

	*o = ImportedAsCloudCredentialList(varImportedAsCloudCredentialList)

	return err
}

type NullableImportedAsCloudCredentialList struct {
	value *ImportedAsCloudCredentialList
	isSet bool
}

func (v NullableImportedAsCloudCredentialList) Get() *ImportedAsCloudCredentialList {
	return v.value
}

func (v *NullableImportedAsCloudCredentialList) Set(val *ImportedAsCloudCredentialList) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedAsCloudCredentialList) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedAsCloudCredentialList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedAsCloudCredentialList(val *ImportedAsCloudCredentialList) *NullableImportedAsCloudCredentialList {
	return &NullableImportedAsCloudCredentialList{value: val, isSet: true}
}

func (v NullableImportedAsCloudCredentialList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedAsCloudCredentialList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


