# VClusterResourceLimitsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUnit** | **NullableString** |  | 
**MaxClusterRequests** | **float64** |  | 
**MaxClusterLimits** | **float64** |  | 
**DefaultContainerLimit** | **float64** |  | 
**DefaultContainerRequest** | **float64** |  | 

## Methods

### NewVClusterResourceLimitsListDto

`func NewVClusterResourceLimitsListDto(resourceUnit NullableString, maxClusterRequests float64, maxClusterLimits float64, defaultContainerLimit float64, defaultContainerRequest float64, ) *VClusterResourceLimitsListDto`

NewVClusterResourceLimitsListDto instantiates a new VClusterResourceLimitsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVClusterResourceLimitsListDtoWithDefaults

`func NewVClusterResourceLimitsListDtoWithDefaults() *VClusterResourceLimitsListDto`

NewVClusterResourceLimitsListDtoWithDefaults instantiates a new VClusterResourceLimitsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUnit

`func (o *VClusterResourceLimitsListDto) GetResourceUnit() string`

GetResourceUnit returns the ResourceUnit field if non-nil, zero value otherwise.

### GetResourceUnitOk

`func (o *VClusterResourceLimitsListDto) GetResourceUnitOk() (*string, bool)`

GetResourceUnitOk returns a tuple with the ResourceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUnit

`func (o *VClusterResourceLimitsListDto) SetResourceUnit(v string)`

SetResourceUnit sets ResourceUnit field to given value.


### SetResourceUnitNil

`func (o *VClusterResourceLimitsListDto) SetResourceUnitNil(b bool)`

 SetResourceUnitNil sets the value for ResourceUnit to be an explicit nil

### UnsetResourceUnit
`func (o *VClusterResourceLimitsListDto) UnsetResourceUnit()`

UnsetResourceUnit ensures that no value is present for ResourceUnit, not even an explicit nil
### GetMaxClusterRequests

`func (o *VClusterResourceLimitsListDto) GetMaxClusterRequests() float64`

GetMaxClusterRequests returns the MaxClusterRequests field if non-nil, zero value otherwise.

### GetMaxClusterRequestsOk

`func (o *VClusterResourceLimitsListDto) GetMaxClusterRequestsOk() (*float64, bool)`

GetMaxClusterRequestsOk returns a tuple with the MaxClusterRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterRequests

`func (o *VClusterResourceLimitsListDto) SetMaxClusterRequests(v float64)`

SetMaxClusterRequests sets MaxClusterRequests field to given value.


### GetMaxClusterLimits

`func (o *VClusterResourceLimitsListDto) GetMaxClusterLimits() float64`

GetMaxClusterLimits returns the MaxClusterLimits field if non-nil, zero value otherwise.

### GetMaxClusterLimitsOk

`func (o *VClusterResourceLimitsListDto) GetMaxClusterLimitsOk() (*float64, bool)`

GetMaxClusterLimitsOk returns a tuple with the MaxClusterLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterLimits

`func (o *VClusterResourceLimitsListDto) SetMaxClusterLimits(v float64)`

SetMaxClusterLimits sets MaxClusterLimits field to given value.


### GetDefaultContainerLimit

`func (o *VClusterResourceLimitsListDto) GetDefaultContainerLimit() float64`

GetDefaultContainerLimit returns the DefaultContainerLimit field if non-nil, zero value otherwise.

### GetDefaultContainerLimitOk

`func (o *VClusterResourceLimitsListDto) GetDefaultContainerLimitOk() (*float64, bool)`

GetDefaultContainerLimitOk returns a tuple with the DefaultContainerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerLimit

`func (o *VClusterResourceLimitsListDto) SetDefaultContainerLimit(v float64)`

SetDefaultContainerLimit sets DefaultContainerLimit field to given value.


### GetDefaultContainerRequest

`func (o *VClusterResourceLimitsListDto) GetDefaultContainerRequest() float64`

GetDefaultContainerRequest returns the DefaultContainerRequest field if non-nil, zero value otherwise.

### GetDefaultContainerRequestOk

`func (o *VClusterResourceLimitsListDto) GetDefaultContainerRequestOk() (*float64, bool)`

GetDefaultContainerRequestOk returns a tuple with the DefaultContainerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerRequest

`func (o *VClusterResourceLimitsListDto) SetDefaultContainerRequest(v float64)`

SetDefaultContainerRequest sets DefaultContainerRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


