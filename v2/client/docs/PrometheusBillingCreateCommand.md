# PrometheusBillingCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int32** |  | 
**PrometheusRuleId** | **int32** |  | 
**StartDate** | **time.Time** |  | 
**Price** | **float64** |  | 

## Methods

### NewPrometheusBillingCreateCommand

`func NewPrometheusBillingCreateCommand(organizationId int32, prometheusRuleId int32, startDate time.Time, price float64, ) *PrometheusBillingCreateCommand`

NewPrometheusBillingCreateCommand instantiates a new PrometheusBillingCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusBillingCreateCommandWithDefaults

`func NewPrometheusBillingCreateCommandWithDefaults() *PrometheusBillingCreateCommand`

NewPrometheusBillingCreateCommandWithDefaults instantiates a new PrometheusBillingCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *PrometheusBillingCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PrometheusBillingCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PrometheusBillingCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetPrometheusRuleId

`func (o *PrometheusBillingCreateCommand) GetPrometheusRuleId() int32`

GetPrometheusRuleId returns the PrometheusRuleId field if non-nil, zero value otherwise.

### GetPrometheusRuleIdOk

`func (o *PrometheusBillingCreateCommand) GetPrometheusRuleIdOk() (*int32, bool)`

GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusRuleId

`func (o *PrometheusBillingCreateCommand) SetPrometheusRuleId(v int32)`

SetPrometheusRuleId sets PrometheusRuleId field to given value.


### GetStartDate

`func (o *PrometheusBillingCreateCommand) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PrometheusBillingCreateCommand) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PrometheusBillingCreateCommand) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetPrice

`func (o *PrometheusBillingCreateCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PrometheusBillingCreateCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PrometheusBillingCreateCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


