# ProjectDeploymentDeleteServersCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**ServerIds** | Pointer to **[]int32** |  | [optional] 
**ForceDeleteVClusters** | Pointer to **bool** |  | [optional] 
**DeleteAutoscalingServers** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectDeploymentDeleteServersCommand

`func NewProjectDeploymentDeleteServersCommand() *ProjectDeploymentDeleteServersCommand`

NewProjectDeploymentDeleteServersCommand instantiates a new ProjectDeploymentDeleteServersCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentDeleteServersCommandWithDefaults

`func NewProjectDeploymentDeleteServersCommandWithDefaults() *ProjectDeploymentDeleteServersCommand`

NewProjectDeploymentDeleteServersCommandWithDefaults instantiates a new ProjectDeploymentDeleteServersCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectDeploymentDeleteServersCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectDeploymentDeleteServersCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectDeploymentDeleteServersCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectDeploymentDeleteServersCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServerIds

`func (o *ProjectDeploymentDeleteServersCommand) GetServerIds() []int32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *ProjectDeploymentDeleteServersCommand) GetServerIdsOk() (*[]int32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *ProjectDeploymentDeleteServersCommand) SetServerIds(v []int32)`

SetServerIds sets ServerIds field to given value.

### HasServerIds

`func (o *ProjectDeploymentDeleteServersCommand) HasServerIds() bool`

HasServerIds returns a boolean if a field has been set.

### SetServerIdsNil

`func (o *ProjectDeploymentDeleteServersCommand) SetServerIdsNil(b bool)`

 SetServerIdsNil sets the value for ServerIds to be an explicit nil

### UnsetServerIds
`func (o *ProjectDeploymentDeleteServersCommand) UnsetServerIds()`

UnsetServerIds ensures that no value is present for ServerIds, not even an explicit nil
### GetForceDeleteVClusters

`func (o *ProjectDeploymentDeleteServersCommand) GetForceDeleteVClusters() bool`

GetForceDeleteVClusters returns the ForceDeleteVClusters field if non-nil, zero value otherwise.

### GetForceDeleteVClustersOk

`func (o *ProjectDeploymentDeleteServersCommand) GetForceDeleteVClustersOk() (*bool, bool)`

GetForceDeleteVClustersOk returns a tuple with the ForceDeleteVClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeleteVClusters

`func (o *ProjectDeploymentDeleteServersCommand) SetForceDeleteVClusters(v bool)`

SetForceDeleteVClusters sets ForceDeleteVClusters field to given value.

### HasForceDeleteVClusters

`func (o *ProjectDeploymentDeleteServersCommand) HasForceDeleteVClusters() bool`

HasForceDeleteVClusters returns a boolean if a field has been set.

### GetDeleteAutoscalingServers

`func (o *ProjectDeploymentDeleteServersCommand) GetDeleteAutoscalingServers() bool`

GetDeleteAutoscalingServers returns the DeleteAutoscalingServers field if non-nil, zero value otherwise.

### GetDeleteAutoscalingServersOk

`func (o *ProjectDeploymentDeleteServersCommand) GetDeleteAutoscalingServersOk() (*bool, bool)`

GetDeleteAutoscalingServersOk returns a tuple with the DeleteAutoscalingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAutoscalingServers

`func (o *ProjectDeploymentDeleteServersCommand) SetDeleteAutoscalingServers(v bool)`

SetDeleteAutoscalingServers sets DeleteAutoscalingServers field to given value.

### HasDeleteAutoscalingServers

`func (o *ProjectDeploymentDeleteServersCommand) HasDeleteAutoscalingServers() bool`

HasDeleteAutoscalingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


