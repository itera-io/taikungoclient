# GoogleFlavorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GoogleFlavorDto**](GoogleFlavorDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewGoogleFlavorList

`func NewGoogleFlavorList(data []GoogleFlavorDto, totalCount int32, ) *GoogleFlavorList`

NewGoogleFlavorList instantiates a new GoogleFlavorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleFlavorListWithDefaults

`func NewGoogleFlavorListWithDefaults() *GoogleFlavorList`

NewGoogleFlavorListWithDefaults instantiates a new GoogleFlavorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GoogleFlavorList) GetData() []GoogleFlavorDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GoogleFlavorList) GetDataOk() (*[]GoogleFlavorDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GoogleFlavorList) SetData(v []GoogleFlavorDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *GoogleFlavorList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GoogleFlavorList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *GoogleFlavorList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GoogleFlavorList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GoogleFlavorList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


