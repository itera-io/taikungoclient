# BillingCredentialsSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **NullableInt32** |  | 
**Offset** | **NullableInt32** |  | 
**SearchTerm** | **NullableString** |  | 

## Methods

### NewBillingCredentialsSearchCommand

`func NewBillingCredentialsSearchCommand(limit NullableInt32, offset NullableInt32, searchTerm NullableString, ) *BillingCredentialsSearchCommand`

NewBillingCredentialsSearchCommand instantiates a new BillingCredentialsSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCredentialsSearchCommandWithDefaults

`func NewBillingCredentialsSearchCommandWithDefaults() *BillingCredentialsSearchCommand`

NewBillingCredentialsSearchCommandWithDefaults instantiates a new BillingCredentialsSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BillingCredentialsSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BillingCredentialsSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BillingCredentialsSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *BillingCredentialsSearchCommand) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *BillingCredentialsSearchCommand) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *BillingCredentialsSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BillingCredentialsSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BillingCredentialsSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### SetOffsetNil

`func (o *BillingCredentialsSearchCommand) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *BillingCredentialsSearchCommand) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearchTerm

`func (o *BillingCredentialsSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *BillingCredentialsSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *BillingCredentialsSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.


### SetSearchTermNil

`func (o *BillingCredentialsSearchCommand) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *BillingCredentialsSearchCommand) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


