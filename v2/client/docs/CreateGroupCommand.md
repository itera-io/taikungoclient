# CreateGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ClaimValue** | Pointer to **NullableString** |  | [optional] 
**DomainId** | **int32** |  | 
**Organizations** | Pointer to [**[]CreateGroupOrganizationDto**](CreateGroupOrganizationDto.md) |  | [optional] 
**Users** | Pointer to [**[]CreateGroupUserDto**](CreateGroupUserDto.md) |  | [optional] 

## Methods

### NewCreateGroupCommand

`func NewCreateGroupCommand(name string, domainId int32, ) *CreateGroupCommand`

NewCreateGroupCommand instantiates a new CreateGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupCommandWithDefaults

`func NewCreateGroupCommandWithDefaults() *CreateGroupCommand`

NewCreateGroupCommandWithDefaults instantiates a new CreateGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupCommand) SetName(v string)`

SetName sets Name field to given value.


### GetClaimValue

`func (o *CreateGroupCommand) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *CreateGroupCommand) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *CreateGroupCommand) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *CreateGroupCommand) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *CreateGroupCommand) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *CreateGroupCommand) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil
### GetDomainId

`func (o *CreateGroupCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateGroupCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateGroupCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetOrganizations

`func (o *CreateGroupCommand) GetOrganizations() []CreateGroupOrganizationDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *CreateGroupCommand) GetOrganizationsOk() (*[]CreateGroupOrganizationDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *CreateGroupCommand) SetOrganizations(v []CreateGroupOrganizationDto)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *CreateGroupCommand) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### SetOrganizationsNil

`func (o *CreateGroupCommand) SetOrganizationsNil(b bool)`

 SetOrganizationsNil sets the value for Organizations to be an explicit nil

### UnsetOrganizations
`func (o *CreateGroupCommand) UnsetOrganizations()`

UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil
### GetUsers

`func (o *CreateGroupCommand) GetUsers() []CreateGroupUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateGroupCommand) GetUsersOk() (*[]CreateGroupUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateGroupCommand) SetUsers(v []CreateGroupUserDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateGroupCommand) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *CreateGroupCommand) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *CreateGroupCommand) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


