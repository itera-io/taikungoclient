# DuplicateProjectAppCheckerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 

## Methods

### NewDuplicateProjectAppCheckerCommand

`func NewDuplicateProjectAppCheckerCommand(projectId int32, name NullableString, namespace NullableString, ) *DuplicateProjectAppCheckerCommand`

NewDuplicateProjectAppCheckerCommand instantiates a new DuplicateProjectAppCheckerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateProjectAppCheckerCommandWithDefaults

`func NewDuplicateProjectAppCheckerCommandWithDefaults() *DuplicateProjectAppCheckerCommand`

NewDuplicateProjectAppCheckerCommandWithDefaults instantiates a new DuplicateProjectAppCheckerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DuplicateProjectAppCheckerCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DuplicateProjectAppCheckerCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DuplicateProjectAppCheckerCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DuplicateProjectAppCheckerCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DuplicateProjectAppCheckerCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DuplicateProjectAppCheckerCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DuplicateProjectAppCheckerCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DuplicateProjectAppCheckerCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *DuplicateProjectAppCheckerCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DuplicateProjectAppCheckerCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DuplicateProjectAppCheckerCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *DuplicateProjectAppCheckerCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DuplicateProjectAppCheckerCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


