# GroupedOrganizationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Role** | [**ENonGlobalRoles**](ENonGlobalRoles.md) |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 

## Methods

### NewGroupedOrganizationDto

`func NewGroupedOrganizationDto(id int32, name NullableString, role ENonGlobalRoles, projects []CommonDropdownDto, ) *GroupedOrganizationDto`

NewGroupedOrganizationDto instantiates a new GroupedOrganizationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedOrganizationDtoWithDefaults

`func NewGroupedOrganizationDtoWithDefaults() *GroupedOrganizationDto`

NewGroupedOrganizationDtoWithDefaults instantiates a new GroupedOrganizationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupedOrganizationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupedOrganizationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupedOrganizationDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GroupedOrganizationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupedOrganizationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupedOrganizationDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupedOrganizationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupedOrganizationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *GroupedOrganizationDto) GetRole() ENonGlobalRoles`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GroupedOrganizationDto) GetRoleOk() (*ENonGlobalRoles, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GroupedOrganizationDto) SetRole(v ENonGlobalRoles)`

SetRole sets Role field to given value.


### GetProjects

`func (o *GroupedOrganizationDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GroupedOrganizationDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GroupedOrganizationDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### SetProjectsNil

`func (o *GroupedOrganizationDto) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *GroupedOrganizationDto) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


