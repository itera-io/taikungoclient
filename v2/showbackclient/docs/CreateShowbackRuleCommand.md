# CreateShowbackRuleCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**MetricName** | Pointer to **NullableString** |  | [optional] 
**ByLabel** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to [**EShowbackType**](EShowbackType.md) |  | [optional] 
**Type** | Pointer to [**EPrometheusType**](EPrometheusType.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**ProjectAlertLimit** | Pointer to **NullableInt32** |  | [optional] 
**GlobalAlertLimit** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to [**[]ShowbackLabelCreateDto**](ShowbackLabelCreateDto.md) |  | [optional] 
**ShowbackCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateShowbackRuleCommand

`func NewCreateShowbackRuleCommand() *CreateShowbackRuleCommand`

NewCreateShowbackRuleCommand instantiates a new CreateShowbackRuleCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShowbackRuleCommandWithDefaults

`func NewCreateShowbackRuleCommandWithDefaults() *CreateShowbackRuleCommand`

NewCreateShowbackRuleCommandWithDefaults instantiates a new CreateShowbackRuleCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateShowbackRuleCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateShowbackRuleCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateShowbackRuleCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateShowbackRuleCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateShowbackRuleCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateShowbackRuleCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetricName

`func (o *CreateShowbackRuleCommand) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *CreateShowbackRuleCommand) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *CreateShowbackRuleCommand) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *CreateShowbackRuleCommand) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *CreateShowbackRuleCommand) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *CreateShowbackRuleCommand) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetByLabel

`func (o *CreateShowbackRuleCommand) GetByLabel() string`

GetByLabel returns the ByLabel field if non-nil, zero value otherwise.

### GetByLabelOk

`func (o *CreateShowbackRuleCommand) GetByLabelOk() (*string, bool)`

GetByLabelOk returns a tuple with the ByLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByLabel

`func (o *CreateShowbackRuleCommand) SetByLabel(v string)`

SetByLabel sets ByLabel field to given value.

### HasByLabel

`func (o *CreateShowbackRuleCommand) HasByLabel() bool`

HasByLabel returns a boolean if a field has been set.

### SetByLabelNil

`func (o *CreateShowbackRuleCommand) SetByLabelNil(b bool)`

 SetByLabelNil sets the value for ByLabel to be an explicit nil

### UnsetByLabel
`func (o *CreateShowbackRuleCommand) UnsetByLabel()`

UnsetByLabel ensures that no value is present for ByLabel, not even an explicit nil
### GetKind

`func (o *CreateShowbackRuleCommand) GetKind() EShowbackType`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateShowbackRuleCommand) GetKindOk() (*EShowbackType, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateShowbackRuleCommand) SetKind(v EShowbackType)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateShowbackRuleCommand) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetType

`func (o *CreateShowbackRuleCommand) GetType() EPrometheusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateShowbackRuleCommand) GetTypeOk() (*EPrometheusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateShowbackRuleCommand) SetType(v EPrometheusType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateShowbackRuleCommand) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *CreateShowbackRuleCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateShowbackRuleCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateShowbackRuleCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateShowbackRuleCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CreateShowbackRuleCommand) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CreateShowbackRuleCommand) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetProjectAlertLimit

`func (o *CreateShowbackRuleCommand) GetProjectAlertLimit() int32`

GetProjectAlertLimit returns the ProjectAlertLimit field if non-nil, zero value otherwise.

### GetProjectAlertLimitOk

`func (o *CreateShowbackRuleCommand) GetProjectAlertLimitOk() (*int32, bool)`

GetProjectAlertLimitOk returns a tuple with the ProjectAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAlertLimit

`func (o *CreateShowbackRuleCommand) SetProjectAlertLimit(v int32)`

SetProjectAlertLimit sets ProjectAlertLimit field to given value.

### HasProjectAlertLimit

`func (o *CreateShowbackRuleCommand) HasProjectAlertLimit() bool`

HasProjectAlertLimit returns a boolean if a field has been set.

### SetProjectAlertLimitNil

`func (o *CreateShowbackRuleCommand) SetProjectAlertLimitNil(b bool)`

 SetProjectAlertLimitNil sets the value for ProjectAlertLimit to be an explicit nil

### UnsetProjectAlertLimit
`func (o *CreateShowbackRuleCommand) UnsetProjectAlertLimit()`

UnsetProjectAlertLimit ensures that no value is present for ProjectAlertLimit, not even an explicit nil
### GetGlobalAlertLimit

`func (o *CreateShowbackRuleCommand) GetGlobalAlertLimit() int32`

GetGlobalAlertLimit returns the GlobalAlertLimit field if non-nil, zero value otherwise.

### GetGlobalAlertLimitOk

`func (o *CreateShowbackRuleCommand) GetGlobalAlertLimitOk() (*int32, bool)`

GetGlobalAlertLimitOk returns a tuple with the GlobalAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAlertLimit

`func (o *CreateShowbackRuleCommand) SetGlobalAlertLimit(v int32)`

SetGlobalAlertLimit sets GlobalAlertLimit field to given value.

### HasGlobalAlertLimit

`func (o *CreateShowbackRuleCommand) HasGlobalAlertLimit() bool`

HasGlobalAlertLimit returns a boolean if a field has been set.

### SetGlobalAlertLimitNil

`func (o *CreateShowbackRuleCommand) SetGlobalAlertLimitNil(b bool)`

 SetGlobalAlertLimitNil sets the value for GlobalAlertLimit to be an explicit nil

### UnsetGlobalAlertLimit
`func (o *CreateShowbackRuleCommand) UnsetGlobalAlertLimit()`

UnsetGlobalAlertLimit ensures that no value is present for GlobalAlertLimit, not even an explicit nil
### GetOrganizationId

`func (o *CreateShowbackRuleCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateShowbackRuleCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateShowbackRuleCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateShowbackRuleCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetLabels

`func (o *CreateShowbackRuleCommand) GetLabels() []ShowbackLabelCreateDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateShowbackRuleCommand) GetLabelsOk() (*[]ShowbackLabelCreateDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateShowbackRuleCommand) SetLabels(v []ShowbackLabelCreateDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateShowbackRuleCommand) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CreateShowbackRuleCommand) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CreateShowbackRuleCommand) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetShowbackCredentialId

`func (o *CreateShowbackRuleCommand) GetShowbackCredentialId() int32`

GetShowbackCredentialId returns the ShowbackCredentialId field if non-nil, zero value otherwise.

### GetShowbackCredentialIdOk

`func (o *CreateShowbackRuleCommand) GetShowbackCredentialIdOk() (*int32, bool)`

GetShowbackCredentialIdOk returns a tuple with the ShowbackCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackCredentialId

`func (o *CreateShowbackRuleCommand) SetShowbackCredentialId(v int32)`

SetShowbackCredentialId sets ShowbackCredentialId field to given value.

### HasShowbackCredentialId

`func (o *CreateShowbackRuleCommand) HasShowbackCredentialId() bool`

HasShowbackCredentialId returns a boolean if a field has been set.

### SetShowbackCredentialIdNil

`func (o *CreateShowbackRuleCommand) SetShowbackCredentialIdNil(b bool)`

 SetShowbackCredentialIdNil sets the value for ShowbackCredentialId to be an explicit nil

### UnsetShowbackCredentialId
`func (o *CreateShowbackRuleCommand) UnsetShowbackCredentialId()`

UnsetShowbackCredentialId ensures that no value is present for ShowbackCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


