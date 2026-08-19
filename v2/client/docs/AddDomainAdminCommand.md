# AddDomainAdminCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **NullableInt32** |  | [optional] 
**UserId** | **NullableString** |  | 

## Methods

### NewAddDomainAdminCommand

`func NewAddDomainAdminCommand(userId NullableString, ) *AddDomainAdminCommand`

NewAddDomainAdminCommand instantiates a new AddDomainAdminCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDomainAdminCommandWithDefaults

`func NewAddDomainAdminCommandWithDefaults() *AddDomainAdminCommand`

NewAddDomainAdminCommandWithDefaults instantiates a new AddDomainAdminCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *AddDomainAdminCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *AddDomainAdminCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *AddDomainAdminCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *AddDomainAdminCommand) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *AddDomainAdminCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *AddDomainAdminCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetUserId

`func (o *AddDomainAdminCommand) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AddDomainAdminCommand) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AddDomainAdminCommand) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *AddDomainAdminCommand) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AddDomainAdminCommand) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


