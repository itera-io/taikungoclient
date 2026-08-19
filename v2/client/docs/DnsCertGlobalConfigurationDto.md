# DnsCertGlobalConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**ExternalDnsChartName** | Pointer to **NullableString** |  | [optional] 
**ExternalDnsRepoUrl** | Pointer to **NullableString** |  | [optional] 
**ExternalDnsVersion** | Pointer to **NullableString** |  | [optional] 
**CertManagerChartName** | Pointer to **NullableString** |  | [optional] 
**CertManagerRepoUrl** | Pointer to **NullableString** |  | [optional] 
**CertManagerVersion** | Pointer to **NullableString** |  | [optional] 
**DefaultCertIssuer** | Pointer to **NullableString** |  | [optional] 
**DefaultAcmeServer** | Pointer to **NullableString** |  | [optional] 
**SupportedProviders** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsCertGlobalConfigurationDto

`func NewDnsCertGlobalConfigurationDto() *DnsCertGlobalConfigurationDto`

NewDnsCertGlobalConfigurationDto instantiates a new DnsCertGlobalConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCertGlobalConfigurationDtoWithDefaults

`func NewDnsCertGlobalConfigurationDtoWithDefaults() *DnsCertGlobalConfigurationDto`

NewDnsCertGlobalConfigurationDtoWithDefaults instantiates a new DnsCertGlobalConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsCertGlobalConfigurationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsCertGlobalConfigurationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsCertGlobalConfigurationDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DnsCertGlobalConfigurationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsCertGlobalConfigurationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCertGlobalConfigurationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCertGlobalConfigurationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsCertGlobalConfigurationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DnsCertGlobalConfigurationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DnsCertGlobalConfigurationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *DnsCertGlobalConfigurationDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DnsCertGlobalConfigurationDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DnsCertGlobalConfigurationDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DnsCertGlobalConfigurationDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetExternalDnsChartName

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsChartName() string`

GetExternalDnsChartName returns the ExternalDnsChartName field if non-nil, zero value otherwise.

### GetExternalDnsChartNameOk

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsChartNameOk() (*string, bool)`

GetExternalDnsChartNameOk returns a tuple with the ExternalDnsChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsChartName

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsChartName(v string)`

SetExternalDnsChartName sets ExternalDnsChartName field to given value.

### HasExternalDnsChartName

`func (o *DnsCertGlobalConfigurationDto) HasExternalDnsChartName() bool`

HasExternalDnsChartName returns a boolean if a field has been set.

### SetExternalDnsChartNameNil

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsChartNameNil(b bool)`

 SetExternalDnsChartNameNil sets the value for ExternalDnsChartName to be an explicit nil

### UnsetExternalDnsChartName
`func (o *DnsCertGlobalConfigurationDto) UnsetExternalDnsChartName()`

UnsetExternalDnsChartName ensures that no value is present for ExternalDnsChartName, not even an explicit nil
### GetExternalDnsRepoUrl

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsRepoUrl() string`

GetExternalDnsRepoUrl returns the ExternalDnsRepoUrl field if non-nil, zero value otherwise.

### GetExternalDnsRepoUrlOk

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsRepoUrlOk() (*string, bool)`

GetExternalDnsRepoUrlOk returns a tuple with the ExternalDnsRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsRepoUrl

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsRepoUrl(v string)`

SetExternalDnsRepoUrl sets ExternalDnsRepoUrl field to given value.

### HasExternalDnsRepoUrl

`func (o *DnsCertGlobalConfigurationDto) HasExternalDnsRepoUrl() bool`

HasExternalDnsRepoUrl returns a boolean if a field has been set.

### SetExternalDnsRepoUrlNil

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsRepoUrlNil(b bool)`

 SetExternalDnsRepoUrlNil sets the value for ExternalDnsRepoUrl to be an explicit nil

### UnsetExternalDnsRepoUrl
`func (o *DnsCertGlobalConfigurationDto) UnsetExternalDnsRepoUrl()`

UnsetExternalDnsRepoUrl ensures that no value is present for ExternalDnsRepoUrl, not even an explicit nil
### GetExternalDnsVersion

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsVersion() string`

GetExternalDnsVersion returns the ExternalDnsVersion field if non-nil, zero value otherwise.

### GetExternalDnsVersionOk

`func (o *DnsCertGlobalConfigurationDto) GetExternalDnsVersionOk() (*string, bool)`

GetExternalDnsVersionOk returns a tuple with the ExternalDnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsVersion

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsVersion(v string)`

SetExternalDnsVersion sets ExternalDnsVersion field to given value.

### HasExternalDnsVersion

`func (o *DnsCertGlobalConfigurationDto) HasExternalDnsVersion() bool`

HasExternalDnsVersion returns a boolean if a field has been set.

### SetExternalDnsVersionNil

`func (o *DnsCertGlobalConfigurationDto) SetExternalDnsVersionNil(b bool)`

 SetExternalDnsVersionNil sets the value for ExternalDnsVersion to be an explicit nil

### UnsetExternalDnsVersion
`func (o *DnsCertGlobalConfigurationDto) UnsetExternalDnsVersion()`

UnsetExternalDnsVersion ensures that no value is present for ExternalDnsVersion, not even an explicit nil
### GetCertManagerChartName

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerChartName() string`

GetCertManagerChartName returns the CertManagerChartName field if non-nil, zero value otherwise.

### GetCertManagerChartNameOk

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerChartNameOk() (*string, bool)`

GetCertManagerChartNameOk returns a tuple with the CertManagerChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerChartName

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerChartName(v string)`

SetCertManagerChartName sets CertManagerChartName field to given value.

### HasCertManagerChartName

`func (o *DnsCertGlobalConfigurationDto) HasCertManagerChartName() bool`

HasCertManagerChartName returns a boolean if a field has been set.

### SetCertManagerChartNameNil

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerChartNameNil(b bool)`

 SetCertManagerChartNameNil sets the value for CertManagerChartName to be an explicit nil

### UnsetCertManagerChartName
`func (o *DnsCertGlobalConfigurationDto) UnsetCertManagerChartName()`

UnsetCertManagerChartName ensures that no value is present for CertManagerChartName, not even an explicit nil
### GetCertManagerRepoUrl

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerRepoUrl() string`

GetCertManagerRepoUrl returns the CertManagerRepoUrl field if non-nil, zero value otherwise.

### GetCertManagerRepoUrlOk

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerRepoUrlOk() (*string, bool)`

GetCertManagerRepoUrlOk returns a tuple with the CertManagerRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerRepoUrl

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerRepoUrl(v string)`

SetCertManagerRepoUrl sets CertManagerRepoUrl field to given value.

### HasCertManagerRepoUrl

`func (o *DnsCertGlobalConfigurationDto) HasCertManagerRepoUrl() bool`

HasCertManagerRepoUrl returns a boolean if a field has been set.

### SetCertManagerRepoUrlNil

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerRepoUrlNil(b bool)`

 SetCertManagerRepoUrlNil sets the value for CertManagerRepoUrl to be an explicit nil

### UnsetCertManagerRepoUrl
`func (o *DnsCertGlobalConfigurationDto) UnsetCertManagerRepoUrl()`

UnsetCertManagerRepoUrl ensures that no value is present for CertManagerRepoUrl, not even an explicit nil
### GetCertManagerVersion

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerVersion() string`

GetCertManagerVersion returns the CertManagerVersion field if non-nil, zero value otherwise.

### GetCertManagerVersionOk

`func (o *DnsCertGlobalConfigurationDto) GetCertManagerVersionOk() (*string, bool)`

GetCertManagerVersionOk returns a tuple with the CertManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerVersion

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerVersion(v string)`

SetCertManagerVersion sets CertManagerVersion field to given value.

### HasCertManagerVersion

`func (o *DnsCertGlobalConfigurationDto) HasCertManagerVersion() bool`

HasCertManagerVersion returns a boolean if a field has been set.

### SetCertManagerVersionNil

`func (o *DnsCertGlobalConfigurationDto) SetCertManagerVersionNil(b bool)`

 SetCertManagerVersionNil sets the value for CertManagerVersion to be an explicit nil

### UnsetCertManagerVersion
`func (o *DnsCertGlobalConfigurationDto) UnsetCertManagerVersion()`

UnsetCertManagerVersion ensures that no value is present for CertManagerVersion, not even an explicit nil
### GetDefaultCertIssuer

`func (o *DnsCertGlobalConfigurationDto) GetDefaultCertIssuer() string`

GetDefaultCertIssuer returns the DefaultCertIssuer field if non-nil, zero value otherwise.

### GetDefaultCertIssuerOk

`func (o *DnsCertGlobalConfigurationDto) GetDefaultCertIssuerOk() (*string, bool)`

GetDefaultCertIssuerOk returns a tuple with the DefaultCertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCertIssuer

`func (o *DnsCertGlobalConfigurationDto) SetDefaultCertIssuer(v string)`

SetDefaultCertIssuer sets DefaultCertIssuer field to given value.

### HasDefaultCertIssuer

`func (o *DnsCertGlobalConfigurationDto) HasDefaultCertIssuer() bool`

HasDefaultCertIssuer returns a boolean if a field has been set.

### SetDefaultCertIssuerNil

`func (o *DnsCertGlobalConfigurationDto) SetDefaultCertIssuerNil(b bool)`

 SetDefaultCertIssuerNil sets the value for DefaultCertIssuer to be an explicit nil

### UnsetDefaultCertIssuer
`func (o *DnsCertGlobalConfigurationDto) UnsetDefaultCertIssuer()`

UnsetDefaultCertIssuer ensures that no value is present for DefaultCertIssuer, not even an explicit nil
### GetDefaultAcmeServer

`func (o *DnsCertGlobalConfigurationDto) GetDefaultAcmeServer() string`

GetDefaultAcmeServer returns the DefaultAcmeServer field if non-nil, zero value otherwise.

### GetDefaultAcmeServerOk

`func (o *DnsCertGlobalConfigurationDto) GetDefaultAcmeServerOk() (*string, bool)`

GetDefaultAcmeServerOk returns a tuple with the DefaultAcmeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAcmeServer

`func (o *DnsCertGlobalConfigurationDto) SetDefaultAcmeServer(v string)`

SetDefaultAcmeServer sets DefaultAcmeServer field to given value.

### HasDefaultAcmeServer

`func (o *DnsCertGlobalConfigurationDto) HasDefaultAcmeServer() bool`

HasDefaultAcmeServer returns a boolean if a field has been set.

### SetDefaultAcmeServerNil

`func (o *DnsCertGlobalConfigurationDto) SetDefaultAcmeServerNil(b bool)`

 SetDefaultAcmeServerNil sets the value for DefaultAcmeServer to be an explicit nil

### UnsetDefaultAcmeServer
`func (o *DnsCertGlobalConfigurationDto) UnsetDefaultAcmeServer()`

UnsetDefaultAcmeServer ensures that no value is present for DefaultAcmeServer, not even an explicit nil
### GetSupportedProviders

`func (o *DnsCertGlobalConfigurationDto) GetSupportedProviders() string`

GetSupportedProviders returns the SupportedProviders field if non-nil, zero value otherwise.

### GetSupportedProvidersOk

`func (o *DnsCertGlobalConfigurationDto) GetSupportedProvidersOk() (*string, bool)`

GetSupportedProvidersOk returns a tuple with the SupportedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProviders

`func (o *DnsCertGlobalConfigurationDto) SetSupportedProviders(v string)`

SetSupportedProviders sets SupportedProviders field to given value.

### HasSupportedProviders

`func (o *DnsCertGlobalConfigurationDto) HasSupportedProviders() bool`

HasSupportedProviders returns a boolean if a field has been set.

### SetSupportedProvidersNil

`func (o *DnsCertGlobalConfigurationDto) SetSupportedProvidersNil(b bool)`

 SetSupportedProvidersNil sets the value for SupportedProviders to be an explicit nil

### UnsetSupportedProviders
`func (o *DnsCertGlobalConfigurationDto) UnsetSupportedProviders()`

UnsetSupportedProviders ensures that no value is present for SupportedProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


