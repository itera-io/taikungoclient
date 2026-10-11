# LokiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | [**Stream**](Stream.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewLokiResult

`func NewLokiResult(stream Stream, values []string, ) *LokiResult`

NewLokiResult instantiates a new LokiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLokiResultWithDefaults

`func NewLokiResultWithDefaults() *LokiResult`

NewLokiResultWithDefaults instantiates a new LokiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *LokiResult) GetStream() Stream`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *LokiResult) GetStreamOk() (*Stream, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *LokiResult) SetStream(v Stream)`

SetStream sets Stream field to given value.


### GetValues

`func (o *LokiResult) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LokiResult) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LokiResult) SetValues(v []string)`

SetValues sets Values field to given value.


### SetValuesNil

`func (o *LokiResult) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *LokiResult) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


