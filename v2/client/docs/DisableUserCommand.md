# DisableUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Disable** | **bool** |  | 

## Methods

### NewDisableUserCommand

`func NewDisableUserCommand(id NullableString, disable bool, ) *DisableUserCommand`

NewDisableUserCommand instantiates a new DisableUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableUserCommandWithDefaults

`func NewDisableUserCommandWithDefaults() *DisableUserCommand`

NewDisableUserCommandWithDefaults instantiates a new DisableUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisableUserCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableUserCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableUserCommand) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DisableUserCommand) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DisableUserCommand) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDisable

`func (o *DisableUserCommand) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DisableUserCommand) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DisableUserCommand) SetDisable(v bool)`

SetDisable sets Disable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


