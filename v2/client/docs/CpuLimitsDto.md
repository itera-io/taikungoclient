# CpuLimitsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxClusterRequests** | **float64** |  | 
**MaxClusterLimits** | **float64** |  | 
**DefaultContainerRequest** | **float64** |  | 
**DefaultContainerLimit** | **float64** |  | 

## Methods

### NewCpuLimitsDto

`func NewCpuLimitsDto(maxClusterRequests float64, maxClusterLimits float64, defaultContainerRequest float64, defaultContainerLimit float64, ) *CpuLimitsDto`

NewCpuLimitsDto instantiates a new CpuLimitsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuLimitsDtoWithDefaults

`func NewCpuLimitsDtoWithDefaults() *CpuLimitsDto`

NewCpuLimitsDtoWithDefaults instantiates a new CpuLimitsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxClusterRequests

`func (o *CpuLimitsDto) GetMaxClusterRequests() float64`

GetMaxClusterRequests returns the MaxClusterRequests field if non-nil, zero value otherwise.

### GetMaxClusterRequestsOk

`func (o *CpuLimitsDto) GetMaxClusterRequestsOk() (*float64, bool)`

GetMaxClusterRequestsOk returns a tuple with the MaxClusterRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterRequests

`func (o *CpuLimitsDto) SetMaxClusterRequests(v float64)`

SetMaxClusterRequests sets MaxClusterRequests field to given value.


### GetMaxClusterLimits

`func (o *CpuLimitsDto) GetMaxClusterLimits() float64`

GetMaxClusterLimits returns the MaxClusterLimits field if non-nil, zero value otherwise.

### GetMaxClusterLimitsOk

`func (o *CpuLimitsDto) GetMaxClusterLimitsOk() (*float64, bool)`

GetMaxClusterLimitsOk returns a tuple with the MaxClusterLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterLimits

`func (o *CpuLimitsDto) SetMaxClusterLimits(v float64)`

SetMaxClusterLimits sets MaxClusterLimits field to given value.


### GetDefaultContainerRequest

`func (o *CpuLimitsDto) GetDefaultContainerRequest() float64`

GetDefaultContainerRequest returns the DefaultContainerRequest field if non-nil, zero value otherwise.

### GetDefaultContainerRequestOk

`func (o *CpuLimitsDto) GetDefaultContainerRequestOk() (*float64, bool)`

GetDefaultContainerRequestOk returns a tuple with the DefaultContainerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerRequest

`func (o *CpuLimitsDto) SetDefaultContainerRequest(v float64)`

SetDefaultContainerRequest sets DefaultContainerRequest field to given value.


### GetDefaultContainerLimit

`func (o *CpuLimitsDto) GetDefaultContainerLimit() float64`

GetDefaultContainerLimit returns the DefaultContainerLimit field if non-nil, zero value otherwise.

### GetDefaultContainerLimitOk

`func (o *CpuLimitsDto) GetDefaultContainerLimitOk() (*float64, bool)`

GetDefaultContainerLimitOk returns a tuple with the DefaultContainerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerLimit

`func (o *CpuLimitsDto) SetDefaultContainerLimit(v float64)`

SetDefaultContainerLimit sets DefaultContainerLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


