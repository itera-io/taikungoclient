# UpdateQuotaCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaId** | Pointer to **int32** |  | [optional] 
**ServerCpu** | Pointer to **int64** |  | [optional] 
**ServerRam** | Pointer to **float64** |  | [optional] 
**ServerDiskSize** | Pointer to **float64** |  | [optional] 
**VmCpu** | Pointer to **int64** |  | [optional] 
**VmRam** | Pointer to **float64** |  | [optional] 
**VmVolumeSize** | Pointer to **float64** |  | [optional] 

## Methods

### NewUpdateQuotaCommand

`func NewUpdateQuotaCommand() *UpdateQuotaCommand`

NewUpdateQuotaCommand instantiates a new UpdateQuotaCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQuotaCommandWithDefaults

`func NewUpdateQuotaCommandWithDefaults() *UpdateQuotaCommand`

NewUpdateQuotaCommandWithDefaults instantiates a new UpdateQuotaCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaId

`func (o *UpdateQuotaCommand) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *UpdateQuotaCommand) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *UpdateQuotaCommand) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *UpdateQuotaCommand) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetServerCpu

`func (o *UpdateQuotaCommand) GetServerCpu() int64`

GetServerCpu returns the ServerCpu field if non-nil, zero value otherwise.

### GetServerCpuOk

`func (o *UpdateQuotaCommand) GetServerCpuOk() (*int64, bool)`

GetServerCpuOk returns a tuple with the ServerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCpu

`func (o *UpdateQuotaCommand) SetServerCpu(v int64)`

SetServerCpu sets ServerCpu field to given value.

### HasServerCpu

`func (o *UpdateQuotaCommand) HasServerCpu() bool`

HasServerCpu returns a boolean if a field has been set.

### GetServerRam

`func (o *UpdateQuotaCommand) GetServerRam() float64`

GetServerRam returns the ServerRam field if non-nil, zero value otherwise.

### GetServerRamOk

`func (o *UpdateQuotaCommand) GetServerRamOk() (*float64, bool)`

GetServerRamOk returns a tuple with the ServerRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRam

`func (o *UpdateQuotaCommand) SetServerRam(v float64)`

SetServerRam sets ServerRam field to given value.

### HasServerRam

`func (o *UpdateQuotaCommand) HasServerRam() bool`

HasServerRam returns a boolean if a field has been set.

### GetServerDiskSize

`func (o *UpdateQuotaCommand) GetServerDiskSize() float64`

GetServerDiskSize returns the ServerDiskSize field if non-nil, zero value otherwise.

### GetServerDiskSizeOk

`func (o *UpdateQuotaCommand) GetServerDiskSizeOk() (*float64, bool)`

GetServerDiskSizeOk returns a tuple with the ServerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDiskSize

`func (o *UpdateQuotaCommand) SetServerDiskSize(v float64)`

SetServerDiskSize sets ServerDiskSize field to given value.

### HasServerDiskSize

`func (o *UpdateQuotaCommand) HasServerDiskSize() bool`

HasServerDiskSize returns a boolean if a field has been set.

### GetVmCpu

`func (o *UpdateQuotaCommand) GetVmCpu() int64`

GetVmCpu returns the VmCpu field if non-nil, zero value otherwise.

### GetVmCpuOk

`func (o *UpdateQuotaCommand) GetVmCpuOk() (*int64, bool)`

GetVmCpuOk returns a tuple with the VmCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpu

`func (o *UpdateQuotaCommand) SetVmCpu(v int64)`

SetVmCpu sets VmCpu field to given value.

### HasVmCpu

`func (o *UpdateQuotaCommand) HasVmCpu() bool`

HasVmCpu returns a boolean if a field has been set.

### GetVmRam

`func (o *UpdateQuotaCommand) GetVmRam() float64`

GetVmRam returns the VmRam field if non-nil, zero value otherwise.

### GetVmRamOk

`func (o *UpdateQuotaCommand) GetVmRamOk() (*float64, bool)`

GetVmRamOk returns a tuple with the VmRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRam

`func (o *UpdateQuotaCommand) SetVmRam(v float64)`

SetVmRam sets VmRam field to given value.

### HasVmRam

`func (o *UpdateQuotaCommand) HasVmRam() bool`

HasVmRam returns a boolean if a field has been set.

### GetVmVolumeSize

`func (o *UpdateQuotaCommand) GetVmVolumeSize() float64`

GetVmVolumeSize returns the VmVolumeSize field if non-nil, zero value otherwise.

### GetVmVolumeSizeOk

`func (o *UpdateQuotaCommand) GetVmVolumeSizeOk() (*float64, bool)`

GetVmVolumeSizeOk returns a tuple with the VmVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVolumeSize

`func (o *UpdateQuotaCommand) SetVmVolumeSize(v float64)`

SetVmVolumeSize sets VmVolumeSize field to given value.

### HasVmVolumeSize

`func (o *UpdateQuotaCommand) HasVmVolumeSize() bool`

HasVmVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


