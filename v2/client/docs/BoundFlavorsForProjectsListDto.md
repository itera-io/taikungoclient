# BoundFlavorsForProjectsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Cpu** | **int32** |  | 
**Ram** | **float64** |  | 
**ProjectId** | **NullableInt32** |  | 
**ProjectName** | **string** |  | 
**MaxDataDiskCount** | **NullableInt32** |  | 
**ExistsLinuxSpotPrice** | **bool** |  | 
**ExistsWindowsSpotPrice** | **bool** |  | 
**LinuxSpotPrice** | **NullableString** |  | 
**LinuxPrice** | **NullableString** |  | 
**WindowsSpotPrice** | **NullableString** |  | 
**WindowsPrice** | **NullableString** |  | 
**HasGpuSupport** | Pointer to **bool** |  | [optional] 
**GpuDetails** | Pointer to [**FlavorGpuDetailsDto**](FlavorGpuDetailsDto.md) |  | [optional] 
**CloudType** | **map[string]interface{}** |  | 
**LocalDiskSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewBoundFlavorsForProjectsListDto

`func NewBoundFlavorsForProjectsListDto(id int32, name string, cpu int32, ram float64, projectId NullableInt32, projectName string, maxDataDiskCount NullableInt32, existsLinuxSpotPrice bool, existsWindowsSpotPrice bool, linuxSpotPrice NullableString, linuxPrice NullableString, windowsSpotPrice NullableString, windowsPrice NullableString, cloudType map[string]interface{}, ) *BoundFlavorsForProjectsListDto`

NewBoundFlavorsForProjectsListDto instantiates a new BoundFlavorsForProjectsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundFlavorsForProjectsListDtoWithDefaults

`func NewBoundFlavorsForProjectsListDtoWithDefaults() *BoundFlavorsForProjectsListDto`

NewBoundFlavorsForProjectsListDtoWithDefaults instantiates a new BoundFlavorsForProjectsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoundFlavorsForProjectsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoundFlavorsForProjectsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoundFlavorsForProjectsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BoundFlavorsForProjectsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundFlavorsForProjectsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundFlavorsForProjectsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetCpu

`func (o *BoundFlavorsForProjectsListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BoundFlavorsForProjectsListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BoundFlavorsForProjectsListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *BoundFlavorsForProjectsListDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *BoundFlavorsForProjectsListDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *BoundFlavorsForProjectsListDto) SetRam(v float64)`

SetRam sets Ram field to given value.


### GetProjectId

`func (o *BoundFlavorsForProjectsListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BoundFlavorsForProjectsListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BoundFlavorsForProjectsListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *BoundFlavorsForProjectsListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BoundFlavorsForProjectsListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BoundFlavorsForProjectsListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BoundFlavorsForProjectsListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BoundFlavorsForProjectsListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetMaxDataDiskCount

`func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCount() int32`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCountOk() (*int32, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCount(v int32)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.


### SetMaxDataDiskCountNil

`func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *BoundFlavorsForProjectsListDto) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
### GetExistsLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPrice() bool`

GetExistsLinuxSpotPrice returns the ExistsLinuxSpotPrice field if non-nil, zero value otherwise.

### GetExistsLinuxSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPriceOk() (*bool, bool)`

GetExistsLinuxSpotPriceOk returns a tuple with the ExistsLinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetExistsLinuxSpotPrice(v bool)`

SetExistsLinuxSpotPrice sets ExistsLinuxSpotPrice field to given value.


### GetExistsWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPrice() bool`

GetExistsWindowsSpotPrice returns the ExistsWindowsSpotPrice field if non-nil, zero value otherwise.

### GetExistsWindowsSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPriceOk() (*bool, bool)`

GetExistsWindowsSpotPriceOk returns a tuple with the ExistsWindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetExistsWindowsSpotPrice(v bool)`

SetExistsWindowsSpotPrice sets ExistsWindowsSpotPrice field to given value.


### GetLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPrice() string`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPriceOk() (*string, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPrice(v string)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.


### SetLinuxSpotPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
### GetLinuxPrice

`func (o *BoundFlavorsForProjectsListDto) GetLinuxPrice() string`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetLinuxPriceOk() (*string, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *BoundFlavorsForProjectsListDto) SetLinuxPrice(v string)`

SetLinuxPrice sets LinuxPrice field to given value.


### SetLinuxPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPrice() string`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPriceOk() (*string, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPrice(v string)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.


### SetWindowsSpotPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
### GetWindowsPrice

`func (o *BoundFlavorsForProjectsListDto) GetWindowsPrice() string`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *BoundFlavorsForProjectsListDto) GetWindowsPriceOk() (*string, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *BoundFlavorsForProjectsListDto) SetWindowsPrice(v string)`

SetWindowsPrice sets WindowsPrice field to given value.


### SetWindowsPriceNil

`func (o *BoundFlavorsForProjectsListDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *BoundFlavorsForProjectsListDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
### GetHasGpuSupport

`func (o *BoundFlavorsForProjectsListDto) GetHasGpuSupport() bool`

GetHasGpuSupport returns the HasGpuSupport field if non-nil, zero value otherwise.

### GetHasGpuSupportOk

`func (o *BoundFlavorsForProjectsListDto) GetHasGpuSupportOk() (*bool, bool)`

GetHasGpuSupportOk returns a tuple with the HasGpuSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpuSupport

`func (o *BoundFlavorsForProjectsListDto) SetHasGpuSupport(v bool)`

SetHasGpuSupport sets HasGpuSupport field to given value.

### HasHasGpuSupport

`func (o *BoundFlavorsForProjectsListDto) HasHasGpuSupport() bool`

HasHasGpuSupport returns a boolean if a field has been set.

### GetGpuDetails

`func (o *BoundFlavorsForProjectsListDto) GetGpuDetails() FlavorGpuDetailsDto`

GetGpuDetails returns the GpuDetails field if non-nil, zero value otherwise.

### GetGpuDetailsOk

`func (o *BoundFlavorsForProjectsListDto) GetGpuDetailsOk() (*FlavorGpuDetailsDto, bool)`

GetGpuDetailsOk returns a tuple with the GpuDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDetails

`func (o *BoundFlavorsForProjectsListDto) SetGpuDetails(v FlavorGpuDetailsDto)`

SetGpuDetails sets GpuDetails field to given value.

### HasGpuDetails

`func (o *BoundFlavorsForProjectsListDto) HasGpuDetails() bool`

HasGpuDetails returns a boolean if a field has been set.

### GetCloudType

`func (o *BoundFlavorsForProjectsListDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *BoundFlavorsForProjectsListDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *BoundFlavorsForProjectsListDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetLocalDiskSize

`func (o *BoundFlavorsForProjectsListDto) GetLocalDiskSize() int32`

GetLocalDiskSize returns the LocalDiskSize field if non-nil, zero value otherwise.

### GetLocalDiskSizeOk

`func (o *BoundFlavorsForProjectsListDto) GetLocalDiskSizeOk() (*int32, bool)`

GetLocalDiskSizeOk returns a tuple with the LocalDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiskSize

`func (o *BoundFlavorsForProjectsListDto) SetLocalDiskSize(v int32)`

SetLocalDiskSize sets LocalDiskSize field to given value.

### HasLocalDiskSize

`func (o *BoundFlavorsForProjectsListDto) HasLocalDiskSize() bool`

HasLocalDiskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


