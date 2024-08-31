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

// checks if the ProjectDeploymentDeleteServersCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDeploymentDeleteServersCommand{}

// ProjectDeploymentDeleteServersCommand struct for ProjectDeploymentDeleteServersCommand
type ProjectDeploymentDeleteServersCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ServerIds []int32 `json:"serverIds,omitempty"`
	ForceDeleteVClusters *bool `json:"forceDeleteVClusters,omitempty"`
	DeleteAutoscalingServers *bool `json:"deleteAutoscalingServers,omitempty"`
}

// NewProjectDeploymentDeleteServersCommand instantiates a new ProjectDeploymentDeleteServersCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeploymentDeleteServersCommand() *ProjectDeploymentDeleteServersCommand {
	this := ProjectDeploymentDeleteServersCommand{}
	return &this
}

// NewProjectDeploymentDeleteServersCommandWithDefaults instantiates a new ProjectDeploymentDeleteServersCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeploymentDeleteServersCommandWithDefaults() *ProjectDeploymentDeleteServersCommand {
	this := ProjectDeploymentDeleteServersCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectDeploymentDeleteServersCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentDeleteServersCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectDeploymentDeleteServersCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ProjectDeploymentDeleteServersCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetServerIds returns the ServerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDeploymentDeleteServersCommand) GetServerIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ServerIds
}

// GetServerIdsOk returns a tuple with the ServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDeploymentDeleteServersCommand) GetServerIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ServerIds) {
		return nil, false
	}
	return o.ServerIds, true
}

// HasServerIds returns a boolean if a field has been set.
func (o *ProjectDeploymentDeleteServersCommand) HasServerIds() bool {
	if o != nil && !IsNil(o.ServerIds) {
		return true
	}

	return false
}

// SetServerIds gets a reference to the given []int32 and assigns it to the ServerIds field.
func (o *ProjectDeploymentDeleteServersCommand) SetServerIds(v []int32) {
	o.ServerIds = v
}

// GetForceDeleteVClusters returns the ForceDeleteVClusters field value if set, zero value otherwise.
func (o *ProjectDeploymentDeleteServersCommand) GetForceDeleteVClusters() bool {
	if o == nil || IsNil(o.ForceDeleteVClusters) {
		var ret bool
		return ret
	}
	return *o.ForceDeleteVClusters
}

// GetForceDeleteVClustersOk returns a tuple with the ForceDeleteVClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentDeleteServersCommand) GetForceDeleteVClustersOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceDeleteVClusters) {
		return nil, false
	}
	return o.ForceDeleteVClusters, true
}

// HasForceDeleteVClusters returns a boolean if a field has been set.
func (o *ProjectDeploymentDeleteServersCommand) HasForceDeleteVClusters() bool {
	if o != nil && !IsNil(o.ForceDeleteVClusters) {
		return true
	}

	return false
}

// SetForceDeleteVClusters gets a reference to the given bool and assigns it to the ForceDeleteVClusters field.
func (o *ProjectDeploymentDeleteServersCommand) SetForceDeleteVClusters(v bool) {
	o.ForceDeleteVClusters = &v
}

// GetDeleteAutoscalingServers returns the DeleteAutoscalingServers field value if set, zero value otherwise.
func (o *ProjectDeploymentDeleteServersCommand) GetDeleteAutoscalingServers() bool {
	if o == nil || IsNil(o.DeleteAutoscalingServers) {
		var ret bool
		return ret
	}
	return *o.DeleteAutoscalingServers
}

// GetDeleteAutoscalingServersOk returns a tuple with the DeleteAutoscalingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentDeleteServersCommand) GetDeleteAutoscalingServersOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAutoscalingServers) {
		return nil, false
	}
	return o.DeleteAutoscalingServers, true
}

// HasDeleteAutoscalingServers returns a boolean if a field has been set.
func (o *ProjectDeploymentDeleteServersCommand) HasDeleteAutoscalingServers() bool {
	if o != nil && !IsNil(o.DeleteAutoscalingServers) {
		return true
	}

	return false
}

// SetDeleteAutoscalingServers gets a reference to the given bool and assigns it to the DeleteAutoscalingServers field.
func (o *ProjectDeploymentDeleteServersCommand) SetDeleteAutoscalingServers(v bool) {
	o.DeleteAutoscalingServers = &v
}

func (o ProjectDeploymentDeleteServersCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDeploymentDeleteServersCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ServerIds != nil {
		toSerialize["serverIds"] = o.ServerIds
	}
	if !IsNil(o.ForceDeleteVClusters) {
		toSerialize["forceDeleteVClusters"] = o.ForceDeleteVClusters
	}
	if !IsNil(o.DeleteAutoscalingServers) {
		toSerialize["deleteAutoscalingServers"] = o.DeleteAutoscalingServers
	}
	return toSerialize, nil
}

type NullableProjectDeploymentDeleteServersCommand struct {
	value *ProjectDeploymentDeleteServersCommand
	isSet bool
}

func (v NullableProjectDeploymentDeleteServersCommand) Get() *ProjectDeploymentDeleteServersCommand {
	return v.value
}

func (v *NullableProjectDeploymentDeleteServersCommand) Set(val *ProjectDeploymentDeleteServersCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeploymentDeleteServersCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeploymentDeleteServersCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeploymentDeleteServersCommand(val *ProjectDeploymentDeleteServersCommand) *NullableProjectDeploymentDeleteServersCommand {
	return &NullableProjectDeploymentDeleteServersCommand{value: val, isSet: true}
}

func (v NullableProjectDeploymentDeleteServersCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeploymentDeleteServersCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

