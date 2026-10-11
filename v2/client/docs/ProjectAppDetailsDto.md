# ProjectAppDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Status** | [**EInstanceStatus**](EInstanceStatus.md) |  | 
**Version** | **string** |  | 
**CatalogId** | **int32** |  | 
**CatalogName** | **string** |  | 
**CatalogAppName** | **string** |  | 
**AppRepoName** | **string** |  | 
**Logo** | **NullableString** |  | 
**Values** | **string** |  | 
**AutoSync** | **bool** |  | 
**ReleaseNotes** | **NullableString** |  | 
**ProjectName** | **string** |  | 
**HelmResult** | **string** |  | 
**ProjectId** | **int32** |  | 
**Ttl** | **int32** |  | 
**HasJsonSchema** | **bool** |  | 
**CatalogAppId** | **int32** |  | 
**PackageId** | **string** |  | 
**Logs** | **NullableString** |  | 
**TaikunLinkUrl** | Pointer to **NullableString** |  | [optional] 
**TaikunLinkEnabled** | Pointer to **bool** |  | [optional] 
**ProjectAppParams** | [**[]ProjectAppParamDto**](ProjectAppParamDto.md) |  | 

## Methods

### NewProjectAppDetailsDto

`func NewProjectAppDetailsDto(id int32, name string, namespace string, status EInstanceStatus, version string, catalogId int32, catalogName string, catalogAppName string, appRepoName string, logo NullableString, values string, autoSync bool, releaseNotes NullableString, projectName string, helmResult string, projectId int32, ttl int32, hasJsonSchema bool, catalogAppId int32, packageId string, logs NullableString, projectAppParams []ProjectAppParamDto, ) *ProjectAppDetailsDto`

NewProjectAppDetailsDto instantiates a new ProjectAppDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppDetailsDtoWithDefaults

`func NewProjectAppDetailsDtoWithDefaults() *ProjectAppDetailsDto`

NewProjectAppDetailsDtoWithDefaults instantiates a new ProjectAppDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectAppDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectAppDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectAppDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectAppDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAppDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAppDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ProjectAppDetailsDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProjectAppDetailsDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProjectAppDetailsDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetStatus

`func (o *ProjectAppDetailsDto) GetStatus() EInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectAppDetailsDto) GetStatusOk() (*EInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectAppDetailsDto) SetStatus(v EInstanceStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *ProjectAppDetailsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectAppDetailsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectAppDetailsDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCatalogId

`func (o *ProjectAppDetailsDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ProjectAppDetailsDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ProjectAppDetailsDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.


### GetCatalogName

`func (o *ProjectAppDetailsDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *ProjectAppDetailsDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *ProjectAppDetailsDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetCatalogAppName

`func (o *ProjectAppDetailsDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *ProjectAppDetailsDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *ProjectAppDetailsDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.


### GetAppRepoName

`func (o *ProjectAppDetailsDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *ProjectAppDetailsDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *ProjectAppDetailsDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.


### GetLogo

`func (o *ProjectAppDetailsDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ProjectAppDetailsDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ProjectAppDetailsDto) SetLogo(v string)`

SetLogo sets Logo field to given value.


### SetLogoNil

`func (o *ProjectAppDetailsDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *ProjectAppDetailsDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetValues

`func (o *ProjectAppDetailsDto) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProjectAppDetailsDto) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProjectAppDetailsDto) SetValues(v string)`

SetValues sets Values field to given value.


### GetAutoSync

`func (o *ProjectAppDetailsDto) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *ProjectAppDetailsDto) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *ProjectAppDetailsDto) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.


### GetReleaseNotes

`func (o *ProjectAppDetailsDto) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ProjectAppDetailsDto) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ProjectAppDetailsDto) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.


### SetReleaseNotesNil

`func (o *ProjectAppDetailsDto) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *ProjectAppDetailsDto) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil
### GetProjectName

`func (o *ProjectAppDetailsDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectAppDetailsDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectAppDetailsDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetHelmResult

`func (o *ProjectAppDetailsDto) GetHelmResult() string`

GetHelmResult returns the HelmResult field if non-nil, zero value otherwise.

### GetHelmResultOk

`func (o *ProjectAppDetailsDto) GetHelmResultOk() (*string, bool)`

GetHelmResultOk returns a tuple with the HelmResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmResult

`func (o *ProjectAppDetailsDto) SetHelmResult(v string)`

SetHelmResult sets HelmResult field to given value.


### GetProjectId

`func (o *ProjectAppDetailsDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectAppDetailsDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectAppDetailsDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetTtl

`func (o *ProjectAppDetailsDto) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ProjectAppDetailsDto) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ProjectAppDetailsDto) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetHasJsonSchema

`func (o *ProjectAppDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *ProjectAppDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *ProjectAppDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.


### GetCatalogAppId

`func (o *ProjectAppDetailsDto) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *ProjectAppDetailsDto) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *ProjectAppDetailsDto) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.


### GetPackageId

`func (o *ProjectAppDetailsDto) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *ProjectAppDetailsDto) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *ProjectAppDetailsDto) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetLogs

`func (o *ProjectAppDetailsDto) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ProjectAppDetailsDto) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ProjectAppDetailsDto) SetLogs(v string)`

SetLogs sets Logs field to given value.


### SetLogsNil

`func (o *ProjectAppDetailsDto) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *ProjectAppDetailsDto) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetTaikunLinkUrl

`func (o *ProjectAppDetailsDto) GetTaikunLinkUrl() string`

GetTaikunLinkUrl returns the TaikunLinkUrl field if non-nil, zero value otherwise.

### GetTaikunLinkUrlOk

`func (o *ProjectAppDetailsDto) GetTaikunLinkUrlOk() (*string, bool)`

GetTaikunLinkUrlOk returns a tuple with the TaikunLinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLinkUrl

`func (o *ProjectAppDetailsDto) SetTaikunLinkUrl(v string)`

SetTaikunLinkUrl sets TaikunLinkUrl field to given value.

### HasTaikunLinkUrl

`func (o *ProjectAppDetailsDto) HasTaikunLinkUrl() bool`

HasTaikunLinkUrl returns a boolean if a field has been set.

### SetTaikunLinkUrlNil

`func (o *ProjectAppDetailsDto) SetTaikunLinkUrlNil(b bool)`

 SetTaikunLinkUrlNil sets the value for TaikunLinkUrl to be an explicit nil

### UnsetTaikunLinkUrl
`func (o *ProjectAppDetailsDto) UnsetTaikunLinkUrl()`

UnsetTaikunLinkUrl ensures that no value is present for TaikunLinkUrl, not even an explicit nil
### GetTaikunLinkEnabled

`func (o *ProjectAppDetailsDto) GetTaikunLinkEnabled() bool`

GetTaikunLinkEnabled returns the TaikunLinkEnabled field if non-nil, zero value otherwise.

### GetTaikunLinkEnabledOk

`func (o *ProjectAppDetailsDto) GetTaikunLinkEnabledOk() (*bool, bool)`

GetTaikunLinkEnabledOk returns a tuple with the TaikunLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLinkEnabled

`func (o *ProjectAppDetailsDto) SetTaikunLinkEnabled(v bool)`

SetTaikunLinkEnabled sets TaikunLinkEnabled field to given value.

### HasTaikunLinkEnabled

`func (o *ProjectAppDetailsDto) HasTaikunLinkEnabled() bool`

HasTaikunLinkEnabled returns a boolean if a field has been set.

### GetProjectAppParams

`func (o *ProjectAppDetailsDto) GetProjectAppParams() []ProjectAppParamDto`

GetProjectAppParams returns the ProjectAppParams field if non-nil, zero value otherwise.

### GetProjectAppParamsOk

`func (o *ProjectAppDetailsDto) GetProjectAppParamsOk() (*[]ProjectAppParamDto, bool)`

GetProjectAppParamsOk returns a tuple with the ProjectAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppParams

`func (o *ProjectAppDetailsDto) SetProjectAppParams(v []ProjectAppParamDto)`

SetProjectAppParams sets ProjectAppParams field to given value.


### SetProjectAppParamsNil

`func (o *ProjectAppDetailsDto) SetProjectAppParamsNil(b bool)`

 SetProjectAppParamsNil sets the value for ProjectAppParams to be an explicit nil

### UnsetProjectAppParams
`func (o *ProjectAppDetailsDto) UnsetProjectAppParams()`

UnsetProjectAppParams ensures that no value is present for ProjectAppParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


