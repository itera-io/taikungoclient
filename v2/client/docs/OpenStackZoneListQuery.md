# OpenStackZoneListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | **NullableString** |  | 
**OpenStackPassword** | **NullableString** |  | 
**OpenStackUrl** | **NullableString** |  | 
**OpenStackDomain** | **NullableString** |  | 
**OpenStackRegion** | **NullableString** |  | 
**ApplicationCredEnabled** | **bool** |  | 
**IsAdmin** | **bool** |  | 
**OpenstackProject** | **NullableString** |  | 

## Methods

### NewOpenStackZoneListQuery

`func NewOpenStackZoneListQuery(openStackUser NullableString, openStackPassword NullableString, openStackUrl NullableString, openStackDomain NullableString, openStackRegion NullableString, applicationCredEnabled bool, isAdmin bool, openstackProject NullableString, ) *OpenStackZoneListQuery`

NewOpenStackZoneListQuery instantiates a new OpenStackZoneListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackZoneListQueryWithDefaults

`func NewOpenStackZoneListQueryWithDefaults() *OpenStackZoneListQuery`

NewOpenStackZoneListQueryWithDefaults instantiates a new OpenStackZoneListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *OpenStackZoneListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenStackZoneListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenStackZoneListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### SetOpenStackUserNil

`func (o *OpenStackZoneListQuery) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *OpenStackZoneListQuery) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *OpenStackZoneListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenStackZoneListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenStackZoneListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### SetOpenStackPasswordNil

`func (o *OpenStackZoneListQuery) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *OpenStackZoneListQuery) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *OpenStackZoneListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenStackZoneListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenStackZoneListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### SetOpenStackUrlNil

`func (o *OpenStackZoneListQuery) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *OpenStackZoneListQuery) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackDomain

`func (o *OpenStackZoneListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenStackZoneListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenStackZoneListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.


### SetOpenStackDomainNil

`func (o *OpenStackZoneListQuery) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *OpenStackZoneListQuery) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetOpenStackRegion

`func (o *OpenStackZoneListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenStackZoneListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenStackZoneListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.


### SetOpenStackRegionNil

`func (o *OpenStackZoneListQuery) SetOpenStackRegionNil(b bool)`

 SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil

### UnsetOpenStackRegion
`func (o *OpenStackZoneListQuery) UnsetOpenStackRegion()`

UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
### GetApplicationCredEnabled

`func (o *OpenStackZoneListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenStackZoneListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenStackZoneListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.


### GetIsAdmin

`func (o *OpenStackZoneListQuery) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenStackZoneListQuery) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenStackZoneListQuery) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetOpenstackProject

`func (o *OpenStackZoneListQuery) GetOpenstackProject() string`

GetOpenstackProject returns the OpenstackProject field if non-nil, zero value otherwise.

### GetOpenstackProjectOk

`func (o *OpenStackZoneListQuery) GetOpenstackProjectOk() (*string, bool)`

GetOpenstackProjectOk returns a tuple with the OpenstackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackProject

`func (o *OpenStackZoneListQuery) SetOpenstackProject(v string)`

SetOpenstackProject sets OpenstackProject field to given value.


### SetOpenstackProjectNil

`func (o *OpenStackZoneListQuery) SetOpenstackProjectNil(b bool)`

 SetOpenstackProjectNil sets the value for OpenstackProject to be an explicit nil

### UnsetOpenstackProject
`func (o *OpenStackZoneListQuery) UnsetOpenstackProject()`

UnsetOpenstackProject ensures that no value is present for OpenstackProject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


