# ValidateTrustBundleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidCount** | **int32** |  | 
**Items** | [**[]ValidateTrustBundleItem**](ValidateTrustBundleItem.md) |  | 

## Methods

### NewValidateTrustBundleResponse

`func NewValidateTrustBundleResponse(validCount int32, items []ValidateTrustBundleItem, ) *ValidateTrustBundleResponse`

NewValidateTrustBundleResponse instantiates a new ValidateTrustBundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateTrustBundleResponseWithDefaults

`func NewValidateTrustBundleResponseWithDefaults() *ValidateTrustBundleResponse`

NewValidateTrustBundleResponseWithDefaults instantiates a new ValidateTrustBundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidCount

`func (o *ValidateTrustBundleResponse) GetValidCount() int32`

GetValidCount returns the ValidCount field if non-nil, zero value otherwise.

### GetValidCountOk

`func (o *ValidateTrustBundleResponse) GetValidCountOk() (*int32, bool)`

GetValidCountOk returns a tuple with the ValidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCount

`func (o *ValidateTrustBundleResponse) SetValidCount(v int32)`

SetValidCount sets ValidCount field to given value.


### GetItems

`func (o *ValidateTrustBundleResponse) GetItems() []ValidateTrustBundleItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ValidateTrustBundleResponse) GetItemsOk() (*[]ValidateTrustBundleItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ValidateTrustBundleResponse) SetItems(v []ValidateTrustBundleItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


