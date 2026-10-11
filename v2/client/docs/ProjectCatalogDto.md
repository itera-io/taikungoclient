# ProjectCatalogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**KubernetesVersion** | **NullableString** |  | 
**IsLocked** | **bool** |  | 
**MaintenanceModeEnabled** | **bool** |  | 
**IsVirtualCluster** | **bool** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 

## Methods

### NewProjectCatalogDto

`func NewProjectCatalogDto(id int32, name NullableString, kubernetesVersion NullableString, isLocked bool, maintenanceModeEnabled bool, isVirtualCluster bool, cloudType ECloudCredentialType, status ProjectStatus, health ProjectHealth, ) *ProjectCatalogDto`

NewProjectCatalogDto instantiates a new ProjectCatalogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCatalogDtoWithDefaults

`func NewProjectCatalogDtoWithDefaults() *ProjectCatalogDto`

NewProjectCatalogDtoWithDefaults instantiates a new ProjectCatalogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCatalogDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCatalogDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCatalogDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectCatalogDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCatalogDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCatalogDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProjectCatalogDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectCatalogDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKubernetesVersion

`func (o *ProjectCatalogDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ProjectCatalogDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ProjectCatalogDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *ProjectCatalogDto) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *ProjectCatalogDto) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetIsLocked

`func (o *ProjectCatalogDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectCatalogDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectCatalogDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetMaintenanceModeEnabled

`func (o *ProjectCatalogDto) GetMaintenanceModeEnabled() bool`

GetMaintenanceModeEnabled returns the MaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetMaintenanceModeEnabledOk

`func (o *ProjectCatalogDto) GetMaintenanceModeEnabledOk() (*bool, bool)`

GetMaintenanceModeEnabledOk returns a tuple with the MaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeEnabled

`func (o *ProjectCatalogDto) SetMaintenanceModeEnabled(v bool)`

SetMaintenanceModeEnabled sets MaintenanceModeEnabled field to given value.


### GetIsVirtualCluster

`func (o *ProjectCatalogDto) GetIsVirtualCluster() bool`

GetIsVirtualCluster returns the IsVirtualCluster field if non-nil, zero value otherwise.

### GetIsVirtualClusterOk

`func (o *ProjectCatalogDto) GetIsVirtualClusterOk() (*bool, bool)`

GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualCluster

`func (o *ProjectCatalogDto) SetIsVirtualCluster(v bool)`

SetIsVirtualCluster sets IsVirtualCluster field to given value.


### GetCloudType

`func (o *ProjectCatalogDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectCatalogDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectCatalogDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *ProjectCatalogDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectCatalogDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectCatalogDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ProjectCatalogDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectCatalogDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectCatalogDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


