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

// checks if the DeleteServerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteServerCommand{}

// DeleteServerCommand struct for DeleteServerCommand
type DeleteServerCommand struct {
	ProjectId                *int32  `json:"projectId,omitempty"`
	ServerIds                []int32 `json:"serverIds,omitempty"`
	DeleteAutoscalingServers *bool   `json:"deleteAutoscalingServers,omitempty"`
}

// NewDeleteServerCommand instantiates a new DeleteServerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteServerCommand() *DeleteServerCommand {
	this := DeleteServerCommand{}
	return &this
}

// NewDeleteServerCommandWithDefaults instantiates a new DeleteServerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteServerCommandWithDefaults() *DeleteServerCommand {
	this := DeleteServerCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteServerCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteServerCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteServerCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteServerCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetServerIds returns the ServerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteServerCommand) GetServerIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ServerIds
}

// GetServerIdsOk returns a tuple with the ServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteServerCommand) GetServerIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ServerIds) {
		return nil, false
	}
	return o.ServerIds, true
}

// HasServerIds returns a boolean if a field has been set.
func (o *DeleteServerCommand) HasServerIds() bool {
	if o != nil && IsNil(o.ServerIds) {
		return true
	}

	return false
}

// SetServerIds gets a reference to the given []int32 and assigns it to the ServerIds field.
func (o *DeleteServerCommand) SetServerIds(v []int32) {
	o.ServerIds = v
}

// GetDeleteAutoscalingServers returns the DeleteAutoscalingServers field value if set, zero value otherwise.
func (o *DeleteServerCommand) GetDeleteAutoscalingServers() bool {
	if o == nil || IsNil(o.DeleteAutoscalingServers) {
		var ret bool
		return ret
	}
	return *o.DeleteAutoscalingServers
}

// GetDeleteAutoscalingServersOk returns a tuple with the DeleteAutoscalingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteServerCommand) GetDeleteAutoscalingServersOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAutoscalingServers) {
		return nil, false
	}
	return o.DeleteAutoscalingServers, true
}

// HasDeleteAutoscalingServers returns a boolean if a field has been set.
func (o *DeleteServerCommand) HasDeleteAutoscalingServers() bool {
	if o != nil && !IsNil(o.DeleteAutoscalingServers) {
		return true
	}

	return false
}

// SetDeleteAutoscalingServers gets a reference to the given bool and assigns it to the DeleteAutoscalingServers field.
func (o *DeleteServerCommand) SetDeleteAutoscalingServers(v bool) {
	o.DeleteAutoscalingServers = &v
}

func (o DeleteServerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteServerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ServerIds != nil {
		toSerialize["serverIds"] = o.ServerIds
	}
	if !IsNil(o.DeleteAutoscalingServers) {
		toSerialize["deleteAutoscalingServers"] = o.DeleteAutoscalingServers
	}
	return toSerialize, nil
}

type NullableDeleteServerCommand struct {
	value *DeleteServerCommand
	isSet bool
}

func (v NullableDeleteServerCommand) Get() *DeleteServerCommand {
	return v.value
}

func (v *NullableDeleteServerCommand) Set(val *DeleteServerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteServerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteServerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteServerCommand(val *DeleteServerCommand) *NullableDeleteServerCommand {
	return &NullableDeleteServerCommand{value: val, isSet: true}
}

func (v NullableDeleteServerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteServerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
