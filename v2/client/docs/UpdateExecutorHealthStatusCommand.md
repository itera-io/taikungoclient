# UpdateExecutorHealthStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutorId** | **int32** |  | 
**Health** | [**ExecutorHealth**](ExecutorHealth.md) |  | 

## Methods

### NewUpdateExecutorHealthStatusCommand

`func NewUpdateExecutorHealthStatusCommand(executorId int32, health ExecutorHealth, ) *UpdateExecutorHealthStatusCommand`

NewUpdateExecutorHealthStatusCommand instantiates a new UpdateExecutorHealthStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExecutorHealthStatusCommandWithDefaults

`func NewUpdateExecutorHealthStatusCommandWithDefaults() *UpdateExecutorHealthStatusCommand`

NewUpdateExecutorHealthStatusCommandWithDefaults instantiates a new UpdateExecutorHealthStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutorId

`func (o *UpdateExecutorHealthStatusCommand) GetExecutorId() int32`

GetExecutorId returns the ExecutorId field if non-nil, zero value otherwise.

### GetExecutorIdOk

`func (o *UpdateExecutorHealthStatusCommand) GetExecutorIdOk() (*int32, bool)`

GetExecutorIdOk returns a tuple with the ExecutorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorId

`func (o *UpdateExecutorHealthStatusCommand) SetExecutorId(v int32)`

SetExecutorId sets ExecutorId field to given value.


### GetHealth

`func (o *UpdateExecutorHealthStatusCommand) GetHealth() ExecutorHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateExecutorHealthStatusCommand) GetHealthOk() (*ExecutorHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateExecutorHealthStatusCommand) SetHealth(v ExecutorHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


