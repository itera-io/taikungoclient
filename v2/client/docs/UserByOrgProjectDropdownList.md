# UserByOrgProjectDropdownList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** |  | 
**NextOffset** | Pointer to **NullableInt32** |  | [optional] 
**Data** | [**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewUserByOrgProjectDropdownList

`func NewUserByOrgProjectDropdownList(offset int32, data []CommonStringBasedDropdownDto, limit int32, hasMore bool, totalCount int64, ) *UserByOrgProjectDropdownList`

NewUserByOrgProjectDropdownList instantiates a new UserByOrgProjectDropdownList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserByOrgProjectDropdownListWithDefaults

`func NewUserByOrgProjectDropdownListWithDefaults() *UserByOrgProjectDropdownList`

NewUserByOrgProjectDropdownListWithDefaults instantiates a new UserByOrgProjectDropdownList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *UserByOrgProjectDropdownList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UserByOrgProjectDropdownList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UserByOrgProjectDropdownList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetNextOffset

`func (o *UserByOrgProjectDropdownList) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *UserByOrgProjectDropdownList) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *UserByOrgProjectDropdownList) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *UserByOrgProjectDropdownList) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *UserByOrgProjectDropdownList) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *UserByOrgProjectDropdownList) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetData

`func (o *UserByOrgProjectDropdownList) GetData() []CommonStringBasedDropdownDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserByOrgProjectDropdownList) GetDataOk() (*[]CommonStringBasedDropdownDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserByOrgProjectDropdownList) SetData(v []CommonStringBasedDropdownDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *UserByOrgProjectDropdownList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserByOrgProjectDropdownList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *UserByOrgProjectDropdownList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserByOrgProjectDropdownList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserByOrgProjectDropdownList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *UserByOrgProjectDropdownList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UserByOrgProjectDropdownList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UserByOrgProjectDropdownList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *UserByOrgProjectDropdownList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserByOrgProjectDropdownList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserByOrgProjectDropdownList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


