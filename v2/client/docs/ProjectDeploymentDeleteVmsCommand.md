# ProjectDeploymentDeleteVmsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**VmIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewProjectDeploymentDeleteVmsCommand

`func NewProjectDeploymentDeleteVmsCommand() *ProjectDeploymentDeleteVmsCommand`

NewProjectDeploymentDeleteVmsCommand instantiates a new ProjectDeploymentDeleteVmsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentDeleteVmsCommandWithDefaults

`func NewProjectDeploymentDeleteVmsCommandWithDefaults() *ProjectDeploymentDeleteVmsCommand`

NewProjectDeploymentDeleteVmsCommandWithDefaults instantiates a new ProjectDeploymentDeleteVmsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectDeploymentDeleteVmsCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectDeploymentDeleteVmsCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectDeploymentDeleteVmsCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectDeploymentDeleteVmsCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVmIds

`func (o *ProjectDeploymentDeleteVmsCommand) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *ProjectDeploymentDeleteVmsCommand) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *ProjectDeploymentDeleteVmsCommand) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *ProjectDeploymentDeleteVmsCommand) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIdsNil

`func (o *ProjectDeploymentDeleteVmsCommand) SetVmIdsNil(b bool)`

 SetVmIdsNil sets the value for VmIds to be an explicit nil

### UnsetVmIds
`func (o *ProjectDeploymentDeleteVmsCommand) UnsetVmIds()`

UnsetVmIds ensures that no value is present for VmIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


