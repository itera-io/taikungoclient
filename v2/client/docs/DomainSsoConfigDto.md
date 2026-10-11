# DomainSsoConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**PartnerId** | **int32** |  | 
**IssuerUrl** | **string** |  | 
**ClientId** | **string** |  | 
**GroupFromIdp** | **bool** |  | 
**GroupClaimName** | Pointer to **NullableString** |  | [optional] 
**GroupScopeName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | **bool** |  | 

## Methods

### NewDomainSsoConfigDto

`func NewDomainSsoConfigDto(id int32, name string, partnerId int32, issuerUrl string, clientId string, groupFromIdp bool, isEnabled bool, ) *DomainSsoConfigDto`

NewDomainSsoConfigDto instantiates a new DomainSsoConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSsoConfigDtoWithDefaults

`func NewDomainSsoConfigDtoWithDefaults() *DomainSsoConfigDto`

NewDomainSsoConfigDtoWithDefaults instantiates a new DomainSsoConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainSsoConfigDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSsoConfigDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSsoConfigDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DomainSsoConfigDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainSsoConfigDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainSsoConfigDto) SetName(v string)`

SetName sets Name field to given value.


### GetPartnerId

`func (o *DomainSsoConfigDto) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DomainSsoConfigDto) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DomainSsoConfigDto) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetIssuerUrl

`func (o *DomainSsoConfigDto) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *DomainSsoConfigDto) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *DomainSsoConfigDto) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetClientId

`func (o *DomainSsoConfigDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DomainSsoConfigDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DomainSsoConfigDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGroupFromIdp

`func (o *DomainSsoConfigDto) GetGroupFromIdp() bool`

GetGroupFromIdp returns the GroupFromIdp field if non-nil, zero value otherwise.

### GetGroupFromIdpOk

`func (o *DomainSsoConfigDto) GetGroupFromIdpOk() (*bool, bool)`

GetGroupFromIdpOk returns a tuple with the GroupFromIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFromIdp

`func (o *DomainSsoConfigDto) SetGroupFromIdp(v bool)`

SetGroupFromIdp sets GroupFromIdp field to given value.


### GetGroupClaimName

`func (o *DomainSsoConfigDto) GetGroupClaimName() string`

GetGroupClaimName returns the GroupClaimName field if non-nil, zero value otherwise.

### GetGroupClaimNameOk

`func (o *DomainSsoConfigDto) GetGroupClaimNameOk() (*string, bool)`

GetGroupClaimNameOk returns a tuple with the GroupClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaimName

`func (o *DomainSsoConfigDto) SetGroupClaimName(v string)`

SetGroupClaimName sets GroupClaimName field to given value.

### HasGroupClaimName

`func (o *DomainSsoConfigDto) HasGroupClaimName() bool`

HasGroupClaimName returns a boolean if a field has been set.

### SetGroupClaimNameNil

`func (o *DomainSsoConfigDto) SetGroupClaimNameNil(b bool)`

 SetGroupClaimNameNil sets the value for GroupClaimName to be an explicit nil

### UnsetGroupClaimName
`func (o *DomainSsoConfigDto) UnsetGroupClaimName()`

UnsetGroupClaimName ensures that no value is present for GroupClaimName, not even an explicit nil
### GetGroupScopeName

`func (o *DomainSsoConfigDto) GetGroupScopeName() string`

GetGroupScopeName returns the GroupScopeName field if non-nil, zero value otherwise.

### GetGroupScopeNameOk

`func (o *DomainSsoConfigDto) GetGroupScopeNameOk() (*string, bool)`

GetGroupScopeNameOk returns a tuple with the GroupScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScopeName

`func (o *DomainSsoConfigDto) SetGroupScopeName(v string)`

SetGroupScopeName sets GroupScopeName field to given value.

### HasGroupScopeName

`func (o *DomainSsoConfigDto) HasGroupScopeName() bool`

HasGroupScopeName returns a boolean if a field has been set.

### SetGroupScopeNameNil

`func (o *DomainSsoConfigDto) SetGroupScopeNameNil(b bool)`

 SetGroupScopeNameNil sets the value for GroupScopeName to be an explicit nil

### UnsetGroupScopeName
`func (o *DomainSsoConfigDto) UnsetGroupScopeName()`

UnsetGroupScopeName ensures that no value is present for GroupScopeName, not even an explicit nil
### GetIsEnabled

`func (o *DomainSsoConfigDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DomainSsoConfigDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DomainSsoConfigDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


