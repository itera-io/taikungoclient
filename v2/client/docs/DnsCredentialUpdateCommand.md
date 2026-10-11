# DnsCredentialUpdateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ProviderType** | **string** |  | 
**DomainFilter** | **string** |  | 
**UseLetsEncrypt** | **bool** |  | 
**AuthMethod** | Pointer to **NullableString** |  | [optional] 
**AuthVariant** | Pointer to **NullableString** |  | [optional] 
**AwsRoleArn** | Pointer to **NullableString** |  | [optional] 
**GcpServiceAccountEmail** | Pointer to **NullableString** |  | [optional] 
**PdnsApiUrl** | Pointer to **NullableString** |  | [optional] 
**PdnsApiKey** | Pointer to **NullableString** |  | [optional] 
**PdnsServerId** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**HostedZoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsCredentialUpdateCommand

`func NewDnsCredentialUpdateCommand(id int32, name string, providerType string, domainFilter string, useLetsEncrypt bool, ) *DnsCredentialUpdateCommand`

NewDnsCredentialUpdateCommand instantiates a new DnsCredentialUpdateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCredentialUpdateCommandWithDefaults

`func NewDnsCredentialUpdateCommandWithDefaults() *DnsCredentialUpdateCommand`

NewDnsCredentialUpdateCommandWithDefaults instantiates a new DnsCredentialUpdateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsCredentialUpdateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsCredentialUpdateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsCredentialUpdateCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DnsCredentialUpdateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCredentialUpdateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCredentialUpdateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *DnsCredentialUpdateCommand) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DnsCredentialUpdateCommand) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DnsCredentialUpdateCommand) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetDomainFilter

`func (o *DnsCredentialUpdateCommand) GetDomainFilter() string`

GetDomainFilter returns the DomainFilter field if non-nil, zero value otherwise.

### GetDomainFilterOk

`func (o *DnsCredentialUpdateCommand) GetDomainFilterOk() (*string, bool)`

GetDomainFilterOk returns a tuple with the DomainFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainFilter

`func (o *DnsCredentialUpdateCommand) SetDomainFilter(v string)`

SetDomainFilter sets DomainFilter field to given value.


### GetUseLetsEncrypt

`func (o *DnsCredentialUpdateCommand) GetUseLetsEncrypt() bool`

GetUseLetsEncrypt returns the UseLetsEncrypt field if non-nil, zero value otherwise.

### GetUseLetsEncryptOk

`func (o *DnsCredentialUpdateCommand) GetUseLetsEncryptOk() (*bool, bool)`

GetUseLetsEncryptOk returns a tuple with the UseLetsEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLetsEncrypt

`func (o *DnsCredentialUpdateCommand) SetUseLetsEncrypt(v bool)`

SetUseLetsEncrypt sets UseLetsEncrypt field to given value.


### GetAuthMethod

`func (o *DnsCredentialUpdateCommand) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *DnsCredentialUpdateCommand) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *DnsCredentialUpdateCommand) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *DnsCredentialUpdateCommand) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *DnsCredentialUpdateCommand) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *DnsCredentialUpdateCommand) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetAuthVariant

`func (o *DnsCredentialUpdateCommand) GetAuthVariant() string`

GetAuthVariant returns the AuthVariant field if non-nil, zero value otherwise.

### GetAuthVariantOk

`func (o *DnsCredentialUpdateCommand) GetAuthVariantOk() (*string, bool)`

GetAuthVariantOk returns a tuple with the AuthVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVariant

`func (o *DnsCredentialUpdateCommand) SetAuthVariant(v string)`

SetAuthVariant sets AuthVariant field to given value.

### HasAuthVariant

`func (o *DnsCredentialUpdateCommand) HasAuthVariant() bool`

HasAuthVariant returns a boolean if a field has been set.

### SetAuthVariantNil

`func (o *DnsCredentialUpdateCommand) SetAuthVariantNil(b bool)`

 SetAuthVariantNil sets the value for AuthVariant to be an explicit nil

### UnsetAuthVariant
`func (o *DnsCredentialUpdateCommand) UnsetAuthVariant()`

UnsetAuthVariant ensures that no value is present for AuthVariant, not even an explicit nil
### GetAwsRoleArn

`func (o *DnsCredentialUpdateCommand) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *DnsCredentialUpdateCommand) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *DnsCredentialUpdateCommand) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *DnsCredentialUpdateCommand) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### SetAwsRoleArnNil

`func (o *DnsCredentialUpdateCommand) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *DnsCredentialUpdateCommand) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetGcpServiceAccountEmail

`func (o *DnsCredentialUpdateCommand) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *DnsCredentialUpdateCommand) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *DnsCredentialUpdateCommand) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *DnsCredentialUpdateCommand) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### SetGcpServiceAccountEmailNil

`func (o *DnsCredentialUpdateCommand) SetGcpServiceAccountEmailNil(b bool)`

 SetGcpServiceAccountEmailNil sets the value for GcpServiceAccountEmail to be an explicit nil

### UnsetGcpServiceAccountEmail
`func (o *DnsCredentialUpdateCommand) UnsetGcpServiceAccountEmail()`

UnsetGcpServiceAccountEmail ensures that no value is present for GcpServiceAccountEmail, not even an explicit nil
### GetPdnsApiUrl

`func (o *DnsCredentialUpdateCommand) GetPdnsApiUrl() string`

GetPdnsApiUrl returns the PdnsApiUrl field if non-nil, zero value otherwise.

### GetPdnsApiUrlOk

`func (o *DnsCredentialUpdateCommand) GetPdnsApiUrlOk() (*string, bool)`

GetPdnsApiUrlOk returns a tuple with the PdnsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiUrl

`func (o *DnsCredentialUpdateCommand) SetPdnsApiUrl(v string)`

SetPdnsApiUrl sets PdnsApiUrl field to given value.

### HasPdnsApiUrl

`func (o *DnsCredentialUpdateCommand) HasPdnsApiUrl() bool`

HasPdnsApiUrl returns a boolean if a field has been set.

### SetPdnsApiUrlNil

`func (o *DnsCredentialUpdateCommand) SetPdnsApiUrlNil(b bool)`

 SetPdnsApiUrlNil sets the value for PdnsApiUrl to be an explicit nil

### UnsetPdnsApiUrl
`func (o *DnsCredentialUpdateCommand) UnsetPdnsApiUrl()`

UnsetPdnsApiUrl ensures that no value is present for PdnsApiUrl, not even an explicit nil
### GetPdnsApiKey

`func (o *DnsCredentialUpdateCommand) GetPdnsApiKey() string`

GetPdnsApiKey returns the PdnsApiKey field if non-nil, zero value otherwise.

### GetPdnsApiKeyOk

`func (o *DnsCredentialUpdateCommand) GetPdnsApiKeyOk() (*string, bool)`

GetPdnsApiKeyOk returns a tuple with the PdnsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiKey

`func (o *DnsCredentialUpdateCommand) SetPdnsApiKey(v string)`

SetPdnsApiKey sets PdnsApiKey field to given value.

### HasPdnsApiKey

`func (o *DnsCredentialUpdateCommand) HasPdnsApiKey() bool`

HasPdnsApiKey returns a boolean if a field has been set.

### SetPdnsApiKeyNil

`func (o *DnsCredentialUpdateCommand) SetPdnsApiKeyNil(b bool)`

 SetPdnsApiKeyNil sets the value for PdnsApiKey to be an explicit nil

### UnsetPdnsApiKey
`func (o *DnsCredentialUpdateCommand) UnsetPdnsApiKey()`

UnsetPdnsApiKey ensures that no value is present for PdnsApiKey, not even an explicit nil
### GetPdnsServerId

`func (o *DnsCredentialUpdateCommand) GetPdnsServerId() string`

GetPdnsServerId returns the PdnsServerId field if non-nil, zero value otherwise.

### GetPdnsServerIdOk

`func (o *DnsCredentialUpdateCommand) GetPdnsServerIdOk() (*string, bool)`

GetPdnsServerIdOk returns a tuple with the PdnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsServerId

`func (o *DnsCredentialUpdateCommand) SetPdnsServerId(v string)`

SetPdnsServerId sets PdnsServerId field to given value.

### HasPdnsServerId

`func (o *DnsCredentialUpdateCommand) HasPdnsServerId() bool`

HasPdnsServerId returns a boolean if a field has been set.

### SetPdnsServerIdNil

`func (o *DnsCredentialUpdateCommand) SetPdnsServerIdNil(b bool)`

 SetPdnsServerIdNil sets the value for PdnsServerId to be an explicit nil

### UnsetPdnsServerId
`func (o *DnsCredentialUpdateCommand) UnsetPdnsServerId()`

UnsetPdnsServerId ensures that no value is present for PdnsServerId, not even an explicit nil
### GetRegion

`func (o *DnsCredentialUpdateCommand) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DnsCredentialUpdateCommand) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DnsCredentialUpdateCommand) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DnsCredentialUpdateCommand) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *DnsCredentialUpdateCommand) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *DnsCredentialUpdateCommand) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetHostedZoneId

`func (o *DnsCredentialUpdateCommand) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *DnsCredentialUpdateCommand) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *DnsCredentialUpdateCommand) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *DnsCredentialUpdateCommand) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### SetHostedZoneIdNil

`func (o *DnsCredentialUpdateCommand) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *DnsCredentialUpdateCommand) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


