# CloudCredentialsForOrganizationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**FullName** | **NullableString** |  | 
**CloudType** | **map[string]interface{}** |  | 
**IsDefault** | **bool** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCredentialsForOrganizationEntity

`func NewCloudCredentialsForOrganizationEntity(id int32, projects []CommonDropdownDto, fullName NullableString, cloudType map[string]interface{}, isDefault bool, ) *CloudCredentialsForOrganizationEntity`

NewCloudCredentialsForOrganizationEntity instantiates a new CloudCredentialsForOrganizationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsForOrganizationEntityWithDefaults

`func NewCloudCredentialsForOrganizationEntityWithDefaults() *CloudCredentialsForOrganizationEntity`

NewCloudCredentialsForOrganizationEntityWithDefaults instantiates a new CloudCredentialsForOrganizationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudCredentialsForOrganizationEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialsForOrganizationEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialsForOrganizationEntity) SetId(v int32)`

SetId sets Id field to given value.


### GetProjects

`func (o *CloudCredentialsForOrganizationEntity) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CloudCredentialsForOrganizationEntity) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CloudCredentialsForOrganizationEntity) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### SetProjectsNil

`func (o *CloudCredentialsForOrganizationEntity) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *CloudCredentialsForOrganizationEntity) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetFullName

`func (o *CloudCredentialsForOrganizationEntity) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CloudCredentialsForOrganizationEntity) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CloudCredentialsForOrganizationEntity) SetFullName(v string)`

SetFullName sets FullName field to given value.


### SetFullNameNil

`func (o *CloudCredentialsForOrganizationEntity) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *CloudCredentialsForOrganizationEntity) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetCloudType

`func (o *CloudCredentialsForOrganizationEntity) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CloudCredentialsForOrganizationEntity) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CloudCredentialsForOrganizationEntity) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetIsDefault

`func (o *CloudCredentialsForOrganizationEntity) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudCredentialsForOrganizationEntity) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudCredentialsForOrganizationEntity) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOrganizationId

`func (o *CloudCredentialsForOrganizationEntity) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CloudCredentialsForOrganizationEntity) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CloudCredentialsForOrganizationEntity) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CloudCredentialsForOrganizationEntity) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


