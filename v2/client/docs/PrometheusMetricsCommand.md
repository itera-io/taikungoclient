# PrometheusMetricsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to **NullableString** |  | [optional] 
**Time** | Pointer to **NullableTime** |  | [optional] 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**IsGraphEnabled** | Pointer to **bool** |  | [optional] 
**Step** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrometheusMetricsCommand

`func NewPrometheusMetricsCommand() *PrometheusMetricsCommand`

NewPrometheusMetricsCommand instantiates a new PrometheusMetricsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMetricsCommandWithDefaults

`func NewPrometheusMetricsCommandWithDefaults() *PrometheusMetricsCommand`

NewPrometheusMetricsCommandWithDefaults instantiates a new PrometheusMetricsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PrometheusMetricsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PrometheusMetricsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PrometheusMetricsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PrometheusMetricsCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParameters

`func (o *PrometheusMetricsCommand) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PrometheusMetricsCommand) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PrometheusMetricsCommand) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PrometheusMetricsCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PrometheusMetricsCommand) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PrometheusMetricsCommand) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetTime

`func (o *PrometheusMetricsCommand) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PrometheusMetricsCommand) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PrometheusMetricsCommand) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *PrometheusMetricsCommand) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *PrometheusMetricsCommand) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *PrometheusMetricsCommand) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetStart

`func (o *PrometheusMetricsCommand) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PrometheusMetricsCommand) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PrometheusMetricsCommand) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *PrometheusMetricsCommand) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *PrometheusMetricsCommand) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *PrometheusMetricsCommand) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *PrometheusMetricsCommand) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PrometheusMetricsCommand) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PrometheusMetricsCommand) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PrometheusMetricsCommand) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *PrometheusMetricsCommand) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *PrometheusMetricsCommand) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetIsGraphEnabled

`func (o *PrometheusMetricsCommand) GetIsGraphEnabled() bool`

GetIsGraphEnabled returns the IsGraphEnabled field if non-nil, zero value otherwise.

### GetIsGraphEnabledOk

`func (o *PrometheusMetricsCommand) GetIsGraphEnabledOk() (*bool, bool)`

GetIsGraphEnabledOk returns a tuple with the IsGraphEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGraphEnabled

`func (o *PrometheusMetricsCommand) SetIsGraphEnabled(v bool)`

SetIsGraphEnabled sets IsGraphEnabled field to given value.

### HasIsGraphEnabled

`func (o *PrometheusMetricsCommand) HasIsGraphEnabled() bool`

HasIsGraphEnabled returns a boolean if a field has been set.

### GetStep

`func (o *PrometheusMetricsCommand) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PrometheusMetricsCommand) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PrometheusMetricsCommand) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *PrometheusMetricsCommand) HasStep() bool`

HasStep returns a boolean if a field has been set.

### SetStepNil

`func (o *PrometheusMetricsCommand) SetStepNil(b bool)`

 SetStepNil sets the value for Step to be an explicit nil

### UnsetStep
`func (o *PrometheusMetricsCommand) UnsetStep()`

UnsetStep ensures that no value is present for Step, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


