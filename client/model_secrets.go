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
)

// checks if the Secrets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Secrets{}

// Secrets struct for Secrets
type Secrets struct {
	Data       []SecretDto `json:"data,omitempty"`
	TotalCount *int32      `json:"totalCount,omitempty"`
}

// NewSecrets instantiates a new Secrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecrets() *Secrets {
	this := Secrets{}
	return &this
}

// NewSecretsWithDefaults instantiates a new Secrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretsWithDefaults() *Secrets {
	this := Secrets{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Secrets) GetData() []SecretDto {
	if o == nil {
		var ret []SecretDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Secrets) GetDataOk() ([]SecretDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Secrets) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecretDto and assigns it to the Data field.
func (o *Secrets) SetData(v []SecretDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Secrets) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secrets) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Secrets) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *Secrets) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o Secrets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
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
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
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
