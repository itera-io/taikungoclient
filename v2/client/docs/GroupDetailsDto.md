# GroupDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**ClaimName** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**Organizations** | [**[]GroupedOrganizationDto**](GroupedOrganizationDto.md) |  | 
**Users** | [**[]GroupedUserDto**](GroupedUserDto.md) |  | 

## Methods

### NewGroupDetailsDto

`func NewGroupDetailsDto(id int32, name NullableString, claimName NullableString, createdAt NullableString, organizations []GroupedOrganizationDto, users []GroupedUserDto, ) *GroupDetailsDto`

NewGroupDetailsDto instantiates a new GroupDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDetailsDtoWithDefaults

`func NewGroupDetailsDtoWithDefaults() *GroupDetailsDto`

NewGroupDetailsDtoWithDefaults instantiates a new GroupDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GroupDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClaimName

`func (o *GroupDetailsDto) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *GroupDetailsDto) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *GroupDetailsDto) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### SetClaimNameNil

`func (o *GroupDetailsDto) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *GroupDetailsDto) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetCreatedAt

`func (o *GroupDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GroupDetailsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupDetailsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetOrganizations

`func (o *GroupDetailsDto) GetOrganizations() []GroupedOrganizationDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GroupDetailsDto) GetOrganizationsOk() (*[]GroupedOrganizationDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GroupDetailsDto) SetOrganizations(v []GroupedOrganizationDto)`

SetOrganizations sets Organizations field to given value.


### SetOrganizationsNil

`func (o *GroupDetailsDto) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *GroupDetailsDto) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil
### GetUsers

`func (o *GroupDetailsDto) GetUsers() []GroupedUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupDetailsDto) GetUsersOk() (*[]GroupedUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupDetailsDto) SetUsers(v []GroupedUserDto)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *GroupDetailsDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *GroupDetailsDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


