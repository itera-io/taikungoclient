# DeleteUserFromGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** |  | 
**Users** | **[]string** |  | 

## Methods

### NewDeleteUserFromGroupCommand

`func NewDeleteUserFromGroupCommand(groupId int32, users []string, ) *DeleteUserFromGroupCommand`

NewDeleteUserFromGroupCommand instantiates a new DeleteUserFromGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserFromGroupCommandWithDefaults

`func NewDeleteUserFromGroupCommandWithDefaults() *DeleteUserFromGroupCommand`

NewDeleteUserFromGroupCommandWithDefaults instantiates a new DeleteUserFromGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *DeleteUserFromGroupCommand) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DeleteUserFromGroupCommand) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DeleteUserFromGroupCommand) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetUsers

`func (o *DeleteUserFromGroupCommand) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DeleteUserFromGroupCommand) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DeleteUserFromGroupCommand) SetUsers(v []string)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *DeleteUserFromGroupCommand) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *DeleteUserFromGroupCommand) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


