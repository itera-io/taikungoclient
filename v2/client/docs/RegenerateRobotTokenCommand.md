# RegenerateRobotTokenCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ExpiresAt** | **NullableTime** |  | 

## Methods

### NewRegenerateRobotTokenCommand

`func NewRegenerateRobotTokenCommand(id NullableString, expiresAt NullableTime, ) *RegenerateRobotTokenCommand`

NewRegenerateRobotTokenCommand instantiates a new RegenerateRobotTokenCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegenerateRobotTokenCommandWithDefaults

`func NewRegenerateRobotTokenCommandWithDefaults() *RegenerateRobotTokenCommand`

NewRegenerateRobotTokenCommandWithDefaults instantiates a new RegenerateRobotTokenCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegenerateRobotTokenCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegenerateRobotTokenCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegenerateRobotTokenCommand) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RegenerateRobotTokenCommand) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RegenerateRobotTokenCommand) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExpiresAt

`func (o *RegenerateRobotTokenCommand) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RegenerateRobotTokenCommand) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RegenerateRobotTokenCommand) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *RegenerateRobotTokenCommand) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *RegenerateRobotTokenCommand) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


