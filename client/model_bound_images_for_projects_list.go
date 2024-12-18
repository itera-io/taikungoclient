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

// checks if the BoundImagesForProjectsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundImagesForProjectsList{}

// BoundImagesForProjectsList struct for BoundImagesForProjectsList
type BoundImagesForProjectsList struct {
	Data []BoundImagesForProjectsListDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _BoundImagesForProjectsList BoundImagesForProjectsList

// NewBoundImagesForProjectsList instantiates a new BoundImagesForProjectsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundImagesForProjectsList(data []BoundImagesForProjectsListDto, totalCount int32) *BoundImagesForProjectsList {
	this := BoundImagesForProjectsList{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewBoundImagesForProjectsListWithDefaults instantiates a new BoundImagesForProjectsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundImagesForProjectsListWithDefaults() *BoundImagesForProjectsList {
	this := BoundImagesForProjectsList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []BoundImagesForProjectsListDto will be returned
func (o *BoundImagesForProjectsList) GetData() []BoundImagesForProjectsListDto {
	if o == nil {
		var ret []BoundImagesForProjectsListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundImagesForProjectsList) GetDataOk() ([]BoundImagesForProjectsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BoundImagesForProjectsList) SetData(v []BoundImagesForProjectsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *BoundImagesForProjectsList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *BoundImagesForProjectsList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *BoundImagesForProjectsList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o BoundImagesForProjectsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoundImagesForProjectsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *BoundImagesForProjectsList) UnmarshalJSON(data []byte) (err error) {
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

	varBoundImagesForProjectsList := _BoundImagesForProjectsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBoundImagesForProjectsList)

	if err != nil {
		return err
	}

	*o = BoundImagesForProjectsList(varBoundImagesForProjectsList)

	return err
}

type NullableBoundImagesForProjectsList struct {
	value *BoundImagesForProjectsList
	isSet bool
}

func (v NullableBoundImagesForProjectsList) Get() *BoundImagesForProjectsList {
	return v.value
}

func (v *NullableBoundImagesForProjectsList) Set(val *BoundImagesForProjectsList) {
	v.value = val
	v.isSet = true
}

func (v NullableBoundImagesForProjectsList) IsSet() bool {
	return v.isSet
}

func (v *NullableBoundImagesForProjectsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoundImagesForProjectsList(val *BoundImagesForProjectsList) *NullableBoundImagesForProjectsList {
	return &NullableBoundImagesForProjectsList{value: val, isSet: true}
}

func (v NullableBoundImagesForProjectsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoundImagesForProjectsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


