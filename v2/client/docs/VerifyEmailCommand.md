# VerifyEmailCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **NullableString** |  | 
**Mode** | [**EmailMode**](EmailMode.md) |  | 

## Methods

### NewVerifyEmailCommand

`func NewVerifyEmailCommand(token NullableString, mode EmailMode, ) *VerifyEmailCommand`

NewVerifyEmailCommand instantiates a new VerifyEmailCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEmailCommandWithDefaults

`func NewVerifyEmailCommandWithDefaults() *VerifyEmailCommand`

NewVerifyEmailCommandWithDefaults instantiates a new VerifyEmailCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *VerifyEmailCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyEmailCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyEmailCommand) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *VerifyEmailCommand) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *VerifyEmailCommand) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetMode

`func (o *VerifyEmailCommand) GetMode() EmailMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VerifyEmailCommand) GetModeOk() (*EmailMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VerifyEmailCommand) SetMode(v EmailMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


