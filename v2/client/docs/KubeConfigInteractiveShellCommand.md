# KubeConfigInteractiveShellCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfigId** | **int32** |  | 
**Token** | **NullableString** |  | 
**ProjectId** | **int32** |  | 

## Methods

### NewKubeConfigInteractiveShellCommand

`func NewKubeConfigInteractiveShellCommand(kubeConfigId int32, token NullableString, projectId int32, ) *KubeConfigInteractiveShellCommand`

NewKubeConfigInteractiveShellCommand instantiates a new KubeConfigInteractiveShellCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigInteractiveShellCommandWithDefaults

`func NewKubeConfigInteractiveShellCommandWithDefaults() *KubeConfigInteractiveShellCommand`

NewKubeConfigInteractiveShellCommandWithDefaults instantiates a new KubeConfigInteractiveShellCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfigId

`func (o *KubeConfigInteractiveShellCommand) GetKubeConfigId() int32`

GetKubeConfigId returns the KubeConfigId field if non-nil, zero value otherwise.

### GetKubeConfigIdOk

`func (o *KubeConfigInteractiveShellCommand) GetKubeConfigIdOk() (*int32, bool)`

GetKubeConfigIdOk returns a tuple with the KubeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigId

`func (o *KubeConfigInteractiveShellCommand) SetKubeConfigId(v int32)`

SetKubeConfigId sets KubeConfigId field to given value.


### GetToken

`func (o *KubeConfigInteractiveShellCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KubeConfigInteractiveShellCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KubeConfigInteractiveShellCommand) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *KubeConfigInteractiveShellCommand) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *KubeConfigInteractiveShellCommand) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetProjectId

`func (o *KubeConfigInteractiveShellCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubeConfigInteractiveShellCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubeConfigInteractiveShellCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


