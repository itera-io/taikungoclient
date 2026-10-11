# DomainList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**OrganizationsCount** | Pointer to **int64** |  | [optional] 
**UsersCount** | Pointer to **int64** |  | [optional] 
**GroupsCount** | Pointer to **int64** |  | [optional] 
**ProjectsCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewDomainList

`func NewDomainList(id int32, name NullableString, ) *DomainList`

NewDomainList instantiates a new DomainList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainListWithDefaults

`func NewDomainListWithDefaults() *DomainList`

NewDomainListWithDefaults instantiates a new DomainList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainList) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DomainList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainList) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DomainList) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DomainList) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationsCount

`func (o *DomainList) GetOrganizationsCount() int64`

GetOrganizationsCount returns the OrganizationsCount field if non-nil, zero value otherwise.

### GetOrganizationsCountOk

`func (o *DomainList) GetOrganizationsCountOk() (*int64, bool)`

GetOrganizationsCountOk returns a tuple with the OrganizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsCount

`func (o *DomainList) SetOrganizationsCount(v int64)`

SetOrganizationsCount sets OrganizationsCount field to given value.

### HasOrganizationsCount

`func (o *DomainList) HasOrganizationsCount() bool`

HasOrganizationsCount returns a boolean if a field has been set.

### GetUsersCount

`func (o *DomainList) GetUsersCount() int64`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *DomainList) GetUsersCountOk() (*int64, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *DomainList) SetUsersCount(v int64)`

SetUsersCount sets UsersCount field to given value.

### HasUsersCount

`func (o *DomainList) HasUsersCount() bool`

HasUsersCount returns a boolean if a field has been set.

### GetGroupsCount

`func (o *DomainList) GetGroupsCount() int64`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *DomainList) GetGroupsCountOk() (*int64, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *DomainList) SetGroupsCount(v int64)`

SetGroupsCount sets GroupsCount field to given value.

### HasGroupsCount

`func (o *DomainList) HasGroupsCount() bool`

HasGroupsCount returns a boolean if a field has been set.

### GetProjectsCount

`func (o *DomainList) GetProjectsCount() int64`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *DomainList) GetProjectsCountOk() (*int64, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *DomainList) SetProjectsCount(v int64)`

SetProjectsCount sets ProjectsCount field to given value.

### HasProjectsCount

`func (o *DomainList) HasProjectsCount() bool`

HasProjectsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


