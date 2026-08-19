# CreateBillingSummaryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icu** | Pointer to **int32** |  | [optional] 
**BeginApply** | Pointer to **time.Time** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateBillingSummaryCommand

`func NewCreateBillingSummaryCommand() *CreateBillingSummaryCommand`

NewCreateBillingSummaryCommand instantiates a new CreateBillingSummaryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingSummaryCommandWithDefaults

`func NewCreateBillingSummaryCommandWithDefaults() *CreateBillingSummaryCommand`

NewCreateBillingSummaryCommandWithDefaults instantiates a new CreateBillingSummaryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcu

`func (o *CreateBillingSummaryCommand) GetIcu() int32`

GetIcu returns the Icu field if non-nil, zero value otherwise.

### GetIcuOk

`func (o *CreateBillingSummaryCommand) GetIcuOk() (*int32, bool)`

GetIcuOk returns a tuple with the Icu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcu

`func (o *CreateBillingSummaryCommand) SetIcu(v int32)`

SetIcu sets Icu field to given value.

### HasIcu

`func (o *CreateBillingSummaryCommand) HasIcu() bool`

HasIcu returns a boolean if a field has been set.

### GetBeginApply

`func (o *CreateBillingSummaryCommand) GetBeginApply() time.Time`

GetBeginApply returns the BeginApply field if non-nil, zero value otherwise.

### GetBeginApplyOk

`func (o *CreateBillingSummaryCommand) GetBeginApplyOk() (*time.Time, bool)`

GetBeginApplyOk returns a tuple with the BeginApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginApply

`func (o *CreateBillingSummaryCommand) SetBeginApply(v time.Time)`

SetBeginApply sets BeginApply field to given value.

### HasBeginApply

`func (o *CreateBillingSummaryCommand) HasBeginApply() bool`

HasBeginApply returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateBillingSummaryCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateBillingSummaryCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateBillingSummaryCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateBillingSummaryCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


