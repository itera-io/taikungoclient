# ButtonStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Reasons** | Pointer to **[]string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 

## Methods

### NewButtonStatusDto

`func NewButtonStatusDto() *ButtonStatusDto`

NewButtonStatusDto instantiates a new ButtonStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonStatusDtoWithDefaults

`func NewButtonStatusDtoWithDefaults() *ButtonStatusDto`

NewButtonStatusDtoWithDefaults instantiates a new ButtonStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ButtonStatusDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ButtonStatusDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ButtonStatusDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ButtonStatusDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetReasons

`func (o *ButtonStatusDto) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ButtonStatusDto) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ButtonStatusDto) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ButtonStatusDto) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *ButtonStatusDto) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *ButtonStatusDto) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil
### GetHidden

`func (o *ButtonStatusDto) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ButtonStatusDto) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ButtonStatusDto) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ButtonStatusDto) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


