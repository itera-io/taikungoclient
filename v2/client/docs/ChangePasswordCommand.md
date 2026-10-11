# ChangePasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **NullableString** |  | 
**NewPassword** | **NullableString** |  | 

## Methods

### NewChangePasswordCommand

`func NewChangePasswordCommand(password NullableString, newPassword NullableString, ) *ChangePasswordCommand`

NewChangePasswordCommand instantiates a new ChangePasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordCommandWithDefaults

`func NewChangePasswordCommandWithDefaults() *ChangePasswordCommand`

NewChangePasswordCommandWithDefaults instantiates a new ChangePasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ChangePasswordCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordCommand) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *ChangePasswordCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ChangePasswordCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetNewPassword

`func (o *ChangePasswordCommand) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ChangePasswordCommand) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ChangePasswordCommand) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### SetNewPasswordNil

`func (o *ChangePasswordCommand) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *ChangePasswordCommand) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


