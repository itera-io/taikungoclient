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

// checks if the ImportedAsReadOnlyList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedAsReadOnlyList{}

// ImportedAsReadOnlyList struct for ImportedAsReadOnlyList
type ImportedAsReadOnlyList struct {
	Visibility ImportedAsReadOnlyVisibility `json:"visibility"`
	Data ImportedClusterDetailsDto `json:"data"`
}

type _ImportedAsReadOnlyList ImportedAsReadOnlyList

// NewImportedAsReadOnlyList instantiates a new ImportedAsReadOnlyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedAsReadOnlyList(visibility ImportedAsReadOnlyVisibility, data ImportedClusterDetailsDto) *ImportedAsReadOnlyList {
	this := ImportedAsReadOnlyList{}
	this.Visibility = visibility
	this.Data = data
	return &this
}

// NewImportedAsReadOnlyListWithDefaults instantiates a new ImportedAsReadOnlyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedAsReadOnlyListWithDefaults() *ImportedAsReadOnlyList {
	this := ImportedAsReadOnlyList{}
	return &this
}

// GetVisibility returns the Visibility field value
func (o *ImportedAsReadOnlyList) GetVisibility() ImportedAsReadOnlyVisibility {
	if o == nil {
		var ret ImportedAsReadOnlyVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ImportedAsReadOnlyList) GetVisibilityOk() (*ImportedAsReadOnlyVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ImportedAsReadOnlyList) SetVisibility(v ImportedAsReadOnlyVisibility) {
	o.Visibility = v
}

// GetData returns the Data field value
func (o *ImportedAsReadOnlyList) GetData() ImportedClusterDetailsDto {
	if o == nil {
		var ret ImportedClusterDetailsDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImportedAsReadOnlyList) GetDataOk() (*ImportedClusterDetailsDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ImportedAsReadOnlyList) SetData(v ImportedClusterDetailsDto) {
	o.Data = v
}

func (o ImportedAsReadOnlyList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedAsReadOnlyList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visibility"] = o.Visibility
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ImportedAsReadOnlyList) UnmarshalJSON(data []byte) (err error) {
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

	varImportedAsReadOnlyList := _ImportedAsReadOnlyList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportedAsReadOnlyList)

	if err != nil {
		return err
	}

	*o = ImportedAsReadOnlyList(varImportedAsReadOnlyList)

	return err
}

type NullableImportedAsReadOnlyList struct {
	value *ImportedAsReadOnlyList
	isSet bool
}

func (v NullableImportedAsReadOnlyList) Get() *ImportedAsReadOnlyList {
	return v.value
}

func (v *NullableImportedAsReadOnlyList) Set(val *ImportedAsReadOnlyList) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedAsReadOnlyList) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedAsReadOnlyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedAsReadOnlyList(val *ImportedAsReadOnlyList) *NullableImportedAsReadOnlyList {
	return &NullableImportedAsReadOnlyList{value: val, isSet: true}
}

func (v NullableImportedAsReadOnlyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedAsReadOnlyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


