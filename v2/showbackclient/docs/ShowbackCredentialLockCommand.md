# ShowbackCredentialLockCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShowbackCredentialLockCommand

`func NewShowbackCredentialLockCommand() *ShowbackCredentialLockCommand`

NewShowbackCredentialLockCommand instantiates a new ShowbackCredentialLockCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowbackCredentialLockCommandWithDefaults

`func NewShowbackCredentialLockCommandWithDefaults() *ShowbackCredentialLockCommand`

NewShowbackCredentialLockCommandWithDefaults instantiates a new ShowbackCredentialLockCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShowbackCredentialLockCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowbackCredentialLockCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowbackCredentialLockCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShowbackCredentialLockCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *ShowbackCredentialLockCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ShowbackCredentialLockCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ShowbackCredentialLockCommand) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ShowbackCredentialLockCommand) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *ShowbackCredentialLockCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *ShowbackCredentialLockCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


