# GroupedBillingListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **NullableInt32** |  | 
**PeriodDuration** | **map[string]interface{}** |  | 
**IsDeleted** | **NullableBool** |  | 

## Methods

### NewGroupedBillingListQuery

`func NewGroupedBillingListQuery(organizationId NullableInt32, periodDuration map[string]interface{}, isDeleted NullableBool, ) *GroupedBillingListQuery`

NewGroupedBillingListQuery instantiates a new GroupedBillingListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedBillingListQueryWithDefaults

`func NewGroupedBillingListQueryWithDefaults() *GroupedBillingListQuery`

NewGroupedBillingListQueryWithDefaults instantiates a new GroupedBillingListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *GroupedBillingListQuery) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GroupedBillingListQuery) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GroupedBillingListQuery) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *GroupedBillingListQuery) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *GroupedBillingListQuery) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetPeriodDuration

`func (o *GroupedBillingListQuery) GetPeriodDuration() map[string]interface{}`

GetPeriodDuration returns the PeriodDuration field if non-nil, zero value otherwise.

### GetPeriodDurationOk

`func (o *GroupedBillingListQuery) GetPeriodDurationOk() (*map[string]interface{}, bool)`

GetPeriodDurationOk returns a tuple with the PeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDuration

`func (o *GroupedBillingListQuery) SetPeriodDuration(v map[string]interface{})`

SetPeriodDuration sets PeriodDuration field to given value.


### GetIsDeleted

`func (o *GroupedBillingListQuery) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *GroupedBillingListQuery) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *GroupedBillingListQuery) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### SetIsDeletedNil

`func (o *GroupedBillingListQuery) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *GroupedBillingListQuery) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


