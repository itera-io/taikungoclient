# BillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingSummaryDto**](BillingSummaryDto.md) |  | 
**TotalTcu** | **float64** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewBillingInfo

`func NewBillingInfo(data []BillingSummaryDto, totalTcu float64, totalCount int32, ) *BillingInfo`

NewBillingInfo instantiates a new BillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoWithDefaults

`func NewBillingInfoWithDefaults() *BillingInfo`

NewBillingInfoWithDefaults instantiates a new BillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BillingInfo) GetData() []BillingSummaryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingInfo) GetDataOk() (*[]BillingSummaryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingInfo) SetData(v []BillingSummaryDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *BillingInfo) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BillingInfo) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalTcu

`func (o *BillingInfo) GetTotalTcu() float64`

GetTotalTcu returns the TotalTcu field if non-nil, zero value otherwise.

### GetTotalTcuOk

`func (o *BillingInfo) GetTotalTcuOk() (*float64, bool)`

GetTotalTcuOk returns a tuple with the TotalTcu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTcu

`func (o *BillingInfo) SetTotalTcu(v float64)`

SetTotalTcu sets TotalTcu field to given value.


### GetTotalCount

`func (o *BillingInfo) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BillingInfo) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BillingInfo) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


