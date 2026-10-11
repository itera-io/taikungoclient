# InternalExporterProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Status** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**CloudType** | **NullableInt32** |  | 
**ProjectAction** | **bool** |  | 

## Methods

### NewInternalExporterProjectDto

`func NewInternalExporterProjectDto(id int32, name NullableString, status int32, organizationId int32, organizationName NullableString, cloudType NullableInt32, projectAction bool, ) *InternalExporterProjectDto`

NewInternalExporterProjectDto instantiates a new InternalExporterProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalExporterProjectDtoWithDefaults

`func NewInternalExporterProjectDtoWithDefaults() *InternalExporterProjectDto`

NewInternalExporterProjectDtoWithDefaults instantiates a new InternalExporterProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalExporterProjectDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalExporterProjectDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalExporterProjectDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InternalExporterProjectDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalExporterProjectDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalExporterProjectDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *InternalExporterProjectDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternalExporterProjectDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *InternalExporterProjectDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalExporterProjectDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalExporterProjectDto) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetOrganizationId

`func (o *InternalExporterProjectDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InternalExporterProjectDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InternalExporterProjectDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *InternalExporterProjectDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InternalExporterProjectDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InternalExporterProjectDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *InternalExporterProjectDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *InternalExporterProjectDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetCloudType

`func (o *InternalExporterProjectDto) GetCloudType() int32`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *InternalExporterProjectDto) GetCloudTypeOk() (*int32, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *InternalExporterProjectDto) SetCloudType(v int32)`

SetCloudType sets CloudType field to given value.


### SetCloudTypeNil

`func (o *InternalExporterProjectDto) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *InternalExporterProjectDto) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetProjectAction

`func (o *InternalExporterProjectDto) GetProjectAction() bool`

GetProjectAction returns the ProjectAction field if non-nil, zero value otherwise.

### GetProjectActionOk

`func (o *InternalExporterProjectDto) GetProjectActionOk() (*bool, bool)`

GetProjectActionOk returns a tuple with the ProjectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAction

`func (o *InternalExporterProjectDto) SetProjectAction(v bool)`

SetProjectAction sets ProjectAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


