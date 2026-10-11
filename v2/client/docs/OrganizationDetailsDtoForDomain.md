# OrganizationDetailsDtoForDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**FullName** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**Projects** | [**[]GroupedProjectDto**](GroupedProjectDto.md) |  | 
**Groups** | [**[]GroupedOrganizationDto**](GroupedOrganizationDto.md) |  | 
**Users** | [**[]GroupedOrganizationUserDto**](GroupedOrganizationUserDto.md) |  | 

## Methods

### NewOrganizationDetailsDtoForDomain

`func NewOrganizationDetailsDtoForDomain(id int32, name NullableString, fullName NullableString, createdAt NullableString, projects []GroupedProjectDto, groups []GroupedOrganizationDto, users []GroupedOrganizationUserDto, ) *OrganizationDetailsDtoForDomain`

NewOrganizationDetailsDtoForDomain instantiates a new OrganizationDetailsDtoForDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsDtoForDomainWithDefaults

`func NewOrganizationDetailsDtoForDomainWithDefaults() *OrganizationDetailsDtoForDomain`

NewOrganizationDetailsDtoForDomainWithDefaults instantiates a new OrganizationDetailsDtoForDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDetailsDtoForDomain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDetailsDtoForDomain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDetailsDtoForDomain) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationDetailsDtoForDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDetailsDtoForDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDetailsDtoForDomain) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *OrganizationDetailsDtoForDomain) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationDetailsDtoForDomain) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *OrganizationDetailsDtoForDomain) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationDetailsDtoForDomain) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationDetailsDtoForDomain) SetFullName(v string)`

SetFullName sets FullName field to given value.


### SetFullNameNil

`func (o *OrganizationDetailsDtoForDomain) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationDetailsDtoForDomain) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetCreatedAt

`func (o *OrganizationDetailsDtoForDomain) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDetailsDtoForDomain) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDetailsDtoForDomain) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *OrganizationDetailsDtoForDomain) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OrganizationDetailsDtoForDomain) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetProjects

`func (o *OrganizationDetailsDtoForDomain) GetProjects() []GroupedProjectDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationDetailsDtoForDomain) GetProjectsOk() (*[]GroupedProjectDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationDetailsDtoForDomain) SetProjects(v []GroupedProjectDto)`

SetProjects sets Projects field to given value.


### SetProjectsNil

`func (o *OrganizationDetailsDtoForDomain) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *OrganizationDetailsDtoForDomain) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetGroups

`func (o *OrganizationDetailsDtoForDomain) GetGroups() []GroupedOrganizationDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OrganizationDetailsDtoForDomain) GetGroupsOk() (*[]GroupedOrganizationDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OrganizationDetailsDtoForDomain) SetGroups(v []GroupedOrganizationDto)`

SetGroups sets Groups field to given value.


### SetGroupsNil

`func (o *OrganizationDetailsDtoForDomain) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *OrganizationDetailsDtoForDomain) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetUsers

`func (o *OrganizationDetailsDtoForDomain) GetUsers() []GroupedOrganizationUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OrganizationDetailsDtoForDomain) GetUsersOk() (*[]GroupedOrganizationUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OrganizationDetailsDtoForDomain) SetUsers(v []GroupedOrganizationUserDto)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *OrganizationDetailsDtoForDomain) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *OrganizationDetailsDtoForDomain) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


