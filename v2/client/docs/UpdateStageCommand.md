# UpdateStageCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Stage** | **NullableString** |  | 

## Methods

### NewUpdateStageCommand

`func NewUpdateStageCommand(projectId int32, stage NullableString, ) *UpdateStageCommand`

NewUpdateStageCommand instantiates a new UpdateStageCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStageCommandWithDefaults

`func NewUpdateStageCommandWithDefaults() *UpdateStageCommand`

NewUpdateStageCommandWithDefaults instantiates a new UpdateStageCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *UpdateStageCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateStageCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateStageCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetStage

`func (o *UpdateStageCommand) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *UpdateStageCommand) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *UpdateStageCommand) SetStage(v string)`

SetStage sets Stage field to given value.


### SetStageNil

`func (o *UpdateStageCommand) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *UpdateStageCommand) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


