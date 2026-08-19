# RuleForUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**MetricName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **map[string]interface{}** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Labels** | Pointer to [**[]PrometheusLabelListDto**](PrometheusLabelListDto.md) |  | [optional] 
**OperationCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRuleForUpdateDto

`func NewRuleForUpdateDto() *RuleForUpdateDto`

NewRuleForUpdateDto instantiates a new RuleForUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleForUpdateDtoWithDefaults

`func NewRuleForUpdateDtoWithDefaults() *RuleForUpdateDto`

NewRuleForUpdateDtoWithDefaults instantiates a new RuleForUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleForUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleForUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleForUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleForUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RuleForUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RuleForUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetricName

`func (o *RuleForUpdateDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *RuleForUpdateDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *RuleForUpdateDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *RuleForUpdateDto) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *RuleForUpdateDto) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *RuleForUpdateDto) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetType

`func (o *RuleForUpdateDto) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleForUpdateDto) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleForUpdateDto) SetType(v map[string]interface{})`

SetType sets Type field to given value.

### HasType

`func (o *RuleForUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *RuleForUpdateDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RuleForUpdateDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RuleForUpdateDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RuleForUpdateDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *RuleForUpdateDto) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *RuleForUpdateDto) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetLabels

`func (o *RuleForUpdateDto) GetLabels() []PrometheusLabelListDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RuleForUpdateDto) GetLabelsOk() (*[]PrometheusLabelListDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RuleForUpdateDto) SetLabels(v []PrometheusLabelListDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RuleForUpdateDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *RuleForUpdateDto) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RuleForUpdateDto) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOperationCredentialId

`func (o *RuleForUpdateDto) GetOperationCredentialId() int32`

GetOperationCredentialId returns the OperationCredentialId field if non-nil, zero value otherwise.

### GetOperationCredentialIdOk

`func (o *RuleForUpdateDto) GetOperationCredentialIdOk() (*int32, bool)`

GetOperationCredentialIdOk returns a tuple with the OperationCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCredentialId

`func (o *RuleForUpdateDto) SetOperationCredentialId(v int32)`

SetOperationCredentialId sets OperationCredentialId field to given value.

### HasOperationCredentialId

`func (o *RuleForUpdateDto) HasOperationCredentialId() bool`

HasOperationCredentialId returns a boolean if a field has been set.

### SetOperationCredentialIdNil

`func (o *RuleForUpdateDto) SetOperationCredentialIdNil(b bool)`

 SetOperationCredentialIdNil sets the value for OperationCredentialId to be an explicit nil

### UnsetOperationCredentialId
`func (o *RuleForUpdateDto) UnsetOperationCredentialId()`

UnsetOperationCredentialId ensures that no value is present for OperationCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


