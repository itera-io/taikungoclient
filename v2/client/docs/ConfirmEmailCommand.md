# ConfirmEmailCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewEmail** | **NullableString** |  | 
**Mode** | [**EmailMode**](EmailMode.md) |  | 

## Methods

### NewConfirmEmailCommand

`func NewConfirmEmailCommand(newEmail NullableString, mode EmailMode, ) *ConfirmEmailCommand`

NewConfirmEmailCommand instantiates a new ConfirmEmailCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmEmailCommandWithDefaults

`func NewConfirmEmailCommandWithDefaults() *ConfirmEmailCommand`

NewConfirmEmailCommandWithDefaults instantiates a new ConfirmEmailCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewEmail

`func (o *ConfirmEmailCommand) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *ConfirmEmailCommand) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *ConfirmEmailCommand) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.


### SetNewEmailNil

`func (o *ConfirmEmailCommand) SetNewEmailNil(b bool)`

 SetNewEmailNil sets the value for NewEmail to be an explicit nil

### UnsetNewEmail
`func (o *ConfirmEmailCommand) UnsetNewEmail()`

UnsetNewEmail ensures that no value is present for NewEmail, not even an explicit nil
### GetMode

`func (o *ConfirmEmailCommand) GetMode() EmailMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConfirmEmailCommand) GetModeOk() (*EmailMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConfirmEmailCommand) SetMode(v EmailMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


