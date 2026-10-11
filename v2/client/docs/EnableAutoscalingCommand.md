# EnableAutoscalingCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MinSize** | **int32** |  | 
**MaxSize** | **int32** |  | 
**DiskSize** | **float64** |  | 
**Flavor** | **NullableString** |  | 
**SpotEnabled** | **bool** |  | 

## Methods

### NewEnableAutoscalingCommand

`func NewEnableAutoscalingCommand(id int32, minSize int32, maxSize int32, diskSize float64, flavor NullableString, spotEnabled bool, ) *EnableAutoscalingCommand`

NewEnableAutoscalingCommand instantiates a new EnableAutoscalingCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableAutoscalingCommandWithDefaults

`func NewEnableAutoscalingCommandWithDefaults() *EnableAutoscalingCommand`

NewEnableAutoscalingCommandWithDefaults instantiates a new EnableAutoscalingCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnableAutoscalingCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnableAutoscalingCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnableAutoscalingCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMinSize

`func (o *EnableAutoscalingCommand) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *EnableAutoscalingCommand) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *EnableAutoscalingCommand) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.


### GetMaxSize

`func (o *EnableAutoscalingCommand) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *EnableAutoscalingCommand) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *EnableAutoscalingCommand) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### GetDiskSize

`func (o *EnableAutoscalingCommand) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *EnableAutoscalingCommand) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *EnableAutoscalingCommand) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.


### GetFlavor

`func (o *EnableAutoscalingCommand) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *EnableAutoscalingCommand) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *EnableAutoscalingCommand) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### SetFlavorNil

`func (o *EnableAutoscalingCommand) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *EnableAutoscalingCommand) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetSpotEnabled

`func (o *EnableAutoscalingCommand) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *EnableAutoscalingCommand) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *EnableAutoscalingCommand) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


