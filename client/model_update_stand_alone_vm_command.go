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

// checks if the UpdateStandAloneVmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStandAloneVmCommand{}

// UpdateStandAloneVmCommand struct for UpdateStandAloneVmCommand
type UpdateStandAloneVmCommand struct {
	StandaloneVms []UpdateStandaloneVmDto `json:"standaloneVms,omitempty"`
}

// NewUpdateStandAloneVmCommand instantiates a new UpdateStandAloneVmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStandAloneVmCommand() *UpdateStandAloneVmCommand {
	this := UpdateStandAloneVmCommand{}
	return &this
}

// NewUpdateStandAloneVmCommandWithDefaults instantiates a new UpdateStandAloneVmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStandAloneVmCommandWithDefaults() *UpdateStandAloneVmCommand {
	this := UpdateStandAloneVmCommand{}
	return &this
}

// GetStandaloneVms returns the StandaloneVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStandAloneVmCommand) GetStandaloneVms() []UpdateStandaloneVmDto {
	if o == nil {
		var ret []UpdateStandaloneVmDto
		return ret
	}
	return o.StandaloneVms
}

// GetStandaloneVmsOk returns a tuple with the StandaloneVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStandAloneVmCommand) GetStandaloneVmsOk() ([]UpdateStandaloneVmDto, bool) {
	if o == nil || IsNil(o.StandaloneVms) {
		return nil, false
	}
	return o.StandaloneVms, true
}

// HasStandaloneVms returns a boolean if a field has been set.
func (o *UpdateStandAloneVmCommand) HasStandaloneVms() bool {
	if o != nil && !IsNil(o.StandaloneVms) {
		return true
	}

	return false
}

// SetStandaloneVms gets a reference to the given []UpdateStandaloneVmDto and assigns it to the StandaloneVms field.
func (o *UpdateStandAloneVmCommand) SetStandaloneVms(v []UpdateStandaloneVmDto) {
	o.StandaloneVms = v
}

func (o UpdateStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStandAloneVmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StandaloneVms != nil {
		toSerialize["standaloneVms"] = o.StandaloneVms
	}
	return toSerialize, nil
}

type NullableUpdateStandAloneVmCommand struct {
	value *UpdateStandAloneVmCommand
	isSet bool
}

func (v NullableUpdateStandAloneVmCommand) Get() *UpdateStandAloneVmCommand {
	return v.value
}

func (v *NullableUpdateStandAloneVmCommand) Set(val *UpdateStandAloneVmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStandAloneVmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStandAloneVmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStandAloneVmCommand(val *UpdateStandAloneVmCommand) *NullableUpdateStandAloneVmCommand {
	return &NullableUpdateStandAloneVmCommand{value: val, isSet: true}
}

func (v NullableUpdateStandAloneVmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStandAloneVmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


