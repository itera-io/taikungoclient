/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the KubernetesOverviewDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesOverviewDto{}

// KubernetesOverviewDto struct for KubernetesOverviewDto
type KubernetesOverviewDto struct {
	ProjectId        *int32         `json:"projectId,omitempty"`
	ProjectName      NullableString `json:"projectName,omitempty"`
	HealthyPods      *int32         `json:"healthyPods,omitempty"`
	UnhealthyPods    *int32         `json:"unhealthyPods,omitempty"`
	HealthyNodes     *int32         `json:"healthyNodes,omitempty"`
	UnhealthyNodes   *int32         `json:"unhealthyNodes,omitempty"`
	AlertsCount      *int32         `json:"alertsCount,omitempty"`
	KubernetesHealth NullableString `json:"kubernetesHealth,omitempty"`
}

// NewKubernetesOverviewDto instantiates a new KubernetesOverviewDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesOverviewDto() *KubernetesOverviewDto {
	this := KubernetesOverviewDto{}
	return &this
}

// NewKubernetesOverviewDtoWithDefaults instantiates a new KubernetesOverviewDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesOverviewDtoWithDefaults() *KubernetesOverviewDto {
	this := KubernetesOverviewDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubernetesOverviewDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesOverviewDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesOverviewDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *KubernetesOverviewDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *KubernetesOverviewDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *KubernetesOverviewDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetHealthyPods returns the HealthyPods field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetHealthyPods() int32 {
	if o == nil || IsNil(o.HealthyPods) {
		var ret int32
		return ret
	}
	return *o.HealthyPods
}

// GetHealthyPodsOk returns a tuple with the HealthyPods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetHealthyPodsOk() (*int32, bool) {
	if o == nil || IsNil(o.HealthyPods) {
		return nil, false
	}
	return o.HealthyPods, true
}

// HasHealthyPods returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasHealthyPods() bool {
	if o != nil && !IsNil(o.HealthyPods) {
		return true
	}

	return false
}

// SetHealthyPods gets a reference to the given int32 and assigns it to the HealthyPods field.
func (o *KubernetesOverviewDto) SetHealthyPods(v int32) {
	o.HealthyPods = &v
}

// GetUnhealthyPods returns the UnhealthyPods field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetUnhealthyPods() int32 {
	if o == nil || IsNil(o.UnhealthyPods) {
		var ret int32
		return ret
	}
	return *o.UnhealthyPods
}

// GetUnhealthyPodsOk returns a tuple with the UnhealthyPods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetUnhealthyPodsOk() (*int32, bool) {
	if o == nil || IsNil(o.UnhealthyPods) {
		return nil, false
	}
	return o.UnhealthyPods, true
}

// HasUnhealthyPods returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasUnhealthyPods() bool {
	if o != nil && !IsNil(o.UnhealthyPods) {
		return true
	}

	return false
}

// SetUnhealthyPods gets a reference to the given int32 and assigns it to the UnhealthyPods field.
func (o *KubernetesOverviewDto) SetUnhealthyPods(v int32) {
	o.UnhealthyPods = &v
}

// GetHealthyNodes returns the HealthyNodes field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetHealthyNodes() int32 {
	if o == nil || IsNil(o.HealthyNodes) {
		var ret int32
		return ret
	}
	return *o.HealthyNodes
}

// GetHealthyNodesOk returns a tuple with the HealthyNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetHealthyNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.HealthyNodes) {
		return nil, false
	}
	return o.HealthyNodes, true
}

// HasHealthyNodes returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasHealthyNodes() bool {
	if o != nil && !IsNil(o.HealthyNodes) {
		return true
	}

	return false
}

// SetHealthyNodes gets a reference to the given int32 and assigns it to the HealthyNodes field.
func (o *KubernetesOverviewDto) SetHealthyNodes(v int32) {
	o.HealthyNodes = &v
}

// GetUnhealthyNodes returns the UnhealthyNodes field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetUnhealthyNodes() int32 {
	if o == nil || IsNil(o.UnhealthyNodes) {
		var ret int32
		return ret
	}
	return *o.UnhealthyNodes
}

// GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetUnhealthyNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.UnhealthyNodes) {
		return nil, false
	}
	return o.UnhealthyNodes, true
}

// HasUnhealthyNodes returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasUnhealthyNodes() bool {
	if o != nil && !IsNil(o.UnhealthyNodes) {
		return true
	}

	return false
}

// SetUnhealthyNodes gets a reference to the given int32 and assigns it to the UnhealthyNodes field.
func (o *KubernetesOverviewDto) SetUnhealthyNodes(v int32) {
	o.UnhealthyNodes = &v
}

// GetAlertsCount returns the AlertsCount field value if set, zero value otherwise.
func (o *KubernetesOverviewDto) GetAlertsCount() int32 {
	if o == nil || IsNil(o.AlertsCount) {
		var ret int32
		return ret
	}
	return *o.AlertsCount
}

// GetAlertsCountOk returns a tuple with the AlertsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetAlertsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AlertsCount) {
		return nil, false
	}
	return o.AlertsCount, true
}

// HasAlertsCount returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasAlertsCount() bool {
	if o != nil && !IsNil(o.AlertsCount) {
		return true
	}

	return false
}

// SetAlertsCount gets a reference to the given int32 and assigns it to the AlertsCount field.
func (o *KubernetesOverviewDto) SetAlertsCount(v int32) {
	o.AlertsCount = &v
}

// GetKubernetesHealth returns the KubernetesHealth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesOverviewDto) GetKubernetesHealth() string {
	if o == nil || IsNil(o.KubernetesHealth.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesHealth.Get()
}

// GetKubernetesHealthOk returns a tuple with the KubernetesHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesOverviewDto) GetKubernetesHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesHealth.Get(), o.KubernetesHealth.IsSet()
}

// HasKubernetesHealth returns a boolean if a field has been set.
func (o *KubernetesOverviewDto) HasKubernetesHealth() bool {
	if o != nil && o.KubernetesHealth.IsSet() {
		return true
	}

	return false
}

// SetKubernetesHealth gets a reference to the given NullableString and assigns it to the KubernetesHealth field.
func (o *KubernetesOverviewDto) SetKubernetesHealth(v string) {
	o.KubernetesHealth.Set(&v)
}

// SetKubernetesHealthNil sets the value for KubernetesHealth to be an explicit nil
func (o *KubernetesOverviewDto) SetKubernetesHealthNil() {
	o.KubernetesHealth.Set(nil)
}

// UnsetKubernetesHealth ensures that no value is present for KubernetesHealth, not even an explicit nil
func (o *KubernetesOverviewDto) UnsetKubernetesHealth() {
	o.KubernetesHealth.Unset()
}

func (o KubernetesOverviewDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesOverviewDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.HealthyPods) {
		toSerialize["healthyPods"] = o.HealthyPods
	}
	if !IsNil(o.UnhealthyPods) {
		toSerialize["unhealthyPods"] = o.UnhealthyPods
	}
	if !IsNil(o.HealthyNodes) {
		toSerialize["healthyNodes"] = o.HealthyNodes
	}
	if !IsNil(o.UnhealthyNodes) {
		toSerialize["unhealthyNodes"] = o.UnhealthyNodes
	}
	if !IsNil(o.AlertsCount) {
		toSerialize["alertsCount"] = o.AlertsCount
	}
	if o.KubernetesHealth.IsSet() {
		toSerialize["kubernetesHealth"] = o.KubernetesHealth.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesOverviewDto struct {
	value *KubernetesOverviewDto
	isSet bool
}

func (v NullableKubernetesOverviewDto) Get() *KubernetesOverviewDto {
	return v.value
}

func (v *NullableKubernetesOverviewDto) Set(val *KubernetesOverviewDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesOverviewDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesOverviewDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesOverviewDto(val *KubernetesOverviewDto) *NullableKubernetesOverviewDto {
	return &NullableKubernetesOverviewDto{value: val, isSet: true}
}

func (v NullableKubernetesOverviewDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesOverviewDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
