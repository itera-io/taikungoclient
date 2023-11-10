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

// checks if the AutoSyncManagementCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoSyncManagementCommand{}

// AutoSyncManagementCommand struct for AutoSyncManagementCommand
type AutoSyncManagementCommand struct {
	Id *int32 `json:"id,omitempty"`
	Mode NullableString `json:"mode,omitempty"`
}

// NewAutoSyncManagementCommand instantiates a new AutoSyncManagementCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoSyncManagementCommand() *AutoSyncManagementCommand {
	this := AutoSyncManagementCommand{}
	return &this
}

// NewAutoSyncManagementCommandWithDefaults instantiates a new AutoSyncManagementCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoSyncManagementCommandWithDefaults() *AutoSyncManagementCommand {
	this := AutoSyncManagementCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoSyncManagementCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoSyncManagementCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoSyncManagementCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoSyncManagementCommand) SetId(v int32) {
	o.Id = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoSyncManagementCommand) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoSyncManagementCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *AutoSyncManagementCommand) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *AutoSyncManagementCommand) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *AutoSyncManagementCommand) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *AutoSyncManagementCommand) UnsetMode() {
	o.Mode.Unset()
}

func (o AutoSyncManagementCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoSyncManagementCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableAutoSyncManagementCommand struct {
	value *AutoSyncManagementCommand
	isSet bool
}

func (v NullableAutoSyncManagementCommand) Get() *AutoSyncManagementCommand {
	return v.value
}

func (v *NullableAutoSyncManagementCommand) Set(val *AutoSyncManagementCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoSyncManagementCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoSyncManagementCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoSyncManagementCommand(val *AutoSyncManagementCommand) *NullableAutoSyncManagementCommand {
	return &NullableAutoSyncManagementCommand{value: val, isSet: true}
}

func (v NullableAutoSyncManagementCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoSyncManagementCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


