# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastBackup** | Pointer to **time.Time** |  | [optional] 
**Phase** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastBackup

`func (o *Status) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *Status) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *Status) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *Status) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetPhase

`func (o *Status) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *Status) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *Status) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *Status) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *Status) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *Status) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


