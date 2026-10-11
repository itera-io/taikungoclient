# UserResourceChartDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **int64** |  | [optional] 
**Ram** | Pointer to **int64** |  | [optional] 
**Cpu** | Pointer to **int64** |  | [optional] 
**MaxRam** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**MaxDiskSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserResourceChartDto

`func NewUserResourceChartDto() *UserResourceChartDto`

NewUserResourceChartDto instantiates a new UserResourceChartDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResourceChartDtoWithDefaults

`func NewUserResourceChartDtoWithDefaults() *UserResourceChartDto`

NewUserResourceChartDtoWithDefaults instantiates a new UserResourceChartDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *UserResourceChartDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *UserResourceChartDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *UserResourceChartDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *UserResourceChartDto) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *UserResourceChartDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *UserResourceChartDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetProjectId

`func (o *UserResourceChartDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UserResourceChartDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UserResourceChartDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UserResourceChartDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDiskSize

`func (o *UserResourceChartDto) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *UserResourceChartDto) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *UserResourceChartDto) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *UserResourceChartDto) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetRam

`func (o *UserResourceChartDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *UserResourceChartDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *UserResourceChartDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *UserResourceChartDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCpu

`func (o *UserResourceChartDto) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UserResourceChartDto) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UserResourceChartDto) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UserResourceChartDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMaxRam

`func (o *UserResourceChartDto) GetMaxRam() int64`

GetMaxRam returns the MaxRam field if non-nil, zero value otherwise.

### GetMaxRamOk

`func (o *UserResourceChartDto) GetMaxRamOk() (*int64, bool)`

GetMaxRamOk returns a tuple with the MaxRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRam

`func (o *UserResourceChartDto) SetMaxRam(v int64)`

SetMaxRam sets MaxRam field to given value.

### HasMaxRam

`func (o *UserResourceChartDto) HasMaxRam() bool`

HasMaxRam returns a boolean if a field has been set.

### GetMaxCpu

`func (o *UserResourceChartDto) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *UserResourceChartDto) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *UserResourceChartDto) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *UserResourceChartDto) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxDiskSize

`func (o *UserResourceChartDto) GetMaxDiskSize() int64`

GetMaxDiskSize returns the MaxDiskSize field if non-nil, zero value otherwise.

### GetMaxDiskSizeOk

`func (o *UserResourceChartDto) GetMaxDiskSizeOk() (*int64, bool)`

GetMaxDiskSizeOk returns a tuple with the MaxDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiskSize

`func (o *UserResourceChartDto) SetMaxDiskSize(v int64)`

SetMaxDiskSize sets MaxDiskSize field to given value.

### HasMaxDiskSize

`func (o *UserResourceChartDto) HasMaxDiskSize() bool`

HasMaxDiskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


