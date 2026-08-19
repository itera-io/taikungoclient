# OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **NullableInt32** |  | [optional] 
**Data** | [**[]OrganizationsWithGroupInfoResultDto**](OrganizationsWithGroupInfoResultDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponse

`func NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponse(data []OrganizationsWithGroupInfoResultDto, limit int32, hasMore bool, totalCount int64, ) *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse`

NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponse instantiates a new OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponseWithDefaults

`func NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponseWithDefaults() *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse`

NewOrganizationsWithGroupInfoResultDtoCursorPaginatedResponseWithDefaults instantiates a new OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetData

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetData() []OrganizationsWithGroupInfoResultDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetDataOk() (*[]OrganizationsWithGroupInfoResultDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetData(v []OrganizationsWithGroupInfoResultDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


