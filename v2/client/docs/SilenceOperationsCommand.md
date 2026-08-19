# SilenceOperationsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Mode** | **NullableString** |  | 
**Reason** | **NullableString** |  | 

## Methods

### NewSilenceOperationsCommand

`func NewSilenceOperationsCommand(id int32, mode NullableString, reason NullableString, ) *SilenceOperationsCommand`

NewSilenceOperationsCommand instantiates a new SilenceOperationsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSilenceOperationsCommandWithDefaults

`func NewSilenceOperationsCommandWithDefaults() *SilenceOperationsCommand`

NewSilenceOperationsCommandWithDefaults instantiates a new SilenceOperationsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SilenceOperationsCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SilenceOperationsCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SilenceOperationsCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMode

`func (o *SilenceOperationsCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SilenceOperationsCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SilenceOperationsCommand) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *SilenceOperationsCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SilenceOperationsCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetReason

`func (o *SilenceOperationsCommand) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SilenceOperationsCommand) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SilenceOperationsCommand) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *SilenceOperationsCommand) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SilenceOperationsCommand) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


