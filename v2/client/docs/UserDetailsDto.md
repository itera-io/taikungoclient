# UserDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Email** | **string** |  | 
**Organizations** | [**[]UserOrganizationDto**](UserOrganizationDto.md) |  | 
**GlobalRole** | [**EGlobalRole**](EGlobalRole.md) |  | 
**CreatedAt** | **NullableString** |  | 
**Is2FAEnabled** | **bool** |  | 

## Methods

### NewUserDetailsDto

`func NewUserDetailsDto(id string, name string, displayName string, email string, organizations []UserOrganizationDto, globalRole EGlobalRole, createdAt NullableString, is2FAEnabled bool, ) *UserDetailsDto`

NewUserDetailsDto instantiates a new UserDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailsDtoWithDefaults

`func NewUserDetailsDtoWithDefaults() *UserDetailsDto`

NewUserDetailsDtoWithDefaults instantiates a new UserDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetailsDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *UserDetailsDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDetailsDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDetailsDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *UserDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizations

`func (o *UserDetailsDto) GetOrganizations() []UserOrganizationDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserDetailsDto) GetOrganizationsOk() (*[]UserOrganizationDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserDetailsDto) SetOrganizations(v []UserOrganizationDto)`

SetOrganizations sets Organizations field to given value.


### GetGlobalRole

`func (o *UserDetailsDto) GetGlobalRole() EGlobalRole`

GetGlobalRole returns the GlobalRole field if non-nil, zero value otherwise.

### GetGlobalRoleOk

`func (o *UserDetailsDto) GetGlobalRoleOk() (*EGlobalRole, bool)`

GetGlobalRoleOk returns a tuple with the GlobalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRole

`func (o *UserDetailsDto) SetGlobalRole(v EGlobalRole)`

SetGlobalRole sets GlobalRole field to given value.


### GetCreatedAt

`func (o *UserDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *UserDetailsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UserDetailsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIs2FAEnabled

`func (o *UserDetailsDto) GetIs2FAEnabled() bool`

GetIs2FAEnabled returns the Is2FAEnabled field if non-nil, zero value otherwise.

### GetIs2FAEnabledOk

`func (o *UserDetailsDto) GetIs2FAEnabledOk() (*bool, bool)`

GetIs2FAEnabledOk returns a tuple with the Is2FAEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2FAEnabled

`func (o *UserDetailsDto) SetIs2FAEnabled(v bool)`

SetIs2FAEnabled sets Is2FAEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


