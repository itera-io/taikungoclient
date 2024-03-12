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

// checks if the UpdateHypervisorsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHypervisorsCommand{}

// UpdateHypervisorsCommand struct for UpdateHypervisorsCommand
type UpdateHypervisorsCommand struct {
	Id *int32 `json:"id,omitempty"`
	Hypervisors []string `json:"hypervisors,omitempty"`
}

// NewUpdateHypervisorsCommand instantiates a new UpdateHypervisorsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHypervisorsCommand() *UpdateHypervisorsCommand {
	this := UpdateHypervisorsCommand{}
	return &this
}

// NewUpdateHypervisorsCommandWithDefaults instantiates a new UpdateHypervisorsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHypervisorsCommandWithDefaults() *UpdateHypervisorsCommand {
	this := UpdateHypervisorsCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateHypervisorsCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHypervisorsCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateHypervisorsCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateHypervisorsCommand) SetId(v int32) {
	o.Id = &v
}

// GetHypervisors returns the Hypervisors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateHypervisorsCommand) GetHypervisors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hypervisors
}

// GetHypervisorsOk returns a tuple with the Hypervisors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateHypervisorsCommand) GetHypervisorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hypervisors) {
		return nil, false
	}
	return o.Hypervisors, true
}

// HasHypervisors returns a boolean if a field has been set.
func (o *UpdateHypervisorsCommand) HasHypervisors() bool {
	if o != nil && !IsNil(o.Hypervisors) {
		return true
	}

	return false
}

// SetHypervisors gets a reference to the given []string and assigns it to the Hypervisors field.
func (o *UpdateHypervisorsCommand) SetHypervisors(v []string) {
	o.Hypervisors = v
}

func (o UpdateHypervisorsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHypervisorsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Hypervisors != nil {
		toSerialize["hypervisors"] = o.Hypervisors
	}
	return toSerialize, nil
}

type NullableUpdateHypervisorsCommand struct {
	value *UpdateHypervisorsCommand
	isSet bool
}

func (v NullableUpdateHypervisorsCommand) Get() *UpdateHypervisorsCommand {
	return v.value
}

func (v *NullableUpdateHypervisorsCommand) Set(val *UpdateHypervisorsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHypervisorsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHypervisorsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHypervisorsCommand(val *UpdateHypervisorsCommand) *NullableUpdateHypervisorsCommand {
	return &NullableUpdateHypervisorsCommand{value: val, isSet: true}
}

func (v NullableUpdateHypervisorsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHypervisorsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


