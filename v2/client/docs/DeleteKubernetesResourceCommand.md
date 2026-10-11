# DeleteKubernetesResourceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Kind** | [**EKubernetesResource**](EKubernetesResource.md) |  | 
**Data** | [**[]KubernetesActionRequest**](KubernetesActionRequest.md) |  | 

## Methods

### NewDeleteKubernetesResourceCommand

`func NewDeleteKubernetesResourceCommand(projectId int32, kind EKubernetesResource, data []KubernetesActionRequest, ) *DeleteKubernetesResourceCommand`

NewDeleteKubernetesResourceCommand instantiates a new DeleteKubernetesResourceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteKubernetesResourceCommandWithDefaults

`func NewDeleteKubernetesResourceCommandWithDefaults() *DeleteKubernetesResourceCommand`

NewDeleteKubernetesResourceCommandWithDefaults instantiates a new DeleteKubernetesResourceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteKubernetesResourceCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteKubernetesResourceCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteKubernetesResourceCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetKind

`func (o *DeleteKubernetesResourceCommand) GetKind() EKubernetesResource`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeleteKubernetesResourceCommand) GetKindOk() (*EKubernetesResource, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeleteKubernetesResourceCommand) SetKind(v EKubernetesResource)`

SetKind sets Kind field to given value.


### GetData

`func (o *DeleteKubernetesResourceCommand) GetData() []KubernetesActionRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteKubernetesResourceCommand) GetDataOk() (*[]KubernetesActionRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteKubernetesResourceCommand) SetData(v []KubernetesActionRequest)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


