# AzureFlavorsWithPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**WindowsPrice** | Pointer to **NullableString** |  | [optional] 
**LinuxPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsSpotPrice** | Pointer to **NullableString** |  | [optional] 
**LinuxSpotPrice** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Ram** | Pointer to **float64** |  | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxDataDiskCount** | Pointer to **NullableFloat64** |  | [optional] 
**HasGpuSupport** | Pointer to **bool** |  | [optional] 
**GpuDetails** | Pointer to [**FlavorGpuDetailsDto**](FlavorGpuDetailsDto.md) |  | [optional] 

## Methods

### NewAzureFlavorsWithPriceDto

`func NewAzureFlavorsWithPriceDto() *AzureFlavorsWithPriceDto`

NewAzureFlavorsWithPriceDto instantiates a new AzureFlavorsWithPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureFlavorsWithPriceDtoWithDefaults

`func NewAzureFlavorsWithPriceDtoWithDefaults() *AzureFlavorsWithPriceDto`

NewAzureFlavorsWithPriceDtoWithDefaults instantiates a new AzureFlavorsWithPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureFlavorsWithPriceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureFlavorsWithPriceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureFlavorsWithPriceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureFlavorsWithPriceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureFlavorsWithPriceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureFlavorsWithPriceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWindowsPrice

`func (o *AzureFlavorsWithPriceDto) GetWindowsPrice() string`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *AzureFlavorsWithPriceDto) GetWindowsPriceOk() (*string, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *AzureFlavorsWithPriceDto) SetWindowsPrice(v string)`

SetWindowsPrice sets WindowsPrice field to given value.

### HasWindowsPrice

`func (o *AzureFlavorsWithPriceDto) HasWindowsPrice() bool`

HasWindowsPrice returns a boolean if a field has been set.

### SetWindowsPriceNil

`func (o *AzureFlavorsWithPriceDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *AzureFlavorsWithPriceDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
### GetLinuxPrice

`func (o *AzureFlavorsWithPriceDto) GetLinuxPrice() string`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *AzureFlavorsWithPriceDto) GetLinuxPriceOk() (*string, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *AzureFlavorsWithPriceDto) SetLinuxPrice(v string)`

SetLinuxPrice sets LinuxPrice field to given value.

### HasLinuxPrice

`func (o *AzureFlavorsWithPriceDto) HasLinuxPrice() bool`

HasLinuxPrice returns a boolean if a field has been set.

### SetLinuxPriceNil

`func (o *AzureFlavorsWithPriceDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *AzureFlavorsWithPriceDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *AzureFlavorsWithPriceDto) GetWindowsSpotPrice() string`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *AzureFlavorsWithPriceDto) GetWindowsSpotPriceOk() (*string, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *AzureFlavorsWithPriceDto) SetWindowsSpotPrice(v string)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.

### HasWindowsSpotPrice

`func (o *AzureFlavorsWithPriceDto) HasWindowsSpotPrice() bool`

HasWindowsSpotPrice returns a boolean if a field has been set.

### SetWindowsSpotPriceNil

`func (o *AzureFlavorsWithPriceDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *AzureFlavorsWithPriceDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
### GetLinuxSpotPrice

`func (o *AzureFlavorsWithPriceDto) GetLinuxSpotPrice() string`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *AzureFlavorsWithPriceDto) GetLinuxSpotPriceOk() (*string, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *AzureFlavorsWithPriceDto) SetLinuxSpotPrice(v string)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.

### HasLinuxSpotPrice

`func (o *AzureFlavorsWithPriceDto) HasLinuxSpotPrice() bool`

HasLinuxSpotPrice returns a boolean if a field has been set.

### SetLinuxSpotPriceNil

`func (o *AzureFlavorsWithPriceDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *AzureFlavorsWithPriceDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
### GetCpu

`func (o *AzureFlavorsWithPriceDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AzureFlavorsWithPriceDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AzureFlavorsWithPriceDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AzureFlavorsWithPriceDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *AzureFlavorsWithPriceDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AzureFlavorsWithPriceDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AzureFlavorsWithPriceDto) SetRam(v float64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *AzureFlavorsWithPriceDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDescription

`func (o *AzureFlavorsWithPriceDto) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureFlavorsWithPriceDto) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureFlavorsWithPriceDto) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureFlavorsWithPriceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxDataDiskCount

`func (o *AzureFlavorsWithPriceDto) GetMaxDataDiskCount() float64`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *AzureFlavorsWithPriceDto) GetMaxDataDiskCountOk() (*float64, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *AzureFlavorsWithPriceDto) SetMaxDataDiskCount(v float64)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.

### HasMaxDataDiskCount

`func (o *AzureFlavorsWithPriceDto) HasMaxDataDiskCount() bool`

HasMaxDataDiskCount returns a boolean if a field has been set.

### SetMaxDataDiskCountNil

`func (o *AzureFlavorsWithPriceDto) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *AzureFlavorsWithPriceDto) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
### GetHasGpuSupport

`func (o *AzureFlavorsWithPriceDto) GetHasGpuSupport() bool`

GetHasGpuSupport returns the HasGpuSupport field if non-nil, zero value otherwise.

### GetHasGpuSupportOk

`func (o *AzureFlavorsWithPriceDto) GetHasGpuSupportOk() (*bool, bool)`

GetHasGpuSupportOk returns a tuple with the HasGpuSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpuSupport

`func (o *AzureFlavorsWithPriceDto) SetHasGpuSupport(v bool)`

SetHasGpuSupport sets HasGpuSupport field to given value.

### HasHasGpuSupport

`func (o *AzureFlavorsWithPriceDto) HasHasGpuSupport() bool`

HasHasGpuSupport returns a boolean if a field has been set.

### GetGpuDetails

`func (o *AzureFlavorsWithPriceDto) GetGpuDetails() FlavorGpuDetailsDto`

GetGpuDetails returns the GpuDetails field if non-nil, zero value otherwise.

### GetGpuDetailsOk

`func (o *AzureFlavorsWithPriceDto) GetGpuDetailsOk() (*FlavorGpuDetailsDto, bool)`

GetGpuDetailsOk returns a tuple with the GpuDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDetails

`func (o *AzureFlavorsWithPriceDto) SetGpuDetails(v FlavorGpuDetailsDto)`

SetGpuDetails sets GpuDetails field to given value.

### HasGpuDetails

`func (o *AzureFlavorsWithPriceDto) HasGpuDetails() bool`

HasGpuDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


