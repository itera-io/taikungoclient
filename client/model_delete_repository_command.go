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

// checks if the DeleteRepositoryCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRepositoryCommand{}

// DeleteRepositoryCommand struct for DeleteRepositoryCommand
type DeleteRepositoryCommand struct {
	AppRepoId *int32 `json:"appRepoId,omitempty"`
}

// NewDeleteRepositoryCommand instantiates a new DeleteRepositoryCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRepositoryCommand() *DeleteRepositoryCommand {
	this := DeleteRepositoryCommand{}
	return &this
}

// NewDeleteRepositoryCommandWithDefaults instantiates a new DeleteRepositoryCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRepositoryCommandWithDefaults() *DeleteRepositoryCommand {
	this := DeleteRepositoryCommand{}
	return &this
}

// GetAppRepoId returns the AppRepoId field value if set, zero value otherwise.
func (o *DeleteRepositoryCommand) GetAppRepoId() int32 {
	if o == nil || IsNil(o.AppRepoId) {
		var ret int32
		return ret
	}
	return *o.AppRepoId
}

// GetAppRepoIdOk returns a tuple with the AppRepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRepositoryCommand) GetAppRepoIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AppRepoId) {
		return nil, false
	}
	return o.AppRepoId, true
}

// HasAppRepoId returns a boolean if a field has been set.
func (o *DeleteRepositoryCommand) HasAppRepoId() bool {
	if o != nil && !IsNil(o.AppRepoId) {
		return true
	}

	return false
}

// SetAppRepoId gets a reference to the given int32 and assigns it to the AppRepoId field.
func (o *DeleteRepositoryCommand) SetAppRepoId(v int32) {
	o.AppRepoId = &v
}

func (o DeleteRepositoryCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRepositoryCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppRepoId) {
		toSerialize["appRepoId"] = o.AppRepoId
	}
	return toSerialize, nil
}

type NullableDeleteRepositoryCommand struct {
	value *DeleteRepositoryCommand
	isSet bool
}

func (v NullableDeleteRepositoryCommand) Get() *DeleteRepositoryCommand {
	return v.value
}

func (v *NullableDeleteRepositoryCommand) Set(val *DeleteRepositoryCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRepositoryCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRepositoryCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRepositoryCommand(val *DeleteRepositoryCommand) *NullableDeleteRepositoryCommand {
	return &NullableDeleteRepositoryCommand{value: val, isSet: true}
}

func (v NullableDeleteRepositoryCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRepositoryCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


