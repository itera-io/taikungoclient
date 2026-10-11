# AutoSyncManagementCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Mode** | **NullableString** |  | 
**Ttl** | **NullableInt32** |  | 

## Methods

### NewAutoSyncManagementCommand

`func NewAutoSyncManagementCommand(id int32, mode NullableString, ttl NullableInt32, ) *AutoSyncManagementCommand`

NewAutoSyncManagementCommand instantiates a new AutoSyncManagementCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSyncManagementCommandWithDefaults

`func NewAutoSyncManagementCommandWithDefaults() *AutoSyncManagementCommand`

NewAutoSyncManagementCommandWithDefaults instantiates a new AutoSyncManagementCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoSyncManagementCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoSyncManagementCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoSyncManagementCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMode

`func (o *AutoSyncManagementCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AutoSyncManagementCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AutoSyncManagementCommand) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *AutoSyncManagementCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AutoSyncManagementCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetTtl

`func (o *AutoSyncManagementCommand) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AutoSyncManagementCommand) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AutoSyncManagementCommand) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *AutoSyncManagementCommand) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *AutoSyncManagementCommand) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


