# ExceededQuotaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExceededQuotaDto**](ExceededQuotaDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewExceededQuotaList

`func NewExceededQuotaList(data []ExceededQuotaDto, totalCount int32, ) *ExceededQuotaList`

NewExceededQuotaList instantiates a new ExceededQuotaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceededQuotaListWithDefaults

`func NewExceededQuotaListWithDefaults() *ExceededQuotaList`

NewExceededQuotaListWithDefaults instantiates a new ExceededQuotaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExceededQuotaList) GetData() []ExceededQuotaDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExceededQuotaList) GetDataOk() (*[]ExceededQuotaDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExceededQuotaList) SetData(v []ExceededQuotaDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ExceededQuotaList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ExceededQuotaList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ExceededQuotaList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExceededQuotaList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExceededQuotaList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


