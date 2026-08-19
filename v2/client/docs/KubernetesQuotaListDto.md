# KubernetesQuotaListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SumOfCpu** | **NullableFloat64** |  | 
**SumOfCpuUsage** | **float64** |  | 
**SumOfCpuRequests** | **NullableFloat64** |  | 
**SumOfCpuRequestsUsage** | **float64** |  | 
**SumOfRamInGb** | **NullableFloat64** |  | 
**SumOfRamUsage** | **float64** |  | 
**SumOfRamRequestsInGb** | **NullableFloat64** |  | 
**SumOfRamRequestsUsage** | **float64** |  | 
**PodsCapacity** | **NullableInt32** |  | 
**PodsTotalCount** | **int32** |  | 

## Methods

### NewKubernetesQuotaListDto

`func NewKubernetesQuotaListDto(sumOfCpu NullableFloat64, sumOfCpuUsage float64, sumOfCpuRequests NullableFloat64, sumOfCpuRequestsUsage float64, sumOfRamInGb NullableFloat64, sumOfRamUsage float64, sumOfRamRequestsInGb NullableFloat64, sumOfRamRequestsUsage float64, podsCapacity NullableInt32, podsTotalCount int32, ) *KubernetesQuotaListDto`

NewKubernetesQuotaListDto instantiates a new KubernetesQuotaListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesQuotaListDtoWithDefaults

`func NewKubernetesQuotaListDtoWithDefaults() *KubernetesQuotaListDto`

NewKubernetesQuotaListDtoWithDefaults instantiates a new KubernetesQuotaListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSumOfCpu

`func (o *KubernetesQuotaListDto) GetSumOfCpu() float64`

GetSumOfCpu returns the SumOfCpu field if non-nil, zero value otherwise.

### GetSumOfCpuOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuOk() (*float64, bool)`

GetSumOfCpuOk returns a tuple with the SumOfCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpu

`func (o *KubernetesQuotaListDto) SetSumOfCpu(v float64)`

SetSumOfCpu sets SumOfCpu field to given value.


### SetSumOfCpuNil

`func (o *KubernetesQuotaListDto) SetSumOfCpuNil(b bool)`

 SetSumOfCpuNil sets the value for SumOfCpu to be an explicit nil

### UnsetSumOfCpu
`func (o *KubernetesQuotaListDto) UnsetSumOfCpu()`

UnsetSumOfCpu ensures that no value is present for SumOfCpu, not even an explicit nil
### GetSumOfCpuUsage

`func (o *KubernetesQuotaListDto) GetSumOfCpuUsage() float64`

GetSumOfCpuUsage returns the SumOfCpuUsage field if non-nil, zero value otherwise.

### GetSumOfCpuUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuUsageOk() (*float64, bool)`

GetSumOfCpuUsageOk returns a tuple with the SumOfCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpuUsage

`func (o *KubernetesQuotaListDto) SetSumOfCpuUsage(v float64)`

SetSumOfCpuUsage sets SumOfCpuUsage field to given value.


### GetSumOfCpuRequests

`func (o *KubernetesQuotaListDto) GetSumOfCpuRequests() float64`

GetSumOfCpuRequests returns the SumOfCpuRequests field if non-nil, zero value otherwise.

### GetSumOfCpuRequestsOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuRequestsOk() (*float64, bool)`

GetSumOfCpuRequestsOk returns a tuple with the SumOfCpuRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpuRequests

`func (o *KubernetesQuotaListDto) SetSumOfCpuRequests(v float64)`

SetSumOfCpuRequests sets SumOfCpuRequests field to given value.


### SetSumOfCpuRequestsNil

`func (o *KubernetesQuotaListDto) SetSumOfCpuRequestsNil(b bool)`

 SetSumOfCpuRequestsNil sets the value for SumOfCpuRequests to be an explicit nil

### UnsetSumOfCpuRequests
`func (o *KubernetesQuotaListDto) UnsetSumOfCpuRequests()`

UnsetSumOfCpuRequests ensures that no value is present for SumOfCpuRequests, not even an explicit nil
### GetSumOfCpuRequestsUsage

`func (o *KubernetesQuotaListDto) GetSumOfCpuRequestsUsage() float64`

GetSumOfCpuRequestsUsage returns the SumOfCpuRequestsUsage field if non-nil, zero value otherwise.

### GetSumOfCpuRequestsUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfCpuRequestsUsageOk() (*float64, bool)`

GetSumOfCpuRequestsUsageOk returns a tuple with the SumOfCpuRequestsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfCpuRequestsUsage

`func (o *KubernetesQuotaListDto) SetSumOfCpuRequestsUsage(v float64)`

SetSumOfCpuRequestsUsage sets SumOfCpuRequestsUsage field to given value.


### GetSumOfRamInGb

`func (o *KubernetesQuotaListDto) GetSumOfRamInGb() float64`

GetSumOfRamInGb returns the SumOfRamInGb field if non-nil, zero value otherwise.

### GetSumOfRamInGbOk

`func (o *KubernetesQuotaListDto) GetSumOfRamInGbOk() (*float64, bool)`

GetSumOfRamInGbOk returns a tuple with the SumOfRamInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamInGb

`func (o *KubernetesQuotaListDto) SetSumOfRamInGb(v float64)`

SetSumOfRamInGb sets SumOfRamInGb field to given value.


### SetSumOfRamInGbNil

`func (o *KubernetesQuotaListDto) SetSumOfRamInGbNil(b bool)`

 SetSumOfRamInGbNil sets the value for SumOfRamInGb to be an explicit nil

### UnsetSumOfRamInGb
`func (o *KubernetesQuotaListDto) UnsetSumOfRamInGb()`

UnsetSumOfRamInGb ensures that no value is present for SumOfRamInGb, not even an explicit nil
### GetSumOfRamUsage

`func (o *KubernetesQuotaListDto) GetSumOfRamUsage() float64`

GetSumOfRamUsage returns the SumOfRamUsage field if non-nil, zero value otherwise.

### GetSumOfRamUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfRamUsageOk() (*float64, bool)`

GetSumOfRamUsageOk returns a tuple with the SumOfRamUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamUsage

`func (o *KubernetesQuotaListDto) SetSumOfRamUsage(v float64)`

SetSumOfRamUsage sets SumOfRamUsage field to given value.


### GetSumOfRamRequestsInGb

`func (o *KubernetesQuotaListDto) GetSumOfRamRequestsInGb() float64`

GetSumOfRamRequestsInGb returns the SumOfRamRequestsInGb field if non-nil, zero value otherwise.

### GetSumOfRamRequestsInGbOk

`func (o *KubernetesQuotaListDto) GetSumOfRamRequestsInGbOk() (*float64, bool)`

GetSumOfRamRequestsInGbOk returns a tuple with the SumOfRamRequestsInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamRequestsInGb

`func (o *KubernetesQuotaListDto) SetSumOfRamRequestsInGb(v float64)`

SetSumOfRamRequestsInGb sets SumOfRamRequestsInGb field to given value.


### SetSumOfRamRequestsInGbNil

`func (o *KubernetesQuotaListDto) SetSumOfRamRequestsInGbNil(b bool)`

 SetSumOfRamRequestsInGbNil sets the value for SumOfRamRequestsInGb to be an explicit nil

### UnsetSumOfRamRequestsInGb
`func (o *KubernetesQuotaListDto) UnsetSumOfRamRequestsInGb()`

UnsetSumOfRamRequestsInGb ensures that no value is present for SumOfRamRequestsInGb, not even an explicit nil
### GetSumOfRamRequestsUsage

`func (o *KubernetesQuotaListDto) GetSumOfRamRequestsUsage() float64`

GetSumOfRamRequestsUsage returns the SumOfRamRequestsUsage field if non-nil, zero value otherwise.

### GetSumOfRamRequestsUsageOk

`func (o *KubernetesQuotaListDto) GetSumOfRamRequestsUsageOk() (*float64, bool)`

GetSumOfRamRequestsUsageOk returns a tuple with the SumOfRamRequestsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOfRamRequestsUsage

`func (o *KubernetesQuotaListDto) SetSumOfRamRequestsUsage(v float64)`

SetSumOfRamRequestsUsage sets SumOfRamRequestsUsage field to given value.


### GetPodsCapacity

`func (o *KubernetesQuotaListDto) GetPodsCapacity() int32`

GetPodsCapacity returns the PodsCapacity field if non-nil, zero value otherwise.

### GetPodsCapacityOk

`func (o *KubernetesQuotaListDto) GetPodsCapacityOk() (*int32, bool)`

GetPodsCapacityOk returns a tuple with the PodsCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsCapacity

`func (o *KubernetesQuotaListDto) SetPodsCapacity(v int32)`

SetPodsCapacity sets PodsCapacity field to given value.


### SetPodsCapacityNil

`func (o *KubernetesQuotaListDto) SetPodsCapacityNil(b bool)`

 SetPodsCapacityNil sets the value for PodsCapacity to be an explicit nil

### UnsetPodsCapacity
`func (o *KubernetesQuotaListDto) UnsetPodsCapacity()`

UnsetPodsCapacity ensures that no value is present for PodsCapacity, not even an explicit nil
### GetPodsTotalCount

`func (o *KubernetesQuotaListDto) GetPodsTotalCount() int32`

GetPodsTotalCount returns the PodsTotalCount field if non-nil, zero value otherwise.

### GetPodsTotalCountOk

`func (o *KubernetesQuotaListDto) GetPodsTotalCountOk() (*int32, bool)`

GetPodsTotalCountOk returns a tuple with the PodsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsTotalCount

`func (o *KubernetesQuotaListDto) SetPodsTotalCount(v int32)`

SetPodsTotalCount sets PodsTotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


