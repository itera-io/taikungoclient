# EnableDnsCertCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**DnsCredentialId** | **NullableInt32** |  | 
**CertIssuerType** | **NullableString** |  | 
**AcmeEmail** | **NullableString** |  | 
**Domain** | **NullableString** |  | 
**CertificateProfileId** | Pointer to **NullableInt32** |  | [optional] 
**EnableTrustManager** | Pointer to **bool** |  | [optional] [default to false]
**StageIntentOnly** | Pointer to **bool** |  | [optional] [default to false]
**InternalCaNoDnsStage** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnableDnsCertCommand

`func NewEnableDnsCertCommand(projectId int32, dnsCredentialId NullableInt32, certIssuerType NullableString, acmeEmail NullableString, domain NullableString, ) *EnableDnsCertCommand`

NewEnableDnsCertCommand instantiates a new EnableDnsCertCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableDnsCertCommandWithDefaults

`func NewEnableDnsCertCommandWithDefaults() *EnableDnsCertCommand`

NewEnableDnsCertCommandWithDefaults instantiates a new EnableDnsCertCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *EnableDnsCertCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnableDnsCertCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnableDnsCertCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetDnsCredentialId

`func (o *EnableDnsCertCommand) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *EnableDnsCertCommand) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *EnableDnsCertCommand) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.


### SetDnsCredentialIdNil

`func (o *EnableDnsCertCommand) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *EnableDnsCertCommand) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetCertIssuerType

`func (o *EnableDnsCertCommand) GetCertIssuerType() string`

GetCertIssuerType returns the CertIssuerType field if non-nil, zero value otherwise.

### GetCertIssuerTypeOk

`func (o *EnableDnsCertCommand) GetCertIssuerTypeOk() (*string, bool)`

GetCertIssuerTypeOk returns a tuple with the CertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerType

`func (o *EnableDnsCertCommand) SetCertIssuerType(v string)`

SetCertIssuerType sets CertIssuerType field to given value.


### SetCertIssuerTypeNil

`func (o *EnableDnsCertCommand) SetCertIssuerTypeNil(b bool)`

 SetCertIssuerTypeNil sets the value for CertIssuerType to be an explicit nil

### UnsetCertIssuerType
`func (o *EnableDnsCertCommand) UnsetCertIssuerType()`

UnsetCertIssuerType ensures that no value is present for CertIssuerType, not even an explicit nil
### GetAcmeEmail

`func (o *EnableDnsCertCommand) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *EnableDnsCertCommand) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *EnableDnsCertCommand) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.


### SetAcmeEmailNil

`func (o *EnableDnsCertCommand) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *EnableDnsCertCommand) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetDomain

`func (o *EnableDnsCertCommand) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnableDnsCertCommand) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnableDnsCertCommand) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *EnableDnsCertCommand) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *EnableDnsCertCommand) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCertificateProfileId

`func (o *EnableDnsCertCommand) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *EnableDnsCertCommand) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *EnableDnsCertCommand) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.

### HasCertificateProfileId

`func (o *EnableDnsCertCommand) HasCertificateProfileId() bool`

HasCertificateProfileId returns a boolean if a field has been set.

### SetCertificateProfileIdNil

`func (o *EnableDnsCertCommand) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *EnableDnsCertCommand) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil
### GetEnableTrustManager

`func (o *EnableDnsCertCommand) GetEnableTrustManager() bool`

GetEnableTrustManager returns the EnableTrustManager field if non-nil, zero value otherwise.

### GetEnableTrustManagerOk

`func (o *EnableDnsCertCommand) GetEnableTrustManagerOk() (*bool, bool)`

GetEnableTrustManagerOk returns a tuple with the EnableTrustManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrustManager

`func (o *EnableDnsCertCommand) SetEnableTrustManager(v bool)`

SetEnableTrustManager sets EnableTrustManager field to given value.

### HasEnableTrustManager

`func (o *EnableDnsCertCommand) HasEnableTrustManager() bool`

HasEnableTrustManager returns a boolean if a field has been set.

### GetStageIntentOnly

`func (o *EnableDnsCertCommand) GetStageIntentOnly() bool`

GetStageIntentOnly returns the StageIntentOnly field if non-nil, zero value otherwise.

### GetStageIntentOnlyOk

`func (o *EnableDnsCertCommand) GetStageIntentOnlyOk() (*bool, bool)`

GetStageIntentOnlyOk returns a tuple with the StageIntentOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageIntentOnly

`func (o *EnableDnsCertCommand) SetStageIntentOnly(v bool)`

SetStageIntentOnly sets StageIntentOnly field to given value.

### HasStageIntentOnly

`func (o *EnableDnsCertCommand) HasStageIntentOnly() bool`

HasStageIntentOnly returns a boolean if a field has been set.

### GetInternalCaNoDnsStage

`func (o *EnableDnsCertCommand) GetInternalCaNoDnsStage() bool`

GetInternalCaNoDnsStage returns the InternalCaNoDnsStage field if non-nil, zero value otherwise.

### GetInternalCaNoDnsStageOk

`func (o *EnableDnsCertCommand) GetInternalCaNoDnsStageOk() (*bool, bool)`

GetInternalCaNoDnsStageOk returns a tuple with the InternalCaNoDnsStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCaNoDnsStage

`func (o *EnableDnsCertCommand) SetInternalCaNoDnsStage(v bool)`

SetInternalCaNoDnsStage sets InternalCaNoDnsStage field to given value.

### HasInternalCaNoDnsStage

`func (o *EnableDnsCertCommand) HasInternalCaNoDnsStage() bool`

HasInternalCaNoDnsStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


