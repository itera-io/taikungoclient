# BindImageToProjectCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBindImageToProjectCommand

`func NewBindImageToProjectCommand() *BindImageToProjectCommand`

NewBindImageToProjectCommand instantiates a new BindImageToProjectCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindImageToProjectCommandWithDefaults

`func NewBindImageToProjectCommandWithDefaults() *BindImageToProjectCommand`

NewBindImageToProjectCommandWithDefaults instantiates a new BindImageToProjectCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *BindImageToProjectCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BindImageToProjectCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BindImageToProjectCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BindImageToProjectCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetImages

`func (o *BindImageToProjectCommand) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *BindImageToProjectCommand) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *BindImageToProjectCommand) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *BindImageToProjectCommand) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *BindImageToProjectCommand) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *BindImageToProjectCommand) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


