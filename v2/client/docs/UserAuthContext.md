# UserAuthContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DomainId** | Pointer to **int32** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**MaintenanceModeEnabled** | Pointer to **bool** |  | [optional] 
**TwoFactorEnabled** | Pointer to **bool** |  | [optional] 
**GlobalRole** | Pointer to [**GlobalRole**](GlobalRole.md) |  | [optional] 
**OrgRoles** | Pointer to [**map[string]RoleClaim**](RoleClaim.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserAuthContext

`func NewUserAuthContext() *UserAuthContext`

NewUserAuthContext instantiates a new UserAuthContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthContextWithDefaults

`func NewUserAuthContextWithDefaults() *UserAuthContext`

NewUserAuthContextWithDefaults instantiates a new UserAuthContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserAuthContext) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserAuthContext) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserAuthContext) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserAuthContext) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserAuthContext) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserAuthContext) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDisplayName

`func (o *UserAuthContext) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserAuthContext) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserAuthContext) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserAuthContext) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserAuthContext) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserAuthContext) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *UserAuthContext) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAuthContext) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAuthContext) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAuthContext) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserAuthContext) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserAuthContext) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDomainId

`func (o *UserAuthContext) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UserAuthContext) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UserAuthContext) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *UserAuthContext) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *UserAuthContext) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UserAuthContext) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UserAuthContext) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UserAuthContext) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *UserAuthContext) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *UserAuthContext) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetLogoUrl

`func (o *UserAuthContext) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UserAuthContext) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UserAuthContext) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UserAuthContext) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *UserAuthContext) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *UserAuthContext) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetDomain

`func (o *UserAuthContext) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserAuthContext) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserAuthContext) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserAuthContext) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserAuthContext) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserAuthContext) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetMaintenanceModeEnabled

`func (o *UserAuthContext) GetMaintenanceModeEnabled() bool`

GetMaintenanceModeEnabled returns the MaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetMaintenanceModeEnabledOk

`func (o *UserAuthContext) GetMaintenanceModeEnabledOk() (*bool, bool)`

GetMaintenanceModeEnabledOk returns a tuple with the MaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeEnabled

`func (o *UserAuthContext) SetMaintenanceModeEnabled(v bool)`

SetMaintenanceModeEnabled sets MaintenanceModeEnabled field to given value.

### HasMaintenanceModeEnabled

`func (o *UserAuthContext) HasMaintenanceModeEnabled() bool`

HasMaintenanceModeEnabled returns a boolean if a field has been set.

### GetTwoFactorEnabled

`func (o *UserAuthContext) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *UserAuthContext) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *UserAuthContext) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.

### HasTwoFactorEnabled

`func (o *UserAuthContext) HasTwoFactorEnabled() bool`

HasTwoFactorEnabled returns a boolean if a field has been set.

### GetGlobalRole

`func (o *UserAuthContext) GetGlobalRole() GlobalRole`

GetGlobalRole returns the GlobalRole field if non-nil, zero value otherwise.

### GetGlobalRoleOk

`func (o *UserAuthContext) GetGlobalRoleOk() (*GlobalRole, bool)`

GetGlobalRoleOk returns a tuple with the GlobalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRole

`func (o *UserAuthContext) SetGlobalRole(v GlobalRole)`

SetGlobalRole sets GlobalRole field to given value.

### HasGlobalRole

`func (o *UserAuthContext) HasGlobalRole() bool`

HasGlobalRole returns a boolean if a field has been set.

### GetOrgRoles

`func (o *UserAuthContext) GetOrgRoles() map[string]RoleClaim`

GetOrgRoles returns the OrgRoles field if non-nil, zero value otherwise.

### GetOrgRolesOk

`func (o *UserAuthContext) GetOrgRolesOk() (*map[string]RoleClaim, bool)`

GetOrgRolesOk returns a tuple with the OrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoles

`func (o *UserAuthContext) SetOrgRoles(v map[string]RoleClaim)`

SetOrgRoles sets OrgRoles field to given value.

### HasOrgRoles

`func (o *UserAuthContext) HasOrgRoles() bool`

HasOrgRoles returns a boolean if a field has been set.

### SetOrgRolesNil

`func (o *UserAuthContext) SetOrgRolesNil(b bool)`

 SetOrgRolesNil sets the value for OrgRoles to be an explicit nil

### UnsetOrgRoles
`func (o *UserAuthContext) UnsetOrgRoles()`

UnsetOrgRoles ensures that no value is present for OrgRoles, not even an explicit nil
### GetCreatedAt

`func (o *UserAuthContext) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAuthContext) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAuthContext) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserAuthContext) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *UserAuthContext) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UserAuthContext) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUserName

`func (o *UserAuthContext) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserAuthContext) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserAuthContext) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserAuthContext) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserAuthContext) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserAuthContext) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


