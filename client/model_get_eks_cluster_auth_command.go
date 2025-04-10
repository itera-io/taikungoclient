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

// checks if the GetEksClusterAuthCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEksClusterAuthCommand{}

// GetEksClusterAuthCommand struct for GetEksClusterAuthCommand
type GetEksClusterAuthCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	KubeConfigName NullableString `json:"kubeConfigName,omitempty"`
	ConfigToken NullableString `json:"configToken,omitempty"`
}

// NewGetEksClusterAuthCommand instantiates a new GetEksClusterAuthCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEksClusterAuthCommand() *GetEksClusterAuthCommand {
	this := GetEksClusterAuthCommand{}
	return &this
}

// NewGetEksClusterAuthCommandWithDefaults instantiates a new GetEksClusterAuthCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEksClusterAuthCommandWithDefaults() *GetEksClusterAuthCommand {
	this := GetEksClusterAuthCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetEksClusterAuthCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEksClusterAuthCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetEksClusterAuthCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *GetEksClusterAuthCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetKubeConfigName returns the KubeConfigName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetEksClusterAuthCommand) GetKubeConfigName() string {
	if o == nil || IsNil(o.KubeConfigName.Get()) {
		var ret string
		return ret
	}
	return *o.KubeConfigName.Get()
}

// GetKubeConfigNameOk returns a tuple with the KubeConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEksClusterAuthCommand) GetKubeConfigNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubeConfigName.Get(), o.KubeConfigName.IsSet()
}

// HasKubeConfigName returns a boolean if a field has been set.
func (o *GetEksClusterAuthCommand) HasKubeConfigName() bool {
	if o != nil && o.KubeConfigName.IsSet() {
		return true
	}

	return false
}

// SetKubeConfigName gets a reference to the given NullableString and assigns it to the KubeConfigName field.
func (o *GetEksClusterAuthCommand) SetKubeConfigName(v string) {
	o.KubeConfigName.Set(&v)
}
// SetKubeConfigNameNil sets the value for KubeConfigName to be an explicit nil
func (o *GetEksClusterAuthCommand) SetKubeConfigNameNil() {
	o.KubeConfigName.Set(nil)
}

// UnsetKubeConfigName ensures that no value is present for KubeConfigName, not even an explicit nil
func (o *GetEksClusterAuthCommand) UnsetKubeConfigName() {
	o.KubeConfigName.Unset()
}

// GetConfigToken returns the ConfigToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetEksClusterAuthCommand) GetConfigToken() string {
	if o == nil || IsNil(o.ConfigToken.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigToken.Get()
}

// GetConfigTokenOk returns a tuple with the ConfigToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEksClusterAuthCommand) GetConfigTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigToken.Get(), o.ConfigToken.IsSet()
}

// HasConfigToken returns a boolean if a field has been set.
func (o *GetEksClusterAuthCommand) HasConfigToken() bool {
	if o != nil && o.ConfigToken.IsSet() {
		return true
	}

	return false
}

// SetConfigToken gets a reference to the given NullableString and assigns it to the ConfigToken field.
func (o *GetEksClusterAuthCommand) SetConfigToken(v string) {
	o.ConfigToken.Set(&v)
}
// SetConfigTokenNil sets the value for ConfigToken to be an explicit nil
func (o *GetEksClusterAuthCommand) SetConfigTokenNil() {
	o.ConfigToken.Set(nil)
}

// UnsetConfigToken ensures that no value is present for ConfigToken, not even an explicit nil
func (o *GetEksClusterAuthCommand) UnsetConfigToken() {
	o.ConfigToken.Unset()
}

func (o GetEksClusterAuthCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEksClusterAuthCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.KubeConfigName.IsSet() {
		toSerialize["kubeConfigName"] = o.KubeConfigName.Get()
	}
	if o.ConfigToken.IsSet() {
		toSerialize["configToken"] = o.ConfigToken.Get()
	}
	return toSerialize, nil
}

type NullableGetEksClusterAuthCommand struct {
	value *GetEksClusterAuthCommand
	isSet bool
}

func (v NullableGetEksClusterAuthCommand) Get() *GetEksClusterAuthCommand {
	return v.value
}

func (v *NullableGetEksClusterAuthCommand) Set(val *GetEksClusterAuthCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEksClusterAuthCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEksClusterAuthCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEksClusterAuthCommand(val *GetEksClusterAuthCommand) *NullableGetEksClusterAuthCommand {
	return &NullableGetEksClusterAuthCommand{value: val, isSet: true}
}

func (v NullableGetEksClusterAuthCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEksClusterAuthCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


