# ResourcePresetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Cpu** | [**CpuLimitsDto**](CpuLimitsDto.md) |  | 
**Ram** | [**RamLimitsDto**](RamLimitsDto.md) |  | 
**EphemeralStorage** | [**EphemeralStorageLimitsDto**](EphemeralStorageLimitsDto.md) |  | 
**WorkloadResources** | [**WorkloadResourceLimitsDto**](WorkloadResourceLimitsDto.md) |  | 

## Methods

### NewResourcePresetDto

`func NewResourcePresetDto(name NullableString, cpu CpuLimitsDto, ram RamLimitsDto, ephemeralStorage EphemeralStorageLimitsDto, workloadResources WorkloadResourceLimitsDto, ) *ResourcePresetDto`

NewResourcePresetDto instantiates a new ResourcePresetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePresetDtoWithDefaults

`func NewResourcePresetDtoWithDefaults() *ResourcePresetDto`

NewResourcePresetDtoWithDefaults instantiates a new ResourcePresetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourcePresetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePresetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePresetDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ResourcePresetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ResourcePresetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *ResourcePresetDto) GetCpu() CpuLimitsDto`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourcePresetDto) GetCpuOk() (*CpuLimitsDto, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourcePresetDto) SetCpu(v CpuLimitsDto)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *ResourcePresetDto) GetRam() RamLimitsDto`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ResourcePresetDto) GetRamOk() (*RamLimitsDto, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ResourcePresetDto) SetRam(v RamLimitsDto)`

SetRam sets Ram field to given value.


### GetEphemeralStorage

`func (o *ResourcePresetDto) GetEphemeralStorage() EphemeralStorageLimitsDto`

GetEphemeralStorage returns the EphemeralStorage field if non-nil, zero value otherwise.

### GetEphemeralStorageOk

`func (o *ResourcePresetDto) GetEphemeralStorageOk() (*EphemeralStorageLimitsDto, bool)`

GetEphemeralStorageOk returns a tuple with the EphemeralStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorage

`func (o *ResourcePresetDto) SetEphemeralStorage(v EphemeralStorageLimitsDto)`

SetEphemeralStorage sets EphemeralStorage field to given value.


### GetWorkloadResources

`func (o *ResourcePresetDto) GetWorkloadResources() WorkloadResourceLimitsDto`

GetWorkloadResources returns the WorkloadResources field if non-nil, zero value otherwise.

### GetWorkloadResourcesOk

`func (o *ResourcePresetDto) GetWorkloadResourcesOk() (*WorkloadResourceLimitsDto, bool)`

GetWorkloadResourcesOk returns a tuple with the WorkloadResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadResources

`func (o *ResourcePresetDto) SetWorkloadResources(v WorkloadResourceLimitsDto)`

SetWorkloadResources sets WorkloadResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


