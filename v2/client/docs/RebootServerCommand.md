# RebootServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int32** |  | 
**Type** | **NullableString** |  | 

## Methods

### NewRebootServerCommand

`func NewRebootServerCommand(serverId int32, type_ NullableString, ) *RebootServerCommand`

NewRebootServerCommand instantiates a new RebootServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootServerCommandWithDefaults

`func NewRebootServerCommandWithDefaults() *RebootServerCommand`

NewRebootServerCommandWithDefaults instantiates a new RebootServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *RebootServerCommand) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *RebootServerCommand) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *RebootServerCommand) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetType

`func (o *RebootServerCommand) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RebootServerCommand) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RebootServerCommand) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *RebootServerCommand) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RebootServerCommand) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


