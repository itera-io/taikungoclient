# CreateKubernetesProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**OctaviaEnabled** | **bool** |  | 
**ExposeNodePortOnBastion** | **bool** |  | 
**OrganizationId** | **NullableInt32** |  | 
**TaikunLBEnabled** | **bool** |  | 
**AllowSchedulingOnMaster** | **bool** |  | 
**UniqueClusterName** | **bool** |  | 
**NvidiaGpuOperatorEnabled** | **bool** |  | 
**WasmEnabled** | **bool** |  | 

## Methods

### NewCreateKubernetesProfileCommand

`func NewCreateKubernetesProfileCommand(name NullableString, octaviaEnabled bool, exposeNodePortOnBastion bool, organizationId NullableInt32, taikunLBEnabled bool, allowSchedulingOnMaster bool, uniqueClusterName bool, nvidiaGpuOperatorEnabled bool, wasmEnabled bool, ) *CreateKubernetesProfileCommand`

NewCreateKubernetesProfileCommand instantiates a new CreateKubernetesProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKubernetesProfileCommandWithDefaults

`func NewCreateKubernetesProfileCommandWithDefaults() *CreateKubernetesProfileCommand`

NewCreateKubernetesProfileCommandWithDefaults instantiates a new CreateKubernetesProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKubernetesProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKubernetesProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKubernetesProfileCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateKubernetesProfileCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateKubernetesProfileCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOctaviaEnabled

`func (o *CreateKubernetesProfileCommand) GetOctaviaEnabled() bool`

GetOctaviaEnabled returns the OctaviaEnabled field if non-nil, zero value otherwise.

### GetOctaviaEnabledOk

`func (o *CreateKubernetesProfileCommand) GetOctaviaEnabledOk() (*bool, bool)`

GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctaviaEnabled

`func (o *CreateKubernetesProfileCommand) SetOctaviaEnabled(v bool)`

SetOctaviaEnabled sets OctaviaEnabled field to given value.


### GetExposeNodePortOnBastion

`func (o *CreateKubernetesProfileCommand) GetExposeNodePortOnBastion() bool`

GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field if non-nil, zero value otherwise.

### GetExposeNodePortOnBastionOk

`func (o *CreateKubernetesProfileCommand) GetExposeNodePortOnBastionOk() (*bool, bool)`

GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeNodePortOnBastion

`func (o *CreateKubernetesProfileCommand) SetExposeNodePortOnBastion(v bool)`

SetExposeNodePortOnBastion sets ExposeNodePortOnBastion field to given value.


### GetOrganizationId

`func (o *CreateKubernetesProfileCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateKubernetesProfileCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateKubernetesProfileCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CreateKubernetesProfileCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateKubernetesProfileCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetTaikunLBEnabled

`func (o *CreateKubernetesProfileCommand) GetTaikunLBEnabled() bool`

GetTaikunLBEnabled returns the TaikunLBEnabled field if non-nil, zero value otherwise.

### GetTaikunLBEnabledOk

`func (o *CreateKubernetesProfileCommand) GetTaikunLBEnabledOk() (*bool, bool)`

GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBEnabled

`func (o *CreateKubernetesProfileCommand) SetTaikunLBEnabled(v bool)`

SetTaikunLBEnabled sets TaikunLBEnabled field to given value.


### GetAllowSchedulingOnMaster

`func (o *CreateKubernetesProfileCommand) GetAllowSchedulingOnMaster() bool`

GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field if non-nil, zero value otherwise.

### GetAllowSchedulingOnMasterOk

`func (o *CreateKubernetesProfileCommand) GetAllowSchedulingOnMasterOk() (*bool, bool)`

GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSchedulingOnMaster

`func (o *CreateKubernetesProfileCommand) SetAllowSchedulingOnMaster(v bool)`

SetAllowSchedulingOnMaster sets AllowSchedulingOnMaster field to given value.


### GetUniqueClusterName

`func (o *CreateKubernetesProfileCommand) GetUniqueClusterName() bool`

GetUniqueClusterName returns the UniqueClusterName field if non-nil, zero value otherwise.

### GetUniqueClusterNameOk

`func (o *CreateKubernetesProfileCommand) GetUniqueClusterNameOk() (*bool, bool)`

GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClusterName

`func (o *CreateKubernetesProfileCommand) SetUniqueClusterName(v bool)`

SetUniqueClusterName sets UniqueClusterName field to given value.


### GetNvidiaGpuOperatorEnabled

`func (o *CreateKubernetesProfileCommand) GetNvidiaGpuOperatorEnabled() bool`

GetNvidiaGpuOperatorEnabled returns the NvidiaGpuOperatorEnabled field if non-nil, zero value otherwise.

### GetNvidiaGpuOperatorEnabledOk

`func (o *CreateKubernetesProfileCommand) GetNvidiaGpuOperatorEnabledOk() (*bool, bool)`

GetNvidiaGpuOperatorEnabledOk returns a tuple with the NvidiaGpuOperatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpuOperatorEnabled

`func (o *CreateKubernetesProfileCommand) SetNvidiaGpuOperatorEnabled(v bool)`

SetNvidiaGpuOperatorEnabled sets NvidiaGpuOperatorEnabled field to given value.


### GetWasmEnabled

`func (o *CreateKubernetesProfileCommand) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *CreateKubernetesProfileCommand) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *CreateKubernetesProfileCommand) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


