# UserBriefDtoCursorTimestampPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **NullableInt64** |  | [optional] 
**Data** | [**[]UserBriefDto**](UserBriefDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewUserBriefDtoCursorTimestampPaginatedResponse

`func NewUserBriefDtoCursorTimestampPaginatedResponse(data []UserBriefDto, limit int32, hasMore bool, totalCount int64, ) *UserBriefDtoCursorTimestampPaginatedResponse`

NewUserBriefDtoCursorTimestampPaginatedResponse instantiates a new UserBriefDtoCursorTimestampPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBriefDtoCursorTimestampPaginatedResponseWithDefaults

`func NewUserBriefDtoCursorTimestampPaginatedResponseWithDefaults() *UserBriefDtoCursorTimestampPaginatedResponse`

NewUserBriefDtoCursorTimestampPaginatedResponseWithDefaults instantiates a new UserBriefDtoCursorTimestampPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetNextCursor() int64`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetNextCursorOk() (*int64, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetNextCursor(v int64)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *UserBriefDtoCursorTimestampPaginatedResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetData

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetData() []UserBriefDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetDataOk() (*[]UserBriefDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetData(v []UserBriefDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserBriefDtoCursorTimestampPaginatedResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserBriefDtoCursorTimestampPaginatedResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


