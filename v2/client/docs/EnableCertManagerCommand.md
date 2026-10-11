# EnableCertManagerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**DnsCredentialId** | **NullableInt32** |  | 
**CertIssuerType** | **NullableString** |  | 
**AcmeEmail** | **NullableString** |  | 
**Domain** | **NullableString** |  | 
**CertificateProfileId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEnableCertManagerCommand

`func NewEnableCertManagerCommand(projectId int32, dnsCredentialId NullableInt32, certIssuerType NullableString, acmeEmail NullableString, domain NullableString, ) *EnableCertManagerCommand`

NewEnableCertManagerCommand instantiates a new EnableCertManagerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableCertManagerCommandWithDefaults

`func NewEnableCertManagerCommandWithDefaults() *EnableCertManagerCommand`

NewEnableCertManagerCommandWithDefaults instantiates a new EnableCertManagerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *EnableCertManagerCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnableCertManagerCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnableCertManagerCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetDnsCredentialId

`func (o *EnableCertManagerCommand) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *EnableCertManagerCommand) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *EnableCertManagerCommand) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.


### SetDnsCredentialIdNil

`func (o *EnableCertManagerCommand) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *EnableCertManagerCommand) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetCertIssuerType

`func (o *EnableCertManagerCommand) GetCertIssuerType() string`

GetCertIssuerType returns the CertIssuerType field if non-nil, zero value otherwise.

### GetCertIssuerTypeOk

`func (o *EnableCertManagerCommand) GetCertIssuerTypeOk() (*string, bool)`

GetCertIssuerTypeOk returns a tuple with the CertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerType

`func (o *EnableCertManagerCommand) SetCertIssuerType(v string)`

SetCertIssuerType sets CertIssuerType field to given value.


### SetCertIssuerTypeNil

`func (o *EnableCertManagerCommand) SetCertIssuerTypeNil(b bool)`

 SetCertIssuerTypeNil sets the value for CertIssuerType to be an explicit nil

### UnsetCertIssuerType
`func (o *EnableCertManagerCommand) UnsetCertIssuerType()`

UnsetCertIssuerType ensures that no value is present for CertIssuerType, not even an explicit nil
### GetAcmeEmail

`func (o *EnableCertManagerCommand) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *EnableCertManagerCommand) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *EnableCertManagerCommand) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.


### SetAcmeEmailNil

`func (o *EnableCertManagerCommand) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *EnableCertManagerCommand) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetDomain

`func (o *EnableCertManagerCommand) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnableCertManagerCommand) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnableCertManagerCommand) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *EnableCertManagerCommand) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *EnableCertManagerCommand) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCertificateProfileId

`func (o *EnableCertManagerCommand) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *EnableCertManagerCommand) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *EnableCertManagerCommand) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.

### HasCertificateProfileId

`func (o *EnableCertManagerCommand) HasCertificateProfileId() bool`

HasCertificateProfileId returns a boolean if a field has been set.

### SetCertificateProfileIdNil

`func (o *EnableCertManagerCommand) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *EnableCertManagerCommand) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


