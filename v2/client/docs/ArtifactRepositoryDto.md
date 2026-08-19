# ArtifactRepositoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryId** | **NullableString** |  | 
**AppRepoId** | **int32** |  | 
**Name** | **string** |  | 
**DisplayName** | **NullableString** |  | 
**Url** | **string** |  | 
**OrganizationName** | **string** |  | 
**Disabled** | **bool** |  | 
**VerifiedPublisher** | **bool** |  | 
**Official** | **bool** |  | 
**IsBound** | **bool** |  | 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**IsTaikun** | **bool** |  | 
**HasCatalogApp** | **bool** |  | 
**PasswordProtected** | Pointer to **bool** |  | [optional] 

## Methods

### NewArtifactRepositoryDto

`func NewArtifactRepositoryDto(repositoryId NullableString, appRepoId int32, name string, displayName NullableString, url string, organizationName string, disabled bool, verifiedPublisher bool, official bool, isBound bool, isTaikun bool, hasCatalogApp bool, ) *ArtifactRepositoryDto`

NewArtifactRepositoryDto instantiates a new ArtifactRepositoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryDtoWithDefaults

`func NewArtifactRepositoryDtoWithDefaults() *ArtifactRepositoryDto`

NewArtifactRepositoryDtoWithDefaults instantiates a new ArtifactRepositoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryId

`func (o *ArtifactRepositoryDto) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ArtifactRepositoryDto) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ArtifactRepositoryDto) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.


### SetRepositoryIdNil

`func (o *ArtifactRepositoryDto) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *ArtifactRepositoryDto) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetAppRepoId

`func (o *ArtifactRepositoryDto) GetAppRepoId() int32`

GetAppRepoId returns the AppRepoId field if non-nil, zero value otherwise.

### GetAppRepoIdOk

`func (o *ArtifactRepositoryDto) GetAppRepoIdOk() (*int32, bool)`

GetAppRepoIdOk returns a tuple with the AppRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRepoId

`func (o *ArtifactRepositoryDto) SetAppRepoId(v int32)`

SetAppRepoId sets AppRepoId field to given value.


### GetName

`func (o *ArtifactRepositoryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactRepositoryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactRepositoryDto) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ArtifactRepositoryDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArtifactRepositoryDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArtifactRepositoryDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *ArtifactRepositoryDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ArtifactRepositoryDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUrl

`func (o *ArtifactRepositoryDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArtifactRepositoryDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArtifactRepositoryDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOrganizationName

`func (o *ArtifactRepositoryDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ArtifactRepositoryDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ArtifactRepositoryDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetDisabled

`func (o *ArtifactRepositoryDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ArtifactRepositoryDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ArtifactRepositoryDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVerifiedPublisher

`func (o *ArtifactRepositoryDto) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *ArtifactRepositoryDto) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *ArtifactRepositoryDto) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.


### GetOfficial

`func (o *ArtifactRepositoryDto) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *ArtifactRepositoryDto) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *ArtifactRepositoryDto) SetOfficial(v bool)`

SetOfficial sets Official field to given value.


### GetIsBound

`func (o *ArtifactRepositoryDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *ArtifactRepositoryDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *ArtifactRepositoryDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.


### GetIsPrivate

`func (o *ArtifactRepositoryDto) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ArtifactRepositoryDto) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ArtifactRepositoryDto) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ArtifactRepositoryDto) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsTaikun

`func (o *ArtifactRepositoryDto) GetIsTaikun() bool`

GetIsTaikun returns the IsTaikun field if non-nil, zero value otherwise.

### GetIsTaikunOk

`func (o *ArtifactRepositoryDto) GetIsTaikunOk() (*bool, bool)`

GetIsTaikunOk returns a tuple with the IsTaikun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaikun

`func (o *ArtifactRepositoryDto) SetIsTaikun(v bool)`

SetIsTaikun sets IsTaikun field to given value.


### GetHasCatalogApp

`func (o *ArtifactRepositoryDto) GetHasCatalogApp() bool`

GetHasCatalogApp returns the HasCatalogApp field if non-nil, zero value otherwise.

### GetHasCatalogAppOk

`func (o *ArtifactRepositoryDto) GetHasCatalogAppOk() (*bool, bool)`

GetHasCatalogAppOk returns a tuple with the HasCatalogApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCatalogApp

`func (o *ArtifactRepositoryDto) SetHasCatalogApp(v bool)`

SetHasCatalogApp sets HasCatalogApp field to given value.


### GetPasswordProtected

`func (o *ArtifactRepositoryDto) GetPasswordProtected() bool`

GetPasswordProtected returns the PasswordProtected field if non-nil, zero value otherwise.

### GetPasswordProtectedOk

`func (o *ArtifactRepositoryDto) GetPasswordProtectedOk() (*bool, bool)`

GetPasswordProtectedOk returns a tuple with the PasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtected

`func (o *ArtifactRepositoryDto) SetPasswordProtected(v bool)`

SetPasswordProtected sets PasswordProtected field to given value.

### HasPasswordProtected

`func (o *ArtifactRepositoryDto) HasPasswordProtected() bool`

HasPasswordProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


