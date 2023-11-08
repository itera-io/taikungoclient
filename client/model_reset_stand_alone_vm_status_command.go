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

// checks if the ResetStandAloneVmStatusCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetStandAloneVmStatusCommand{}

// ResetStandAloneVmStatusCommand struct for ResetStandAloneVmStatusCommand
type ResetStandAloneVmStatusCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	VmIds []int32 `json:"vmIds,omitempty"`
	Status *StandAloneVmStatus `json:"status,omitempty"`
}

// NewResetStandAloneVmStatusCommand instantiates a new ResetStandAloneVmStatusCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetStandAloneVmStatusCommand() *ResetStandAloneVmStatusCommand {
	this := ResetStandAloneVmStatusCommand{}
	return &this
}

// NewResetStandAloneVmStatusCommandWithDefaults instantiates a new ResetStandAloneVmStatusCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetStandAloneVmStatusCommandWithDefaults() *ResetStandAloneVmStatusCommand {
	this := ResetStandAloneVmStatusCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ResetStandAloneVmStatusCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetStandAloneVmStatusCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ResetStandAloneVmStatusCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ResetStandAloneVmStatusCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResetStandAloneVmStatusCommand) GetVmIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResetStandAloneVmStatusCommand) GetVmIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.VmIds) {
		return nil, false
	}
	return o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *ResetStandAloneVmStatusCommand) HasVmIds() bool {
	if o != nil && IsNil(o.VmIds) {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []int32 and assigns it to the VmIds field.
func (o *ResetStandAloneVmStatusCommand) SetVmIds(v []int32) {
	o.VmIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResetStandAloneVmStatusCommand) GetStatus() StandAloneVmStatus {
	if o == nil || IsNil(o.Status) {
		var ret StandAloneVmStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetStandAloneVmStatusCommand) GetStatusOk() (*StandAloneVmStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResetStandAloneVmStatusCommand) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StandAloneVmStatus and assigns it to the Status field.
func (o *ResetStandAloneVmStatusCommand) SetStatus(v StandAloneVmStatus) {
	o.Status = &v
}

func (o ResetStandAloneVmStatusCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetStandAloneVmStatusCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.VmIds != nil {
		toSerialize["vmIds"] = o.VmIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResetStandAloneVmStatusCommand struct {
	value *ResetStandAloneVmStatusCommand
	isSet bool
}

func (v NullableResetStandAloneVmStatusCommand) Get() *ResetStandAloneVmStatusCommand {
	return v.value
}

func (v *NullableResetStandAloneVmStatusCommand) Set(val *ResetStandAloneVmStatusCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableResetStandAloneVmStatusCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableResetStandAloneVmStatusCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetStandAloneVmStatusCommand(val *ResetStandAloneVmStatusCommand) *NullableResetStandAloneVmStatusCommand {
	return &NullableResetStandAloneVmStatusCommand{value: val, isSet: true}
}

func (v NullableResetStandAloneVmStatusCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetStandAloneVmStatusCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


