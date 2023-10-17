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

// checks if the CreateKubeConfigCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKubeConfigCommand{}

// CreateKubeConfigCommand struct for CreateKubeConfigCommand
type CreateKubeConfigCommand struct {
	Name                   NullableString `json:"name,omitempty"`
	ProjectId              *int32         `json:"projectId,omitempty"`
	IsAccessibleForAll     *bool          `json:"isAccessibleForAll,omitempty"`
	IsAccessibleForManager *bool          `json:"isAccessibleForManager,omitempty"`
	KubeConfigRoleId       *int32         `json:"kubeConfigRoleId,omitempty"`
	UserId                 NullableString `json:"userId,omitempty"`
	Namespace              NullableString `json:"namespace,omitempty"`
	Ttl                    NullableInt32  `json:"ttl,omitempty"`
}

// NewCreateKubeConfigCommand instantiates a new CreateKubeConfigCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKubeConfigCommand() *CreateKubeConfigCommand {
	this := CreateKubeConfigCommand{}
	return &this
}

// NewCreateKubeConfigCommandWithDefaults instantiates a new CreateKubeConfigCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKubeConfigCommandWithDefaults() *CreateKubeConfigCommand {
	this := CreateKubeConfigCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKubeConfigCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKubeConfigCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateKubeConfigCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateKubeConfigCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateKubeConfigCommand) UnsetName() {
	o.Name.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateKubeConfigCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubeConfigCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CreateKubeConfigCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetIsAccessibleForAll returns the IsAccessibleForAll field value if set, zero value otherwise.
func (o *CreateKubeConfigCommand) GetIsAccessibleForAll() bool {
	if o == nil || IsNil(o.IsAccessibleForAll) {
		var ret bool
		return ret
	}
	return *o.IsAccessibleForAll
}

// GetIsAccessibleForAllOk returns a tuple with the IsAccessibleForAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubeConfigCommand) GetIsAccessibleForAllOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccessibleForAll) {
		return nil, false
	}
	return o.IsAccessibleForAll, true
}

// HasIsAccessibleForAll returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasIsAccessibleForAll() bool {
	if o != nil && !IsNil(o.IsAccessibleForAll) {
		return true
	}

	return false
}

// SetIsAccessibleForAll gets a reference to the given bool and assigns it to the IsAccessibleForAll field.
func (o *CreateKubeConfigCommand) SetIsAccessibleForAll(v bool) {
	o.IsAccessibleForAll = &v
}

// GetIsAccessibleForManager returns the IsAccessibleForManager field value if set, zero value otherwise.
func (o *CreateKubeConfigCommand) GetIsAccessibleForManager() bool {
	if o == nil || IsNil(o.IsAccessibleForManager) {
		var ret bool
		return ret
	}
	return *o.IsAccessibleForManager
}

// GetIsAccessibleForManagerOk returns a tuple with the IsAccessibleForManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubeConfigCommand) GetIsAccessibleForManagerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccessibleForManager) {
		return nil, false
	}
	return o.IsAccessibleForManager, true
}

// HasIsAccessibleForManager returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasIsAccessibleForManager() bool {
	if o != nil && !IsNil(o.IsAccessibleForManager) {
		return true
	}

	return false
}

// SetIsAccessibleForManager gets a reference to the given bool and assigns it to the IsAccessibleForManager field.
func (o *CreateKubeConfigCommand) SetIsAccessibleForManager(v bool) {
	o.IsAccessibleForManager = &v
}

// GetKubeConfigRoleId returns the KubeConfigRoleId field value if set, zero value otherwise.
func (o *CreateKubeConfigCommand) GetKubeConfigRoleId() int32 {
	if o == nil || IsNil(o.KubeConfigRoleId) {
		var ret int32
		return ret
	}
	return *o.KubeConfigRoleId
}

// GetKubeConfigRoleIdOk returns a tuple with the KubeConfigRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubeConfigCommand) GetKubeConfigRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.KubeConfigRoleId) {
		return nil, false
	}
	return o.KubeConfigRoleId, true
}

// HasKubeConfigRoleId returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasKubeConfigRoleId() bool {
	if o != nil && !IsNil(o.KubeConfigRoleId) {
		return true
	}

	return false
}

// SetKubeConfigRoleId gets a reference to the given int32 and assigns it to the KubeConfigRoleId field.
func (o *CreateKubeConfigCommand) SetKubeConfigRoleId(v int32) {
	o.KubeConfigRoleId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKubeConfigCommand) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKubeConfigCommand) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *CreateKubeConfigCommand) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CreateKubeConfigCommand) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CreateKubeConfigCommand) UnsetUserId() {
	o.UserId.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKubeConfigCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKubeConfigCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *CreateKubeConfigCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *CreateKubeConfigCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *CreateKubeConfigCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetTtl returns the Ttl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKubeConfigCommand) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl.Get()) {
		var ret int32
		return ret
	}
	return *o.Ttl.Get()
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKubeConfigCommand) GetTtlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ttl.Get(), o.Ttl.IsSet()
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateKubeConfigCommand) HasTtl() bool {
	if o != nil && o.Ttl.IsSet() {
		return true
	}

	return false
}

// SetTtl gets a reference to the given NullableInt32 and assigns it to the Ttl field.
func (o *CreateKubeConfigCommand) SetTtl(v int32) {
	o.Ttl.Set(&v)
}

// SetTtlNil sets the value for Ttl to be an explicit nil
func (o *CreateKubeConfigCommand) SetTtlNil() {
	o.Ttl.Set(nil)
}

// UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
func (o *CreateKubeConfigCommand) UnsetTtl() {
	o.Ttl.Unset()
}

func (o CreateKubeConfigCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKubeConfigCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsAccessibleForAll) {
		toSerialize["isAccessibleForAll"] = o.IsAccessibleForAll
	}
	if !IsNil(o.IsAccessibleForManager) {
		toSerialize["isAccessibleForManager"] = o.IsAccessibleForManager
	}
	if !IsNil(o.KubeConfigRoleId) {
		toSerialize["kubeConfigRoleId"] = o.KubeConfigRoleId
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Ttl.IsSet() {
		toSerialize["ttl"] = o.Ttl.Get()
	}
	return toSerialize, nil
}

type NullableCreateKubeConfigCommand struct {
	value *CreateKubeConfigCommand
	isSet bool
}

func (v NullableCreateKubeConfigCommand) Get() *CreateKubeConfigCommand {
	return v.value
}

func (v *NullableCreateKubeConfigCommand) Set(val *CreateKubeConfigCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKubeConfigCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKubeConfigCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKubeConfigCommand(val *CreateKubeConfigCommand) *NullableCreateKubeConfigCommand {
	return &NullableCreateKubeConfigCommand{value: val, isSet: true}
}

func (v NullableCreateKubeConfigCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKubeConfigCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
