# UpdateRepoPasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **NullableString** |  | 
**Password** | **NullableString** |  | 
**RepositoryId** | **NullableString** |  | 
**OrganizationId** | **int32** |  | 

## Methods

### NewUpdateRepoPasswordCommand

`func NewUpdateRepoPasswordCommand(username NullableString, password NullableString, repositoryId NullableString, organizationId int32, ) *UpdateRepoPasswordCommand`

NewUpdateRepoPasswordCommand instantiates a new UpdateRepoPasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRepoPasswordCommandWithDefaults

`func NewUpdateRepoPasswordCommandWithDefaults() *UpdateRepoPasswordCommand`

NewUpdateRepoPasswordCommandWithDefaults instantiates a new UpdateRepoPasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateRepoPasswordCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateRepoPasswordCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateRepoPasswordCommand) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *UpdateRepoPasswordCommand) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateRepoPasswordCommand) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateRepoPasswordCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRepoPasswordCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRepoPasswordCommand) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *UpdateRepoPasswordCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateRepoPasswordCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRepositoryId

`func (o *UpdateRepoPasswordCommand) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *UpdateRepoPasswordCommand) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *UpdateRepoPasswordCommand) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.


### SetRepositoryIdNil

`func (o *UpdateRepoPasswordCommand) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *UpdateRepoPasswordCommand) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetOrganizationId

`func (o *UpdateRepoPasswordCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateRepoPasswordCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateRepoPasswordCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


