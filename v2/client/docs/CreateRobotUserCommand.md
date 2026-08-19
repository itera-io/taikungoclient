# CreateRobotUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **NullableInt32** |  | 
**OrganizationId** | **NullableInt32** |  | 
**Name** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**ExpiresAt** | **NullableTime** |  | 
**Scopes** | **[]string** |  | 
**Ips** | **[]string** |  | 

## Methods

### NewCreateRobotUserCommand

`func NewCreateRobotUserCommand(domainId NullableInt32, organizationId NullableInt32, name NullableString, description NullableString, expiresAt NullableTime, scopes []string, ips []string, ) *CreateRobotUserCommand`

NewCreateRobotUserCommand instantiates a new CreateRobotUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRobotUserCommandWithDefaults

`func NewCreateRobotUserCommandWithDefaults() *CreateRobotUserCommand`

NewCreateRobotUserCommandWithDefaults instantiates a new CreateRobotUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *CreateRobotUserCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateRobotUserCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateRobotUserCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### SetDomainIdNil

`func (o *CreateRobotUserCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *CreateRobotUserCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetOrganizationId

`func (o *CreateRobotUserCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateRobotUserCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateRobotUserCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CreateRobotUserCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateRobotUserCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetName

`func (o *CreateRobotUserCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRobotUserCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRobotUserCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateRobotUserCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateRobotUserCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateRobotUserCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRobotUserCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRobotUserCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CreateRobotUserCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateRobotUserCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpiresAt

`func (o *CreateRobotUserCommand) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateRobotUserCommand) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateRobotUserCommand) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CreateRobotUserCommand) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateRobotUserCommand) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetScopes

`func (o *CreateRobotUserCommand) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateRobotUserCommand) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateRobotUserCommand) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *CreateRobotUserCommand) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CreateRobotUserCommand) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetIps

`func (o *CreateRobotUserCommand) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *CreateRobotUserCommand) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *CreateRobotUserCommand) SetIps(v []string)`

SetIps sets Ips field to given value.


### SetIpsNil

`func (o *CreateRobotUserCommand) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *CreateRobotUserCommand) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


