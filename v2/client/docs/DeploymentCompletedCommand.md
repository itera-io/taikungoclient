# DeploymentCompletedCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Result** | **NullableString** |  | 
**FromCronJob** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDeploymentCompletedCommand

`func NewDeploymentCompletedCommand(projectId int32, result NullableString, ) *DeploymentCompletedCommand`

NewDeploymentCompletedCommand instantiates a new DeploymentCompletedCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCompletedCommandWithDefaults

`func NewDeploymentCompletedCommandWithDefaults() *DeploymentCompletedCommand`

NewDeploymentCompletedCommandWithDefaults instantiates a new DeploymentCompletedCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeploymentCompletedCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeploymentCompletedCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeploymentCompletedCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetResult

`func (o *DeploymentCompletedCommand) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeploymentCompletedCommand) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeploymentCompletedCommand) SetResult(v string)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *DeploymentCompletedCommand) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *DeploymentCompletedCommand) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetFromCronJob

`func (o *DeploymentCompletedCommand) GetFromCronJob() bool`

GetFromCronJob returns the FromCronJob field if non-nil, zero value otherwise.

### GetFromCronJobOk

`func (o *DeploymentCompletedCommand) GetFromCronJobOk() (*bool, bool)`

GetFromCronJobOk returns a tuple with the FromCronJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCronJob

`func (o *DeploymentCompletedCommand) SetFromCronJob(v bool)`

SetFromCronJob sets FromCronJob field to given value.

### HasFromCronJob

`func (o *DeploymentCompletedCommand) HasFromCronJob() bool`

HasFromCronJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


