# VirtualClusterQuotasDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Presets** | [**[]ResourcePresetDto**](ResourcePresetDto.md) |  | 

## Methods

### NewVirtualClusterQuotasDto

`func NewVirtualClusterQuotasDto(presets []ResourcePresetDto, ) *VirtualClusterQuotasDto`

NewVirtualClusterQuotasDto instantiates a new VirtualClusterQuotasDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualClusterQuotasDtoWithDefaults

`func NewVirtualClusterQuotasDtoWithDefaults() *VirtualClusterQuotasDto`

NewVirtualClusterQuotasDtoWithDefaults instantiates a new VirtualClusterQuotasDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresets

`func (o *VirtualClusterQuotasDto) GetPresets() []ResourcePresetDto`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *VirtualClusterQuotasDto) GetPresetsOk() (*[]ResourcePresetDto, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *VirtualClusterQuotasDto) SetPresets(v []ResourcePresetDto)`

SetPresets sets Presets field to given value.


### SetPresetsNil

`func (o *VirtualClusterQuotasDto) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *VirtualClusterQuotasDto) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


