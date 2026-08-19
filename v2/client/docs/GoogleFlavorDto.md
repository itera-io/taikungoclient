# GoogleFlavorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cpu** | **NullableInt32** |  | 
**Ram** | **NullableFloat64** |  | 
**LinuxPrice** | **NullableFloat64** |  | 
**WindowsPrice** | **NullableFloat64** |  | 
**LinuxSpotPrice** | **NullableFloat64** |  | 
**WindowsSpotPrice** | **NullableFloat64** |  | 
**HasGpuSupport** | Pointer to **bool** |  | [optional] 
**Description** | **map[string]interface{}** |  | 
**GpuDetails** | Pointer to [**FlavorGpuDetailsDto**](FlavorGpuDetailsDto.md) |  | [optional] 

## Methods

### NewGoogleFlavorDto

`func NewGoogleFlavorDto(name string, cpu NullableInt32, ram NullableFloat64, linuxPrice NullableFloat64, windowsPrice NullableFloat64, linuxSpotPrice NullableFloat64, windowsSpotPrice NullableFloat64, description map[string]interface{}, ) *GoogleFlavorDto`

NewGoogleFlavorDto instantiates a new GoogleFlavorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleFlavorDtoWithDefaults

`func NewGoogleFlavorDtoWithDefaults() *GoogleFlavorDto`

NewGoogleFlavorDtoWithDefaults instantiates a new GoogleFlavorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GoogleFlavorDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleFlavorDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleFlavorDto) SetName(v string)`

SetName sets Name field to given value.


### GetCpu

`func (o *GoogleFlavorDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GoogleFlavorDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GoogleFlavorDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### SetCpuNil

`func (o *GoogleFlavorDto) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *GoogleFlavorDto) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetRam

`func (o *GoogleFlavorDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *GoogleFlavorDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *GoogleFlavorDto) SetRam(v float64)`

SetRam sets Ram field to given value.


### SetRamNil

`func (o *GoogleFlavorDto) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *GoogleFlavorDto) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetLinuxPrice

`func (o *GoogleFlavorDto) GetLinuxPrice() float64`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *GoogleFlavorDto) GetLinuxPriceOk() (*float64, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *GoogleFlavorDto) SetLinuxPrice(v float64)`

SetLinuxPrice sets LinuxPrice field to given value.


### SetLinuxPriceNil

`func (o *GoogleFlavorDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *GoogleFlavorDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsPrice

`func (o *GoogleFlavorDto) GetWindowsPrice() float64`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *GoogleFlavorDto) GetWindowsPriceOk() (*float64, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *GoogleFlavorDto) SetWindowsPrice(v float64)`

SetWindowsPrice sets WindowsPrice field to given value.


### SetWindowsPriceNil

`func (o *GoogleFlavorDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *GoogleFlavorDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
### GetLinuxSpotPrice

`func (o *GoogleFlavorDto) GetLinuxSpotPrice() float64`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *GoogleFlavorDto) GetLinuxSpotPriceOk() (*float64, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *GoogleFlavorDto) SetLinuxSpotPrice(v float64)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.


### SetLinuxSpotPriceNil

`func (o *GoogleFlavorDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *GoogleFlavorDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *GoogleFlavorDto) GetWindowsSpotPrice() float64`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *GoogleFlavorDto) GetWindowsSpotPriceOk() (*float64, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *GoogleFlavorDto) SetWindowsSpotPrice(v float64)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.


### SetWindowsSpotPriceNil

`func (o *GoogleFlavorDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *GoogleFlavorDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
### GetHasGpuSupport

`func (o *GoogleFlavorDto) GetHasGpuSupport() bool`

GetHasGpuSupport returns the HasGpuSupport field if non-nil, zero value otherwise.

### GetHasGpuSupportOk

`func (o *GoogleFlavorDto) GetHasGpuSupportOk() (*bool, bool)`

GetHasGpuSupportOk returns a tuple with the HasGpuSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpuSupport

`func (o *GoogleFlavorDto) SetHasGpuSupport(v bool)`

SetHasGpuSupport sets HasGpuSupport field to given value.

### HasHasGpuSupport

`func (o *GoogleFlavorDto) HasHasGpuSupport() bool`

HasHasGpuSupport returns a boolean if a field has been set.

### GetDescription

`func (o *GoogleFlavorDto) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoogleFlavorDto) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoogleFlavorDto) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.


### GetGpuDetails

`func (o *GoogleFlavorDto) GetGpuDetails() FlavorGpuDetailsDto`

GetGpuDetails returns the GpuDetails field if non-nil, zero value otherwise.

### GetGpuDetailsOk

`func (o *GoogleFlavorDto) GetGpuDetailsOk() (*FlavorGpuDetailsDto, bool)`

GetGpuDetailsOk returns a tuple with the GpuDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDetails

`func (o *GoogleFlavorDto) SetGpuDetails(v FlavorGpuDetailsDto)`

SetGpuDetails sets GpuDetails field to given value.

### HasGpuDetails

`func (o *GoogleFlavorDto) HasGpuDetails() bool`

HasGpuDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


