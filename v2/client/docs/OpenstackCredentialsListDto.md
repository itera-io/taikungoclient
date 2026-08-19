# OpenstackCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ProjectCount** | **int32** |  | 
**IsLocked** | **bool** |  | 
**Name** | **string** |  | 
**User** | **string** |  | 
**Url** | **string** |  | 
**Project** | **string** |  | 
**Domain** | **string** |  | 
**Region** | **string** |  | 
**PublicNetwork** | **string** |  | 
**ImportNetwork** | **bool** |  | 
**TenantId** | **string** |  | 
**AvailabilityZone** | **NullableString** |  | 
**VolumeType** | **NullableString** |  | 
**InternalSubnetId** | **NullableString** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **string** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**IsDefault** | **bool** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**CreatedAt** | **string** |  | 
**ContinentName** | **string** |  | 
**IsAdmin** | **bool** |  | 
**IsInfra** | **bool** |  | 
**ApplicationCredEnabled** | **bool** |  | 
**SkipTlsFlag** | **bool** |  | 
**ManilaEnabled** | Pointer to **bool** |  | [optional] 
**ManilaStorageType** | Pointer to **NullableString** |  | [optional] 
**LbProvider** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenstackCredentialsListDto

`func NewOpenstackCredentialsListDto(id int32, projectCount int32, isLocked bool, name string, user string, url string, project string, domain string, region string, publicNetwork string, importNetwork bool, tenantId string, availabilityZone NullableString, volumeType NullableString, internalSubnetId NullableString, projects []CommonDropdownDto, createdBy AuditUserDto, lastModified string, lastModifiedBy AuditUserDto, isDefault bool, organizationId int32, organizationName string, createdAt string, continentName string, isAdmin bool, isInfra bool, applicationCredEnabled bool, skipTlsFlag bool, ) *OpenstackCredentialsListDto`

NewOpenstackCredentialsListDto instantiates a new OpenstackCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackCredentialsListDtoWithDefaults

`func NewOpenstackCredentialsListDtoWithDefaults() *OpenstackCredentialsListDto`

NewOpenstackCredentialsListDtoWithDefaults instantiates a new OpenstackCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenstackCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenstackCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenstackCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectCount

`func (o *OpenstackCredentialsListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *OpenstackCredentialsListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *OpenstackCredentialsListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.


### GetIsLocked

`func (o *OpenstackCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OpenstackCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OpenstackCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetName

`func (o *OpenstackCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenstackCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenstackCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *OpenstackCredentialsListDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OpenstackCredentialsListDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OpenstackCredentialsListDto) SetUser(v string)`

SetUser sets User field to given value.


### GetUrl

`func (o *OpenstackCredentialsListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenstackCredentialsListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenstackCredentialsListDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProject

`func (o *OpenstackCredentialsListDto) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OpenstackCredentialsListDto) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OpenstackCredentialsListDto) SetProject(v string)`

SetProject sets Project field to given value.


### GetDomain

`func (o *OpenstackCredentialsListDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OpenstackCredentialsListDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OpenstackCredentialsListDto) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRegion

`func (o *OpenstackCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OpenstackCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OpenstackCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetPublicNetwork

`func (o *OpenstackCredentialsListDto) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *OpenstackCredentialsListDto) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *OpenstackCredentialsListDto) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.


### GetImportNetwork

`func (o *OpenstackCredentialsListDto) GetImportNetwork() bool`

GetImportNetwork returns the ImportNetwork field if non-nil, zero value otherwise.

### GetImportNetworkOk

`func (o *OpenstackCredentialsListDto) GetImportNetworkOk() (*bool, bool)`

GetImportNetworkOk returns a tuple with the ImportNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNetwork

`func (o *OpenstackCredentialsListDto) SetImportNetwork(v bool)`

SetImportNetwork sets ImportNetwork field to given value.


### GetTenantId

`func (o *OpenstackCredentialsListDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OpenstackCredentialsListDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OpenstackCredentialsListDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAvailabilityZone

`func (o *OpenstackCredentialsListDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *OpenstackCredentialsListDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *OpenstackCredentialsListDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *OpenstackCredentialsListDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *OpenstackCredentialsListDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetVolumeType

`func (o *OpenstackCredentialsListDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *OpenstackCredentialsListDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *OpenstackCredentialsListDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### SetVolumeTypeNil

`func (o *OpenstackCredentialsListDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *OpenstackCredentialsListDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetInternalSubnetId

`func (o *OpenstackCredentialsListDto) GetInternalSubnetId() string`

GetInternalSubnetId returns the InternalSubnetId field if non-nil, zero value otherwise.

### GetInternalSubnetIdOk

`func (o *OpenstackCredentialsListDto) GetInternalSubnetIdOk() (*string, bool)`

GetInternalSubnetIdOk returns a tuple with the InternalSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSubnetId

`func (o *OpenstackCredentialsListDto) SetInternalSubnetId(v string)`

SetInternalSubnetId sets InternalSubnetId field to given value.


### SetInternalSubnetIdNil

`func (o *OpenstackCredentialsListDto) SetInternalSubnetIdNil(b bool)`

 SetInternalSubnetIdNil sets the value for InternalSubnetId to be an explicit nil

### UnsetInternalSubnetId
`func (o *OpenstackCredentialsListDto) UnsetInternalSubnetId()`

UnsetInternalSubnetId ensures that no value is present for InternalSubnetId, not even an explicit nil
### GetProjects

`func (o *OpenstackCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OpenstackCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OpenstackCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *OpenstackCredentialsListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OpenstackCredentialsListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OpenstackCredentialsListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *OpenstackCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *OpenstackCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *OpenstackCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### GetLastModifiedBy

`func (o *OpenstackCredentialsListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *OpenstackCredentialsListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *OpenstackCredentialsListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsDefault

`func (o *OpenstackCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OpenstackCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OpenstackCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOrganizationId

`func (o *OpenstackCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpenstackCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpenstackCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *OpenstackCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OpenstackCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OpenstackCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCreatedAt

`func (o *OpenstackCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenstackCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenstackCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetContinentName

`func (o *OpenstackCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *OpenstackCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *OpenstackCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.


### GetIsAdmin

`func (o *OpenstackCredentialsListDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenstackCredentialsListDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenstackCredentialsListDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetIsInfra

`func (o *OpenstackCredentialsListDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *OpenstackCredentialsListDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *OpenstackCredentialsListDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.


### GetApplicationCredEnabled

`func (o *OpenstackCredentialsListDto) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenstackCredentialsListDto) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenstackCredentialsListDto) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.


### GetSkipTlsFlag

`func (o *OpenstackCredentialsListDto) GetSkipTlsFlag() bool`

GetSkipTlsFlag returns the SkipTlsFlag field if non-nil, zero value otherwise.

### GetSkipTlsFlagOk

`func (o *OpenstackCredentialsListDto) GetSkipTlsFlagOk() (*bool, bool)`

GetSkipTlsFlagOk returns a tuple with the SkipTlsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsFlag

`func (o *OpenstackCredentialsListDto) SetSkipTlsFlag(v bool)`

SetSkipTlsFlag sets SkipTlsFlag field to given value.


### GetManilaEnabled

`func (o *OpenstackCredentialsListDto) GetManilaEnabled() bool`

GetManilaEnabled returns the ManilaEnabled field if non-nil, zero value otherwise.

### GetManilaEnabledOk

`func (o *OpenstackCredentialsListDto) GetManilaEnabledOk() (*bool, bool)`

GetManilaEnabledOk returns a tuple with the ManilaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManilaEnabled

`func (o *OpenstackCredentialsListDto) SetManilaEnabled(v bool)`

SetManilaEnabled sets ManilaEnabled field to given value.

### HasManilaEnabled

`func (o *OpenstackCredentialsListDto) HasManilaEnabled() bool`

HasManilaEnabled returns a boolean if a field has been set.

### GetManilaStorageType

`func (o *OpenstackCredentialsListDto) GetManilaStorageType() string`

GetManilaStorageType returns the ManilaStorageType field if non-nil, zero value otherwise.

### GetManilaStorageTypeOk

`func (o *OpenstackCredentialsListDto) GetManilaStorageTypeOk() (*string, bool)`

GetManilaStorageTypeOk returns a tuple with the ManilaStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManilaStorageType

`func (o *OpenstackCredentialsListDto) SetManilaStorageType(v string)`

SetManilaStorageType sets ManilaStorageType field to given value.

### HasManilaStorageType

`func (o *OpenstackCredentialsListDto) HasManilaStorageType() bool`

HasManilaStorageType returns a boolean if a field has been set.

### SetManilaStorageTypeNil

`func (o *OpenstackCredentialsListDto) SetManilaStorageTypeNil(b bool)`

 SetManilaStorageTypeNil sets the value for ManilaStorageType to be an explicit nil

### UnsetManilaStorageType
`func (o *OpenstackCredentialsListDto) UnsetManilaStorageType()`

UnsetManilaStorageType ensures that no value is present for ManilaStorageType, not even an explicit nil
### GetLbProvider

`func (o *OpenstackCredentialsListDto) GetLbProvider() string`

GetLbProvider returns the LbProvider field if non-nil, zero value otherwise.

### GetLbProviderOk

`func (o *OpenstackCredentialsListDto) GetLbProviderOk() (*string, bool)`

GetLbProviderOk returns a tuple with the LbProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbProvider

`func (o *OpenstackCredentialsListDto) SetLbProvider(v string)`

SetLbProvider sets LbProvider field to given value.

### HasLbProvider

`func (o *OpenstackCredentialsListDto) HasLbProvider() bool`

HasLbProvider returns a boolean if a field has been set.

### SetLbProviderNil

`func (o *OpenstackCredentialsListDto) SetLbProviderNil(b bool)`

 SetLbProviderNil sets the value for LbProvider to be an explicit nil

### UnsetLbProvider
`func (o *OpenstackCredentialsListDto) UnsetLbProvider()`

UnsetLbProvider ensures that no value is present for LbProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


