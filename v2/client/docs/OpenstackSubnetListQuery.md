# OpenstackSubnetListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | **NullableString** |  | 
**OpenStackPassword** | **NullableString** |  | 
**OpenStackUrl** | **NullableString** |  | 
**OpenStackProject** | **NullableString** |  | 
**OpenStackProjectId** | **NullableString** |  | 
**OpenStackDomain** | **NullableString** |  | 
**OpenStackRegion** | **NullableString** |  | 
**ApplicationCredEnabled** | **bool** |  | 

## Methods

### NewOpenstackSubnetListQuery

`func NewOpenstackSubnetListQuery(openStackUser NullableString, openStackPassword NullableString, openStackUrl NullableString, openStackProject NullableString, openStackProjectId NullableString, openStackDomain NullableString, openStackRegion NullableString, applicationCredEnabled bool, ) *OpenstackSubnetListQuery`

NewOpenstackSubnetListQuery instantiates a new OpenstackSubnetListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackSubnetListQueryWithDefaults

`func NewOpenstackSubnetListQueryWithDefaults() *OpenstackSubnetListQuery`

NewOpenstackSubnetListQueryWithDefaults instantiates a new OpenstackSubnetListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *OpenstackSubnetListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenstackSubnetListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenstackSubnetListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### SetOpenStackUserNil

`func (o *OpenstackSubnetListQuery) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *OpenstackSubnetListQuery) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *OpenstackSubnetListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenstackSubnetListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenstackSubnetListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### SetOpenStackPasswordNil

`func (o *OpenstackSubnetListQuery) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *OpenstackSubnetListQuery) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *OpenstackSubnetListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenstackSubnetListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenstackSubnetListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### SetOpenStackUrlNil

`func (o *OpenstackSubnetListQuery) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *OpenstackSubnetListQuery) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackProject

`func (o *OpenstackSubnetListQuery) GetOpenStackProject() string`

GetOpenStackProject returns the OpenStackProject field if non-nil, zero value otherwise.

### GetOpenStackProjectOk

`func (o *OpenstackSubnetListQuery) GetOpenStackProjectOk() (*string, bool)`

GetOpenStackProjectOk returns a tuple with the OpenStackProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackProject

`func (o *OpenstackSubnetListQuery) SetOpenStackProject(v string)`

SetOpenStackProject sets OpenStackProject field to given value.


### SetOpenStackProjectNil

`func (o *OpenstackSubnetListQuery) SetOpenStackProjectNil(b bool)`

 SetOpenStackProjectNil sets the value for OpenStackProject to be an explicit nil

### UnsetOpenStackProject
`func (o *OpenstackSubnetListQuery) UnsetOpenStackProject()`

UnsetOpenStackProject ensures that no value is present for OpenStackProject, not even an explicit nil
### GetOpenStackProjectId

`func (o *OpenstackSubnetListQuery) GetOpenStackProjectId() string`

GetOpenStackProjectId returns the OpenStackProjectId field if non-nil, zero value otherwise.

### GetOpenStackProjectIdOk

`func (o *OpenstackSubnetListQuery) GetOpenStackProjectIdOk() (*string, bool)`

GetOpenStackProjectIdOk returns a tuple with the OpenStackProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackProjectId

`func (o *OpenstackSubnetListQuery) SetOpenStackProjectId(v string)`

SetOpenStackProjectId sets OpenStackProjectId field to given value.


### SetOpenStackProjectIdNil

`func (o *OpenstackSubnetListQuery) SetOpenStackProjectIdNil(b bool)`

 SetOpenStackProjectIdNil sets the value for OpenStackProjectId to be an explicit nil

### UnsetOpenStackProjectId
`func (o *OpenstackSubnetListQuery) UnsetOpenStackProjectId()`

UnsetOpenStackProjectId ensures that no value is present for OpenStackProjectId, not even an explicit nil
### GetOpenStackDomain

`func (o *OpenstackSubnetListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenstackSubnetListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenstackSubnetListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.


### SetOpenStackDomainNil

`func (o *OpenstackSubnetListQuery) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *OpenstackSubnetListQuery) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetOpenStackRegion

`func (o *OpenstackSubnetListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenstackSubnetListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenstackSubnetListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.


### SetOpenStackRegionNil

`func (o *OpenstackSubnetListQuery) SetOpenStackRegionNil(b bool)`

 SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil

### UnsetOpenStackRegion
`func (o *OpenstackSubnetListQuery) UnsetOpenStackRegion()`

UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
### GetApplicationCredEnabled

`func (o *OpenstackSubnetListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenstackSubnetListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenstackSubnetListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


