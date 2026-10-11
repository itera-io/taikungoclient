# GroupedPrometheusBillingListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **NullableInt32** |  | 
**PeriodDuration** | **map[string]interface{}** |  | 

## Methods

### NewGroupedPrometheusBillingListQuery

`func NewGroupedPrometheusBillingListQuery(organizationId NullableInt32, periodDuration map[string]interface{}, ) *GroupedPrometheusBillingListQuery`

NewGroupedPrometheusBillingListQuery instantiates a new GroupedPrometheusBillingListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedPrometheusBillingListQueryWithDefaults

`func NewGroupedPrometheusBillingListQueryWithDefaults() *GroupedPrometheusBillingListQuery`

NewGroupedPrometheusBillingListQueryWithDefaults instantiates a new GroupedPrometheusBillingListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *GroupedPrometheusBillingListQuery) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GroupedPrometheusBillingListQuery) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GroupedPrometheusBillingListQuery) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *GroupedPrometheusBillingListQuery) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *GroupedPrometheusBillingListQuery) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetPeriodDuration

`func (o *GroupedPrometheusBillingListQuery) GetPeriodDuration() map[string]interface{}`

GetPeriodDuration returns the PeriodDuration field if non-nil, zero value otherwise.

### GetPeriodDurationOk

`func (o *GroupedPrometheusBillingListQuery) GetPeriodDurationOk() (*map[string]interface{}, bool)`

GetPeriodDurationOk returns a tuple with the PeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDuration

`func (o *GroupedPrometheusBillingListQuery) SetPeriodDuration(v map[string]interface{})`

SetPeriodDuration sets PeriodDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


