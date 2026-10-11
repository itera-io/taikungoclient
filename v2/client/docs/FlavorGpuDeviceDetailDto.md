# FlavorGpuDeviceDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceleratorType** | **string** |  | 
**Vendor** | Pointer to **NullableString** |  | [optional] 
**DeviceCount** | **int32** |  | 
**MemoryMiBPerUnit** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFlavorGpuDeviceDetailDto

`func NewFlavorGpuDeviceDetailDto(acceleratorType string, deviceCount int32, ) *FlavorGpuDeviceDetailDto`

NewFlavorGpuDeviceDetailDto instantiates a new FlavorGpuDeviceDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorGpuDeviceDetailDtoWithDefaults

`func NewFlavorGpuDeviceDetailDtoWithDefaults() *FlavorGpuDeviceDetailDto`

NewFlavorGpuDeviceDetailDtoWithDefaults instantiates a new FlavorGpuDeviceDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceleratorType

`func (o *FlavorGpuDeviceDetailDto) GetAcceleratorType() string`

GetAcceleratorType returns the AcceleratorType field if non-nil, zero value otherwise.

### GetAcceleratorTypeOk

`func (o *FlavorGpuDeviceDetailDto) GetAcceleratorTypeOk() (*string, bool)`

GetAcceleratorTypeOk returns a tuple with the AcceleratorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorType

`func (o *FlavorGpuDeviceDetailDto) SetAcceleratorType(v string)`

SetAcceleratorType sets AcceleratorType field to given value.


### GetVendor

`func (o *FlavorGpuDeviceDetailDto) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FlavorGpuDeviceDetailDto) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FlavorGpuDeviceDetailDto) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FlavorGpuDeviceDetailDto) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *FlavorGpuDeviceDetailDto) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *FlavorGpuDeviceDetailDto) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetDeviceCount

`func (o *FlavorGpuDeviceDetailDto) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *FlavorGpuDeviceDetailDto) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *FlavorGpuDeviceDetailDto) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.


### GetMemoryMiBPerUnit

`func (o *FlavorGpuDeviceDetailDto) GetMemoryMiBPerUnit() int32`

GetMemoryMiBPerUnit returns the MemoryMiBPerUnit field if non-nil, zero value otherwise.

### GetMemoryMiBPerUnitOk

`func (o *FlavorGpuDeviceDetailDto) GetMemoryMiBPerUnitOk() (*int32, bool)`

GetMemoryMiBPerUnitOk returns a tuple with the MemoryMiBPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiBPerUnit

`func (o *FlavorGpuDeviceDetailDto) SetMemoryMiBPerUnit(v int32)`

SetMemoryMiBPerUnit sets MemoryMiBPerUnit field to given value.

### HasMemoryMiBPerUnit

`func (o *FlavorGpuDeviceDetailDto) HasMemoryMiBPerUnit() bool`

HasMemoryMiBPerUnit returns a boolean if a field has been set.

### SetMemoryMiBPerUnitNil

`func (o *FlavorGpuDeviceDetailDto) SetMemoryMiBPerUnitNil(b bool)`

 SetMemoryMiBPerUnitNil sets the value for MemoryMiBPerUnit to be an explicit nil

### UnsetMemoryMiBPerUnit
`func (o *FlavorGpuDeviceDetailDto) UnsetMemoryMiBPerUnit()`

UnsetMemoryMiBPerUnit ensures that no value is present for MemoryMiBPerUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


