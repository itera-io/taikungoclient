# OpenstackFlavorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OpenstackFlavorListDto**](OpenstackFlavorListDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewOpenstackFlavorList

`func NewOpenstackFlavorList(data []OpenstackFlavorListDto, totalCount int32, ) *OpenstackFlavorList`

NewOpenstackFlavorList instantiates a new OpenstackFlavorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackFlavorListWithDefaults

`func NewOpenstackFlavorListWithDefaults() *OpenstackFlavorList`

NewOpenstackFlavorListWithDefaults instantiates a new OpenstackFlavorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OpenstackFlavorList) GetData() []OpenstackFlavorListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenstackFlavorList) GetDataOk() (*[]OpenstackFlavorListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenstackFlavorList) SetData(v []OpenstackFlavorListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *OpenstackFlavorList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OpenstackFlavorList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *OpenstackFlavorList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpenstackFlavorList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpenstackFlavorList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


