# RamLimitsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxClusterRequests** | **int64** |  | 
**MaxClusterLimits** | **int64** |  | 
**DefaultContainerRequest** | **int64** |  | 
**DefaultContainerLimit** | **int64** |  | 

## Methods

### NewRamLimitsDto

`func NewRamLimitsDto(maxClusterRequests int64, maxClusterLimits int64, defaultContainerRequest int64, defaultContainerLimit int64, ) *RamLimitsDto`

NewRamLimitsDto instantiates a new RamLimitsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRamLimitsDtoWithDefaults

`func NewRamLimitsDtoWithDefaults() *RamLimitsDto`

NewRamLimitsDtoWithDefaults instantiates a new RamLimitsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxClusterRequests

`func (o *RamLimitsDto) GetMaxClusterRequests() int64`

GetMaxClusterRequests returns the MaxClusterRequests field if non-nil, zero value otherwise.

### GetMaxClusterRequestsOk

`func (o *RamLimitsDto) GetMaxClusterRequestsOk() (*int64, bool)`

GetMaxClusterRequestsOk returns a tuple with the MaxClusterRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterRequests

`func (o *RamLimitsDto) SetMaxClusterRequests(v int64)`

SetMaxClusterRequests sets MaxClusterRequests field to given value.


### GetMaxClusterLimits

`func (o *RamLimitsDto) GetMaxClusterLimits() int64`

GetMaxClusterLimits returns the MaxClusterLimits field if non-nil, zero value otherwise.

### GetMaxClusterLimitsOk

`func (o *RamLimitsDto) GetMaxClusterLimitsOk() (*int64, bool)`

GetMaxClusterLimitsOk returns a tuple with the MaxClusterLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterLimits

`func (o *RamLimitsDto) SetMaxClusterLimits(v int64)`

SetMaxClusterLimits sets MaxClusterLimits field to given value.


### GetDefaultContainerRequest

`func (o *RamLimitsDto) GetDefaultContainerRequest() int64`

GetDefaultContainerRequest returns the DefaultContainerRequest field if non-nil, zero value otherwise.

### GetDefaultContainerRequestOk

`func (o *RamLimitsDto) GetDefaultContainerRequestOk() (*int64, bool)`

GetDefaultContainerRequestOk returns a tuple with the DefaultContainerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerRequest

`func (o *RamLimitsDto) SetDefaultContainerRequest(v int64)`

SetDefaultContainerRequest sets DefaultContainerRequest field to given value.


### GetDefaultContainerLimit

`func (o *RamLimitsDto) GetDefaultContainerLimit() int64`

GetDefaultContainerLimit returns the DefaultContainerLimit field if non-nil, zero value otherwise.

### GetDefaultContainerLimitOk

`func (o *RamLimitsDto) GetDefaultContainerLimitOk() (*int64, bool)`

GetDefaultContainerLimitOk returns a tuple with the DefaultContainerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerLimit

`func (o *RamLimitsDto) SetDefaultContainerLimit(v int64)`

SetDefaultContainerLimit sets DefaultContainerLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


