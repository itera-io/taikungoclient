# NodeActionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **NullableString** |  | 
**Action** | [**ENodeAction**](ENodeAction.md) |  | 

## Methods

### NewNodeActionCommand

`func NewNodeActionCommand(projectId int32, name NullableString, action ENodeAction, ) *NodeActionCommand`

NewNodeActionCommand instantiates a new NodeActionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeActionCommandWithDefaults

`func NewNodeActionCommandWithDefaults() *NodeActionCommand`

NewNodeActionCommandWithDefaults instantiates a new NodeActionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *NodeActionCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NodeActionCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NodeActionCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *NodeActionCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeActionCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeActionCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *NodeActionCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NodeActionCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAction

`func (o *NodeActionCommand) GetAction() ENodeAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NodeActionCommand) GetActionOk() (*ENodeAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NodeActionCommand) SetAction(v ENodeAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


