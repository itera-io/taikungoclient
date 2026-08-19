# DeleteImportedProjectCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**IsForceDelete** | **bool** |  | 

## Methods

### NewDeleteImportedProjectCommand

`func NewDeleteImportedProjectCommand(projectId int32, isForceDelete bool, ) *DeleteImportedProjectCommand`

NewDeleteImportedProjectCommand instantiates a new DeleteImportedProjectCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteImportedProjectCommandWithDefaults

`func NewDeleteImportedProjectCommandWithDefaults() *DeleteImportedProjectCommand`

NewDeleteImportedProjectCommandWithDefaults instantiates a new DeleteImportedProjectCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteImportedProjectCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteImportedProjectCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteImportedProjectCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetIsForceDelete

`func (o *DeleteImportedProjectCommand) GetIsForceDelete() bool`

GetIsForceDelete returns the IsForceDelete field if non-nil, zero value otherwise.

### GetIsForceDeleteOk

`func (o *DeleteImportedProjectCommand) GetIsForceDeleteOk() (*bool, bool)`

GetIsForceDeleteOk returns a tuple with the IsForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceDelete

`func (o *DeleteImportedProjectCommand) SetIsForceDelete(v bool)`

SetIsForceDelete sets IsForceDelete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


