# ServerChartDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Azure** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Openstack** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Google** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**GenericK8S** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Failed** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Succeeded** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Waiting** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Updating** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Deleting** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**Purging** | Pointer to [**[]ServerCommonRecordDto**](ServerCommonRecordDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**TotalCpu** | Pointer to **int32** |  | [optional] 
**TotalRam** | Pointer to **int64** |  | [optional] 
**TotalDiskSize** | Pointer to **int64** |  | [optional] 
**TotalFailedCount** | Pointer to **int32** |  | [optional] 
**TotalSucceededCount** | Pointer to **int32** |  | [optional] 
**TotalUpdatingCount** | Pointer to **int32** |  | [optional] 
**TotalPendingCount** | Pointer to **int32** |  | [optional] 
**TotalAwsCount** | Pointer to **int32** |  | [optional] 
**TotalAzureCount** | Pointer to **int32** |  | [optional] 
**TotalOpenstackCount** | Pointer to **int32** |  | [optional] 
**TotalGoogleCount** | Pointer to **int32** |  | [optional] 
**TotalGenericK8SCount** | Pointer to **int32** |  | [optional] 
**UsedResources** | Pointer to [**[]UserResourceChartDto**](UserResourceChartDto.md) |  | [optional] 

## Methods

### NewServerChartDto

`func NewServerChartDto() *ServerChartDto`

NewServerChartDto instantiates a new ServerChartDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerChartDtoWithDefaults

`func NewServerChartDtoWithDefaults() *ServerChartDto`

NewServerChartDtoWithDefaults instantiates a new ServerChartDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *ServerChartDto) GetAws() []ServerCommonRecordDto`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *ServerChartDto) GetAwsOk() (*[]ServerCommonRecordDto, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *ServerChartDto) SetAws(v []ServerCommonRecordDto)`

SetAws sets Aws field to given value.

### HasAws

`func (o *ServerChartDto) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *ServerChartDto) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *ServerChartDto) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetAzure

`func (o *ServerChartDto) GetAzure() []ServerCommonRecordDto`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *ServerChartDto) GetAzureOk() (*[]ServerCommonRecordDto, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *ServerChartDto) SetAzure(v []ServerCommonRecordDto)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *ServerChartDto) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *ServerChartDto) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *ServerChartDto) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil
### GetOpenstack

`func (o *ServerChartDto) GetOpenstack() []ServerCommonRecordDto`

GetOpenstack returns the Openstack field if non-nil, zero value otherwise.

### GetOpenstackOk

`func (o *ServerChartDto) GetOpenstackOk() (*[]ServerCommonRecordDto, bool)`

GetOpenstackOk returns a tuple with the Openstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstack

`func (o *ServerChartDto) SetOpenstack(v []ServerCommonRecordDto)`

SetOpenstack sets Openstack field to given value.

### HasOpenstack

`func (o *ServerChartDto) HasOpenstack() bool`

HasOpenstack returns a boolean if a field has been set.

### SetOpenstackNil

`func (o *ServerChartDto) SetOpenstackNil(b bool)`

 SetOpenstackNil sets the value for Openstack to be an explicit nil

### UnsetOpenstack
`func (o *ServerChartDto) UnsetOpenstack()`

UnsetOpenstack ensures that no value is present for Openstack, not even an explicit nil
### GetGoogle

`func (o *ServerChartDto) GetGoogle() []ServerCommonRecordDto`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *ServerChartDto) GetGoogleOk() (*[]ServerCommonRecordDto, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *ServerChartDto) SetGoogle(v []ServerCommonRecordDto)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *ServerChartDto) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### SetGoogleNil

`func (o *ServerChartDto) SetGoogleNil(b bool)`

 SetGoogleNil sets the value for Google to be an explicit nil

### UnsetGoogle
`func (o *ServerChartDto) UnsetGoogle()`

UnsetGoogle ensures that no value is present for Google, not even an explicit nil
### GetGenericK8S

`func (o *ServerChartDto) GetGenericK8S() []ServerCommonRecordDto`

GetGenericK8S returns the GenericK8S field if non-nil, zero value otherwise.

### GetGenericK8SOk

`func (o *ServerChartDto) GetGenericK8SOk() (*[]ServerCommonRecordDto, bool)`

GetGenericK8SOk returns a tuple with the GenericK8S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericK8S

`func (o *ServerChartDto) SetGenericK8S(v []ServerCommonRecordDto)`

SetGenericK8S sets GenericK8S field to given value.

### HasGenericK8S

`func (o *ServerChartDto) HasGenericK8S() bool`

HasGenericK8S returns a boolean if a field has been set.

### SetGenericK8SNil

`func (o *ServerChartDto) SetGenericK8SNil(b bool)`

 SetGenericK8SNil sets the value for GenericK8S to be an explicit nil

### UnsetGenericK8S
`func (o *ServerChartDto) UnsetGenericK8S()`

UnsetGenericK8S ensures that no value is present for GenericK8S, not even an explicit nil
### GetFailed

`func (o *ServerChartDto) GetFailed() []ServerCommonRecordDto`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ServerChartDto) GetFailedOk() (*[]ServerCommonRecordDto, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ServerChartDto) SetFailed(v []ServerCommonRecordDto)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ServerChartDto) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ServerChartDto) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ServerChartDto) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetSucceeded

`func (o *ServerChartDto) GetSucceeded() []ServerCommonRecordDto`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ServerChartDto) GetSucceededOk() (*[]ServerCommonRecordDto, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ServerChartDto) SetSucceeded(v []ServerCommonRecordDto)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *ServerChartDto) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *ServerChartDto) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *ServerChartDto) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetWaiting

`func (o *ServerChartDto) GetWaiting() []ServerCommonRecordDto`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *ServerChartDto) GetWaitingOk() (*[]ServerCommonRecordDto, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *ServerChartDto) SetWaiting(v []ServerCommonRecordDto)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *ServerChartDto) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.

### SetWaitingNil

`func (o *ServerChartDto) SetWaitingNil(b bool)`

 SetWaitingNil sets the value for Waiting to be an explicit nil

### UnsetWaiting
`func (o *ServerChartDto) UnsetWaiting()`

UnsetWaiting ensures that no value is present for Waiting, not even an explicit nil
### GetUpdating

`func (o *ServerChartDto) GetUpdating() []ServerCommonRecordDto`

GetUpdating returns the Updating field if non-nil, zero value otherwise.

### GetUpdatingOk

`func (o *ServerChartDto) GetUpdatingOk() (*[]ServerCommonRecordDto, bool)`

GetUpdatingOk returns a tuple with the Updating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdating

`func (o *ServerChartDto) SetUpdating(v []ServerCommonRecordDto)`

SetUpdating sets Updating field to given value.

### HasUpdating

`func (o *ServerChartDto) HasUpdating() bool`

HasUpdating returns a boolean if a field has been set.

### SetUpdatingNil

`func (o *ServerChartDto) SetUpdatingNil(b bool)`

 SetUpdatingNil sets the value for Updating to be an explicit nil

### UnsetUpdating
`func (o *ServerChartDto) UnsetUpdating()`

UnsetUpdating ensures that no value is present for Updating, not even an explicit nil
### GetDeleting

`func (o *ServerChartDto) GetDeleting() []ServerCommonRecordDto`

GetDeleting returns the Deleting field if non-nil, zero value otherwise.

### GetDeletingOk

`func (o *ServerChartDto) GetDeletingOk() (*[]ServerCommonRecordDto, bool)`

GetDeletingOk returns a tuple with the Deleting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleting

`func (o *ServerChartDto) SetDeleting(v []ServerCommonRecordDto)`

SetDeleting sets Deleting field to given value.

### HasDeleting

`func (o *ServerChartDto) HasDeleting() bool`

HasDeleting returns a boolean if a field has been set.

### SetDeletingNil

`func (o *ServerChartDto) SetDeletingNil(b bool)`

 SetDeletingNil sets the value for Deleting to be an explicit nil

### UnsetDeleting
`func (o *ServerChartDto) UnsetDeleting()`

UnsetDeleting ensures that no value is present for Deleting, not even an explicit nil
### GetPurging

`func (o *ServerChartDto) GetPurging() []ServerCommonRecordDto`

GetPurging returns the Purging field if non-nil, zero value otherwise.

### GetPurgingOk

`func (o *ServerChartDto) GetPurgingOk() (*[]ServerCommonRecordDto, bool)`

GetPurgingOk returns a tuple with the Purging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurging

`func (o *ServerChartDto) SetPurging(v []ServerCommonRecordDto)`

SetPurging sets Purging field to given value.

### HasPurging

`func (o *ServerChartDto) HasPurging() bool`

HasPurging returns a boolean if a field has been set.

### SetPurgingNil

`func (o *ServerChartDto) SetPurgingNil(b bool)`

 SetPurgingNil sets the value for Purging to be an explicit nil

### UnsetPurging
`func (o *ServerChartDto) UnsetPurging()`

UnsetPurging ensures that no value is present for Purging, not even an explicit nil
### GetTotalCount

`func (o *ServerChartDto) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServerChartDto) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServerChartDto) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServerChartDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalCpu

`func (o *ServerChartDto) GetTotalCpu() int32`

GetTotalCpu returns the TotalCpu field if non-nil, zero value otherwise.

### GetTotalCpuOk

`func (o *ServerChartDto) GetTotalCpuOk() (*int32, bool)`

GetTotalCpuOk returns a tuple with the TotalCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpu

`func (o *ServerChartDto) SetTotalCpu(v int32)`

SetTotalCpu sets TotalCpu field to given value.

### HasTotalCpu

`func (o *ServerChartDto) HasTotalCpu() bool`

HasTotalCpu returns a boolean if a field has been set.

### GetTotalRam

`func (o *ServerChartDto) GetTotalRam() int64`

GetTotalRam returns the TotalRam field if non-nil, zero value otherwise.

### GetTotalRamOk

`func (o *ServerChartDto) GetTotalRamOk() (*int64, bool)`

GetTotalRamOk returns a tuple with the TotalRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRam

`func (o *ServerChartDto) SetTotalRam(v int64)`

SetTotalRam sets TotalRam field to given value.

### HasTotalRam

`func (o *ServerChartDto) HasTotalRam() bool`

HasTotalRam returns a boolean if a field has been set.

### GetTotalDiskSize

`func (o *ServerChartDto) GetTotalDiskSize() int64`

GetTotalDiskSize returns the TotalDiskSize field if non-nil, zero value otherwise.

### GetTotalDiskSizeOk

`func (o *ServerChartDto) GetTotalDiskSizeOk() (*int64, bool)`

GetTotalDiskSizeOk returns a tuple with the TotalDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiskSize

`func (o *ServerChartDto) SetTotalDiskSize(v int64)`

SetTotalDiskSize sets TotalDiskSize field to given value.

### HasTotalDiskSize

`func (o *ServerChartDto) HasTotalDiskSize() bool`

HasTotalDiskSize returns a boolean if a field has been set.

### GetTotalFailedCount

`func (o *ServerChartDto) GetTotalFailedCount() int32`

GetTotalFailedCount returns the TotalFailedCount field if non-nil, zero value otherwise.

### GetTotalFailedCountOk

`func (o *ServerChartDto) GetTotalFailedCountOk() (*int32, bool)`

GetTotalFailedCountOk returns a tuple with the TotalFailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailedCount

`func (o *ServerChartDto) SetTotalFailedCount(v int32)`

SetTotalFailedCount sets TotalFailedCount field to given value.

### HasTotalFailedCount

`func (o *ServerChartDto) HasTotalFailedCount() bool`

HasTotalFailedCount returns a boolean if a field has been set.

### GetTotalSucceededCount

`func (o *ServerChartDto) GetTotalSucceededCount() int32`

GetTotalSucceededCount returns the TotalSucceededCount field if non-nil, zero value otherwise.

### GetTotalSucceededCountOk

`func (o *ServerChartDto) GetTotalSucceededCountOk() (*int32, bool)`

GetTotalSucceededCountOk returns a tuple with the TotalSucceededCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSucceededCount

`func (o *ServerChartDto) SetTotalSucceededCount(v int32)`

SetTotalSucceededCount sets TotalSucceededCount field to given value.

### HasTotalSucceededCount

`func (o *ServerChartDto) HasTotalSucceededCount() bool`

HasTotalSucceededCount returns a boolean if a field has been set.

### GetTotalUpdatingCount

`func (o *ServerChartDto) GetTotalUpdatingCount() int32`

GetTotalUpdatingCount returns the TotalUpdatingCount field if non-nil, zero value otherwise.

### GetTotalUpdatingCountOk

`func (o *ServerChartDto) GetTotalUpdatingCountOk() (*int32, bool)`

GetTotalUpdatingCountOk returns a tuple with the TotalUpdatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpdatingCount

`func (o *ServerChartDto) SetTotalUpdatingCount(v int32)`

SetTotalUpdatingCount sets TotalUpdatingCount field to given value.

### HasTotalUpdatingCount

`func (o *ServerChartDto) HasTotalUpdatingCount() bool`

HasTotalUpdatingCount returns a boolean if a field has been set.

### GetTotalPendingCount

`func (o *ServerChartDto) GetTotalPendingCount() int32`

GetTotalPendingCount returns the TotalPendingCount field if non-nil, zero value otherwise.

### GetTotalPendingCountOk

`func (o *ServerChartDto) GetTotalPendingCountOk() (*int32, bool)`

GetTotalPendingCountOk returns a tuple with the TotalPendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPendingCount

`func (o *ServerChartDto) SetTotalPendingCount(v int32)`

SetTotalPendingCount sets TotalPendingCount field to given value.

### HasTotalPendingCount

`func (o *ServerChartDto) HasTotalPendingCount() bool`

HasTotalPendingCount returns a boolean if a field has been set.

### GetTotalAwsCount

`func (o *ServerChartDto) GetTotalAwsCount() int32`

GetTotalAwsCount returns the TotalAwsCount field if non-nil, zero value otherwise.

### GetTotalAwsCountOk

`func (o *ServerChartDto) GetTotalAwsCountOk() (*int32, bool)`

GetTotalAwsCountOk returns a tuple with the TotalAwsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAwsCount

`func (o *ServerChartDto) SetTotalAwsCount(v int32)`

SetTotalAwsCount sets TotalAwsCount field to given value.

### HasTotalAwsCount

`func (o *ServerChartDto) HasTotalAwsCount() bool`

HasTotalAwsCount returns a boolean if a field has been set.

### GetTotalAzureCount

`func (o *ServerChartDto) GetTotalAzureCount() int32`

GetTotalAzureCount returns the TotalAzureCount field if non-nil, zero value otherwise.

### GetTotalAzureCountOk

`func (o *ServerChartDto) GetTotalAzureCountOk() (*int32, bool)`

GetTotalAzureCountOk returns a tuple with the TotalAzureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAzureCount

`func (o *ServerChartDto) SetTotalAzureCount(v int32)`

SetTotalAzureCount sets TotalAzureCount field to given value.

### HasTotalAzureCount

`func (o *ServerChartDto) HasTotalAzureCount() bool`

HasTotalAzureCount returns a boolean if a field has been set.

### GetTotalOpenstackCount

`func (o *ServerChartDto) GetTotalOpenstackCount() int32`

GetTotalOpenstackCount returns the TotalOpenstackCount field if non-nil, zero value otherwise.

### GetTotalOpenstackCountOk

`func (o *ServerChartDto) GetTotalOpenstackCountOk() (*int32, bool)`

GetTotalOpenstackCountOk returns a tuple with the TotalOpenstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenstackCount

`func (o *ServerChartDto) SetTotalOpenstackCount(v int32)`

SetTotalOpenstackCount sets TotalOpenstackCount field to given value.

### HasTotalOpenstackCount

`func (o *ServerChartDto) HasTotalOpenstackCount() bool`

HasTotalOpenstackCount returns a boolean if a field has been set.

### GetTotalGoogleCount

`func (o *ServerChartDto) GetTotalGoogleCount() int32`

GetTotalGoogleCount returns the TotalGoogleCount field if non-nil, zero value otherwise.

### GetTotalGoogleCountOk

`func (o *ServerChartDto) GetTotalGoogleCountOk() (*int32, bool)`

GetTotalGoogleCountOk returns a tuple with the TotalGoogleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGoogleCount

`func (o *ServerChartDto) SetTotalGoogleCount(v int32)`

SetTotalGoogleCount sets TotalGoogleCount field to given value.

### HasTotalGoogleCount

`func (o *ServerChartDto) HasTotalGoogleCount() bool`

HasTotalGoogleCount returns a boolean if a field has been set.

### GetTotalGenericK8SCount

`func (o *ServerChartDto) GetTotalGenericK8SCount() int32`

GetTotalGenericK8SCount returns the TotalGenericK8SCount field if non-nil, zero value otherwise.

### GetTotalGenericK8SCountOk

`func (o *ServerChartDto) GetTotalGenericK8SCountOk() (*int32, bool)`

GetTotalGenericK8SCountOk returns a tuple with the TotalGenericK8SCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGenericK8SCount

`func (o *ServerChartDto) SetTotalGenericK8SCount(v int32)`

SetTotalGenericK8SCount sets TotalGenericK8SCount field to given value.

### HasTotalGenericK8SCount

`func (o *ServerChartDto) HasTotalGenericK8SCount() bool`

HasTotalGenericK8SCount returns a boolean if a field has been set.

### GetUsedResources

`func (o *ServerChartDto) GetUsedResources() []UserResourceChartDto`

GetUsedResources returns the UsedResources field if non-nil, zero value otherwise.

### GetUsedResourcesOk

`func (o *ServerChartDto) GetUsedResourcesOk() (*[]UserResourceChartDto, bool)`

GetUsedResourcesOk returns a tuple with the UsedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedResources

`func (o *ServerChartDto) SetUsedResources(v []UserResourceChartDto)`

SetUsedResources sets UsedResources field to given value.

### HasUsedResources

`func (o *ServerChartDto) HasUsedResources() bool`

HasUsedResources returns a boolean if a field has been set.

### SetUsedResourcesNil

`func (o *ServerChartDto) SetUsedResourcesNil(b bool)`

 SetUsedResourcesNil sets the value for UsedResources to be an explicit nil

### UnsetUsedResources
`func (o *ServerChartDto) UnsetUsedResources()`

UnsetUsedResources ensures that no value is present for UsedResources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


