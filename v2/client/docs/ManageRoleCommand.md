# ManageRoleCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** |  | 
**TargetRole** | [**EGlobalRole**](EGlobalRole.md) |  | 

## Methods

### NewManageRoleCommand

`func NewManageRoleCommand(userId NullableString, targetRole EGlobalRole, ) *ManageRoleCommand`

NewManageRoleCommand instantiates a new ManageRoleCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageRoleCommandWithDefaults

`func NewManageRoleCommandWithDefaults() *ManageRoleCommand`

NewManageRoleCommandWithDefaults instantiates a new ManageRoleCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ManageRoleCommand) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ManageRoleCommand) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ManageRoleCommand) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *ManageRoleCommand) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ManageRoleCommand) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTargetRole

`func (o *ManageRoleCommand) GetTargetRole() EGlobalRole`

GetTargetRole returns the TargetRole field if non-nil, zero value otherwise.

### GetTargetRoleOk

`func (o *ManageRoleCommand) GetTargetRoleOk() (*EGlobalRole, bool)`

GetTargetRoleOk returns a tuple with the TargetRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRole

`func (o *ManageRoleCommand) SetTargetRole(v EGlobalRole)`

SetTargetRole sets TargetRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


