# MetricResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **map[string]string** |  | 
**Value** | **map[string]interface{}** |  | 
**Values** | **[][]float64** |  | 

## Methods

### NewMetricResult

`func NewMetricResult(metric map[string]string, value map[string]interface{}, values [][]float64, ) *MetricResult`

NewMetricResult instantiates a new MetricResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResultWithDefaults

`func NewMetricResultWithDefaults() *MetricResult`

NewMetricResultWithDefaults instantiates a new MetricResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.


### SetMetricNil

`func (o *MetricResult) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *MetricResult) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetValue

`func (o *MetricResult) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricResult) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricResult) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetValues

`func (o *MetricResult) GetValues() [][]float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricResult) GetValuesOk() (*[][]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricResult) SetValues(v [][]float64)`

SetValues sets Values field to given value.


### SetValuesNil

`func (o *MetricResult) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *MetricResult) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


