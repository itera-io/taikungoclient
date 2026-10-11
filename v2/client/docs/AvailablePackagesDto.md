# AvailablePackagesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | Pointer to **NullableString** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**CatalogAppId** | Pointer to **NullableInt32** |  | [optional] 
**InstalledInstanceCount** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**NormalizedName** | Pointer to **NullableString** |  | [optional] 
**LogoImageId** | Pointer to **NullableString** |  | [optional] 
**Stars** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**AppVersion** | Pointer to **NullableString** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Signed** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**SecurityReportSummary** | Pointer to [**SecurityReportSummary**](SecurityReportSummary.md) |  | [optional] 
**Ts** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**IsAdded** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAvailablePackagesDto

`func NewAvailablePackagesDto() *AvailablePackagesDto`

NewAvailablePackagesDto instantiates a new AvailablePackagesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePackagesDtoWithDefaults

`func NewAvailablePackagesDtoWithDefaults() *AvailablePackagesDto`

NewAvailablePackagesDtoWithDefaults instantiates a new AvailablePackagesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *AvailablePackagesDto) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AvailablePackagesDto) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AvailablePackagesDto) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *AvailablePackagesDto) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageIdNil

`func (o *AvailablePackagesDto) SetPackageIdNil(b bool)`

 SetPackageIdNil sets the value for PackageId to be an explicit nil

### UnsetPackageId
`func (o *AvailablePackagesDto) UnsetPackageId()`

UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
### GetCatalogId

`func (o *AvailablePackagesDto) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AvailablePackagesDto) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *AvailablePackagesDto) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *AvailablePackagesDto) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogAppId

`func (o *AvailablePackagesDto) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *AvailablePackagesDto) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *AvailablePackagesDto) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.

### HasCatalogAppId

`func (o *AvailablePackagesDto) HasCatalogAppId() bool`

HasCatalogAppId returns a boolean if a field has been set.

### SetCatalogAppIdNil

`func (o *AvailablePackagesDto) SetCatalogAppIdNil(b bool)`

 SetCatalogAppIdNil sets the value for CatalogAppId to be an explicit nil

### UnsetCatalogAppId
`func (o *AvailablePackagesDto) UnsetCatalogAppId()`

UnsetCatalogAppId ensures that no value is present for CatalogAppId, not even an explicit nil
### GetInstalledInstanceCount

`func (o *AvailablePackagesDto) GetInstalledInstanceCount() int32`

GetInstalledInstanceCount returns the InstalledInstanceCount field if non-nil, zero value otherwise.

### GetInstalledInstanceCountOk

`func (o *AvailablePackagesDto) GetInstalledInstanceCountOk() (*int32, bool)`

GetInstalledInstanceCountOk returns a tuple with the InstalledInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledInstanceCount

`func (o *AvailablePackagesDto) SetInstalledInstanceCount(v int32)`

SetInstalledInstanceCount sets InstalledInstanceCount field to given value.

### HasInstalledInstanceCount

`func (o *AvailablePackagesDto) HasInstalledInstanceCount() bool`

HasInstalledInstanceCount returns a boolean if a field has been set.

### SetInstalledInstanceCountNil

`func (o *AvailablePackagesDto) SetInstalledInstanceCountNil(b bool)`

 SetInstalledInstanceCountNil sets the value for InstalledInstanceCount to be an explicit nil

### UnsetInstalledInstanceCount
`func (o *AvailablePackagesDto) UnsetInstalledInstanceCount()`

UnsetInstalledInstanceCount ensures that no value is present for InstalledInstanceCount, not even an explicit nil
### GetName

`func (o *AvailablePackagesDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailablePackagesDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailablePackagesDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AvailablePackagesDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AvailablePackagesDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AvailablePackagesDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNormalizedName

`func (o *AvailablePackagesDto) GetNormalizedName() string`

GetNormalizedName returns the NormalizedName field if non-nil, zero value otherwise.

### GetNormalizedNameOk

`func (o *AvailablePackagesDto) GetNormalizedNameOk() (*string, bool)`

GetNormalizedNameOk returns a tuple with the NormalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedName

`func (o *AvailablePackagesDto) SetNormalizedName(v string)`

SetNormalizedName sets NormalizedName field to given value.

### HasNormalizedName

`func (o *AvailablePackagesDto) HasNormalizedName() bool`

HasNormalizedName returns a boolean if a field has been set.

### SetNormalizedNameNil

`func (o *AvailablePackagesDto) SetNormalizedNameNil(b bool)`

 SetNormalizedNameNil sets the value for NormalizedName to be an explicit nil

### UnsetNormalizedName
`func (o *AvailablePackagesDto) UnsetNormalizedName()`

UnsetNormalizedName ensures that no value is present for NormalizedName, not even an explicit nil
### GetLogoImageId

`func (o *AvailablePackagesDto) GetLogoImageId() string`

GetLogoImageId returns the LogoImageId field if non-nil, zero value otherwise.

### GetLogoImageIdOk

`func (o *AvailablePackagesDto) GetLogoImageIdOk() (*string, bool)`

GetLogoImageIdOk returns a tuple with the LogoImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImageId

`func (o *AvailablePackagesDto) SetLogoImageId(v string)`

SetLogoImageId sets LogoImageId field to given value.

### HasLogoImageId

`func (o *AvailablePackagesDto) HasLogoImageId() bool`

HasLogoImageId returns a boolean if a field has been set.

### SetLogoImageIdNil

`func (o *AvailablePackagesDto) SetLogoImageIdNil(b bool)`

 SetLogoImageIdNil sets the value for LogoImageId to be an explicit nil

### UnsetLogoImageId
`func (o *AvailablePackagesDto) UnsetLogoImageId()`

UnsetLogoImageId ensures that no value is present for LogoImageId, not even an explicit nil
### GetStars

`func (o *AvailablePackagesDto) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *AvailablePackagesDto) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *AvailablePackagesDto) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *AvailablePackagesDto) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetDescription

`func (o *AvailablePackagesDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailablePackagesDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailablePackagesDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailablePackagesDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AvailablePackagesDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AvailablePackagesDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *AvailablePackagesDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AvailablePackagesDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AvailablePackagesDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AvailablePackagesDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AvailablePackagesDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AvailablePackagesDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAppVersion

`func (o *AvailablePackagesDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AvailablePackagesDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AvailablePackagesDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *AvailablePackagesDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *AvailablePackagesDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *AvailablePackagesDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetDeprecated

`func (o *AvailablePackagesDto) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AvailablePackagesDto) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AvailablePackagesDto) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AvailablePackagesDto) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetSigned

`func (o *AvailablePackagesDto) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *AvailablePackagesDto) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *AvailablePackagesDto) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *AvailablePackagesDto) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetIsLocked

`func (o *AvailablePackagesDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AvailablePackagesDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AvailablePackagesDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AvailablePackagesDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetSecurityReportSummary

`func (o *AvailablePackagesDto) GetSecurityReportSummary() SecurityReportSummary`

GetSecurityReportSummary returns the SecurityReportSummary field if non-nil, zero value otherwise.

### GetSecurityReportSummaryOk

`func (o *AvailablePackagesDto) GetSecurityReportSummaryOk() (*SecurityReportSummary, bool)`

GetSecurityReportSummaryOk returns a tuple with the SecurityReportSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReportSummary

`func (o *AvailablePackagesDto) SetSecurityReportSummary(v SecurityReportSummary)`

SetSecurityReportSummary sets SecurityReportSummary field to given value.

### HasSecurityReportSummary

`func (o *AvailablePackagesDto) HasSecurityReportSummary() bool`

HasSecurityReportSummary returns a boolean if a field has been set.

### GetTs

`func (o *AvailablePackagesDto) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *AvailablePackagesDto) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *AvailablePackagesDto) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *AvailablePackagesDto) HasTs() bool`

HasTs returns a boolean if a field has been set.

### SetTsNil

`func (o *AvailablePackagesDto) SetTsNil(b bool)`

 SetTsNil sets the value for Ts to be an explicit nil

### UnsetTs
`func (o *AvailablePackagesDto) UnsetTs()`

UnsetTs ensures that no value is present for Ts, not even an explicit nil
### GetRepository

`func (o *AvailablePackagesDto) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AvailablePackagesDto) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AvailablePackagesDto) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AvailablePackagesDto) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetIsAdded

`func (o *AvailablePackagesDto) GetIsAdded() bool`

GetIsAdded returns the IsAdded field if non-nil, zero value otherwise.

### GetIsAddedOk

`func (o *AvailablePackagesDto) GetIsAddedOk() (*bool, bool)`

GetIsAddedOk returns a tuple with the IsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdded

`func (o *AvailablePackagesDto) SetIsAdded(v bool)`

SetIsAdded sets IsAdded field to given value.

### HasIsAdded

`func (o *AvailablePackagesDto) HasIsAdded() bool`

HasIsAdded returns a boolean if a field has been set.

### SetIsAddedNil

`func (o *AvailablePackagesDto) SetIsAddedNil(b bool)`

 SetIsAddedNil sets the value for IsAdded to be an explicit nil

### UnsetIsAdded
`func (o *AvailablePackagesDto) UnsetIsAdded()`

UnsetIsAdded ensures that no value is present for IsAdded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


