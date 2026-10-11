# OpenstackFlavorListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ram** | **float64** |  | 
**Cpu** | **int64** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**HasGpuSupport** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenstackFlavorListDto

`func NewOpenstackFlavorListDto(ram float64, cpu int64, name string, description NullableString, ) *OpenstackFlavorListDto`

NewOpenstackFlavorListDto instantiates a new OpenstackFlavorListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackFlavorListDtoWithDefaults

`func NewOpenstackFlavorListDtoWithDefaults() *OpenstackFlavorListDto`

NewOpenstackFlavorListDtoWithDefaults instantiates a new OpenstackFlavorListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRam

`func (o *OpenstackFlavorListDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *OpenstackFlavorListDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *OpenstackFlavorListDto) SetRam(v float64)`

SetRam sets Ram field to given value.


### GetCpu

`func (o *OpenstackFlavorListDto) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OpenstackFlavorListDto) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OpenstackFlavorListDto) SetCpu(v int64)`

SetCpu sets Cpu field to given value.


### GetName

`func (o *OpenstackFlavorListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenstackFlavorListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenstackFlavorListDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OpenstackFlavorListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenstackFlavorListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenstackFlavorListDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *OpenstackFlavorListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenstackFlavorListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHasGpuSupport

`func (o *OpenstackFlavorListDto) GetHasGpuSupport() bool`

GetHasGpuSupport returns the HasGpuSupport field if non-nil, zero value otherwise.

### GetHasGpuSupportOk

`func (o *OpenstackFlavorListDto) GetHasGpuSupportOk() (*bool, bool)`

GetHasGpuSupportOk returns a tuple with the HasGpuSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGpuSupport

`func (o *OpenstackFlavorListDto) SetHasGpuSupport(v bool)`

SetHasGpuSupport sets HasGpuSupport field to given value.

### HasHasGpuSupport

`func (o *OpenstackFlavorListDto) HasHasGpuSupport() bool`

HasHasGpuSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


