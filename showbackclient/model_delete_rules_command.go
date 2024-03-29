/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the DeleteRulesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRulesCommand{}

// DeleteRulesCommand struct for DeleteRulesCommand
type DeleteRulesCommand struct {
	Ids []int32 `json:"ids,omitempty"`
}

// NewDeleteRulesCommand instantiates a new DeleteRulesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRulesCommand() *DeleteRulesCommand {
	this := DeleteRulesCommand{}
	return &this
}

// NewDeleteRulesCommandWithDefaults instantiates a new DeleteRulesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRulesCommandWithDefaults() *DeleteRulesCommand {
	this := DeleteRulesCommand{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteRulesCommand) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteRulesCommand) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteRulesCommand) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *DeleteRulesCommand) SetIds(v []int32) {
	o.Ids = v
}

func (o DeleteRulesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRulesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableDeleteRulesCommand struct {
	value *DeleteRulesCommand
	isSet bool
}

func (v NullableDeleteRulesCommand) Get() *DeleteRulesCommand {
	return v.value
}

func (v *NullableDeleteRulesCommand) Set(val *DeleteRulesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRulesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRulesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRulesCommand(val *DeleteRulesCommand) *NullableDeleteRulesCommand {
	return &NullableDeleteRulesCommand{value: val, isSet: true}
}

func (v NullableDeleteRulesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRulesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


