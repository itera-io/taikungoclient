# GenericKubernetesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ProjectCount** | **int32** |  | 
**IsLocked** | **bool** |  | 
**Name** | **NullableString** |  | 
**MainProject** | [**MainProjectDto**](MainProjectDto.md) |  | 
**AssociatedVClusters** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**IsDefault** | **bool** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**ContinentName** | **NullableString** |  | 

## Methods

### NewGenericKubernetesListDto

`func NewGenericKubernetesListDto(id int32, projectCount int32, isLocked bool, name NullableString, mainProject MainProjectDto, associatedVClusters []CommonDropdownDto, createdBy AuditUserDto, createdAt NullableString, lastModified NullableString, lastModifiedBy AuditUserDto, isDefault bool, organizationId int32, organizationName NullableString, continentName NullableString, ) *GenericKubernetesListDto`

NewGenericKubernetesListDto instantiates a new GenericKubernetesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericKubernetesListDtoWithDefaults

`func NewGenericKubernetesListDtoWithDefaults() *GenericKubernetesListDto`

NewGenericKubernetesListDtoWithDefaults instantiates a new GenericKubernetesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericKubernetesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericKubernetesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericKubernetesListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectCount

`func (o *GenericKubernetesListDto) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *GenericKubernetesListDto) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *GenericKubernetesListDto) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.


### GetIsLocked

`func (o *GenericKubernetesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GenericKubernetesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GenericKubernetesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetName

`func (o *GenericKubernetesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericKubernetesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericKubernetesListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GenericKubernetesListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GenericKubernetesListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMainProject

`func (o *GenericKubernetesListDto) GetMainProject() MainProjectDto`

GetMainProject returns the MainProject field if non-nil, zero value otherwise.

### GetMainProjectOk

`func (o *GenericKubernetesListDto) GetMainProjectOk() (*MainProjectDto, bool)`

GetMainProjectOk returns a tuple with the MainProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainProject

`func (o *GenericKubernetesListDto) SetMainProject(v MainProjectDto)`

SetMainProject sets MainProject field to given value.


### GetAssociatedVClusters

`func (o *GenericKubernetesListDto) GetAssociatedVClusters() []CommonDropdownDto`

GetAssociatedVClusters returns the AssociatedVClusters field if non-nil, zero value otherwise.

### GetAssociatedVClustersOk

`func (o *GenericKubernetesListDto) GetAssociatedVClustersOk() (*[]CommonDropdownDto, bool)`

GetAssociatedVClustersOk returns a tuple with the AssociatedVClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedVClusters

`func (o *GenericKubernetesListDto) SetAssociatedVClusters(v []CommonDropdownDto)`

SetAssociatedVClusters sets AssociatedVClusters field to given value.


### SetAssociatedVClustersNil

`func (o *GenericKubernetesListDto) SetAssociatedVClustersNil(b bool)`

 SetAssociatedVClustersNil sets the value for AssociatedVClusters to be an explicit nil

### UnsetAssociatedVClusters
`func (o *GenericKubernetesListDto) UnsetAssociatedVClusters()`

UnsetAssociatedVClusters ensures that no value is present for AssociatedVClusters, not even an explicit nil
### GetCreatedBy

`func (o *GenericKubernetesListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GenericKubernetesListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GenericKubernetesListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *GenericKubernetesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GenericKubernetesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GenericKubernetesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GenericKubernetesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GenericKubernetesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *GenericKubernetesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *GenericKubernetesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *GenericKubernetesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *GenericKubernetesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *GenericKubernetesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *GenericKubernetesListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *GenericKubernetesListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *GenericKubernetesListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsDefault

`func (o *GenericKubernetesListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GenericKubernetesListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GenericKubernetesListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOrganizationId

`func (o *GenericKubernetesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GenericKubernetesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GenericKubernetesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *GenericKubernetesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GenericKubernetesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GenericKubernetesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *GenericKubernetesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *GenericKubernetesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetContinentName

`func (o *GenericKubernetesListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *GenericKubernetesListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *GenericKubernetesListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.


### SetContinentNameNil

`func (o *GenericKubernetesListDto) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *GenericKubernetesListDto) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


