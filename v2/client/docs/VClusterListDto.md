# VClusterListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**IsVirtualCluster** | **bool** |  | 
**IsLocked** | **bool** |  | 
**HasKubeConfigFile** | **bool** |  | 
**IsMaintenanceModeEnabled** | **bool** |  | 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**KubernetesVersion** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **string** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**AlertsCount** | **int32** |  | 
**ExpiredAt** | **string** |  | 
**DeleteOnExpiration** | **bool** |  | 
**WasmEnabled** | **bool** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**AlertingProfileName** | **NullableString** |  | 
**AccessIp** | **string** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**WorkloadResources** | Pointer to [**VClusterWorkloadResourcesListDto**](VClusterWorkloadResourcesListDto.md) |  | [optional] 
**CpuLimits** | Pointer to [**VClusterResourceLimitsListDto**](VClusterResourceLimitsListDto.md) |  | [optional] 
**RamLimits** | Pointer to [**VClusterResourceLimitsListDto**](VClusterResourceLimitsListDto.md) |  | [optional] 
**EphemeralStorageLimits** | Pointer to [**VClusterResourceLimitsListDto**](VClusterResourceLimitsListDto.md) |  | [optional] 
**LockButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**UnlockButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DeleteButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**KubeInfoButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**SetExpirationDateButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**ResetStatusButton** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 

## Methods

### NewVClusterListDto

`func NewVClusterListDto(id int32, name string, isVirtualCluster bool, isLocked bool, hasKubeConfigFile bool, isMaintenanceModeEnabled bool, organizationName string, organizationId int32, kubernetesVersion string, createdAt string, createdBy AuditUserDto, lastModified string, lastModifiedBy AuditUserDto, alertsCount int32, expiredAt string, deleteOnExpiration bool, wasmEnabled bool, alertingProfileId NullableInt32, alertingProfileName NullableString, accessIp string, cloudType ECloudCredentialType, status ProjectStatus, health ProjectHealth, lockButton ButtonStatusDto, unlockButton ButtonStatusDto, deleteButton ButtonStatusDto, kubeInfoButton ButtonStatusDto, setExpirationDateButton ButtonStatusDto, resetStatusButton ButtonStatusDto, ) *VClusterListDto`

NewVClusterListDto instantiates a new VClusterListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVClusterListDtoWithDefaults

`func NewVClusterListDtoWithDefaults() *VClusterListDto`

NewVClusterListDtoWithDefaults instantiates a new VClusterListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VClusterListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VClusterListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VClusterListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *VClusterListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VClusterListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VClusterListDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsVirtualCluster

`func (o *VClusterListDto) GetIsVirtualCluster() bool`

GetIsVirtualCluster returns the IsVirtualCluster field if non-nil, zero value otherwise.

### GetIsVirtualClusterOk

`func (o *VClusterListDto) GetIsVirtualClusterOk() (*bool, bool)`

GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualCluster

`func (o *VClusterListDto) SetIsVirtualCluster(v bool)`

SetIsVirtualCluster sets IsVirtualCluster field to given value.


### GetIsLocked

`func (o *VClusterListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *VClusterListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *VClusterListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetHasKubeConfigFile

`func (o *VClusterListDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *VClusterListDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *VClusterListDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetIsMaintenanceModeEnabled

`func (o *VClusterListDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *VClusterListDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *VClusterListDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.


### GetOrganizationName

`func (o *VClusterListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *VClusterListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *VClusterListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *VClusterListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *VClusterListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *VClusterListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetKubernetesVersion

`func (o *VClusterListDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *VClusterListDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *VClusterListDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetCreatedAt

`func (o *VClusterListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VClusterListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VClusterListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VClusterListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VClusterListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VClusterListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *VClusterListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VClusterListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VClusterListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### GetLastModifiedBy

`func (o *VClusterListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *VClusterListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *VClusterListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetAlertsCount

`func (o *VClusterListDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *VClusterListDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *VClusterListDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetExpiredAt

`func (o *VClusterListDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *VClusterListDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *VClusterListDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### GetDeleteOnExpiration

`func (o *VClusterListDto) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *VClusterListDto) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *VClusterListDto) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.


### GetWasmEnabled

`func (o *VClusterListDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *VClusterListDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *VClusterListDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.


### GetAlertingProfileId

`func (o *VClusterListDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *VClusterListDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *VClusterListDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *VClusterListDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *VClusterListDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetAlertingProfileName

`func (o *VClusterListDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *VClusterListDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *VClusterListDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.


### SetAlertingProfileNameNil

`func (o *VClusterListDto) SetAlertingProfileNameNil(b bool)`

 SetAlertingProfileNameNil sets the value for AlertingProfileName to be an explicit nil

### UnsetAlertingProfileName
`func (o *VClusterListDto) UnsetAlertingProfileName()`

UnsetAlertingProfileName ensures that no value is present for AlertingProfileName, not even an explicit nil
### GetAccessIp

`func (o *VClusterListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *VClusterListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *VClusterListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### GetCloudType

`func (o *VClusterListDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *VClusterListDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *VClusterListDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *VClusterListDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VClusterListDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VClusterListDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *VClusterListDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VClusterListDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VClusterListDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetWorkloadResources

`func (o *VClusterListDto) GetWorkloadResources() VClusterWorkloadResourcesListDto`

GetWorkloadResources returns the WorkloadResources field if non-nil, zero value otherwise.

### GetWorkloadResourcesOk

`func (o *VClusterListDto) GetWorkloadResourcesOk() (*VClusterWorkloadResourcesListDto, bool)`

GetWorkloadResourcesOk returns a tuple with the WorkloadResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadResources

`func (o *VClusterListDto) SetWorkloadResources(v VClusterWorkloadResourcesListDto)`

SetWorkloadResources sets WorkloadResources field to given value.

### HasWorkloadResources

`func (o *VClusterListDto) HasWorkloadResources() bool`

HasWorkloadResources returns a boolean if a field has been set.

### GetCpuLimits

`func (o *VClusterListDto) GetCpuLimits() VClusterResourceLimitsListDto`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *VClusterListDto) GetCpuLimitsOk() (*VClusterResourceLimitsListDto, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *VClusterListDto) SetCpuLimits(v VClusterResourceLimitsListDto)`

SetCpuLimits sets CpuLimits field to given value.

### HasCpuLimits

`func (o *VClusterListDto) HasCpuLimits() bool`

HasCpuLimits returns a boolean if a field has been set.

### GetRamLimits

`func (o *VClusterListDto) GetRamLimits() VClusterResourceLimitsListDto`

GetRamLimits returns the RamLimits field if non-nil, zero value otherwise.

### GetRamLimitsOk

`func (o *VClusterListDto) GetRamLimitsOk() (*VClusterResourceLimitsListDto, bool)`

GetRamLimitsOk returns a tuple with the RamLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamLimits

`func (o *VClusterListDto) SetRamLimits(v VClusterResourceLimitsListDto)`

SetRamLimits sets RamLimits field to given value.

### HasRamLimits

`func (o *VClusterListDto) HasRamLimits() bool`

HasRamLimits returns a boolean if a field has been set.

### GetEphemeralStorageLimits

`func (o *VClusterListDto) GetEphemeralStorageLimits() VClusterResourceLimitsListDto`

GetEphemeralStorageLimits returns the EphemeralStorageLimits field if non-nil, zero value otherwise.

### GetEphemeralStorageLimitsOk

`func (o *VClusterListDto) GetEphemeralStorageLimitsOk() (*VClusterResourceLimitsListDto, bool)`

GetEphemeralStorageLimitsOk returns a tuple with the EphemeralStorageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageLimits

`func (o *VClusterListDto) SetEphemeralStorageLimits(v VClusterResourceLimitsListDto)`

SetEphemeralStorageLimits sets EphemeralStorageLimits field to given value.

### HasEphemeralStorageLimits

`func (o *VClusterListDto) HasEphemeralStorageLimits() bool`

HasEphemeralStorageLimits returns a boolean if a field has been set.

### GetLockButton

`func (o *VClusterListDto) GetLockButton() ButtonStatusDto`

GetLockButton returns the LockButton field if non-nil, zero value otherwise.

### GetLockButtonOk

`func (o *VClusterListDto) GetLockButtonOk() (*ButtonStatusDto, bool)`

GetLockButtonOk returns a tuple with the LockButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockButton

`func (o *VClusterListDto) SetLockButton(v ButtonStatusDto)`

SetLockButton sets LockButton field to given value.


### GetUnlockButton

`func (o *VClusterListDto) GetUnlockButton() ButtonStatusDto`

GetUnlockButton returns the UnlockButton field if non-nil, zero value otherwise.

### GetUnlockButtonOk

`func (o *VClusterListDto) GetUnlockButtonOk() (*ButtonStatusDto, bool)`

GetUnlockButtonOk returns a tuple with the UnlockButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockButton

`func (o *VClusterListDto) SetUnlockButton(v ButtonStatusDto)`

SetUnlockButton sets UnlockButton field to given value.


### GetDeleteButton

`func (o *VClusterListDto) GetDeleteButton() ButtonStatusDto`

GetDeleteButton returns the DeleteButton field if non-nil, zero value otherwise.

### GetDeleteButtonOk

`func (o *VClusterListDto) GetDeleteButtonOk() (*ButtonStatusDto, bool)`

GetDeleteButtonOk returns a tuple with the DeleteButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteButton

`func (o *VClusterListDto) SetDeleteButton(v ButtonStatusDto)`

SetDeleteButton sets DeleteButton field to given value.


### GetKubeInfoButton

`func (o *VClusterListDto) GetKubeInfoButton() ButtonStatusDto`

GetKubeInfoButton returns the KubeInfoButton field if non-nil, zero value otherwise.

### GetKubeInfoButtonOk

`func (o *VClusterListDto) GetKubeInfoButtonOk() (*ButtonStatusDto, bool)`

GetKubeInfoButtonOk returns a tuple with the KubeInfoButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeInfoButton

`func (o *VClusterListDto) SetKubeInfoButton(v ButtonStatusDto)`

SetKubeInfoButton sets KubeInfoButton field to given value.


### GetSetExpirationDateButton

`func (o *VClusterListDto) GetSetExpirationDateButton() ButtonStatusDto`

GetSetExpirationDateButton returns the SetExpirationDateButton field if non-nil, zero value otherwise.

### GetSetExpirationDateButtonOk

`func (o *VClusterListDto) GetSetExpirationDateButtonOk() (*ButtonStatusDto, bool)`

GetSetExpirationDateButtonOk returns a tuple with the SetExpirationDateButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetExpirationDateButton

`func (o *VClusterListDto) SetSetExpirationDateButton(v ButtonStatusDto)`

SetSetExpirationDateButton sets SetExpirationDateButton field to given value.


### GetResetStatusButton

`func (o *VClusterListDto) GetResetStatusButton() ButtonStatusDto`

GetResetStatusButton returns the ResetStatusButton field if non-nil, zero value otherwise.

### GetResetStatusButtonOk

`func (o *VClusterListDto) GetResetStatusButtonOk() (*ButtonStatusDto, bool)`

GetResetStatusButtonOk returns a tuple with the ResetStatusButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetStatusButton

`func (o *VClusterListDto) SetResetStatusButton(v ButtonStatusDto)`

SetResetStatusButton sets ResetStatusButton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


