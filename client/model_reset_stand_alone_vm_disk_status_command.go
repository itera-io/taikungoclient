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

// checks if the ResetStandAloneVmDiskStatusCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetStandAloneVmDiskStatusCommand{}

// ResetStandAloneVmDiskStatusCommand struct for ResetStandAloneVmDiskStatusCommand
type ResetStandAloneVmDiskStatusCommand struct {
	StandAloneVmId *int32 `json:"standAloneVmId,omitempty"`
	DiskIds []int32 `json:"diskIds,omitempty"`
	Status *StandAloneVmDiskStatus `json:"status,omitempty"`
}

// NewResetStandAloneVmDiskStatusCommand instantiates a new ResetStandAloneVmDiskStatusCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetStandAloneVmDiskStatusCommand() *ResetStandAloneVmDiskStatusCommand {
	this := ResetStandAloneVmDiskStatusCommand{}
	return &this
}

// NewResetStandAloneVmDiskStatusCommandWithDefaults instantiates a new ResetStandAloneVmDiskStatusCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetStandAloneVmDiskStatusCommandWithDefaults() *ResetStandAloneVmDiskStatusCommand {
	this := ResetStandAloneVmDiskStatusCommand{}
	return &this
}

// GetStandAloneVmId returns the StandAloneVmId field value if set, zero value otherwise.
func (o *ResetStandAloneVmDiskStatusCommand) GetStandAloneVmId() int32 {
	if o == nil || IsNil(o.StandAloneVmId) {
		var ret int32
		return ret
	}
	return *o.StandAloneVmId
}

// GetStandAloneVmIdOk returns a tuple with the StandAloneVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetStandAloneVmDiskStatusCommand) GetStandAloneVmIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StandAloneVmId) {
		return nil, false
	}
	return o.StandAloneVmId, true
}

// HasStandAloneVmId returns a boolean if a field has been set.
func (o *ResetStandAloneVmDiskStatusCommand) HasStandAloneVmId() bool {
	if o != nil && !IsNil(o.StandAloneVmId) {
		return true
	}

	return false
}

// SetStandAloneVmId gets a reference to the given int32 and assigns it to the StandAloneVmId field.
func (o *ResetStandAloneVmDiskStatusCommand) SetStandAloneVmId(v int32) {
	o.StandAloneVmId = &v
}

// GetDiskIds returns the DiskIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResetStandAloneVmDiskStatusCommand) GetDiskIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.DiskIds
}

// GetDiskIdsOk returns a tuple with the DiskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResetStandAloneVmDiskStatusCommand) GetDiskIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DiskIds) {
		return nil, false
	}
	return o.DiskIds, true
}

// HasDiskIds returns a boolean if a field has been set.
func (o *ResetStandAloneVmDiskStatusCommand) HasDiskIds() bool {
	if o != nil && !IsNil(o.DiskIds) {
		return true
	}

	return false
}

// SetDiskIds gets a reference to the given []int32 and assigns it to the DiskIds field.
func (o *ResetStandAloneVmDiskStatusCommand) SetDiskIds(v []int32) {
	o.DiskIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResetStandAloneVmDiskStatusCommand) GetStatus() StandAloneVmDiskStatus {
	if o == nil || IsNil(o.Status) {
		var ret StandAloneVmDiskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetStandAloneVmDiskStatusCommand) GetStatusOk() (*StandAloneVmDiskStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResetStandAloneVmDiskStatusCommand) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StandAloneVmDiskStatus and assigns it to the Status field.
func (o *ResetStandAloneVmDiskStatusCommand) SetStatus(v StandAloneVmDiskStatus) {
	o.Status = &v
}

func (o ResetStandAloneVmDiskStatusCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetStandAloneVmDiskStatusCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StandAloneVmId) {
		toSerialize["standAloneVmId"] = o.StandAloneVmId
	}
	if o.DiskIds != nil {
		toSerialize["diskIds"] = o.DiskIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResetStandAloneVmDiskStatusCommand struct {
	value *ResetStandAloneVmDiskStatusCommand
	isSet bool
}

func (v NullableResetStandAloneVmDiskStatusCommand) Get() *ResetStandAloneVmDiskStatusCommand {
	return v.value
}

func (v *NullableResetStandAloneVmDiskStatusCommand) Set(val *ResetStandAloneVmDiskStatusCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableResetStandAloneVmDiskStatusCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableResetStandAloneVmDiskStatusCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetStandAloneVmDiskStatusCommand(val *ResetStandAloneVmDiskStatusCommand) *NullableResetStandAloneVmDiskStatusCommand {
	return &NullableResetStandAloneVmDiskStatusCommand{value: val, isSet: true}
}

func (v NullableResetStandAloneVmDiskStatusCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetStandAloneVmDiskStatusCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


