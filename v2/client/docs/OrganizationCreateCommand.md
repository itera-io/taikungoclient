# OrganizationCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DomainId** | Pointer to **NullableInt32** |  | [optional] 
**AdminCloudCredentialId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewOrganizationCreateCommand

`func NewOrganizationCreateCommand() *OrganizationCreateCommand`

NewOrganizationCreateCommand instantiates a new OrganizationCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateCommandWithDefaults

`func NewOrganizationCreateCommandWithDefaults() *OrganizationCreateCommand`

NewOrganizationCreateCommandWithDefaults instantiates a new OrganizationCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreateCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationCreateCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationCreateCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationCreateCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *OrganizationCreateCommand) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationCreateCommand) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationCreateCommand) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OrganizationCreateCommand) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *OrganizationCreateCommand) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationCreateCommand) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetEmail

`func (o *OrganizationCreateCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationCreateCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationCreateCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationCreateCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationCreateCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationCreateCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDomainId

`func (o *OrganizationCreateCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *OrganizationCreateCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *OrganizationCreateCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *OrganizationCreateCommand) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *OrganizationCreateCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *OrganizationCreateCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialId() int32`

GetAdminCloudCredentialId returns the AdminCloudCredentialId field if non-nil, zero value otherwise.

### GetAdminCloudCredentialIdOk

`func (o *OrganizationCreateCommand) GetAdminCloudCredentialIdOk() (*int32, bool)`

GetAdminCloudCredentialIdOk returns a tuple with the AdminCloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCloudCredentialId

`func (o *OrganizationCreateCommand) SetAdminCloudCredentialId(v int32)`

SetAdminCloudCredentialId sets AdminCloudCredentialId field to given value.

### HasAdminCloudCredentialId

`func (o *OrganizationCreateCommand) HasAdminCloudCredentialId() bool`

HasAdminCloudCredentialId returns a boolean if a field has been set.

### SetAdminCloudCredentialIdNil

`func (o *OrganizationCreateCommand) SetAdminCloudCredentialIdNil(b bool)`

 SetAdminCloudCredentialIdNil sets the value for AdminCloudCredentialId to be an explicit nil

### UnsetAdminCloudCredentialId
`func (o *OrganizationCreateCommand) UnsetAdminCloudCredentialId()`

UnsetAdminCloudCredentialId ensures that no value is present for AdminCloudCredentialId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


