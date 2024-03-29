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

// checks if the EditCatalogAppParamCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditCatalogAppParamCommand{}

// EditCatalogAppParamCommand struct for EditCatalogAppParamCommand
type EditCatalogAppParamCommand struct {
	CatalogAppId *int32 `json:"catalogAppId,omitempty"`
	Parameters []CatalogAppParamsDto `json:"parameters,omitempty"`
}

// NewEditCatalogAppParamCommand instantiates a new EditCatalogAppParamCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditCatalogAppParamCommand() *EditCatalogAppParamCommand {
	this := EditCatalogAppParamCommand{}
	return &this
}

// NewEditCatalogAppParamCommandWithDefaults instantiates a new EditCatalogAppParamCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditCatalogAppParamCommandWithDefaults() *EditCatalogAppParamCommand {
	this := EditCatalogAppParamCommand{}
	return &this
}

// GetCatalogAppId returns the CatalogAppId field value if set, zero value otherwise.
func (o *EditCatalogAppParamCommand) GetCatalogAppId() int32 {
	if o == nil || IsNil(o.CatalogAppId) {
		var ret int32
		return ret
	}
	return *o.CatalogAppId
}

// GetCatalogAppIdOk returns a tuple with the CatalogAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditCatalogAppParamCommand) GetCatalogAppIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogAppId) {
		return nil, false
	}
	return o.CatalogAppId, true
}

// HasCatalogAppId returns a boolean if a field has been set.
func (o *EditCatalogAppParamCommand) HasCatalogAppId() bool {
	if o != nil && !IsNil(o.CatalogAppId) {
		return true
	}

	return false
}

// SetCatalogAppId gets a reference to the given int32 and assigns it to the CatalogAppId field.
func (o *EditCatalogAppParamCommand) SetCatalogAppId(v int32) {
	o.CatalogAppId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditCatalogAppParamCommand) GetParameters() []CatalogAppParamsDto {
	if o == nil {
		var ret []CatalogAppParamsDto
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditCatalogAppParamCommand) GetParametersOk() ([]CatalogAppParamsDto, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *EditCatalogAppParamCommand) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []CatalogAppParamsDto and assigns it to the Parameters field.
func (o *EditCatalogAppParamCommand) SetParameters(v []CatalogAppParamsDto) {
	o.Parameters = v
}

func (o EditCatalogAppParamCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditCatalogAppParamCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogAppId) {
		toSerialize["catalogAppId"] = o.CatalogAppId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableEditCatalogAppParamCommand struct {
	value *EditCatalogAppParamCommand
	isSet bool
}

func (v NullableEditCatalogAppParamCommand) Get() *EditCatalogAppParamCommand {
	return v.value
}

func (v *NullableEditCatalogAppParamCommand) Set(val *EditCatalogAppParamCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditCatalogAppParamCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditCatalogAppParamCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditCatalogAppParamCommand(val *EditCatalogAppParamCommand) *NullableEditCatalogAppParamCommand {
	return &NullableEditCatalogAppParamCommand{value: val, isSet: true}
}

func (v NullableEditCatalogAppParamCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditCatalogAppParamCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


