# TwoFactorAuthRecoveryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryCode** | **NullableString** |  | 

## Methods

### NewTwoFactorAuthRecoveryCommand

`func NewTwoFactorAuthRecoveryCommand(recoveryCode NullableString, ) *TwoFactorAuthRecoveryCommand`

NewTwoFactorAuthRecoveryCommand instantiates a new TwoFactorAuthRecoveryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorAuthRecoveryCommandWithDefaults

`func NewTwoFactorAuthRecoveryCommandWithDefaults() *TwoFactorAuthRecoveryCommand`

NewTwoFactorAuthRecoveryCommandWithDefaults instantiates a new TwoFactorAuthRecoveryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryCode

`func (o *TwoFactorAuthRecoveryCommand) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *TwoFactorAuthRecoveryCommand) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *TwoFactorAuthRecoveryCommand) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.


### SetRecoveryCodeNil

`func (o *TwoFactorAuthRecoveryCommand) SetRecoveryCodeNil(b bool)`

 SetRecoveryCodeNil sets the value for RecoveryCode to be an explicit nil

### UnsetRecoveryCode
`func (o *TwoFactorAuthRecoveryCommand) UnsetRecoveryCode()`

UnsetRecoveryCode ensures that no value is present for RecoveryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


