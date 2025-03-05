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

// checks if the VClusterListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VClusterListDto{}

// VClusterListDto struct for VClusterListDto
type VClusterListDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	IsVirtualCluster bool `json:"isVirtualCluster"`
	IsLocked bool `json:"isLocked"`
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`
	IsMaintenanceModeEnabled bool `json:"isMaintenanceModeEnabled"`
	OrganizationName string `json:"organizationName"`
	OrganizationId int32 `json:"organizationId"`
	KubernetesVersion string `json:"kubernetesVersion"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	LastModified string `json:"lastModified"`
	LastModifiedBy string `json:"lastModifiedBy"`
	AlertsCount int32 `json:"alertsCount"`
	ExpiredAt string `json:"expiredAt"`
	DeleteOnExpiration bool `json:"deleteOnExpiration"`
	WasmEnabled bool `json:"wasmEnabled"`
	AlertingProfileId NullableInt32 `json:"alertingProfileId"`
	AlertingProfileName NullableString `json:"alertingProfileName"`
	AccessIp string `json:"accessIp"`
	CloudType ECloudCredentialType `json:"cloudType"`
	Status ProjectStatus `json:"status"`
	Health ProjectHealth `json:"health"`
	LockButton ButtonStatusDto `json:"lockButton"`
	UnlockButton ButtonStatusDto `json:"unlockButton"`
	DeleteButton ButtonStatusDto `json:"deleteButton"`
	KubeInfoButton ButtonStatusDto `json:"kubeInfoButton"`
	SetExpirationDateButton ButtonStatusDto `json:"setExpirationDateButton"`
	ResetStatusButton ButtonStatusDto `json:"resetStatusButton"`
}

type _VClusterListDto VClusterListDto

// NewVClusterListDto instantiates a new VClusterListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVClusterListDto(id int32, name string, isVirtualCluster bool, isLocked bool, hasKubeConfigFile bool, isMaintenanceModeEnabled bool, organizationName string, organizationId int32, kubernetesVersion string, createdAt string, createdBy string, lastModified string, lastModifiedBy string, alertsCount int32, expiredAt string, deleteOnExpiration bool, wasmEnabled bool, alertingProfileId NullableInt32, alertingProfileName NullableString, accessIp string, cloudType ECloudCredentialType, status ProjectStatus, health ProjectHealth, lockButton ButtonStatusDto, unlockButton ButtonStatusDto, deleteButton ButtonStatusDto, kubeInfoButton ButtonStatusDto, setExpirationDateButton ButtonStatusDto, resetStatusButton ButtonStatusDto) *VClusterListDto {
	this := VClusterListDto{}
	this.Id = id
	this.Name = name
	this.IsVirtualCluster = isVirtualCluster
	this.IsLocked = isLocked
	this.HasKubeConfigFile = hasKubeConfigFile
	this.IsMaintenanceModeEnabled = isMaintenanceModeEnabled
	this.OrganizationName = organizationName
	this.OrganizationId = organizationId
	this.KubernetesVersion = kubernetesVersion
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.LastModified = lastModified
	this.LastModifiedBy = lastModifiedBy
	this.AlertsCount = alertsCount
	this.ExpiredAt = expiredAt
	this.DeleteOnExpiration = deleteOnExpiration
	this.WasmEnabled = wasmEnabled
	this.AlertingProfileId = alertingProfileId
	this.AlertingProfileName = alertingProfileName
	this.AccessIp = accessIp
	this.CloudType = cloudType
	this.Status = status
	this.Health = health
	this.LockButton = lockButton
	this.UnlockButton = unlockButton
	this.DeleteButton = deleteButton
	this.KubeInfoButton = kubeInfoButton
	this.SetExpirationDateButton = setExpirationDateButton
	this.ResetStatusButton = resetStatusButton
	return &this
}

// NewVClusterListDtoWithDefaults instantiates a new VClusterListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVClusterListDtoWithDefaults() *VClusterListDto {
	this := VClusterListDto{}
	return &this
}

// GetId returns the Id field value
func (o *VClusterListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VClusterListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VClusterListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VClusterListDto) SetName(v string) {
	o.Name = v
}

// GetIsVirtualCluster returns the IsVirtualCluster field value
func (o *VClusterListDto) GetIsVirtualCluster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVirtualCluster
}

// GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetIsVirtualClusterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVirtualCluster, true
}

// SetIsVirtualCluster sets field value
func (o *VClusterListDto) SetIsVirtualCluster(v bool) {
	o.IsVirtualCluster = v
}

// GetIsLocked returns the IsLocked field value
func (o *VClusterListDto) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *VClusterListDto) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetHasKubeConfigFile returns the HasKubeConfigFile field value
func (o *VClusterListDto) GetHasKubeConfigFile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasKubeConfigFile
}

// GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetHasKubeConfigFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasKubeConfigFile, true
}

// SetHasKubeConfigFile sets field value
func (o *VClusterListDto) SetHasKubeConfigFile(v bool) {
	o.HasKubeConfigFile = v
}

// GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field value
func (o *VClusterListDto) GetIsMaintenanceModeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMaintenanceModeEnabled
}

// GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetIsMaintenanceModeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMaintenanceModeEnabled, true
}

// SetIsMaintenanceModeEnabled sets field value
func (o *VClusterListDto) SetIsMaintenanceModeEnabled(v bool) {
	o.IsMaintenanceModeEnabled = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *VClusterListDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *VClusterListDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *VClusterListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *VClusterListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetKubernetesVersion returns the KubernetesVersion field value
func (o *VClusterListDto) GetKubernetesVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesVersion, true
}

// SetKubernetesVersion sets field value
func (o *VClusterListDto) SetKubernetesVersion(v string) {
	o.KubernetesVersion = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VClusterListDto) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VClusterListDto) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *VClusterListDto) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *VClusterListDto) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLastModified returns the LastModified field value
func (o *VClusterListDto) GetLastModified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *VClusterListDto) SetLastModified(v string) {
	o.LastModified = v
}

// GetLastModifiedBy returns the LastModifiedBy field value
func (o *VClusterListDto) GetLastModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value
func (o *VClusterListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy = v
}

// GetAlertsCount returns the AlertsCount field value
func (o *VClusterListDto) GetAlertsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertsCount
}

// GetAlertsCountOk returns a tuple with the AlertsCount field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetAlertsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertsCount, true
}

// SetAlertsCount sets field value
func (o *VClusterListDto) SetAlertsCount(v int32) {
	o.AlertsCount = v
}

// GetExpiredAt returns the ExpiredAt field value
func (o *VClusterListDto) GetExpiredAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetExpiredAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredAt, true
}

// SetExpiredAt sets field value
func (o *VClusterListDto) SetExpiredAt(v string) {
	o.ExpiredAt = v
}

// GetDeleteOnExpiration returns the DeleteOnExpiration field value
func (o *VClusterListDto) GetDeleteOnExpiration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeleteOnExpiration
}

// GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetDeleteOnExpirationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeleteOnExpiration, true
}

// SetDeleteOnExpiration sets field value
func (o *VClusterListDto) SetDeleteOnExpiration(v bool) {
	o.DeleteOnExpiration = v
}

// GetWasmEnabled returns the WasmEnabled field value
func (o *VClusterListDto) GetWasmEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WasmEnabled
}

// GetWasmEnabledOk returns a tuple with the WasmEnabled field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetWasmEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WasmEnabled, true
}

// SetWasmEnabled sets field value
func (o *VClusterListDto) SetWasmEnabled(v bool) {
	o.WasmEnabled = v
}

// GetAlertingProfileId returns the AlertingProfileId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VClusterListDto) GetAlertingProfileId() int32 {
	if o == nil || o.AlertingProfileId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AlertingProfileId.Get()
}

// GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VClusterListDto) GetAlertingProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertingProfileId.Get(), o.AlertingProfileId.IsSet()
}

// SetAlertingProfileId sets field value
func (o *VClusterListDto) SetAlertingProfileId(v int32) {
	o.AlertingProfileId.Set(&v)
}

// GetAlertingProfileName returns the AlertingProfileName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VClusterListDto) GetAlertingProfileName() string {
	if o == nil || o.AlertingProfileName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AlertingProfileName.Get()
}

// GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VClusterListDto) GetAlertingProfileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertingProfileName.Get(), o.AlertingProfileName.IsSet()
}

// SetAlertingProfileName sets field value
func (o *VClusterListDto) SetAlertingProfileName(v string) {
	o.AlertingProfileName.Set(&v)
}

// GetAccessIp returns the AccessIp field value
func (o *VClusterListDto) GetAccessIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessIp
}

// GetAccessIpOk returns a tuple with the AccessIp field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetAccessIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessIp, true
}

// SetAccessIp sets field value
func (o *VClusterListDto) SetAccessIp(v string) {
	o.AccessIp = v
}

// GetCloudType returns the CloudType field value
func (o *VClusterListDto) GetCloudType() ECloudCredentialType {
	if o == nil {
		var ret ECloudCredentialType
		return ret
	}

	return o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetCloudTypeOk() (*ECloudCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudType, true
}

// SetCloudType sets field value
func (o *VClusterListDto) SetCloudType(v ECloudCredentialType) {
	o.CloudType = v
}

// GetStatus returns the Status field value
func (o *VClusterListDto) GetStatus() ProjectStatus {
	if o == nil {
		var ret ProjectStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetStatusOk() (*ProjectStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VClusterListDto) SetStatus(v ProjectStatus) {
	o.Status = v
}

// GetHealth returns the Health field value
func (o *VClusterListDto) GetHealth() ProjectHealth {
	if o == nil {
		var ret ProjectHealth
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetHealthOk() (*ProjectHealth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *VClusterListDto) SetHealth(v ProjectHealth) {
	o.Health = v
}

// GetLockButton returns the LockButton field value
func (o *VClusterListDto) GetLockButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.LockButton
}

// GetLockButtonOk returns a tuple with the LockButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetLockButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockButton, true
}

// SetLockButton sets field value
func (o *VClusterListDto) SetLockButton(v ButtonStatusDto) {
	o.LockButton = v
}

// GetUnlockButton returns the UnlockButton field value
func (o *VClusterListDto) GetUnlockButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.UnlockButton
}

// GetUnlockButtonOk returns a tuple with the UnlockButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetUnlockButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockButton, true
}

// SetUnlockButton sets field value
func (o *VClusterListDto) SetUnlockButton(v ButtonStatusDto) {
	o.UnlockButton = v
}

// GetDeleteButton returns the DeleteButton field value
func (o *VClusterListDto) GetDeleteButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.DeleteButton
}

// GetDeleteButtonOk returns a tuple with the DeleteButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetDeleteButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeleteButton, true
}

// SetDeleteButton sets field value
func (o *VClusterListDto) SetDeleteButton(v ButtonStatusDto) {
	o.DeleteButton = v
}

// GetKubeInfoButton returns the KubeInfoButton field value
func (o *VClusterListDto) GetKubeInfoButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.KubeInfoButton
}

// GetKubeInfoButtonOk returns a tuple with the KubeInfoButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetKubeInfoButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeInfoButton, true
}

// SetKubeInfoButton sets field value
func (o *VClusterListDto) SetKubeInfoButton(v ButtonStatusDto) {
	o.KubeInfoButton = v
}

// GetSetExpirationDateButton returns the SetExpirationDateButton field value
func (o *VClusterListDto) GetSetExpirationDateButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.SetExpirationDateButton
}

// GetSetExpirationDateButtonOk returns a tuple with the SetExpirationDateButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetSetExpirationDateButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetExpirationDateButton, true
}

// SetSetExpirationDateButton sets field value
func (o *VClusterListDto) SetSetExpirationDateButton(v ButtonStatusDto) {
	o.SetExpirationDateButton = v
}

// GetResetStatusButton returns the ResetStatusButton field value
func (o *VClusterListDto) GetResetStatusButton() ButtonStatusDto {
	if o == nil {
		var ret ButtonStatusDto
		return ret
	}

	return o.ResetStatusButton
}

// GetResetStatusButtonOk returns a tuple with the ResetStatusButton field value
// and a boolean to check if the value has been set.
func (o *VClusterListDto) GetResetStatusButtonOk() (*ButtonStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetStatusButton, true
}

// SetResetStatusButton sets field value
func (o *VClusterListDto) SetResetStatusButton(v ButtonStatusDto) {
	o.ResetStatusButton = v
}

func (o VClusterListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VClusterListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["isVirtualCluster"] = o.IsVirtualCluster
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["hasKubeConfigFile"] = o.HasKubeConfigFile
	toSerialize["isMaintenanceModeEnabled"] = o.IsMaintenanceModeEnabled
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["kubernetesVersion"] = o.KubernetesVersion
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["lastModified"] = o.LastModified
	toSerialize["lastModifiedBy"] = o.LastModifiedBy
	toSerialize["alertsCount"] = o.AlertsCount
	toSerialize["expiredAt"] = o.ExpiredAt
	toSerialize["deleteOnExpiration"] = o.DeleteOnExpiration
	toSerialize["wasmEnabled"] = o.WasmEnabled
	toSerialize["alertingProfileId"] = o.AlertingProfileId.Get()
	toSerialize["alertingProfileName"] = o.AlertingProfileName.Get()
	toSerialize["accessIp"] = o.AccessIp
	toSerialize["cloudType"] = o.CloudType
	toSerialize["status"] = o.Status
	toSerialize["health"] = o.Health
	toSerialize["lockButton"] = o.LockButton
	toSerialize["unlockButton"] = o.UnlockButton
	toSerialize["deleteButton"] = o.DeleteButton
	toSerialize["kubeInfoButton"] = o.KubeInfoButton
	toSerialize["setExpirationDateButton"] = o.SetExpirationDateButton
	toSerialize["resetStatusButton"] = o.ResetStatusButton
	return toSerialize, nil
}

func (o *VClusterListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"isVirtualCluster",
		"isLocked",
		"hasKubeConfigFile",
		"isMaintenanceModeEnabled",
		"organizationName",
		"organizationId",
		"kubernetesVersion",
		"createdAt",
		"createdBy",
		"lastModified",
		"lastModifiedBy",
		"alertsCount",
		"expiredAt",
		"deleteOnExpiration",
		"wasmEnabled",
		"alertingProfileId",
		"alertingProfileName",
		"accessIp",
		"cloudType",
		"status",
		"health",
		"lockButton",
		"unlockButton",
		"deleteButton",
		"kubeInfoButton",
		"setExpirationDateButton",
		"resetStatusButton",
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

	varVClusterListDto := _VClusterListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVClusterListDto)

	if err != nil {
		return err
	}

	*o = VClusterListDto(varVClusterListDto)

	return err
}

type NullableVClusterListDto struct {
	value *VClusterListDto
	isSet bool
}

func (v NullableVClusterListDto) Get() *VClusterListDto {
	return v.value
}

func (v *NullableVClusterListDto) Set(val *VClusterListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVClusterListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVClusterListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVClusterListDto(val *VClusterListDto) *NullableVClusterListDto {
	return &NullableVClusterListDto{value: val, isSet: true}
}

func (v NullableVClusterListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVClusterListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


