# AwsFlavorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AwsFlavorListDto**](AwsFlavorListDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAwsFlavorList

`func NewAwsFlavorList(data []AwsFlavorListDto, totalCount int32, ) *AwsFlavorList`

NewAwsFlavorList instantiates a new AwsFlavorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsFlavorListWithDefaults

`func NewAwsFlavorListWithDefaults() *AwsFlavorList`

NewAwsFlavorListWithDefaults instantiates a new AwsFlavorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AwsFlavorList) GetData() []AwsFlavorListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AwsFlavorList) GetDataOk() (*[]AwsFlavorListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AwsFlavorList) SetData(v []AwsFlavorListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AwsFlavorList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AwsFlavorList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AwsFlavorList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AwsFlavorList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AwsFlavorList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


