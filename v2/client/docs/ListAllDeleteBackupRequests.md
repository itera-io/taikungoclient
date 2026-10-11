# ListAllDeleteBackupRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CDeleteBackupRequestDto**](CDeleteBackupRequestDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewListAllDeleteBackupRequests

`func NewListAllDeleteBackupRequests(data []CDeleteBackupRequestDto, totalCount int32, ) *ListAllDeleteBackupRequests`

NewListAllDeleteBackupRequests instantiates a new ListAllDeleteBackupRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllDeleteBackupRequestsWithDefaults

`func NewListAllDeleteBackupRequestsWithDefaults() *ListAllDeleteBackupRequests`

NewListAllDeleteBackupRequestsWithDefaults instantiates a new ListAllDeleteBackupRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllDeleteBackupRequests) GetData() []CDeleteBackupRequestDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllDeleteBackupRequests) GetDataOk() (*[]CDeleteBackupRequestDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllDeleteBackupRequests) SetData(v []CDeleteBackupRequestDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ListAllDeleteBackupRequests) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListAllDeleteBackupRequests) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ListAllDeleteBackupRequests) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListAllDeleteBackupRequests) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListAllDeleteBackupRequests) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


