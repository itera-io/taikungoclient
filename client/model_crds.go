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

// checks if the Crds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Crds{}

// Crds struct for Crds
type Crds struct {
	Data []CrdListDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _Crds Crds

// NewCrds instantiates a new Crds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrds(data []CrdListDto, totalCount int32) *Crds {
	this := Crds{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewCrdsWithDefaults instantiates a new Crds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrdsWithDefaults() *Crds {
	this := Crds{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []CrdListDto will be returned
func (o *Crds) GetData() []CrdListDto {
	if o == nil {
		var ret []CrdListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Crds) GetDataOk() ([]CrdListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Crds) SetData(v []CrdListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *Crds) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *Crds) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *Crds) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o Crds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Crds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *Crds) UnmarshalJSON(data []byte) (err error) {
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

	varCrds := _Crds{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCrds)

	if err != nil {
		return err
	}

	*o = Crds(varCrds)

	return err
}

type NullableCrds struct {
	value *Crds
	isSet bool
}

func (v NullableCrds) Get() *Crds {
	return v.value
}

func (v *NullableCrds) Set(val *Crds) {
	v.value = val
	v.isSet = true
}

func (v NullableCrds) IsSet() bool {
	return v.isSet
}

func (v *NullableCrds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrds(val *Crds) *NullableCrds {
	return &NullableCrds{value: val, isSet: true}
}

func (v NullableCrds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


