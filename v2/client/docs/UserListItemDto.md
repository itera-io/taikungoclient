# UserListItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** |  | 
**UserName** | **NullableString** |  | 
**GlobalRole** | [**EGlobalRole**](EGlobalRole.md) |  | 
**NonGlobalRoles** | [**[]ENonGlobalRoles**](ENonGlobalRoles.md) |  | 
**GroupsCount** | **int32** |  | 
**OrganizationsCount** | **int32** |  | 
**ProjectsCount** | **int32** |  | 

## Methods

### NewUserListItemDto

`func NewUserListItemDto(userId NullableString, userName NullableString, globalRole EGlobalRole, nonGlobalRoles []ENonGlobalRoles, groupsCount int32, organizationsCount int32, projectsCount int32, ) *UserListItemDto`

NewUserListItemDto instantiates a new UserListItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListItemDtoWithDefaults

`func NewUserListItemDtoWithDefaults() *UserListItemDto`

NewUserListItemDtoWithDefaults instantiates a new UserListItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserListItemDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserListItemDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserListItemDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *UserListItemDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserListItemDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *UserListItemDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserListItemDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserListItemDto) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *UserListItemDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserListItemDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetGlobalRole

`func (o *UserListItemDto) GetGlobalRole() EGlobalRole`

GetGlobalRole returns the GlobalRole field if non-nil, zero value otherwise.

### GetGlobalRoleOk

`func (o *UserListItemDto) GetGlobalRoleOk() (*EGlobalRole, bool)`

GetGlobalRoleOk returns a tuple with the GlobalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRole

`func (o *UserListItemDto) SetGlobalRole(v EGlobalRole)`

SetGlobalRole sets GlobalRole field to given value.


### GetNonGlobalRoles

`func (o *UserListItemDto) GetNonGlobalRoles() []ENonGlobalRoles`

GetNonGlobalRoles returns the NonGlobalRoles field if non-nil, zero value otherwise.

### GetNonGlobalRolesOk

`func (o *UserListItemDto) GetNonGlobalRolesOk() (*[]ENonGlobalRoles, bool)`

GetNonGlobalRolesOk returns a tuple with the NonGlobalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonGlobalRoles

`func (o *UserListItemDto) SetNonGlobalRoles(v []ENonGlobalRoles)`

SetNonGlobalRoles sets NonGlobalRoles field to given value.


### SetNonGlobalRolesNil

`func (o *UserListItemDto) SetNonGlobalRolesNil(b bool)`

 SetNonGlobalRolesNil sets the value for NonGlobalRoles to be an explicit nil

### UnsetNonGlobalRoles
`func (o *UserListItemDto) UnsetNonGlobalRoles()`

UnsetNonGlobalRoles ensures that no value is present for NonGlobalRoles, not even an explicit nil
### GetGroupsCount

`func (o *UserListItemDto) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *UserListItemDto) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *UserListItemDto) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.


### GetOrganizationsCount

`func (o *UserListItemDto) GetOrganizationsCount() int32`

GetOrganizationsCount returns the OrganizationsCount field if non-nil, zero value otherwise.

### GetOrganizationsCountOk

`func (o *UserListItemDto) GetOrganizationsCountOk() (*int32, bool)`

GetOrganizationsCountOk returns a tuple with the OrganizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsCount

`func (o *UserListItemDto) SetOrganizationsCount(v int32)`

SetOrganizationsCount sets OrganizationsCount field to given value.


### GetProjectsCount

`func (o *UserListItemDto) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *UserListItemDto) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *UserListItemDto) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


