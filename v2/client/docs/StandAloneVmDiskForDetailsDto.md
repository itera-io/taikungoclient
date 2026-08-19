# StandAloneVmDiskForDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**CurrentSize** | **int64** |  | 
**TargetSize** | **int64** |  | 
**VolumeType** | **NullableString** |  | 
**DeviceName** | **NullableString** |  | 
**LunId** | **NullableString** |  | 
**Status** | [**StandAloneVmDiskStatus**](StandAloneVmDiskStatus.md) |  | 

## Methods

### NewStandAloneVmDiskForDetailsDto

`func NewStandAloneVmDiskForDetailsDto(id int32, name NullableString, currentSize int64, targetSize int64, volumeType NullableString, deviceName NullableString, lunId NullableString, status StandAloneVmDiskStatus, ) *StandAloneVmDiskForDetailsDto`

NewStandAloneVmDiskForDetailsDto instantiates a new StandAloneVmDiskForDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneVmDiskForDetailsDtoWithDefaults

`func NewStandAloneVmDiskForDetailsDtoWithDefaults() *StandAloneVmDiskForDetailsDto`

NewStandAloneVmDiskForDetailsDtoWithDefaults instantiates a new StandAloneVmDiskForDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneVmDiskForDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneVmDiskForDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneVmDiskForDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandAloneVmDiskForDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneVmDiskForDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneVmDiskForDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StandAloneVmDiskForDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneVmDiskForDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrentSize

`func (o *StandAloneVmDiskForDetailsDto) GetCurrentSize() int64`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *StandAloneVmDiskForDetailsDto) GetCurrentSizeOk() (*int64, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *StandAloneVmDiskForDetailsDto) SetCurrentSize(v int64)`

SetCurrentSize sets CurrentSize field to given value.


### GetTargetSize

`func (o *StandAloneVmDiskForDetailsDto) GetTargetSize() int64`

GetTargetSize returns the TargetSize field if non-nil, zero value otherwise.

### GetTargetSizeOk

`func (o *StandAloneVmDiskForDetailsDto) GetTargetSizeOk() (*int64, bool)`

GetTargetSizeOk returns a tuple with the TargetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSize

`func (o *StandAloneVmDiskForDetailsDto) SetTargetSize(v int64)`

SetTargetSize sets TargetSize field to given value.


### GetVolumeType

`func (o *StandAloneVmDiskForDetailsDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandAloneVmDiskForDetailsDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandAloneVmDiskForDetailsDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### SetVolumeTypeNil

`func (o *StandAloneVmDiskForDetailsDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *StandAloneVmDiskForDetailsDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetDeviceName

`func (o *StandAloneVmDiskForDetailsDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StandAloneVmDiskForDetailsDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StandAloneVmDiskForDetailsDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### SetDeviceNameNil

`func (o *StandAloneVmDiskForDetailsDto) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *StandAloneVmDiskForDetailsDto) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetLunId

`func (o *StandAloneVmDiskForDetailsDto) GetLunId() string`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *StandAloneVmDiskForDetailsDto) GetLunIdOk() (*string, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *StandAloneVmDiskForDetailsDto) SetLunId(v string)`

SetLunId sets LunId field to given value.


### SetLunIdNil

`func (o *StandAloneVmDiskForDetailsDto) SetLunIdNil(b bool)`

 SetLunIdNil sets the value for LunId to be an explicit nil

### UnsetLunId
`func (o *StandAloneVmDiskForDetailsDto) UnsetLunId()`

UnsetLunId ensures that no value is present for LunId, not even an explicit nil
### GetStatus

`func (o *StandAloneVmDiskForDetailsDto) GetStatus() StandAloneVmDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandAloneVmDiskForDetailsDto) GetStatusOk() (*StandAloneVmDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandAloneVmDiskForDetailsDto) SetStatus(v StandAloneVmDiskStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


