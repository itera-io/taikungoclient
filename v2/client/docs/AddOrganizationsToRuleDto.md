# AddOrganizationsToRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int32** |  | [optional] 
**RuleDiscountRate** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewAddOrganizationsToRuleDto

`func NewAddOrganizationsToRuleDto() *AddOrganizationsToRuleDto`

NewAddOrganizationsToRuleDto instantiates a new AddOrganizationsToRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOrganizationsToRuleDtoWithDefaults

`func NewAddOrganizationsToRuleDtoWithDefaults() *AddOrganizationsToRuleDto`

NewAddOrganizationsToRuleDtoWithDefaults instantiates a new AddOrganizationsToRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AddOrganizationsToRuleDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AddOrganizationsToRuleDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AddOrganizationsToRuleDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AddOrganizationsToRuleDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRuleDiscountRate

`func (o *AddOrganizationsToRuleDto) GetRuleDiscountRate() float64`

GetRuleDiscountRate returns the RuleDiscountRate field if non-nil, zero value otherwise.

### GetRuleDiscountRateOk

`func (o *AddOrganizationsToRuleDto) GetRuleDiscountRateOk() (*float64, bool)`

GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDiscountRate

`func (o *AddOrganizationsToRuleDto) SetRuleDiscountRate(v float64)`

SetRuleDiscountRate sets RuleDiscountRate field to given value.

### HasRuleDiscountRate

`func (o *AddOrganizationsToRuleDto) HasRuleDiscountRate() bool`

HasRuleDiscountRate returns a boolean if a field has been set.

### SetRuleDiscountRateNil

`func (o *AddOrganizationsToRuleDto) SetRuleDiscountRateNil(b bool)`

 SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil

### UnsetRuleDiscountRate
`func (o *AddOrganizationsToRuleDto) UnsetRuleDiscountRate()`

UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


