# OrganizationsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** |  | 
**NextOffset** | Pointer to **NullableInt32** |  | [optional] 
**Data** | [**[]OrganizationDetailsDto**](OrganizationDetailsDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewOrganizationsList

`func NewOrganizationsList(offset int32, data []OrganizationDetailsDto, limit int32, hasMore bool, totalCount int64, ) *OrganizationsList`

NewOrganizationsList instantiates a new OrganizationsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsListWithDefaults

`func NewOrganizationsListWithDefaults() *OrganizationsList`

NewOrganizationsListWithDefaults instantiates a new OrganizationsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *OrganizationsList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *OrganizationsList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *OrganizationsList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetNextOffset

`func (o *OrganizationsList) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *OrganizationsList) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *OrganizationsList) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *OrganizationsList) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *OrganizationsList) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *OrganizationsList) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetData

`func (o *OrganizationsList) GetData() []OrganizationDetailsDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationsList) GetDataOk() (*[]OrganizationDetailsDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationsList) SetData(v []OrganizationDetailsDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *OrganizationsList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrganizationsList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *OrganizationsList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OrganizationsList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OrganizationsList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *OrganizationsList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrganizationsList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrganizationsList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *OrganizationsList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrganizationsList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrganizationsList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


