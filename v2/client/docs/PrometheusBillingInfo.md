# PrometheusBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PrometheusBillingSummaryDto**](PrometheusBillingSummaryDto.md) |  | 
**TotalPrice** | **float64** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPrometheusBillingInfo

`func NewPrometheusBillingInfo(data []PrometheusBillingSummaryDto, totalPrice float64, totalCount int32, ) *PrometheusBillingInfo`

NewPrometheusBillingInfo instantiates a new PrometheusBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusBillingInfoWithDefaults

`func NewPrometheusBillingInfoWithDefaults() *PrometheusBillingInfo`

NewPrometheusBillingInfoWithDefaults instantiates a new PrometheusBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PrometheusBillingInfo) GetData() []PrometheusBillingSummaryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrometheusBillingInfo) GetDataOk() (*[]PrometheusBillingSummaryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrometheusBillingInfo) SetData(v []PrometheusBillingSummaryDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *PrometheusBillingInfo) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PrometheusBillingInfo) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalPrice

`func (o *PrometheusBillingInfo) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *PrometheusBillingInfo) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *PrometheusBillingInfo) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.


### GetTotalCount

`func (o *PrometheusBillingInfo) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PrometheusBillingInfo) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PrometheusBillingInfo) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


