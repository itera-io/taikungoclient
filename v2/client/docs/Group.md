# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**File** | **NullableString** |  | 
**Rules** | [**[]Rule**](Rule.md) |  | 
**Interval** | **int64** |  | 
**EvaluationTime** | **float64** |  | 
**LastEvaluation** | **time.Time** |  | 

## Methods

### NewGroup

`func NewGroup(name NullableString, file NullableString, rules []Rule, interval int64, evaluationTime float64, lastEvaluation time.Time, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Group) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Group) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFile

`func (o *Group) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Group) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Group) SetFile(v string)`

SetFile sets File field to given value.


### SetFileNil

`func (o *Group) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *Group) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetRules

`func (o *Group) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Group) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Group) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### SetRulesNil

`func (o *Group) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *Group) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetInterval

`func (o *Group) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Group) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Group) SetInterval(v int64)`

SetInterval sets Interval field to given value.


### GetEvaluationTime

`func (o *Group) GetEvaluationTime() float64`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *Group) GetEvaluationTimeOk() (*float64, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *Group) SetEvaluationTime(v float64)`

SetEvaluationTime sets EvaluationTime field to given value.


### GetLastEvaluation

`func (o *Group) GetLastEvaluation() time.Time`

GetLastEvaluation returns the LastEvaluation field if non-nil, zero value otherwise.

### GetLastEvaluationOk

`func (o *Group) GetLastEvaluationOk() (*time.Time, bool)`

GetLastEvaluationOk returns a tuple with the LastEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluation

`func (o *Group) SetLastEvaluation(v time.Time)`

SetLastEvaluation sets LastEvaluation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


