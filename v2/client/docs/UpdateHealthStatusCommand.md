# UpdateHealthStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Health** | Pointer to [**ProjectHealth**](ProjectHealth.md) |  | [optional] 

## Methods

### NewUpdateHealthStatusCommand

`func NewUpdateHealthStatusCommand() *UpdateHealthStatusCommand`

NewUpdateHealthStatusCommand instantiates a new UpdateHealthStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHealthStatusCommandWithDefaults

`func NewUpdateHealthStatusCommandWithDefaults() *UpdateHealthStatusCommand`

NewUpdateHealthStatusCommandWithDefaults instantiates a new UpdateHealthStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *UpdateHealthStatusCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateHealthStatusCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateHealthStatusCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateHealthStatusCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetHealth

`func (o *UpdateHealthStatusCommand) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateHealthStatusCommand) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateHealthStatusCommand) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *UpdateHealthStatusCommand) HasHealth() bool`

HasHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


