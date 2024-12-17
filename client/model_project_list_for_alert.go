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

// checks if the ProjectListForAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListForAlert{}

// ProjectListForAlert struct for ProjectListForAlert
type ProjectListForAlert struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	Token NullableString `json:"token"`
	Status NullableString `json:"status"`
	OrganizationId int32 `json:"organizationId"`
	Health NullableString `json:"health"`
	IsKubernetes bool `json:"isKubernetes"`
	IsLocked bool `json:"isLocked"`
	IsMonitoringEnabled bool `json:"isMonitoringEnabled"`
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`
	MonitoringCredential MonitoringCredentialsListDto `json:"monitoringCredential"`
	KubernetesAlerts []KubernetesAlertDtoForPoller `json:"kubernetesAlerts"`
	KubernetesCurrentVersion NullableString `json:"kubernetesCurrentVersion"`
	TotalServersCount int32 `json:"totalServersCount"`
}

type _ProjectListForAlert ProjectListForAlert

// NewProjectListForAlert instantiates a new ProjectListForAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListForAlert(id int32, name NullableString, token NullableString, status NullableString, organizationId int32, health NullableString, isKubernetes bool, isLocked bool, isMonitoringEnabled bool, hasKubeConfigFile bool, monitoringCredential MonitoringCredentialsListDto, kubernetesAlerts []KubernetesAlertDtoForPoller, kubernetesCurrentVersion NullableString, totalServersCount int32) *ProjectListForAlert {
	this := ProjectListForAlert{}
	this.Id = id
	this.Name = name
	this.Token = token
	this.Status = status
	this.OrganizationId = organizationId
	this.Health = health
	this.IsKubernetes = isKubernetes
	this.IsLocked = isLocked
	this.IsMonitoringEnabled = isMonitoringEnabled
	this.HasKubeConfigFile = hasKubeConfigFile
	this.MonitoringCredential = monitoringCredential
	this.KubernetesAlerts = kubernetesAlerts
	this.KubernetesCurrentVersion = kubernetesCurrentVersion
	this.TotalServersCount = totalServersCount
	return &this
}

// NewProjectListForAlertWithDefaults instantiates a new ProjectListForAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListForAlertWithDefaults() *ProjectListForAlert {
	this := ProjectListForAlert{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectListForAlert) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectListForAlert) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectListForAlert) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ProjectListForAlert) SetName(v string) {
	o.Name.Set(&v)
}

// GetToken returns the Token field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectListForAlert) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}

	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// SetToken sets field value
func (o *ProjectListForAlert) SetToken(v string) {
	o.Token.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectListForAlert) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *ProjectListForAlert) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetOrganizationId returns the OrganizationId field value
func (o *ProjectListForAlert) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ProjectListForAlert) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetHealth returns the Health field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectListForAlert) GetHealth() string {
	if o == nil || o.Health.Get() == nil {
		var ret string
		return ret
	}

	return *o.Health.Get()
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Health.Get(), o.Health.IsSet()
}

// SetHealth sets field value
func (o *ProjectListForAlert) SetHealth(v string) {
	o.Health.Set(&v)
}

// GetIsKubernetes returns the IsKubernetes field value
func (o *ProjectListForAlert) GetIsKubernetes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKubernetes
}

// GetIsKubernetesOk returns a tuple with the IsKubernetes field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsKubernetesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKubernetes, true
}

// SetIsKubernetes sets field value
func (o *ProjectListForAlert) SetIsKubernetes(v bool) {
	o.IsKubernetes = v
}

// GetIsLocked returns the IsLocked field value
func (o *ProjectListForAlert) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *ProjectListForAlert) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetIsMonitoringEnabled returns the IsMonitoringEnabled field value
func (o *ProjectListForAlert) GetIsMonitoringEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMonitoringEnabled
}

// GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsMonitoringEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMonitoringEnabled, true
}

// SetIsMonitoringEnabled sets field value
func (o *ProjectListForAlert) SetIsMonitoringEnabled(v bool) {
	o.IsMonitoringEnabled = v
}

// GetHasKubeConfigFile returns the HasKubeConfigFile field value
func (o *ProjectListForAlert) GetHasKubeConfigFile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasKubeConfigFile
}

// GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetHasKubeConfigFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasKubeConfigFile, true
}

// SetHasKubeConfigFile sets field value
func (o *ProjectListForAlert) SetHasKubeConfigFile(v bool) {
	o.HasKubeConfigFile = v
}

// GetMonitoringCredential returns the MonitoringCredential field value
func (o *ProjectListForAlert) GetMonitoringCredential() MonitoringCredentialsListDto {
	if o == nil {
		var ret MonitoringCredentialsListDto
		return ret
	}

	return o.MonitoringCredential
}

// GetMonitoringCredentialOk returns a tuple with the MonitoringCredential field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetMonitoringCredentialOk() (*MonitoringCredentialsListDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringCredential, true
}

// SetMonitoringCredential sets field value
func (o *ProjectListForAlert) SetMonitoringCredential(v MonitoringCredentialsListDto) {
	o.MonitoringCredential = v
}

// GetKubernetesAlerts returns the KubernetesAlerts field value
// If the value is explicit nil, the zero value for []KubernetesAlertDtoForPoller will be returned
func (o *ProjectListForAlert) GetKubernetesAlerts() []KubernetesAlertDtoForPoller {
	if o == nil {
		var ret []KubernetesAlertDtoForPoller
		return ret
	}

	return o.KubernetesAlerts
}

// GetKubernetesAlertsOk returns a tuple with the KubernetesAlerts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetKubernetesAlertsOk() ([]KubernetesAlertDtoForPoller, bool) {
	if o == nil || IsNil(o.KubernetesAlerts) {
		return nil, false
	}
	return o.KubernetesAlerts, true
}

// SetKubernetesAlerts sets field value
func (o *ProjectListForAlert) SetKubernetesAlerts(v []KubernetesAlertDtoForPoller) {
	o.KubernetesAlerts = v
}

// GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectListForAlert) GetKubernetesCurrentVersion() string {
	if o == nil || o.KubernetesCurrentVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.KubernetesCurrentVersion.Get()
}

// GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetKubernetesCurrentVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesCurrentVersion.Get(), o.KubernetesCurrentVersion.IsSet()
}

// SetKubernetesCurrentVersion sets field value
func (o *ProjectListForAlert) SetKubernetesCurrentVersion(v string) {
	o.KubernetesCurrentVersion.Set(&v)
}

// GetTotalServersCount returns the TotalServersCount field value
func (o *ProjectListForAlert) GetTotalServersCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalServersCount
}

// GetTotalServersCountOk returns a tuple with the TotalServersCount field value
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetTotalServersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalServersCount, true
}

// SetTotalServersCount sets field value
func (o *ProjectListForAlert) SetTotalServersCount(v int32) {
	o.TotalServersCount = v
}

func (o ProjectListForAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListForAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["token"] = o.Token.Get()
	toSerialize["status"] = o.Status.Get()
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["health"] = o.Health.Get()
	toSerialize["isKubernetes"] = o.IsKubernetes
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["isMonitoringEnabled"] = o.IsMonitoringEnabled
	toSerialize["hasKubeConfigFile"] = o.HasKubeConfigFile
	toSerialize["monitoringCredential"] = o.MonitoringCredential
	if o.KubernetesAlerts != nil {
		toSerialize["kubernetesAlerts"] = o.KubernetesAlerts
	}
	toSerialize["kubernetesCurrentVersion"] = o.KubernetesCurrentVersion.Get()
	toSerialize["totalServersCount"] = o.TotalServersCount
	return toSerialize, nil
}

func (o *ProjectListForAlert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"token",
		"status",
		"organizationId",
		"health",
		"isKubernetes",
		"isLocked",
		"isMonitoringEnabled",
		"hasKubeConfigFile",
		"monitoringCredential",
		"kubernetesAlerts",
		"kubernetesCurrentVersion",
		"totalServersCount",
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

	varProjectListForAlert := _ProjectListForAlert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectListForAlert)

	if err != nil {
		return err
	}

	*o = ProjectListForAlert(varProjectListForAlert)

	return err
}

type NullableProjectListForAlert struct {
	value *ProjectListForAlert
	isSet bool
}

func (v NullableProjectListForAlert) Get() *ProjectListForAlert {
	return v.value
}

func (v *NullableProjectListForAlert) Set(val *ProjectListForAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListForAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListForAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListForAlert(val *ProjectListForAlert) *NullableProjectListForAlert {
	return &NullableProjectListForAlert{value: val, isSet: true}
}

func (v NullableProjectListForAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListForAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


