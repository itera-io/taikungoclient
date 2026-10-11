# CreateOpenstackCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**OpenStackUser** | Pointer to **NullableString** |  | [optional] 
**OpenStackPassword** | Pointer to **NullableString** |  | [optional] 
**OpenStackUrl** | Pointer to **NullableString** |  | [optional] 
**OpenStackProject** | Pointer to **NullableString** |  | [optional] 
**OpenStackPublicNetwork** | Pointer to **NullableString** |  | [optional] 
**OpenStackAvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**OpenStackDomain** | Pointer to **NullableString** |  | [optional] 
**OpenStackRegion** | Pointer to **NullableString** |  | [optional] 
**OpenStackContinent** | Pointer to **NullableString** |  | [optional] 
**OpenStackVolumeType** | Pointer to **NullableString** |  | [optional] 
**OpenStackImportNetwork** | Pointer to **bool** |  | [optional] 
**OpenStackInternalSubnetId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **bool** |  | [optional] 
**ManilaEnabled** | Pointer to **bool** |  | [optional] 
**ManilaStorageType** | Pointer to **NullableString** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**SkipTlsFlag** | Pointer to **bool** |  | [optional] 
**LbProvider** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateOpenstackCloudCommand

`func NewCreateOpenstackCloudCommand() *CreateOpenstackCloudCommand`

NewCreateOpenstackCloudCommand instantiates a new CreateOpenstackCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpenstackCloudCommandWithDefaults

`func NewCreateOpenstackCloudCommandWithDefaults() *CreateOpenstackCloudCommand`

NewCreateOpenstackCloudCommandWithDefaults instantiates a new CreateOpenstackCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOpenstackCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOpenstackCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOpenstackCloudCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOpenstackCloudCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateOpenstackCloudCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOpenstackCloudCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOpenStackUser

`func (o *CreateOpenstackCloudCommand) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *CreateOpenstackCloudCommand) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.

### HasOpenStackUser

`func (o *CreateOpenstackCloudCommand) HasOpenStackUser() bool`

HasOpenStackUser returns a boolean if a field has been set.

### SetOpenStackUserNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *CreateOpenstackCloudCommand) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *CreateOpenstackCloudCommand) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.

### HasOpenStackPassword

`func (o *CreateOpenstackCloudCommand) HasOpenStackPassword() bool`

HasOpenStackPassword returns a boolean if a field has been set.

### SetOpenStackPasswordNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *CreateOpenstackCloudCommand) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *CreateOpenstackCloudCommand) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.

### HasOpenStackUrl

`func (o *CreateOpenstackCloudCommand) HasOpenStackUrl() bool`

HasOpenStackUrl returns a boolean if a field has been set.

### SetOpenStackUrlNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackProject

`func (o *CreateOpenstackCloudCommand) GetOpenStackProject() string`

GetOpenStackProject returns the OpenStackProject field if non-nil, zero value otherwise.

### GetOpenStackProjectOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackProjectOk() (*string, bool)`

GetOpenStackProjectOk returns a tuple with the OpenStackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackProject

`func (o *CreateOpenstackCloudCommand) SetOpenStackProject(v string)`

SetOpenStackProject sets OpenStackProject field to given value.

### HasOpenStackProject

`func (o *CreateOpenstackCloudCommand) HasOpenStackProject() bool`

HasOpenStackProject returns a boolean if a field has been set.

### SetOpenStackProjectNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackProjectNil(b bool)`

 SetOpenStackProjectNil sets the value for OpenStackProject to be an explicit nil

### UnsetOpenStackProject
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackProject()`

UnsetOpenStackProject ensures that no value is present for OpenStackProject, not even an explicit nil
### GetOpenStackPublicNetwork

`func (o *CreateOpenstackCloudCommand) GetOpenStackPublicNetwork() string`

GetOpenStackPublicNetwork returns the OpenStackPublicNetwork field if non-nil, zero value otherwise.

### GetOpenStackPublicNetworkOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackPublicNetworkOk() (*string, bool)`

GetOpenStackPublicNetworkOk returns a tuple with the OpenStackPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPublicNetwork

`func (o *CreateOpenstackCloudCommand) SetOpenStackPublicNetwork(v string)`

SetOpenStackPublicNetwork sets OpenStackPublicNetwork field to given value.

### HasOpenStackPublicNetwork

`func (o *CreateOpenstackCloudCommand) HasOpenStackPublicNetwork() bool`

HasOpenStackPublicNetwork returns a boolean if a field has been set.

### SetOpenStackPublicNetworkNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackPublicNetworkNil(b bool)`

 SetOpenStackPublicNetworkNil sets the value for OpenStackPublicNetwork to be an explicit nil

### UnsetOpenStackPublicNetwork
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackPublicNetwork()`

UnsetOpenStackPublicNetwork ensures that no value is present for OpenStackPublicNetwork, not even an explicit nil
### GetOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) GetOpenStackAvailabilityZone() string`

GetOpenStackAvailabilityZone returns the OpenStackAvailabilityZone field if non-nil, zero value otherwise.

### GetOpenStackAvailabilityZoneOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackAvailabilityZoneOk() (*string, bool)`

GetOpenStackAvailabilityZoneOk returns a tuple with the OpenStackAvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) SetOpenStackAvailabilityZone(v string)`

SetOpenStackAvailabilityZone sets OpenStackAvailabilityZone field to given value.

### HasOpenStackAvailabilityZone

`func (o *CreateOpenstackCloudCommand) HasOpenStackAvailabilityZone() bool`

HasOpenStackAvailabilityZone returns a boolean if a field has been set.

### SetOpenStackAvailabilityZoneNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackAvailabilityZoneNil(b bool)`

 SetOpenStackAvailabilityZoneNil sets the value for OpenStackAvailabilityZone to be an explicit nil

### UnsetOpenStackAvailabilityZone
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackAvailabilityZone()`

UnsetOpenStackAvailabilityZone ensures that no value is present for OpenStackAvailabilityZone, not even an explicit nil
### GetOpenStackDomain

`func (o *CreateOpenstackCloudCommand) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *CreateOpenstackCloudCommand) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.

### HasOpenStackDomain

`func (o *CreateOpenstackCloudCommand) HasOpenStackDomain() bool`

HasOpenStackDomain returns a boolean if a field has been set.

### SetOpenStackDomainNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetOpenStackRegion

`func (o *CreateOpenstackCloudCommand) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *CreateOpenstackCloudCommand) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.

### HasOpenStackRegion

`func (o *CreateOpenstackCloudCommand) HasOpenStackRegion() bool`

HasOpenStackRegion returns a boolean if a field has been set.

### SetOpenStackRegionNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackRegionNil(b bool)`

 SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil

### UnsetOpenStackRegion
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackRegion()`

UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
### GetOpenStackContinent

`func (o *CreateOpenstackCloudCommand) GetOpenStackContinent() string`

GetOpenStackContinent returns the OpenStackContinent field if non-nil, zero value otherwise.

### GetOpenStackContinentOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackContinentOk() (*string, bool)`

GetOpenStackContinentOk returns a tuple with the OpenStackContinent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackContinent

`func (o *CreateOpenstackCloudCommand) SetOpenStackContinent(v string)`

SetOpenStackContinent sets OpenStackContinent field to given value.

### HasOpenStackContinent

`func (o *CreateOpenstackCloudCommand) HasOpenStackContinent() bool`

HasOpenStackContinent returns a boolean if a field has been set.

### SetOpenStackContinentNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackContinentNil(b bool)`

 SetOpenStackContinentNil sets the value for OpenStackContinent to be an explicit nil

### UnsetOpenStackContinent
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackContinent()`

UnsetOpenStackContinent ensures that no value is present for OpenStackContinent, not even an explicit nil
### GetOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) GetOpenStackVolumeType() string`

GetOpenStackVolumeType returns the OpenStackVolumeType field if non-nil, zero value otherwise.

### GetOpenStackVolumeTypeOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackVolumeTypeOk() (*string, bool)`

GetOpenStackVolumeTypeOk returns a tuple with the OpenStackVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) SetOpenStackVolumeType(v string)`

SetOpenStackVolumeType sets OpenStackVolumeType field to given value.

### HasOpenStackVolumeType

`func (o *CreateOpenstackCloudCommand) HasOpenStackVolumeType() bool`

HasOpenStackVolumeType returns a boolean if a field has been set.

### SetOpenStackVolumeTypeNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackVolumeTypeNil(b bool)`

 SetOpenStackVolumeTypeNil sets the value for OpenStackVolumeType to be an explicit nil

### UnsetOpenStackVolumeType
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackVolumeType()`

UnsetOpenStackVolumeType ensures that no value is present for OpenStackVolumeType, not even an explicit nil
### GetOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) GetOpenStackImportNetwork() bool`

GetOpenStackImportNetwork returns the OpenStackImportNetwork field if non-nil, zero value otherwise.

### GetOpenStackImportNetworkOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackImportNetworkOk() (*bool, bool)`

GetOpenStackImportNetworkOk returns a tuple with the OpenStackImportNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) SetOpenStackImportNetwork(v bool)`

SetOpenStackImportNetwork sets OpenStackImportNetwork field to given value.

### HasOpenStackImportNetwork

`func (o *CreateOpenstackCloudCommand) HasOpenStackImportNetwork() bool`

HasOpenStackImportNetwork returns a boolean if a field has been set.

### GetOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) GetOpenStackInternalSubnetId() string`

GetOpenStackInternalSubnetId returns the OpenStackInternalSubnetId field if non-nil, zero value otherwise.

### GetOpenStackInternalSubnetIdOk

`func (o *CreateOpenstackCloudCommand) GetOpenStackInternalSubnetIdOk() (*string, bool)`

GetOpenStackInternalSubnetIdOk returns a tuple with the OpenStackInternalSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) SetOpenStackInternalSubnetId(v string)`

SetOpenStackInternalSubnetId sets OpenStackInternalSubnetId field to given value.

### HasOpenStackInternalSubnetId

`func (o *CreateOpenstackCloudCommand) HasOpenStackInternalSubnetId() bool`

HasOpenStackInternalSubnetId returns a boolean if a field has been set.

### SetOpenStackInternalSubnetIdNil

`func (o *CreateOpenstackCloudCommand) SetOpenStackInternalSubnetIdNil(b bool)`

 SetOpenStackInternalSubnetIdNil sets the value for OpenStackInternalSubnetId to be an explicit nil

### UnsetOpenStackInternalSubnetId
`func (o *CreateOpenstackCloudCommand) UnsetOpenStackInternalSubnetId()`

UnsetOpenStackInternalSubnetId ensures that no value is present for OpenStackInternalSubnetId, not even an explicit nil
### GetOrganizationId

`func (o *CreateOpenstackCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOpenstackCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOpenstackCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOpenstackCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateOpenstackCloudCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateOpenstackCloudCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *CreateOpenstackCloudCommand) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *CreateOpenstackCloudCommand) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.

### GetManilaEnabled

`func (o *CreateOpenstackCloudCommand) GetManilaEnabled() bool`

GetManilaEnabled returns the ManilaEnabled field if non-nil, zero value otherwise.

### GetManilaEnabledOk

`func (o *CreateOpenstackCloudCommand) GetManilaEnabledOk() (*bool, bool)`

GetManilaEnabledOk returns a tuple with the ManilaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManilaEnabled

`func (o *CreateOpenstackCloudCommand) SetManilaEnabled(v bool)`

SetManilaEnabled sets ManilaEnabled field to given value.

### HasManilaEnabled

`func (o *CreateOpenstackCloudCommand) HasManilaEnabled() bool`

HasManilaEnabled returns a boolean if a field has been set.

### GetManilaStorageType

`func (o *CreateOpenstackCloudCommand) GetManilaStorageType() string`

GetManilaStorageType returns the ManilaStorageType field if non-nil, zero value otherwise.

### GetManilaStorageTypeOk

`func (o *CreateOpenstackCloudCommand) GetManilaStorageTypeOk() (*string, bool)`

GetManilaStorageTypeOk returns a tuple with the ManilaStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManilaStorageType

`func (o *CreateOpenstackCloudCommand) SetManilaStorageType(v string)`

SetManilaStorageType sets ManilaStorageType field to given value.

### HasManilaStorageType

`func (o *CreateOpenstackCloudCommand) HasManilaStorageType() bool`

HasManilaStorageType returns a boolean if a field has been set.

### SetManilaStorageTypeNil

`func (o *CreateOpenstackCloudCommand) SetManilaStorageTypeNil(b bool)`

 SetManilaStorageTypeNil sets the value for ManilaStorageType to be an explicit nil

### UnsetManilaStorageType
`func (o *CreateOpenstackCloudCommand) UnsetManilaStorageType()`

UnsetManilaStorageType ensures that no value is present for ManilaStorageType, not even an explicit nil
### GetIsAdmin

`func (o *CreateOpenstackCloudCommand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CreateOpenstackCloudCommand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CreateOpenstackCloudCommand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CreateOpenstackCloudCommand) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetSkipTlsFlag

`func (o *CreateOpenstackCloudCommand) GetSkipTlsFlag() bool`

GetSkipTlsFlag returns the SkipTlsFlag field if non-nil, zero value otherwise.

### GetSkipTlsFlagOk

`func (o *CreateOpenstackCloudCommand) GetSkipTlsFlagOk() (*bool, bool)`

GetSkipTlsFlagOk returns a tuple with the SkipTlsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsFlag

`func (o *CreateOpenstackCloudCommand) SetSkipTlsFlag(v bool)`

SetSkipTlsFlag sets SkipTlsFlag field to given value.

### HasSkipTlsFlag

`func (o *CreateOpenstackCloudCommand) HasSkipTlsFlag() bool`

HasSkipTlsFlag returns a boolean if a field has been set.

### GetLbProvider

`func (o *CreateOpenstackCloudCommand) GetLbProvider() map[string]interface{}`

GetLbProvider returns the LbProvider field if non-nil, zero value otherwise.

### GetLbProviderOk

`func (o *CreateOpenstackCloudCommand) GetLbProviderOk() (*map[string]interface{}, bool)`

GetLbProviderOk returns a tuple with the LbProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbProvider

`func (o *CreateOpenstackCloudCommand) SetLbProvider(v map[string]interface{})`

SetLbProvider sets LbProvider field to given value.

### HasLbProvider

`func (o *CreateOpenstackCloudCommand) HasLbProvider() bool`

HasLbProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


