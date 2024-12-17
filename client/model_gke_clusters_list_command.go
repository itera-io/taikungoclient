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

// checks if the GkeClustersListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GkeClustersListCommand{}

// GkeClustersListCommand struct for GkeClustersListCommand
type GkeClustersListCommand struct {
	CloudId *int32 `json:"cloudId,omitempty"`
}

// NewGkeClustersListCommand instantiates a new GkeClustersListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGkeClustersListCommand() *GkeClustersListCommand {
	this := GkeClustersListCommand{}
	return &this
}

// NewGkeClustersListCommandWithDefaults instantiates a new GkeClustersListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGkeClustersListCommandWithDefaults() *GkeClustersListCommand {
	this := GkeClustersListCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *GkeClustersListCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GkeClustersListCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *GkeClustersListCommand) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *GkeClustersListCommand) SetCloudId(v int32) {
	o.CloudId = &v
}

func (o GkeClustersListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GkeClustersListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	return toSerialize, nil
}

type NullableGkeClustersListCommand struct {
	value *GkeClustersListCommand
	isSet bool
}

func (v NullableGkeClustersListCommand) Get() *GkeClustersListCommand {
	return v.value
}

func (v *NullableGkeClustersListCommand) Set(val *GkeClustersListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableGkeClustersListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableGkeClustersListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGkeClustersListCommand(val *GkeClustersListCommand) *NullableGkeClustersListCommand {
	return &NullableGkeClustersListCommand{value: val, isSet: true}
}

func (v NullableGkeClustersListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGkeClustersListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


