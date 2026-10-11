# UserOrganizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**OrganizationRole** | [**ENonGlobalRoles**](ENonGlobalRoles.md) |  | 
**GroupId** | **int32** |  | 
**GroupName** | **string** |  | 

## Methods

### NewUserOrganizationDto

`func NewUserOrganizationDto(id string, name string, projects []CommonDropdownDto, organizationRole ENonGlobalRoles, groupId int32, groupName string, ) *UserOrganizationDto`

NewUserOrganizationDto instantiates a new UserOrganizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOrganizationDtoWithDefaults

`func NewUserOrganizationDtoWithDefaults() *UserOrganizationDto`

NewUserOrganizationDtoWithDefaults instantiates a new UserOrganizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserOrganizationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserOrganizationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserOrganizationDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserOrganizationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserOrganizationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserOrganizationDto) SetName(v string)`

SetName sets Name field to given value.


### GetProjects

`func (o *UserOrganizationDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UserOrganizationDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UserOrganizationDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetOrganizationRole

`func (o *UserOrganizationDto) GetOrganizationRole() ENonGlobalRoles`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *UserOrganizationDto) GetOrganizationRoleOk() (*ENonGlobalRoles, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *UserOrganizationDto) SetOrganizationRole(v ENonGlobalRoles)`

SetOrganizationRole sets OrganizationRole field to given value.


### GetGroupId

`func (o *UserOrganizationDto) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserOrganizationDto) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserOrganizationDto) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *UserOrganizationDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UserOrganizationDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UserOrganizationDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


