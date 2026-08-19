# UpdateGroupOrganizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**Projects** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUpdateGroupOrganizationDto

`func NewUpdateGroupOrganizationDto(role string, ) *UpdateGroupOrganizationDto`

NewUpdateGroupOrganizationDto instantiates a new UpdateGroupOrganizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupOrganizationDtoWithDefaults

`func NewUpdateGroupOrganizationDtoWithDefaults() *UpdateGroupOrganizationDto`

NewUpdateGroupOrganizationDtoWithDefaults instantiates a new UpdateGroupOrganizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UpdateGroupOrganizationDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateGroupOrganizationDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateGroupOrganizationDto) SetRole(v string)`

SetRole sets Role field to given value.


### GetProjects

`func (o *UpdateGroupOrganizationDto) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UpdateGroupOrganizationDto) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UpdateGroupOrganizationDto) SetProjects(v []int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UpdateGroupOrganizationDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *UpdateGroupOrganizationDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *UpdateGroupOrganizationDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


