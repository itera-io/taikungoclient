# InstanceAppListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Status** | [**EInstanceStatus**](EInstanceStatus.md) |  | 
**Version** | **NullableString** |  | 
**CatalogId** | **int32** |  | 
**CatalogName** | **NullableString** |  | 
**CatalogAppName** | **NullableString** |  | 
**CatalogAppId** | **int32** |  | 
**AppRepoName** | **NullableString** |  | 
**Logo** | **NullableString** |  | 
**AutoSync** | **bool** |  | 
**Created** | **NullableString** |  | 
**CreatedBy** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | **NullableString** |  | 
**ProjectId** | **int32** |  | 
**ProjectName** | **NullableString** |  | 
**Logs** | **NullableString** |  | 

## Methods

### NewInstanceAppListDto

`func NewInstanceAppListDto(id int32, name NullableString, namespace NullableString, status EInstanceStatus, version NullableString, catalogId int32, catalogName NullableString, catalogAppName NullableString, catalogAppId int32, appRepoName NullableString, logo NullableString, autoSync bool, created NullableString, createdBy NullableString, lastModified NullableString, lastModifiedBy NullableString, projectId int32, projectName NullableString, logs NullableString, ) *InstanceAppListDto`

NewInstanceAppListDto instantiates a new InstanceAppListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAppListDtoWithDefaults

`func NewInstanceAppListDtoWithDefaults() *InstanceAppListDto`

NewInstanceAppListDtoWithDefaults instantiates a new InstanceAppListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceAppListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceAppListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceAppListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InstanceAppListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceAppListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceAppListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *InstanceAppListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceAppListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *InstanceAppListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *InstanceAppListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *InstanceAppListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *InstanceAppListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *InstanceAppListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStatus

`func (o *InstanceAppListDto) GetStatus() EInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceAppListDto) GetStatusOk() (*EInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceAppListDto) SetStatus(v EInstanceStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *InstanceAppListDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceAppListDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceAppListDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *InstanceAppListDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *InstanceAppListDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCatalogId

`func (o *InstanceAppListDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *InstanceAppListDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *InstanceAppListDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.


### GetCatalogName

`func (o *InstanceAppListDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *InstanceAppListDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *InstanceAppListDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### SetCatalogNameNil

`func (o *InstanceAppListDto) SetCatalogNameNil(b bool)`

 SetCatalogNameNil sets the value for CatalogName to be an explicit nil

### UnsetCatalogName
`func (o *InstanceAppListDto) UnsetCatalogName()`

UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
### GetCatalogAppName

`func (o *InstanceAppListDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *InstanceAppListDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *InstanceAppListDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.


### SetCatalogAppNameNil

`func (o *InstanceAppListDto) SetCatalogAppNameNil(b bool)`

 SetCatalogAppNameNil sets the value for CatalogAppName to be an explicit nil

### UnsetCatalogAppName
`func (o *InstanceAppListDto) UnsetCatalogAppName()`

UnsetCatalogAppName ensures that no value is present for CatalogAppName, not even an explicit nil
### GetCatalogAppId

`func (o *InstanceAppListDto) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *InstanceAppListDto) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *InstanceAppListDto) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.


### GetAppRepoName

`func (o *InstanceAppListDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *InstanceAppListDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *InstanceAppListDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.


### SetAppRepoNameNil

`func (o *InstanceAppListDto) SetAppRepoNameNil(b bool)`

 SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil

### UnsetAppRepoName
`func (o *InstanceAppListDto) UnsetAppRepoName()`

UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
### GetLogo

`func (o *InstanceAppListDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InstanceAppListDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InstanceAppListDto) SetLogo(v string)`

SetLogo sets Logo field to given value.


### SetLogoNil

`func (o *InstanceAppListDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *InstanceAppListDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetAutoSync

`func (o *InstanceAppListDto) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *InstanceAppListDto) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *InstanceAppListDto) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.


### GetCreated

`func (o *InstanceAppListDto) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InstanceAppListDto) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InstanceAppListDto) SetCreated(v string)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *InstanceAppListDto) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InstanceAppListDto) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetCreatedBy

`func (o *InstanceAppListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InstanceAppListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InstanceAppListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *InstanceAppListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InstanceAppListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *InstanceAppListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *InstanceAppListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *InstanceAppListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *InstanceAppListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *InstanceAppListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *InstanceAppListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *InstanceAppListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *InstanceAppListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### SetLastModifiedByNil

`func (o *InstanceAppListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *InstanceAppListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetProjectId

`func (o *InstanceAppListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InstanceAppListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InstanceAppListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *InstanceAppListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InstanceAppListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InstanceAppListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *InstanceAppListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *InstanceAppListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetLogs

`func (o *InstanceAppListDto) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *InstanceAppListDto) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *InstanceAppListDto) SetLogs(v string)`

SetLogs sets Logs field to given value.


### SetLogsNil

`func (o *InstanceAppListDto) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *InstanceAppListDto) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


