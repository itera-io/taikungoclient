# OpenStackNetworkListQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenStackUser** | **NullableString** |  | 
**OpenStackPassword** | **NullableString** |  | 
**OpenStackUrl** | **NullableString** |  | 
**OpenStackProjectId** | **NullableString** |  | 
**OpenStackDomain** | **NullableString** |  | 
**OpenStackRegion** | **NullableString** |  | 
**ApplicationCredEnabled** | **bool** |  | 
**IsAdmin** | **bool** |  | 

## Methods

### NewOpenStackNetworkListQuery

`func NewOpenStackNetworkListQuery(openStackUser NullableString, openStackPassword NullableString, openStackUrl NullableString, openStackProjectId NullableString, openStackDomain NullableString, openStackRegion NullableString, applicationCredEnabled bool, isAdmin bool, ) *OpenStackNetworkListQuery`

NewOpenStackNetworkListQuery instantiates a new OpenStackNetworkListQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenStackNetworkListQueryWithDefaults

`func NewOpenStackNetworkListQueryWithDefaults() *OpenStackNetworkListQuery`

NewOpenStackNetworkListQueryWithDefaults instantiates a new OpenStackNetworkListQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenStackUser

`func (o *OpenStackNetworkListQuery) GetOpenStackUser() string`

GetOpenStackUser returns the OpenStackUser field if non-nil, zero value otherwise.

### GetOpenStackUserOk

`func (o *OpenStackNetworkListQuery) GetOpenStackUserOk() (*string, bool)`

GetOpenStackUserOk returns a tuple with the OpenStackUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUser

`func (o *OpenStackNetworkListQuery) SetOpenStackUser(v string)`

SetOpenStackUser sets OpenStackUser field to given value.


### SetOpenStackUserNil

`func (o *OpenStackNetworkListQuery) SetOpenStackUserNil(b bool)`

 SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil

### UnsetOpenStackUser
`func (o *OpenStackNetworkListQuery) UnsetOpenStackUser()`

UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
### GetOpenStackPassword

`func (o *OpenStackNetworkListQuery) GetOpenStackPassword() string`

GetOpenStackPassword returns the OpenStackPassword field if non-nil, zero value otherwise.

### GetOpenStackPasswordOk

`func (o *OpenStackNetworkListQuery) GetOpenStackPasswordOk() (*string, bool)`

GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackPassword

`func (o *OpenStackNetworkListQuery) SetOpenStackPassword(v string)`

SetOpenStackPassword sets OpenStackPassword field to given value.


### SetOpenStackPasswordNil

`func (o *OpenStackNetworkListQuery) SetOpenStackPasswordNil(b bool)`

 SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil

### UnsetOpenStackPassword
`func (o *OpenStackNetworkListQuery) UnsetOpenStackPassword()`

UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
### GetOpenStackUrl

`func (o *OpenStackNetworkListQuery) GetOpenStackUrl() string`

GetOpenStackUrl returns the OpenStackUrl field if non-nil, zero value otherwise.

### GetOpenStackUrlOk

`func (o *OpenStackNetworkListQuery) GetOpenStackUrlOk() (*string, bool)`

GetOpenStackUrlOk returns a tuple with the OpenStackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackUrl

`func (o *OpenStackNetworkListQuery) SetOpenStackUrl(v string)`

SetOpenStackUrl sets OpenStackUrl field to given value.


### SetOpenStackUrlNil

`func (o *OpenStackNetworkListQuery) SetOpenStackUrlNil(b bool)`

 SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil

### UnsetOpenStackUrl
`func (o *OpenStackNetworkListQuery) UnsetOpenStackUrl()`

UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
### GetOpenStackProjectId

`func (o *OpenStackNetworkListQuery) GetOpenStackProjectId() string`

GetOpenStackProjectId returns the OpenStackProjectId field if non-nil, zero value otherwise.

### GetOpenStackProjectIdOk

`func (o *OpenStackNetworkListQuery) GetOpenStackProjectIdOk() (*string, bool)`

GetOpenStackProjectIdOk returns a tuple with the OpenStackProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackProjectId

`func (o *OpenStackNetworkListQuery) SetOpenStackProjectId(v string)`

SetOpenStackProjectId sets OpenStackProjectId field to given value.


### SetOpenStackProjectIdNil

`func (o *OpenStackNetworkListQuery) SetOpenStackProjectIdNil(b bool)`

 SetOpenStackProjectIdNil sets the value for OpenStackProjectId to be an explicit nil

### UnsetOpenStackProjectId
`func (o *OpenStackNetworkListQuery) UnsetOpenStackProjectId()`

UnsetOpenStackProjectId ensures that no value is present for OpenStackProjectId, not even an explicit nil
### GetOpenStackDomain

`func (o *OpenStackNetworkListQuery) GetOpenStackDomain() string`

GetOpenStackDomain returns the OpenStackDomain field if non-nil, zero value otherwise.

### GetOpenStackDomainOk

`func (o *OpenStackNetworkListQuery) GetOpenStackDomainOk() (*string, bool)`

GetOpenStackDomainOk returns a tuple with the OpenStackDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackDomain

`func (o *OpenStackNetworkListQuery) SetOpenStackDomain(v string)`

SetOpenStackDomain sets OpenStackDomain field to given value.


### SetOpenStackDomainNil

`func (o *OpenStackNetworkListQuery) SetOpenStackDomainNil(b bool)`

 SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil

### UnsetOpenStackDomain
`func (o *OpenStackNetworkListQuery) UnsetOpenStackDomain()`

UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
### GetOpenStackRegion

`func (o *OpenStackNetworkListQuery) GetOpenStackRegion() string`

GetOpenStackRegion returns the OpenStackRegion field if non-nil, zero value otherwise.

### GetOpenStackRegionOk

`func (o *OpenStackNetworkListQuery) GetOpenStackRegionOk() (*string, bool)`

GetOpenStackRegionOk returns a tuple with the OpenStackRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStackRegion

`func (o *OpenStackNetworkListQuery) SetOpenStackRegion(v string)`

SetOpenStackRegion sets OpenStackRegion field to given value.


### SetOpenStackRegionNil

`func (o *OpenStackNetworkListQuery) SetOpenStackRegionNil(b bool)`

 SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil

### UnsetOpenStackRegion
`func (o *OpenStackNetworkListQuery) UnsetOpenStackRegion()`

UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
### GetApplicationCredEnabled

`func (o *OpenStackNetworkListQuery) GetApplicationCredEnabled() bool`

GetApplicationCredEnabled returns the ApplicationCredEnabled field if non-nil, zero value otherwise.

### GetApplicationCredEnabledOk

`func (o *OpenStackNetworkListQuery) GetApplicationCredEnabledOk() (*bool, bool)`

GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredEnabled

`func (o *OpenStackNetworkListQuery) SetApplicationCredEnabled(v bool)`

SetApplicationCredEnabled sets ApplicationCredEnabled field to given value.


### GetIsAdmin

`func (o *OpenStackNetworkListQuery) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OpenStackNetworkListQuery) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OpenStackNetworkListQuery) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


