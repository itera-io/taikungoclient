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

// checks if the AllFlavorsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllFlavorsList{}

// AllFlavorsList struct for AllFlavorsList
type AllFlavorsList struct {
	Data []FlavorsListDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
	CloudType NullableString `json:"cloudType"`
}

type _AllFlavorsList AllFlavorsList

// NewAllFlavorsList instantiates a new AllFlavorsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllFlavorsList(data []FlavorsListDto, totalCount int32, cloudType NullableString) *AllFlavorsList {
	this := AllFlavorsList{}
	this.Data = data
	this.TotalCount = totalCount
	this.CloudType = cloudType
	return &this
}

// NewAllFlavorsListWithDefaults instantiates a new AllFlavorsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllFlavorsListWithDefaults() *AllFlavorsList {
	this := AllFlavorsList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []FlavorsListDto will be returned
func (o *AllFlavorsList) GetData() []FlavorsListDto {
	if o == nil {
		var ret []FlavorsListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllFlavorsList) GetDataOk() ([]FlavorsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AllFlavorsList) SetData(v []FlavorsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *AllFlavorsList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AllFlavorsList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AllFlavorsList) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetCloudType returns the CloudType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AllFlavorsList) GetCloudType() string {
	if o == nil || o.CloudType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllFlavorsList) GetCloudTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// SetCloudType sets field value
func (o *AllFlavorsList) SetCloudType(v string) {
	o.CloudType.Set(&v)
}

func (o AllFlavorsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllFlavorsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["cloudType"] = o.CloudType.Get()
	return toSerialize, nil
}

func (o *AllFlavorsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"totalCount",
		"cloudType",
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

	varAllFlavorsList := _AllFlavorsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAllFlavorsList)

	if err != nil {
		return err
	}

	*o = AllFlavorsList(varAllFlavorsList)

	return err
}

type NullableAllFlavorsList struct {
	value *AllFlavorsList
	isSet bool
}

func (v NullableAllFlavorsList) Get() *AllFlavorsList {
	return v.value
}

func (v *NullableAllFlavorsList) Set(val *AllFlavorsList) {
	v.value = val
	v.isSet = true
}

func (v NullableAllFlavorsList) IsSet() bool {
	return v.isSet
}

func (v *NullableAllFlavorsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllFlavorsList(val *AllFlavorsList) *NullableAllFlavorsList {
	return &NullableAllFlavorsList{value: val, isSet: true}
}

func (v NullableAllFlavorsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllFlavorsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


