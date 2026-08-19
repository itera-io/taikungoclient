# DisableTwoFactorAuthCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** |  | 
**IsRecoveryCode** | **bool** |  | 

## Methods

### NewDisableTwoFactorAuthCommand

`func NewDisableTwoFactorAuthCommand(code NullableString, isRecoveryCode bool, ) *DisableTwoFactorAuthCommand`

NewDisableTwoFactorAuthCommand instantiates a new DisableTwoFactorAuthCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableTwoFactorAuthCommandWithDefaults

`func NewDisableTwoFactorAuthCommandWithDefaults() *DisableTwoFactorAuthCommand`

NewDisableTwoFactorAuthCommandWithDefaults instantiates a new DisableTwoFactorAuthCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DisableTwoFactorAuthCommand) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DisableTwoFactorAuthCommand) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DisableTwoFactorAuthCommand) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *DisableTwoFactorAuthCommand) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *DisableTwoFactorAuthCommand) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIsRecoveryCode

`func (o *DisableTwoFactorAuthCommand) GetIsRecoveryCode() bool`

GetIsRecoveryCode returns the IsRecoveryCode field if non-nil, zero value otherwise.

### GetIsRecoveryCodeOk

`func (o *DisableTwoFactorAuthCommand) GetIsRecoveryCodeOk() (*bool, bool)`

GetIsRecoveryCodeOk returns a tuple with the IsRecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecoveryCode

`func (o *DisableTwoFactorAuthCommand) SetIsRecoveryCode(v bool)`

SetIsRecoveryCode sets IsRecoveryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


