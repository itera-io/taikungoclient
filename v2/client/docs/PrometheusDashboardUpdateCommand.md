# PrometheusDashboardUpdateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Expression** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**CategoryName** | **NullableString** |  | 

## Methods

### NewPrometheusDashboardUpdateCommand

`func NewPrometheusDashboardUpdateCommand(id int32, name NullableString, expression NullableString, description NullableString, categoryName NullableString, ) *PrometheusDashboardUpdateCommand`

NewPrometheusDashboardUpdateCommand instantiates a new PrometheusDashboardUpdateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusDashboardUpdateCommandWithDefaults

`func NewPrometheusDashboardUpdateCommandWithDefaults() *PrometheusDashboardUpdateCommand`

NewPrometheusDashboardUpdateCommandWithDefaults instantiates a new PrometheusDashboardUpdateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusDashboardUpdateCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusDashboardUpdateCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusDashboardUpdateCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PrometheusDashboardUpdateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrometheusDashboardUpdateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrometheusDashboardUpdateCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PrometheusDashboardUpdateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PrometheusDashboardUpdateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpression

`func (o *PrometheusDashboardUpdateCommand) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PrometheusDashboardUpdateCommand) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PrometheusDashboardUpdateCommand) SetExpression(v string)`

SetExpression sets Expression field to given value.


### SetExpressionNil

`func (o *PrometheusDashboardUpdateCommand) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *PrometheusDashboardUpdateCommand) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil
### GetDescription

`func (o *PrometheusDashboardUpdateCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrometheusDashboardUpdateCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrometheusDashboardUpdateCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PrometheusDashboardUpdateCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrometheusDashboardUpdateCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategoryName

`func (o *PrometheusDashboardUpdateCommand) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *PrometheusDashboardUpdateCommand) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *PrometheusDashboardUpdateCommand) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### SetCategoryNameNil

`func (o *PrometheusDashboardUpdateCommand) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *PrometheusDashboardUpdateCommand) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


