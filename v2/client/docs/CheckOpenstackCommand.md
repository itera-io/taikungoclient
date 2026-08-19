# CheckOpenstackCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | Pointer to **NullableString** |  | [optional] 
**OpenStackPassword** | Pointer to **NullableString** |  | [optional] 
**OpenStackUrl** | Pointer to **NullableString** |  | [optional] 
**OpenStackDomain** | Pointer to **NullableString** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**ApplicationCredEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCheckOpenstackCommand

`func NewCheckOpenstackCommand() *CheckOpenstackCommand`

NewCheckOpenstackCommand instantiates a new CheckOpenstackCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckOpenstackCommandWithDefaults

`func NewCheckOpenstackCommandWithDefaults() *CheckOpenstackCommand`

NewCheckOpenstackCommandWithDefaults instantiates a new CheckOpenstackCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *CheckOpenstackCommand) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *CheckOpenstackCommand) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *CheckOpenstackCommand) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.

### HasOpenStackUser

`func (o *CheckOpenstackCommand) HasOpenStackUser() bool`

HasOpenStackUser returns a boolean if a field has been set.

### SetOpenStackUserNil

`func (o *CheckOpenstackCommand) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *CheckOpenstackCommand) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *CheckOpenstackCommand) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *CheckOpenstackCommand) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *CheckOpenstackCommand) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.

### HasOpenStackPassword

`func (o *CheckOpenstackCommand) HasOpenStackPassword() bool`

HasOpenStackPassword returns a boolean if a field has been set.

### SetOpenStackPasswordNil

`func (o *CheckOpenstackCommand) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *CheckOpenstackCommand) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *CheckOpenstackCommand) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *CheckOpenstackCommand) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *CheckOpenstackCommand) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.

### HasOpenStackUrl

`func (o *CheckOpenstackCommand) HasOpenStackUrl() bool`

HasOpenStackUrl returns a boolean if a field has been set.

### SetOpenStackUrlNil

`func (o *CheckOpenstackCommand) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *CheckOpenstackCommand) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackDomain

`func (o *CheckOpenstackCommand) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *CheckOpenstackCommand) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *CheckOpenstackCommand) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.

### HasOpenStackDomain

`func (o *CheckOpenstackCommand) HasOpenStackDomain() bool`

HasOpenStackDomain returns a boolean if a field has been set.

### SetOpenStackDomainNil

`func (o *CheckOpenstackCommand) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *CheckOpenstackCommand) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetIsAdmin

`func (o *CheckOpenstackCommand) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CheckOpenstackCommand) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CheckOpenstackCommand) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CheckOpenstackCommand) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetApplicationCredEnabled

`func (o *CheckOpenstackCommand) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *CheckOpenstackCommand) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *CheckOpenstackCommand) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.

### HasApplicationCredEnabled

`func (o *CheckOpenstackCommand) HasApplicationCredEnabled() bool`

HasApplicationCredEnabled returns a boolean if a field has been set.

### SetApplicationCredEnabledNil

`func (o *CheckOpenstackCommand) SetApplicationCredEnabledNil(b bool)`

 SetApplicationCredEnabledNil sets the value for ApplicationCredEnabled to be an explicit nil

### UnsetApplicationCredEnabled
`func (o *CheckOpenstackCommand) UnsetApplicationCredEnabled()`

UnsetApplicationCredEnabled ensures that no value is present for ApplicationCredEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


