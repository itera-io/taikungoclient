# AccessProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**HttpProxy** | **NullableString** |  | 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **string** |  | 
**IsLocked** | **bool** |  | 
**DnsServers** | [**[]DnsServerListDto**](DnsServerListDto.md) |  | 
**TrustedRegistries** | [**[]TrustedRegistryListDto**](TrustedRegistryListDto.md) |  | 
**NtpServers** | [**[]NtpServerListDto**](NtpServerListDto.md) |  | 
**AllowedHosts** | [**[]AllowedHostListDto**](AllowedHostListDto.md) |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewAccessProfilesListDto

`func NewAccessProfilesListDto(id int32, name string, httpProxy NullableString, organizationId NullableInt32, organizationName string, isLocked bool, dnsServers []DnsServerListDto, trustedRegistries []TrustedRegistryListDto, ntpServers []NtpServerListDto, allowedHosts []AllowedHostListDto, projects []CommonDropdownDto, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, createdAt string, ) *AccessProfilesListDto`

NewAccessProfilesListDto instantiates a new AccessProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfilesListDtoWithDefaults

`func NewAccessProfilesListDtoWithDefaults() *AccessProfilesListDto`

NewAccessProfilesListDtoWithDefaults instantiates a new AccessProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AccessProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfilesListDto) SetName(v string)`

SetName sets Name field to given value.


### GetHttpProxy

`func (o *AccessProfilesListDto) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *AccessProfilesListDto) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *AccessProfilesListDto) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.


### SetHttpProxyNil

`func (o *AccessProfilesListDto) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *AccessProfilesListDto) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetOrganizationId

`func (o *AccessProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccessProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccessProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *AccessProfilesListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AccessProfilesListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *AccessProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AccessProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AccessProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetIsLocked

`func (o *AccessProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AccessProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AccessProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetDnsServers

`func (o *AccessProfilesListDto) GetDnsServers() []DnsServerListDto`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AccessProfilesListDto) GetDnsServersOk() (*[]DnsServerListDto, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AccessProfilesListDto) SetDnsServers(v []DnsServerListDto)`

SetDnsServers sets DnsServers field to given value.


### GetTrustedRegistries

`func (o *AccessProfilesListDto) GetTrustedRegistries() []TrustedRegistryListDto`

GetTrustedRegistries returns the TrustedRegistries field if non-nil, zero value otherwise.

### GetTrustedRegistriesOk

`func (o *AccessProfilesListDto) GetTrustedRegistriesOk() (*[]TrustedRegistryListDto, bool)`

GetTrustedRegistriesOk returns a tuple with the TrustedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRegistries

`func (o *AccessProfilesListDto) SetTrustedRegistries(v []TrustedRegistryListDto)`

SetTrustedRegistries sets TrustedRegistries field to given value.


### GetNtpServers

`func (o *AccessProfilesListDto) GetNtpServers() []NtpServerListDto`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *AccessProfilesListDto) GetNtpServersOk() (*[]NtpServerListDto, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *AccessProfilesListDto) SetNtpServers(v []NtpServerListDto)`

SetNtpServers sets NtpServers field to given value.


### GetAllowedHosts

`func (o *AccessProfilesListDto) GetAllowedHosts() []AllowedHostListDto`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *AccessProfilesListDto) GetAllowedHostsOk() (*[]AllowedHostListDto, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *AccessProfilesListDto) SetAllowedHosts(v []AllowedHostListDto)`

SetAllowedHosts sets AllowedHosts field to given value.


### GetProjects

`func (o *AccessProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AccessProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AccessProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *AccessProfilesListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccessProfilesListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccessProfilesListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *AccessProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AccessProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AccessProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *AccessProfilesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AccessProfilesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *AccessProfilesListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AccessProfilesListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AccessProfilesListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetCreatedAt

`func (o *AccessProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


