# EnableExternalDnsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**DnsCredentialId** | **NullableInt32** |  | 
**Domain** | **NullableString** |  | 

## Methods

### NewEnableExternalDnsCommand

`func NewEnableExternalDnsCommand(projectId int32, dnsCredentialId NullableInt32, domain NullableString, ) *EnableExternalDnsCommand`

NewEnableExternalDnsCommand instantiates a new EnableExternalDnsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableExternalDnsCommandWithDefaults

`func NewEnableExternalDnsCommandWithDefaults() *EnableExternalDnsCommand`

NewEnableExternalDnsCommandWithDefaults instantiates a new EnableExternalDnsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *EnableExternalDnsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnableExternalDnsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnableExternalDnsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetDnsCredentialId

`func (o *EnableExternalDnsCommand) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *EnableExternalDnsCommand) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *EnableExternalDnsCommand) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.


### SetDnsCredentialIdNil

`func (o *EnableExternalDnsCommand) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *EnableExternalDnsCommand) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetDomain

`func (o *EnableExternalDnsCommand) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnableExternalDnsCommand) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnableExternalDnsCommand) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *EnableExternalDnsCommand) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *EnableExternalDnsCommand) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


