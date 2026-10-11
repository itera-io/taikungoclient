# DownloadKubernetesResourceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Kind** | [**EKubernetesResource**](EKubernetesResource.md) |  | 

## Methods

### NewDownloadKubernetesResourceCommand

`func NewDownloadKubernetesResourceCommand(projectId int32, name string, kind EKubernetesResource, ) *DownloadKubernetesResourceCommand`

NewDownloadKubernetesResourceCommand instantiates a new DownloadKubernetesResourceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadKubernetesResourceCommandWithDefaults

`func NewDownloadKubernetesResourceCommandWithDefaults() *DownloadKubernetesResourceCommand`

NewDownloadKubernetesResourceCommandWithDefaults instantiates a new DownloadKubernetesResourceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DownloadKubernetesResourceCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DownloadKubernetesResourceCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DownloadKubernetesResourceCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DownloadKubernetesResourceCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadKubernetesResourceCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadKubernetesResourceCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DownloadKubernetesResourceCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DownloadKubernetesResourceCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DownloadKubernetesResourceCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DownloadKubernetesResourceCommand) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *DownloadKubernetesResourceCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DownloadKubernetesResourceCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetKind

`func (o *DownloadKubernetesResourceCommand) GetKind() EKubernetesResource`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DownloadKubernetesResourceCommand) GetKindOk() (*EKubernetesResource, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DownloadKubernetesResourceCommand) SetKind(v EKubernetesResource)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


