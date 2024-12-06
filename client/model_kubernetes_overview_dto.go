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
	"bytes"
	"fmt"
)

// checks if the KubernetesOverviewDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesOverviewDto{}

// KubernetesOverviewDto struct for KubernetesOverviewDto
type KubernetesOverviewDto struct {
	ProjectId int32 `json:"projectId"`
	ProjectName string `json:"projectName"`
	HealthyPods int32 `json:"healthyPods"`
	UnhealthyPods int32 `json:"unhealthyPods"`
	HealthyNodes int32 `json:"healthyNodes"`
	UnhealthyNodes int32 `json:"unhealthyNodes"`
	AlertsCount int32 `json:"alertsCount"`
	KubernetesHealth NullableString `json:"kubernetesHealth"`
}

type _KubernetesOverviewDto KubernetesOverviewDto

// NewKubernetesOverviewDto instantiates a new KubernetesOverviewDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesOverviewDto(projectId int32, projectName string, healthyPods int32, unhealthyPods int32, healthyNodes int32, unhealthyNodes int32, alertsCount int32, kubernetesHealth NullableString) *KubernetesOverviewDto {
	this := KubernetesOverviewDto{}
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.HealthyPods = healthyPods
	this.UnhealthyPods = unhealthyPods
	this.HealthyNodes = healthyNodes
	this.UnhealthyNodes = unhealthyNodes
	this.AlertsCount = alertsCount
	this.KubernetesHealth = kubernetesHealth
	return &this
}

// NewKubernetesOverviewDtoWithDefaults instantiates a new KubernetesOverviewDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesOverviewDtoWithDefaults() *KubernetesOverviewDto {
	this := KubernetesOverviewDto{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *KubernetesOverviewDto) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *KubernetesOverviewDto) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
func (o *KubernetesOverviewDto) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *KubernetesOverviewDto) SetProjectName(v string) {
	o.ProjectName = v
}

// GetHealthyPods returns the HealthyPods field value
func (o *KubernetesOverviewDto) GetHealthyPods() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HealthyPods
}

// GetHealthyPodsOk returns a tuple with the HealthyPods field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetHealthyPodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthyPods, true
}

// SetHealthyPods sets field value
func (o *KubernetesOverviewDto) SetHealthyPods(v int32) {
	o.HealthyPods = v
}

// GetUnhealthyPods returns the UnhealthyPods field value
func (o *KubernetesOverviewDto) GetUnhealthyPods() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnhealthyPods
}

// GetUnhealthyPodsOk returns a tuple with the UnhealthyPods field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetUnhealthyPodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnhealthyPods, true
}

// SetUnhealthyPods sets field value
func (o *KubernetesOverviewDto) SetUnhealthyPods(v int32) {
	o.UnhealthyPods = v
}

// GetHealthyNodes returns the HealthyNodes field value
func (o *KubernetesOverviewDto) GetHealthyNodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HealthyNodes
}

// GetHealthyNodesOk returns a tuple with the HealthyNodes field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetHealthyNodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthyNodes, true
}

// SetHealthyNodes sets field value
func (o *KubernetesOverviewDto) SetHealthyNodes(v int32) {
	o.HealthyNodes = v
}

// GetUnhealthyNodes returns the UnhealthyNodes field value
func (o *KubernetesOverviewDto) GetUnhealthyNodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnhealthyNodes
}

// GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetUnhealthyNodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnhealthyNodes, true
}

// SetUnhealthyNodes sets field value
func (o *KubernetesOverviewDto) SetUnhealthyNodes(v int32) {
	o.UnhealthyNodes = v
}

// GetAlertsCount returns the AlertsCount field value
func (o *KubernetesOverviewDto) GetAlertsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertsCount
}

// GetAlertsCountOk returns a tuple with the AlertsCount field value
// and a boolean to check if the value has been set.
func (o *KubernetesOverviewDto) GetAlertsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertsCount, true
}

// SetAlertsCount sets field value
func (o *KubernetesOverviewDto) SetAlertsCount(v int32) {
	o.AlertsCount = v
}

// GetKubernetesHealth returns the KubernetesHealth field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesOverviewDto) GetKubernetesHealth() string {
	if o == nil || o.KubernetesHealth.Get() == nil {
		var ret string
		return ret
	}

	return *o.KubernetesHealth.Get()
}

// GetKubernetesHealthOk returns a tuple with the KubernetesHealth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesOverviewDto) GetKubernetesHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesHealth.Get(), o.KubernetesHealth.IsSet()
}

// SetKubernetesHealth sets field value
func (o *KubernetesOverviewDto) SetKubernetesHealth(v string) {
	o.KubernetesHealth.Set(&v)
}

func (o KubernetesOverviewDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesOverviewDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["projectName"] = o.ProjectName
	toSerialize["healthyPods"] = o.HealthyPods
	toSerialize["unhealthyPods"] = o.UnhealthyPods
	toSerialize["healthyNodes"] = o.HealthyNodes
	toSerialize["unhealthyNodes"] = o.UnhealthyNodes
	toSerialize["alertsCount"] = o.AlertsCount
	toSerialize["kubernetesHealth"] = o.KubernetesHealth.Get()
	return toSerialize, nil
}

func (o *KubernetesOverviewDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
		"projectName",
		"healthyPods",
		"unhealthyPods",
		"healthyNodes",
		"unhealthyNodes",
		"alertsCount",
		"kubernetesHealth",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubernetesOverviewDto := _KubernetesOverviewDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesOverviewDto)

	if err != nil {
		return err
	}

	*o = KubernetesOverviewDto(varKubernetesOverviewDto)

	return err
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


