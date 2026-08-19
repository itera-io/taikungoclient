# CommonDropdownIsBoundDtoForProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**IsBound** | **bool** |  | 
**HasKubeConfigFile** | **bool** |  | 
**KubernetesVersion** | **NullableString** |  | 
**IsLocked** | **bool** |  | 
**MaintenanceModeEnabled** | **bool** |  | 
**IsVirtualCluster** | **bool** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**ImportClusterType** | [**ImportClusterType**](ImportClusterType.md) |  | 

## Methods

### NewCommonDropdownIsBoundDtoForProject

`func NewCommonDropdownIsBoundDtoForProject(id int32, name NullableString, isBound bool, hasKubeConfigFile bool, kubernetesVersion NullableString, isLocked bool, maintenanceModeEnabled bool, isVirtualCluster bool, alertingProfileId NullableInt32, cloudType ECloudCredentialType, status ProjectStatus, health ProjectHealth, importClusterType ImportClusterType, ) *CommonDropdownIsBoundDtoForProject`

NewCommonDropdownIsBoundDtoForProject instantiates a new CommonDropdownIsBoundDtoForProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDropdownIsBoundDtoForProjectWithDefaults

`func NewCommonDropdownIsBoundDtoForProjectWithDefaults() *CommonDropdownIsBoundDtoForProject`

NewCommonDropdownIsBoundDtoForProjectWithDefaults instantiates a new CommonDropdownIsBoundDtoForProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonDropdownIsBoundDtoForProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonDropdownIsBoundDtoForProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonDropdownIsBoundDtoForProject) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CommonDropdownIsBoundDtoForProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonDropdownIsBoundDtoForProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonDropdownIsBoundDtoForProject) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonDropdownIsBoundDtoForProject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonDropdownIsBoundDtoForProject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsBound

`func (o *CommonDropdownIsBoundDtoForProject) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *CommonDropdownIsBoundDtoForProject) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *CommonDropdownIsBoundDtoForProject) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.


### GetHasKubeConfigFile

`func (o *CommonDropdownIsBoundDtoForProject) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *CommonDropdownIsBoundDtoForProject) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *CommonDropdownIsBoundDtoForProject) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetKubernetesVersion

`func (o *CommonDropdownIsBoundDtoForProject) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *CommonDropdownIsBoundDtoForProject) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *CommonDropdownIsBoundDtoForProject) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *CommonDropdownIsBoundDtoForProject) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *CommonDropdownIsBoundDtoForProject) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetIsLocked

`func (o *CommonDropdownIsBoundDtoForProject) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CommonDropdownIsBoundDtoForProject) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CommonDropdownIsBoundDtoForProject) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetMaintenanceModeEnabled

`func (o *CommonDropdownIsBoundDtoForProject) GetMaintenanceModeEnabled() bool`

GetMaintenanceModeEnabled returns the MaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetMaintenanceModeEnabledOk

`func (o *CommonDropdownIsBoundDtoForProject) GetMaintenanceModeEnabledOk() (*bool, bool)`

GetMaintenanceModeEnabledOk returns a tuple with the MaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeEnabled

`func (o *CommonDropdownIsBoundDtoForProject) SetMaintenanceModeEnabled(v bool)`

SetMaintenanceModeEnabled sets MaintenanceModeEnabled field to given value.


### GetIsVirtualCluster

`func (o *CommonDropdownIsBoundDtoForProject) GetIsVirtualCluster() bool`

GetIsVirtualCluster returns the IsVirtualCluster field if non-nil, zero value otherwise.

### GetIsVirtualClusterOk

`func (o *CommonDropdownIsBoundDtoForProject) GetIsVirtualClusterOk() (*bool, bool)`

GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualCluster

`func (o *CommonDropdownIsBoundDtoForProject) SetIsVirtualCluster(v bool)`

SetIsVirtualCluster sets IsVirtualCluster field to given value.


### GetAlertingProfileId

`func (o *CommonDropdownIsBoundDtoForProject) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *CommonDropdownIsBoundDtoForProject) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *CommonDropdownIsBoundDtoForProject) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *CommonDropdownIsBoundDtoForProject) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *CommonDropdownIsBoundDtoForProject) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetCloudType

`func (o *CommonDropdownIsBoundDtoForProject) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CommonDropdownIsBoundDtoForProject) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CommonDropdownIsBoundDtoForProject) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *CommonDropdownIsBoundDtoForProject) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonDropdownIsBoundDtoForProject) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonDropdownIsBoundDtoForProject) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *CommonDropdownIsBoundDtoForProject) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CommonDropdownIsBoundDtoForProject) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CommonDropdownIsBoundDtoForProject) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetImportClusterType

`func (o *CommonDropdownIsBoundDtoForProject) GetImportClusterType() ImportClusterType`

GetImportClusterType returns the ImportClusterType field if non-nil, zero value otherwise.

### GetImportClusterTypeOk

`func (o *CommonDropdownIsBoundDtoForProject) GetImportClusterTypeOk() (*ImportClusterType, bool)`

GetImportClusterTypeOk returns a tuple with the ImportClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportClusterType

`func (o *CommonDropdownIsBoundDtoForProject) SetImportClusterType(v ImportClusterType)`

SetImportClusterType sets ImportClusterType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


