# CreateUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DomainId** | Pointer to **NullableInt32** |  | [optional] 
**IsDomainAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUserCommand

`func NewCreateUserCommand() *CreateUserCommand`

NewCreateUserCommand instantiates a new CreateUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserCommandWithDefaults

`func NewCreateUserCommandWithDefaults() *CreateUserCommand`

NewCreateUserCommandWithDefaults instantiates a new CreateUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateUserCommand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateUserCommand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateUserCommand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateUserCommand) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateUserCommand) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateUserCommand) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *CreateUserCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateUserCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateUserCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDomainId

`func (o *CreateUserCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateUserCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateUserCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *CreateUserCommand) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *CreateUserCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *CreateUserCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetIsDomainAdmin

`func (o *CreateUserCommand) GetIsDomainAdmin() bool`

GetIsDomainAdmin returns the IsDomainAdmin field if non-nil, zero value otherwise.

### GetIsDomainAdminOk

`func (o *CreateUserCommand) GetIsDomainAdminOk() (*bool, bool)`

GetIsDomainAdminOk returns a tuple with the IsDomainAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainAdmin

`func (o *CreateUserCommand) SetIsDomainAdmin(v bool)`

SetIsDomainAdmin sets IsDomainAdmin field to given value.

### HasIsDomainAdmin

`func (o *CreateUserCommand) HasIsDomainAdmin() bool`

HasIsDomainAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


