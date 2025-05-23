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

// checks if the PrometheusDashboardDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusDashboardDto{}

// PrometheusDashboardDto struct for PrometheusDashboardDto
type PrometheusDashboardDto struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	ExpressionDecoded NullableString `json:"expressionDecoded"`
	ExpressionEncoded NullableString `json:"expressionEncoded"`
	Description NullableString `json:"description"`
	IsReadonly bool `json:"isReadonly"`
}

type _PrometheusDashboardDto PrometheusDashboardDto

// NewPrometheusDashboardDto instantiates a new PrometheusDashboardDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusDashboardDto(id int32, name NullableString, expressionDecoded NullableString, expressionEncoded NullableString, description NullableString, isReadonly bool) *PrometheusDashboardDto {
	this := PrometheusDashboardDto{}
	this.Id = id
	this.Name = name
	this.ExpressionDecoded = expressionDecoded
	this.ExpressionEncoded = expressionEncoded
	this.Description = description
	this.IsReadonly = isReadonly
	return &this
}

// NewPrometheusDashboardDtoWithDefaults instantiates a new PrometheusDashboardDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusDashboardDtoWithDefaults() *PrometheusDashboardDto {
	this := PrometheusDashboardDto{}
	return &this
}

// GetId returns the Id field value
func (o *PrometheusDashboardDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrometheusDashboardDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrometheusDashboardDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PrometheusDashboardDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetExpressionDecoded returns the ExpressionDecoded field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrometheusDashboardDto) GetExpressionDecoded() string {
	if o == nil || o.ExpressionDecoded.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpressionDecoded.Get()
}

// GetExpressionDecodedOk returns a tuple with the ExpressionDecoded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardDto) GetExpressionDecodedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpressionDecoded.Get(), o.ExpressionDecoded.IsSet()
}

// SetExpressionDecoded sets field value
func (o *PrometheusDashboardDto) SetExpressionDecoded(v string) {
	o.ExpressionDecoded.Set(&v)
}

// GetExpressionEncoded returns the ExpressionEncoded field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrometheusDashboardDto) GetExpressionEncoded() string {
	if o == nil || o.ExpressionEncoded.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpressionEncoded.Get()
}

// GetExpressionEncodedOk returns a tuple with the ExpressionEncoded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardDto) GetExpressionEncodedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpressionEncoded.Get(), o.ExpressionEncoded.IsSet()
}

// SetExpressionEncoded sets field value
func (o *PrometheusDashboardDto) SetExpressionEncoded(v string) {
	o.ExpressionEncoded.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrometheusDashboardDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *PrometheusDashboardDto) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetIsReadonly returns the IsReadonly field value
func (o *PrometheusDashboardDto) GetIsReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReadonly
}

// GetIsReadonlyOk returns a tuple with the IsReadonly field value
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardDto) GetIsReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReadonly, true
}

// SetIsReadonly sets field value
func (o *PrometheusDashboardDto) SetIsReadonly(v bool) {
	o.IsReadonly = v
}

func (o PrometheusDashboardDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusDashboardDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["expressionDecoded"] = o.ExpressionDecoded.Get()
	toSerialize["expressionEncoded"] = o.ExpressionEncoded.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["isReadonly"] = o.IsReadonly
	return toSerialize, nil
}

func (o *PrometheusDashboardDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"expressionDecoded",
		"expressionEncoded",
		"description",
		"isReadonly",
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

	varPrometheusDashboardDto := _PrometheusDashboardDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrometheusDashboardDto)

	if err != nil {
		return err
	}

	*o = PrometheusDashboardDto(varPrometheusDashboardDto)

	return err
}

type NullablePrometheusDashboardDto struct {
	value *PrometheusDashboardDto
	isSet bool
}

func (v NullablePrometheusDashboardDto) Get() *PrometheusDashboardDto {
	return v.value
}

func (v *NullablePrometheusDashboardDto) Set(val *PrometheusDashboardDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusDashboardDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusDashboardDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusDashboardDto(val *PrometheusDashboardDto) *NullablePrometheusDashboardDto {
	return &NullablePrometheusDashboardDto{value: val, isSet: true}
}

func (v NullablePrometheusDashboardDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusDashboardDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


