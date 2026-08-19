# EditVirtualClusterQuotasCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VClusterId** | **int32** |  | 
**WorkloadResources** | [**EditVirtualClusterWorkloadResources**](EditVirtualClusterWorkloadResources.md) |  | 
**CpuLimits** | [**EditVirtualClusterResourceLimits**](EditVirtualClusterResourceLimits.md) |  | 
**RamLimits** | [**EditVirtualClusterResourceLimits**](EditVirtualClusterResourceLimits.md) |  | 
**EphemeralStorageLimits** | [**EditVirtualClusterResourceLimits**](EditVirtualClusterResourceLimits.md) |  | 

## Methods

### NewEditVirtualClusterQuotasCommand

`func NewEditVirtualClusterQuotasCommand(vClusterId int32, workloadResources EditVirtualClusterWorkloadResources, cpuLimits EditVirtualClusterResourceLimits, ramLimits EditVirtualClusterResourceLimits, ephemeralStorageLimits EditVirtualClusterResourceLimits, ) *EditVirtualClusterQuotasCommand`

NewEditVirtualClusterQuotasCommand instantiates a new EditVirtualClusterQuotasCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditVirtualClusterQuotasCommandWithDefaults

`func NewEditVirtualClusterQuotasCommandWithDefaults() *EditVirtualClusterQuotasCommand`

NewEditVirtualClusterQuotasCommandWithDefaults instantiates a new EditVirtualClusterQuotasCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVClusterId

`func (o *EditVirtualClusterQuotasCommand) GetVClusterId() int32`

GetVClusterId returns the VClusterId field if non-nil, zero value otherwise.

### GetVClusterIdOk

`func (o *EditVirtualClusterQuotasCommand) GetVClusterIdOk() (*int32, bool)`

GetVClusterIdOk returns a tuple with the VClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVClusterId

`func (o *EditVirtualClusterQuotasCommand) SetVClusterId(v int32)`

SetVClusterId sets VClusterId field to given value.


### GetWorkloadResources

`func (o *EditVirtualClusterQuotasCommand) GetWorkloadResources() EditVirtualClusterWorkloadResources`

GetWorkloadResources returns the WorkloadResources field if non-nil, zero value otherwise.

### GetWorkloadResourcesOk

`func (o *EditVirtualClusterQuotasCommand) GetWorkloadResourcesOk() (*EditVirtualClusterWorkloadResources, bool)`

GetWorkloadResourcesOk returns a tuple with the WorkloadResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadResources

`func (o *EditVirtualClusterQuotasCommand) SetWorkloadResources(v EditVirtualClusterWorkloadResources)`

SetWorkloadResources sets WorkloadResources field to given value.


### GetCpuLimits

`func (o *EditVirtualClusterQuotasCommand) GetCpuLimits() EditVirtualClusterResourceLimits`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *EditVirtualClusterQuotasCommand) GetCpuLimitsOk() (*EditVirtualClusterResourceLimits, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *EditVirtualClusterQuotasCommand) SetCpuLimits(v EditVirtualClusterResourceLimits)`

SetCpuLimits sets CpuLimits field to given value.


### GetRamLimits

`func (o *EditVirtualClusterQuotasCommand) GetRamLimits() EditVirtualClusterResourceLimits`

GetRamLimits returns the RamLimits field if non-nil, zero value otherwise.

### GetRamLimitsOk

`func (o *EditVirtualClusterQuotasCommand) GetRamLimitsOk() (*EditVirtualClusterResourceLimits, bool)`

GetRamLimitsOk returns a tuple with the RamLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamLimits

`func (o *EditVirtualClusterQuotasCommand) SetRamLimits(v EditVirtualClusterResourceLimits)`

SetRamLimits sets RamLimits field to given value.


### GetEphemeralStorageLimits

`func (o *EditVirtualClusterQuotasCommand) GetEphemeralStorageLimits() EditVirtualClusterResourceLimits`

GetEphemeralStorageLimits returns the EphemeralStorageLimits field if non-nil, zero value otherwise.

### GetEphemeralStorageLimitsOk

`func (o *EditVirtualClusterQuotasCommand) GetEphemeralStorageLimitsOk() (*EditVirtualClusterResourceLimits, bool)`

GetEphemeralStorageLimitsOk returns a tuple with the EphemeralStorageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageLimits

`func (o *EditVirtualClusterQuotasCommand) SetEphemeralStorageLimits(v EditVirtualClusterResourceLimits)`

SetEphemeralStorageLimits sets EphemeralStorageLimits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


