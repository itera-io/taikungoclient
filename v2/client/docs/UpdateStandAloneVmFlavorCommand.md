# UpdateStandAloneVmFlavorCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Flavor** | **NullableString** |  | 

## Methods

### NewUpdateStandAloneVmFlavorCommand

`func NewUpdateStandAloneVmFlavorCommand(id int32, flavor NullableString, ) *UpdateStandAloneVmFlavorCommand`

NewUpdateStandAloneVmFlavorCommand instantiates a new UpdateStandAloneVmFlavorCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStandAloneVmFlavorCommandWithDefaults

`func NewUpdateStandAloneVmFlavorCommandWithDefaults() *UpdateStandAloneVmFlavorCommand`

NewUpdateStandAloneVmFlavorCommandWithDefaults instantiates a new UpdateStandAloneVmFlavorCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStandAloneVmFlavorCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStandAloneVmFlavorCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStandAloneVmFlavorCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetFlavor

`func (o *UpdateStandAloneVmFlavorCommand) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *UpdateStandAloneVmFlavorCommand) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *UpdateStandAloneVmFlavorCommand) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### SetFlavorNil

`func (o *UpdateStandAloneVmFlavorCommand) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *UpdateStandAloneVmFlavorCommand) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


