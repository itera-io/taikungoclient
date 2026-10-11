# CreateVirtualClusterResourceLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUnit** | [**EProjectResourceUnit**](EProjectResourceUnit.md) |  | 
**MaxClusterRequests** | **float64** |  | 
**MaxClusterLimits** | **float64** |  | 
**DefaultContainerLimit** | **float64** |  | 
**DefaultContainerRequest** | **float64** |  | 

## Methods

### NewCreateVirtualClusterResourceLimits

`func NewCreateVirtualClusterResourceLimits(resourceUnit EProjectResourceUnit, maxClusterRequests float64, maxClusterLimits float64, defaultContainerLimit float64, defaultContainerRequest float64, ) *CreateVirtualClusterResourceLimits`

NewCreateVirtualClusterResourceLimits instantiates a new CreateVirtualClusterResourceLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualClusterResourceLimitsWithDefaults

`func NewCreateVirtualClusterResourceLimitsWithDefaults() *CreateVirtualClusterResourceLimits`

NewCreateVirtualClusterResourceLimitsWithDefaults instantiates a new CreateVirtualClusterResourceLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUnit

`func (o *CreateVirtualClusterResourceLimits) GetResourceUnit() EProjectResourceUnit`

GetResourceUnit returns the ResourceUnit field if non-nil, zero value otherwise.

### GetResourceUnitOk

`func (o *CreateVirtualClusterResourceLimits) GetResourceUnitOk() (*EProjectResourceUnit, bool)`

GetResourceUnitOk returns a tuple with the ResourceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUnit

`func (o *CreateVirtualClusterResourceLimits) SetResourceUnit(v EProjectResourceUnit)`

SetResourceUnit sets ResourceUnit field to given value.


### GetMaxClusterRequests

`func (o *CreateVirtualClusterResourceLimits) GetMaxClusterRequests() float64`

GetMaxClusterRequests returns the MaxClusterRequests field if non-nil, zero value otherwise.

### GetMaxClusterRequestsOk

`func (o *CreateVirtualClusterResourceLimits) GetMaxClusterRequestsOk() (*float64, bool)`

GetMaxClusterRequestsOk returns a tuple with the MaxClusterRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterRequests

`func (o *CreateVirtualClusterResourceLimits) SetMaxClusterRequests(v float64)`

SetMaxClusterRequests sets MaxClusterRequests field to given value.


### GetMaxClusterLimits

`func (o *CreateVirtualClusterResourceLimits) GetMaxClusterLimits() float64`

GetMaxClusterLimits returns the MaxClusterLimits field if non-nil, zero value otherwise.

### GetMaxClusterLimitsOk

`func (o *CreateVirtualClusterResourceLimits) GetMaxClusterLimitsOk() (*float64, bool)`

GetMaxClusterLimitsOk returns a tuple with the MaxClusterLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClusterLimits

`func (o *CreateVirtualClusterResourceLimits) SetMaxClusterLimits(v float64)`

SetMaxClusterLimits sets MaxClusterLimits field to given value.


### GetDefaultContainerLimit

`func (o *CreateVirtualClusterResourceLimits) GetDefaultContainerLimit() float64`

GetDefaultContainerLimit returns the DefaultContainerLimit field if non-nil, zero value otherwise.

### GetDefaultContainerLimitOk

`func (o *CreateVirtualClusterResourceLimits) GetDefaultContainerLimitOk() (*float64, bool)`

GetDefaultContainerLimitOk returns a tuple with the DefaultContainerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerLimit

`func (o *CreateVirtualClusterResourceLimits) SetDefaultContainerLimit(v float64)`

SetDefaultContainerLimit sets DefaultContainerLimit field to given value.


### GetDefaultContainerRequest

`func (o *CreateVirtualClusterResourceLimits) GetDefaultContainerRequest() float64`

GetDefaultContainerRequest returns the DefaultContainerRequest field if non-nil, zero value otherwise.

### GetDefaultContainerRequestOk

`func (o *CreateVirtualClusterResourceLimits) GetDefaultContainerRequestOk() (*float64, bool)`

GetDefaultContainerRequestOk returns a tuple with the DefaultContainerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContainerRequest

`func (o *CreateVirtualClusterResourceLimits) SetDefaultContainerRequest(v float64)`

SetDefaultContainerRequest sets DefaultContainerRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


