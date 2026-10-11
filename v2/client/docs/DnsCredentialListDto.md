# DnsCredentialListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ProviderType** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **NullableString** |  | 
**DomainFilter** | **string** |  | 
**UseLetsEncrypt** | **bool** |  | 
**IsLocked** | **bool** |  | 
**IsDefault** | **bool** |  | 
**IsInfra** | **bool** |  | 
**AuthMethod** | **string** |  | 
**AuthVariant** | **NullableString** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 

## Methods

### NewDnsCredentialListDto

`func NewDnsCredentialListDto(id int32, name string, providerType string, organizationId NullableInt32, organizationName NullableString, domainFilter string, useLetsEncrypt bool, isLocked bool, isDefault bool, isInfra bool, authMethod string, authVariant NullableString, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, projects []CommonDropdownDto, ) *DnsCredentialListDto`

NewDnsCredentialListDto instantiates a new DnsCredentialListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCredentialListDtoWithDefaults

`func NewDnsCredentialListDtoWithDefaults() *DnsCredentialListDto`

NewDnsCredentialListDtoWithDefaults instantiates a new DnsCredentialListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsCredentialListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsCredentialListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsCredentialListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DnsCredentialListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCredentialListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCredentialListDto) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *DnsCredentialListDto) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DnsCredentialListDto) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DnsCredentialListDto) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetOrganizationId

`func (o *DnsCredentialListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DnsCredentialListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DnsCredentialListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *DnsCredentialListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DnsCredentialListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *DnsCredentialListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *DnsCredentialListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *DnsCredentialListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *DnsCredentialListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *DnsCredentialListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetDomainFilter

`func (o *DnsCredentialListDto) GetDomainFilter() string`

GetDomainFilter returns the DomainFilter field if non-nil, zero value otherwise.

### GetDomainFilterOk

`func (o *DnsCredentialListDto) GetDomainFilterOk() (*string, bool)`

GetDomainFilterOk returns a tuple with the DomainFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainFilter

`func (o *DnsCredentialListDto) SetDomainFilter(v string)`

SetDomainFilter sets DomainFilter field to given value.


### GetUseLetsEncrypt

`func (o *DnsCredentialListDto) GetUseLetsEncrypt() bool`

GetUseLetsEncrypt returns the UseLetsEncrypt field if non-nil, zero value otherwise.

### GetUseLetsEncryptOk

`func (o *DnsCredentialListDto) GetUseLetsEncryptOk() (*bool, bool)`

GetUseLetsEncryptOk returns a tuple with the UseLetsEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLetsEncrypt

`func (o *DnsCredentialListDto) SetUseLetsEncrypt(v bool)`

SetUseLetsEncrypt sets UseLetsEncrypt field to given value.


### GetIsLocked

`func (o *DnsCredentialListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *DnsCredentialListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *DnsCredentialListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsDefault

`func (o *DnsCredentialListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DnsCredentialListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DnsCredentialListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsInfra

`func (o *DnsCredentialListDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *DnsCredentialListDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *DnsCredentialListDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.


### GetAuthMethod

`func (o *DnsCredentialListDto) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *DnsCredentialListDto) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *DnsCredentialListDto) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetAuthVariant

`func (o *DnsCredentialListDto) GetAuthVariant() string`

GetAuthVariant returns the AuthVariant field if non-nil, zero value otherwise.

### GetAuthVariantOk

`func (o *DnsCredentialListDto) GetAuthVariantOk() (*string, bool)`

GetAuthVariantOk returns a tuple with the AuthVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVariant

`func (o *DnsCredentialListDto) SetAuthVariant(v string)`

SetAuthVariant sets AuthVariant field to given value.


### SetAuthVariantNil

`func (o *DnsCredentialListDto) SetAuthVariantNil(b bool)`

 SetAuthVariantNil sets the value for AuthVariant to be an explicit nil

### UnsetAuthVariant
`func (o *DnsCredentialListDto) UnsetAuthVariant()`

UnsetAuthVariant ensures that no value is present for AuthVariant, not even an explicit nil
### GetCreatedBy

`func (o *DnsCredentialListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DnsCredentialListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DnsCredentialListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *DnsCredentialListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DnsCredentialListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DnsCredentialListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *DnsCredentialListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *DnsCredentialListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *DnsCredentialListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DnsCredentialListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DnsCredentialListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetProjects

`func (o *DnsCredentialListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *DnsCredentialListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *DnsCredentialListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


