# UserForListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Email** | **string** |  | 
**DisplayName** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**Is2FAEnabled** | **bool** |  | 
**Domain** | [**DomainDetailsForUserDto**](DomainDetailsForUserDto.md) |  | 
**Role** | [**EGlobalRole**](EGlobalRole.md) |  | 
**Organizations** | [**map[string]OrganizationContextDto**](OrganizationContextDto.md) |  | 
**IsEmailConfirmed** | **bool** |  | 
**IsEmailNotificationEnabled** | **bool** |  | 
**IsForcedToResetPassword** | **bool** |  | 
**IsLocked** | **bool** |  | 
**Owner** | **bool** |  | 
**IsForcedToEnableTwoFactorAuthentication** | **bool** |  | 

## Methods

### NewUserForListDto

`func NewUserForListDto(id string, username string, email string, displayName NullableString, createdAt string, is2FAEnabled bool, domain DomainDetailsForUserDto, role EGlobalRole, organizations map[string]OrganizationContextDto, isEmailConfirmed bool, isEmailNotificationEnabled bool, isForcedToResetPassword bool, isLocked bool, owner bool, isForcedToEnableTwoFactorAuthentication bool, ) *UserForListDto`

NewUserForListDto instantiates a new UserForListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserForListDtoWithDefaults

`func NewUserForListDtoWithDefaults() *UserForListDto`

NewUserForListDtoWithDefaults instantiates a new UserForListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserForListDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserForListDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserForListDto) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserForListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserForListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserForListDto) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *UserForListDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserForListDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserForListDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *UserForListDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserForListDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserForListDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *UserForListDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserForListDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCreatedAt

`func (o *UserForListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserForListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserForListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetIs2FAEnabled

`func (o *UserForListDto) GetIs2FAEnabled() bool`

GetIs2FAEnabled returns the Is2FAEnabled field if non-nil, zero value otherwise.

### GetIs2FAEnabledOk

`func (o *UserForListDto) GetIs2FAEnabledOk() (*bool, bool)`

GetIs2FAEnabledOk returns a tuple with the Is2FAEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2FAEnabled

`func (o *UserForListDto) SetIs2FAEnabled(v bool)`

SetIs2FAEnabled sets Is2FAEnabled field to given value.


### GetDomain

`func (o *UserForListDto) GetDomain() DomainDetailsForUserDto`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserForListDto) GetDomainOk() (*DomainDetailsForUserDto, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserForListDto) SetDomain(v DomainDetailsForUserDto)`

SetDomain sets Domain field to given value.


### GetRole

`func (o *UserForListDto) GetRole() EGlobalRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserForListDto) GetRoleOk() (*EGlobalRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserForListDto) SetRole(v EGlobalRole)`

SetRole sets Role field to given value.


### GetOrganizations

`func (o *UserForListDto) GetOrganizations() map[string]OrganizationContextDto`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserForListDto) GetOrganizationsOk() (*map[string]OrganizationContextDto, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserForListDto) SetOrganizations(v map[string]OrganizationContextDto)`

SetOrganizations sets Organizations field to given value.


### GetIsEmailConfirmed

`func (o *UserForListDto) GetIsEmailConfirmed() bool`

GetIsEmailConfirmed returns the IsEmailConfirmed field if non-nil, zero value otherwise.

### GetIsEmailConfirmedOk

`func (o *UserForListDto) GetIsEmailConfirmedOk() (*bool, bool)`

GetIsEmailConfirmedOk returns a tuple with the IsEmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailConfirmed

`func (o *UserForListDto) SetIsEmailConfirmed(v bool)`

SetIsEmailConfirmed sets IsEmailConfirmed field to given value.


### GetIsEmailNotificationEnabled

`func (o *UserForListDto) GetIsEmailNotificationEnabled() bool`

GetIsEmailNotificationEnabled returns the IsEmailNotificationEnabled field if non-nil, zero value otherwise.

### GetIsEmailNotificationEnabledOk

`func (o *UserForListDto) GetIsEmailNotificationEnabledOk() (*bool, bool)`

GetIsEmailNotificationEnabledOk returns a tuple with the IsEmailNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailNotificationEnabled

`func (o *UserForListDto) SetIsEmailNotificationEnabled(v bool)`

SetIsEmailNotificationEnabled sets IsEmailNotificationEnabled field to given value.


### GetIsForcedToResetPassword

`func (o *UserForListDto) GetIsForcedToResetPassword() bool`

GetIsForcedToResetPassword returns the IsForcedToResetPassword field if non-nil, zero value otherwise.

### GetIsForcedToResetPasswordOk

`func (o *UserForListDto) GetIsForcedToResetPasswordOk() (*bool, bool)`

GetIsForcedToResetPasswordOk returns a tuple with the IsForcedToResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcedToResetPassword

`func (o *UserForListDto) SetIsForcedToResetPassword(v bool)`

SetIsForcedToResetPassword sets IsForcedToResetPassword field to given value.


### GetIsLocked

`func (o *UserForListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserForListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserForListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetOwner

`func (o *UserForListDto) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserForListDto) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserForListDto) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetIsForcedToEnableTwoFactorAuthentication

`func (o *UserForListDto) GetIsForcedToEnableTwoFactorAuthentication() bool`

GetIsForcedToEnableTwoFactorAuthentication returns the IsForcedToEnableTwoFactorAuthentication field if non-nil, zero value otherwise.

### GetIsForcedToEnableTwoFactorAuthenticationOk

`func (o *UserForListDto) GetIsForcedToEnableTwoFactorAuthenticationOk() (*bool, bool)`

GetIsForcedToEnableTwoFactorAuthenticationOk returns a tuple with the IsForcedToEnableTwoFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcedToEnableTwoFactorAuthentication

`func (o *UserForListDto) SetIsForcedToEnableTwoFactorAuthentication(v bool)`

SetIsForcedToEnableTwoFactorAuthentication sets IsForcedToEnableTwoFactorAuthentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


