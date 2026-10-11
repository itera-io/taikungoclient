# CreateDnsCertGlobalConfigurationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**ExternalDnsChartName** | **NullableString** |  | 
**ExternalDnsRepoUrl** | **NullableString** |  | 
**ExternalDnsVersion** | **NullableString** |  | 
**CertManagerChartName** | **NullableString** |  | 
**CertManagerRepoUrl** | **NullableString** |  | 
**CertManagerVersion** | **NullableString** |  | 
**DefaultCertIssuer** | **NullableString** |  | 
**DefaultAcmeServer** | **NullableString** |  | 
**SupportedProviders** | **NullableString** |  | 

## Methods

### NewCreateDnsCertGlobalConfigurationCommand

`func NewCreateDnsCertGlobalConfigurationCommand(name NullableString, externalDnsChartName NullableString, externalDnsRepoUrl NullableString, externalDnsVersion NullableString, certManagerChartName NullableString, certManagerRepoUrl NullableString, certManagerVersion NullableString, defaultCertIssuer NullableString, defaultAcmeServer NullableString, supportedProviders NullableString, ) *CreateDnsCertGlobalConfigurationCommand`

NewCreateDnsCertGlobalConfigurationCommand instantiates a new CreateDnsCertGlobalConfigurationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsCertGlobalConfigurationCommandWithDefaults

`func NewCreateDnsCertGlobalConfigurationCommandWithDefaults() *CreateDnsCertGlobalConfigurationCommand`

NewCreateDnsCertGlobalConfigurationCommandWithDefaults instantiates a new CreateDnsCertGlobalConfigurationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDnsCertGlobalConfigurationCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDnsCertGlobalConfigurationCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExternalDnsChartName

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsChartName() string`

GetExternalDnsChartName returns the ExternalDnsChartName field if non-nil, zero value otherwise.

### GetExternalDnsChartNameOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsChartNameOk() (*string, bool)`

GetExternalDnsChartNameOk returns a tuple with the ExternalDnsChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsChartName

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsChartName(v string)`

SetExternalDnsChartName sets ExternalDnsChartName field to given value.


### SetExternalDnsChartNameNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsChartNameNil(b bool)`

 SetExternalDnsChartNameNil sets the value for ExternalDnsChartName to be an explicit nil

### UnsetExternalDnsChartName
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetExternalDnsChartName()`

UnsetExternalDnsChartName ensures that no value is present for ExternalDnsChartName, not even an explicit nil
### GetExternalDnsRepoUrl

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsRepoUrl() string`

GetExternalDnsRepoUrl returns the ExternalDnsRepoUrl field if non-nil, zero value otherwise.

### GetExternalDnsRepoUrlOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsRepoUrlOk() (*string, bool)`

GetExternalDnsRepoUrlOk returns a tuple with the ExternalDnsRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsRepoUrl

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsRepoUrl(v string)`

SetExternalDnsRepoUrl sets ExternalDnsRepoUrl field to given value.


### SetExternalDnsRepoUrlNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsRepoUrlNil(b bool)`

 SetExternalDnsRepoUrlNil sets the value for ExternalDnsRepoUrl to be an explicit nil

### UnsetExternalDnsRepoUrl
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetExternalDnsRepoUrl()`

UnsetExternalDnsRepoUrl ensures that no value is present for ExternalDnsRepoUrl, not even an explicit nil
### GetExternalDnsVersion

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsVersion() string`

GetExternalDnsVersion returns the ExternalDnsVersion field if non-nil, zero value otherwise.

### GetExternalDnsVersionOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetExternalDnsVersionOk() (*string, bool)`

GetExternalDnsVersionOk returns a tuple with the ExternalDnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsVersion

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsVersion(v string)`

SetExternalDnsVersion sets ExternalDnsVersion field to given value.


### SetExternalDnsVersionNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetExternalDnsVersionNil(b bool)`

 SetExternalDnsVersionNil sets the value for ExternalDnsVersion to be an explicit nil

### UnsetExternalDnsVersion
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetExternalDnsVersion()`

UnsetExternalDnsVersion ensures that no value is present for ExternalDnsVersion, not even an explicit nil
### GetCertManagerChartName

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerChartName() string`

GetCertManagerChartName returns the CertManagerChartName field if non-nil, zero value otherwise.

### GetCertManagerChartNameOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerChartNameOk() (*string, bool)`

GetCertManagerChartNameOk returns a tuple with the CertManagerChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerChartName

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerChartName(v string)`

SetCertManagerChartName sets CertManagerChartName field to given value.


### SetCertManagerChartNameNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerChartNameNil(b bool)`

 SetCertManagerChartNameNil sets the value for CertManagerChartName to be an explicit nil

### UnsetCertManagerChartName
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetCertManagerChartName()`

UnsetCertManagerChartName ensures that no value is present for CertManagerChartName, not even an explicit nil
### GetCertManagerRepoUrl

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerRepoUrl() string`

GetCertManagerRepoUrl returns the CertManagerRepoUrl field if non-nil, zero value otherwise.

### GetCertManagerRepoUrlOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerRepoUrlOk() (*string, bool)`

GetCertManagerRepoUrlOk returns a tuple with the CertManagerRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerRepoUrl

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerRepoUrl(v string)`

SetCertManagerRepoUrl sets CertManagerRepoUrl field to given value.


### SetCertManagerRepoUrlNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerRepoUrlNil(b bool)`

 SetCertManagerRepoUrlNil sets the value for CertManagerRepoUrl to be an explicit nil

### UnsetCertManagerRepoUrl
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetCertManagerRepoUrl()`

UnsetCertManagerRepoUrl ensures that no value is present for CertManagerRepoUrl, not even an explicit nil
### GetCertManagerVersion

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerVersion() string`

GetCertManagerVersion returns the CertManagerVersion field if non-nil, zero value otherwise.

### GetCertManagerVersionOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetCertManagerVersionOk() (*string, bool)`

GetCertManagerVersionOk returns a tuple with the CertManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerVersion

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerVersion(v string)`

SetCertManagerVersion sets CertManagerVersion field to given value.


### SetCertManagerVersionNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetCertManagerVersionNil(b bool)`

 SetCertManagerVersionNil sets the value for CertManagerVersion to be an explicit nil

### UnsetCertManagerVersion
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetCertManagerVersion()`

UnsetCertManagerVersion ensures that no value is present for CertManagerVersion, not even an explicit nil
### GetDefaultCertIssuer

`func (o *CreateDnsCertGlobalConfigurationCommand) GetDefaultCertIssuer() string`

GetDefaultCertIssuer returns the DefaultCertIssuer field if non-nil, zero value otherwise.

### GetDefaultCertIssuerOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetDefaultCertIssuerOk() (*string, bool)`

GetDefaultCertIssuerOk returns a tuple with the DefaultCertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertIssuer

`func (o *CreateDnsCertGlobalConfigurationCommand) SetDefaultCertIssuer(v string)`

SetDefaultCertIssuer sets DefaultCertIssuer field to given value.


### SetDefaultCertIssuerNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetDefaultCertIssuerNil(b bool)`

 SetDefaultCertIssuerNil sets the value for DefaultCertIssuer to be an explicit nil

### UnsetDefaultCertIssuer
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetDefaultCertIssuer()`

UnsetDefaultCertIssuer ensures that no value is present for DefaultCertIssuer, not even an explicit nil
### GetDefaultAcmeServer

`func (o *CreateDnsCertGlobalConfigurationCommand) GetDefaultAcmeServer() string`

GetDefaultAcmeServer returns the DefaultAcmeServer field if non-nil, zero value otherwise.

### GetDefaultAcmeServerOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetDefaultAcmeServerOk() (*string, bool)`

GetDefaultAcmeServerOk returns a tuple with the DefaultAcmeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAcmeServer

`func (o *CreateDnsCertGlobalConfigurationCommand) SetDefaultAcmeServer(v string)`

SetDefaultAcmeServer sets DefaultAcmeServer field to given value.


### SetDefaultAcmeServerNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetDefaultAcmeServerNil(b bool)`

 SetDefaultAcmeServerNil sets the value for DefaultAcmeServer to be an explicit nil

### UnsetDefaultAcmeServer
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetDefaultAcmeServer()`

UnsetDefaultAcmeServer ensures that no value is present for DefaultAcmeServer, not even an explicit nil
### GetSupportedProviders

`func (o *CreateDnsCertGlobalConfigurationCommand) GetSupportedProviders() string`

GetSupportedProviders returns the SupportedProviders field if non-nil, zero value otherwise.

### GetSupportedProvidersOk

`func (o *CreateDnsCertGlobalConfigurationCommand) GetSupportedProvidersOk() (*string, bool)`

GetSupportedProvidersOk returns a tuple with the SupportedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProviders

`func (o *CreateDnsCertGlobalConfigurationCommand) SetSupportedProviders(v string)`

SetSupportedProviders sets SupportedProviders field to given value.


### SetSupportedProvidersNil

`func (o *CreateDnsCertGlobalConfigurationCommand) SetSupportedProvidersNil(b bool)`

 SetSupportedProvidersNil sets the value for SupportedProviders to be an explicit nil

### UnsetSupportedProviders
`func (o *CreateDnsCertGlobalConfigurationCommand) UnsetSupportedProviders()`

UnsetSupportedProviders ensures that no value is present for SupportedProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


