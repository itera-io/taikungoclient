# SpotVmOperationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Mode** | **NullableString** |  | 

## Methods

### NewSpotVmOperationCommand

`func NewSpotVmOperationCommand(id int32, mode NullableString, ) *SpotVmOperationCommand`

NewSpotVmOperationCommand instantiates a new SpotVmOperationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotVmOperationCommandWithDefaults

`func NewSpotVmOperationCommandWithDefaults() *SpotVmOperationCommand`

NewSpotVmOperationCommandWithDefaults instantiates a new SpotVmOperationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpotVmOperationCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpotVmOperationCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpotVmOperationCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMode

`func (o *SpotVmOperationCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SpotVmOperationCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SpotVmOperationCommand) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *SpotVmOperationCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SpotVmOperationCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


