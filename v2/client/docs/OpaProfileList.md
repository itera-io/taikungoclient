# OpaProfileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OpaProfileListDto**](OpaProfileListDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewOpaProfileList

`func NewOpaProfileList(data []OpaProfileListDto, totalCount int32, ) *OpaProfileList`

NewOpaProfileList instantiates a new OpaProfileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaProfileListWithDefaults

`func NewOpaProfileListWithDefaults() *OpaProfileList`

NewOpaProfileListWithDefaults instantiates a new OpaProfileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OpaProfileList) GetData() []OpaProfileListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpaProfileList) GetDataOk() (*[]OpaProfileListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpaProfileList) SetData(v []OpaProfileListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *OpaProfileList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OpaProfileList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *OpaProfileList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpaProfileList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpaProfileList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


