# ResetServerStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ServerIds** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to [**CloudStatus**](CloudStatus.md) |  | [optional] 

## Methods

### NewResetServerStatusCommand

`func NewResetServerStatusCommand() *ResetServerStatusCommand`

NewResetServerStatusCommand instantiates a new ResetServerStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetServerStatusCommandWithDefaults

`func NewResetServerStatusCommandWithDefaults() *ResetServerStatusCommand`

NewResetServerStatusCommandWithDefaults instantiates a new ResetServerStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ResetServerStatusCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ResetServerStatusCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ResetServerStatusCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ResetServerStatusCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServerIds

`func (o *ResetServerStatusCommand) GetServerIds() []int32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *ResetServerStatusCommand) GetServerIdsOk() (*[]int32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *ResetServerStatusCommand) SetServerIds(v []int32)`

SetServerIds sets ServerIds field to given value.

### HasServerIds

`func (o *ResetServerStatusCommand) HasServerIds() bool`

HasServerIds returns a boolean if a field has been set.

### SetServerIdsNil

`func (o *ResetServerStatusCommand) SetServerIdsNil(b bool)`

 SetServerIdsNil sets the value for ServerIds to be an explicit nil

### UnsetServerIds
`func (o *ResetServerStatusCommand) UnsetServerIds()`

UnsetServerIds ensures that no value is present for ServerIds, not even an explicit nil
### GetStatus

`func (o *ResetServerStatusCommand) GetStatus() CloudStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResetServerStatusCommand) GetStatusOk() (*CloudStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResetServerStatusCommand) SetStatus(v CloudStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResetServerStatusCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


