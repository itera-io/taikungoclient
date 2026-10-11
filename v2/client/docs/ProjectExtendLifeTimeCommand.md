# ProjectExtendLifeTimeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**ExpireAt** | **NullableTime** |  | 
**DeleteOnExpiration** | **bool** |  | 

## Methods

### NewProjectExtendLifeTimeCommand

`func NewProjectExtendLifeTimeCommand(projectId int32, expireAt NullableTime, deleteOnExpiration bool, ) *ProjectExtendLifeTimeCommand`

NewProjectExtendLifeTimeCommand instantiates a new ProjectExtendLifeTimeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectExtendLifeTimeCommandWithDefaults

`func NewProjectExtendLifeTimeCommandWithDefaults() *ProjectExtendLifeTimeCommand`

NewProjectExtendLifeTimeCommandWithDefaults instantiates a new ProjectExtendLifeTimeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectExtendLifeTimeCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectExtendLifeTimeCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectExtendLifeTimeCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetExpireAt

`func (o *ProjectExtendLifeTimeCommand) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *ProjectExtendLifeTimeCommand) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *ProjectExtendLifeTimeCommand) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.


### SetExpireAtNil

`func (o *ProjectExtendLifeTimeCommand) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *ProjectExtendLifeTimeCommand) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetDeleteOnExpiration

`func (o *ProjectExtendLifeTimeCommand) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *ProjectExtendLifeTimeCommand) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *ProjectExtendLifeTimeCommand) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


