# CatalogDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**PackageIds** | Pointer to **[]string** |  | [optional] 
**BoundProjects** | Pointer to [**[]ProjectCatalogDto**](ProjectCatalogDto.md) |  | [optional] 
**BoundApplications** | Pointer to [**[]AvailablePackagesDto**](AvailablePackagesDto.md) |  | [optional] 

## Methods

### NewCatalogDetails

`func NewCatalogDetails() *CatalogDetails`

NewCatalogDetails instantiates a new CatalogDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogDetailsWithDefaults

`func NewCatalogDetailsWithDefaults() *CatalogDetails`

NewCatalogDetailsWithDefaults instantiates a new CatalogDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CatalogDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogDetails) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogDetails) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CatalogDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsLocked

`func (o *CatalogDetails) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CatalogDetails) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CatalogDetails) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *CatalogDetails) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsDefault

`func (o *CatalogDetails) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CatalogDetails) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CatalogDetails) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CatalogDetails) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CatalogDetails) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CatalogDetails) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CatalogDetails) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CatalogDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPackageIds

`func (o *CatalogDetails) GetPackageIds() []string`

GetPackageIds returns the PackageIds field if non-nil, zero value otherwise.

### GetPackageIdsOk

`func (o *CatalogDetails) GetPackageIdsOk() (*[]string, bool)`

GetPackageIdsOk returns a tuple with the PackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageIds

`func (o *CatalogDetails) SetPackageIds(v []string)`

SetPackageIds sets PackageIds field to given value.

### HasPackageIds

`func (o *CatalogDetails) HasPackageIds() bool`

HasPackageIds returns a boolean if a field has been set.

### SetPackageIdsNil

`func (o *CatalogDetails) SetPackageIdsNil(b bool)`

 SetPackageIdsNil sets the value for PackageIds to be an explicit nil

### UnsetPackageIds
`func (o *CatalogDetails) UnsetPackageIds()`

UnsetPackageIds ensures that no value is present for PackageIds, not even an explicit nil
### GetBoundProjects

`func (o *CatalogDetails) GetBoundProjects() []ProjectCatalogDto`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *CatalogDetails) GetBoundProjectsOk() (*[]ProjectCatalogDto, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *CatalogDetails) SetBoundProjects(v []ProjectCatalogDto)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *CatalogDetails) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### SetBoundProjectsNil

`func (o *CatalogDetails) SetBoundProjectsNil(b bool)`

 SetBoundProjectsNil sets the value for BoundProjects to be an explicit nil

### UnsetBoundProjects
`func (o *CatalogDetails) UnsetBoundProjects()`

UnsetBoundProjects ensures that no value is present for BoundProjects, not even an explicit nil
### GetBoundApplications

`func (o *CatalogDetails) GetBoundApplications() []AvailablePackagesDto`

GetBoundApplications returns the BoundApplications field if non-nil, zero value otherwise.

### GetBoundApplicationsOk

`func (o *CatalogDetails) GetBoundApplicationsOk() (*[]AvailablePackagesDto, bool)`

GetBoundApplicationsOk returns a tuple with the BoundApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundApplications

`func (o *CatalogDetails) SetBoundApplications(v []AvailablePackagesDto)`

SetBoundApplications sets BoundApplications field to given value.

### HasBoundApplications

`func (o *CatalogDetails) HasBoundApplications() bool`

HasBoundApplications returns a boolean if a field has been set.

### SetBoundApplicationsNil

`func (o *CatalogDetails) SetBoundApplicationsNil(b bool)`

 SetBoundApplicationsNil sets the value for BoundApplications to be an explicit nil

### UnsetBoundApplications
`func (o *CatalogDetails) UnsetBoundApplications()`

UnsetBoundApplications ensures that no value is present for BoundApplications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


