# DnsCredentialListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** |  | 
**NextOffset** | Pointer to **NullableInt32** |  | [optional] 
**Data** | [**[]DnsCredentialListDto**](DnsCredentialListDto.md) |  | 
**Limit** | **int32** |  | 
**HasMore** | **bool** |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewDnsCredentialListResponse

`func NewDnsCredentialListResponse(offset int32, data []DnsCredentialListDto, limit int32, hasMore bool, totalCount int64, ) *DnsCredentialListResponse`

NewDnsCredentialListResponse instantiates a new DnsCredentialListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCredentialListResponseWithDefaults

`func NewDnsCredentialListResponseWithDefaults() *DnsCredentialListResponse`

NewDnsCredentialListResponseWithDefaults instantiates a new DnsCredentialListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *DnsCredentialListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DnsCredentialListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DnsCredentialListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetNextOffset

`func (o *DnsCredentialListResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *DnsCredentialListResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *DnsCredentialListResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *DnsCredentialListResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### SetNextOffsetNil

`func (o *DnsCredentialListResponse) SetNextOffsetNil(b bool)`

 SetNextOffsetNil sets the value for NextOffset to be an explicit nil

### UnsetNextOffset
`func (o *DnsCredentialListResponse) UnsetNextOffset()`

UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil
### GetData

`func (o *DnsCredentialListResponse) GetData() []DnsCredentialListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsCredentialListResponse) GetDataOk() (*[]DnsCredentialListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsCredentialListResponse) SetData(v []DnsCredentialListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *DnsCredentialListResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DnsCredentialListResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *DnsCredentialListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DnsCredentialListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DnsCredentialListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetHasMore

`func (o *DnsCredentialListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DnsCredentialListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DnsCredentialListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *DnsCredentialListResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DnsCredentialListResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DnsCredentialListResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


