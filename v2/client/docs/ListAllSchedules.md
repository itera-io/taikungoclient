# ListAllSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CScheduleDto**](CScheduleDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewListAllSchedules

`func NewListAllSchedules(data []CScheduleDto, totalCount int32, ) *ListAllSchedules`

NewListAllSchedules instantiates a new ListAllSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllSchedulesWithDefaults

`func NewListAllSchedulesWithDefaults() *ListAllSchedules`

NewListAllSchedulesWithDefaults instantiates a new ListAllSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllSchedules) GetData() []CScheduleDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllSchedules) GetDataOk() (*[]CScheduleDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllSchedules) SetData(v []CScheduleDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ListAllSchedules) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListAllSchedules) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ListAllSchedules) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAllSchedules) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAllSchedules) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


