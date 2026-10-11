# FlavorGpuFilterMetadataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** |  | 
**GpuFilterMetadata** | Pointer to [**FlavorGpuFilterMetadataDto**](FlavorGpuFilterMetadataDto.md) |  | [optional] 

## Methods

### NewFlavorGpuFilterMetadataList

`func NewFlavorGpuFilterMetadataList(cloudType string, ) *FlavorGpuFilterMetadataList`

NewFlavorGpuFilterMetadataList instantiates a new FlavorGpuFilterMetadataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorGpuFilterMetadataListWithDefaults

`func NewFlavorGpuFilterMetadataListWithDefaults() *FlavorGpuFilterMetadataList`

NewFlavorGpuFilterMetadataListWithDefaults instantiates a new FlavorGpuFilterMetadataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *FlavorGpuFilterMetadataList) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *FlavorGpuFilterMetadataList) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *FlavorGpuFilterMetadataList) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### GetGpuFilterMetadata

`func (o *FlavorGpuFilterMetadataList) GetGpuFilterMetadata() FlavorGpuFilterMetadataDto`

GetGpuFilterMetadata returns the GpuFilterMetadata field if non-nil, zero value otherwise.

### GetGpuFilterMetadataOk

`func (o *FlavorGpuFilterMetadataList) GetGpuFilterMetadataOk() (*FlavorGpuFilterMetadataDto, bool)`

GetGpuFilterMetadataOk returns a tuple with the GpuFilterMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuFilterMetadata

`func (o *FlavorGpuFilterMetadataList) SetGpuFilterMetadata(v FlavorGpuFilterMetadataDto)`

SetGpuFilterMetadata sets GpuFilterMetadata field to given value.

### HasGpuFilterMetadata

`func (o *FlavorGpuFilterMetadataList) HasGpuFilterMetadata() bool`

HasGpuFilterMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


