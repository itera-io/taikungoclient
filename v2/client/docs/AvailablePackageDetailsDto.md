# AvailablePackageDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**AppRepoName** | Pointer to **NullableString** |  | [optional] 
**AppRepoOrganizationName** | Pointer to **NullableString** |  | [optional] 
**AppRepoId** | Pointer to **NullableString** |  | [optional] 
**PackageId** | Pointer to **NullableString** |  | [optional] 
**LogoId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Readme** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**SecurityReport** | Pointer to [**SecurityReportSummaryDto**](SecurityReportSummaryDto.md) |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**Stars** | Pointer to **int32** |  | [optional] 
**VerifiedPublisher** | Pointer to **bool** |  | [optional] 
**Official** | Pointer to **bool** |  | [optional] 
**BoundCatalogs** | Pointer to [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | [optional] 
**HasJsonSchema** | Pointer to **bool** |  | [optional] 

## Methods

### NewAvailablePackageDetailsDto

`func NewAvailablePackageDetailsDto() *AvailablePackageDetailsDto`

NewAvailablePackageDetailsDto instantiates a new AvailablePackageDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePackageDetailsDtoWithDefaults

`func NewAvailablePackageDetailsDtoWithDefaults() *AvailablePackageDetailsDto`

NewAvailablePackageDetailsDtoWithDefaults instantiates a new AvailablePackageDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AvailablePackageDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailablePackageDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailablePackageDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AvailablePackageDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AvailablePackageDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AvailablePackageDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAppRepoName

`func (o *AvailablePackageDetailsDto) GetAppRepoName() string`

GetAppRepoName returns the AppRepoName field if non-nil, zero value otherwise.

### GetAppRepoNameOk

`func (o *AvailablePackageDetailsDto) GetAppRepoNameOk() (*string, bool)`

GetAppRepoNameOk returns a tuple with the AppRepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoName

`func (o *AvailablePackageDetailsDto) SetAppRepoName(v string)`

SetAppRepoName sets AppRepoName field to given value.

### HasAppRepoName

`func (o *AvailablePackageDetailsDto) HasAppRepoName() bool`

HasAppRepoName returns a boolean if a field has been set.

### SetAppRepoNameNil

`func (o *AvailablePackageDetailsDto) SetAppRepoNameNil(b bool)`

 SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil

### UnsetAppRepoName
`func (o *AvailablePackageDetailsDto) UnsetAppRepoName()`

UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
### GetAppRepoOrganizationName

`func (o *AvailablePackageDetailsDto) GetAppRepoOrganizationName() string`

GetAppRepoOrganizationName returns the AppRepoOrganizationName field if non-nil, zero value otherwise.

### GetAppRepoOrganizationNameOk

`func (o *AvailablePackageDetailsDto) GetAppRepoOrganizationNameOk() (*string, bool)`

GetAppRepoOrganizationNameOk returns a tuple with the AppRepoOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoOrganizationName

`func (o *AvailablePackageDetailsDto) SetAppRepoOrganizationName(v string)`

SetAppRepoOrganizationName sets AppRepoOrganizationName field to given value.

### HasAppRepoOrganizationName

`func (o *AvailablePackageDetailsDto) HasAppRepoOrganizationName() bool`

HasAppRepoOrganizationName returns a boolean if a field has been set.

### SetAppRepoOrganizationNameNil

`func (o *AvailablePackageDetailsDto) SetAppRepoOrganizationNameNil(b bool)`

 SetAppRepoOrganizationNameNil sets the value for AppRepoOrganizationName to be an explicit nil

### UnsetAppRepoOrganizationName
`func (o *AvailablePackageDetailsDto) UnsetAppRepoOrganizationName()`

UnsetAppRepoOrganizationName ensures that no value is present for AppRepoOrganizationName, not even an explicit nil
### GetAppRepoId

`func (o *AvailablePackageDetailsDto) GetAppRepoId() string`

GetAppRepoId returns the AppRepoId field if non-nil, zero value otherwise.

### GetAppRepoIdOk

`func (o *AvailablePackageDetailsDto) GetAppRepoIdOk() (*string, bool)`

GetAppRepoIdOk returns a tuple with the AppRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoId

`func (o *AvailablePackageDetailsDto) SetAppRepoId(v string)`

SetAppRepoId sets AppRepoId field to given value.

### HasAppRepoId

`func (o *AvailablePackageDetailsDto) HasAppRepoId() bool`

HasAppRepoId returns a boolean if a field has been set.

### SetAppRepoIdNil

`func (o *AvailablePackageDetailsDto) SetAppRepoIdNil(b bool)`

 SetAppRepoIdNil sets the value for AppRepoId to be an explicit nil

### UnsetAppRepoId
`func (o *AvailablePackageDetailsDto) UnsetAppRepoId()`

UnsetAppRepoId ensures that no value is present for AppRepoId, not even an explicit nil
### GetPackageId

`func (o *AvailablePackageDetailsDto) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AvailablePackageDetailsDto) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AvailablePackageDetailsDto) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *AvailablePackageDetailsDto) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageIdNil

`func (o *AvailablePackageDetailsDto) SetPackageIdNil(b bool)`

 SetPackageIdNil sets the value for PackageId to be an explicit nil

### UnsetPackageId
`func (o *AvailablePackageDetailsDto) UnsetPackageId()`

UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
### GetLogoId

`func (o *AvailablePackageDetailsDto) GetLogoId() string`

GetLogoId returns the LogoId field if non-nil, zero value otherwise.

### GetLogoIdOk

`func (o *AvailablePackageDetailsDto) GetLogoIdOk() (*string, bool)`

GetLogoIdOk returns a tuple with the LogoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoId

`func (o *AvailablePackageDetailsDto) SetLogoId(v string)`

SetLogoId sets LogoId field to given value.

### HasLogoId

`func (o *AvailablePackageDetailsDto) HasLogoId() bool`

HasLogoId returns a boolean if a field has been set.

### SetLogoIdNil

`func (o *AvailablePackageDetailsDto) SetLogoIdNil(b bool)`

 SetLogoIdNil sets the value for LogoId to be an explicit nil

### UnsetLogoId
`func (o *AvailablePackageDetailsDto) UnsetLogoId()`

UnsetLogoId ensures that no value is present for LogoId, not even an explicit nil
### GetDescription

`func (o *AvailablePackageDetailsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailablePackageDetailsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailablePackageDetailsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailablePackageDetailsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AvailablePackageDetailsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AvailablePackageDetailsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReadme

`func (o *AvailablePackageDetailsDto) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *AvailablePackageDetailsDto) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *AvailablePackageDetailsDto) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *AvailablePackageDetailsDto) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### SetReadmeNil

`func (o *AvailablePackageDetailsDto) SetReadmeNil(b bool)`

 SetReadmeNil sets the value for Readme to be an explicit nil

### UnsetReadme
`func (o *AvailablePackageDetailsDto) UnsetReadme()`

UnsetReadme ensures that no value is present for Readme, not even an explicit nil
### GetVersion

`func (o *AvailablePackageDetailsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AvailablePackageDetailsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AvailablePackageDetailsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AvailablePackageDetailsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AvailablePackageDetailsDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AvailablePackageDetailsDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSecurityReport

`func (o *AvailablePackageDetailsDto) GetSecurityReport() SecurityReportSummaryDto`

GetSecurityReport returns the SecurityReport field if non-nil, zero value otherwise.

### GetSecurityReportOk

`func (o *AvailablePackageDetailsDto) GetSecurityReportOk() (*SecurityReportSummaryDto, bool)`

GetSecurityReportOk returns a tuple with the SecurityReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReport

`func (o *AvailablePackageDetailsDto) SetSecurityReport(v SecurityReportSummaryDto)`

SetSecurityReport sets SecurityReport field to given value.

### HasSecurityReport

`func (o *AvailablePackageDetailsDto) HasSecurityReport() bool`

HasSecurityReport returns a boolean if a field has been set.

### GetAppVersion

`func (o *AvailablePackageDetailsDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AvailablePackageDetailsDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AvailablePackageDetailsDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *AvailablePackageDetailsDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *AvailablePackageDetailsDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *AvailablePackageDetailsDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetStars

`func (o *AvailablePackageDetailsDto) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *AvailablePackageDetailsDto) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *AvailablePackageDetailsDto) SetStars(v int32)`

SetStars sets Stars field to given value.

### HasStars

`func (o *AvailablePackageDetailsDto) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetVerifiedPublisher

`func (o *AvailablePackageDetailsDto) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *AvailablePackageDetailsDto) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *AvailablePackageDetailsDto) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *AvailablePackageDetailsDto) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### GetOfficial

`func (o *AvailablePackageDetailsDto) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *AvailablePackageDetailsDto) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *AvailablePackageDetailsDto) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *AvailablePackageDetailsDto) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetBoundCatalogs

`func (o *AvailablePackageDetailsDto) GetBoundCatalogs() []CommonDropdownDto`

GetBoundCatalogs returns the BoundCatalogs field if non-nil, zero value otherwise.

### GetBoundCatalogsOk

`func (o *AvailablePackageDetailsDto) GetBoundCatalogsOk() (*[]CommonDropdownDto, bool)`

GetBoundCatalogsOk returns a tuple with the BoundCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCatalogs

`func (o *AvailablePackageDetailsDto) SetBoundCatalogs(v []CommonDropdownDto)`

SetBoundCatalogs sets BoundCatalogs field to given value.

### HasBoundCatalogs

`func (o *AvailablePackageDetailsDto) HasBoundCatalogs() bool`

HasBoundCatalogs returns a boolean if a field has been set.

### SetBoundCatalogsNil

`func (o *AvailablePackageDetailsDto) SetBoundCatalogsNil(b bool)`

 SetBoundCatalogsNil sets the value for BoundCatalogs to be an explicit nil

### UnsetBoundCatalogs
`func (o *AvailablePackageDetailsDto) UnsetBoundCatalogs()`

UnsetBoundCatalogs ensures that no value is present for BoundCatalogs, not even an explicit nil
### GetHasJsonSchema

`func (o *AvailablePackageDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *AvailablePackageDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *AvailablePackageDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.

### HasHasJsonSchema

`func (o *AvailablePackageDetailsDto) HasHasJsonSchema() bool`

HasHasJsonSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


