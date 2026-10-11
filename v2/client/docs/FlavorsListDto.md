# FlavorsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ram** | **float64** |  | 
**Cpu** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **map[string]interface{}** |  | 
**MaxDataDiskCount** | **NullableFloat64** |  | 
**HasGpuSupport** | Pointer to **bool** |  | [optional] 
**GpuDetails** | Pointer to [**FlavorGpuDetailsDto**](FlavorGpuDetailsDto.md) |  | [optional] 

## Methods

### NewFlavorsListDto

`func NewFlavorsListDto(ram float64, cpu int32, name string, description map[string]interface{}, maxDataDiskCount NullableFloat64, ) *FlavorsListDto`

NewFlavorsListDto instantiates a new FlavorsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorsListDtoWithDefaults

`func NewFlavorsListDtoWithDefaults() *FlavorsListDto`

NewFlavorsListDtoWithDefaults instantiates a new FlavorsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRam

`func (o *FlavorsListDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *FlavorsListDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *FlavorsListDto) SetRam(v float64)`

SetRam sets Ram field to given value.


### GetCpu

`func (o *FlavorsListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *FlavorsListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *FlavorsListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetName

`func (o *FlavorsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FlavorsListDto) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlavorsListDto) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlavorsListDto) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.


### GetMaxDataDiskCount

`func (o *FlavorsListDto) GetMaxDataDiskCount() float64`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *FlavorsListDto) GetMaxDataDiskCountOk() (*float64, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *FlavorsListDto) SetMaxDataDiskCount(v float64)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.


### SetMaxDataDiskCountNil

`func (o *FlavorsListDto) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *FlavorsListDto) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
### GetHasGpuSupport

`func (o *FlavorsListDto) GetHasGpuSupport() bool`

GetHasGpuSupport returns the HasGpuSupport field if non-nil, zero value otherwise.

### GetHasGpuSupportOk

`func (o *FlavorsListDto) GetHasGpuSupportOk() (*bool, bool)`

GetHasGpuSupportOk returns a tuple with the HasGpuSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpuSupport

`func (o *FlavorsListDto) SetHasGpuSupport(v bool)`

SetHasGpuSupport sets HasGpuSupport field to given value.

### HasHasGpuSupport

`func (o *FlavorsListDto) HasHasGpuSupport() bool`

HasHasGpuSupport returns a boolean if a field has been set.

### GetGpuDetails

`func (o *FlavorsListDto) GetGpuDetails() FlavorGpuDetailsDto`

GetGpuDetails returns the GpuDetails field if non-nil, zero value otherwise.

### GetGpuDetailsOk

`func (o *FlavorsListDto) GetGpuDetailsOk() (*FlavorGpuDetailsDto, bool)`

GetGpuDetailsOk returns a tuple with the GpuDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDetails

`func (o *FlavorsListDto) SetGpuDetails(v FlavorGpuDetailsDto)`

SetGpuDetails sets GpuDetails field to given value.

### HasGpuDetails

`func (o *FlavorsListDto) HasGpuDetails() bool`

HasGpuDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


