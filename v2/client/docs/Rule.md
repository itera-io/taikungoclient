# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Query** | **NullableString** |  | 
**Duration** | **int64** |  | 
**Labels** | [**RuleLabels**](RuleLabels.md) |  | 
**Annotations** | [**Annotations**](Annotations.md) |  | 
**Alerts** | [**[]Alert**](Alert.md) |  | 
**Health** | **NullableString** |  | 
**EvaluationTime** | **float64** |  | 
**LastEvaluation** | **time.Time** |  | 
**Type** | **NullableString** |  | 

## Methods

### NewRule

`func NewRule(state NullableString, name NullableString, query NullableString, duration int64, labels RuleLabels, annotations Annotations, alerts []Alert, health NullableString, evaluationTime float64, lastEvaluation time.Time, type_ NullableString, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Rule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Rule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Rule) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *Rule) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Rule) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Rule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Rule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *Rule) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Rule) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Rule) SetQuery(v string)`

SetQuery sets Query field to given value.


### SetQueryNil

`func (o *Rule) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *Rule) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuration

`func (o *Rule) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Rule) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Rule) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetLabels

`func (o *Rule) GetLabels() RuleLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Rule) GetLabelsOk() (*RuleLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Rule) SetLabels(v RuleLabels)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *Rule) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Rule) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Rule) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.


### GetAlerts

`func (o *Rule) GetAlerts() []Alert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Rule) GetAlertsOk() (*[]Alert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Rule) SetAlerts(v []Alert)`

SetAlerts sets Alerts field to given value.


### SetAlertsNil

`func (o *Rule) SetAlertsNil(b bool)`

 SetAlertsNil sets the value for Alerts to be an explicit nil

### UnsetAlerts
`func (o *Rule) UnsetAlerts()`

UnsetAlerts ensures that no value is present for Alerts, not even an explicit nil
### GetHealth

`func (o *Rule) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Rule) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Rule) SetHealth(v string)`

SetHealth sets Health field to given value.


### SetHealthNil

`func (o *Rule) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *Rule) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil
### GetEvaluationTime

`func (o *Rule) GetEvaluationTime() float64`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *Rule) GetEvaluationTimeOk() (*float64, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *Rule) SetEvaluationTime(v float64)`

SetEvaluationTime sets EvaluationTime field to given value.


### GetLastEvaluation

`func (o *Rule) GetLastEvaluation() time.Time`

GetLastEvaluation returns the LastEvaluation field if non-nil, zero value otherwise.

### GetLastEvaluationOk

`func (o *Rule) GetLastEvaluationOk() (*time.Time, bool)`

GetLastEvaluationOk returns a tuple with the LastEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluation

`func (o *Rule) SetLastEvaluation(v time.Time)`

SetLastEvaluation sets LastEvaluation field to given value.


### GetType

`func (o *Rule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rule) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Rule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Rule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


