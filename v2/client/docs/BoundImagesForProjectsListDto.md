# BoundImagesForProjectsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**ProjectId** | **NullableInt32** |  | 
**ProjectName** | **NullableString** |  | 
**Size** | **NullableFloat64** |  | 
**ImageId** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 
**IsWindows** | **bool** |  | 
**CloudType** | **map[string]interface{}** |  | 

## Methods

### NewBoundImagesForProjectsListDto

`func NewBoundImagesForProjectsListDto(id int32, name NullableString, projectId NullableInt32, projectName NullableString, size NullableFloat64, imageId NullableString, cloudId NullableInt32, isWindows bool, cloudType map[string]interface{}, ) *BoundImagesForProjectsListDto`

NewBoundImagesForProjectsListDto instantiates a new BoundImagesForProjectsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundImagesForProjectsListDtoWithDefaults

`func NewBoundImagesForProjectsListDtoWithDefaults() *BoundImagesForProjectsListDto`

NewBoundImagesForProjectsListDtoWithDefaults instantiates a new BoundImagesForProjectsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoundImagesForProjectsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoundImagesForProjectsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoundImagesForProjectsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BoundImagesForProjectsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundImagesForProjectsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundImagesForProjectsListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BoundImagesForProjectsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BoundImagesForProjectsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *BoundImagesForProjectsListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BoundImagesForProjectsListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BoundImagesForProjectsListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *BoundImagesForProjectsListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BoundImagesForProjectsListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BoundImagesForProjectsListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BoundImagesForProjectsListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BoundImagesForProjectsListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *BoundImagesForProjectsListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BoundImagesForProjectsListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetSize

`func (o *BoundImagesForProjectsListDto) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BoundImagesForProjectsListDto) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BoundImagesForProjectsListDto) SetSize(v float64)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *BoundImagesForProjectsListDto) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BoundImagesForProjectsListDto) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetImageId

`func (o *BoundImagesForProjectsListDto) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *BoundImagesForProjectsListDto) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *BoundImagesForProjectsListDto) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### SetImageIdNil

`func (o *BoundImagesForProjectsListDto) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *BoundImagesForProjectsListDto) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetCloudId

`func (o *BoundImagesForProjectsListDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *BoundImagesForProjectsListDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *BoundImagesForProjectsListDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *BoundImagesForProjectsListDto) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *BoundImagesForProjectsListDto) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetIsWindows

`func (o *BoundImagesForProjectsListDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *BoundImagesForProjectsListDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *BoundImagesForProjectsListDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.


### GetCloudType

`func (o *BoundImagesForProjectsListDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *BoundImagesForProjectsListDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *BoundImagesForProjectsListDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


