# DeleteBackupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewDeleteBackupCommand

`func NewDeleteBackupCommand(projectId int32, name NullableString, ) *DeleteBackupCommand`

NewDeleteBackupCommand instantiates a new DeleteBackupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupCommandWithDefaults

`func NewDeleteBackupCommandWithDefaults() *DeleteBackupCommand`

NewDeleteBackupCommandWithDefaults instantiates a new DeleteBackupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteBackupCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteBackupCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteBackupCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DeleteBackupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteBackupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteBackupCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeleteBackupCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeleteBackupCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


