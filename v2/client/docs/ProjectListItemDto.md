# ProjectListItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**ProjectName** | **NullableString** |  | 
**CloudType** | **map[string]interface{}** |  | 
**DomainId** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**UsersCount** | **int32** |  | 
**GroupsCount** | **int32** |  | 

## Methods

### NewProjectListItemDto

`func NewProjectListItemDto(projectId int32, projectName NullableString, cloudType map[string]interface{}, domainId int32, organizationId int32, organizationName NullableString, status ProjectStatus, health ProjectHealth, usersCount int32, groupsCount int32, ) *ProjectListItemDto`

NewProjectListItemDto instantiates a new ProjectListItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListItemDtoWithDefaults

`func NewProjectListItemDtoWithDefaults() *ProjectListItemDto`

NewProjectListItemDtoWithDefaults instantiates a new ProjectListItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectListItemDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectListItemDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectListItemDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *ProjectListItemDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectListItemDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectListItemDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *ProjectListItemDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectListItemDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetCloudType

`func (o *ProjectListItemDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectListItemDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectListItemDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetDomainId

`func (o *ProjectListItemDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ProjectListItemDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ProjectListItemDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetOrganizationId

`func (o *ProjectListItemDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectListItemDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectListItemDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *ProjectListItemDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectListItemDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectListItemDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *ProjectListItemDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectListItemDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetStatus

`func (o *ProjectListItemDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListItemDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListItemDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ProjectListItemDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectListItemDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectListItemDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetUsersCount

`func (o *ProjectListItemDto) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *ProjectListItemDto) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *ProjectListItemDto) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.


### GetGroupsCount

`func (o *ProjectListItemDto) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *ProjectListItemDto) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *ProjectListItemDto) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


