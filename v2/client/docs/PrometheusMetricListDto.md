# PrometheusMetricListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to [**MetricData**](MetricData.md) |  | [optional] 

## Methods

### NewPrometheusMetricListDto

`func NewPrometheusMetricListDto() *PrometheusMetricListDto`

NewPrometheusMetricListDto instantiates a new PrometheusMetricListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMetricListDtoWithDefaults

`func NewPrometheusMetricListDtoWithDefaults() *PrometheusMetricListDto`

NewPrometheusMetricListDtoWithDefaults instantiates a new PrometheusMetricListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PrometheusMetricListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrometheusMetricListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrometheusMetricListDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrometheusMetricListDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PrometheusMetricListDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PrometheusMetricListDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetData

`func (o *PrometheusMetricListDto) GetData() MetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrometheusMetricListDto) GetDataOk() (*MetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrometheusMetricListDto) SetData(v MetricData)`

SetData sets Data field to given value.

### HasData

`func (o *PrometheusMetricListDto) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


