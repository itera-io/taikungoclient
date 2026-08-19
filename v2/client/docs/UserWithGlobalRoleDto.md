# UserWithGlobalRoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Username** | **NullableString** |  | 
**Role** | [**EGlobalRole**](EGlobalRole.md) |  | 

## Methods

### NewUserWithGlobalRoleDto

`func NewUserWithGlobalRoleDto(id NullableString, username NullableString, role EGlobalRole, ) *UserWithGlobalRoleDto`

NewUserWithGlobalRoleDto instantiates a new UserWithGlobalRoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithGlobalRoleDtoWithDefaults

`func NewUserWithGlobalRoleDtoWithDefaults() *UserWithGlobalRoleDto`

NewUserWithGlobalRoleDtoWithDefaults instantiates a new UserWithGlobalRoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserWithGlobalRoleDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWithGlobalRoleDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWithGlobalRoleDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UserWithGlobalRoleDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserWithGlobalRoleDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUsername

`func (o *UserWithGlobalRoleDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserWithGlobalRoleDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserWithGlobalRoleDto) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *UserWithGlobalRoleDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserWithGlobalRoleDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetRole

`func (o *UserWithGlobalRoleDto) GetRole() EGlobalRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserWithGlobalRoleDto) GetRoleOk() (*EGlobalRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserWithGlobalRoleDto) SetRole(v EGlobalRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


