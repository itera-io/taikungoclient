# PatchKubernetesResourceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Yaml** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchKubernetesResourceCommand

`func NewPatchKubernetesResourceCommand(projectId int32, yaml string, name string, ) *PatchKubernetesResourceCommand`

NewPatchKubernetesResourceCommand instantiates a new PatchKubernetesResourceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchKubernetesResourceCommandWithDefaults

`func NewPatchKubernetesResourceCommandWithDefaults() *PatchKubernetesResourceCommand`

NewPatchKubernetesResourceCommandWithDefaults instantiates a new PatchKubernetesResourceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PatchKubernetesResourceCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PatchKubernetesResourceCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PatchKubernetesResourceCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetYaml

`func (o *PatchKubernetesResourceCommand) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *PatchKubernetesResourceCommand) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *PatchKubernetesResourceCommand) SetYaml(v string)`

SetYaml sets Yaml field to given value.


### GetName

`func (o *PatchKubernetesResourceCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchKubernetesResourceCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchKubernetesResourceCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PatchKubernetesResourceCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchKubernetesResourceCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchKubernetesResourceCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchKubernetesResourceCommand) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *PatchKubernetesResourceCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *PatchKubernetesResourceCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


