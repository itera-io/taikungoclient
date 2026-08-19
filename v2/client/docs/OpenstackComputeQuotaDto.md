# OpenstackComputeQuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTotalRamSize** | Pointer to **int64** |  | [optional] 
**MaxServerGroups** | Pointer to **int64** |  | [optional] 
**MaxTotalInstances** | Pointer to **int64** |  | [optional] 
**MaxTotalCores** | Pointer to **int64** |  | [optional] 
**UsedRamSize** | Pointer to **int64** |  | [optional] 
**UsedCpuSize** | Pointer to **int64** |  | [optional] 
**UsedInstanceSize** | Pointer to **int64** |  | [optional] 
**UsedServerGroups** | Pointer to **int64** |  | [optional] 

## Methods

### NewOpenstackComputeQuotaDto

`func NewOpenstackComputeQuotaDto() *OpenstackComputeQuotaDto`

NewOpenstackComputeQuotaDto instantiates a new OpenstackComputeQuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackComputeQuotaDtoWithDefaults

`func NewOpenstackComputeQuotaDtoWithDefaults() *OpenstackComputeQuotaDto`

NewOpenstackComputeQuotaDtoWithDefaults instantiates a new OpenstackComputeQuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTotalRamSize

`func (o *OpenstackComputeQuotaDto) GetMaxTotalRamSize() int64`

GetMaxTotalRamSize returns the MaxTotalRamSize field if non-nil, zero value otherwise.

### GetMaxTotalRamSizeOk

`func (o *OpenstackComputeQuotaDto) GetMaxTotalRamSizeOk() (*int64, bool)`

GetMaxTotalRamSizeOk returns a tuple with the MaxTotalRamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalRamSize

`func (o *OpenstackComputeQuotaDto) SetMaxTotalRamSize(v int64)`

SetMaxTotalRamSize sets MaxTotalRamSize field to given value.

### HasMaxTotalRamSize

`func (o *OpenstackComputeQuotaDto) HasMaxTotalRamSize() bool`

HasMaxTotalRamSize returns a boolean if a field has been set.

### GetMaxServerGroups

`func (o *OpenstackComputeQuotaDto) GetMaxServerGroups() int64`

GetMaxServerGroups returns the MaxServerGroups field if non-nil, zero value otherwise.

### GetMaxServerGroupsOk

`func (o *OpenstackComputeQuotaDto) GetMaxServerGroupsOk() (*int64, bool)`

GetMaxServerGroupsOk returns a tuple with the MaxServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerGroups

`func (o *OpenstackComputeQuotaDto) SetMaxServerGroups(v int64)`

SetMaxServerGroups sets MaxServerGroups field to given value.

### HasMaxServerGroups

`func (o *OpenstackComputeQuotaDto) HasMaxServerGroups() bool`

HasMaxServerGroups returns a boolean if a field has been set.

### GetMaxTotalInstances

`func (o *OpenstackComputeQuotaDto) GetMaxTotalInstances() int64`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *OpenstackComputeQuotaDto) GetMaxTotalInstancesOk() (*int64, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *OpenstackComputeQuotaDto) SetMaxTotalInstances(v int64)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *OpenstackComputeQuotaDto) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### GetMaxTotalCores

`func (o *OpenstackComputeQuotaDto) GetMaxTotalCores() int64`

GetMaxTotalCores returns the MaxTotalCores field if non-nil, zero value otherwise.

### GetMaxTotalCoresOk

`func (o *OpenstackComputeQuotaDto) GetMaxTotalCoresOk() (*int64, bool)`

GetMaxTotalCoresOk returns a tuple with the MaxTotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalCores

`func (o *OpenstackComputeQuotaDto) SetMaxTotalCores(v int64)`

SetMaxTotalCores sets MaxTotalCores field to given value.

### HasMaxTotalCores

`func (o *OpenstackComputeQuotaDto) HasMaxTotalCores() bool`

HasMaxTotalCores returns a boolean if a field has been set.

### GetUsedRamSize

`func (o *OpenstackComputeQuotaDto) GetUsedRamSize() int64`

GetUsedRamSize returns the UsedRamSize field if non-nil, zero value otherwise.

### GetUsedRamSizeOk

`func (o *OpenstackComputeQuotaDto) GetUsedRamSizeOk() (*int64, bool)`

GetUsedRamSizeOk returns a tuple with the UsedRamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRamSize

`func (o *OpenstackComputeQuotaDto) SetUsedRamSize(v int64)`

SetUsedRamSize sets UsedRamSize field to given value.

### HasUsedRamSize

`func (o *OpenstackComputeQuotaDto) HasUsedRamSize() bool`

HasUsedRamSize returns a boolean if a field has been set.

### GetUsedCpuSize

`func (o *OpenstackComputeQuotaDto) GetUsedCpuSize() int64`

GetUsedCpuSize returns the UsedCpuSize field if non-nil, zero value otherwise.

### GetUsedCpuSizeOk

`func (o *OpenstackComputeQuotaDto) GetUsedCpuSizeOk() (*int64, bool)`

GetUsedCpuSizeOk returns a tuple with the UsedCpuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCpuSize

`func (o *OpenstackComputeQuotaDto) SetUsedCpuSize(v int64)`

SetUsedCpuSize sets UsedCpuSize field to given value.

### HasUsedCpuSize

`func (o *OpenstackComputeQuotaDto) HasUsedCpuSize() bool`

HasUsedCpuSize returns a boolean if a field has been set.

### GetUsedInstanceSize

`func (o *OpenstackComputeQuotaDto) GetUsedInstanceSize() int64`

GetUsedInstanceSize returns the UsedInstanceSize field if non-nil, zero value otherwise.

### GetUsedInstanceSizeOk

`func (o *OpenstackComputeQuotaDto) GetUsedInstanceSizeOk() (*int64, bool)`

GetUsedInstanceSizeOk returns a tuple with the UsedInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInstanceSize

`func (o *OpenstackComputeQuotaDto) SetUsedInstanceSize(v int64)`

SetUsedInstanceSize sets UsedInstanceSize field to given value.

### HasUsedInstanceSize

`func (o *OpenstackComputeQuotaDto) HasUsedInstanceSize() bool`

HasUsedInstanceSize returns a boolean if a field has been set.

### GetUsedServerGroups

`func (o *OpenstackComputeQuotaDto) GetUsedServerGroups() int64`

GetUsedServerGroups returns the UsedServerGroups field if non-nil, zero value otherwise.

### GetUsedServerGroupsOk

`func (o *OpenstackComputeQuotaDto) GetUsedServerGroupsOk() (*int64, bool)`

GetUsedServerGroupsOk returns a tuple with the UsedServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedServerGroups

`func (o *OpenstackComputeQuotaDto) SetUsedServerGroups(v int64)`

SetUsedServerGroups sets UsedServerGroups field to given value.

### HasUsedServerGroups

`func (o *OpenstackComputeQuotaDto) HasUsedServerGroups() bool`

HasUsedServerGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


