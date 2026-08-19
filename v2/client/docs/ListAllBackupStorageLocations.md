# ListAllBackupStorageLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BackupStorageLocationDto**](BackupStorageLocationDto.md) |  | 
**TotalCount** | **int32** |  | 
**Projects** | **[]int32** |  | 

## Methods

### NewListAllBackupStorageLocations

`func NewListAllBackupStorageLocations(data []BackupStorageLocationDto, totalCount int32, projects []int32, ) *ListAllBackupStorageLocations`

NewListAllBackupStorageLocations instantiates a new ListAllBackupStorageLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllBackupStorageLocationsWithDefaults

`func NewListAllBackupStorageLocationsWithDefaults() *ListAllBackupStorageLocations`

NewListAllBackupStorageLocationsWithDefaults instantiates a new ListAllBackupStorageLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllBackupStorageLocations) GetData() []BackupStorageLocationDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllBackupStorageLocations) GetDataOk() (*[]BackupStorageLocationDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllBackupStorageLocations) SetData(v []BackupStorageLocationDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ListAllBackupStorageLocations) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListAllBackupStorageLocations) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ListAllBackupStorageLocations) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAllBackupStorageLocations) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAllBackupStorageLocations) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetProjects

`func (o *ListAllBackupStorageLocations) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ListAllBackupStorageLocations) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ListAllBackupStorageLocations) SetProjects(v []int32)`

SetProjects sets Projects field to given value.


### SetProjectsNil

`func (o *ListAllBackupStorageLocations) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *ListAllBackupStorageLocations) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


