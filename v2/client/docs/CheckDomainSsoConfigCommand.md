# CheckDomainSsoConfigCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerUrl** | **NullableString** |  | 
**ClientId** | **NullableString** |  | 
**ClientSecret** | **NullableString** |  | 

## Methods

### NewCheckDomainSsoConfigCommand

`func NewCheckDomainSsoConfigCommand(issuerUrl NullableString, clientId NullableString, clientSecret NullableString, ) *CheckDomainSsoConfigCommand`

NewCheckDomainSsoConfigCommand instantiates a new CheckDomainSsoConfigCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDomainSsoConfigCommandWithDefaults

`func NewCheckDomainSsoConfigCommandWithDefaults() *CheckDomainSsoConfigCommand`

NewCheckDomainSsoConfigCommandWithDefaults instantiates a new CheckDomainSsoConfigCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerUrl

`func (o *CheckDomainSsoConfigCommand) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CheckDomainSsoConfigCommand) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CheckDomainSsoConfigCommand) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### SetIssuerUrlNil

`func (o *CheckDomainSsoConfigCommand) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *CheckDomainSsoConfigCommand) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetClientId

`func (o *CheckDomainSsoConfigCommand) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CheckDomainSsoConfigCommand) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CheckDomainSsoConfigCommand) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### SetClientIdNil

`func (o *CheckDomainSsoConfigCommand) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CheckDomainSsoConfigCommand) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *CheckDomainSsoConfigCommand) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CheckDomainSsoConfigCommand) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CheckDomainSsoConfigCommand) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### SetClientSecretNil

`func (o *CheckDomainSsoConfigCommand) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *CheckDomainSsoConfigCommand) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


