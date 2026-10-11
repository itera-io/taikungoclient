# ShowbackCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**Rules** | Pointer to [**[]ShowbackRuleEntityDto**](ShowbackRuleEntityDto.md) |  | [optional] 

## Methods

### NewShowbackCredentialsListDto

`func NewShowbackCredentialsListDto() *ShowbackCredentialsListDto`

NewShowbackCredentialsListDto instantiates a new ShowbackCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowbackCredentialsListDtoWithDefaults

`func NewShowbackCredentialsListDtoWithDefaults() *ShowbackCredentialsListDto`

NewShowbackCredentialsListDtoWithDefaults instantiates a new ShowbackCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShowbackCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowbackCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowbackCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShowbackCredentialsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShowbackCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShowbackCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShowbackCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShowbackCredentialsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShowbackCredentialsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShowbackCredentialsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *ShowbackCredentialsListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ShowbackCredentialsListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ShowbackCredentialsListDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ShowbackCredentialsListDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ShowbackCredentialsListDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ShowbackCredentialsListDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCreatedAt

`func (o *ShowbackCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShowbackCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShowbackCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShowbackCredentialsListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ShowbackCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ShowbackCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUsername

`func (o *ShowbackCredentialsListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ShowbackCredentialsListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ShowbackCredentialsListDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ShowbackCredentialsListDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ShowbackCredentialsListDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ShowbackCredentialsListDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ShowbackCredentialsListDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ShowbackCredentialsListDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ShowbackCredentialsListDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ShowbackCredentialsListDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ShowbackCredentialsListDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ShowbackCredentialsListDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetIsLocked

`func (o *ShowbackCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ShowbackCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ShowbackCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ShowbackCredentialsListDto) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ShowbackCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ShowbackCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ShowbackCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ShowbackCredentialsListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ShowbackCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ShowbackCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ShowbackCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ShowbackCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ShowbackCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ShowbackCredentialsListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ShowbackCredentialsListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ShowbackCredentialsListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ShowbackCredentialsListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ShowbackCredentialsListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ShowbackCredentialsListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ShowbackCredentialsListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *ShowbackCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ShowbackCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ShowbackCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ShowbackCredentialsListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *ShowbackCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ShowbackCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *ShowbackCredentialsListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ShowbackCredentialsListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ShowbackCredentialsListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ShowbackCredentialsListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *ShowbackCredentialsListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *ShowbackCredentialsListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetRules

`func (o *ShowbackCredentialsListDto) GetRules() []ShowbackRuleEntityDto`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ShowbackCredentialsListDto) GetRulesOk() (*[]ShowbackRuleEntityDto, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ShowbackCredentialsListDto) SetRules(v []ShowbackRuleEntityDto)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ShowbackCredentialsListDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *ShowbackCredentialsListDto) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *ShowbackCredentialsListDto) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


