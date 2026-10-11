# AppRepositoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArtifactRepositoryDto**](ArtifactRepositoryDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAppRepositoryList

`func NewAppRepositoryList(data []ArtifactRepositoryDto, totalCount int32, ) *AppRepositoryList`

NewAppRepositoryList instantiates a new AppRepositoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRepositoryListWithDefaults

`func NewAppRepositoryListWithDefaults() *AppRepositoryList`

NewAppRepositoryListWithDefaults instantiates a new AppRepositoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppRepositoryList) GetData() []ArtifactRepositoryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppRepositoryList) GetDataOk() (*[]ArtifactRepositoryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppRepositoryList) SetData(v []ArtifactRepositoryDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AppRepositoryList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AppRepositoryList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AppRepositoryList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AppRepositoryList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AppRepositoryList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


