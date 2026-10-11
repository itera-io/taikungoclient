# UpdateDomainCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Email** | **NullableString** |  | 

## Methods

### NewUpdateDomainCommand

`func NewUpdateDomainCommand(id int32, name NullableString, email NullableString, ) *UpdateDomainCommand`

NewUpdateDomainCommand instantiates a new UpdateDomainCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainCommandWithDefaults

`func NewUpdateDomainCommandWithDefaults() *UpdateDomainCommand`

NewUpdateDomainCommandWithDefaults instantiates a new UpdateDomainCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateDomainCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDomainCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDomainCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateDomainCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDomainCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDomainCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateDomainCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDomainCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *UpdateDomainCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateDomainCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateDomainCommand) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *UpdateDomainCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateDomainCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


