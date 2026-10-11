# DaemonsetListDtoCursorStringPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**Data** | [**[]DaemonsetListDto**](DaemonsetListDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewDaemonsetListDtoCursorStringPaginatedResponse

`func NewDaemonsetListDtoCursorStringPaginatedResponse(data []DaemonsetListDto, limit int32, hasMore bool, totalCount int64, ) *DaemonsetListDtoCursorStringPaginatedResponse`

NewDaemonsetListDtoCursorStringPaginatedResponse instantiates a new DaemonsetListDtoCursorStringPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonsetListDtoCursorStringPaginatedResponseWithDefaults

`func NewDaemonsetListDtoCursorStringPaginatedResponseWithDefaults() *DaemonsetListDtoCursorStringPaginatedResponse`

NewDaemonsetListDtoCursorStringPaginatedResponseWithDefaults instantiates a new DaemonsetListDtoCursorStringPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *DaemonsetListDtoCursorStringPaginatedResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetData

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetData() []DaemonsetListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetDataOk() (*[]DaemonsetListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetData(v []DaemonsetListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DaemonsetListDtoCursorStringPaginatedResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DaemonsetListDtoCursorStringPaginatedResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


