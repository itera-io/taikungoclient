# AmazonCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ProjectCount** | **int32** |  | 
**IsLocked** | **bool** |  | 
**Name** | **string** |  | 
**Region** | **string** |  | 
**AvailabilityZones** | **[]string** |  | 
**AvailabilityZonesCount** | **int32** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**IsDefault** | **bool** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**ContinentName** | **NullableString** |  | 
**EfsFileSystemId** | Pointer to **NullableString** |  | [optional] 
**EfsFileSystemName** | Pointer to **NullableString** |  | [optional] 
**SharedFileSystemEnabled** | Pointer to **bool** |  | [optional] 
**VpcMode** | Pointer to **map[string]interface{}** |  | [optional] 
**IpMode** | Pointer to [**IpMode**](IpMode.md) |  | [optional] 
**ImportedVpcId** | Pointer to **NullableString** |  | [optional] 
**ImportedVpcCidr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAmazonCredentialsListDto

`func NewAmazonCredentialsListDto(id int32, projectCount int32, isLocked bool, name string, region string, availabilityZones []string, availabilityZonesCount int32, projects []CommonDropdownDto, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, isDefault bool, organizationId int32, organizationName string, createdAt NullableString, continentName NullableString, ) *AmazonCredentialsListDto`

NewAmazonCredentialsListDto instantiates a new AmazonCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonCredentialsListDtoWithDefaults

`func NewAmazonCredentialsListDtoWithDefaults() *AmazonCredentialsListDto`

NewAmazonCredentialsListDtoWithDefaults instantiates a new AmazonCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectCount

`func (o *AmazonCredentialsListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *AmazonCredentialsListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *AmazonCredentialsListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.


### GetIsLocked

`func (o *AmazonCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AmazonCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AmazonCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetName

`func (o *AmazonCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *AmazonCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AmazonCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AmazonCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAvailabilityZones

`func (o *AmazonCredentialsListDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AmazonCredentialsListDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetAvailabilityZonesCount

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesCount() int32`

GetAvailabilityZonesCount returns the AvailabilityZonesCount field if non-nil, zero value otherwise.

### GetAvailabilityZonesCountOk

`func (o *AmazonCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool)`

GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZonesCount

`func (o *AmazonCredentialsListDto) SetAvailabilityZonesCount(v int32)`

SetAvailabilityZonesCount sets AvailabilityZonesCount field to given value.


### GetProjects

`func (o *AmazonCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AmazonCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AmazonCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *AmazonCredentialsListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AmazonCredentialsListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AmazonCredentialsListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *AmazonCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AmazonCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AmazonCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *AmazonCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AmazonCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *AmazonCredentialsListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AmazonCredentialsListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AmazonCredentialsListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsDefault

`func (o *AmazonCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AmazonCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AmazonCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOrganizationId

`func (o *AmazonCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AmazonCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AmazonCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *AmazonCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AmazonCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AmazonCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCreatedAt

`func (o *AmazonCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmazonCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmazonCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AmazonCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AmazonCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetContinentName

`func (o *AmazonCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *AmazonCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *AmazonCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.


### SetContinentNameNil

`func (o *AmazonCredentialsListDto) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *AmazonCredentialsListDto) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
### GetEfsFileSystemId

`func (o *AmazonCredentialsListDto) GetEfsFileSystemId() string`

GetEfsFileSystemId returns the EfsFileSystemId field if non-nil, zero value otherwise.

### GetEfsFileSystemIdOk

`func (o *AmazonCredentialsListDto) GetEfsFileSystemIdOk() (*string, bool)`

GetEfsFileSystemIdOk returns a tuple with the EfsFileSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfsFileSystemId

`func (o *AmazonCredentialsListDto) SetEfsFileSystemId(v string)`

SetEfsFileSystemId sets EfsFileSystemId field to given value.

### HasEfsFileSystemId

`func (o *AmazonCredentialsListDto) HasEfsFileSystemId() bool`

HasEfsFileSystemId returns a boolean if a field has been set.

### SetEfsFileSystemIdNil

`func (o *AmazonCredentialsListDto) SetEfsFileSystemIdNil(b bool)`

 SetEfsFileSystemIdNil sets the value for EfsFileSystemId to be an explicit nil

### UnsetEfsFileSystemId
`func (o *AmazonCredentialsListDto) UnsetEfsFileSystemId()`

UnsetEfsFileSystemId ensures that no value is present for EfsFileSystemId, not even an explicit nil
### GetEfsFileSystemName

`func (o *AmazonCredentialsListDto) GetEfsFileSystemName() string`

GetEfsFileSystemName returns the EfsFileSystemName field if non-nil, zero value otherwise.

### GetEfsFileSystemNameOk

`func (o *AmazonCredentialsListDto) GetEfsFileSystemNameOk() (*string, bool)`

GetEfsFileSystemNameOk returns a tuple with the EfsFileSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfsFileSystemName

`func (o *AmazonCredentialsListDto) SetEfsFileSystemName(v string)`

SetEfsFileSystemName sets EfsFileSystemName field to given value.

### HasEfsFileSystemName

`func (o *AmazonCredentialsListDto) HasEfsFileSystemName() bool`

HasEfsFileSystemName returns a boolean if a field has been set.

### SetEfsFileSystemNameNil

`func (o *AmazonCredentialsListDto) SetEfsFileSystemNameNil(b bool)`

 SetEfsFileSystemNameNil sets the value for EfsFileSystemName to be an explicit nil

### UnsetEfsFileSystemName
`func (o *AmazonCredentialsListDto) UnsetEfsFileSystemName()`

UnsetEfsFileSystemName ensures that no value is present for EfsFileSystemName, not even an explicit nil
### GetSharedFileSystemEnabled

`func (o *AmazonCredentialsListDto) GetSharedFileSystemEnabled() bool`

GetSharedFileSystemEnabled returns the SharedFileSystemEnabled field if non-nil, zero value otherwise.

### GetSharedFileSystemEnabledOk

`func (o *AmazonCredentialsListDto) GetSharedFileSystemEnabledOk() (*bool, bool)`

GetSharedFileSystemEnabledOk returns a tuple with the SharedFileSystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFileSystemEnabled

`func (o *AmazonCredentialsListDto) SetSharedFileSystemEnabled(v bool)`

SetSharedFileSystemEnabled sets SharedFileSystemEnabled field to given value.

### HasSharedFileSystemEnabled

`func (o *AmazonCredentialsListDto) HasSharedFileSystemEnabled() bool`

HasSharedFileSystemEnabled returns a boolean if a field has been set.

### GetVpcMode

`func (o *AmazonCredentialsListDto) GetVpcMode() map[string]interface{}`

GetVpcMode returns the VpcMode field if non-nil, zero value otherwise.

### GetVpcModeOk

`func (o *AmazonCredentialsListDto) GetVpcModeOk() (*map[string]interface{}, bool)`

GetVpcModeOk returns a tuple with the VpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMode

`func (o *AmazonCredentialsListDto) SetVpcMode(v map[string]interface{})`

SetVpcMode sets VpcMode field to given value.

### HasVpcMode

`func (o *AmazonCredentialsListDto) HasVpcMode() bool`

HasVpcMode returns a boolean if a field has been set.

### GetIpMode

`func (o *AmazonCredentialsListDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AmazonCredentialsListDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AmazonCredentialsListDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *AmazonCredentialsListDto) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetImportedVpcId

`func (o *AmazonCredentialsListDto) GetImportedVpcId() string`

GetImportedVpcId returns the ImportedVpcId field if non-nil, zero value otherwise.

### GetImportedVpcIdOk

`func (o *AmazonCredentialsListDto) GetImportedVpcIdOk() (*string, bool)`

GetImportedVpcIdOk returns a tuple with the ImportedVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedVpcId

`func (o *AmazonCredentialsListDto) SetImportedVpcId(v string)`

SetImportedVpcId sets ImportedVpcId field to given value.

### HasImportedVpcId

`func (o *AmazonCredentialsListDto) HasImportedVpcId() bool`

HasImportedVpcId returns a boolean if a field has been set.

### SetImportedVpcIdNil

`func (o *AmazonCredentialsListDto) SetImportedVpcIdNil(b bool)`

 SetImportedVpcIdNil sets the value for ImportedVpcId to be an explicit nil

### UnsetImportedVpcId
`func (o *AmazonCredentialsListDto) UnsetImportedVpcId()`

UnsetImportedVpcId ensures that no value is present for ImportedVpcId, not even an explicit nil
### GetImportedVpcCidr

`func (o *AmazonCredentialsListDto) GetImportedVpcCidr() string`

GetImportedVpcCidr returns the ImportedVpcCidr field if non-nil, zero value otherwise.

### GetImportedVpcCidrOk

`func (o *AmazonCredentialsListDto) GetImportedVpcCidrOk() (*string, bool)`

GetImportedVpcCidrOk returns a tuple with the ImportedVpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedVpcCidr

`func (o *AmazonCredentialsListDto) SetImportedVpcCidr(v string)`

SetImportedVpcCidr sets ImportedVpcCidr field to given value.

### HasImportedVpcCidr

`func (o *AmazonCredentialsListDto) HasImportedVpcCidr() bool`

HasImportedVpcCidr returns a boolean if a field has been set.

### SetImportedVpcCidrNil

`func (o *AmazonCredentialsListDto) SetImportedVpcCidrNil(b bool)`

 SetImportedVpcCidrNil sets the value for ImportedVpcCidr to be an explicit nil

### UnsetImportedVpcCidr
`func (o *AmazonCredentialsListDto) UnsetImportedVpcCidr()`

UnsetImportedVpcCidr ensures that no value is present for ImportedVpcCidr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


