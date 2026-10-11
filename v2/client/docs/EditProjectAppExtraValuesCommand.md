# EditProjectAppExtraValuesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectAppId** | **int32** |  | 
**ExtraValues** | **NullableString** |  | 
**Timeout** | **NullableInt32** |  | 

## Methods

### NewEditProjectAppExtraValuesCommand

`func NewEditProjectAppExtraValuesCommand(projectAppId int32, extraValues NullableString, timeout NullableInt32, ) *EditProjectAppExtraValuesCommand`

NewEditProjectAppExtraValuesCommand instantiates a new EditProjectAppExtraValuesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditProjectAppExtraValuesCommandWithDefaults

`func NewEditProjectAppExtraValuesCommandWithDefaults() *EditProjectAppExtraValuesCommand`

NewEditProjectAppExtraValuesCommandWithDefaults instantiates a new EditProjectAppExtraValuesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectAppId

`func (o *EditProjectAppExtraValuesCommand) GetProjectAppId() int32`

GetProjectAppId returns the ProjectAppId field if non-nil, zero value otherwise.

### GetProjectAppIdOk

`func (o *EditProjectAppExtraValuesCommand) GetProjectAppIdOk() (*int32, bool)`

GetProjectAppIdOk returns a tuple with the ProjectAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppId

`func (o *EditProjectAppExtraValuesCommand) SetProjectAppId(v int32)`

SetProjectAppId sets ProjectAppId field to given value.


### GetExtraValues

`func (o *EditProjectAppExtraValuesCommand) GetExtraValues() string`

GetExtraValues returns the ExtraValues field if non-nil, zero value otherwise.

### GetExtraValuesOk

`func (o *EditProjectAppExtraValuesCommand) GetExtraValuesOk() (*string, bool)`

GetExtraValuesOk returns a tuple with the ExtraValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValues

`func (o *EditProjectAppExtraValuesCommand) SetExtraValues(v string)`

SetExtraValues sets ExtraValues field to given value.


### SetExtraValuesNil

`func (o *EditProjectAppExtraValuesCommand) SetExtraValuesNil(b bool)`

 SetExtraValuesNil sets the value for ExtraValues to be an explicit nil

### UnsetExtraValues
`func (o *EditProjectAppExtraValuesCommand) UnsetExtraValues()`

UnsetExtraValues ensures that no value is present for ExtraValues, not even an explicit nil
### GetTimeout

`func (o *EditProjectAppExtraValuesCommand) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EditProjectAppExtraValuesCommand) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EditProjectAppExtraValuesCommand) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *EditProjectAppExtraValuesCommand) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EditProjectAppExtraValuesCommand) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


