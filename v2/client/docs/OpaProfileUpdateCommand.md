# OpaProfileUpdateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ForbidNodePort** | Pointer to **NullableBool** |  | [optional] 
**ForbidHttpIngress** | Pointer to **NullableBool** |  | [optional] 
**RequireProbe** | Pointer to **NullableBool** |  | [optional] 
**UniqueIngresses** | Pointer to **NullableBool** |  | [optional] 
**UniqueServiceSelector** | Pointer to **NullableBool** |  | [optional] 
**IsNodeNameForbiddenInVC** | Pointer to **NullableBool** |  | [optional] 
**IsMasterTaintEnforced** | Pointer to **NullableBool** |  | [optional] 
**ForcePodResource** | Pointer to **NullableBool** |  | [optional] 
**AllowedRepo** | Pointer to **[]string** |  | [optional] 
**ForbidSpecificTags** | Pointer to **[]string** |  | [optional] 
**IngressWhitelist** | Pointer to **[]string** |  | [optional] 
**WhitelistMasterTaintNamespaces** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOpaProfileUpdateCommand

`func NewOpaProfileUpdateCommand() *OpaProfileUpdateCommand`

NewOpaProfileUpdateCommand instantiates a new OpaProfileUpdateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaProfileUpdateCommandWithDefaults

`func NewOpaProfileUpdateCommandWithDefaults() *OpaProfileUpdateCommand`

NewOpaProfileUpdateCommandWithDefaults instantiates a new OpaProfileUpdateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpaProfileUpdateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaProfileUpdateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaProfileUpdateCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpaProfileUpdateCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OpaProfileUpdateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpaProfileUpdateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpaProfileUpdateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpaProfileUpdateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpaProfileUpdateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpaProfileUpdateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForbidNodePort

`func (o *OpaProfileUpdateCommand) GetForbidNodePort() bool`

GetForbidNodePort returns the ForbidNodePort field if non-nil, zero value otherwise.

### GetForbidNodePortOk

`func (o *OpaProfileUpdateCommand) GetForbidNodePortOk() (*bool, bool)`

GetForbidNodePortOk returns a tuple with the ForbidNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidNodePort

`func (o *OpaProfileUpdateCommand) SetForbidNodePort(v bool)`

SetForbidNodePort sets ForbidNodePort field to given value.

### HasForbidNodePort

`func (o *OpaProfileUpdateCommand) HasForbidNodePort() bool`

HasForbidNodePort returns a boolean if a field has been set.

### SetForbidNodePortNil

`func (o *OpaProfileUpdateCommand) SetForbidNodePortNil(b bool)`

 SetForbidNodePortNil sets the value for ForbidNodePort to be an explicit nil

### UnsetForbidNodePort
`func (o *OpaProfileUpdateCommand) UnsetForbidNodePort()`

UnsetForbidNodePort ensures that no value is present for ForbidNodePort, not even an explicit nil
### GetForbidHttpIngress

`func (o *OpaProfileUpdateCommand) GetForbidHttpIngress() bool`

GetForbidHttpIngress returns the ForbidHttpIngress field if non-nil, zero value otherwise.

### GetForbidHttpIngressOk

`func (o *OpaProfileUpdateCommand) GetForbidHttpIngressOk() (*bool, bool)`

GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidHttpIngress

`func (o *OpaProfileUpdateCommand) SetForbidHttpIngress(v bool)`

SetForbidHttpIngress sets ForbidHttpIngress field to given value.

### HasForbidHttpIngress

`func (o *OpaProfileUpdateCommand) HasForbidHttpIngress() bool`

HasForbidHttpIngress returns a boolean if a field has been set.

### SetForbidHttpIngressNil

`func (o *OpaProfileUpdateCommand) SetForbidHttpIngressNil(b bool)`

 SetForbidHttpIngressNil sets the value for ForbidHttpIngress to be an explicit nil

### UnsetForbidHttpIngress
`func (o *OpaProfileUpdateCommand) UnsetForbidHttpIngress()`

UnsetForbidHttpIngress ensures that no value is present for ForbidHttpIngress, not even an explicit nil
### GetRequireProbe

`func (o *OpaProfileUpdateCommand) GetRequireProbe() bool`

GetRequireProbe returns the RequireProbe field if non-nil, zero value otherwise.

### GetRequireProbeOk

`func (o *OpaProfileUpdateCommand) GetRequireProbeOk() (*bool, bool)`

GetRequireProbeOk returns a tuple with the RequireProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProbe

`func (o *OpaProfileUpdateCommand) SetRequireProbe(v bool)`

SetRequireProbe sets RequireProbe field to given value.

### HasRequireProbe

`func (o *OpaProfileUpdateCommand) HasRequireProbe() bool`

HasRequireProbe returns a boolean if a field has been set.

### SetRequireProbeNil

`func (o *OpaProfileUpdateCommand) SetRequireProbeNil(b bool)`

 SetRequireProbeNil sets the value for RequireProbe to be an explicit nil

### UnsetRequireProbe
`func (o *OpaProfileUpdateCommand) UnsetRequireProbe()`

UnsetRequireProbe ensures that no value is present for RequireProbe, not even an explicit nil
### GetUniqueIngresses

`func (o *OpaProfileUpdateCommand) GetUniqueIngresses() bool`

GetUniqueIngresses returns the UniqueIngresses field if non-nil, zero value otherwise.

### GetUniqueIngressesOk

`func (o *OpaProfileUpdateCommand) GetUniqueIngressesOk() (*bool, bool)`

GetUniqueIngressesOk returns a tuple with the UniqueIngresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIngresses

`func (o *OpaProfileUpdateCommand) SetUniqueIngresses(v bool)`

SetUniqueIngresses sets UniqueIngresses field to given value.

### HasUniqueIngresses

`func (o *OpaProfileUpdateCommand) HasUniqueIngresses() bool`

HasUniqueIngresses returns a boolean if a field has been set.

### SetUniqueIngressesNil

`func (o *OpaProfileUpdateCommand) SetUniqueIngressesNil(b bool)`

 SetUniqueIngressesNil sets the value for UniqueIngresses to be an explicit nil

### UnsetUniqueIngresses
`func (o *OpaProfileUpdateCommand) UnsetUniqueIngresses()`

UnsetUniqueIngresses ensures that no value is present for UniqueIngresses, not even an explicit nil
### GetUniqueServiceSelector

`func (o *OpaProfileUpdateCommand) GetUniqueServiceSelector() bool`

GetUniqueServiceSelector returns the UniqueServiceSelector field if non-nil, zero value otherwise.

### GetUniqueServiceSelectorOk

`func (o *OpaProfileUpdateCommand) GetUniqueServiceSelectorOk() (*bool, bool)`

GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueServiceSelector

`func (o *OpaProfileUpdateCommand) SetUniqueServiceSelector(v bool)`

SetUniqueServiceSelector sets UniqueServiceSelector field to given value.

### HasUniqueServiceSelector

`func (o *OpaProfileUpdateCommand) HasUniqueServiceSelector() bool`

HasUniqueServiceSelector returns a boolean if a field has been set.

### SetUniqueServiceSelectorNil

`func (o *OpaProfileUpdateCommand) SetUniqueServiceSelectorNil(b bool)`

 SetUniqueServiceSelectorNil sets the value for UniqueServiceSelector to be an explicit nil

### UnsetUniqueServiceSelector
`func (o *OpaProfileUpdateCommand) UnsetUniqueServiceSelector()`

UnsetUniqueServiceSelector ensures that no value is present for UniqueServiceSelector, not even an explicit nil
### GetIsNodeNameForbiddenInVC

`func (o *OpaProfileUpdateCommand) GetIsNodeNameForbiddenInVC() bool`

GetIsNodeNameForbiddenInVC returns the IsNodeNameForbiddenInVC field if non-nil, zero value otherwise.

### GetIsNodeNameForbiddenInVCOk

`func (o *OpaProfileUpdateCommand) GetIsNodeNameForbiddenInVCOk() (*bool, bool)`

GetIsNodeNameForbiddenInVCOk returns a tuple with the IsNodeNameForbiddenInVC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNodeNameForbiddenInVC

`func (o *OpaProfileUpdateCommand) SetIsNodeNameForbiddenInVC(v bool)`

SetIsNodeNameForbiddenInVC sets IsNodeNameForbiddenInVC field to given value.

### HasIsNodeNameForbiddenInVC

`func (o *OpaProfileUpdateCommand) HasIsNodeNameForbiddenInVC() bool`

HasIsNodeNameForbiddenInVC returns a boolean if a field has been set.

### SetIsNodeNameForbiddenInVCNil

`func (o *OpaProfileUpdateCommand) SetIsNodeNameForbiddenInVCNil(b bool)`

 SetIsNodeNameForbiddenInVCNil sets the value for IsNodeNameForbiddenInVC to be an explicit nil

### UnsetIsNodeNameForbiddenInVC
`func (o *OpaProfileUpdateCommand) UnsetIsNodeNameForbiddenInVC()`

UnsetIsNodeNameForbiddenInVC ensures that no value is present for IsNodeNameForbiddenInVC, not even an explicit nil
### GetIsMasterTaintEnforced

`func (o *OpaProfileUpdateCommand) GetIsMasterTaintEnforced() bool`

GetIsMasterTaintEnforced returns the IsMasterTaintEnforced field if non-nil, zero value otherwise.

### GetIsMasterTaintEnforcedOk

`func (o *OpaProfileUpdateCommand) GetIsMasterTaintEnforcedOk() (*bool, bool)`

GetIsMasterTaintEnforcedOk returns a tuple with the IsMasterTaintEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterTaintEnforced

`func (o *OpaProfileUpdateCommand) SetIsMasterTaintEnforced(v bool)`

SetIsMasterTaintEnforced sets IsMasterTaintEnforced field to given value.

### HasIsMasterTaintEnforced

`func (o *OpaProfileUpdateCommand) HasIsMasterTaintEnforced() bool`

HasIsMasterTaintEnforced returns a boolean if a field has been set.

### SetIsMasterTaintEnforcedNil

`func (o *OpaProfileUpdateCommand) SetIsMasterTaintEnforcedNil(b bool)`

 SetIsMasterTaintEnforcedNil sets the value for IsMasterTaintEnforced to be an explicit nil

### UnsetIsMasterTaintEnforced
`func (o *OpaProfileUpdateCommand) UnsetIsMasterTaintEnforced()`

UnsetIsMasterTaintEnforced ensures that no value is present for IsMasterTaintEnforced, not even an explicit nil
### GetForcePodResource

`func (o *OpaProfileUpdateCommand) GetForcePodResource() bool`

GetForcePodResource returns the ForcePodResource field if non-nil, zero value otherwise.

### GetForcePodResourceOk

`func (o *OpaProfileUpdateCommand) GetForcePodResourceOk() (*bool, bool)`

GetForcePodResourceOk returns a tuple with the ForcePodResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePodResource

`func (o *OpaProfileUpdateCommand) SetForcePodResource(v bool)`

SetForcePodResource sets ForcePodResource field to given value.

### HasForcePodResource

`func (o *OpaProfileUpdateCommand) HasForcePodResource() bool`

HasForcePodResource returns a boolean if a field has been set.

### SetForcePodResourceNil

`func (o *OpaProfileUpdateCommand) SetForcePodResourceNil(b bool)`

 SetForcePodResourceNil sets the value for ForcePodResource to be an explicit nil

### UnsetForcePodResource
`func (o *OpaProfileUpdateCommand) UnsetForcePodResource()`

UnsetForcePodResource ensures that no value is present for ForcePodResource, not even an explicit nil
### GetAllowedRepo

`func (o *OpaProfileUpdateCommand) GetAllowedRepo() []string`

GetAllowedRepo returns the AllowedRepo field if non-nil, zero value otherwise.

### GetAllowedRepoOk

`func (o *OpaProfileUpdateCommand) GetAllowedRepoOk() (*[]string, bool)`

GetAllowedRepoOk returns a tuple with the AllowedRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRepo

`func (o *OpaProfileUpdateCommand) SetAllowedRepo(v []string)`

SetAllowedRepo sets AllowedRepo field to given value.

### HasAllowedRepo

`func (o *OpaProfileUpdateCommand) HasAllowedRepo() bool`

HasAllowedRepo returns a boolean if a field has been set.

### SetAllowedRepoNil

`func (o *OpaProfileUpdateCommand) SetAllowedRepoNil(b bool)`

 SetAllowedRepoNil sets the value for AllowedRepo to be an explicit nil

### UnsetAllowedRepo
`func (o *OpaProfileUpdateCommand) UnsetAllowedRepo()`

UnsetAllowedRepo ensures that no value is present for AllowedRepo, not even an explicit nil
### GetForbidSpecificTags

`func (o *OpaProfileUpdateCommand) GetForbidSpecificTags() []string`

GetForbidSpecificTags returns the ForbidSpecificTags field if non-nil, zero value otherwise.

### GetForbidSpecificTagsOk

`func (o *OpaProfileUpdateCommand) GetForbidSpecificTagsOk() (*[]string, bool)`

GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidSpecificTags

`func (o *OpaProfileUpdateCommand) SetForbidSpecificTags(v []string)`

SetForbidSpecificTags sets ForbidSpecificTags field to given value.

### HasForbidSpecificTags

`func (o *OpaProfileUpdateCommand) HasForbidSpecificTags() bool`

HasForbidSpecificTags returns a boolean if a field has been set.

### SetForbidSpecificTagsNil

`func (o *OpaProfileUpdateCommand) SetForbidSpecificTagsNil(b bool)`

 SetForbidSpecificTagsNil sets the value for ForbidSpecificTags to be an explicit nil

### UnsetForbidSpecificTags
`func (o *OpaProfileUpdateCommand) UnsetForbidSpecificTags()`

UnsetForbidSpecificTags ensures that no value is present for ForbidSpecificTags, not even an explicit nil
### GetIngressWhitelist

`func (o *OpaProfileUpdateCommand) GetIngressWhitelist() []string`

GetIngressWhitelist returns the IngressWhitelist field if non-nil, zero value otherwise.

### GetIngressWhitelistOk

`func (o *OpaProfileUpdateCommand) GetIngressWhitelistOk() (*[]string, bool)`

GetIngressWhitelistOk returns a tuple with the IngressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressWhitelist

`func (o *OpaProfileUpdateCommand) SetIngressWhitelist(v []string)`

SetIngressWhitelist sets IngressWhitelist field to given value.

### HasIngressWhitelist

`func (o *OpaProfileUpdateCommand) HasIngressWhitelist() bool`

HasIngressWhitelist returns a boolean if a field has been set.

### SetIngressWhitelistNil

`func (o *OpaProfileUpdateCommand) SetIngressWhitelistNil(b bool)`

 SetIngressWhitelistNil sets the value for IngressWhitelist to be an explicit nil

### UnsetIngressWhitelist
`func (o *OpaProfileUpdateCommand) UnsetIngressWhitelist()`

UnsetIngressWhitelist ensures that no value is present for IngressWhitelist, not even an explicit nil
### GetWhitelistMasterTaintNamespaces

`func (o *OpaProfileUpdateCommand) GetWhitelistMasterTaintNamespaces() []string`

GetWhitelistMasterTaintNamespaces returns the WhitelistMasterTaintNamespaces field if non-nil, zero value otherwise.

### GetWhitelistMasterTaintNamespacesOk

`func (o *OpaProfileUpdateCommand) GetWhitelistMasterTaintNamespacesOk() (*[]string, bool)`

GetWhitelistMasterTaintNamespacesOk returns a tuple with the WhitelistMasterTaintNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistMasterTaintNamespaces

`func (o *OpaProfileUpdateCommand) SetWhitelistMasterTaintNamespaces(v []string)`

SetWhitelistMasterTaintNamespaces sets WhitelistMasterTaintNamespaces field to given value.

### HasWhitelistMasterTaintNamespaces

`func (o *OpaProfileUpdateCommand) HasWhitelistMasterTaintNamespaces() bool`

HasWhitelistMasterTaintNamespaces returns a boolean if a field has been set.

### SetWhitelistMasterTaintNamespacesNil

`func (o *OpaProfileUpdateCommand) SetWhitelistMasterTaintNamespacesNil(b bool)`

 SetWhitelistMasterTaintNamespacesNil sets the value for WhitelistMasterTaintNamespaces to be an explicit nil

### UnsetWhitelistMasterTaintNamespaces
`func (o *OpaProfileUpdateCommand) UnsetWhitelistMasterTaintNamespaces()`

UnsetWhitelistMasterTaintNamespaces ensures that no value is present for WhitelistMasterTaintNamespaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


