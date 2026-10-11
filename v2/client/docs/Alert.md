# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**AlertLabels**](AlertLabels.md) |  | 
**Annotations** | [**Annotations**](Annotations.md) |  | 
**State** | **NullableString** |  | 
**ActiveAt** | **time.Time** |  | 
**Value** | **NullableString** |  | 

## Methods

### NewAlert

`func NewAlert(labels AlertLabels, annotations Annotations, state NullableString, activeAt time.Time, value NullableString, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *Alert) GetLabels() AlertLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Alert) GetLabelsOk() (*AlertLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Alert) SetLabels(v AlertLabels)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *Alert) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Alert) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Alert) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.


### GetState

`func (o *Alert) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Alert) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Alert) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *Alert) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Alert) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetActiveAt

`func (o *Alert) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *Alert) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *Alert) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.


### GetValue

`func (o *Alert) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Alert) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Alert) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Alert) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Alert) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


