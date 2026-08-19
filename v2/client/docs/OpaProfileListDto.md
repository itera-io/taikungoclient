# OpaProfileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ForbidNodePort** | **bool** |  | 
**ForbidHttpIngress** | **bool** |  | 
**RequireProbe** | **bool** |  | 
**UniqueIngresses** | **bool** |  | 
**UniqueServiceSelector** | **bool** |  | 
**ForcePodResource** | **bool** |  | 
**IsNodeNameForbiddenInVC** | **bool** |  | 
**IsMasterTaintEnforced** | **bool** |  | 
**WhitelistMasterTaintNamespaces** | **[]string** |  | 
**AllowedRepo** | **[]string** |  | 
**ForbidSpecificTags** | **[]string** |  | 
**IngressWhitelist** | **[]string** |  | 
**IsLocked** | **bool** |  | 
**Revision** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**IsDefault** | **bool** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 

## Methods

### NewOpaProfileListDto

`func NewOpaProfileListDto(id int32, name string, forbidNodePort bool, forbidHttpIngress bool, requireProbe bool, uniqueIngresses bool, uniqueServiceSelector bool, forcePodResource bool, isNodeNameForbiddenInVC bool, isMasterTaintEnforced bool, whitelistMasterTaintNamespaces []string, allowedRepo []string, forbidSpecificTags []string, ingressWhitelist []string, isLocked bool, revision int32, organizationId int32, organizationName string, createdAt NullableString, isDefault bool, projects []CommonDropdownDto, ) *OpaProfileListDto`

NewOpaProfileListDto instantiates a new OpaProfileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaProfileListDtoWithDefaults

`func NewOpaProfileListDtoWithDefaults() *OpaProfileListDto`

NewOpaProfileListDtoWithDefaults instantiates a new OpaProfileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpaProfileListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaProfileListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaProfileListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OpaProfileListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpaProfileListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpaProfileListDto) SetName(v string)`

SetName sets Name field to given value.


### GetForbidNodePort

`func (o *OpaProfileListDto) GetForbidNodePort() bool`

GetForbidNodePort returns the ForbidNodePort field if non-nil, zero value otherwise.

### GetForbidNodePortOk

`func (o *OpaProfileListDto) GetForbidNodePortOk() (*bool, bool)`

GetForbidNodePortOk returns a tuple with the ForbidNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidNodePort

`func (o *OpaProfileListDto) SetForbidNodePort(v bool)`

SetForbidNodePort sets ForbidNodePort field to given value.


### GetForbidHttpIngress

`func (o *OpaProfileListDto) GetForbidHttpIngress() bool`

GetForbidHttpIngress returns the ForbidHttpIngress field if non-nil, zero value otherwise.

### GetForbidHttpIngressOk

`func (o *OpaProfileListDto) GetForbidHttpIngressOk() (*bool, bool)`

GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidHttpIngress

`func (o *OpaProfileListDto) SetForbidHttpIngress(v bool)`

SetForbidHttpIngress sets ForbidHttpIngress field to given value.


### GetRequireProbe

`func (o *OpaProfileListDto) GetRequireProbe() bool`

GetRequireProbe returns the RequireProbe field if non-nil, zero value otherwise.

### GetRequireProbeOk

`func (o *OpaProfileListDto) GetRequireProbeOk() (*bool, bool)`

GetRequireProbeOk returns a tuple with the RequireProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProbe

`func (o *OpaProfileListDto) SetRequireProbe(v bool)`

SetRequireProbe sets RequireProbe field to given value.


### GetUniqueIngresses

`func (o *OpaProfileListDto) GetUniqueIngresses() bool`

GetUniqueIngresses returns the UniqueIngresses field if non-nil, zero value otherwise.

### GetUniqueIngressesOk

`func (o *OpaProfileListDto) GetUniqueIngressesOk() (*bool, bool)`

GetUniqueIngressesOk returns a tuple with the UniqueIngresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIngresses

`func (o *OpaProfileListDto) SetUniqueIngresses(v bool)`

SetUniqueIngresses sets UniqueIngresses field to given value.


### GetUniqueServiceSelector

`func (o *OpaProfileListDto) GetUniqueServiceSelector() bool`

GetUniqueServiceSelector returns the UniqueServiceSelector field if non-nil, zero value otherwise.

### GetUniqueServiceSelectorOk

`func (o *OpaProfileListDto) GetUniqueServiceSelectorOk() (*bool, bool)`

GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueServiceSelector

`func (o *OpaProfileListDto) SetUniqueServiceSelector(v bool)`

SetUniqueServiceSelector sets UniqueServiceSelector field to given value.


### GetForcePodResource

`func (o *OpaProfileListDto) GetForcePodResource() bool`

GetForcePodResource returns the ForcePodResource field if non-nil, zero value otherwise.

### GetForcePodResourceOk

`func (o *OpaProfileListDto) GetForcePodResourceOk() (*bool, bool)`

GetForcePodResourceOk returns a tuple with the ForcePodResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePodResource

`func (o *OpaProfileListDto) SetForcePodResource(v bool)`

SetForcePodResource sets ForcePodResource field to given value.


### GetIsNodeNameForbiddenInVC

`func (o *OpaProfileListDto) GetIsNodeNameForbiddenInVC() bool`

GetIsNodeNameForbiddenInVC returns the IsNodeNameForbiddenInVC field if non-nil, zero value otherwise.

### GetIsNodeNameForbiddenInVCOk

`func (o *OpaProfileListDto) GetIsNodeNameForbiddenInVCOk() (*bool, bool)`

GetIsNodeNameForbiddenInVCOk returns a tuple with the IsNodeNameForbiddenInVC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNodeNameForbiddenInVC

`func (o *OpaProfileListDto) SetIsNodeNameForbiddenInVC(v bool)`

SetIsNodeNameForbiddenInVC sets IsNodeNameForbiddenInVC field to given value.


### GetIsMasterTaintEnforced

`func (o *OpaProfileListDto) GetIsMasterTaintEnforced() bool`

GetIsMasterTaintEnforced returns the IsMasterTaintEnforced field if non-nil, zero value otherwise.

### GetIsMasterTaintEnforcedOk

`func (o *OpaProfileListDto) GetIsMasterTaintEnforcedOk() (*bool, bool)`

GetIsMasterTaintEnforcedOk returns a tuple with the IsMasterTaintEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterTaintEnforced

`func (o *OpaProfileListDto) SetIsMasterTaintEnforced(v bool)`

SetIsMasterTaintEnforced sets IsMasterTaintEnforced field to given value.


### GetWhitelistMasterTaintNamespaces

`func (o *OpaProfileListDto) GetWhitelistMasterTaintNamespaces() []string`

GetWhitelistMasterTaintNamespaces returns the WhitelistMasterTaintNamespaces field if non-nil, zero value otherwise.

### GetWhitelistMasterTaintNamespacesOk

`func (o *OpaProfileListDto) GetWhitelistMasterTaintNamespacesOk() (*[]string, bool)`

GetWhitelistMasterTaintNamespacesOk returns a tuple with the WhitelistMasterTaintNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistMasterTaintNamespaces

`func (o *OpaProfileListDto) SetWhitelistMasterTaintNamespaces(v []string)`

SetWhitelistMasterTaintNamespaces sets WhitelistMasterTaintNamespaces field to given value.


### GetAllowedRepo

`func (o *OpaProfileListDto) GetAllowedRepo() []string`

GetAllowedRepo returns the AllowedRepo field if non-nil, zero value otherwise.

### GetAllowedRepoOk

`func (o *OpaProfileListDto) GetAllowedRepoOk() (*[]string, bool)`

GetAllowedRepoOk returns a tuple with the AllowedRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRepo

`func (o *OpaProfileListDto) SetAllowedRepo(v []string)`

SetAllowedRepo sets AllowedRepo field to given value.


### GetForbidSpecificTags

`func (o *OpaProfileListDto) GetForbidSpecificTags() []string`

GetForbidSpecificTags returns the ForbidSpecificTags field if non-nil, zero value otherwise.

### GetForbidSpecificTagsOk

`func (o *OpaProfileListDto) GetForbidSpecificTagsOk() (*[]string, bool)`

GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidSpecificTags

`func (o *OpaProfileListDto) SetForbidSpecificTags(v []string)`

SetForbidSpecificTags sets ForbidSpecificTags field to given value.


### GetIngressWhitelist

`func (o *OpaProfileListDto) GetIngressWhitelist() []string`

GetIngressWhitelist returns the IngressWhitelist field if non-nil, zero value otherwise.

### GetIngressWhitelistOk

`func (o *OpaProfileListDto) GetIngressWhitelistOk() (*[]string, bool)`

GetIngressWhitelistOk returns a tuple with the IngressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressWhitelist

`func (o *OpaProfileListDto) SetIngressWhitelist(v []string)`

SetIngressWhitelist sets IngressWhitelist field to given value.


### GetIsLocked

`func (o *OpaProfileListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OpaProfileListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OpaProfileListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetRevision

`func (o *OpaProfileListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OpaProfileListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OpaProfileListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetOrganizationId

`func (o *OpaProfileListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OpaProfileListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OpaProfileListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *OpaProfileListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OpaProfileListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OpaProfileListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCreatedAt

`func (o *OpaProfileListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpaProfileListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpaProfileListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *OpaProfileListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OpaProfileListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsDefault

`func (o *OpaProfileListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OpaProfileListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OpaProfileListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetProjects

`func (o *OpaProfileListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OpaProfileListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OpaProfileListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


