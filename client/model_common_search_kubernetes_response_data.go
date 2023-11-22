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

// checks if the CommonSearchKubernetesResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonSearchKubernetesResponseData{}

// CommonSearchKubernetesResponseData struct for CommonSearchKubernetesResponseData
type CommonSearchKubernetesResponseData struct {
	MetadataName     NullableString `json:"metadataName,omitempty"`
	Namespace        NullableString `json:"namespace,omitempty"`
	ProjectId        *int32         `json:"projectId,omitempty"`
	ProjectName      NullableString `json:"projectName,omitempty"`
	OrganizationId   *int32         `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
}

// NewCommonSearchKubernetesResponseData instantiates a new CommonSearchKubernetesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonSearchKubernetesResponseData() *CommonSearchKubernetesResponseData {
	this := CommonSearchKubernetesResponseData{}
	return &this
}

// NewCommonSearchKubernetesResponseDataWithDefaults instantiates a new CommonSearchKubernetesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonSearchKubernetesResponseDataWithDefaults() *CommonSearchKubernetesResponseData {
	this := CommonSearchKubernetesResponseData{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSearchKubernetesResponseData) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSearchKubernetesResponseData) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *CommonSearchKubernetesResponseData) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}

// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *CommonSearchKubernetesResponseData) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *CommonSearchKubernetesResponseData) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSearchKubernetesResponseData) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSearchKubernetesResponseData) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *CommonSearchKubernetesResponseData) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *CommonSearchKubernetesResponseData) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *CommonSearchKubernetesResponseData) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CommonSearchKubernetesResponseData) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSearchKubernetesResponseData) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CommonSearchKubernetesResponseData) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSearchKubernetesResponseData) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSearchKubernetesResponseData) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *CommonSearchKubernetesResponseData) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *CommonSearchKubernetesResponseData) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *CommonSearchKubernetesResponseData) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CommonSearchKubernetesResponseData) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSearchKubernetesResponseData) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *CommonSearchKubernetesResponseData) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSearchKubernetesResponseData) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSearchKubernetesResponseData) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *CommonSearchKubernetesResponseData) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *CommonSearchKubernetesResponseData) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *CommonSearchKubernetesResponseData) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *CommonSearchKubernetesResponseData) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

func (o CommonSearchKubernetesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonSearchKubernetesResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	return toSerialize, nil
}

type NullableCommonSearchKubernetesResponseData struct {
	value *CommonSearchKubernetesResponseData
	isSet bool
}

func (v NullableCommonSearchKubernetesResponseData) Get() *CommonSearchKubernetesResponseData {
	return v.value
}

func (v *NullableCommonSearchKubernetesResponseData) Set(val *CommonSearchKubernetesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonSearchKubernetesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonSearchKubernetesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonSearchKubernetesResponseData(val *CommonSearchKubernetesResponseData) *NullableCommonSearchKubernetesResponseData {
	return &NullableCommonSearchKubernetesResponseData{value: val, isSet: true}
}

func (v NullableCommonSearchKubernetesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonSearchKubernetesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
