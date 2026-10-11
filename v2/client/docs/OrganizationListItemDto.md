# OrganizationListItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**DomainId** | **int32** |  | 
**ProjectsCount** | **int32** |  | 
**GroupsCount** | **int32** |  | 
**UsersCount** | **int32** |  | 

## Methods

### NewOrganizationListItemDto

`func NewOrganizationListItemDto(organizationId int32, organizationName NullableString, domainId int32, projectsCount int32, groupsCount int32, usersCount int32, ) *OrganizationListItemDto`

NewOrganizationListItemDto instantiates a new OrganizationListItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationListItemDtoWithDefaults

`func NewOrganizationListItemDtoWithDefaults() *OrganizationListItemDto`

NewOrganizationListItemDtoWithDefaults instantiates a new OrganizationListItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *OrganizationListItemDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationListItemDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationListItemDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *OrganizationListItemDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationListItemDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationListItemDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *OrganizationListItemDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OrganizationListItemDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetDomainId

`func (o *OrganizationListItemDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *OrganizationListItemDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *OrganizationListItemDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetProjectsCount

`func (o *OrganizationListItemDto) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *OrganizationListItemDto) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *OrganizationListItemDto) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.


### GetGroupsCount

`func (o *OrganizationListItemDto) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *OrganizationListItemDto) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *OrganizationListItemDto) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.


### GetUsersCount

`func (o *OrganizationListItemDto) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *OrganizationListItemDto) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *OrganizationListItemDto) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


