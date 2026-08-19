# CreateShowbackSummaryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginApply** | Pointer to **time.Time** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**ShowbackRuleId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**ByLabelValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateShowbackSummaryCommand

`func NewCreateShowbackSummaryCommand() *CreateShowbackSummaryCommand`

NewCreateShowbackSummaryCommand instantiates a new CreateShowbackSummaryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShowbackSummaryCommandWithDefaults

`func NewCreateShowbackSummaryCommandWithDefaults() *CreateShowbackSummaryCommand`

NewCreateShowbackSummaryCommandWithDefaults instantiates a new CreateShowbackSummaryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeginApply

`func (o *CreateShowbackSummaryCommand) GetBeginApply() time.Time`

GetBeginApply returns the BeginApply field if non-nil, zero value otherwise.

### GetBeginApplyOk

`func (o *CreateShowbackSummaryCommand) GetBeginApplyOk() (*time.Time, bool)`

GetBeginApplyOk returns a tuple with the BeginApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginApply

`func (o *CreateShowbackSummaryCommand) SetBeginApply(v time.Time)`

SetBeginApply sets BeginApply field to given value.

### HasBeginApply

`func (o *CreateShowbackSummaryCommand) HasBeginApply() bool`

HasBeginApply returns a boolean if a field has been set.

### GetPrice

`func (o *CreateShowbackSummaryCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateShowbackSummaryCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateShowbackSummaryCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateShowbackSummaryCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetShowbackRuleId

`func (o *CreateShowbackSummaryCommand) GetShowbackRuleId() int32`

GetShowbackRuleId returns the ShowbackRuleId field if non-nil, zero value otherwise.

### GetShowbackRuleIdOk

`func (o *CreateShowbackSummaryCommand) GetShowbackRuleIdOk() (*int32, bool)`

GetShowbackRuleIdOk returns a tuple with the ShowbackRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackRuleId

`func (o *CreateShowbackSummaryCommand) SetShowbackRuleId(v int32)`

SetShowbackRuleId sets ShowbackRuleId field to given value.

### HasShowbackRuleId

`func (o *CreateShowbackSummaryCommand) HasShowbackRuleId() bool`

HasShowbackRuleId returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateShowbackSummaryCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateShowbackSummaryCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateShowbackSummaryCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateShowbackSummaryCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *CreateShowbackSummaryCommand) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *CreateShowbackSummaryCommand) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetByLabelValue

`func (o *CreateShowbackSummaryCommand) GetByLabelValue() string`

GetByLabelValue returns the ByLabelValue field if non-nil, zero value otherwise.

### GetByLabelValueOk

`func (o *CreateShowbackSummaryCommand) GetByLabelValueOk() (*string, bool)`

GetByLabelValueOk returns a tuple with the ByLabelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByLabelValue

`func (o *CreateShowbackSummaryCommand) SetByLabelValue(v string)`

SetByLabelValue sets ByLabelValue field to given value.

### HasByLabelValue

`func (o *CreateShowbackSummaryCommand) HasByLabelValue() bool`

HasByLabelValue returns a boolean if a field has been set.

### SetByLabelValueNil

`func (o *CreateShowbackSummaryCommand) SetByLabelValueNil(b bool)`

 SetByLabelValueNil sets the value for ByLabelValue to be an explicit nil

### UnsetByLabelValue
`func (o *CreateShowbackSummaryCommand) UnsetByLabelValue()`

UnsetByLabelValue ensures that no value is present for ByLabelValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


