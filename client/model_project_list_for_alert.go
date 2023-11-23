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

// checks if the ProjectListForAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListForAlert{}

// ProjectListForAlert struct for ProjectListForAlert
type ProjectListForAlert struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Token NullableString `json:"token,omitempty"`
	Status NullableString `json:"status,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	Health NullableString `json:"health,omitempty"`
	IsKubernetes *bool `json:"isKubernetes,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	IsMonitoringEnabled *bool `json:"isMonitoringEnabled,omitempty"`
	HasKubeConfigFile *bool `json:"hasKubeConfigFile,omitempty"`
	MonitoringCredential *MonitoringCredentialsListDto `json:"monitoringCredential,omitempty"`
	KubernetesAlerts []KubernetesAlertDto `json:"kubernetesAlerts,omitempty"`
}

// NewProjectListForAlert instantiates a new ProjectListForAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListForAlert() *ProjectListForAlert {
	this := ProjectListForAlert{}
	return &this
}

// NewProjectListForAlertWithDefaults instantiates a new ProjectListForAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListForAlertWithDefaults() *ProjectListForAlert {
	this := ProjectListForAlert{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectListForAlert) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForAlert) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectListForAlert) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectListForAlert) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectListForAlert) UnsetName() {
	o.Name.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForAlert) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ProjectListForAlert) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ProjectListForAlert) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ProjectListForAlert) UnsetToken() {
	o.Token.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForAlert) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ProjectListForAlert) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ProjectListForAlert) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ProjectListForAlert) UnsetStatus() {
	o.Status.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *ProjectListForAlert) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForAlert) GetHealth() string {
	if o == nil || IsNil(o.Health.Get()) {
		var ret string
		return ret
	}
	return *o.Health.Get()
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Health.Get(), o.Health.IsSet()
}

// HasHealth returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasHealth() bool {
	if o != nil && o.Health.IsSet() {
		return true
	}

	return false
}

// SetHealth gets a reference to the given NullableString and assigns it to the Health field.
func (o *ProjectListForAlert) SetHealth(v string) {
	o.Health.Set(&v)
}
// SetHealthNil sets the value for Health to be an explicit nil
func (o *ProjectListForAlert) SetHealthNil() {
	o.Health.Set(nil)
}

// UnsetHealth ensures that no value is present for Health, not even an explicit nil
func (o *ProjectListForAlert) UnsetHealth() {
	o.Health.Unset()
}

// GetIsKubernetes returns the IsKubernetes field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetIsKubernetes() bool {
	if o == nil || IsNil(o.IsKubernetes) {
		var ret bool
		return ret
	}
	return *o.IsKubernetes
}

// GetIsKubernetesOk returns a tuple with the IsKubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsKubernetesOk() (*bool, bool) {
	if o == nil || IsNil(o.IsKubernetes) {
		return nil, false
	}
	return o.IsKubernetes, true
}

// HasIsKubernetes returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasIsKubernetes() bool {
	if o != nil && !IsNil(o.IsKubernetes) {
		return true
	}

	return false
}

// SetIsKubernetes gets a reference to the given bool and assigns it to the IsKubernetes field.
func (o *ProjectListForAlert) SetIsKubernetes(v bool) {
	o.IsKubernetes = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *ProjectListForAlert) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsMonitoringEnabled returns the IsMonitoringEnabled field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetIsMonitoringEnabled() bool {
	if o == nil || IsNil(o.IsMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMonitoringEnabled
}

// GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetIsMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMonitoringEnabled) {
		return nil, false
	}
	return o.IsMonitoringEnabled, true
}

// HasIsMonitoringEnabled returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasIsMonitoringEnabled() bool {
	if o != nil && !IsNil(o.IsMonitoringEnabled) {
		return true
	}

	return false
}

// SetIsMonitoringEnabled gets a reference to the given bool and assigns it to the IsMonitoringEnabled field.
func (o *ProjectListForAlert) SetIsMonitoringEnabled(v bool) {
	o.IsMonitoringEnabled = &v
}

// GetHasKubeConfigFile returns the HasKubeConfigFile field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetHasKubeConfigFile() bool {
	if o == nil || IsNil(o.HasKubeConfigFile) {
		var ret bool
		return ret
	}
	return *o.HasKubeConfigFile
}

// GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetHasKubeConfigFileOk() (*bool, bool) {
	if o == nil || IsNil(o.HasKubeConfigFile) {
		return nil, false
	}
	return o.HasKubeConfigFile, true
}

// HasHasKubeConfigFile returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasHasKubeConfigFile() bool {
	if o != nil && !IsNil(o.HasKubeConfigFile) {
		return true
	}

	return false
}

// SetHasKubeConfigFile gets a reference to the given bool and assigns it to the HasKubeConfigFile field.
func (o *ProjectListForAlert) SetHasKubeConfigFile(v bool) {
	o.HasKubeConfigFile = &v
}

// GetMonitoringCredential returns the MonitoringCredential field value if set, zero value otherwise.
func (o *ProjectListForAlert) GetMonitoringCredential() MonitoringCredentialsListDto {
	if o == nil || IsNil(o.MonitoringCredential) {
		var ret MonitoringCredentialsListDto
		return ret
	}
	return *o.MonitoringCredential
}

// GetMonitoringCredentialOk returns a tuple with the MonitoringCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListForAlert) GetMonitoringCredentialOk() (*MonitoringCredentialsListDto, bool) {
	if o == nil || IsNil(o.MonitoringCredential) {
		return nil, false
	}
	return o.MonitoringCredential, true
}

// HasMonitoringCredential returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasMonitoringCredential() bool {
	if o != nil && !IsNil(o.MonitoringCredential) {
		return true
	}

	return false
}

// SetMonitoringCredential gets a reference to the given MonitoringCredentialsListDto and assigns it to the MonitoringCredential field.
func (o *ProjectListForAlert) SetMonitoringCredential(v MonitoringCredentialsListDto) {
	o.MonitoringCredential = &v
}

// GetKubernetesAlerts returns the KubernetesAlerts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListForAlert) GetKubernetesAlerts() []KubernetesAlertDto {
	if o == nil {
		var ret []KubernetesAlertDto
		return ret
	}
	return o.KubernetesAlerts
}

// GetKubernetesAlertsOk returns a tuple with the KubernetesAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListForAlert) GetKubernetesAlertsOk() ([]KubernetesAlertDto, bool) {
	if o == nil || IsNil(o.KubernetesAlerts) {
		return nil, false
	}
	return o.KubernetesAlerts, true
}

// HasKubernetesAlerts returns a boolean if a field has been set.
func (o *ProjectListForAlert) HasKubernetesAlerts() bool {
	if o != nil && IsNil(o.KubernetesAlerts) {
		return true
	}

	return false
}

// SetKubernetesAlerts gets a reference to the given []KubernetesAlertDto and assigns it to the KubernetesAlerts field.
func (o *ProjectListForAlert) SetKubernetesAlerts(v []KubernetesAlertDto) {
	o.KubernetesAlerts = v
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.Health.IsSet() {
		toSerialize["health"] = o.Health.Get()
	}
	if !IsNil(o.IsKubernetes) {
		toSerialize["isKubernetes"] = o.IsKubernetes
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsMonitoringEnabled) {
		toSerialize["isMonitoringEnabled"] = o.IsMonitoringEnabled
	}
	if !IsNil(o.HasKubeConfigFile) {
		toSerialize["hasKubeConfigFile"] = o.HasKubeConfigFile
	}
	if !IsNil(o.MonitoringCredential) {
		toSerialize["monitoringCredential"] = o.MonitoringCredential
	}
	if o.KubernetesAlerts != nil {
		toSerialize["kubernetesAlerts"] = o.KubernetesAlerts
	}
	return toSerialize, nil
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


