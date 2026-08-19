# DomainDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**OrganizationsCount** | **int64** |  | 
**UsersCount** | **int64** |  | 
**GroupsCount** | **int64** |  | 
**ProjectsCount** | **int64** |  | 
**CreatedAt** | **NullableString** |  | 
**Is2FAEnabled** | **bool** |  | 
**UserWithGlobalRoles** | [**[]UserWithGlobalRoleDto**](UserWithGlobalRoleDto.md) |  | 

## Methods

### NewDomainDetailsDto

`func NewDomainDetailsDto(id int32, name NullableString, email NullableString, organizationsCount int64, usersCount int64, groupsCount int64, projectsCount int64, createdAt NullableString, is2FAEnabled bool, userWithGlobalRoles []UserWithGlobalRoleDto, ) *DomainDetailsDto`

NewDomainDetailsDto instantiates a new DomainDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailsDtoWithDefaults

`func NewDomainDetailsDtoWithDefaults() *DomainDetailsDto`

NewDomainDetailsDtoWithDefaults instantiates a new DomainDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DomainDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DomainDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DomainDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *DomainDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomainDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomainDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *DomainDetailsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *DomainDetailsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOrganizationsCount

`func (o *DomainDetailsDto) GetOrganizationsCount() int64`

GetOrganizationsCount returns the OrganizationsCount field if non-nil, zero value otherwise.

### GetOrganizationsCountOk

`func (o *DomainDetailsDto) GetOrganizationsCountOk() (*int64, bool)`

GetOrganizationsCountOk returns a tuple with the OrganizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsCount

`func (o *DomainDetailsDto) SetOrganizationsCount(v int64)`

SetOrganizationsCount sets OrganizationsCount field to given value.


### GetUsersCount

`func (o *DomainDetailsDto) GetUsersCount() int64`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *DomainDetailsDto) GetUsersCountOk() (*int64, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *DomainDetailsDto) SetUsersCount(v int64)`

SetUsersCount sets UsersCount field to given value.


### GetGroupsCount

`func (o *DomainDetailsDto) GetGroupsCount() int64`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *DomainDetailsDto) GetGroupsCountOk() (*int64, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *DomainDetailsDto) SetGroupsCount(v int64)`

SetGroupsCount sets GroupsCount field to given value.


### GetProjectsCount

`func (o *DomainDetailsDto) GetProjectsCount() int64`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *DomainDetailsDto) GetProjectsCountOk() (*int64, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *DomainDetailsDto) SetProjectsCount(v int64)`

SetProjectsCount sets ProjectsCount field to given value.


### GetCreatedAt

`func (o *DomainDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DomainDetailsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DomainDetailsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIs2FAEnabled

`func (o *DomainDetailsDto) GetIs2FAEnabled() bool`

GetIs2FAEnabled returns the Is2FAEnabled field if non-nil, zero value otherwise.

### GetIs2FAEnabledOk

`func (o *DomainDetailsDto) GetIs2FAEnabledOk() (*bool, bool)`

GetIs2FAEnabledOk returns a tuple with the Is2FAEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2FAEnabled

`func (o *DomainDetailsDto) SetIs2FAEnabled(v bool)`

SetIs2FAEnabled sets Is2FAEnabled field to given value.


### GetUserWithGlobalRoles

`func (o *DomainDetailsDto) GetUserWithGlobalRoles() []UserWithGlobalRoleDto`

GetUserWithGlobalRoles returns the UserWithGlobalRoles field if non-nil, zero value otherwise.

### GetUserWithGlobalRolesOk

`func (o *DomainDetailsDto) GetUserWithGlobalRolesOk() (*[]UserWithGlobalRoleDto, bool)`

GetUserWithGlobalRolesOk returns a tuple with the UserWithGlobalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWithGlobalRoles

`func (o *DomainDetailsDto) SetUserWithGlobalRoles(v []UserWithGlobalRoleDto)`

SetUserWithGlobalRoles sets UserWithGlobalRoles field to given value.


### SetUserWithGlobalRolesNil

`func (o *DomainDetailsDto) SetUserWithGlobalRolesNil(b bool)`

 SetUserWithGlobalRolesNil sets the value for UserWithGlobalRoles to be an explicit nil

### UnsetUserWithGlobalRoles
`func (o *DomainDetailsDto) UnsetUserWithGlobalRoles()`

UnsetUserWithGlobalRoles ensures that no value is present for UserWithGlobalRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


