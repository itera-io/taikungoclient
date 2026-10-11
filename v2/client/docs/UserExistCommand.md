# UserExistCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** |  | 
**UserName** | **NullableString** |  | 
**DomainId** | **NullableInt32** |  | 

## Methods

### NewUserExistCommand

`func NewUserExistCommand(email NullableString, userName NullableString, domainId NullableInt32, ) *UserExistCommand`

NewUserExistCommand instantiates a new UserExistCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExistCommandWithDefaults

`func NewUserExistCommandWithDefaults() *UserExistCommand`

NewUserExistCommandWithDefaults instantiates a new UserExistCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserExistCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserExistCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserExistCommand) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *UserExistCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserExistCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUserName

`func (o *UserExistCommand) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserExistCommand) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserExistCommand) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *UserExistCommand) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserExistCommand) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetDomainId

`func (o *UserExistCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UserExistCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UserExistCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### SetDomainIdNil

`func (o *UserExistCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *UserExistCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


