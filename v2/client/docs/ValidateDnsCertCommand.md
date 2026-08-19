# ValidateDnsCertCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsProvider** | **string** |  | 
**CertIssuerType** | **string** |  | 
**Domain** | **string** |  | 
**ApiToken** | **NullableString** |  | 
**ZoneId** | **NullableString** |  | 
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
**AuthMethod** | Pointer to **NullableString** |  | [optional] 
**AuthVariant** | Pointer to **NullableString** |  | [optional] 
**AwsRoleArn** | Pointer to **NullableString** |  | [optional] 
**GcpServiceAccountEmail** | Pointer to **NullableString** |  | [optional] 
**PdnsApiUrl** | Pointer to **NullableString** |  | [optional] 
**PdnsApiKey** | Pointer to **NullableString** |  | [optional] 
**PdnsServerId** | Pointer to **NullableString** |  | [optional] 
**DnsCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**CertificateProfileId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewValidateDnsCertCommand

`func NewValidateDnsCertCommand(dnsProvider string, certIssuerType string, domain string, apiToken NullableString, zoneId NullableString, accessKeyId NullableString, secretAccessKey NullableString, region NullableString, hostedZoneId NullableString, tsigKeyName NullableString, tsigSecret NullableString, tsigAlgorithm NullableString, nameServer NullableString, azureTenantId NullableString, azureSubscriptionId NullableString, azureResourceGroup NullableString, azureClientId NullableString, azureClientSecret NullableString, gcpProject NullableString, gcpServiceAccountJson NullableString, ) *ValidateDnsCertCommand`

NewValidateDnsCertCommand instantiates a new ValidateDnsCertCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateDnsCertCommandWithDefaults

`func NewValidateDnsCertCommandWithDefaults() *ValidateDnsCertCommand`

NewValidateDnsCertCommandWithDefaults instantiates a new ValidateDnsCertCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsProvider

`func (o *ValidateDnsCertCommand) GetDnsProvider() string`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *ValidateDnsCertCommand) GetDnsProviderOk() (*string, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *ValidateDnsCertCommand) SetDnsProvider(v string)`

SetDnsProvider sets DnsProvider field to given value.


### GetCertIssuerType

`func (o *ValidateDnsCertCommand) GetCertIssuerType() string`

GetCertIssuerType returns the CertIssuerType field if non-nil, zero value otherwise.

### GetCertIssuerTypeOk

`func (o *ValidateDnsCertCommand) GetCertIssuerTypeOk() (*string, bool)`

GetCertIssuerTypeOk returns a tuple with the CertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerType

`func (o *ValidateDnsCertCommand) SetCertIssuerType(v string)`

SetCertIssuerType sets CertIssuerType field to given value.


### GetDomain

`func (o *ValidateDnsCertCommand) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ValidateDnsCertCommand) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ValidateDnsCertCommand) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetApiToken

`func (o *ValidateDnsCertCommand) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ValidateDnsCertCommand) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ValidateDnsCertCommand) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### SetApiTokenNil

`func (o *ValidateDnsCertCommand) SetApiTokenNil(b bool)`

 SetApiTokenNil sets the value for ApiToken to be an explicit nil

### UnsetApiToken
`func (o *ValidateDnsCertCommand) UnsetApiToken()`

UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
### GetZoneId

`func (o *ValidateDnsCertCommand) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ValidateDnsCertCommand) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ValidateDnsCertCommand) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### SetZoneIdNil

`func (o *ValidateDnsCertCommand) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ValidateDnsCertCommand) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetAccessKeyId

`func (o *ValidateDnsCertCommand) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ValidateDnsCertCommand) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ValidateDnsCertCommand) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *ValidateDnsCertCommand) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *ValidateDnsCertCommand) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetSecretAccessKey

`func (o *ValidateDnsCertCommand) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ValidateDnsCertCommand) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ValidateDnsCertCommand) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### SetSecretAccessKeyNil

`func (o *ValidateDnsCertCommand) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *ValidateDnsCertCommand) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetRegion

`func (o *ValidateDnsCertCommand) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ValidateDnsCertCommand) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ValidateDnsCertCommand) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *ValidateDnsCertCommand) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ValidateDnsCertCommand) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetHostedZoneId

`func (o *ValidateDnsCertCommand) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ValidateDnsCertCommand) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ValidateDnsCertCommand) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.


### SetHostedZoneIdNil

`func (o *ValidateDnsCertCommand) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *ValidateDnsCertCommand) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil
### GetTsigKeyName

`func (o *ValidateDnsCertCommand) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *ValidateDnsCertCommand) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *ValidateDnsCertCommand) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.


### SetTsigKeyNameNil

`func (o *ValidateDnsCertCommand) SetTsigKeyNameNil(b bool)`

 SetTsigKeyNameNil sets the value for TsigKeyName to be an explicit nil

### UnsetTsigKeyName
`func (o *ValidateDnsCertCommand) UnsetTsigKeyName()`

UnsetTsigKeyName ensures that no value is present for TsigKeyName, not even an explicit nil
### GetTsigSecret

`func (o *ValidateDnsCertCommand) GetTsigSecret() string`

GetTsigSecret returns the TsigSecret field if non-nil, zero value otherwise.

### GetTsigSecretOk

`func (o *ValidateDnsCertCommand) GetTsigSecretOk() (*string, bool)`

GetTsigSecretOk returns a tuple with the TsigSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigSecret

`func (o *ValidateDnsCertCommand) SetTsigSecret(v string)`

SetTsigSecret sets TsigSecret field to given value.


### SetTsigSecretNil

`func (o *ValidateDnsCertCommand) SetTsigSecretNil(b bool)`

 SetTsigSecretNil sets the value for TsigSecret to be an explicit nil

### UnsetTsigSecret
`func (o *ValidateDnsCertCommand) UnsetTsigSecret()`

UnsetTsigSecret ensures that no value is present for TsigSecret, not even an explicit nil
### GetTsigAlgorithm

`func (o *ValidateDnsCertCommand) GetTsigAlgorithm() string`

GetTsigAlgorithm returns the TsigAlgorithm field if non-nil, zero value otherwise.

### GetTsigAlgorithmOk

`func (o *ValidateDnsCertCommand) GetTsigAlgorithmOk() (*string, bool)`

GetTsigAlgorithmOk returns a tuple with the TsigAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigAlgorithm

`func (o *ValidateDnsCertCommand) SetTsigAlgorithm(v string)`

SetTsigAlgorithm sets TsigAlgorithm field to given value.


### SetTsigAlgorithmNil

`func (o *ValidateDnsCertCommand) SetTsigAlgorithmNil(b bool)`

 SetTsigAlgorithmNil sets the value for TsigAlgorithm to be an explicit nil

### UnsetTsigAlgorithm
`func (o *ValidateDnsCertCommand) UnsetTsigAlgorithm()`

UnsetTsigAlgorithm ensures that no value is present for TsigAlgorithm, not even an explicit nil
### GetNameServer

`func (o *ValidateDnsCertCommand) GetNameServer() string`

GetNameServer returns the NameServer field if non-nil, zero value otherwise.

### GetNameServerOk

`func (o *ValidateDnsCertCommand) GetNameServerOk() (*string, bool)`

GetNameServerOk returns a tuple with the NameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServer

`func (o *ValidateDnsCertCommand) SetNameServer(v string)`

SetNameServer sets NameServer field to given value.


### SetNameServerNil

`func (o *ValidateDnsCertCommand) SetNameServerNil(b bool)`

 SetNameServerNil sets the value for NameServer to be an explicit nil

### UnsetNameServer
`func (o *ValidateDnsCertCommand) UnsetNameServer()`

UnsetNameServer ensures that no value is present for NameServer, not even an explicit nil
### GetAzureTenantId

`func (o *ValidateDnsCertCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *ValidateDnsCertCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *ValidateDnsCertCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *ValidateDnsCertCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *ValidateDnsCertCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureSubscriptionId

`func (o *ValidateDnsCertCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ValidateDnsCertCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ValidateDnsCertCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### SetAzureSubscriptionIdNil

`func (o *ValidateDnsCertCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *ValidateDnsCertCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetAzureResourceGroup

`func (o *ValidateDnsCertCommand) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *ValidateDnsCertCommand) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *ValidateDnsCertCommand) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.


### SetAzureResourceGroupNil

`func (o *ValidateDnsCertCommand) SetAzureResourceGroupNil(b bool)`

 SetAzureResourceGroupNil sets the value for AzureResourceGroup to be an explicit nil

### UnsetAzureResourceGroup
`func (o *ValidateDnsCertCommand) UnsetAzureResourceGroup()`

UnsetAzureResourceGroup ensures that no value is present for AzureResourceGroup, not even an explicit nil
### GetAzureClientId

`func (o *ValidateDnsCertCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *ValidateDnsCertCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *ValidateDnsCertCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *ValidateDnsCertCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *ValidateDnsCertCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *ValidateDnsCertCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *ValidateDnsCertCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *ValidateDnsCertCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *ValidateDnsCertCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *ValidateDnsCertCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetGcpProject

`func (o *ValidateDnsCertCommand) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *ValidateDnsCertCommand) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *ValidateDnsCertCommand) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.


### SetGcpProjectNil

`func (o *ValidateDnsCertCommand) SetGcpProjectNil(b bool)`

 SetGcpProjectNil sets the value for GcpProject to be an explicit nil

### UnsetGcpProject
`func (o *ValidateDnsCertCommand) UnsetGcpProject()`

UnsetGcpProject ensures that no value is present for GcpProject, not even an explicit nil
### GetGcpServiceAccountJson

`func (o *ValidateDnsCertCommand) GetGcpServiceAccountJson() string`

GetGcpServiceAccountJson returns the GcpServiceAccountJson field if non-nil, zero value otherwise.

### GetGcpServiceAccountJsonOk

`func (o *ValidateDnsCertCommand) GetGcpServiceAccountJsonOk() (*string, bool)`

GetGcpServiceAccountJsonOk returns a tuple with the GcpServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountJson

`func (o *ValidateDnsCertCommand) SetGcpServiceAccountJson(v string)`

SetGcpServiceAccountJson sets GcpServiceAccountJson field to given value.


### SetGcpServiceAccountJsonNil

`func (o *ValidateDnsCertCommand) SetGcpServiceAccountJsonNil(b bool)`

 SetGcpServiceAccountJsonNil sets the value for GcpServiceAccountJson to be an explicit nil

### UnsetGcpServiceAccountJson
`func (o *ValidateDnsCertCommand) UnsetGcpServiceAccountJson()`

UnsetGcpServiceAccountJson ensures that no value is present for GcpServiceAccountJson, not even an explicit nil
### GetAuthMethod

`func (o *ValidateDnsCertCommand) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *ValidateDnsCertCommand) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *ValidateDnsCertCommand) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *ValidateDnsCertCommand) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *ValidateDnsCertCommand) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *ValidateDnsCertCommand) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetAuthVariant

`func (o *ValidateDnsCertCommand) GetAuthVariant() string`

GetAuthVariant returns the AuthVariant field if non-nil, zero value otherwise.

### GetAuthVariantOk

`func (o *ValidateDnsCertCommand) GetAuthVariantOk() (*string, bool)`

GetAuthVariantOk returns a tuple with the AuthVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVariant

`func (o *ValidateDnsCertCommand) SetAuthVariant(v string)`

SetAuthVariant sets AuthVariant field to given value.

### HasAuthVariant

`func (o *ValidateDnsCertCommand) HasAuthVariant() bool`

HasAuthVariant returns a boolean if a field has been set.

### SetAuthVariantNil

`func (o *ValidateDnsCertCommand) SetAuthVariantNil(b bool)`

 SetAuthVariantNil sets the value for AuthVariant to be an explicit nil

### UnsetAuthVariant
`func (o *ValidateDnsCertCommand) UnsetAuthVariant()`

UnsetAuthVariant ensures that no value is present for AuthVariant, not even an explicit nil
### GetAwsRoleArn

`func (o *ValidateDnsCertCommand) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *ValidateDnsCertCommand) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *ValidateDnsCertCommand) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *ValidateDnsCertCommand) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### SetAwsRoleArnNil

`func (o *ValidateDnsCertCommand) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *ValidateDnsCertCommand) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetGcpServiceAccountEmail

`func (o *ValidateDnsCertCommand) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *ValidateDnsCertCommand) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *ValidateDnsCertCommand) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *ValidateDnsCertCommand) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### SetGcpServiceAccountEmailNil

`func (o *ValidateDnsCertCommand) SetGcpServiceAccountEmailNil(b bool)`

 SetGcpServiceAccountEmailNil sets the value for GcpServiceAccountEmail to be an explicit nil

### UnsetGcpServiceAccountEmail
`func (o *ValidateDnsCertCommand) UnsetGcpServiceAccountEmail()`

UnsetGcpServiceAccountEmail ensures that no value is present for GcpServiceAccountEmail, not even an explicit nil
### GetPdnsApiUrl

`func (o *ValidateDnsCertCommand) GetPdnsApiUrl() string`

GetPdnsApiUrl returns the PdnsApiUrl field if non-nil, zero value otherwise.

### GetPdnsApiUrlOk

`func (o *ValidateDnsCertCommand) GetPdnsApiUrlOk() (*string, bool)`

GetPdnsApiUrlOk returns a tuple with the PdnsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiUrl

`func (o *ValidateDnsCertCommand) SetPdnsApiUrl(v string)`

SetPdnsApiUrl sets PdnsApiUrl field to given value.

### HasPdnsApiUrl

`func (o *ValidateDnsCertCommand) HasPdnsApiUrl() bool`

HasPdnsApiUrl returns a boolean if a field has been set.

### SetPdnsApiUrlNil

`func (o *ValidateDnsCertCommand) SetPdnsApiUrlNil(b bool)`

 SetPdnsApiUrlNil sets the value for PdnsApiUrl to be an explicit nil

### UnsetPdnsApiUrl
`func (o *ValidateDnsCertCommand) UnsetPdnsApiUrl()`

UnsetPdnsApiUrl ensures that no value is present for PdnsApiUrl, not even an explicit nil
### GetPdnsApiKey

`func (o *ValidateDnsCertCommand) GetPdnsApiKey() string`

GetPdnsApiKey returns the PdnsApiKey field if non-nil, zero value otherwise.

### GetPdnsApiKeyOk

`func (o *ValidateDnsCertCommand) GetPdnsApiKeyOk() (*string, bool)`

GetPdnsApiKeyOk returns a tuple with the PdnsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsApiKey

`func (o *ValidateDnsCertCommand) SetPdnsApiKey(v string)`

SetPdnsApiKey sets PdnsApiKey field to given value.

### HasPdnsApiKey

`func (o *ValidateDnsCertCommand) HasPdnsApiKey() bool`

HasPdnsApiKey returns a boolean if a field has been set.

### SetPdnsApiKeyNil

`func (o *ValidateDnsCertCommand) SetPdnsApiKeyNil(b bool)`

 SetPdnsApiKeyNil sets the value for PdnsApiKey to be an explicit nil

### UnsetPdnsApiKey
`func (o *ValidateDnsCertCommand) UnsetPdnsApiKey()`

UnsetPdnsApiKey ensures that no value is present for PdnsApiKey, not even an explicit nil
### GetPdnsServerId

`func (o *ValidateDnsCertCommand) GetPdnsServerId() string`

GetPdnsServerId returns the PdnsServerId field if non-nil, zero value otherwise.

### GetPdnsServerIdOk

`func (o *ValidateDnsCertCommand) GetPdnsServerIdOk() (*string, bool)`

GetPdnsServerIdOk returns a tuple with the PdnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnsServerId

`func (o *ValidateDnsCertCommand) SetPdnsServerId(v string)`

SetPdnsServerId sets PdnsServerId field to given value.

### HasPdnsServerId

`func (o *ValidateDnsCertCommand) HasPdnsServerId() bool`

HasPdnsServerId returns a boolean if a field has been set.

### SetPdnsServerIdNil

`func (o *ValidateDnsCertCommand) SetPdnsServerIdNil(b bool)`

 SetPdnsServerIdNil sets the value for PdnsServerId to be an explicit nil

### UnsetPdnsServerId
`func (o *ValidateDnsCertCommand) UnsetPdnsServerId()`

UnsetPdnsServerId ensures that no value is present for PdnsServerId, not even an explicit nil
### GetDnsCredentialId

`func (o *ValidateDnsCertCommand) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *ValidateDnsCertCommand) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *ValidateDnsCertCommand) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.

### HasDnsCredentialId

`func (o *ValidateDnsCertCommand) HasDnsCredentialId() bool`

HasDnsCredentialId returns a boolean if a field has been set.

### SetDnsCredentialIdNil

`func (o *ValidateDnsCertCommand) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *ValidateDnsCertCommand) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetCertificateProfileId

`func (o *ValidateDnsCertCommand) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *ValidateDnsCertCommand) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *ValidateDnsCertCommand) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.

### HasCertificateProfileId

`func (o *ValidateDnsCertCommand) HasCertificateProfileId() bool`

HasCertificateProfileId returns a boolean if a field has been set.

### SetCertificateProfileIdNil

`func (o *ValidateDnsCertCommand) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *ValidateDnsCertCommand) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


