# PromoteUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** |  | 
**IsAdmin** | **bool** |  | 

## Methods

### NewPromoteUserCommand

`func NewPromoteUserCommand(userId NullableString, isAdmin bool, ) *PromoteUserCommand`

NewPromoteUserCommand instantiates a new PromoteUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteUserCommandWithDefaults

`func NewPromoteUserCommandWithDefaults() *PromoteUserCommand`

NewPromoteUserCommandWithDefaults instantiates a new PromoteUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PromoteUserCommand) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PromoteUserCommand) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PromoteUserCommand) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *PromoteUserCommand) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PromoteUserCommand) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetIsAdmin

`func (o *PromoteUserCommand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *PromoteUserCommand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *PromoteUserCommand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


