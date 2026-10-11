# ServerForCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to [**CloudRole**](CloudRole.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **int64** |  | [optional] 
**Flavor** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**SpotPrice** | Pointer to **NullableFloat64** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**WasmEnabled** | Pointer to **bool** |  | [optional] 
**AutoscalingGroup** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**Hypervisor** | Pointer to **NullableString** |  | [optional] 
**KubernetesNodeLabels** | Pointer to [**[]KubernetesNodeLabelsDto**](KubernetesNodeLabelsDto.md) |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**ReplicaCount** | Pointer to **NullableInt32** |  | [optional] 
**UseLocalDisk** | Pointer to **bool** |  | [optional] 
**CskVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerForCreateDto

`func NewServerForCreateDto() *ServerForCreateDto`

NewServerForCreateDto instantiates a new ServerForCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerForCreateDtoWithDefaults

`func NewServerForCreateDtoWithDefaults() *ServerForCreateDto`

NewServerForCreateDtoWithDefaults instantiates a new ServerForCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerForCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerForCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerForCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerForCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServerForCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServerForCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *ServerForCreateDto) GetRole() CloudRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServerForCreateDto) GetRoleOk() (*CloudRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServerForCreateDto) SetRole(v CloudRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServerForCreateDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetProjectId

`func (o *ServerForCreateDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServerForCreateDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServerForCreateDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServerForCreateDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDiskSize

`func (o *ServerForCreateDto) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerForCreateDto) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerForCreateDto) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ServerForCreateDto) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetFlavor

`func (o *ServerForCreateDto) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ServerForCreateDto) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ServerForCreateDto) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *ServerForCreateDto) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *ServerForCreateDto) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *ServerForCreateDto) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetCount

`func (o *ServerForCreateDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServerForCreateDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServerForCreateDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ServerForCreateDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSpotPrice

`func (o *ServerForCreateDto) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *ServerForCreateDto) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *ServerForCreateDto) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *ServerForCreateDto) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### SetSpotPriceNil

`func (o *ServerForCreateDto) SetSpotPriceNil(b bool)`

 SetSpotPriceNil sets the value for SpotPrice to be an explicit nil

### UnsetSpotPrice
`func (o *ServerForCreateDto) UnsetSpotPrice()`

UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
### GetSpotInstance

`func (o *ServerForCreateDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *ServerForCreateDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *ServerForCreateDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *ServerForCreateDto) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetWasmEnabled

`func (o *ServerForCreateDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *ServerForCreateDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *ServerForCreateDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.

### HasWasmEnabled

`func (o *ServerForCreateDto) HasWasmEnabled() bool`

HasWasmEnabled returns a boolean if a field has been set.

### GetAutoscalingGroup

`func (o *ServerForCreateDto) GetAutoscalingGroup() string`

GetAutoscalingGroup returns the AutoscalingGroup field if non-nil, zero value otherwise.

### GetAutoscalingGroupOk

`func (o *ServerForCreateDto) GetAutoscalingGroupOk() (*string, bool)`

GetAutoscalingGroupOk returns a tuple with the AutoscalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroup

`func (o *ServerForCreateDto) SetAutoscalingGroup(v string)`

SetAutoscalingGroup sets AutoscalingGroup field to given value.

### HasAutoscalingGroup

`func (o *ServerForCreateDto) HasAutoscalingGroup() bool`

HasAutoscalingGroup returns a boolean if a field has been set.

### SetAutoscalingGroupNil

`func (o *ServerForCreateDto) SetAutoscalingGroupNil(b bool)`

 SetAutoscalingGroupNil sets the value for AutoscalingGroup to be an explicit nil

### UnsetAutoscalingGroup
`func (o *ServerForCreateDto) UnsetAutoscalingGroup()`

UnsetAutoscalingGroup ensures that no value is present for AutoscalingGroup, not even an explicit nil
### GetAvailabilityZone

`func (o *ServerForCreateDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ServerForCreateDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ServerForCreateDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ServerForCreateDto) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *ServerForCreateDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *ServerForCreateDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetHypervisor

`func (o *ServerForCreateDto) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *ServerForCreateDto) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *ServerForCreateDto) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *ServerForCreateDto) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### SetHypervisorNil

`func (o *ServerForCreateDto) SetHypervisorNil(b bool)`

 SetHypervisorNil sets the value for Hypervisor to be an explicit nil

### UnsetHypervisor
`func (o *ServerForCreateDto) UnsetHypervisor()`

UnsetHypervisor ensures that no value is present for Hypervisor, not even an explicit nil
### GetKubernetesNodeLabels

`func (o *ServerForCreateDto) GetKubernetesNodeLabels() []KubernetesNodeLabelsDto`

GetKubernetesNodeLabels returns the KubernetesNodeLabels field if non-nil, zero value otherwise.

### GetKubernetesNodeLabelsOk

`func (o *ServerForCreateDto) GetKubernetesNodeLabelsOk() (*[]KubernetesNodeLabelsDto, bool)`

GetKubernetesNodeLabelsOk returns a tuple with the KubernetesNodeLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNodeLabels

`func (o *ServerForCreateDto) SetKubernetesNodeLabels(v []KubernetesNodeLabelsDto)`

SetKubernetesNodeLabels sets KubernetesNodeLabels field to given value.

### HasKubernetesNodeLabels

`func (o *ServerForCreateDto) HasKubernetesNodeLabels() bool`

HasKubernetesNodeLabels returns a boolean if a field has been set.

### SetKubernetesNodeLabelsNil

`func (o *ServerForCreateDto) SetKubernetesNodeLabelsNil(b bool)`

 SetKubernetesNodeLabelsNil sets the value for KubernetesNodeLabels to be an explicit nil

### UnsetKubernetesNodeLabels
`func (o *ServerForCreateDto) UnsetKubernetesNodeLabels()`

UnsetKubernetesNodeLabels ensures that no value is present for KubernetesNodeLabels, not even an explicit nil
### GetSubnetId

`func (o *ServerForCreateDto) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ServerForCreateDto) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ServerForCreateDto) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *ServerForCreateDto) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *ServerForCreateDto) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *ServerForCreateDto) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetReplicaCount

`func (o *ServerForCreateDto) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ServerForCreateDto) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ServerForCreateDto) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *ServerForCreateDto) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### SetReplicaCountNil

`func (o *ServerForCreateDto) SetReplicaCountNil(b bool)`

 SetReplicaCountNil sets the value for ReplicaCount to be an explicit nil

### UnsetReplicaCount
`func (o *ServerForCreateDto) UnsetReplicaCount()`

UnsetReplicaCount ensures that no value is present for ReplicaCount, not even an explicit nil
### GetUseLocalDisk

`func (o *ServerForCreateDto) GetUseLocalDisk() bool`

GetUseLocalDisk returns the UseLocalDisk field if non-nil, zero value otherwise.

### GetUseLocalDiskOk

`func (o *ServerForCreateDto) GetUseLocalDiskOk() (*bool, bool)`

GetUseLocalDiskOk returns a tuple with the UseLocalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalDisk

`func (o *ServerForCreateDto) SetUseLocalDisk(v bool)`

SetUseLocalDisk sets UseLocalDisk field to given value.

### HasUseLocalDisk

`func (o *ServerForCreateDto) HasUseLocalDisk() bool`

HasUseLocalDisk returns a boolean if a field has been set.

### GetCskVersion

`func (o *ServerForCreateDto) GetCskVersion() string`

GetCskVersion returns the CskVersion field if non-nil, zero value otherwise.

### GetCskVersionOk

`func (o *ServerForCreateDto) GetCskVersionOk() (*string, bool)`

GetCskVersionOk returns a tuple with the CskVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCskVersion

`func (o *ServerForCreateDto) SetCskVersion(v string)`

SetCskVersion sets CskVersion field to given value.

### HasCskVersion

`func (o *ServerForCreateDto) HasCskVersion() bool`

HasCskVersion returns a boolean if a field has been set.

### SetCskVersionNil

`func (o *ServerForCreateDto) SetCskVersionNil(b bool)`

 SetCskVersionNil sets the value for CskVersion to be an explicit nil

### UnsetCskVersion
`func (o *ServerForCreateDto) UnsetCskVersion()`

UnsetCskVersion ensures that no value is present for CskVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


