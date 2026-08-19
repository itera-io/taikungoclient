# ProjectDetailsForVmsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Name** | **string** |  | 
**Id** | **int32** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**CloudName** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**IsLocked** | **bool** |  | 
**IsProjectMaintenanceModeEnabled** | **bool** |  | 
**HasSelectedFlavors** | **NullableBool** |  | 
**IsMaintenanceModeEnabled** | **bool** |  | 
**ProjectCloudRevision** | **NullableInt32** |  | 
**CloudCredentialRevision** | **NullableInt32** |  | 
**AllowFullSpotKubernetes** | **bool** |  | 
**AllowSpotWorkers** | **bool** |  | 
**AllowSpotVMs** | **bool** |  | 
**MaxSpotPrice** | **NullableFloat64** |  | 
**AvailabilityZones** | **[]string** |  | 
**Hypervisors** | **[]string** |  | 
**ExpiredAt** | **NullableString** |  | 
**VpcMode** | Pointer to **map[string]interface{}** |  | [optional] 
**CloudSubnets** | Pointer to [**[]CloudSubnetDetailsDto**](CloudSubnetDetailsDto.md) |  | [optional] 
**PrivateOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProjectDetailsForVmsDto

`func NewProjectDetailsForVmsDto(status ProjectStatus, name string, id int32, cloudType ECloudCredentialType, cloudName NullableString, cloudId NullableInt32, organizationName string, organizationId int32, isLocked bool, isProjectMaintenanceModeEnabled bool, hasSelectedFlavors NullableBool, isMaintenanceModeEnabled bool, projectCloudRevision NullableInt32, cloudCredentialRevision NullableInt32, allowFullSpotKubernetes bool, allowSpotWorkers bool, allowSpotVMs bool, maxSpotPrice NullableFloat64, availabilityZones []string, hypervisors []string, expiredAt NullableString, ) *ProjectDetailsForVmsDto`

NewProjectDetailsForVmsDto instantiates a new ProjectDetailsForVmsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsForVmsDtoWithDefaults

`func NewProjectDetailsForVmsDtoWithDefaults() *ProjectDetailsForVmsDto`

NewProjectDetailsForVmsDtoWithDefaults instantiates a new ProjectDetailsForVmsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectDetailsForVmsDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectDetailsForVmsDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectDetailsForVmsDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *ProjectDetailsForVmsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetailsForVmsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetailsForVmsDto) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ProjectDetailsForVmsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetailsForVmsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetailsForVmsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCloudType

`func (o *ProjectDetailsForVmsDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectDetailsForVmsDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectDetailsForVmsDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetCloudName

`func (o *ProjectDetailsForVmsDto) GetCloudName() string`

GetCloudName returns the CloudName field if non-nil, zero value otherwise.

### GetCloudNameOk

`func (o *ProjectDetailsForVmsDto) GetCloudNameOk() (*string, bool)`

GetCloudNameOk returns a tuple with the CloudName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudName

`func (o *ProjectDetailsForVmsDto) SetCloudName(v string)`

SetCloudName sets CloudName field to given value.


### SetCloudNameNil

`func (o *ProjectDetailsForVmsDto) SetCloudNameNil(b bool)`

 SetCloudNameNil sets the value for CloudName to be an explicit nil

### UnsetCloudName
`func (o *ProjectDetailsForVmsDto) UnsetCloudName()`

UnsetCloudName ensures that no value is present for CloudName, not even an explicit nil
### GetCloudId

`func (o *ProjectDetailsForVmsDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ProjectDetailsForVmsDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ProjectDetailsForVmsDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *ProjectDetailsForVmsDto) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *ProjectDetailsForVmsDto) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetOrganizationName

`func (o *ProjectDetailsForVmsDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectDetailsForVmsDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectDetailsForVmsDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *ProjectDetailsForVmsDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectDetailsForVmsDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectDetailsForVmsDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetIsLocked

`func (o *ProjectDetailsForVmsDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectDetailsForVmsDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectDetailsForVmsDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsProjectMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) GetIsProjectMaintenanceModeEnabled() bool`

GetIsProjectMaintenanceModeEnabled returns the IsProjectMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsProjectMaintenanceModeEnabledOk

`func (o *ProjectDetailsForVmsDto) GetIsProjectMaintenanceModeEnabledOk() (*bool, bool)`

GetIsProjectMaintenanceModeEnabledOk returns a tuple with the IsProjectMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) SetIsProjectMaintenanceModeEnabled(v bool)`

SetIsProjectMaintenanceModeEnabled sets IsProjectMaintenanceModeEnabled field to given value.


### GetHasSelectedFlavors

`func (o *ProjectDetailsForVmsDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectDetailsForVmsDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectDetailsForVmsDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.


### SetHasSelectedFlavorsNil

`func (o *ProjectDetailsForVmsDto) SetHasSelectedFlavorsNil(b bool)`

 SetHasSelectedFlavorsNil sets the value for HasSelectedFlavors to be an explicit nil

### UnsetHasSelectedFlavors
`func (o *ProjectDetailsForVmsDto) UnsetHasSelectedFlavors()`

UnsetHasSelectedFlavors ensures that no value is present for HasSelectedFlavors, not even an explicit nil
### GetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ProjectDetailsForVmsDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForVmsDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.


### GetProjectCloudRevision

`func (o *ProjectDetailsForVmsDto) GetProjectCloudRevision() int32`

GetProjectCloudRevision returns the ProjectCloudRevision field if non-nil, zero value otherwise.

### GetProjectCloudRevisionOk

`func (o *ProjectDetailsForVmsDto) GetProjectCloudRevisionOk() (*int32, bool)`

GetProjectCloudRevisionOk returns a tuple with the ProjectCloudRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCloudRevision

`func (o *ProjectDetailsForVmsDto) SetProjectCloudRevision(v int32)`

SetProjectCloudRevision sets ProjectCloudRevision field to given value.


### SetProjectCloudRevisionNil

`func (o *ProjectDetailsForVmsDto) SetProjectCloudRevisionNil(b bool)`

 SetProjectCloudRevisionNil sets the value for ProjectCloudRevision to be an explicit nil

### UnsetProjectCloudRevision
`func (o *ProjectDetailsForVmsDto) UnsetProjectCloudRevision()`

UnsetProjectCloudRevision ensures that no value is present for ProjectCloudRevision, not even an explicit nil
### GetCloudCredentialRevision

`func (o *ProjectDetailsForVmsDto) GetCloudCredentialRevision() int32`

GetCloudCredentialRevision returns the CloudCredentialRevision field if non-nil, zero value otherwise.

### GetCloudCredentialRevisionOk

`func (o *ProjectDetailsForVmsDto) GetCloudCredentialRevisionOk() (*int32, bool)`

GetCloudCredentialRevisionOk returns a tuple with the CloudCredentialRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialRevision

`func (o *ProjectDetailsForVmsDto) SetCloudCredentialRevision(v int32)`

SetCloudCredentialRevision sets CloudCredentialRevision field to given value.


### SetCloudCredentialRevisionNil

`func (o *ProjectDetailsForVmsDto) SetCloudCredentialRevisionNil(b bool)`

 SetCloudCredentialRevisionNil sets the value for CloudCredentialRevision to be an explicit nil

### UnsetCloudCredentialRevision
`func (o *ProjectDetailsForVmsDto) UnsetCloudCredentialRevision()`

UnsetCloudCredentialRevision ensures that no value is present for CloudCredentialRevision, not even an explicit nil
### GetAllowFullSpotKubernetes

`func (o *ProjectDetailsForVmsDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectDetailsForVmsDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectDetailsForVmsDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.


### GetAllowSpotWorkers

`func (o *ProjectDetailsForVmsDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectDetailsForVmsDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectDetailsForVmsDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.


### GetAllowSpotVMs

`func (o *ProjectDetailsForVmsDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectDetailsForVmsDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectDetailsForVmsDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.


### GetMaxSpotPrice

`func (o *ProjectDetailsForVmsDto) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *ProjectDetailsForVmsDto) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *ProjectDetailsForVmsDto) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.


### SetMaxSpotPriceNil

`func (o *ProjectDetailsForVmsDto) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *ProjectDetailsForVmsDto) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetAvailabilityZones

`func (o *ProjectDetailsForVmsDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectDetailsForVmsDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectDetailsForVmsDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### SetAvailabilityZonesNil

`func (o *ProjectDetailsForVmsDto) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *ProjectDetailsForVmsDto) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetHypervisors

`func (o *ProjectDetailsForVmsDto) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *ProjectDetailsForVmsDto) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *ProjectDetailsForVmsDto) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.


### SetHypervisorsNil

`func (o *ProjectDetailsForVmsDto) SetHypervisorsNil(b bool)`

 SetHypervisorsNil sets the value for Hypervisors to be an explicit nil

### UnsetHypervisors
`func (o *ProjectDetailsForVmsDto) UnsetHypervisors()`

UnsetHypervisors ensures that no value is present for Hypervisors, not even an explicit nil
### GetExpiredAt

`func (o *ProjectDetailsForVmsDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectDetailsForVmsDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectDetailsForVmsDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *ProjectDetailsForVmsDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ProjectDetailsForVmsDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetVpcMode

`func (o *ProjectDetailsForVmsDto) GetVpcMode() map[string]interface{}`

GetVpcMode returns the VpcMode field if non-nil, zero value otherwise.

### GetVpcModeOk

`func (o *ProjectDetailsForVmsDto) GetVpcModeOk() (*map[string]interface{}, bool)`

GetVpcModeOk returns a tuple with the VpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMode

`func (o *ProjectDetailsForVmsDto) SetVpcMode(v map[string]interface{})`

SetVpcMode sets VpcMode field to given value.

### HasVpcMode

`func (o *ProjectDetailsForVmsDto) HasVpcMode() bool`

HasVpcMode returns a boolean if a field has been set.

### GetCloudSubnets

`func (o *ProjectDetailsForVmsDto) GetCloudSubnets() []CloudSubnetDetailsDto`

GetCloudSubnets returns the CloudSubnets field if non-nil, zero value otherwise.

### GetCloudSubnetsOk

`func (o *ProjectDetailsForVmsDto) GetCloudSubnetsOk() (*[]CloudSubnetDetailsDto, bool)`

GetCloudSubnetsOk returns a tuple with the CloudSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSubnets

`func (o *ProjectDetailsForVmsDto) SetCloudSubnets(v []CloudSubnetDetailsDto)`

SetCloudSubnets sets CloudSubnets field to given value.

### HasCloudSubnets

`func (o *ProjectDetailsForVmsDto) HasCloudSubnets() bool`

HasCloudSubnets returns a boolean if a field has been set.

### SetCloudSubnetsNil

`func (o *ProjectDetailsForVmsDto) SetCloudSubnetsNil(b bool)`

 SetCloudSubnetsNil sets the value for CloudSubnets to be an explicit nil

### UnsetCloudSubnets
`func (o *ProjectDetailsForVmsDto) UnsetCloudSubnets()`

UnsetCloudSubnets ensures that no value is present for CloudSubnets, not even an explicit nil
### GetPrivateOnly

`func (o *ProjectDetailsForVmsDto) GetPrivateOnly() bool`

GetPrivateOnly returns the PrivateOnly field if non-nil, zero value otherwise.

### GetPrivateOnlyOk

`func (o *ProjectDetailsForVmsDto) GetPrivateOnlyOk() (*bool, bool)`

GetPrivateOnlyOk returns a tuple with the PrivateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOnly

`func (o *ProjectDetailsForVmsDto) SetPrivateOnly(v bool)`

SetPrivateOnly sets PrivateOnly field to given value.

### HasPrivateOnly

`func (o *ProjectDetailsForVmsDto) HasPrivateOnly() bool`

HasPrivateOnly returns a boolean if a field has been set.

### SetPrivateOnlyNil

`func (o *ProjectDetailsForVmsDto) SetPrivateOnlyNil(b bool)`

 SetPrivateOnlyNil sets the value for PrivateOnly to be an explicit nil

### UnsetPrivateOnly
`func (o *ProjectDetailsForVmsDto) UnsetPrivateOnly()`

UnsetPrivateOnly ensures that no value is present for PrivateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


