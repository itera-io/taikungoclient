# DnsCredentialCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ProviderType** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**DomainFilter** | **string** |  | 
**UseLetsEncrypt** | **bool** |  | 
**ApiToken** | **NullableString** |  | 
**AccessKeyId** | **NullableString** |  | 
**SecretAccessKey** | **NullableString** |  | 
**Region** | **NullableString** |  | 
**HostedZoneId** | **NullableString** |  | 
**TsigKeyName** | **NullableString** |  | 
**TsigSecret** | **NullableString** |  | 
**TsigAlgorithm** | **NullableString** |  | 
**NameServer** | **NullableString** |  | 
**AzureTenantId** | **NullableString** |  | 
**AzureSubscriptionId** | **NullableString** |  | 
**AzureResourceGroup** | **NullableString** |  | 
**AzureClientId** | **NullableString** |  | 
**AzureClientSecret** | **NullableString** |  | 
**GcpProject** | **NullableString** |  | 
**GcpServiceAccountJson** | **NullableString** |  | 
**AuthMethod** | **NullableString** |  | 
**AuthVariant** | **NullableString** |  | 
**AwsRoleArn** | **NullableString** |  | 
**GcpServiceAccountEmail** | **NullableString** |  | 
**PdnsApiUrl** | Pointer to **NullableString** |  | [optional] 
**PdnsApiKey** | Pointer to **NullableString** |  | [optional] 
**PdnsServerId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsCredentialCreateCommand

`func NewDnsCredentialCreateCommand(name string, providerType string, organizationId NullableInt32, domainFilter string, useLetsEncrypt bool, apiToken NullableString, accessKeyId NullableString, secretAccessKey NullableString, region NullableString, hostedZoneId NullableString, tsigKeyName NullableString, tsigSecret NullableString, tsigAlgorithm NullableString, nameServer NullableString, azureTenantId NullableString, azureSubscriptionId NullableString, azureResourceGroup NullableString, azureClientId NullableString, azureClientSecret NullableString, gcpProject NullableString, gcpServiceAccountJson NullableString, authMethod NullableString, authVariant NullableString, awsRoleArn NullableString, gcpServiceAccountEmail NullableString, ) *DnsCredentialCreateCommand`

NewDnsCredentialCreateCommand instantiates a new DnsCredentialCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCredentialCreateCommandWithDefaults

`func NewDnsCredentialCreateCommandWithDefaults() *DnsCredentialCreateCommand`

NewDnsCredentialCreateCommandWithDefaults instantiates a new DnsCredentialCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DnsCredentialCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCredentialCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCredentialCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *DnsCredentialCreateCommand) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DnsCredentialCreateCommand) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DnsCredentialCreateCommand) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetOrganizationId

`func (o *DnsCredentialCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DnsCredentialCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DnsCredentialCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *DnsCredentialCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DnsCredentialCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetDomainFilter

`func (o *DnsCredentialCreateCommand) GetDomainFilter() string`

GetDomainFilter returns the DomainFilter field if non-nil, zero value otherwise.

### GetDomainFilterOk

`func (o *DnsCredentialCreateCommand) GetDomainFilterOk() (*string, bool)`

GetDomainFilterOk returns a tuple with the DomainFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainFilter

`func (o *DnsCredentialCreateCommand) SetDomainFilter(v string)`

SetDomainFilter sets DomainFilter field to given value.


### GetUseLetsEncrypt

`func (o *DnsCredentialCreateCommand) GetUseLetsEncrypt() bool`

GetUseLetsEncrypt returns the UseLetsEncrypt field if non-nil, zero value otherwise.

### GetUseLetsEncryptOk

`func (o *DnsCredentialCreateCommand) GetUseLetsEncryptOk() (*bool, bool)`

GetUseLetsEncryptOk returns a tuple with the UseLetsEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLetsEncrypt

`func (o *DnsCredentialCreateCommand) SetUseLetsEncrypt(v bool)`

SetUseLetsEncrypt sets UseLetsEncrypt field to given value.


### GetApiToken

`func (o *DnsCredentialCreateCommand) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *DnsCredentialCreateCommand) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *DnsCredentialCreateCommand) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### SetApiTokenNil

`func (o *DnsCredentialCreateCommand) SetApiTokenNil(b bool)`

 SetApiTokenNil sets the value for ApiToken to be an explicit nil

### UnsetApiToken
`func (o *DnsCredentialCreateCommand) UnsetApiToken()`

UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
### GetAccessKeyId

`func (o *DnsCredentialCreateCommand) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *DnsCredentialCreateCommand) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *DnsCredentialCreateCommand) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *DnsCredentialCreateCommand) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *DnsCredentialCreateCommand) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetSecretAccessKey

`func (o *DnsCredentialCreateCommand) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *DnsCredentialCreateCommand) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *DnsCredentialCreateCommand) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### SetSecretAccessKeyNil

`func (o *DnsCredentialCreateCommand) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *DnsCredentialCreateCommand) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetRegion

`func (o *DnsCredentialCreateCommand) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DnsCredentialCreateCommand) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DnsCredentialCreateCommand) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *DnsCredentialCreateCommand) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *DnsCredentialCreateCommand) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetHostedZoneId

`func (o *DnsCredentialCreateCommand) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *DnsCredentialCreateCommand) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *DnsCredentialCreateCommand) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.


### SetHostedZoneIdNil

`func (o *DnsCredentialCreateCommand) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *DnsCredentialCreateCommand) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil
### GetTsigKeyName

`func (o *DnsCredentialCreateCommand) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *DnsCredentialCreateCommand) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *DnsCredentialCreateCommand) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.


### SetTsigKeyNameNil

`func (o *DnsCredentialCreateCommand) SetTsigKeyNameNil(b bool)`

 SetTsigKeyNameNil sets the value for TsigKeyName to be an explicit nil

### UnsetTsigKeyName
`func (o *DnsCredentialCreateCommand) UnsetTsigKeyName()`

UnsetTsigKeyName ensures that no value is present for TsigKeyName, not even an explicit nil
### GetTsigSecret

`func (o *DnsCredentialCreateCommand) GetTsigSecret() string`

GetTsigSecret returns the TsigSecret field if non-nil, zero value otherwise.

### GetTsigSecretOk

`func (o *DnsCredentialCreateCommand) GetTsigSecretOk() (*string, bool)`

GetTsigSecretOk returns a tuple with the TsigSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigSecret

`func (o *DnsCredentialCreateCommand) SetTsigSecret(v string)`

SetTsigSecret sets TsigSecret field to given value.


### SetTsigSecretNil

`func (o *DnsCredentialCreateCommand) SetTsigSecretNil(b bool)`

 SetTsigSecretNil sets the value for TsigSecret to be an explicit nil

### UnsetTsigSecret
`func (o *DnsCredentialCreateCommand) UnsetTsigSecret()`

UnsetTsigSecret ensures that no value is present for TsigSecret, not even an explicit nil
### GetTsigAlgorithm

`func (o *DnsCredentialCreateCommand) GetTsigAlgorithm() string`

GetTsigAlgorithm returns the TsigAlgorithm field if non-nil, zero value otherwise.

### GetTsigAlgorithmOk

`func (o *DnsCredentialCreateCommand) GetTsigAlgorithmOk() (*string, bool)`

GetTsigAlgorithmOk returns a tuple with the TsigAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigAlgorithm

`func (o *DnsCredentialCreateCommand) SetTsigAlgorithm(v string)`

SetTsigAlgorithm sets TsigAlgorithm field to given value.


### SetTsigAlgorithmNil

`func (o *DnsCredentialCreateCommand) SetTsigAlgorithmNil(b bool)`

 SetTsigAlgorithmNil sets the value for TsigAlgorithm to be an explicit nil

### UnsetTsigAlgorithm
`func (o *DnsCredentialCreateCommand) UnsetTsigAlgorithm()`

UnsetTsigAlgorithm ensures that no value is present for TsigAlgorithm, not even an explicit nil
### GetNameServer

`func (o *DnsCredentialCreateCommand) GetNameServer() string`

GetNameServer returns the NameServer field if non-nil, zero value otherwise.

### GetNameServerOk

`func (o *DnsCredentialCreateCommand) GetNameServerOk() (*string, bool)`

GetNameServerOk returns a tuple with the NameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServer

`func (o *DnsCredentialCreateCommand) SetNameServer(v string)`

SetNameServer sets NameServer field to given value.


### SetNameServerNil

`func (o *DnsCredentialCreateCommand) SetNameServerNil(b bool)`

 SetNameServerNil sets the value for NameServer to be an explicit nil

### UnsetNameServer
`func (o *DnsCredentialCreateCommand) UnsetNameServer()`

UnsetNameServer ensures that no value is present for NameServer, not even an explicit nil
### GetAzureTenantId

`func (o *DnsCredentialCreateCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DnsCredentialCreateCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DnsCredentialCreateCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *DnsCredentialCreateCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *DnsCredentialCreateCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureSubscriptionId

`func (o *DnsCredentialCreateCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *DnsCredentialCreateCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *DnsCredentialCreateCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### SetAzureSubscriptionIdNil

`func (o *DnsCredentialCreateCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *DnsCredentialCreateCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetAzureResourceGroup

`func (o *DnsCredentialCreateCommand) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *DnsCredentialCreateCommand) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *DnsCredentialCreateCommand) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.


### SetAzureResourceGroupNil

`func (o *DnsCredentialCreateCommand) SetAzureResourceGroupNil(b bool)`

 SetAzureResourceGroupNil sets the value for AzureResourceGroup to be an explicit nil

### UnsetAzureResourceGroup
`func (o *DnsCredentialCreateCommand) UnsetAzureResourceGroup()`

UnsetAzureResourceGroup ensures that no value is present for AzureResourceGroup, not even an explicit nil
### GetAzureClientId

`func (o *DnsCredentialCreateCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *DnsCredentialCreateCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *DnsCredentialCreateCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *DnsCredentialCreateCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *DnsCredentialCreateCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *DnsCredentialCreateCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *DnsCredentialCreateCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *DnsCredentialCreateCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *DnsCredentialCreateCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *DnsCredentialCreateCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetGcpProject

`func (o *DnsCredentialCreateCommand) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *DnsCredentialCreateCommand) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *DnsCredentialCreateCommand) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.


### SetGcpProjectNil

`func (o *DnsCredentialCreateCommand) SetGcpProjectNil(b bool)`

 SetGcpProjectNil sets the value for GcpProject to be an explicit nil

### UnsetGcpProject
`func (o *DnsCredentialCreateCommand) UnsetGcpProject()`

UnsetGcpProject ensures that no value is present for GcpProject, not even an explicit nil
### GetGcpServiceAccountJson

`func (o *DnsCredentialCreateCommand) GetGcpServiceAccountJson() string`

GetGcpServiceAccountJson returns the GcpServiceAccountJson field if non-nil, zero value otherwise.

### GetGcpServiceAccountJsonOk

`func (o *DnsCredentialCreateCommand) GetGcpServiceAccountJsonOk() (*string, bool)`

GetGcpServiceAccountJsonOk returns a tuple with the GcpServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountJson

`func (o *DnsCredentialCreateCommand) SetGcpServiceAccountJson(v string)`

SetGcpServiceAccountJson sets GcpServiceAccountJson field to given value.


### SetGcpServiceAccountJsonNil

`func (o *DnsCredentialCreateCommand) SetGcpServiceAccountJsonNil(b bool)`

 SetGcpServiceAccountJsonNil sets the value for GcpServiceAccountJson to be an explicit nil

### UnsetGcpServiceAccountJson
`func (o *DnsCredentialCreateCommand) UnsetGcpServiceAccountJson()`

UnsetGcpServiceAccountJson ensures that no value is present for GcpServiceAccountJson, not even an explicit nil
### GetAuthMethod

`func (o *DnsCredentialCreateCommand) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *DnsCredentialCreateCommand) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *DnsCredentialCreateCommand) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### SetAuthMethodNil

`func (o *DnsCredentialCreateCommand) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *DnsCredentialCreateCommand) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetAuthVariant

`func (o *DnsCredentialCreateCommand) GetAuthVariant() string`

GetAuthVariant returns the AuthVariant field if non-nil, zero value otherwise.

### GetAuthVariantOk

`func (o *DnsCredentialCreateCommand) GetAuthVariantOk() (*string, bool)`

GetAuthVariantOk returns a tuple with the AuthVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVariant

`func (o *DnsCredentialCreateCommand) SetAuthVariant(v string)`

SetAuthVariant sets AuthVariant field to given value.


### SetAuthVariantNil

`func (o *DnsCredentialCreateCommand) SetAuthVariantNil(b bool)`

 SetAuthVariantNil sets the value for AuthVariant to be an explicit nil

### UnsetAuthVariant
`func (o *DnsCredentialCreateCommand) UnsetAuthVariant()`

UnsetAuthVariant ensures that no value is present for AuthVariant, not even an explicit nil
### GetAwsRoleArn

`func (o *DnsCredentialCreateCommand) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *DnsCredentialCreateCommand) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *DnsCredentialCreateCommand) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.


### SetAwsRoleArnNil

`func (o *DnsCredentialCreateCommand) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *DnsCredentialCreateCommand) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetGcpServiceAccountEmail

`func (o *DnsCredentialCreateCommand) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *DnsCredentialCreateCommand) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *DnsCredentialCreateCommand) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.


### SetGcpServiceAccountEmailNil

`func (o *DnsCredentialCreateCommand) SetGcpServiceAccountEmailNil(b bool)`

 SetGcpServiceAccountEmailNil sets the value for GcpServiceAccountEmail to be an explicit nil

### UnsetGcpServiceAccountEmail
`func (o *DnsCredentialCreateCommand) UnsetGcpServiceAccountEmail()`

UnsetGcpServiceAccountEmail ensures that no value is present for GcpServiceAccountEmail, not even an explicit nil
### GetPdnsApiUrl

`func (o *DnsCredentialCreateCommand) GetPdnsApiUrl() string`

GetPdnsApiUrl returns the PdnsApiUrl field if non-nil, zero value otherwise.

### GetPdnsApiUrlOk

`func (o *DnsCredentialCreateCommand) GetPdnsApiUrlOk() (*string, bool)`

GetPdnsApiUrlOk returns a tuple with the PdnsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiUrl

`func (o *DnsCredentialCreateCommand) SetPdnsApiUrl(v string)`

SetPdnsApiUrl sets PdnsApiUrl field to given value.

### HasPdnsApiUrl

`func (o *DnsCredentialCreateCommand) HasPdnsApiUrl() bool`

HasPdnsApiUrl returns a boolean if a field has been set.

### SetPdnsApiUrlNil

`func (o *DnsCredentialCreateCommand) SetPdnsApiUrlNil(b bool)`

 SetPdnsApiUrlNil sets the value for PdnsApiUrl to be an explicit nil

### UnsetPdnsApiUrl
`func (o *DnsCredentialCreateCommand) UnsetPdnsApiUrl()`

UnsetPdnsApiUrl ensures that no value is present for PdnsApiUrl, not even an explicit nil
### GetPdnsApiKey

`func (o *DnsCredentialCreateCommand) GetPdnsApiKey() string`

GetPdnsApiKey returns the PdnsApiKey field if non-nil, zero value otherwise.

### GetPdnsApiKeyOk

`func (o *DnsCredentialCreateCommand) GetPdnsApiKeyOk() (*string, bool)`

GetPdnsApiKeyOk returns a tuple with the PdnsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiKey

`func (o *DnsCredentialCreateCommand) SetPdnsApiKey(v string)`

SetPdnsApiKey sets PdnsApiKey field to given value.

### HasPdnsApiKey

`func (o *DnsCredentialCreateCommand) HasPdnsApiKey() bool`

HasPdnsApiKey returns a boolean if a field has been set.

### SetPdnsApiKeyNil

`func (o *DnsCredentialCreateCommand) SetPdnsApiKeyNil(b bool)`

 SetPdnsApiKeyNil sets the value for PdnsApiKey to be an explicit nil

### UnsetPdnsApiKey
`func (o *DnsCredentialCreateCommand) UnsetPdnsApiKey()`

UnsetPdnsApiKey ensures that no value is present for PdnsApiKey, not even an explicit nil
### GetPdnsServerId

`func (o *DnsCredentialCreateCommand) GetPdnsServerId() string`

GetPdnsServerId returns the PdnsServerId field if non-nil, zero value otherwise.

### GetPdnsServerIdOk

`func (o *DnsCredentialCreateCommand) GetPdnsServerIdOk() (*string, bool)`

GetPdnsServerIdOk returns a tuple with the PdnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsServerId

`func (o *DnsCredentialCreateCommand) SetPdnsServerId(v string)`

SetPdnsServerId sets PdnsServerId field to given value.

### HasPdnsServerId

`func (o *DnsCredentialCreateCommand) HasPdnsServerId() bool`

HasPdnsServerId returns a boolean if a field has been set.

### SetPdnsServerIdNil

`func (o *DnsCredentialCreateCommand) SetPdnsServerIdNil(b bool)`

 SetPdnsServerIdNil sets the value for PdnsServerId to be an explicit nil

### UnsetPdnsServerId
`func (o *DnsCredentialCreateCommand) UnsetPdnsServerId()`

UnsetPdnsServerId ensures that no value is present for PdnsServerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


