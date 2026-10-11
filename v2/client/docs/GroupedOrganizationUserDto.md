# GroupedOrganizationUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**OrganizationRole** | [**ENonGlobalRoles**](ENonGlobalRoles.md) |  | 
**GroupId** | **int32** |  | 
**GroupName** | **NullableString** |  | 

## Methods

### NewGroupedOrganizationUserDto

`func NewGroupedOrganizationUserDto(id NullableString, name NullableString, organizationRole ENonGlobalRoles, groupId int32, groupName NullableString, ) *GroupedOrganizationUserDto`

NewGroupedOrganizationUserDto instantiates a new GroupedOrganizationUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedOrganizationUserDtoWithDefaults

`func NewGroupedOrganizationUserDtoWithDefaults() *GroupedOrganizationUserDto`

NewGroupedOrganizationUserDtoWithDefaults instantiates a new GroupedOrganizationUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupedOrganizationUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupedOrganizationUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupedOrganizationUserDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GroupedOrganizationUserDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GroupedOrganizationUserDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *GroupedOrganizationUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupedOrganizationUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupedOrganizationUserDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupedOrganizationUserDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupedOrganizationUserDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationRole

`func (o *GroupedOrganizationUserDto) GetOrganizationRole() ENonGlobalRoles`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *GroupedOrganizationUserDto) GetOrganizationRoleOk() (*ENonGlobalRoles, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *GroupedOrganizationUserDto) SetOrganizationRole(v ENonGlobalRoles)`

SetOrganizationRole sets OrganizationRole field to given value.


### GetGroupId

`func (o *GroupedOrganizationUserDto) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupedOrganizationUserDto) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupedOrganizationUserDto) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *GroupedOrganizationUserDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupedOrganizationUserDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupedOrganizationUserDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### SetGroupNameNil

`func (o *GroupedOrganizationUserDto) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *GroupedOrganizationUserDto) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


