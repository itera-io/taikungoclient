# ForgotPasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**DomainName** | **NullableString** |  | 

## Methods

### NewForgotPasswordCommand

`func NewForgotPasswordCommand(email NullableString, domainName NullableString, ) *ForgotPasswordCommand`

NewForgotPasswordCommand instantiates a new ForgotPasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordCommandWithDefaults

`func NewForgotPasswordCommandWithDefaults() *ForgotPasswordCommand`

NewForgotPasswordCommandWithDefaults instantiates a new ForgotPasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ForgotPasswordCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ForgotPasswordCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ForgotPasswordCommand) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ForgotPasswordCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ForgotPasswordCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDomainName

`func (o *ForgotPasswordCommand) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ForgotPasswordCommand) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ForgotPasswordCommand) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *ForgotPasswordCommand) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ForgotPasswordCommand) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


