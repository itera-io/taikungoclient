# GroupedShowbackSummaryListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **int32** |  | [optional] 
**RuleName** | Pointer to **NullableString** |  | [optional] 
**TotalPrice** | Pointer to **float64** |  | [optional] 
**Projects** | Pointer to [**[]GroupedProjectForShowbackSummary**](GroupedProjectForShowbackSummary.md) |  | [optional] 
**Credential** | Pointer to [**GroupedCredentialForShowbackSummary**](GroupedCredentialForShowbackSummary.md) |  | [optional] 

## Methods

### NewGroupedShowbackSummaryListDto

`func NewGroupedShowbackSummaryListDto() *GroupedShowbackSummaryListDto`

NewGroupedShowbackSummaryListDto instantiates a new GroupedShowbackSummaryListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedShowbackSummaryListDtoWithDefaults

`func NewGroupedShowbackSummaryListDtoWithDefaults() *GroupedShowbackSummaryListDto`

NewGroupedShowbackSummaryListDtoWithDefaults instantiates a new GroupedShowbackSummaryListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *GroupedShowbackSummaryListDto) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GroupedShowbackSummaryListDto) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GroupedShowbackSummaryListDto) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GroupedShowbackSummaryListDto) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *GroupedShowbackSummaryListDto) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *GroupedShowbackSummaryListDto) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *GroupedShowbackSummaryListDto) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *GroupedShowbackSummaryListDto) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleNameNil

`func (o *GroupedShowbackSummaryListDto) SetRuleNameNil(b bool)`

 SetRuleNameNil sets the value for RuleName to be an explicit nil

### UnsetRuleName
`func (o *GroupedShowbackSummaryListDto) UnsetRuleName()`

UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
### GetTotalPrice

`func (o *GroupedShowbackSummaryListDto) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *GroupedShowbackSummaryListDto) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *GroupedShowbackSummaryListDto) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *GroupedShowbackSummaryListDto) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetProjects

`func (o *GroupedShowbackSummaryListDto) GetProjects() []GroupedProjectForShowbackSummary`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GroupedShowbackSummaryListDto) GetProjectsOk() (*[]GroupedProjectForShowbackSummary, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GroupedShowbackSummaryListDto) SetProjects(v []GroupedProjectForShowbackSummary)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GroupedShowbackSummaryListDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *GroupedShowbackSummaryListDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *GroupedShowbackSummaryListDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetCredential

`func (o *GroupedShowbackSummaryListDto) GetCredential() GroupedCredentialForShowbackSummary`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GroupedShowbackSummaryListDto) GetCredentialOk() (*GroupedCredentialForShowbackSummary, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GroupedShowbackSummaryListDto) SetCredential(v GroupedCredentialForShowbackSummary)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GroupedShowbackSummaryListDto) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


