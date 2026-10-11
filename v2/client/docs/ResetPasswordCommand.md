# ResetPasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**NewPassword** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResetPasswordCommand

`func NewResetPasswordCommand() *ResetPasswordCommand`

NewResetPasswordCommand instantiates a new ResetPasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordCommandWithDefaults

`func NewResetPasswordCommandWithDefaults() *ResetPasswordCommand`

NewResetPasswordCommandWithDefaults instantiates a new ResetPasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResetPasswordCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetPasswordCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetPasswordCommand) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ResetPasswordCommand) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ResetPasswordCommand) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ResetPasswordCommand) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetEmail

`func (o *ResetPasswordCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResetPasswordCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ResetPasswordCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ResetPasswordCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDomainName

`func (o *ResetPasswordCommand) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ResetPasswordCommand) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ResetPasswordCommand) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ResetPasswordCommand) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ResetPasswordCommand) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ResetPasswordCommand) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetNewPassword

`func (o *ResetPasswordCommand) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordCommand) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordCommand) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *ResetPasswordCommand) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *ResetPasswordCommand) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *ResetPasswordCommand) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


