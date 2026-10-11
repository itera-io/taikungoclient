# CronjobActionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Data** | [**[]KubernetesCronjobActionRequest**](KubernetesCronjobActionRequest.md) |  | 
**Action** | [**ECronJobAction**](ECronJobAction.md) |  | 

## Methods

### NewCronjobActionCommand

`func NewCronjobActionCommand(projectId int32, data []KubernetesCronjobActionRequest, action ECronJobAction, ) *CronjobActionCommand`

NewCronjobActionCommand instantiates a new CronjobActionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronjobActionCommandWithDefaults

`func NewCronjobActionCommandWithDefaults() *CronjobActionCommand`

NewCronjobActionCommandWithDefaults instantiates a new CronjobActionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CronjobActionCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CronjobActionCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CronjobActionCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetData

`func (o *CronjobActionCommand) GetData() []KubernetesCronjobActionRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CronjobActionCommand) GetDataOk() (*[]KubernetesCronjobActionRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CronjobActionCommand) SetData(v []KubernetesCronjobActionRequest)`

SetData sets Data field to given value.


### SetDataNil

`func (o *CronjobActionCommand) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CronjobActionCommand) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetAction

`func (o *CronjobActionCommand) GetAction() ECronJobAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CronjobActionCommand) GetActionOk() (*ECronJobAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CronjobActionCommand) SetAction(v ECronJobAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


