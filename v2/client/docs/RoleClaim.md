# RoleClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **NullableString** |  | [optional] 
**Claim** | Pointer to **[]string** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoleClaim

`func NewRoleClaim() *RoleClaim`

NewRoleClaim instantiates a new RoleClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleClaimWithDefaults

`func NewRoleClaimWithDefaults() *RoleClaim`

NewRoleClaimWithDefaults instantiates a new RoleClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RoleClaim) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleClaim) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleClaim) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleClaim) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *RoleClaim) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *RoleClaim) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetClaim

`func (o *RoleClaim) GetClaim() []string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *RoleClaim) GetClaimOk() (*[]string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *RoleClaim) SetClaim(v []string)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *RoleClaim) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### SetClaimNil

`func (o *RoleClaim) SetClaimNil(b bool)`

 SetClaimNil sets the value for Claim to be an explicit nil

### UnsetClaim
`func (o *RoleClaim) UnsetClaim()`

UnsetClaim ensures that no value is present for Claim, not even an explicit nil
### GetOrganizationName

`func (o *RoleClaim) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *RoleClaim) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *RoleClaim) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *RoleClaim) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *RoleClaim) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *RoleClaim) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


