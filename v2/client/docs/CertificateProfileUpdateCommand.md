# CertificateProfileUpdateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IssuerMode** | Pointer to [**CertificateProfileIssuerMode**](CertificateProfileIssuerMode.md) |  | [optional] 
**TrustBundle** | Pointer to [**[]TrustBundleUpdateItem**](TrustBundleUpdateItem.md) |  | [optional] 
**AcmeEmail** | Pointer to **NullableString** |  | [optional] 
**AcmeServer** | Pointer to **NullableString** |  | [optional] 
**VenafiType** | Pointer to [**CertificateProfileVenafiType**](CertificateProfileVenafiType.md) |  | [optional] 
**VenafiBaseUrl** | Pointer to **NullableString** |  | [optional] 
**VenafiZone** | Pointer to **NullableString** |  | [optional] 
**VenafiAccessToken** | Pointer to **NullableString** |  | [optional] 
**VenafiApiKey** | Pointer to **NullableString** |  | [optional] 
**VenafiZoneUuid** | Pointer to **NullableString** |  | [optional] 
**CustomCaCert** | Pointer to **NullableString** |  | [optional] 
**CustomCaPrivateKey** | Pointer to **NullableString** |  | [optional] 
**ChainPem** | Pointer to **NullableString** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**ValidityDays** | Pointer to **NullableInt32** |  | [optional] 
**CertPem** | Pointer to **NullableString** |  | [optional] 
**KeyPem** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateProfileUpdateCommand

`func NewCertificateProfileUpdateCommand(id int32, name string, ) *CertificateProfileUpdateCommand`

NewCertificateProfileUpdateCommand instantiates a new CertificateProfileUpdateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileUpdateCommandWithDefaults

`func NewCertificateProfileUpdateCommandWithDefaults() *CertificateProfileUpdateCommand`

NewCertificateProfileUpdateCommandWithDefaults instantiates a new CertificateProfileUpdateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileUpdateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileUpdateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileUpdateCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateProfileUpdateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileUpdateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileUpdateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateProfileUpdateCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileUpdateCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileUpdateCommand) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileUpdateCommand) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileUpdateCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileUpdateCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIssuerMode

`func (o *CertificateProfileUpdateCommand) GetIssuerMode() CertificateProfileIssuerMode`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *CertificateProfileUpdateCommand) GetIssuerModeOk() (*CertificateProfileIssuerMode, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *CertificateProfileUpdateCommand) SetIssuerMode(v CertificateProfileIssuerMode)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *CertificateProfileUpdateCommand) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetTrustBundle

`func (o *CertificateProfileUpdateCommand) GetTrustBundle() []TrustBundleUpdateItem`

GetTrustBundle returns the TrustBundle field if non-nil, zero value otherwise.

### GetTrustBundleOk

`func (o *CertificateProfileUpdateCommand) GetTrustBundleOk() (*[]TrustBundleUpdateItem, bool)`

GetTrustBundleOk returns a tuple with the TrustBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustBundle

`func (o *CertificateProfileUpdateCommand) SetTrustBundle(v []TrustBundleUpdateItem)`

SetTrustBundle sets TrustBundle field to given value.

### HasTrustBundle

`func (o *CertificateProfileUpdateCommand) HasTrustBundle() bool`

HasTrustBundle returns a boolean if a field has been set.

### SetTrustBundleNil

`func (o *CertificateProfileUpdateCommand) SetTrustBundleNil(b bool)`

 SetTrustBundleNil sets the value for TrustBundle to be an explicit nil

### UnsetTrustBundle
`func (o *CertificateProfileUpdateCommand) UnsetTrustBundle()`

UnsetTrustBundle ensures that no value is present for TrustBundle, not even an explicit nil
### GetAcmeEmail

`func (o *CertificateProfileUpdateCommand) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *CertificateProfileUpdateCommand) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *CertificateProfileUpdateCommand) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *CertificateProfileUpdateCommand) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### SetAcmeEmailNil

`func (o *CertificateProfileUpdateCommand) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *CertificateProfileUpdateCommand) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetAcmeServer

`func (o *CertificateProfileUpdateCommand) GetAcmeServer() string`

GetAcmeServer returns the AcmeServer field if non-nil, zero value otherwise.

### GetAcmeServerOk

`func (o *CertificateProfileUpdateCommand) GetAcmeServerOk() (*string, bool)`

GetAcmeServerOk returns a tuple with the AcmeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeServer

`func (o *CertificateProfileUpdateCommand) SetAcmeServer(v string)`

SetAcmeServer sets AcmeServer field to given value.

### HasAcmeServer

`func (o *CertificateProfileUpdateCommand) HasAcmeServer() bool`

HasAcmeServer returns a boolean if a field has been set.

### SetAcmeServerNil

`func (o *CertificateProfileUpdateCommand) SetAcmeServerNil(b bool)`

 SetAcmeServerNil sets the value for AcmeServer to be an explicit nil

### UnsetAcmeServer
`func (o *CertificateProfileUpdateCommand) UnsetAcmeServer()`

UnsetAcmeServer ensures that no value is present for AcmeServer, not even an explicit nil
### GetVenafiType

`func (o *CertificateProfileUpdateCommand) GetVenafiType() CertificateProfileVenafiType`

GetVenafiType returns the VenafiType field if non-nil, zero value otherwise.

### GetVenafiTypeOk

`func (o *CertificateProfileUpdateCommand) GetVenafiTypeOk() (*CertificateProfileVenafiType, bool)`

GetVenafiTypeOk returns a tuple with the VenafiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiType

`func (o *CertificateProfileUpdateCommand) SetVenafiType(v CertificateProfileVenafiType)`

SetVenafiType sets VenafiType field to given value.

### HasVenafiType

`func (o *CertificateProfileUpdateCommand) HasVenafiType() bool`

HasVenafiType returns a boolean if a field has been set.

### GetVenafiBaseUrl

`func (o *CertificateProfileUpdateCommand) GetVenafiBaseUrl() string`

GetVenafiBaseUrl returns the VenafiBaseUrl field if non-nil, zero value otherwise.

### GetVenafiBaseUrlOk

`func (o *CertificateProfileUpdateCommand) GetVenafiBaseUrlOk() (*string, bool)`

GetVenafiBaseUrlOk returns a tuple with the VenafiBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiBaseUrl

`func (o *CertificateProfileUpdateCommand) SetVenafiBaseUrl(v string)`

SetVenafiBaseUrl sets VenafiBaseUrl field to given value.

### HasVenafiBaseUrl

`func (o *CertificateProfileUpdateCommand) HasVenafiBaseUrl() bool`

HasVenafiBaseUrl returns a boolean if a field has been set.

### SetVenafiBaseUrlNil

`func (o *CertificateProfileUpdateCommand) SetVenafiBaseUrlNil(b bool)`

 SetVenafiBaseUrlNil sets the value for VenafiBaseUrl to be an explicit nil

### UnsetVenafiBaseUrl
`func (o *CertificateProfileUpdateCommand) UnsetVenafiBaseUrl()`

UnsetVenafiBaseUrl ensures that no value is present for VenafiBaseUrl, not even an explicit nil
### GetVenafiZone

`func (o *CertificateProfileUpdateCommand) GetVenafiZone() string`

GetVenafiZone returns the VenafiZone field if non-nil, zero value otherwise.

### GetVenafiZoneOk

`func (o *CertificateProfileUpdateCommand) GetVenafiZoneOk() (*string, bool)`

GetVenafiZoneOk returns a tuple with the VenafiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZone

`func (o *CertificateProfileUpdateCommand) SetVenafiZone(v string)`

SetVenafiZone sets VenafiZone field to given value.

### HasVenafiZone

`func (o *CertificateProfileUpdateCommand) HasVenafiZone() bool`

HasVenafiZone returns a boolean if a field has been set.

### SetVenafiZoneNil

`func (o *CertificateProfileUpdateCommand) SetVenafiZoneNil(b bool)`

 SetVenafiZoneNil sets the value for VenafiZone to be an explicit nil

### UnsetVenafiZone
`func (o *CertificateProfileUpdateCommand) UnsetVenafiZone()`

UnsetVenafiZone ensures that no value is present for VenafiZone, not even an explicit nil
### GetVenafiAccessToken

`func (o *CertificateProfileUpdateCommand) GetVenafiAccessToken() string`

GetVenafiAccessToken returns the VenafiAccessToken field if non-nil, zero value otherwise.

### GetVenafiAccessTokenOk

`func (o *CertificateProfileUpdateCommand) GetVenafiAccessTokenOk() (*string, bool)`

GetVenafiAccessTokenOk returns a tuple with the VenafiAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAccessToken

`func (o *CertificateProfileUpdateCommand) SetVenafiAccessToken(v string)`

SetVenafiAccessToken sets VenafiAccessToken field to given value.

### HasVenafiAccessToken

`func (o *CertificateProfileUpdateCommand) HasVenafiAccessToken() bool`

HasVenafiAccessToken returns a boolean if a field has been set.

### SetVenafiAccessTokenNil

`func (o *CertificateProfileUpdateCommand) SetVenafiAccessTokenNil(b bool)`

 SetVenafiAccessTokenNil sets the value for VenafiAccessToken to be an explicit nil

### UnsetVenafiAccessToken
`func (o *CertificateProfileUpdateCommand) UnsetVenafiAccessToken()`

UnsetVenafiAccessToken ensures that no value is present for VenafiAccessToken, not even an explicit nil
### GetVenafiApiKey

`func (o *CertificateProfileUpdateCommand) GetVenafiApiKey() string`

GetVenafiApiKey returns the VenafiApiKey field if non-nil, zero value otherwise.

### GetVenafiApiKeyOk

`func (o *CertificateProfileUpdateCommand) GetVenafiApiKeyOk() (*string, bool)`

GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiApiKey

`func (o *CertificateProfileUpdateCommand) SetVenafiApiKey(v string)`

SetVenafiApiKey sets VenafiApiKey field to given value.

### HasVenafiApiKey

`func (o *CertificateProfileUpdateCommand) HasVenafiApiKey() bool`

HasVenafiApiKey returns a boolean if a field has been set.

### SetVenafiApiKeyNil

`func (o *CertificateProfileUpdateCommand) SetVenafiApiKeyNil(b bool)`

 SetVenafiApiKeyNil sets the value for VenafiApiKey to be an explicit nil

### UnsetVenafiApiKey
`func (o *CertificateProfileUpdateCommand) UnsetVenafiApiKey()`

UnsetVenafiApiKey ensures that no value is present for VenafiApiKey, not even an explicit nil
### GetVenafiZoneUuid

`func (o *CertificateProfileUpdateCommand) GetVenafiZoneUuid() string`

GetVenafiZoneUuid returns the VenafiZoneUuid field if non-nil, zero value otherwise.

### GetVenafiZoneUuidOk

`func (o *CertificateProfileUpdateCommand) GetVenafiZoneUuidOk() (*string, bool)`

GetVenafiZoneUuidOk returns a tuple with the VenafiZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZoneUuid

`func (o *CertificateProfileUpdateCommand) SetVenafiZoneUuid(v string)`

SetVenafiZoneUuid sets VenafiZoneUuid field to given value.

### HasVenafiZoneUuid

`func (o *CertificateProfileUpdateCommand) HasVenafiZoneUuid() bool`

HasVenafiZoneUuid returns a boolean if a field has been set.

### SetVenafiZoneUuidNil

`func (o *CertificateProfileUpdateCommand) SetVenafiZoneUuidNil(b bool)`

 SetVenafiZoneUuidNil sets the value for VenafiZoneUuid to be an explicit nil

### UnsetVenafiZoneUuid
`func (o *CertificateProfileUpdateCommand) UnsetVenafiZoneUuid()`

UnsetVenafiZoneUuid ensures that no value is present for VenafiZoneUuid, not even an explicit nil
### GetCustomCaCert

`func (o *CertificateProfileUpdateCommand) GetCustomCaCert() string`

GetCustomCaCert returns the CustomCaCert field if non-nil, zero value otherwise.

### GetCustomCaCertOk

`func (o *CertificateProfileUpdateCommand) GetCustomCaCertOk() (*string, bool)`

GetCustomCaCertOk returns a tuple with the CustomCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaCert

`func (o *CertificateProfileUpdateCommand) SetCustomCaCert(v string)`

SetCustomCaCert sets CustomCaCert field to given value.

### HasCustomCaCert

`func (o *CertificateProfileUpdateCommand) HasCustomCaCert() bool`

HasCustomCaCert returns a boolean if a field has been set.

### SetCustomCaCertNil

`func (o *CertificateProfileUpdateCommand) SetCustomCaCertNil(b bool)`

 SetCustomCaCertNil sets the value for CustomCaCert to be an explicit nil

### UnsetCustomCaCert
`func (o *CertificateProfileUpdateCommand) UnsetCustomCaCert()`

UnsetCustomCaCert ensures that no value is present for CustomCaCert, not even an explicit nil
### GetCustomCaPrivateKey

`func (o *CertificateProfileUpdateCommand) GetCustomCaPrivateKey() string`

GetCustomCaPrivateKey returns the CustomCaPrivateKey field if non-nil, zero value otherwise.

### GetCustomCaPrivateKeyOk

`func (o *CertificateProfileUpdateCommand) GetCustomCaPrivateKeyOk() (*string, bool)`

GetCustomCaPrivateKeyOk returns a tuple with the CustomCaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaPrivateKey

`func (o *CertificateProfileUpdateCommand) SetCustomCaPrivateKey(v string)`

SetCustomCaPrivateKey sets CustomCaPrivateKey field to given value.

### HasCustomCaPrivateKey

`func (o *CertificateProfileUpdateCommand) HasCustomCaPrivateKey() bool`

HasCustomCaPrivateKey returns a boolean if a field has been set.

### SetCustomCaPrivateKeyNil

`func (o *CertificateProfileUpdateCommand) SetCustomCaPrivateKeyNil(b bool)`

 SetCustomCaPrivateKeyNil sets the value for CustomCaPrivateKey to be an explicit nil

### UnsetCustomCaPrivateKey
`func (o *CertificateProfileUpdateCommand) UnsetCustomCaPrivateKey()`

UnsetCustomCaPrivateKey ensures that no value is present for CustomCaPrivateKey, not even an explicit nil
### GetChainPem

`func (o *CertificateProfileUpdateCommand) GetChainPem() string`

GetChainPem returns the ChainPem field if non-nil, zero value otherwise.

### GetChainPemOk

`func (o *CertificateProfileUpdateCommand) GetChainPemOk() (*string, bool)`

GetChainPemOk returns a tuple with the ChainPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainPem

`func (o *CertificateProfileUpdateCommand) SetChainPem(v string)`

SetChainPem sets ChainPem field to given value.

### HasChainPem

`func (o *CertificateProfileUpdateCommand) HasChainPem() bool`

HasChainPem returns a boolean if a field has been set.

### SetChainPemNil

`func (o *CertificateProfileUpdateCommand) SetChainPemNil(b bool)`

 SetChainPemNil sets the value for ChainPem to be an explicit nil

### UnsetChainPem
`func (o *CertificateProfileUpdateCommand) UnsetChainPem()`

UnsetChainPem ensures that no value is present for ChainPem, not even an explicit nil
### GetCommonName

`func (o *CertificateProfileUpdateCommand) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateProfileUpdateCommand) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateProfileUpdateCommand) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateProfileUpdateCommand) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CertificateProfileUpdateCommand) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CertificateProfileUpdateCommand) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetValidityDays

`func (o *CertificateProfileUpdateCommand) GetValidityDays() int32`

GetValidityDays returns the ValidityDays field if non-nil, zero value otherwise.

### GetValidityDaysOk

`func (o *CertificateProfileUpdateCommand) GetValidityDaysOk() (*int32, bool)`

GetValidityDaysOk returns a tuple with the ValidityDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDays

`func (o *CertificateProfileUpdateCommand) SetValidityDays(v int32)`

SetValidityDays sets ValidityDays field to given value.

### HasValidityDays

`func (o *CertificateProfileUpdateCommand) HasValidityDays() bool`

HasValidityDays returns a boolean if a field has been set.

### SetValidityDaysNil

`func (o *CertificateProfileUpdateCommand) SetValidityDaysNil(b bool)`

 SetValidityDaysNil sets the value for ValidityDays to be an explicit nil

### UnsetValidityDays
`func (o *CertificateProfileUpdateCommand) UnsetValidityDays()`

UnsetValidityDays ensures that no value is present for ValidityDays, not even an explicit nil
### GetCertPem

`func (o *CertificateProfileUpdateCommand) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *CertificateProfileUpdateCommand) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *CertificateProfileUpdateCommand) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *CertificateProfileUpdateCommand) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### SetCertPemNil

`func (o *CertificateProfileUpdateCommand) SetCertPemNil(b bool)`

 SetCertPemNil sets the value for CertPem to be an explicit nil

### UnsetCertPem
`func (o *CertificateProfileUpdateCommand) UnsetCertPem()`

UnsetCertPem ensures that no value is present for CertPem, not even an explicit nil
### GetKeyPem

`func (o *CertificateProfileUpdateCommand) GetKeyPem() string`

GetKeyPem returns the KeyPem field if non-nil, zero value otherwise.

### GetKeyPemOk

`func (o *CertificateProfileUpdateCommand) GetKeyPemOk() (*string, bool)`

GetKeyPemOk returns a tuple with the KeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPem

`func (o *CertificateProfileUpdateCommand) SetKeyPem(v string)`

SetKeyPem sets KeyPem field to given value.

### HasKeyPem

`func (o *CertificateProfileUpdateCommand) HasKeyPem() bool`

HasKeyPem returns a boolean if a field has been set.

### SetKeyPemNil

`func (o *CertificateProfileUpdateCommand) SetKeyPemNil(b bool)`

 SetKeyPemNil sets the value for KeyPem to be an explicit nil

### UnsetKeyPem
`func (o *CertificateProfileUpdateCommand) UnsetKeyPem()`

UnsetKeyPem ensures that no value is present for KeyPem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


