# DaemonsetActionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Action** | [**EDaemonSetAction**](EDaemonSetAction.md) |  | 

## Methods

### NewDaemonsetActionCommand

`func NewDaemonsetActionCommand(projectId int32, name NullableString, namespace NullableString, action EDaemonSetAction, ) *DaemonsetActionCommand`

NewDaemonsetActionCommand instantiates a new DaemonsetActionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonsetActionCommandWithDefaults

`func NewDaemonsetActionCommandWithDefaults() *DaemonsetActionCommand`

NewDaemonsetActionCommandWithDefaults instantiates a new DaemonsetActionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DaemonsetActionCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DaemonsetActionCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DaemonsetActionCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DaemonsetActionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaemonsetActionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaemonsetActionCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DaemonsetActionCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DaemonsetActionCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *DaemonsetActionCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DaemonsetActionCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DaemonsetActionCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *DaemonsetActionCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DaemonsetActionCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetAction

`func (o *DaemonsetActionCommand) GetAction() EDaemonSetAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DaemonsetActionCommand) GetActionOk() (*EDaemonSetAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DaemonsetActionCommand) SetAction(v EDaemonSetAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


