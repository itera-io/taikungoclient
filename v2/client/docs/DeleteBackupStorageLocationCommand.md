# DeleteBackupStorageLocationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**StorageLocation** | **NullableString** |  | 

## Methods

### NewDeleteBackupStorageLocationCommand

`func NewDeleteBackupStorageLocationCommand(projectId int32, storageLocation NullableString, ) *DeleteBackupStorageLocationCommand`

NewDeleteBackupStorageLocationCommand instantiates a new DeleteBackupStorageLocationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupStorageLocationCommandWithDefaults

`func NewDeleteBackupStorageLocationCommandWithDefaults() *DeleteBackupStorageLocationCommand`

NewDeleteBackupStorageLocationCommandWithDefaults instantiates a new DeleteBackupStorageLocationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteBackupStorageLocationCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteBackupStorageLocationCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteBackupStorageLocationCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetStorageLocation

`func (o *DeleteBackupStorageLocationCommand) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *DeleteBackupStorageLocationCommand) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *DeleteBackupStorageLocationCommand) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### SetStorageLocationNil

`func (o *DeleteBackupStorageLocationCommand) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *DeleteBackupStorageLocationCommand) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


