# MetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **NullableString** |  | 
**Result** | [**[]MetricResult**](MetricResult.md) |  | 

## Methods

### NewMetricData

`func NewMetricData(resultType NullableString, result []MetricResult, ) *MetricData`

NewMetricData instantiates a new MetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataWithDefaults

`func NewMetricDataWithDefaults() *MetricData`

NewMetricDataWithDefaults instantiates a new MetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *MetricData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *MetricData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *MetricData) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### SetResultTypeNil

`func (o *MetricData) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *MetricData) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetResult

`func (o *MetricData) GetResult() []MetricResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MetricData) GetResultOk() (*[]MetricResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MetricData) SetResult(v []MetricResult)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *MetricData) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *MetricData) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


