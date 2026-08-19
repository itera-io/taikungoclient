# ServersSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServersSearchResponseData**](ServersSearchResponseData.md) |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewServersSearchList

`func NewServersSearchList(data []ServersSearchResponseData, totalCount int64, ) *ServersSearchList`

NewServersSearchList instantiates a new ServersSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersSearchListWithDefaults

`func NewServersSearchListWithDefaults() *ServersSearchList`

NewServersSearchListWithDefaults instantiates a new ServersSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServersSearchList) GetData() []ServersSearchResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServersSearchList) GetDataOk() (*[]ServersSearchResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServersSearchList) SetData(v []ServersSearchResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ServersSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ServersSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ServersSearchList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServersSearchList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServersSearchList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


