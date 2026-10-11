# ImageByIdCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **int32** |  | 
**ImageId** | **NullableString** |  | 

## Methods

### NewImageByIdCommand

`func NewImageByIdCommand(cloudId int32, imageId NullableString, ) *ImageByIdCommand`

NewImageByIdCommand instantiates a new ImageByIdCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageByIdCommandWithDefaults

`func NewImageByIdCommandWithDefaults() *ImageByIdCommand`

NewImageByIdCommandWithDefaults instantiates a new ImageByIdCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *ImageByIdCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ImageByIdCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ImageByIdCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetImageId

`func (o *ImageByIdCommand) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageByIdCommand) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageByIdCommand) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### SetImageIdNil

`func (o *ImageByIdCommand) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *ImageByIdCommand) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


