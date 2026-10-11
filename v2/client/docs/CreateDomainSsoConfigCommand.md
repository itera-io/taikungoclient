# CreateDomainSsoConfigCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DomainId** | **int32** |  | 
**IssuerUrl** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**GroupFromIdp** | **bool** |  | 
**GroupClaimName** | Pointer to **NullableString** |  | [optional] 
**GroupScopeName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | **bool** |  | 

## Methods

### NewCreateDomainSsoConfigCommand

`func NewCreateDomainSsoConfigCommand(name string, domainId int32, issuerUrl string, clientId string, clientSecret string, groupFromIdp bool, isEnabled bool, ) *CreateDomainSsoConfigCommand`

NewCreateDomainSsoConfigCommand instantiates a new CreateDomainSsoConfigCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainSsoConfigCommandWithDefaults

`func NewCreateDomainSsoConfigCommandWithDefaults() *CreateDomainSsoConfigCommand`

NewCreateDomainSsoConfigCommandWithDefaults instantiates a new CreateDomainSsoConfigCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDomainSsoConfigCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDomainSsoConfigCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDomainSsoConfigCommand) SetName(v string)`

SetName sets Name field to given value.


### GetDomainId

`func (o *CreateDomainSsoConfigCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateDomainSsoConfigCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateDomainSsoConfigCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetIssuerUrl

`func (o *CreateDomainSsoConfigCommand) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CreateDomainSsoConfigCommand) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CreateDomainSsoConfigCommand) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetClientId

`func (o *CreateDomainSsoConfigCommand) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateDomainSsoConfigCommand) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateDomainSsoConfigCommand) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateDomainSsoConfigCommand) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateDomainSsoConfigCommand) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateDomainSsoConfigCommand) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGroupFromIdp

`func (o *CreateDomainSsoConfigCommand) GetGroupFromIdp() bool`

GetGroupFromIdp returns the GroupFromIdp field if non-nil, zero value otherwise.

### GetGroupFromIdpOk

`func (o *CreateDomainSsoConfigCommand) GetGroupFromIdpOk() (*bool, bool)`

GetGroupFromIdpOk returns a tuple with the GroupFromIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFromIdp

`func (o *CreateDomainSsoConfigCommand) SetGroupFromIdp(v bool)`

SetGroupFromIdp sets GroupFromIdp field to given value.


### GetGroupClaimName

`func (o *CreateDomainSsoConfigCommand) GetGroupClaimName() string`

GetGroupClaimName returns the GroupClaimName field if non-nil, zero value otherwise.

### GetGroupClaimNameOk

`func (o *CreateDomainSsoConfigCommand) GetGroupClaimNameOk() (*string, bool)`

GetGroupClaimNameOk returns a tuple with the GroupClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaimName

`func (o *CreateDomainSsoConfigCommand) SetGroupClaimName(v string)`

SetGroupClaimName sets GroupClaimName field to given value.

### HasGroupClaimName

`func (o *CreateDomainSsoConfigCommand) HasGroupClaimName() bool`

HasGroupClaimName returns a boolean if a field has been set.

### SetGroupClaimNameNil

`func (o *CreateDomainSsoConfigCommand) SetGroupClaimNameNil(b bool)`

 SetGroupClaimNameNil sets the value for GroupClaimName to be an explicit nil

### UnsetGroupClaimName
`func (o *CreateDomainSsoConfigCommand) UnsetGroupClaimName()`

UnsetGroupClaimName ensures that no value is present for GroupClaimName, not even an explicit nil
### GetGroupScopeName

`func (o *CreateDomainSsoConfigCommand) GetGroupScopeName() string`

GetGroupScopeName returns the GroupScopeName field if non-nil, zero value otherwise.

### GetGroupScopeNameOk

`func (o *CreateDomainSsoConfigCommand) GetGroupScopeNameOk() (*string, bool)`

GetGroupScopeNameOk returns a tuple with the GroupScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScopeName

`func (o *CreateDomainSsoConfigCommand) SetGroupScopeName(v string)`

SetGroupScopeName sets GroupScopeName field to given value.

### HasGroupScopeName

`func (o *CreateDomainSsoConfigCommand) HasGroupScopeName() bool`

HasGroupScopeName returns a boolean if a field has been set.

### SetGroupScopeNameNil

`func (o *CreateDomainSsoConfigCommand) SetGroupScopeNameNil(b bool)`

 SetGroupScopeNameNil sets the value for GroupScopeName to be an explicit nil

### UnsetGroupScopeName
`func (o *CreateDomainSsoConfigCommand) UnsetGroupScopeName()`

UnsetGroupScopeName ensures that no value is present for GroupScopeName, not even an explicit nil
### GetIsEnabled

`func (o *CreateDomainSsoConfigCommand) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateDomainSsoConfigCommand) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateDomainSsoConfigCommand) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


