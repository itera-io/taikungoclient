# InternalExporterStandaloneVmDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Status** | **int32** |  | 
**FlavorId** | **NullableString** |  | 
**ProjectId** | **int32** |  | 
**ProjectName** | **NullableString** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**CloudType** | **int32** |  | 

## Methods

### NewInternalExporterStandaloneVmDto

`func NewInternalExporterStandaloneVmDto(id int32, name NullableString, status int32, flavorId NullableString, projectId int32, projectName NullableString, organizationId int32, organizationName NullableString, cloudType int32, ) *InternalExporterStandaloneVmDto`

NewInternalExporterStandaloneVmDto instantiates a new InternalExporterStandaloneVmDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalExporterStandaloneVmDtoWithDefaults

`func NewInternalExporterStandaloneVmDtoWithDefaults() *InternalExporterStandaloneVmDto`

NewInternalExporterStandaloneVmDtoWithDefaults instantiates a new InternalExporterStandaloneVmDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalExporterStandaloneVmDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalExporterStandaloneVmDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalExporterStandaloneVmDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InternalExporterStandaloneVmDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalExporterStandaloneVmDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalExporterStandaloneVmDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *InternalExporterStandaloneVmDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternalExporterStandaloneVmDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *InternalExporterStandaloneVmDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalExporterStandaloneVmDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalExporterStandaloneVmDto) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetFlavorId

`func (o *InternalExporterStandaloneVmDto) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *InternalExporterStandaloneVmDto) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *InternalExporterStandaloneVmDto) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### SetFlavorIdNil

`func (o *InternalExporterStandaloneVmDto) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *InternalExporterStandaloneVmDto) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetProjectId

`func (o *InternalExporterStandaloneVmDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InternalExporterStandaloneVmDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InternalExporterStandaloneVmDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *InternalExporterStandaloneVmDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InternalExporterStandaloneVmDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InternalExporterStandaloneVmDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *InternalExporterStandaloneVmDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *InternalExporterStandaloneVmDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetOrganizationId

`func (o *InternalExporterStandaloneVmDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InternalExporterStandaloneVmDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InternalExporterStandaloneVmDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *InternalExporterStandaloneVmDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InternalExporterStandaloneVmDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InternalExporterStandaloneVmDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *InternalExporterStandaloneVmDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *InternalExporterStandaloneVmDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCloudType

`func (o *InternalExporterStandaloneVmDto) GetCloudType() int32`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *InternalExporterStandaloneVmDto) GetCloudTypeOk() (*int32, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *InternalExporterStandaloneVmDto) SetCloudType(v int32)`

SetCloudType sets CloudType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


