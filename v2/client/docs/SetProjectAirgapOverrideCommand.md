# SetProjectAirgapOverrideCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**OciRegistryHost** | **NullableString** |  | 
**OciPlainHttp** | **bool** |  | 

## Methods

### NewSetProjectAirgapOverrideCommand

`func NewSetProjectAirgapOverrideCommand(projectId int32, ociRegistryHost NullableString, ociPlainHttp bool, ) *SetProjectAirgapOverrideCommand`

NewSetProjectAirgapOverrideCommand instantiates a new SetProjectAirgapOverrideCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetProjectAirgapOverrideCommandWithDefaults

`func NewSetProjectAirgapOverrideCommandWithDefaults() *SetProjectAirgapOverrideCommand`

NewSetProjectAirgapOverrideCommandWithDefaults instantiates a new SetProjectAirgapOverrideCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SetProjectAirgapOverrideCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetProjectAirgapOverrideCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetProjectAirgapOverrideCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetOciRegistryHost

`func (o *SetProjectAirgapOverrideCommand) GetOciRegistryHost() string`

GetOciRegistryHost returns the OciRegistryHost field if non-nil, zero value otherwise.

### GetOciRegistryHostOk

`func (o *SetProjectAirgapOverrideCommand) GetOciRegistryHostOk() (*string, bool)`

GetOciRegistryHostOk returns a tuple with the OciRegistryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegistryHost

`func (o *SetProjectAirgapOverrideCommand) SetOciRegistryHost(v string)`

SetOciRegistryHost sets OciRegistryHost field to given value.


### SetOciRegistryHostNil

`func (o *SetProjectAirgapOverrideCommand) SetOciRegistryHostNil(b bool)`

 SetOciRegistryHostNil sets the value for OciRegistryHost to be an explicit nil

### UnsetOciRegistryHost
`func (o *SetProjectAirgapOverrideCommand) UnsetOciRegistryHost()`

UnsetOciRegistryHost ensures that no value is present for OciRegistryHost, not even an explicit nil
### GetOciPlainHttp

`func (o *SetProjectAirgapOverrideCommand) GetOciPlainHttp() bool`

GetOciPlainHttp returns the OciPlainHttp field if non-nil, zero value otherwise.

### GetOciPlainHttpOk

`func (o *SetProjectAirgapOverrideCommand) GetOciPlainHttpOk() (*bool, bool)`

GetOciPlainHttpOk returns a tuple with the OciPlainHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciPlainHttp

`func (o *SetProjectAirgapOverrideCommand) SetOciPlainHttp(v bool)`

SetOciPlainHttp sets OciPlainHttp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


