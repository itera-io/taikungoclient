# OpenstackVolumeTypeListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **NullableInt32** |  | 
**OpenStackUser** | **NullableString** |  | 
**OpenStackPassword** | **NullableString** |  | 
**OpenStackUrl** | **NullableString** |  | 
**OpenStackDomain** | **NullableString** |  | 
**OpenStackRegion** | **NullableString** |  | 
**ApplicationCredEnabled** | **bool** |  | 
**IsAdmin** | **bool** |  | 
**OpenstackProject** | **NullableString** |  | 

## Methods

### NewOpenstackVolumeTypeListQuery

`func NewOpenstackVolumeTypeListQuery(projectId NullableInt32, openStackUser NullableString, openStackPassword NullableString, openStackUrl NullableString, openStackDomain NullableString, openStackRegion NullableString, applicationCredEnabled bool, isAdmin bool, openstackProject NullableString, ) *OpenstackVolumeTypeListQuery`

NewOpenstackVolumeTypeListQuery instantiates a new OpenstackVolumeTypeListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackVolumeTypeListQueryWithDefaults

`func NewOpenstackVolumeTypeListQueryWithDefaults() *OpenstackVolumeTypeListQuery`

NewOpenstackVolumeTypeListQueryWithDefaults instantiates a new OpenstackVolumeTypeListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OpenstackVolumeTypeListQuery) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpenstackVolumeTypeListQuery) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpenstackVolumeTypeListQuery) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *OpenstackVolumeTypeListQuery) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *OpenstackVolumeTypeListQuery) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOpenStackUser

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### SetOpenStackUserNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### SetOpenStackPasswordNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### SetOpenStackUrlNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackDomain

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.


### SetOpenStackDomainNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetOpenStackRegion

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.


### SetOpenStackRegionNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenStackRegionNil(b bool)`

 SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil

### UnsetOpenStackRegion
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenStackRegion()`

UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
### GetApplicationCredEnabled

`func (o *OpenstackVolumeTypeListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenstackVolumeTypeListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenstackVolumeTypeListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.


### GetIsAdmin

`func (o *OpenstackVolumeTypeListQuery) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenstackVolumeTypeListQuery) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenstackVolumeTypeListQuery) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.


### GetOpenstackProject

`func (o *OpenstackVolumeTypeListQuery) GetOpenstackProject() string`

GetOpenstackProject returns the OpenstackProject field if non-nil, zero value otherwise.

### GetOpenstackProjectOk

`func (o *OpenstackVolumeTypeListQuery) GetOpenstackProjectOk() (*string, bool)`

GetOpenstackProjectOk returns a tuple with the OpenstackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackProject

`func (o *OpenstackVolumeTypeListQuery) SetOpenstackProject(v string)`

SetOpenstackProject sets OpenstackProject field to given value.


### SetOpenstackProjectNil

`func (o *OpenstackVolumeTypeListQuery) SetOpenstackProjectNil(b bool)`

 SetOpenstackProjectNil sets the value for OpenstackProject to be an explicit nil

### UnsetOpenstackProject
`func (o *OpenstackVolumeTypeListQuery) UnsetOpenstackProject()`

UnsetOpenstackProject ensures that no value is present for OpenstackProject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


