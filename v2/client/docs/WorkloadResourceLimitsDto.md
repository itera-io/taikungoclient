# WorkloadResourceLimitsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPods** | **int32** |  | 
**MaxPvcs** | **int32** |  | 
**MaxTotalPvcSize** | **int64** |  | 
**MaxIngresses** | **int32** |  | 
**MaxLoadBalancers** | **int32** |  | 

## Methods

### NewWorkloadResourceLimitsDto

`func NewWorkloadResourceLimitsDto(maxPods int32, maxPvcs int32, maxTotalPvcSize int64, maxIngresses int32, maxLoadBalancers int32, ) *WorkloadResourceLimitsDto`

NewWorkloadResourceLimitsDto instantiates a new WorkloadResourceLimitsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadResourceLimitsDtoWithDefaults

`func NewWorkloadResourceLimitsDtoWithDefaults() *WorkloadResourceLimitsDto`

NewWorkloadResourceLimitsDtoWithDefaults instantiates a new WorkloadResourceLimitsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPods

`func (o *WorkloadResourceLimitsDto) GetMaxPods() int32`

GetMaxPods returns the MaxPods field if non-nil, zero value otherwise.

### GetMaxPodsOk

`func (o *WorkloadResourceLimitsDto) GetMaxPodsOk() (*int32, bool)`

GetMaxPodsOk returns a tuple with the MaxPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPods

`func (o *WorkloadResourceLimitsDto) SetMaxPods(v int32)`

SetMaxPods sets MaxPods field to given value.


### GetMaxPvcs

`func (o *WorkloadResourceLimitsDto) GetMaxPvcs() int32`

GetMaxPvcs returns the MaxPvcs field if non-nil, zero value otherwise.

### GetMaxPvcsOk

`func (o *WorkloadResourceLimitsDto) GetMaxPvcsOk() (*int32, bool)`

GetMaxPvcsOk returns a tuple with the MaxPvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPvcs

`func (o *WorkloadResourceLimitsDto) SetMaxPvcs(v int32)`

SetMaxPvcs sets MaxPvcs field to given value.


### GetMaxTotalPvcSize

`func (o *WorkloadResourceLimitsDto) GetMaxTotalPvcSize() int64`

GetMaxTotalPvcSize returns the MaxTotalPvcSize field if non-nil, zero value otherwise.

### GetMaxTotalPvcSizeOk

`func (o *WorkloadResourceLimitsDto) GetMaxTotalPvcSizeOk() (*int64, bool)`

GetMaxTotalPvcSizeOk returns a tuple with the MaxTotalPvcSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalPvcSize

`func (o *WorkloadResourceLimitsDto) SetMaxTotalPvcSize(v int64)`

SetMaxTotalPvcSize sets MaxTotalPvcSize field to given value.


### GetMaxIngresses

`func (o *WorkloadResourceLimitsDto) GetMaxIngresses() int32`

GetMaxIngresses returns the MaxIngresses field if non-nil, zero value otherwise.

### GetMaxIngressesOk

`func (o *WorkloadResourceLimitsDto) GetMaxIngressesOk() (*int32, bool)`

GetMaxIngressesOk returns a tuple with the MaxIngresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIngresses

`func (o *WorkloadResourceLimitsDto) SetMaxIngresses(v int32)`

SetMaxIngresses sets MaxIngresses field to given value.


### GetMaxLoadBalancers

`func (o *WorkloadResourceLimitsDto) GetMaxLoadBalancers() int32`

GetMaxLoadBalancers returns the MaxLoadBalancers field if non-nil, zero value otherwise.

### GetMaxLoadBalancersOk

`func (o *WorkloadResourceLimitsDto) GetMaxLoadBalancersOk() (*int32, bool)`

GetMaxLoadBalancersOk returns a tuple with the MaxLoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoadBalancers

`func (o *WorkloadResourceLimitsDto) SetMaxLoadBalancers(v int32)`

SetMaxLoadBalancers sets MaxLoadBalancers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


