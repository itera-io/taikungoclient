# FlavorGpuDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuDeviceTotal** | **int32** |  | 
**Devices** | [**[]FlavorGpuDeviceDetailDto**](FlavorGpuDeviceDetailDto.md) |  | 

## Methods

### NewFlavorGpuDetailsDto

`func NewFlavorGpuDetailsDto(gpuDeviceTotal int32, devices []FlavorGpuDeviceDetailDto, ) *FlavorGpuDetailsDto`

NewFlavorGpuDetailsDto instantiates a new FlavorGpuDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorGpuDetailsDtoWithDefaults

`func NewFlavorGpuDetailsDtoWithDefaults() *FlavorGpuDetailsDto`

NewFlavorGpuDetailsDtoWithDefaults instantiates a new FlavorGpuDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuDeviceTotal

`func (o *FlavorGpuDetailsDto) GetGpuDeviceTotal() int32`

GetGpuDeviceTotal returns the GpuDeviceTotal field if non-nil, zero value otherwise.

### GetGpuDeviceTotalOk

`func (o *FlavorGpuDetailsDto) GetGpuDeviceTotalOk() (*int32, bool)`

GetGpuDeviceTotalOk returns a tuple with the GpuDeviceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDeviceTotal

`func (o *FlavorGpuDetailsDto) SetGpuDeviceTotal(v int32)`

SetGpuDeviceTotal sets GpuDeviceTotal field to given value.


### GetDevices

`func (o *FlavorGpuDetailsDto) GetDevices() []FlavorGpuDeviceDetailDto`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *FlavorGpuDetailsDto) GetDevicesOk() (*[]FlavorGpuDeviceDetailDto, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *FlavorGpuDetailsDto) SetDevices(v []FlavorGpuDeviceDetailDto)`

SetDevices sets Devices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


