# CatalogAppListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogAppId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RepoId** | Pointer to **int32** |  | [optional] 
**RepoName** | Pointer to **NullableString** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**CatalogName** | Pointer to **NullableString** |  | [optional] 
**PackageId** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**LogoImageId** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SecurityReportSummary** | Pointer to [**SecurityReportSummary**](SecurityReportSummary.md) |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Stars** | Pointer to **int64** |  | [optional] 
**InstalledInstanceCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCatalogAppListDto

`func NewCatalogAppListDto() *CatalogAppListDto`

NewCatalogAppListDto instantiates a new CatalogAppListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAppListDtoWithDefaults

`func NewCatalogAppListDtoWithDefaults() *CatalogAppListDto`

NewCatalogAppListDtoWithDefaults instantiates a new CatalogAppListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogAppId

`func (o *CatalogAppListDto) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *CatalogAppListDto) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *CatalogAppListDto) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.

### HasCatalogAppId

`func (o *CatalogAppListDto) HasCatalogAppId() bool`

HasCatalogAppId returns a boolean if a field has been set.

### GetName

`func (o *CatalogAppListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogAppListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogAppListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogAppListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogAppListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogAppListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRepoId

`func (o *CatalogAppListDto) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *CatalogAppListDto) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *CatalogAppListDto) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *CatalogAppListDto) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetRepoName

`func (o *CatalogAppListDto) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *CatalogAppListDto) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *CatalogAppListDto) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *CatalogAppListDto) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### SetRepoNameNil

`func (o *CatalogAppListDto) SetRepoNameNil(b bool)`

 SetRepoNameNil sets the value for RepoName to be an explicit nil

### UnsetRepoName
`func (o *CatalogAppListDto) UnsetRepoName()`

UnsetRepoName ensures that no value is present for RepoName, not even an explicit nil
### GetCatalogId

`func (o *CatalogAppListDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CatalogAppListDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CatalogAppListDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *CatalogAppListDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogName

`func (o *CatalogAppListDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *CatalogAppListDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *CatalogAppListDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *CatalogAppListDto) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### SetCatalogNameNil

`func (o *CatalogAppListDto) SetCatalogNameNil(b bool)`

 SetCatalogNameNil sets the value for CatalogName to be an explicit nil

### UnsetCatalogName
`func (o *CatalogAppListDto) UnsetCatalogName()`

UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
### GetPackageId

`func (o *CatalogAppListDto) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CatalogAppListDto) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CatalogAppListDto) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CatalogAppListDto) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageIdNil

`func (o *CatalogAppListDto) SetPackageIdNil(b bool)`

 SetPackageIdNil sets the value for PackageId to be an explicit nil

### UnsetPackageId
`func (o *CatalogAppListDto) UnsetPackageId()`

UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
### GetVersion

`func (o *CatalogAppListDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogAppListDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogAppListDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogAppListDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CatalogAppListDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CatalogAppListDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetLogoImageId

`func (o *CatalogAppListDto) GetLogoImageId() string`

GetLogoImageId returns the LogoImageId field if non-nil, zero value otherwise.

### GetLogoImageIdOk

`func (o *CatalogAppListDto) GetLogoImageIdOk() (*string, bool)`

GetLogoImageIdOk returns a tuple with the LogoImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImageId

`func (o *CatalogAppListDto) SetLogoImageId(v string)`

SetLogoImageId sets LogoImageId field to given value.

### HasLogoImageId

`func (o *CatalogAppListDto) HasLogoImageId() bool`

HasLogoImageId returns a boolean if a field has been set.

### SetLogoImageIdNil

`func (o *CatalogAppListDto) SetLogoImageIdNil(b bool)`

 SetLogoImageIdNil sets the value for LogoImageId to be an explicit nil

### UnsetLogoImageId
`func (o *CatalogAppListDto) UnsetLogoImageId()`

UnsetLogoImageId ensures that no value is present for LogoImageId, not even an explicit nil
### GetIsLocked

`func (o *CatalogAppListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CatalogAppListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CatalogAppListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *CatalogAppListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetAppVersion

`func (o *CatalogAppListDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *CatalogAppListDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *CatalogAppListDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *CatalogAppListDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *CatalogAppListDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *CatalogAppListDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetDescription

`func (o *CatalogAppListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogAppListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogAppListDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogAppListDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogAppListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogAppListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecurityReportSummary

`func (o *CatalogAppListDto) GetSecurityReportSummary() SecurityReportSummary`

GetSecurityReportSummary returns the SecurityReportSummary field if non-nil, zero value otherwise.

### GetSecurityReportSummaryOk

`func (o *CatalogAppListDto) GetSecurityReportSummaryOk() (*SecurityReportSummary, bool)`

GetSecurityReportSummaryOk returns a tuple with the SecurityReportSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReportSummary

`func (o *CatalogAppListDto) SetSecurityReportSummary(v SecurityReportSummary)`

SetSecurityReportSummary sets SecurityReportSummary field to given value.

### HasSecurityReportSummary

`func (o *CatalogAppListDto) HasSecurityReportSummary() bool`

HasSecurityReportSummary returns a boolean if a field has been set.

### GetRepository

`func (o *CatalogAppListDto) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CatalogAppListDto) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CatalogAppListDto) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CatalogAppListDto) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetStars

`func (o *CatalogAppListDto) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *CatalogAppListDto) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *CatalogAppListDto) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *CatalogAppListDto) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetInstalledInstanceCount

`func (o *CatalogAppListDto) GetInstalledInstanceCount() int32`

GetInstalledInstanceCount returns the InstalledInstanceCount field if non-nil, zero value otherwise.

### GetInstalledInstanceCountOk

`func (o *CatalogAppListDto) GetInstalledInstanceCountOk() (*int32, bool)`

GetInstalledInstanceCountOk returns a tuple with the InstalledInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledInstanceCount

`func (o *CatalogAppListDto) SetInstalledInstanceCount(v int32)`

SetInstalledInstanceCount sets InstalledInstanceCount field to given value.

### HasInstalledInstanceCount

`func (o *CatalogAppListDto) HasInstalledInstanceCount() bool`

HasInstalledInstanceCount returns a boolean if a field has been set.

### SetInstalledInstanceCountNil

`func (o *CatalogAppListDto) SetInstalledInstanceCountNil(b bool)`

 SetInstalledInstanceCountNil sets the value for InstalledInstanceCount to be an explicit nil

### UnsetInstalledInstanceCount
`func (o *CatalogAppListDto) UnsetInstalledInstanceCount()`

UnsetInstalledInstanceCount ensures that no value is present for InstalledInstanceCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


