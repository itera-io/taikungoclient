# OrganizationDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to **int32** |  | [optional] 
**Servers** | Pointer to **int32** |  | [optional] 
**CloudCredentials** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**BoundRules** | Pointer to [**[]PrometheusEntity**](PrometheusEntity.md) |  | [optional] 

## Methods

### NewOrganizationDetailsDto

`func NewOrganizationDetailsDto() *OrganizationDetailsDto`

NewOrganizationDetailsDto instantiates a new OrganizationDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsDtoWithDefaults

`func NewOrganizationDetailsDtoWithDefaults() *OrganizationDetailsDto`

NewOrganizationDetailsDtoWithDefaults instantiates a new OrganizationDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *OrganizationDetailsDto) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationDetailsDto) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationDetailsDto) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OrganizationDetailsDto) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *OrganizationDetailsDto) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationDetailsDto) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetEmail

`func (o *OrganizationDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationDetailsDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationDetailsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationDetailsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetProjects

`func (o *OrganizationDetailsDto) GetProjects() int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationDetailsDto) GetProjectsOk() (*int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationDetailsDto) SetProjects(v int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationDetailsDto) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetServers

`func (o *OrganizationDetailsDto) GetServers() int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OrganizationDetailsDto) GetServersOk() (*int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OrganizationDetailsDto) SetServers(v int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OrganizationDetailsDto) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetCloudCredentials

`func (o *OrganizationDetailsDto) GetCloudCredentials() int32`

GetCloudCredentials returns the CloudCredentials field if non-nil, zero value otherwise.

### GetCloudCredentialsOk

`func (o *OrganizationDetailsDto) GetCloudCredentialsOk() (*int32, bool)`

GetCloudCredentialsOk returns a tuple with the CloudCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentials

`func (o *OrganizationDetailsDto) SetCloudCredentials(v int32)`

SetCloudCredentials sets CloudCredentials field to given value.

### HasCloudCredentials

`func (o *OrganizationDetailsDto) HasCloudCredentials() bool`

HasCloudCredentials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationDetailsDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *OrganizationDetailsDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OrganizationDetailsDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetBoundRules

`func (o *OrganizationDetailsDto) GetBoundRules() []PrometheusEntity`

GetBoundRules returns the BoundRules field if non-nil, zero value otherwise.

### GetBoundRulesOk

`func (o *OrganizationDetailsDto) GetBoundRulesOk() (*[]PrometheusEntity, bool)`

GetBoundRulesOk returns a tuple with the BoundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRules

`func (o *OrganizationDetailsDto) SetBoundRules(v []PrometheusEntity)`

SetBoundRules sets BoundRules field to given value.

### HasBoundRules

`func (o *OrganizationDetailsDto) HasBoundRules() bool`

HasBoundRules returns a boolean if a field has been set.

### SetBoundRulesNil

`func (o *OrganizationDetailsDto) SetBoundRulesNil(b bool)`

 SetBoundRulesNil sets the value for BoundRules to be an explicit nil

### UnsetBoundRules
`func (o *OrganizationDetailsDto) UnsetBoundRules()`

UnsetBoundRules ensures that no value is present for BoundRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


