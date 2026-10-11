# CatalogAppDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AppRepoName** | Pointer to **NullableString** |  | [optional] 
**AppRepoOrganizationName** | Pointer to **NullableString** |  | [optional] 
**AppRepoId** | Pointer to **int32** |  | [optional] 
**CatalogName** | Pointer to **NullableString** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**PackageId** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**LogoId** | Pointer to **NullableString** |  | [optional] 
**ProjectApps** | Pointer to [**[]ProjectAppDto**](ProjectAppDto.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Readme** | Pointer to **NullableString** |  | [optional] 
**SecurityReport** | Pointer to [**SecurityReportSummaryDto**](SecurityReportSummaryDto.md) |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**Stars** | Pointer to **int32** |  | [optional] 
**VerifiedPublisher** | Pointer to **bool** |  | [optional] 
**Official** | Pointer to **bool** |  | [optional] 
**HasJsonSchema** | Pointer to **bool** |  | [optional] 

## Methods

### NewCatalogAppDetailsDto

`func NewCatalogAppDetailsDto() *CatalogAppDetailsDto`

NewCatalogAppDetailsDto instantiates a new CatalogAppDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAppDetailsDtoWithDefaults

`func NewCatalogAppDetailsDtoWithDefaults() *CatalogAppDetailsDto`

NewCatalogAppDetailsDtoWithDefaults instantiates a new CatalogAppDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogAppDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogAppDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogAppDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogAppDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CatalogAppDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogAppDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogAppDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogAppDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogAppDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogAppDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAppRepoName

`func (o *CatalogAppDetailsDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *CatalogAppDetailsDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *CatalogAppDetailsDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.

### HasAppRepoName

`func (o *CatalogAppDetailsDto) HasAppRepoName() bool`

HasAppRepoName returns a boolean if a field has been set.

### SetAppRepoNameNil

`func (o *CatalogAppDetailsDto) SetAppRepoNameNil(b bool)`

 SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil

### UnsetAppRepoName
`func (o *CatalogAppDetailsDto) UnsetAppRepoName()`

UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
### GetAppRepoOrganizationName

`func (o *CatalogAppDetailsDto) GetAppRepoOrganizationName() string`

GetAppRepoOrganizationName returns the AppRepoOrganizationName field if non-nil, zero value otherwise.

### GetAppRepoOrganizationNameOk

`func (o *CatalogAppDetailsDto) GetAppRepoOrganizationNameOk() (*string, bool)`

GetAppRepoOrganizationNameOk returns a tuple with the AppRepoOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoOrganizationName

`func (o *CatalogAppDetailsDto) SetAppRepoOrganizationName(v string)`

SetAppRepoOrganizationName sets AppRepoOrganizationName field to given value.

### HasAppRepoOrganizationName

`func (o *CatalogAppDetailsDto) HasAppRepoOrganizationName() bool`

HasAppRepoOrganizationName returns a boolean if a field has been set.

### SetAppRepoOrganizationNameNil

`func (o *CatalogAppDetailsDto) SetAppRepoOrganizationNameNil(b bool)`

 SetAppRepoOrganizationNameNil sets the value for AppRepoOrganizationName to be an explicit nil

### UnsetAppRepoOrganizationName
`func (o *CatalogAppDetailsDto) UnsetAppRepoOrganizationName()`

UnsetAppRepoOrganizationName ensures that no value is present for AppRepoOrganizationName, not even an explicit nil
### GetAppRepoId

`func (o *CatalogAppDetailsDto) GetAppRepoId() int32`

GetAppRepoId returns the AppRepoId field if non-nil, zero value otherwise.

### GetAppRepoIdOk

`func (o *CatalogAppDetailsDto) GetAppRepoIdOk() (*int32, bool)`

GetAppRepoIdOk returns a tuple with the AppRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoId

`func (o *CatalogAppDetailsDto) SetAppRepoId(v int32)`

SetAppRepoId sets AppRepoId field to given value.

### HasAppRepoId

`func (o *CatalogAppDetailsDto) HasAppRepoId() bool`

HasAppRepoId returns a boolean if a field has been set.

### GetCatalogName

`func (o *CatalogAppDetailsDto) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *CatalogAppDetailsDto) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *CatalogAppDetailsDto) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *CatalogAppDetailsDto) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### SetCatalogNameNil

`func (o *CatalogAppDetailsDto) SetCatalogNameNil(b bool)`

 SetCatalogNameNil sets the value for CatalogName to be an explicit nil

### UnsetCatalogName
`func (o *CatalogAppDetailsDto) UnsetCatalogName()`

UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
### GetCatalogId

`func (o *CatalogAppDetailsDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CatalogAppDetailsDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CatalogAppDetailsDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *CatalogAppDetailsDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetPackageId

`func (o *CatalogAppDetailsDto) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CatalogAppDetailsDto) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CatalogAppDetailsDto) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CatalogAppDetailsDto) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageIdNil

`func (o *CatalogAppDetailsDto) SetPackageIdNil(b bool)`

 SetPackageIdNil sets the value for PackageId to be an explicit nil

### UnsetPackageId
`func (o *CatalogAppDetailsDto) UnsetPackageId()`

UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
### GetVersion

`func (o *CatalogAppDetailsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogAppDetailsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogAppDetailsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogAppDetailsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CatalogAppDetailsDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CatalogAppDetailsDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetLogoId

`func (o *CatalogAppDetailsDto) GetLogoId() string`

GetLogoId returns the LogoId field if non-nil, zero value otherwise.

### GetLogoIdOk

`func (o *CatalogAppDetailsDto) GetLogoIdOk() (*string, bool)`

GetLogoIdOk returns a tuple with the LogoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoId

`func (o *CatalogAppDetailsDto) SetLogoId(v string)`

SetLogoId sets LogoId field to given value.

### HasLogoId

`func (o *CatalogAppDetailsDto) HasLogoId() bool`

HasLogoId returns a boolean if a field has been set.

### SetLogoIdNil

`func (o *CatalogAppDetailsDto) SetLogoIdNil(b bool)`

 SetLogoIdNil sets the value for LogoId to be an explicit nil

### UnsetLogoId
`func (o *CatalogAppDetailsDto) UnsetLogoId()`

UnsetLogoId ensures that no value is present for LogoId, not even an explicit nil
### GetProjectApps

`func (o *CatalogAppDetailsDto) GetProjectApps() []ProjectAppDto`

GetProjectApps returns the ProjectApps field if non-nil, zero value otherwise.

### GetProjectAppsOk

`func (o *CatalogAppDetailsDto) GetProjectAppsOk() (*[]ProjectAppDto, bool)`

GetProjectAppsOk returns a tuple with the ProjectApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectApps

`func (o *CatalogAppDetailsDto) SetProjectApps(v []ProjectAppDto)`

SetProjectApps sets ProjectApps field to given value.

### HasProjectApps

`func (o *CatalogAppDetailsDto) HasProjectApps() bool`

HasProjectApps returns a boolean if a field has been set.

### SetProjectAppsNil

`func (o *CatalogAppDetailsDto) SetProjectAppsNil(b bool)`

 SetProjectAppsNil sets the value for ProjectApps to be an explicit nil

### UnsetProjectApps
`func (o *CatalogAppDetailsDto) UnsetProjectApps()`

UnsetProjectApps ensures that no value is present for ProjectApps, not even an explicit nil
### GetDescription

`func (o *CatalogAppDetailsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogAppDetailsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogAppDetailsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogAppDetailsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogAppDetailsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogAppDetailsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReadme

`func (o *CatalogAppDetailsDto) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *CatalogAppDetailsDto) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *CatalogAppDetailsDto) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *CatalogAppDetailsDto) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### SetReadmeNil

`func (o *CatalogAppDetailsDto) SetReadmeNil(b bool)`

 SetReadmeNil sets the value for Readme to be an explicit nil

### UnsetReadme
`func (o *CatalogAppDetailsDto) UnsetReadme()`

UnsetReadme ensures that no value is present for Readme, not even an explicit nil
### GetSecurityReport

`func (o *CatalogAppDetailsDto) GetSecurityReport() SecurityReportSummaryDto`

GetSecurityReport returns the SecurityReport field if non-nil, zero value otherwise.

### GetSecurityReportOk

`func (o *CatalogAppDetailsDto) GetSecurityReportOk() (*SecurityReportSummaryDto, bool)`

GetSecurityReportOk returns a tuple with the SecurityReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReport

`func (o *CatalogAppDetailsDto) SetSecurityReport(v SecurityReportSummaryDto)`

SetSecurityReport sets SecurityReport field to given value.

### HasSecurityReport

`func (o *CatalogAppDetailsDto) HasSecurityReport() bool`

HasSecurityReport returns a boolean if a field has been set.

### GetAppVersion

`func (o *CatalogAppDetailsDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *CatalogAppDetailsDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *CatalogAppDetailsDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *CatalogAppDetailsDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *CatalogAppDetailsDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *CatalogAppDetailsDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetStars

`func (o *CatalogAppDetailsDto) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *CatalogAppDetailsDto) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *CatalogAppDetailsDto) SetStars(v int32)`

SetStars sets Stars field to given value.

### HasStars

`func (o *CatalogAppDetailsDto) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetVerifiedPublisher

`func (o *CatalogAppDetailsDto) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *CatalogAppDetailsDto) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *CatalogAppDetailsDto) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *CatalogAppDetailsDto) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### GetOfficial

`func (o *CatalogAppDetailsDto) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *CatalogAppDetailsDto) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *CatalogAppDetailsDto) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *CatalogAppDetailsDto) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetHasJsonSchema

`func (o *CatalogAppDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *CatalogAppDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *CatalogAppDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.

### HasHasJsonSchema

`func (o *CatalogAppDetailsDto) HasHasJsonSchema() bool`

HasHasJsonSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


