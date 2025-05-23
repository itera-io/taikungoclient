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

// checks if the Secrets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Secrets{}

// Secrets struct for Secrets
type Secrets struct {
	Data []SecretDto `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _Secrets Secrets

// NewSecrets instantiates a new Secrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecrets(data []SecretDto, totalCount int32) *Secrets {
	this := Secrets{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewSecretsWithDefaults instantiates a new Secrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretsWithDefaults() *Secrets {
	this := Secrets{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []SecretDto will be returned
func (o *Secrets) GetData() []SecretDto {
	if o == nil {
		var ret []SecretDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Secrets) GetDataOk() ([]SecretDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Secrets) SetData(v []SecretDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *Secrets) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *Secrets) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *Secrets) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o Secrets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secrets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *Secrets) UnmarshalJSON(data []byte) (err error) {
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

	varSecrets := _Secrets{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecrets)

	if err != nil {
		return err
	}

	*o = Secrets(varSecrets)

	return err
}

type NullableSecrets struct {
	value *Secrets
	isSet bool
}

func (v NullableSecrets) Get() *Secrets {
	return v.value
}

func (v *NullableSecrets) Set(val *Secrets) {
	v.value = val
	v.isSet = true
}

func (v NullableSecrets) IsSet() bool {
	return v.isSet
}

func (v *NullableSecrets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecrets(val *Secrets) *NullableSecrets {
	return &NullableSecrets{value: val, isSet: true}
}

func (v NullableSecrets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecrets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


