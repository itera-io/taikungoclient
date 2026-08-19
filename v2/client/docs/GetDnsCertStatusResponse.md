# GetDnsCertStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertManagerEnabled** | **bool** |  | 
**ExternalDnsEnabled** | **bool** |  | 
**TrustManagerEnabled** | **bool** |  | 
**Enabled** | **bool** |  | 
**DnsProvider** | **NullableString** |  | 
**CertIssuerType** | **NullableString** |  | 
**Domain** | **NullableString** |  | 
**AcmeEmail** | **NullableString** |  | 
**IssuerName** | **NullableString** |  | 
**DesiredConfigVersion** | **int32** |  | 
**AppliedConfigVersion** | **int32** |  | 
**DesiredCertManagerVersion** | **int32** |  | 
**AppliedCertManagerVersion** | **int32** |  | 
**DesiredTrustManagerVersion** | **int32** |  | 
**AppliedTrustManagerVersion** | **int32** |  | 
**DesiredExternalDnsVersion** | **int32** |  | 
**AppliedExternalDnsVersion** | **int32** |  | 
**DesiredIssuerVersion** | **int32** |  | 
**AppliedIssuerVersion** | **int32** |  | 
**Synced** | **bool** |  | 
**DnsCredentialId** | **NullableInt32** |  | 
**DnsCredentialName** | **NullableString** |  | 
**CertificateProfileId** | **NullableInt32** |  | 
**CertificateProfileName** | **NullableString** |  | 
**ProjectStatus** | [**ProjectStatus**](ProjectStatus.md) |  | 
**FailureReason** | **NullableString** |  | 

## Methods

### NewGetDnsCertStatusResponse

`func NewGetDnsCertStatusResponse(certManagerEnabled bool, externalDnsEnabled bool, trustManagerEnabled bool, enabled bool, dnsProvider NullableString, certIssuerType NullableString, domain NullableString, acmeEmail NullableString, issuerName NullableString, desiredConfigVersion int32, appliedConfigVersion int32, desiredCertManagerVersion int32, appliedCertManagerVersion int32, desiredTrustManagerVersion int32, appliedTrustManagerVersion int32, desiredExternalDnsVersion int32, appliedExternalDnsVersion int32, desiredIssuerVersion int32, appliedIssuerVersion int32, synced bool, dnsCredentialId NullableInt32, dnsCredentialName NullableString, certificateProfileId NullableInt32, certificateProfileName NullableString, projectStatus ProjectStatus, failureReason NullableString, ) *GetDnsCertStatusResponse`

NewGetDnsCertStatusResponse instantiates a new GetDnsCertStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDnsCertStatusResponseWithDefaults

`func NewGetDnsCertStatusResponseWithDefaults() *GetDnsCertStatusResponse`

NewGetDnsCertStatusResponseWithDefaults instantiates a new GetDnsCertStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertManagerEnabled

`func (o *GetDnsCertStatusResponse) GetCertManagerEnabled() bool`

GetCertManagerEnabled returns the CertManagerEnabled field if non-nil, zero value otherwise.

### GetCertManagerEnabledOk

`func (o *GetDnsCertStatusResponse) GetCertManagerEnabledOk() (*bool, bool)`

GetCertManagerEnabledOk returns a tuple with the CertManagerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerEnabled

`func (o *GetDnsCertStatusResponse) SetCertManagerEnabled(v bool)`

SetCertManagerEnabled sets CertManagerEnabled field to given value.


### GetExternalDnsEnabled

`func (o *GetDnsCertStatusResponse) GetExternalDnsEnabled() bool`

GetExternalDnsEnabled returns the ExternalDnsEnabled field if non-nil, zero value otherwise.

### GetExternalDnsEnabledOk

`func (o *GetDnsCertStatusResponse) GetExternalDnsEnabledOk() (*bool, bool)`

GetExternalDnsEnabledOk returns a tuple with the ExternalDnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDnsEnabled

`func (o *GetDnsCertStatusResponse) SetExternalDnsEnabled(v bool)`

SetExternalDnsEnabled sets ExternalDnsEnabled field to given value.


### GetTrustManagerEnabled

`func (o *GetDnsCertStatusResponse) GetTrustManagerEnabled() bool`

GetTrustManagerEnabled returns the TrustManagerEnabled field if non-nil, zero value otherwise.

### GetTrustManagerEnabledOk

`func (o *GetDnsCertStatusResponse) GetTrustManagerEnabledOk() (*bool, bool)`

GetTrustManagerEnabledOk returns a tuple with the TrustManagerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerEnabled

`func (o *GetDnsCertStatusResponse) SetTrustManagerEnabled(v bool)`

SetTrustManagerEnabled sets TrustManagerEnabled field to given value.


### GetEnabled

`func (o *GetDnsCertStatusResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDnsCertStatusResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDnsCertStatusResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDnsProvider

`func (o *GetDnsCertStatusResponse) GetDnsProvider() string`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *GetDnsCertStatusResponse) GetDnsProviderOk() (*string, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *GetDnsCertStatusResponse) SetDnsProvider(v string)`

SetDnsProvider sets DnsProvider field to given value.


### SetDnsProviderNil

`func (o *GetDnsCertStatusResponse) SetDnsProviderNil(b bool)`

 SetDnsProviderNil sets the value for DnsProvider to be an explicit nil

### UnsetDnsProvider
`func (o *GetDnsCertStatusResponse) UnsetDnsProvider()`

UnsetDnsProvider ensures that no value is present for DnsProvider, not even an explicit nil
### GetCertIssuerType

`func (o *GetDnsCertStatusResponse) GetCertIssuerType() string`

GetCertIssuerType returns the CertIssuerType field if non-nil, zero value otherwise.

### GetCertIssuerTypeOk

`func (o *GetDnsCertStatusResponse) GetCertIssuerTypeOk() (*string, bool)`

GetCertIssuerTypeOk returns a tuple with the CertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerType

`func (o *GetDnsCertStatusResponse) SetCertIssuerType(v string)`

SetCertIssuerType sets CertIssuerType field to given value.


### SetCertIssuerTypeNil

`func (o *GetDnsCertStatusResponse) SetCertIssuerTypeNil(b bool)`

 SetCertIssuerTypeNil sets the value for CertIssuerType to be an explicit nil

### UnsetCertIssuerType
`func (o *GetDnsCertStatusResponse) UnsetCertIssuerType()`

UnsetCertIssuerType ensures that no value is present for CertIssuerType, not even an explicit nil
### GetDomain

`func (o *GetDnsCertStatusResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDnsCertStatusResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDnsCertStatusResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *GetDnsCertStatusResponse) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GetDnsCertStatusResponse) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetAcmeEmail

`func (o *GetDnsCertStatusResponse) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *GetDnsCertStatusResponse) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *GetDnsCertStatusResponse) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.


### SetAcmeEmailNil

`func (o *GetDnsCertStatusResponse) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *GetDnsCertStatusResponse) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetIssuerName

`func (o *GetDnsCertStatusResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *GetDnsCertStatusResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *GetDnsCertStatusResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### SetIssuerNameNil

`func (o *GetDnsCertStatusResponse) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *GetDnsCertStatusResponse) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetDesiredConfigVersion

`func (o *GetDnsCertStatusResponse) GetDesiredConfigVersion() int32`

GetDesiredConfigVersion returns the DesiredConfigVersion field if non-nil, zero value otherwise.

### GetDesiredConfigVersionOk

`func (o *GetDnsCertStatusResponse) GetDesiredConfigVersionOk() (*int32, bool)`

GetDesiredConfigVersionOk returns a tuple with the DesiredConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredConfigVersion

`func (o *GetDnsCertStatusResponse) SetDesiredConfigVersion(v int32)`

SetDesiredConfigVersion sets DesiredConfigVersion field to given value.


### GetAppliedConfigVersion

`func (o *GetDnsCertStatusResponse) GetAppliedConfigVersion() int32`

GetAppliedConfigVersion returns the AppliedConfigVersion field if non-nil, zero value otherwise.

### GetAppliedConfigVersionOk

`func (o *GetDnsCertStatusResponse) GetAppliedConfigVersionOk() (*int32, bool)`

GetAppliedConfigVersionOk returns a tuple with the AppliedConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedConfigVersion

`func (o *GetDnsCertStatusResponse) SetAppliedConfigVersion(v int32)`

SetAppliedConfigVersion sets AppliedConfigVersion field to given value.


### GetDesiredCertManagerVersion

`func (o *GetDnsCertStatusResponse) GetDesiredCertManagerVersion() int32`

GetDesiredCertManagerVersion returns the DesiredCertManagerVersion field if non-nil, zero value otherwise.

### GetDesiredCertManagerVersionOk

`func (o *GetDnsCertStatusResponse) GetDesiredCertManagerVersionOk() (*int32, bool)`

GetDesiredCertManagerVersionOk returns a tuple with the DesiredCertManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCertManagerVersion

`func (o *GetDnsCertStatusResponse) SetDesiredCertManagerVersion(v int32)`

SetDesiredCertManagerVersion sets DesiredCertManagerVersion field to given value.


### GetAppliedCertManagerVersion

`func (o *GetDnsCertStatusResponse) GetAppliedCertManagerVersion() int32`

GetAppliedCertManagerVersion returns the AppliedCertManagerVersion field if non-nil, zero value otherwise.

### GetAppliedCertManagerVersionOk

`func (o *GetDnsCertStatusResponse) GetAppliedCertManagerVersionOk() (*int32, bool)`

GetAppliedCertManagerVersionOk returns a tuple with the AppliedCertManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCertManagerVersion

`func (o *GetDnsCertStatusResponse) SetAppliedCertManagerVersion(v int32)`

SetAppliedCertManagerVersion sets AppliedCertManagerVersion field to given value.


### GetDesiredTrustManagerVersion

`func (o *GetDnsCertStatusResponse) GetDesiredTrustManagerVersion() int32`

GetDesiredTrustManagerVersion returns the DesiredTrustManagerVersion field if non-nil, zero value otherwise.

### GetDesiredTrustManagerVersionOk

`func (o *GetDnsCertStatusResponse) GetDesiredTrustManagerVersionOk() (*int32, bool)`

GetDesiredTrustManagerVersionOk returns a tuple with the DesiredTrustManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTrustManagerVersion

`func (o *GetDnsCertStatusResponse) SetDesiredTrustManagerVersion(v int32)`

SetDesiredTrustManagerVersion sets DesiredTrustManagerVersion field to given value.


### GetAppliedTrustManagerVersion

`func (o *GetDnsCertStatusResponse) GetAppliedTrustManagerVersion() int32`

GetAppliedTrustManagerVersion returns the AppliedTrustManagerVersion field if non-nil, zero value otherwise.

### GetAppliedTrustManagerVersionOk

`func (o *GetDnsCertStatusResponse) GetAppliedTrustManagerVersionOk() (*int32, bool)`

GetAppliedTrustManagerVersionOk returns a tuple with the AppliedTrustManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTrustManagerVersion

`func (o *GetDnsCertStatusResponse) SetAppliedTrustManagerVersion(v int32)`

SetAppliedTrustManagerVersion sets AppliedTrustManagerVersion field to given value.


### GetDesiredExternalDnsVersion

`func (o *GetDnsCertStatusResponse) GetDesiredExternalDnsVersion() int32`

GetDesiredExternalDnsVersion returns the DesiredExternalDnsVersion field if non-nil, zero value otherwise.

### GetDesiredExternalDnsVersionOk

`func (o *GetDnsCertStatusResponse) GetDesiredExternalDnsVersionOk() (*int32, bool)`

GetDesiredExternalDnsVersionOk returns a tuple with the DesiredExternalDnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredExternalDnsVersion

`func (o *GetDnsCertStatusResponse) SetDesiredExternalDnsVersion(v int32)`

SetDesiredExternalDnsVersion sets DesiredExternalDnsVersion field to given value.


### GetAppliedExternalDnsVersion

`func (o *GetDnsCertStatusResponse) GetAppliedExternalDnsVersion() int32`

GetAppliedExternalDnsVersion returns the AppliedExternalDnsVersion field if non-nil, zero value otherwise.

### GetAppliedExternalDnsVersionOk

`func (o *GetDnsCertStatusResponse) GetAppliedExternalDnsVersionOk() (*int32, bool)`

GetAppliedExternalDnsVersionOk returns a tuple with the AppliedExternalDnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedExternalDnsVersion

`func (o *GetDnsCertStatusResponse) SetAppliedExternalDnsVersion(v int32)`

SetAppliedExternalDnsVersion sets AppliedExternalDnsVersion field to given value.


### GetDesiredIssuerVersion

`func (o *GetDnsCertStatusResponse) GetDesiredIssuerVersion() int32`

GetDesiredIssuerVersion returns the DesiredIssuerVersion field if non-nil, zero value otherwise.

### GetDesiredIssuerVersionOk

`func (o *GetDnsCertStatusResponse) GetDesiredIssuerVersionOk() (*int32, bool)`

GetDesiredIssuerVersionOk returns a tuple with the DesiredIssuerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredIssuerVersion

`func (o *GetDnsCertStatusResponse) SetDesiredIssuerVersion(v int32)`

SetDesiredIssuerVersion sets DesiredIssuerVersion field to given value.


### GetAppliedIssuerVersion

`func (o *GetDnsCertStatusResponse) GetAppliedIssuerVersion() int32`

GetAppliedIssuerVersion returns the AppliedIssuerVersion field if non-nil, zero value otherwise.

### GetAppliedIssuerVersionOk

`func (o *GetDnsCertStatusResponse) GetAppliedIssuerVersionOk() (*int32, bool)`

GetAppliedIssuerVersionOk returns a tuple with the AppliedIssuerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedIssuerVersion

`func (o *GetDnsCertStatusResponse) SetAppliedIssuerVersion(v int32)`

SetAppliedIssuerVersion sets AppliedIssuerVersion field to given value.


### GetSynced

`func (o *GetDnsCertStatusResponse) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *GetDnsCertStatusResponse) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *GetDnsCertStatusResponse) SetSynced(v bool)`

SetSynced sets Synced field to given value.


### GetDnsCredentialId

`func (o *GetDnsCertStatusResponse) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *GetDnsCertStatusResponse) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *GetDnsCertStatusResponse) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.


### SetDnsCredentialIdNil

`func (o *GetDnsCertStatusResponse) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *GetDnsCertStatusResponse) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetDnsCredentialName

`func (o *GetDnsCertStatusResponse) GetDnsCredentialName() string`

GetDnsCredentialName returns the DnsCredentialName field if non-nil, zero value otherwise.

### GetDnsCredentialNameOk

`func (o *GetDnsCertStatusResponse) GetDnsCredentialNameOk() (*string, bool)`

GetDnsCredentialNameOk returns a tuple with the DnsCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialName

`func (o *GetDnsCertStatusResponse) SetDnsCredentialName(v string)`

SetDnsCredentialName sets DnsCredentialName field to given value.


### SetDnsCredentialNameNil

`func (o *GetDnsCertStatusResponse) SetDnsCredentialNameNil(b bool)`

 SetDnsCredentialNameNil sets the value for DnsCredentialName to be an explicit nil

### UnsetDnsCredentialName
`func (o *GetDnsCertStatusResponse) UnsetDnsCredentialName()`

UnsetDnsCredentialName ensures that no value is present for DnsCredentialName, not even an explicit nil
### GetCertificateProfileId

`func (o *GetDnsCertStatusResponse) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *GetDnsCertStatusResponse) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *GetDnsCertStatusResponse) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.


### SetCertificateProfileIdNil

`func (o *GetDnsCertStatusResponse) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *GetDnsCertStatusResponse) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil
### GetCertificateProfileName

`func (o *GetDnsCertStatusResponse) GetCertificateProfileName() string`

GetCertificateProfileName returns the CertificateProfileName field if non-nil, zero value otherwise.

### GetCertificateProfileNameOk

`func (o *GetDnsCertStatusResponse) GetCertificateProfileNameOk() (*string, bool)`

GetCertificateProfileNameOk returns a tuple with the CertificateProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileName

`func (o *GetDnsCertStatusResponse) SetCertificateProfileName(v string)`

SetCertificateProfileName sets CertificateProfileName field to given value.


### SetCertificateProfileNameNil

`func (o *GetDnsCertStatusResponse) SetCertificateProfileNameNil(b bool)`

 SetCertificateProfileNameNil sets the value for CertificateProfileName to be an explicit nil

### UnsetCertificateProfileName
`func (o *GetDnsCertStatusResponse) UnsetCertificateProfileName()`

UnsetCertificateProfileName ensures that no value is present for CertificateProfileName, not even an explicit nil
### GetProjectStatus

`func (o *GetDnsCertStatusResponse) GetProjectStatus() ProjectStatus`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *GetDnsCertStatusResponse) GetProjectStatusOk() (*ProjectStatus, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *GetDnsCertStatusResponse) SetProjectStatus(v ProjectStatus)`

SetProjectStatus sets ProjectStatus field to given value.


### GetFailureReason

`func (o *GetDnsCertStatusResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetDnsCertStatusResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetDnsCertStatusResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *GetDnsCertStatusResponse) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *GetDnsCertStatusResponse) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


