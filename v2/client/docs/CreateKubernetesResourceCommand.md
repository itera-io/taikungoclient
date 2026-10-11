# CreateKubernetesResourceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Yaml** | **NullableString** |  | 

## Methods

### NewCreateKubernetesResourceCommand

`func NewCreateKubernetesResourceCommand(projectId int32, yaml NullableString, ) *CreateKubernetesResourceCommand`

NewCreateKubernetesResourceCommand instantiates a new CreateKubernetesResourceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKubernetesResourceCommandWithDefaults

`func NewCreateKubernetesResourceCommandWithDefaults() *CreateKubernetesResourceCommand`

NewCreateKubernetesResourceCommandWithDefaults instantiates a new CreateKubernetesResourceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateKubernetesResourceCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateKubernetesResourceCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateKubernetesResourceCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetYaml

`func (o *CreateKubernetesResourceCommand) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *CreateKubernetesResourceCommand) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *CreateKubernetesResourceCommand) SetYaml(v string)`

SetYaml sets Yaml field to given value.


### SetYamlNil

`func (o *CreateKubernetesResourceCommand) SetYamlNil(b bool)`

 SetYamlNil sets the value for Yaml to be an explicit nil

### UnsetYaml
`func (o *CreateKubernetesResourceCommand) UnsetYaml()`

UnsetYaml ensures that no value is present for Yaml, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


