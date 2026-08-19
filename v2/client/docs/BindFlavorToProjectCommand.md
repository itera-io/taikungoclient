# BindFlavorToProjectCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Flavors** | **[]string** |  | 

## Methods

### NewBindFlavorToProjectCommand

`func NewBindFlavorToProjectCommand(projectId int32, flavors []string, ) *BindFlavorToProjectCommand`

NewBindFlavorToProjectCommand instantiates a new BindFlavorToProjectCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindFlavorToProjectCommandWithDefaults

`func NewBindFlavorToProjectCommandWithDefaults() *BindFlavorToProjectCommand`

NewBindFlavorToProjectCommandWithDefaults instantiates a new BindFlavorToProjectCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *BindFlavorToProjectCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BindFlavorToProjectCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BindFlavorToProjectCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetFlavors

`func (o *BindFlavorToProjectCommand) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *BindFlavorToProjectCommand) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *BindFlavorToProjectCommand) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.


### SetFlavorsNil

`func (o *BindFlavorToProjectCommand) SetFlavorsNil(b bool)`

 SetFlavorsNil sets the value for Flavors to be an explicit nil

### UnsetFlavors
`func (o *BindFlavorToProjectCommand) UnsetFlavors()`

UnsetFlavors ensures that no value is present for Flavors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


