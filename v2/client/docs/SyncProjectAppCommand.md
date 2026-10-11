# SyncProjectAppCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectAppId** | **int32** |  | 
**Timeout** | **NullableInt32** |  | 

## Methods

### NewSyncProjectAppCommand

`func NewSyncProjectAppCommand(projectAppId int32, timeout NullableInt32, ) *SyncProjectAppCommand`

NewSyncProjectAppCommand instantiates a new SyncProjectAppCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncProjectAppCommandWithDefaults

`func NewSyncProjectAppCommandWithDefaults() *SyncProjectAppCommand`

NewSyncProjectAppCommandWithDefaults instantiates a new SyncProjectAppCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectAppId

`func (o *SyncProjectAppCommand) GetProjectAppId() int32`

GetProjectAppId returns the ProjectAppId field if non-nil, zero value otherwise.

### GetProjectAppIdOk

`func (o *SyncProjectAppCommand) GetProjectAppIdOk() (*int32, bool)`

GetProjectAppIdOk returns a tuple with the ProjectAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppId

`func (o *SyncProjectAppCommand) SetProjectAppId(v int32)`

SetProjectAppId sets ProjectAppId field to given value.


### GetTimeout

`func (o *SyncProjectAppCommand) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyncProjectAppCommand) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SyncProjectAppCommand) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *SyncProjectAppCommand) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SyncProjectAppCommand) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


