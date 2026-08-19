# AwsImagesPostListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **int32** |  | 
**Limit** | **NullableInt32** |  | 
**Offset** | **NullableInt32** |  | 
**SortBy** | **NullableString** |  | 
**SortDirection** | **NullableString** |  | 
**Search** | **NullableString** |  | 
**Latest** | **bool** |  | 
**Owners** | **[]string** |  | 
**ProjectId** | **NullableInt32** |  | 

## Methods

### NewAwsImagesPostListCommand

`func NewAwsImagesPostListCommand(cloudId int32, limit NullableInt32, offset NullableInt32, sortBy NullableString, sortDirection NullableString, search NullableString, latest bool, owners []string, projectId NullableInt32, ) *AwsImagesPostListCommand`

NewAwsImagesPostListCommand instantiates a new AwsImagesPostListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsImagesPostListCommandWithDefaults

`func NewAwsImagesPostListCommandWithDefaults() *AwsImagesPostListCommand`

NewAwsImagesPostListCommandWithDefaults instantiates a new AwsImagesPostListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *AwsImagesPostListCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AwsImagesPostListCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AwsImagesPostListCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetLimit

`func (o *AwsImagesPostListCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AwsImagesPostListCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AwsImagesPostListCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *AwsImagesPostListCommand) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AwsImagesPostListCommand) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *AwsImagesPostListCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsImagesPostListCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsImagesPostListCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### SetOffsetNil

`func (o *AwsImagesPostListCommand) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *AwsImagesPostListCommand) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSortBy

`func (o *AwsImagesPostListCommand) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *AwsImagesPostListCommand) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *AwsImagesPostListCommand) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.


### SetSortByNil

`func (o *AwsImagesPostListCommand) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *AwsImagesPostListCommand) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortDirection

`func (o *AwsImagesPostListCommand) GetSortDirection() string`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *AwsImagesPostListCommand) GetSortDirectionOk() (*string, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *AwsImagesPostListCommand) SetSortDirection(v string)`

SetSortDirection sets SortDirection field to given value.


### SetSortDirectionNil

`func (o *AwsImagesPostListCommand) SetSortDirectionNil(b bool)`

 SetSortDirectionNil sets the value for SortDirection to be an explicit nil

### UnsetSortDirection
`func (o *AwsImagesPostListCommand) UnsetSortDirection()`

UnsetSortDirection ensures that no value is present for SortDirection, not even an explicit nil
### GetSearch

`func (o *AwsImagesPostListCommand) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AwsImagesPostListCommand) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AwsImagesPostListCommand) SetSearch(v string)`

SetSearch sets Search field to given value.


### SetSearchNil

`func (o *AwsImagesPostListCommand) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *AwsImagesPostListCommand) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetLatest

`func (o *AwsImagesPostListCommand) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *AwsImagesPostListCommand) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *AwsImagesPostListCommand) SetLatest(v bool)`

SetLatest sets Latest field to given value.


### GetOwners

`func (o *AwsImagesPostListCommand) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AwsImagesPostListCommand) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AwsImagesPostListCommand) SetOwners(v []string)`

SetOwners sets Owners field to given value.


### SetOwnersNil

`func (o *AwsImagesPostListCommand) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *AwsImagesPostListCommand) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil
### GetProjectId

`func (o *AwsImagesPostListCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AwsImagesPostListCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AwsImagesPostListCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *AwsImagesPostListCommand) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AwsImagesPostListCommand) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


