# CertificateProfileCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IssuerMode** | Pointer to [**CertificateProfileIssuerMode**](CertificateProfileIssuerMode.md) |  | [optional] 
**TrustBundle** | Pointer to [**[]TrustBundleItem**](TrustBundleItem.md) |  | [optional] 
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
**Kind** | Pointer to **NullableString** |  | [optional] 
**CertPem** | Pointer to **NullableString** |  | [optional] 
**KeyPem** | Pointer to **NullableString** |  | [optional] 
**ChainPem** | Pointer to **NullableString** |  | [optional] 
**GenerateSelfSigned** | Pointer to **bool** |  | [optional] [default to false]
**CommonName** | Pointer to **NullableString** |  | [optional] 
**ValidityDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCertificateProfileCreateCommand

`func NewCertificateProfileCreateCommand(name string, organizationId NullableInt32, ) *CertificateProfileCreateCommand`

NewCertificateProfileCreateCommand instantiates a new CertificateProfileCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileCreateCommandWithDefaults

`func NewCertificateProfileCreateCommandWithDefaults() *CertificateProfileCreateCommand`

NewCertificateProfileCreateCommandWithDefaults instantiates a new CertificateProfileCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateProfileCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CertificateProfileCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CertificateProfileCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CertificateProfileCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CertificateProfileCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CertificateProfileCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetDescription

`func (o *CertificateProfileCreateCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileCreateCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileCreateCommand) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateProfileCreateCommand) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateProfileCreateCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileCreateCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIssuerMode

`func (o *CertificateProfileCreateCommand) GetIssuerMode() CertificateProfileIssuerMode`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *CertificateProfileCreateCommand) GetIssuerModeOk() (*CertificateProfileIssuerMode, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *CertificateProfileCreateCommand) SetIssuerMode(v CertificateProfileIssuerMode)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *CertificateProfileCreateCommand) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetTrustBundle

`func (o *CertificateProfileCreateCommand) GetTrustBundle() []TrustBundleItem`

GetTrustBundle returns the TrustBundle field if non-nil, zero value otherwise.

### GetTrustBundleOk

`func (o *CertificateProfileCreateCommand) GetTrustBundleOk() (*[]TrustBundleItem, bool)`

GetTrustBundleOk returns a tuple with the TrustBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustBundle

`func (o *CertificateProfileCreateCommand) SetTrustBundle(v []TrustBundleItem)`

SetTrustBundle sets TrustBundle field to given value.

### HasTrustBundle

`func (o *CertificateProfileCreateCommand) HasTrustBundle() bool`

HasTrustBundle returns a boolean if a field has been set.

### SetTrustBundleNil

`func (o *CertificateProfileCreateCommand) SetTrustBundleNil(b bool)`

 SetTrustBundleNil sets the value for TrustBundle to be an explicit nil

### UnsetTrustBundle
`func (o *CertificateProfileCreateCommand) UnsetTrustBundle()`

UnsetTrustBundle ensures that no value is present for TrustBundle, not even an explicit nil
### GetAcmeEmail

`func (o *CertificateProfileCreateCommand) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *CertificateProfileCreateCommand) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *CertificateProfileCreateCommand) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *CertificateProfileCreateCommand) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### SetAcmeEmailNil

`func (o *CertificateProfileCreateCommand) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *CertificateProfileCreateCommand) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetAcmeServer

`func (o *CertificateProfileCreateCommand) GetAcmeServer() string`

GetAcmeServer returns the AcmeServer field if non-nil, zero value otherwise.

### GetAcmeServerOk

`func (o *CertificateProfileCreateCommand) GetAcmeServerOk() (*string, bool)`

GetAcmeServerOk returns a tuple with the AcmeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeServer

`func (o *CertificateProfileCreateCommand) SetAcmeServer(v string)`

SetAcmeServer sets AcmeServer field to given value.

### HasAcmeServer

`func (o *CertificateProfileCreateCommand) HasAcmeServer() bool`

HasAcmeServer returns a boolean if a field has been set.

### SetAcmeServerNil

`func (o *CertificateProfileCreateCommand) SetAcmeServerNil(b bool)`

 SetAcmeServerNil sets the value for AcmeServer to be an explicit nil

### UnsetAcmeServer
`func (o *CertificateProfileCreateCommand) UnsetAcmeServer()`

UnsetAcmeServer ensures that no value is present for AcmeServer, not even an explicit nil
### GetVenafiType

`func (o *CertificateProfileCreateCommand) GetVenafiType() CertificateProfileVenafiType`

GetVenafiType returns the VenafiType field if non-nil, zero value otherwise.

### GetVenafiTypeOk

`func (o *CertificateProfileCreateCommand) GetVenafiTypeOk() (*CertificateProfileVenafiType, bool)`

GetVenafiTypeOk returns a tuple with the VenafiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiType

`func (o *CertificateProfileCreateCommand) SetVenafiType(v CertificateProfileVenafiType)`

SetVenafiType sets VenafiType field to given value.

### HasVenafiType

`func (o *CertificateProfileCreateCommand) HasVenafiType() bool`

HasVenafiType returns a boolean if a field has been set.

### GetVenafiBaseUrl

`func (o *CertificateProfileCreateCommand) GetVenafiBaseUrl() string`

GetVenafiBaseUrl returns the VenafiBaseUrl field if non-nil, zero value otherwise.

### GetVenafiBaseUrlOk

`func (o *CertificateProfileCreateCommand) GetVenafiBaseUrlOk() (*string, bool)`

GetVenafiBaseUrlOk returns a tuple with the VenafiBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiBaseUrl

`func (o *CertificateProfileCreateCommand) SetVenafiBaseUrl(v string)`

SetVenafiBaseUrl sets VenafiBaseUrl field to given value.

### HasVenafiBaseUrl

`func (o *CertificateProfileCreateCommand) HasVenafiBaseUrl() bool`

HasVenafiBaseUrl returns a boolean if a field has been set.

### SetVenafiBaseUrlNil

`func (o *CertificateProfileCreateCommand) SetVenafiBaseUrlNil(b bool)`

 SetVenafiBaseUrlNil sets the value for VenafiBaseUrl to be an explicit nil

### UnsetVenafiBaseUrl
`func (o *CertificateProfileCreateCommand) UnsetVenafiBaseUrl()`

UnsetVenafiBaseUrl ensures that no value is present for VenafiBaseUrl, not even an explicit nil
### GetVenafiZone

`func (o *CertificateProfileCreateCommand) GetVenafiZone() string`

GetVenafiZone returns the VenafiZone field if non-nil, zero value otherwise.

### GetVenafiZoneOk

`func (o *CertificateProfileCreateCommand) GetVenafiZoneOk() (*string, bool)`

GetVenafiZoneOk returns a tuple with the VenafiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZone

`func (o *CertificateProfileCreateCommand) SetVenafiZone(v string)`

SetVenafiZone sets VenafiZone field to given value.

### HasVenafiZone

`func (o *CertificateProfileCreateCommand) HasVenafiZone() bool`

HasVenafiZone returns a boolean if a field has been set.

### SetVenafiZoneNil

`func (o *CertificateProfileCreateCommand) SetVenafiZoneNil(b bool)`

 SetVenafiZoneNil sets the value for VenafiZone to be an explicit nil

### UnsetVenafiZone
`func (o *CertificateProfileCreateCommand) UnsetVenafiZone()`

UnsetVenafiZone ensures that no value is present for VenafiZone, not even an explicit nil
### GetVenafiAccessToken

`func (o *CertificateProfileCreateCommand) GetVenafiAccessToken() string`

GetVenafiAccessToken returns the VenafiAccessToken field if non-nil, zero value otherwise.

### GetVenafiAccessTokenOk

`func (o *CertificateProfileCreateCommand) GetVenafiAccessTokenOk() (*string, bool)`

GetVenafiAccessTokenOk returns a tuple with the VenafiAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAccessToken

`func (o *CertificateProfileCreateCommand) SetVenafiAccessToken(v string)`

SetVenafiAccessToken sets VenafiAccessToken field to given value.

### HasVenafiAccessToken

`func (o *CertificateProfileCreateCommand) HasVenafiAccessToken() bool`

HasVenafiAccessToken returns a boolean if a field has been set.

### SetVenafiAccessTokenNil

`func (o *CertificateProfileCreateCommand) SetVenafiAccessTokenNil(b bool)`

 SetVenafiAccessTokenNil sets the value for VenafiAccessToken to be an explicit nil

### UnsetVenafiAccessToken
`func (o *CertificateProfileCreateCommand) UnsetVenafiAccessToken()`

UnsetVenafiAccessToken ensures that no value is present for VenafiAccessToken, not even an explicit nil
### GetVenafiApiKey

`func (o *CertificateProfileCreateCommand) GetVenafiApiKey() string`

GetVenafiApiKey returns the VenafiApiKey field if non-nil, zero value otherwise.

### GetVenafiApiKeyOk

`func (o *CertificateProfileCreateCommand) GetVenafiApiKeyOk() (*string, bool)`

GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiApiKey

`func (o *CertificateProfileCreateCommand) SetVenafiApiKey(v string)`

SetVenafiApiKey sets VenafiApiKey field to given value.

### HasVenafiApiKey

`func (o *CertificateProfileCreateCommand) HasVenafiApiKey() bool`

HasVenafiApiKey returns a boolean if a field has been set.

### SetVenafiApiKeyNil

`func (o *CertificateProfileCreateCommand) SetVenafiApiKeyNil(b bool)`

 SetVenafiApiKeyNil sets the value for VenafiApiKey to be an explicit nil

### UnsetVenafiApiKey
`func (o *CertificateProfileCreateCommand) UnsetVenafiApiKey()`

UnsetVenafiApiKey ensures that no value is present for VenafiApiKey, not even an explicit nil
### GetVenafiZoneUuid

`func (o *CertificateProfileCreateCommand) GetVenafiZoneUuid() string`

GetVenafiZoneUuid returns the VenafiZoneUuid field if non-nil, zero value otherwise.

### GetVenafiZoneUuidOk

`func (o *CertificateProfileCreateCommand) GetVenafiZoneUuidOk() (*string, bool)`

GetVenafiZoneUuidOk returns a tuple with the VenafiZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZoneUuid

`func (o *CertificateProfileCreateCommand) SetVenafiZoneUuid(v string)`

SetVenafiZoneUuid sets VenafiZoneUuid field to given value.

### HasVenafiZoneUuid

`func (o *CertificateProfileCreateCommand) HasVenafiZoneUuid() bool`

HasVenafiZoneUuid returns a boolean if a field has been set.

### SetVenafiZoneUuidNil

`func (o *CertificateProfileCreateCommand) SetVenafiZoneUuidNil(b bool)`

 SetVenafiZoneUuidNil sets the value for VenafiZoneUuid to be an explicit nil

### UnsetVenafiZoneUuid
`func (o *CertificateProfileCreateCommand) UnsetVenafiZoneUuid()`

UnsetVenafiZoneUuid ensures that no value is present for VenafiZoneUuid, not even an explicit nil
### GetCustomCaCert

`func (o *CertificateProfileCreateCommand) GetCustomCaCert() string`

GetCustomCaCert returns the CustomCaCert field if non-nil, zero value otherwise.

### GetCustomCaCertOk

`func (o *CertificateProfileCreateCommand) GetCustomCaCertOk() (*string, bool)`

GetCustomCaCertOk returns a tuple with the CustomCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaCert

`func (o *CertificateProfileCreateCommand) SetCustomCaCert(v string)`

SetCustomCaCert sets CustomCaCert field to given value.

### HasCustomCaCert

`func (o *CertificateProfileCreateCommand) HasCustomCaCert() bool`

HasCustomCaCert returns a boolean if a field has been set.

### SetCustomCaCertNil

`func (o *CertificateProfileCreateCommand) SetCustomCaCertNil(b bool)`

 SetCustomCaCertNil sets the value for CustomCaCert to be an explicit nil

### UnsetCustomCaCert
`func (o *CertificateProfileCreateCommand) UnsetCustomCaCert()`

UnsetCustomCaCert ensures that no value is present for CustomCaCert, not even an explicit nil
### GetCustomCaPrivateKey

`func (o *CertificateProfileCreateCommand) GetCustomCaPrivateKey() string`

GetCustomCaPrivateKey returns the CustomCaPrivateKey field if non-nil, zero value otherwise.

### GetCustomCaPrivateKeyOk

`func (o *CertificateProfileCreateCommand) GetCustomCaPrivateKeyOk() (*string, bool)`

GetCustomCaPrivateKeyOk returns a tuple with the CustomCaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaPrivateKey

`func (o *CertificateProfileCreateCommand) SetCustomCaPrivateKey(v string)`

SetCustomCaPrivateKey sets CustomCaPrivateKey field to given value.

### HasCustomCaPrivateKey

`func (o *CertificateProfileCreateCommand) HasCustomCaPrivateKey() bool`

HasCustomCaPrivateKey returns a boolean if a field has been set.

### SetCustomCaPrivateKeyNil

`func (o *CertificateProfileCreateCommand) SetCustomCaPrivateKeyNil(b bool)`

 SetCustomCaPrivateKeyNil sets the value for CustomCaPrivateKey to be an explicit nil

### UnsetCustomCaPrivateKey
`func (o *CertificateProfileCreateCommand) UnsetCustomCaPrivateKey()`

UnsetCustomCaPrivateKey ensures that no value is present for CustomCaPrivateKey, not even an explicit nil
### GetKind

`func (o *CertificateProfileCreateCommand) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CertificateProfileCreateCommand) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CertificateProfileCreateCommand) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CertificateProfileCreateCommand) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *CertificateProfileCreateCommand) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *CertificateProfileCreateCommand) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetCertPem

`func (o *CertificateProfileCreateCommand) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *CertificateProfileCreateCommand) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *CertificateProfileCreateCommand) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *CertificateProfileCreateCommand) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### SetCertPemNil

`func (o *CertificateProfileCreateCommand) SetCertPemNil(b bool)`

 SetCertPemNil sets the value for CertPem to be an explicit nil

### UnsetCertPem
`func (o *CertificateProfileCreateCommand) UnsetCertPem()`

UnsetCertPem ensures that no value is present for CertPem, not even an explicit nil
### GetKeyPem

`func (o *CertificateProfileCreateCommand) GetKeyPem() string`

GetKeyPem returns the KeyPem field if non-nil, zero value otherwise.

### GetKeyPemOk

`func (o *CertificateProfileCreateCommand) GetKeyPemOk() (*string, bool)`

GetKeyPemOk returns a tuple with the KeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPem

`func (o *CertificateProfileCreateCommand) SetKeyPem(v string)`

SetKeyPem sets KeyPem field to given value.

### HasKeyPem

`func (o *CertificateProfileCreateCommand) HasKeyPem() bool`

HasKeyPem returns a boolean if a field has been set.

### SetKeyPemNil

`func (o *CertificateProfileCreateCommand) SetKeyPemNil(b bool)`

 SetKeyPemNil sets the value for KeyPem to be an explicit nil

### UnsetKeyPem
`func (o *CertificateProfileCreateCommand) UnsetKeyPem()`

UnsetKeyPem ensures that no value is present for KeyPem, not even an explicit nil
### GetChainPem

`func (o *CertificateProfileCreateCommand) GetChainPem() string`

GetChainPem returns the ChainPem field if non-nil, zero value otherwise.

### GetChainPemOk

`func (o *CertificateProfileCreateCommand) GetChainPemOk() (*string, bool)`

GetChainPemOk returns a tuple with the ChainPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainPem

`func (o *CertificateProfileCreateCommand) SetChainPem(v string)`

SetChainPem sets ChainPem field to given value.

### HasChainPem

`func (o *CertificateProfileCreateCommand) HasChainPem() bool`

HasChainPem returns a boolean if a field has been set.

### SetChainPemNil

`func (o *CertificateProfileCreateCommand) SetChainPemNil(b bool)`

 SetChainPemNil sets the value for ChainPem to be an explicit nil

### UnsetChainPem
`func (o *CertificateProfileCreateCommand) UnsetChainPem()`

UnsetChainPem ensures that no value is present for ChainPem, not even an explicit nil
### GetGenerateSelfSigned

`func (o *CertificateProfileCreateCommand) GetGenerateSelfSigned() bool`

GetGenerateSelfSigned returns the GenerateSelfSigned field if non-nil, zero value otherwise.

### GetGenerateSelfSignedOk

`func (o *CertificateProfileCreateCommand) GetGenerateSelfSignedOk() (*bool, bool)`

GetGenerateSelfSignedOk returns a tuple with the GenerateSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSelfSigned

`func (o *CertificateProfileCreateCommand) SetGenerateSelfSigned(v bool)`

SetGenerateSelfSigned sets GenerateSelfSigned field to given value.

### HasGenerateSelfSigned

`func (o *CertificateProfileCreateCommand) HasGenerateSelfSigned() bool`

HasGenerateSelfSigned returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificateProfileCreateCommand) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateProfileCreateCommand) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateProfileCreateCommand) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateProfileCreateCommand) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CertificateProfileCreateCommand) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CertificateProfileCreateCommand) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetValidityDays

`func (o *CertificateProfileCreateCommand) GetValidityDays() int32`

GetValidityDays returns the ValidityDays field if non-nil, zero value otherwise.

### GetValidityDaysOk

`func (o *CertificateProfileCreateCommand) GetValidityDaysOk() (*int32, bool)`

GetValidityDaysOk returns a tuple with the ValidityDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDays

`func (o *CertificateProfileCreateCommand) SetValidityDays(v int32)`

SetValidityDays sets ValidityDays field to given value.

### HasValidityDays

`func (o *CertificateProfileCreateCommand) HasValidityDays() bool`

HasValidityDays returns a boolean if a field has been set.

### SetValidityDaysNil

`func (o *CertificateProfileCreateCommand) SetValidityDaysNil(b bool)`

 SetValidityDaysNil sets the value for ValidityDays to be an explicit nil

### UnsetValidityDays
`func (o *CertificateProfileCreateCommand) UnsetValidityDays()`

UnsetValidityDays ensures that no value is present for ValidityDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


