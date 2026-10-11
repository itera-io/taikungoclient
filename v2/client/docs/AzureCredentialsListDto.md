# AzureCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ProjectCount** | **int32** |  | 
**IsLocked** | **bool** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**Location** | **string** |  | 
**AvailabilityZones** | **[]string** |  | 
**AvailabilityZonesCount** | **int32** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**IsDefault** | **bool** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**ContinentName** | **NullableString** |  | 
**VnetDeploymentMode** | **string** |  | 
**ExistingVirtualNetworkResourceGroupName** | **NullableString** |  | 
**ExistingVirtualNetworkName** | **NullableString** |  | 
**SharedFileSystemEnabled** | **bool** |  | 
**ImportedFileShareId** | **NullableString** |  | 
**ImportedFileShareName** | **NullableString** |  | 
**ImportedFileShareResourceGroupName** | **NullableString** |  | 

## Methods

### NewAzureCredentialsListDto

`func NewAzureCredentialsListDto(id int32, projectCount int32, isLocked bool, name string, tenantId string, location string, availabilityZones []string, availabilityZonesCount int32, projects []CommonDropdownDto, createdBy AuditUserDto, createdAt NullableString, lastModified NullableString, lastModifiedBy AuditUserDto, isDefault bool, organizationId int32, organizationName string, continentName NullableString, vnetDeploymentMode string, existingVirtualNetworkResourceGroupName NullableString, existingVirtualNetworkName NullableString, sharedFileSystemEnabled bool, importedFileShareId NullableString, importedFileShareName NullableString, importedFileShareResourceGroupName NullableString, ) *AzureCredentialsListDto`

NewAzureCredentialsListDto instantiates a new AzureCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsListDtoWithDefaults

`func NewAzureCredentialsListDtoWithDefaults() *AzureCredentialsListDto`

NewAzureCredentialsListDtoWithDefaults instantiates a new AzureCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectCount

`func (o *AzureCredentialsListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *AzureCredentialsListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *AzureCredentialsListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.


### GetIsLocked

`func (o *AzureCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AzureCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AzureCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetName

`func (o *AzureCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *AzureCredentialsListDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureCredentialsListDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureCredentialsListDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetLocation

`func (o *AzureCredentialsListDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureCredentialsListDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureCredentialsListDto) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetAvailabilityZones

`func (o *AzureCredentialsListDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AzureCredentialsListDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AzureCredentialsListDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetAvailabilityZonesCount

`func (o *AzureCredentialsListDto) GetAvailabilityZonesCount() int32`

GetAvailabilityZonesCount returns the AvailabilityZonesCount field if non-nil, zero value otherwise.

### GetAvailabilityZonesCountOk

`func (o *AzureCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool)`

GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZonesCount

`func (o *AzureCredentialsListDto) SetAvailabilityZonesCount(v int32)`

SetAvailabilityZonesCount sets AvailabilityZonesCount field to given value.


### GetProjects

`func (o *AzureCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AzureCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AzureCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *AzureCredentialsListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AzureCredentialsListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AzureCredentialsListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *AzureCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AzureCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AzureCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AzureCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AzureCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *AzureCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AzureCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AzureCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *AzureCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AzureCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *AzureCredentialsListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AzureCredentialsListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AzureCredentialsListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsDefault

`func (o *AzureCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AzureCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AzureCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOrganizationId

`func (o *AzureCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AzureCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AzureCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *AzureCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AzureCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AzureCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetContinentName

`func (o *AzureCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *AzureCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *AzureCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.


### SetContinentNameNil

`func (o *AzureCredentialsListDto) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *AzureCredentialsListDto) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
### GetVnetDeploymentMode

`func (o *AzureCredentialsListDto) GetVnetDeploymentMode() string`

GetVnetDeploymentMode returns the VnetDeploymentMode field if non-nil, zero value otherwise.

### GetVnetDeploymentModeOk

`func (o *AzureCredentialsListDto) GetVnetDeploymentModeOk() (*string, bool)`

GetVnetDeploymentModeOk returns a tuple with the VnetDeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetDeploymentMode

`func (o *AzureCredentialsListDto) SetVnetDeploymentMode(v string)`

SetVnetDeploymentMode sets VnetDeploymentMode field to given value.


### GetExistingVirtualNetworkResourceGroupName

`func (o *AzureCredentialsListDto) GetExistingVirtualNetworkResourceGroupName() string`

GetExistingVirtualNetworkResourceGroupName returns the ExistingVirtualNetworkResourceGroupName field if non-nil, zero value otherwise.

### GetExistingVirtualNetworkResourceGroupNameOk

`func (o *AzureCredentialsListDto) GetExistingVirtualNetworkResourceGroupNameOk() (*string, bool)`

GetExistingVirtualNetworkResourceGroupNameOk returns a tuple with the ExistingVirtualNetworkResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingVirtualNetworkResourceGroupName

`func (o *AzureCredentialsListDto) SetExistingVirtualNetworkResourceGroupName(v string)`

SetExistingVirtualNetworkResourceGroupName sets ExistingVirtualNetworkResourceGroupName field to given value.


### SetExistingVirtualNetworkResourceGroupNameNil

`func (o *AzureCredentialsListDto) SetExistingVirtualNetworkResourceGroupNameNil(b bool)`

 SetExistingVirtualNetworkResourceGroupNameNil sets the value for ExistingVirtualNetworkResourceGroupName to be an explicit nil

### UnsetExistingVirtualNetworkResourceGroupName
`func (o *AzureCredentialsListDto) UnsetExistingVirtualNetworkResourceGroupName()`

UnsetExistingVirtualNetworkResourceGroupName ensures that no value is present for ExistingVirtualNetworkResourceGroupName, not even an explicit nil
### GetExistingVirtualNetworkName

`func (o *AzureCredentialsListDto) GetExistingVirtualNetworkName() string`

GetExistingVirtualNetworkName returns the ExistingVirtualNetworkName field if non-nil, zero value otherwise.

### GetExistingVirtualNetworkNameOk

`func (o *AzureCredentialsListDto) GetExistingVirtualNetworkNameOk() (*string, bool)`

GetExistingVirtualNetworkNameOk returns a tuple with the ExistingVirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingVirtualNetworkName

`func (o *AzureCredentialsListDto) SetExistingVirtualNetworkName(v string)`

SetExistingVirtualNetworkName sets ExistingVirtualNetworkName field to given value.


### SetExistingVirtualNetworkNameNil

`func (o *AzureCredentialsListDto) SetExistingVirtualNetworkNameNil(b bool)`

 SetExistingVirtualNetworkNameNil sets the value for ExistingVirtualNetworkName to be an explicit nil

### UnsetExistingVirtualNetworkName
`func (o *AzureCredentialsListDto) UnsetExistingVirtualNetworkName()`

UnsetExistingVirtualNetworkName ensures that no value is present for ExistingVirtualNetworkName, not even an explicit nil
### GetSharedFileSystemEnabled

`func (o *AzureCredentialsListDto) GetSharedFileSystemEnabled() bool`

GetSharedFileSystemEnabled returns the SharedFileSystemEnabled field if non-nil, zero value otherwise.

### GetSharedFileSystemEnabledOk

`func (o *AzureCredentialsListDto) GetSharedFileSystemEnabledOk() (*bool, bool)`

GetSharedFileSystemEnabledOk returns a tuple with the SharedFileSystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFileSystemEnabled

`func (o *AzureCredentialsListDto) SetSharedFileSystemEnabled(v bool)`

SetSharedFileSystemEnabled sets SharedFileSystemEnabled field to given value.


### GetImportedFileShareId

`func (o *AzureCredentialsListDto) GetImportedFileShareId() string`

GetImportedFileShareId returns the ImportedFileShareId field if non-nil, zero value otherwise.

### GetImportedFileShareIdOk

`func (o *AzureCredentialsListDto) GetImportedFileShareIdOk() (*string, bool)`

GetImportedFileShareIdOk returns a tuple with the ImportedFileShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFileShareId

`func (o *AzureCredentialsListDto) SetImportedFileShareId(v string)`

SetImportedFileShareId sets ImportedFileShareId field to given value.


### SetImportedFileShareIdNil

`func (o *AzureCredentialsListDto) SetImportedFileShareIdNil(b bool)`

 SetImportedFileShareIdNil sets the value for ImportedFileShareId to be an explicit nil

### UnsetImportedFileShareId
`func (o *AzureCredentialsListDto) UnsetImportedFileShareId()`

UnsetImportedFileShareId ensures that no value is present for ImportedFileShareId, not even an explicit nil
### GetImportedFileShareName

`func (o *AzureCredentialsListDto) GetImportedFileShareName() string`

GetImportedFileShareName returns the ImportedFileShareName field if non-nil, zero value otherwise.

### GetImportedFileShareNameOk

`func (o *AzureCredentialsListDto) GetImportedFileShareNameOk() (*string, bool)`

GetImportedFileShareNameOk returns a tuple with the ImportedFileShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFileShareName

`func (o *AzureCredentialsListDto) SetImportedFileShareName(v string)`

SetImportedFileShareName sets ImportedFileShareName field to given value.


### SetImportedFileShareNameNil

`func (o *AzureCredentialsListDto) SetImportedFileShareNameNil(b bool)`

 SetImportedFileShareNameNil sets the value for ImportedFileShareName to be an explicit nil

### UnsetImportedFileShareName
`func (o *AzureCredentialsListDto) UnsetImportedFileShareName()`

UnsetImportedFileShareName ensures that no value is present for ImportedFileShareName, not even an explicit nil
### GetImportedFileShareResourceGroupName

`func (o *AzureCredentialsListDto) GetImportedFileShareResourceGroupName() string`

GetImportedFileShareResourceGroupName returns the ImportedFileShareResourceGroupName field if non-nil, zero value otherwise.

### GetImportedFileShareResourceGroupNameOk

`func (o *AzureCredentialsListDto) GetImportedFileShareResourceGroupNameOk() (*string, bool)`

GetImportedFileShareResourceGroupNameOk returns a tuple with the ImportedFileShareResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFileShareResourceGroupName

`func (o *AzureCredentialsListDto) SetImportedFileShareResourceGroupName(v string)`

SetImportedFileShareResourceGroupName sets ImportedFileShareResourceGroupName field to given value.


### SetImportedFileShareResourceGroupNameNil

`func (o *AzureCredentialsListDto) SetImportedFileShareResourceGroupNameNil(b bool)`

 SetImportedFileShareResourceGroupNameNil sets the value for ImportedFileShareResourceGroupName to be an explicit nil

### UnsetImportedFileShareResourceGroupName
`func (o *AzureCredentialsListDto) UnsetImportedFileShareResourceGroupName()`

UnsetImportedFileShareResourceGroupName ensures that no value is present for ImportedFileShareResourceGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


