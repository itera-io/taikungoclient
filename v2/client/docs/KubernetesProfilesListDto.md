# KubernetesProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **string** |  | 
**Cni** | [**CNI**](CNI.md) |  | 
**OctaviaEnabled** | **bool** |  | 
**ExposeNodePortOnBastion** | **bool** |  | 
**IsLocked** | **bool** |  | 
**TaikunLBEnabled** | **bool** |  | 
**AllowSchedulingOnMaster** | **bool** |  | 
**UniqueClusterName** | **bool** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**NvidiaGpuOperatorEnabled** | **bool** |  | 
**WasmEnabled** | **bool** |  | 

## Methods

### NewKubernetesProfilesListDto

`func NewKubernetesProfilesListDto(id int32, name string, organizationId NullableInt32, organizationName string, cni CNI, octaviaEnabled bool, exposeNodePortOnBastion bool, isLocked bool, taikunLBEnabled bool, allowSchedulingOnMaster bool, uniqueClusterName bool, projects []CommonDropdownDto, createdBy AuditUserDto, createdAt NullableString, lastModified NullableString, lastModifiedBy AuditUserDto, nvidiaGpuOperatorEnabled bool, wasmEnabled bool, ) *KubernetesProfilesListDto`

NewKubernetesProfilesListDto instantiates a new KubernetesProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProfilesListDtoWithDefaults

`func NewKubernetesProfilesListDtoWithDefaults() *KubernetesProfilesListDto`

NewKubernetesProfilesListDtoWithDefaults instantiates a new KubernetesProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesProfilesListDto) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *KubernetesProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KubernetesProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KubernetesProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *KubernetesProfilesListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *KubernetesProfilesListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *KubernetesProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *KubernetesProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *KubernetesProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCni

`func (o *KubernetesProfilesListDto) GetCni() CNI`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *KubernetesProfilesListDto) GetCniOk() (*CNI, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *KubernetesProfilesListDto) SetCni(v CNI)`

SetCni sets Cni field to given value.


### GetOctaviaEnabled

`func (o *KubernetesProfilesListDto) GetOctaviaEnabled() bool`

GetOctaviaEnabled returns the OctaviaEnabled field if non-nil, zero value otherwise.

### GetOctaviaEnabledOk

`func (o *KubernetesProfilesListDto) GetOctaviaEnabledOk() (*bool, bool)`

GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctaviaEnabled

`func (o *KubernetesProfilesListDto) SetOctaviaEnabled(v bool)`

SetOctaviaEnabled sets OctaviaEnabled field to given value.


### GetExposeNodePortOnBastion

`func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastion() bool`

GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field if non-nil, zero value otherwise.

### GetExposeNodePortOnBastionOk

`func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastionOk() (*bool, bool)`

GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeNodePortOnBastion

`func (o *KubernetesProfilesListDto) SetExposeNodePortOnBastion(v bool)`

SetExposeNodePortOnBastion sets ExposeNodePortOnBastion field to given value.


### GetIsLocked

`func (o *KubernetesProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *KubernetesProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *KubernetesProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetTaikunLBEnabled

`func (o *KubernetesProfilesListDto) GetTaikunLBEnabled() bool`

GetTaikunLBEnabled returns the TaikunLBEnabled field if non-nil, zero value otherwise.

### GetTaikunLBEnabledOk

`func (o *KubernetesProfilesListDto) GetTaikunLBEnabledOk() (*bool, bool)`

GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBEnabled

`func (o *KubernetesProfilesListDto) SetTaikunLBEnabled(v bool)`

SetTaikunLBEnabled sets TaikunLBEnabled field to given value.


### GetAllowSchedulingOnMaster

`func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMaster() bool`

GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field if non-nil, zero value otherwise.

### GetAllowSchedulingOnMasterOk

`func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMasterOk() (*bool, bool)`

GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSchedulingOnMaster

`func (o *KubernetesProfilesListDto) SetAllowSchedulingOnMaster(v bool)`

SetAllowSchedulingOnMaster sets AllowSchedulingOnMaster field to given value.


### GetUniqueClusterName

`func (o *KubernetesProfilesListDto) GetUniqueClusterName() bool`

GetUniqueClusterName returns the UniqueClusterName field if non-nil, zero value otherwise.

### GetUniqueClusterNameOk

`func (o *KubernetesProfilesListDto) GetUniqueClusterNameOk() (*bool, bool)`

GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClusterName

`func (o *KubernetesProfilesListDto) SetUniqueClusterName(v bool)`

SetUniqueClusterName sets UniqueClusterName field to given value.


### GetProjects

`func (o *KubernetesProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *KubernetesProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *KubernetesProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *KubernetesProfilesListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KubernetesProfilesListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KubernetesProfilesListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *KubernetesProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *KubernetesProfilesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *KubernetesProfilesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *KubernetesProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *KubernetesProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *KubernetesProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *KubernetesProfilesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *KubernetesProfilesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *KubernetesProfilesListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *KubernetesProfilesListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *KubernetesProfilesListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabled() bool`

GetNvidiaGpuOperatorEnabled returns the NvidiaGpuOperatorEnabled field if non-nil, zero value otherwise.

### GetNvidiaGpuOperatorEnabledOk

`func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabledOk() (*bool, bool)`

GetNvidiaGpuOperatorEnabledOk returns a tuple with the NvidiaGpuOperatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpuOperatorEnabled

`func (o *KubernetesProfilesListDto) SetNvidiaGpuOperatorEnabled(v bool)`

SetNvidiaGpuOperatorEnabled sets NvidiaGpuOperatorEnabled field to given value.


### GetWasmEnabled

`func (o *KubernetesProfilesListDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *KubernetesProfilesListDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *KubernetesProfilesListDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


