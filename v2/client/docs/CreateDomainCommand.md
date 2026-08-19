# CreateDomainCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**CreateOrganization** | **bool** |  | 

## Methods

### NewCreateDomainCommand

`func NewCreateDomainCommand(name NullableString, email NullableString, createOrganization bool, ) *CreateDomainCommand`

NewCreateDomainCommand instantiates a new CreateDomainCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainCommandWithDefaults

`func NewCreateDomainCommandWithDefaults() *CreateDomainCommand`

NewCreateDomainCommandWithDefaults instantiates a new CreateDomainCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDomainCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDomainCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDomainCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateDomainCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateDomainCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *CreateDomainCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateDomainCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateDomainCommand) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *CreateDomainCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateDomainCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCreateOrganization

`func (o *CreateDomainCommand) GetCreateOrganization() bool`

GetCreateOrganization returns the CreateOrganization field if non-nil, zero value otherwise.

### GetCreateOrganizationOk

`func (o *CreateDomainCommand) GetCreateOrganizationOk() (*bool, bool)`

GetCreateOrganizationOk returns a tuple with the CreateOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOrganization

`func (o *CreateDomainCommand) SetCreateOrganization(v bool)`

SetCreateOrganization sets CreateOrganization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


