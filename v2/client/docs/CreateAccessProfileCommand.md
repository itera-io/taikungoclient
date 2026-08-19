# CreateAccessProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**HttpProxy** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**SshUsers** | Pointer to [**[]SshUserCreateDto**](SshUserCreateDto.md) |  | [optional] 
**DnsServers** | Pointer to [**[]DnsServerCreateDto**](DnsServerCreateDto.md) |  | [optional] 
**TrustedRegistries** | Pointer to [**[]TrustedRegisteredCreateDto**](TrustedRegisteredCreateDto.md) |  | [optional] 
**NtpServers** | Pointer to [**[]NtpServerCreateDto**](NtpServerCreateDto.md) |  | [optional] 
**AllowedHosts** | Pointer to [**[]AllowedHostCreateDto**](AllowedHostCreateDto.md) |  | [optional] 

## Methods

### NewCreateAccessProfileCommand

`func NewCreateAccessProfileCommand() *CreateAccessProfileCommand`

NewCreateAccessProfileCommand instantiates a new CreateAccessProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessProfileCommandWithDefaults

`func NewCreateAccessProfileCommandWithDefaults() *CreateAccessProfileCommand`

NewCreateAccessProfileCommandWithDefaults instantiates a new CreateAccessProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAccessProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessProfileCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAccessProfileCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateAccessProfileCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAccessProfileCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHttpProxy

`func (o *CreateAccessProfileCommand) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *CreateAccessProfileCommand) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *CreateAccessProfileCommand) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *CreateAccessProfileCommand) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *CreateAccessProfileCommand) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *CreateAccessProfileCommand) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetOrganizationId

`func (o *CreateAccessProfileCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAccessProfileCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAccessProfileCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAccessProfileCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAccessProfileCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAccessProfileCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetSshUsers

`func (o *CreateAccessProfileCommand) GetSshUsers() []SshUserCreateDto`

GetSshUsers returns the SshUsers field if non-nil, zero value otherwise.

### GetSshUsersOk

`func (o *CreateAccessProfileCommand) GetSshUsersOk() (*[]SshUserCreateDto, bool)`

GetSshUsersOk returns a tuple with the SshUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsers

`func (o *CreateAccessProfileCommand) SetSshUsers(v []SshUserCreateDto)`

SetSshUsers sets SshUsers field to given value.

### HasSshUsers

`func (o *CreateAccessProfileCommand) HasSshUsers() bool`

HasSshUsers returns a boolean if a field has been set.

### SetSshUsersNil

`func (o *CreateAccessProfileCommand) SetSshUsersNil(b bool)`

 SetSshUsersNil sets the value for SshUsers to be an explicit nil

### UnsetSshUsers
`func (o *CreateAccessProfileCommand) UnsetSshUsers()`

UnsetSshUsers ensures that no value is present for SshUsers, not even an explicit nil
### GetDnsServers

`func (o *CreateAccessProfileCommand) GetDnsServers() []DnsServerCreateDto`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *CreateAccessProfileCommand) GetDnsServersOk() (*[]DnsServerCreateDto, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *CreateAccessProfileCommand) SetDnsServers(v []DnsServerCreateDto)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *CreateAccessProfileCommand) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *CreateAccessProfileCommand) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *CreateAccessProfileCommand) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetTrustedRegistries

`func (o *CreateAccessProfileCommand) GetTrustedRegistries() []TrustedRegisteredCreateDto`

GetTrustedRegistries returns the TrustedRegistries field if non-nil, zero value otherwise.

### GetTrustedRegistriesOk

`func (o *CreateAccessProfileCommand) GetTrustedRegistriesOk() (*[]TrustedRegisteredCreateDto, bool)`

GetTrustedRegistriesOk returns a tuple with the TrustedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRegistries

`func (o *CreateAccessProfileCommand) SetTrustedRegistries(v []TrustedRegisteredCreateDto)`

SetTrustedRegistries sets TrustedRegistries field to given value.

### HasTrustedRegistries

`func (o *CreateAccessProfileCommand) HasTrustedRegistries() bool`

HasTrustedRegistries returns a boolean if a field has been set.

### SetTrustedRegistriesNil

`func (o *CreateAccessProfileCommand) SetTrustedRegistriesNil(b bool)`

 SetTrustedRegistriesNil sets the value for TrustedRegistries to be an explicit nil

### UnsetTrustedRegistries
`func (o *CreateAccessProfileCommand) UnsetTrustedRegistries()`

UnsetTrustedRegistries ensures that no value is present for TrustedRegistries, not even an explicit nil
### GetNtpServers

`func (o *CreateAccessProfileCommand) GetNtpServers() []NtpServerCreateDto`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *CreateAccessProfileCommand) GetNtpServersOk() (*[]NtpServerCreateDto, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *CreateAccessProfileCommand) SetNtpServers(v []NtpServerCreateDto)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *CreateAccessProfileCommand) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *CreateAccessProfileCommand) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *CreateAccessProfileCommand) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetAllowedHosts

`func (o *CreateAccessProfileCommand) GetAllowedHosts() []AllowedHostCreateDto`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *CreateAccessProfileCommand) GetAllowedHostsOk() (*[]AllowedHostCreateDto, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *CreateAccessProfileCommand) SetAllowedHosts(v []AllowedHostCreateDto)`

SetAllowedHosts sets AllowedHosts field to given value.

### HasAllowedHosts

`func (o *CreateAccessProfileCommand) HasAllowedHosts() bool`

HasAllowedHosts returns a boolean if a field has been set.

### SetAllowedHostsNil

`func (o *CreateAccessProfileCommand) SetAllowedHostsNil(b bool)`

 SetAllowedHostsNil sets the value for AllowedHosts to be an explicit nil

### UnsetAllowedHosts
`func (o *CreateAccessProfileCommand) UnsetAllowedHosts()`

UnsetAllowedHosts ensures that no value is present for AllowedHosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


