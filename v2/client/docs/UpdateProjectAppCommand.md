# UpdateProjectAppCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectAppId** | **int32** |  | 
**Status** | **NullableString** |  | 

## Methods

### NewUpdateProjectAppCommand

`func NewUpdateProjectAppCommand(projectAppId int32, status NullableString, ) *UpdateProjectAppCommand`

NewUpdateProjectAppCommand instantiates a new UpdateProjectAppCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectAppCommandWithDefaults

`func NewUpdateProjectAppCommandWithDefaults() *UpdateProjectAppCommand`

NewUpdateProjectAppCommandWithDefaults instantiates a new UpdateProjectAppCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectAppId

`func (o *UpdateProjectAppCommand) GetProjectAppId() int32`

GetProjectAppId returns the ProjectAppId field if non-nil, zero value otherwise.

### GetProjectAppIdOk

`func (o *UpdateProjectAppCommand) GetProjectAppIdOk() (*int32, bool)`

GetProjectAppIdOk returns a tuple with the ProjectAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAppId

`func (o *UpdateProjectAppCommand) SetProjectAppId(v int32)`

SetProjectAppId sets ProjectAppId field to given value.


### GetStatus

`func (o *UpdateProjectAppCommand) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateProjectAppCommand) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateProjectAppCommand) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *UpdateProjectAppCommand) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateProjectAppCommand) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


