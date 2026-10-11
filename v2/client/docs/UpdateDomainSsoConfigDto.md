# UpdateDomainSsoConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**IssuerUrl** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**GroupFromIdp** | Pointer to **NullableBool** |  | [optional] 
**GroupClaimName** | Pointer to **NullableString** |  | [optional] 
**GroupScopeName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateDomainSsoConfigDto

`func NewUpdateDomainSsoConfigDto() *UpdateDomainSsoConfigDto`

NewUpdateDomainSsoConfigDto instantiates a new UpdateDomainSsoConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainSsoConfigDtoWithDefaults

`func NewUpdateDomainSsoConfigDtoWithDefaults() *UpdateDomainSsoConfigDto`

NewUpdateDomainSsoConfigDtoWithDefaults instantiates a new UpdateDomainSsoConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDomainSsoConfigDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDomainSsoConfigDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDomainSsoConfigDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDomainSsoConfigDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateDomainSsoConfigDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDomainSsoConfigDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIssuerUrl

`func (o *UpdateDomainSsoConfigDto) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *UpdateDomainSsoConfigDto) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *UpdateDomainSsoConfigDto) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *UpdateDomainSsoConfigDto) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### SetIssuerUrlNil

`func (o *UpdateDomainSsoConfigDto) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *UpdateDomainSsoConfigDto) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetClientId

`func (o *UpdateDomainSsoConfigDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateDomainSsoConfigDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateDomainSsoConfigDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateDomainSsoConfigDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *UpdateDomainSsoConfigDto) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *UpdateDomainSsoConfigDto) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *UpdateDomainSsoConfigDto) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateDomainSsoConfigDto) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateDomainSsoConfigDto) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateDomainSsoConfigDto) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *UpdateDomainSsoConfigDto) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *UpdateDomainSsoConfigDto) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetGroupFromIdp

`func (o *UpdateDomainSsoConfigDto) GetGroupFromIdp() bool`

GetGroupFromIdp returns the GroupFromIdp field if non-nil, zero value otherwise.

### GetGroupFromIdpOk

`func (o *UpdateDomainSsoConfigDto) GetGroupFromIdpOk() (*bool, bool)`

GetGroupFromIdpOk returns a tuple with the GroupFromIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFromIdp

`func (o *UpdateDomainSsoConfigDto) SetGroupFromIdp(v bool)`

SetGroupFromIdp sets GroupFromIdp field to given value.

### HasGroupFromIdp

`func (o *UpdateDomainSsoConfigDto) HasGroupFromIdp() bool`

HasGroupFromIdp returns a boolean if a field has been set.

### SetGroupFromIdpNil

`func (o *UpdateDomainSsoConfigDto) SetGroupFromIdpNil(b bool)`

 SetGroupFromIdpNil sets the value for GroupFromIdp to be an explicit nil

### UnsetGroupFromIdp
`func (o *UpdateDomainSsoConfigDto) UnsetGroupFromIdp()`

UnsetGroupFromIdp ensures that no value is present for GroupFromIdp, not even an explicit nil
### GetGroupClaimName

`func (o *UpdateDomainSsoConfigDto) GetGroupClaimName() string`

GetGroupClaimName returns the GroupClaimName field if non-nil, zero value otherwise.

### GetGroupClaimNameOk

`func (o *UpdateDomainSsoConfigDto) GetGroupClaimNameOk() (*string, bool)`

GetGroupClaimNameOk returns a tuple with the GroupClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaimName

`func (o *UpdateDomainSsoConfigDto) SetGroupClaimName(v string)`

SetGroupClaimName sets GroupClaimName field to given value.

### HasGroupClaimName

`func (o *UpdateDomainSsoConfigDto) HasGroupClaimName() bool`

HasGroupClaimName returns a boolean if a field has been set.

### SetGroupClaimNameNil

`func (o *UpdateDomainSsoConfigDto) SetGroupClaimNameNil(b bool)`

 SetGroupClaimNameNil sets the value for GroupClaimName to be an explicit nil

### UnsetGroupClaimName
`func (o *UpdateDomainSsoConfigDto) UnsetGroupClaimName()`

UnsetGroupClaimName ensures that no value is present for GroupClaimName, not even an explicit nil
### GetGroupScopeName

`func (o *UpdateDomainSsoConfigDto) GetGroupScopeName() string`

GetGroupScopeName returns the GroupScopeName field if non-nil, zero value otherwise.

### GetGroupScopeNameOk

`func (o *UpdateDomainSsoConfigDto) GetGroupScopeNameOk() (*string, bool)`

GetGroupScopeNameOk returns a tuple with the GroupScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScopeName

`func (o *UpdateDomainSsoConfigDto) SetGroupScopeName(v string)`

SetGroupScopeName sets GroupScopeName field to given value.

### HasGroupScopeName

`func (o *UpdateDomainSsoConfigDto) HasGroupScopeName() bool`

HasGroupScopeName returns a boolean if a field has been set.

### SetGroupScopeNameNil

`func (o *UpdateDomainSsoConfigDto) SetGroupScopeNameNil(b bool)`

 SetGroupScopeNameNil sets the value for GroupScopeName to be an explicit nil

### UnsetGroupScopeName
`func (o *UpdateDomainSsoConfigDto) UnsetGroupScopeName()`

UnsetGroupScopeName ensures that no value is present for GroupScopeName, not even an explicit nil
### GetIsEnabled

`func (o *UpdateDomainSsoConfigDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateDomainSsoConfigDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateDomainSsoConfigDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateDomainSsoConfigDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UpdateDomainSsoConfigDto) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UpdateDomainSsoConfigDto) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


