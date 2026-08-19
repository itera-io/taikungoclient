# OrganizationsWithGroupInfoResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **int32** |  | 
**OrgName** | **string** |  | 
**IsBound** | **bool** |  | 
**AccessLevel** | **NullableString** |  | 
**Projects** | [**[]OrganizationDropdownProjectDto**](OrganizationDropdownProjectDto.md) |  | 

## Methods

### NewOrganizationsWithGroupInfoResultDto

`func NewOrganizationsWithGroupInfoResultDto(orgId int32, orgName string, isBound bool, accessLevel NullableString, projects []OrganizationDropdownProjectDto, ) *OrganizationsWithGroupInfoResultDto`

NewOrganizationsWithGroupInfoResultDto instantiates a new OrganizationsWithGroupInfoResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsWithGroupInfoResultDtoWithDefaults

`func NewOrganizationsWithGroupInfoResultDtoWithDefaults() *OrganizationsWithGroupInfoResultDto`

NewOrganizationsWithGroupInfoResultDtoWithDefaults instantiates a new OrganizationsWithGroupInfoResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrganizationsWithGroupInfoResultDto) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrganizationsWithGroupInfoResultDto) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrganizationsWithGroupInfoResultDto) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *OrganizationsWithGroupInfoResultDto) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrganizationsWithGroupInfoResultDto) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrganizationsWithGroupInfoResultDto) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetIsBound

`func (o *OrganizationsWithGroupInfoResultDto) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *OrganizationsWithGroupInfoResultDto) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *OrganizationsWithGroupInfoResultDto) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.


### GetAccessLevel

`func (o *OrganizationsWithGroupInfoResultDto) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *OrganizationsWithGroupInfoResultDto) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *OrganizationsWithGroupInfoResultDto) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### SetAccessLevelNil

`func (o *OrganizationsWithGroupInfoResultDto) SetAccessLevelNil(b bool)`

 SetAccessLevelNil sets the value for AccessLevel to be an explicit nil

### UnsetAccessLevel
`func (o *OrganizationsWithGroupInfoResultDto) UnsetAccessLevel()`

UnsetAccessLevel ensures that no value is present for AccessLevel, not even an explicit nil
### GetProjects

`func (o *OrganizationsWithGroupInfoResultDto) GetProjects() []OrganizationDropdownProjectDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationsWithGroupInfoResultDto) GetProjectsOk() (*[]OrganizationDropdownProjectDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationsWithGroupInfoResultDto) SetProjects(v []OrganizationDropdownProjectDto)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


