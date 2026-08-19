# DeploymentActionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**ScaleReplicaCount** | Pointer to **NullableInt32** |  | [optional] 
**Action** | [**EDeploymentAction**](EDeploymentAction.md) |  | 

## Methods

### NewDeploymentActionCommand

`func NewDeploymentActionCommand(projectId int32, name NullableString, namespace NullableString, action EDeploymentAction, ) *DeploymentActionCommand`

NewDeploymentActionCommand instantiates a new DeploymentActionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentActionCommandWithDefaults

`func NewDeploymentActionCommandWithDefaults() *DeploymentActionCommand`

NewDeploymentActionCommandWithDefaults instantiates a new DeploymentActionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeploymentActionCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeploymentActionCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeploymentActionCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DeploymentActionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentActionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentActionCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeploymentActionCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentActionCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *DeploymentActionCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeploymentActionCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeploymentActionCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *DeploymentActionCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DeploymentActionCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetScaleReplicaCount

`func (o *DeploymentActionCommand) GetScaleReplicaCount() int32`

GetScaleReplicaCount returns the ScaleReplicaCount field if non-nil, zero value otherwise.

### GetScaleReplicaCountOk

`func (o *DeploymentActionCommand) GetScaleReplicaCountOk() (*int32, bool)`

GetScaleReplicaCountOk returns a tuple with the ScaleReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleReplicaCount

`func (o *DeploymentActionCommand) SetScaleReplicaCount(v int32)`

SetScaleReplicaCount sets ScaleReplicaCount field to given value.

### HasScaleReplicaCount

`func (o *DeploymentActionCommand) HasScaleReplicaCount() bool`

HasScaleReplicaCount returns a boolean if a field has been set.

### SetScaleReplicaCountNil

`func (o *DeploymentActionCommand) SetScaleReplicaCountNil(b bool)`

 SetScaleReplicaCountNil sets the value for ScaleReplicaCount to be an explicit nil

### UnsetScaleReplicaCount
`func (o *DeploymentActionCommand) UnsetScaleReplicaCount()`

UnsetScaleReplicaCount ensures that no value is present for ScaleReplicaCount, not even an explicit nil
### GetAction

`func (o *DeploymentActionCommand) GetAction() EDeploymentAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentActionCommand) GetActionOk() (*EDeploymentAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentActionCommand) SetAction(v EDeploymentAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


