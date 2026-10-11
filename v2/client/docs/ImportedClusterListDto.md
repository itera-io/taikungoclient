# ImportedClusterListDto

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
**CreatedBy** | **string** |  | 
**LastModified** | **string** |  | 
**LastModifiedBy** | **string** |  | 
**AlertsCount** | **int32** |  | 
**ExpiredAt** | **string** |  | 
**DeleteOnExpiration** | **bool** |  | 
**WasmEnabled** | **bool** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**AlertingProfileName** | **NullableString** |  | 
**AccessIp** | **string** |  | 
**CloudType** | **map[string]interface{}** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 

## Methods

### NewImportedClusterListDto

`func NewImportedClusterListDto(id int32, name string, isVirtualCluster bool, isLocked bool, hasKubeConfigFile bool, isMaintenanceModeEnabled bool, organizationName string, organizationId int32, kubernetesVersion string, createdAt string, createdBy string, lastModified string, lastModifiedBy string, alertsCount int32, expiredAt string, deleteOnExpiration bool, wasmEnabled bool, alertingProfileId NullableInt32, alertingProfileName NullableString, accessIp string, cloudType map[string]interface{}, status ProjectStatus, health ProjectHealth, ) *ImportedClusterListDto`

NewImportedClusterListDto instantiates a new ImportedClusterListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedClusterListDtoWithDefaults

`func NewImportedClusterListDtoWithDefaults() *ImportedClusterListDto`

NewImportedClusterListDtoWithDefaults instantiates a new ImportedClusterListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportedClusterListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportedClusterListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportedClusterListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ImportedClusterListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportedClusterListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportedClusterListDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsVirtualCluster

`func (o *ImportedClusterListDto) GetIsVirtualCluster() bool`

GetIsVirtualCluster returns the IsVirtualCluster field if non-nil, zero value otherwise.

### GetIsVirtualClusterOk

`func (o *ImportedClusterListDto) GetIsVirtualClusterOk() (*bool, bool)`

GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualCluster

`func (o *ImportedClusterListDto) SetIsVirtualCluster(v bool)`

SetIsVirtualCluster sets IsVirtualCluster field to given value.


### GetIsLocked

`func (o *ImportedClusterListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ImportedClusterListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ImportedClusterListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetHasKubeConfigFile

`func (o *ImportedClusterListDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ImportedClusterListDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ImportedClusterListDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetIsMaintenanceModeEnabled

`func (o *ImportedClusterListDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ImportedClusterListDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ImportedClusterListDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.


### GetOrganizationName

`func (o *ImportedClusterListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ImportedClusterListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ImportedClusterListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *ImportedClusterListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ImportedClusterListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ImportedClusterListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetKubernetesVersion

`func (o *ImportedClusterListDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ImportedClusterListDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ImportedClusterListDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetCreatedAt

`func (o *ImportedClusterListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImportedClusterListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImportedClusterListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ImportedClusterListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ImportedClusterListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ImportedClusterListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *ImportedClusterListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ImportedClusterListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ImportedClusterListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### GetLastModifiedBy

`func (o *ImportedClusterListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ImportedClusterListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ImportedClusterListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetAlertsCount

`func (o *ImportedClusterListDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ImportedClusterListDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ImportedClusterListDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetExpiredAt

`func (o *ImportedClusterListDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ImportedClusterListDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ImportedClusterListDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### GetDeleteOnExpiration

`func (o *ImportedClusterListDto) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *ImportedClusterListDto) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *ImportedClusterListDto) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.


### GetWasmEnabled

`func (o *ImportedClusterListDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *ImportedClusterListDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *ImportedClusterListDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.


### GetAlertingProfileId

`func (o *ImportedClusterListDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *ImportedClusterListDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *ImportedClusterListDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *ImportedClusterListDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *ImportedClusterListDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetAlertingProfileName

`func (o *ImportedClusterListDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *ImportedClusterListDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *ImportedClusterListDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.


### SetAlertingProfileNameNil

`func (o *ImportedClusterListDto) SetAlertingProfileNameNil(b bool)`

 SetAlertingProfileNameNil sets the value for AlertingProfileName to be an explicit nil

### UnsetAlertingProfileName
`func (o *ImportedClusterListDto) UnsetAlertingProfileName()`

UnsetAlertingProfileName ensures that no value is present for AlertingProfileName, not even an explicit nil
### GetAccessIp

`func (o *ImportedClusterListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ImportedClusterListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ImportedClusterListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### GetCloudType

`func (o *ImportedClusterListDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ImportedClusterListDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ImportedClusterListDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *ImportedClusterListDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportedClusterListDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportedClusterListDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ImportedClusterListDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ImportedClusterListDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ImportedClusterListDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


