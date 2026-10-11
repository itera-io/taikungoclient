# DeleteOrganizationFromGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** |  | 
**Organizations** | **[]int32** |  | 

## Methods

### NewDeleteOrganizationFromGroupCommand

`func NewDeleteOrganizationFromGroupCommand(groupId int32, organizations []int32, ) *DeleteOrganizationFromGroupCommand`

NewDeleteOrganizationFromGroupCommand instantiates a new DeleteOrganizationFromGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteOrganizationFromGroupCommandWithDefaults

`func NewDeleteOrganizationFromGroupCommandWithDefaults() *DeleteOrganizationFromGroupCommand`

NewDeleteOrganizationFromGroupCommandWithDefaults instantiates a new DeleteOrganizationFromGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *DeleteOrganizationFromGroupCommand) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DeleteOrganizationFromGroupCommand) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DeleteOrganizationFromGroupCommand) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetOrganizations

`func (o *DeleteOrganizationFromGroupCommand) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *DeleteOrganizationFromGroupCommand) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *DeleteOrganizationFromGroupCommand) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.


### SetOrganizationsNil

`func (o *DeleteOrganizationFromGroupCommand) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *DeleteOrganizationFromGroupCommand) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


