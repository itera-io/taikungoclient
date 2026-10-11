# CreateGroupOrganizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Role** | [**AccessLevelRoles**](AccessLevelRoles.md) |  | 
**Projects** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCreateGroupOrganizationDto

`func NewCreateGroupOrganizationDto(id int32, role AccessLevelRoles, ) *CreateGroupOrganizationDto`

NewCreateGroupOrganizationDto instantiates a new CreateGroupOrganizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupOrganizationDtoWithDefaults

`func NewCreateGroupOrganizationDtoWithDefaults() *CreateGroupOrganizationDto`

NewCreateGroupOrganizationDtoWithDefaults instantiates a new CreateGroupOrganizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateGroupOrganizationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGroupOrganizationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGroupOrganizationDto) SetId(v int32)`

SetId sets Id field to given value.


### GetRole

`func (o *CreateGroupOrganizationDto) GetRole() AccessLevelRoles`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateGroupOrganizationDto) GetRoleOk() (*AccessLevelRoles, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateGroupOrganizationDto) SetRole(v AccessLevelRoles)`

SetRole sets Role field to given value.


### GetProjects

`func (o *CreateGroupOrganizationDto) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateGroupOrganizationDto) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateGroupOrganizationDto) SetProjects(v []int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateGroupOrganizationDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *CreateGroupOrganizationDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *CreateGroupOrganizationDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


