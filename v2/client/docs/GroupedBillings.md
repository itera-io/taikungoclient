# GroupedBillings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **NullableString** |  | [optional] 
**Tcu** | Pointer to **int64** |  | [optional] 

## Methods

### NewGroupedBillings

`func NewGroupedBillings() *GroupedBillings`

NewGroupedBillings instantiates a new GroupedBillings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedBillingsWithDefaults

`func NewGroupedBillingsWithDefaults() *GroupedBillings`

NewGroupedBillingsWithDefaults instantiates a new GroupedBillings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *GroupedBillings) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GroupedBillings) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GroupedBillings) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GroupedBillings) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GroupedBillings) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GroupedBillings) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetTcu

`func (o *GroupedBillings) GetTcu() int64`

GetTcu returns the Tcu field if non-nil, zero value otherwise.

### GetTcuOk

`func (o *GroupedBillings) GetTcuOk() (*int64, bool)`

GetTcuOk returns a tuple with the Tcu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcu

`func (o *GroupedBillings) SetTcu(v int64)`

SetTcu sets Tcu field to given value.

### HasTcu

`func (o *GroupedBillings) HasTcu() bool`

HasTcu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


