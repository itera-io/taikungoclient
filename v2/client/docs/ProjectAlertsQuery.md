# ProjectAlertsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Mode** | [**ProjectType**](ProjectType.md) |  | 

## Methods

### NewProjectAlertsQuery

`func NewProjectAlertsQuery(projectId int32, mode ProjectType, ) *ProjectAlertsQuery`

NewProjectAlertsQuery instantiates a new ProjectAlertsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAlertsQueryWithDefaults

`func NewProjectAlertsQueryWithDefaults() *ProjectAlertsQuery`

NewProjectAlertsQueryWithDefaults instantiates a new ProjectAlertsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectAlertsQuery) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectAlertsQuery) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectAlertsQuery) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetMode

`func (o *ProjectAlertsQuery) GetMode() ProjectType`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectAlertsQuery) GetModeOk() (*ProjectType, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectAlertsQuery) SetMode(v ProjectType)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


