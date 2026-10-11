# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Official** | Pointer to **bool** |  | [optional] 
**RepositoryId** | Pointer to **NullableString** |  | [optional] 
**ScannerDisabled** | Pointer to **bool** |  | [optional] 
**IsImported** | Pointer to **bool** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**VerifiedPublisher** | Pointer to **bool** |  | [optional] 
**OrganizationDisplayName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRepository

`func NewRepository() *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Repository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Repository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Repository) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Repository) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetKind

`func (o *Repository) GetKind() int64`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Repository) GetKindOk() (*int64, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Repository) SetKind(v int64)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Repository) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Repository) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Repository) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOfficial

`func (o *Repository) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *Repository) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *Repository) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *Repository) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetRepositoryId

`func (o *Repository) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Repository) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Repository) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Repository) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### SetRepositoryIdNil

`func (o *Repository) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *Repository) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetScannerDisabled

`func (o *Repository) GetScannerDisabled() bool`

GetScannerDisabled returns the ScannerDisabled field if non-nil, zero value otherwise.

### GetScannerDisabledOk

`func (o *Repository) GetScannerDisabledOk() (*bool, bool)`

GetScannerDisabledOk returns a tuple with the ScannerDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannerDisabled

`func (o *Repository) SetScannerDisabled(v bool)`

SetScannerDisabled sets ScannerDisabled field to given value.

### HasScannerDisabled

`func (o *Repository) HasScannerDisabled() bool`

HasScannerDisabled returns a boolean if a field has been set.

### GetIsImported

`func (o *Repository) GetIsImported() bool`

GetIsImported returns the IsImported field if non-nil, zero value otherwise.

### GetIsImportedOk

`func (o *Repository) GetIsImportedOk() (*bool, bool)`

GetIsImportedOk returns a tuple with the IsImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImported

`func (o *Repository) SetIsImported(v bool)`

SetIsImported sets IsImported field to given value.

### HasIsImported

`func (o *Repository) HasIsImported() bool`

HasIsImported returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Repository) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Repository) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Repository) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Repository) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *Repository) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *Repository) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetVerifiedPublisher

`func (o *Repository) GetVerifiedPublisher() bool`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *Repository) GetVerifiedPublisherOk() (*bool, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *Repository) SetVerifiedPublisher(v bool)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *Repository) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### GetOrganizationDisplayName

`func (o *Repository) GetOrganizationDisplayName() string`

GetOrganizationDisplayName returns the OrganizationDisplayName field if non-nil, zero value otherwise.

### GetOrganizationDisplayNameOk

`func (o *Repository) GetOrganizationDisplayNameOk() (*string, bool)`

GetOrganizationDisplayNameOk returns a tuple with the OrganizationDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDisplayName

`func (o *Repository) SetOrganizationDisplayName(v string)`

SetOrganizationDisplayName sets OrganizationDisplayName field to given value.

### HasOrganizationDisplayName

`func (o *Repository) HasOrganizationDisplayName() bool`

HasOrganizationDisplayName returns a boolean if a field has been set.

### SetOrganizationDisplayNameNil

`func (o *Repository) SetOrganizationDisplayNameNil(b bool)`

 SetOrganizationDisplayNameNil sets the value for OrganizationDisplayName to be an explicit nil

### UnsetOrganizationDisplayName
`func (o *Repository) UnsetOrganizationDisplayName()`

UnsetOrganizationDisplayName ensures that no value is present for OrganizationDisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


