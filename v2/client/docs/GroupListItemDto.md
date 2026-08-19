# GroupListItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** |  | 
**GroupName** | **string** |  | 
**DomainId** | **int32** |  | 
**OrganizationsCount** | **int32** |  | 
**UsersCount** | **int32** |  | 
**Roles** | [**[]ENonGlobalRoles**](ENonGlobalRoles.md) |  | 

## Methods

### NewGroupListItemDto

`func NewGroupListItemDto(groupId int32, groupName string, domainId int32, organizationsCount int32, usersCount int32, roles []ENonGlobalRoles, ) *GroupListItemDto`

NewGroupListItemDto instantiates a new GroupListItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupListItemDtoWithDefaults

`func NewGroupListItemDtoWithDefaults() *GroupListItemDto`

NewGroupListItemDtoWithDefaults instantiates a new GroupListItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupListItemDto) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupListItemDto) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupListItemDto) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *GroupListItemDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupListItemDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupListItemDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDomainId

`func (o *GroupListItemDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GroupListItemDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GroupListItemDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetOrganizationsCount

`func (o *GroupListItemDto) GetOrganizationsCount() int32`

GetOrganizationsCount returns the OrganizationsCount field if non-nil, zero value otherwise.

### GetOrganizationsCountOk

`func (o *GroupListItemDto) GetOrganizationsCountOk() (*int32, bool)`

GetOrganizationsCountOk returns a tuple with the OrganizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsCount

`func (o *GroupListItemDto) SetOrganizationsCount(v int32)`

SetOrganizationsCount sets OrganizationsCount field to given value.


### GetUsersCount

`func (o *GroupListItemDto) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *GroupListItemDto) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *GroupListItemDto) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.


### GetRoles

`func (o *GroupListItemDto) GetRoles() []ENonGlobalRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupListItemDto) GetRolesOk() (*[]ENonGlobalRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupListItemDto) SetRoles(v []ENonGlobalRoles)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


