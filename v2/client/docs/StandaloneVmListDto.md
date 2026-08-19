# StandaloneVmListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**FlavorId** | **string** |  | 
**VolumeSize** | **int64** |  | 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**Ram** | **int64** |  | 
**Cpu** | **int32** |  | 
**VolumeType** | **NullableString** |  | 
**PublicIpEnabled** | **bool** |  | 
**PublicIp** | **NullableString** |  | 
**IpAddress** | **NullableString** |  | 
**CloudType** | **map[string]interface{}** |  | 
**ImageName** | **NullableString** |  | 
**Revision** | **int32** |  | 
**IsWindows** | **bool** |  | 
**Status** | [**StandAloneVmStatus**](StandAloneVmStatus.md) |  | 
**ProjectName** | **string** |  | 
**ProjectId** | **int32** |  | 
**StandAloneProfile** | [**StandaloneProfileListDto**](StandaloneProfileListDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**CreatedBy** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 

## Methods

### NewStandaloneVmListDto

`func NewStandaloneVmListDto(id int32, name string, flavorId string, volumeSize int64, organizationName string, organizationId int32, ram int64, cpu int32, volumeType NullableString, publicIpEnabled bool, publicIp NullableString, ipAddress NullableString, cloudType map[string]interface{}, imageName NullableString, revision int32, isWindows bool, status StandAloneVmStatus, projectName string, projectId int32, standAloneProfile StandaloneProfileListDto, createdAt NullableString, createdBy NullableString, lastModified NullableString, ) *StandaloneVmListDto`

NewStandaloneVmListDto instantiates a new StandaloneVmListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneVmListDtoWithDefaults

`func NewStandaloneVmListDtoWithDefaults() *StandaloneVmListDto`

NewStandaloneVmListDtoWithDefaults instantiates a new StandaloneVmListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneVmListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneVmListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneVmListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandaloneVmListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneVmListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneVmListDto) SetName(v string)`

SetName sets Name field to given value.


### GetFlavorId

`func (o *StandaloneVmListDto) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *StandaloneVmListDto) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *StandaloneVmListDto) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetVolumeSize

`func (o *StandaloneVmListDto) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *StandaloneVmListDto) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *StandaloneVmListDto) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.


### GetOrganizationName

`func (o *StandaloneVmListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *StandaloneVmListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *StandaloneVmListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *StandaloneVmListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *StandaloneVmListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *StandaloneVmListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetRam

`func (o *StandaloneVmListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *StandaloneVmListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *StandaloneVmListDto) SetRam(v int64)`

SetRam sets Ram field to given value.


### GetCpu

`func (o *StandaloneVmListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *StandaloneVmListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *StandaloneVmListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetVolumeType

`func (o *StandaloneVmListDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandaloneVmListDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandaloneVmListDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### SetVolumeTypeNil

`func (o *StandaloneVmListDto) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *StandaloneVmListDto) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetPublicIpEnabled

`func (o *StandaloneVmListDto) GetPublicIpEnabled() bool`

GetPublicIpEnabled returns the PublicIpEnabled field if non-nil, zero value otherwise.

### GetPublicIpEnabledOk

`func (o *StandaloneVmListDto) GetPublicIpEnabledOk() (*bool, bool)`

GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpEnabled

`func (o *StandaloneVmListDto) SetPublicIpEnabled(v bool)`

SetPublicIpEnabled sets PublicIpEnabled field to given value.


### GetPublicIp

`func (o *StandaloneVmListDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *StandaloneVmListDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *StandaloneVmListDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### SetPublicIpNil

`func (o *StandaloneVmListDto) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *StandaloneVmListDto) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetIpAddress

`func (o *StandaloneVmListDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StandaloneVmListDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StandaloneVmListDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *StandaloneVmListDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *StandaloneVmListDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetCloudType

`func (o *StandaloneVmListDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *StandaloneVmListDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *StandaloneVmListDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetImageName

`func (o *StandaloneVmListDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *StandaloneVmListDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *StandaloneVmListDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### SetImageNameNil

`func (o *StandaloneVmListDto) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *StandaloneVmListDto) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetRevision

`func (o *StandaloneVmListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StandaloneVmListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StandaloneVmListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetIsWindows

`func (o *StandaloneVmListDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *StandaloneVmListDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *StandaloneVmListDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.


### GetStatus

`func (o *StandaloneVmListDto) GetStatus() StandAloneVmStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandaloneVmListDto) GetStatusOk() (*StandAloneVmStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandaloneVmListDto) SetStatus(v StandAloneVmStatus)`

SetStatus sets Status field to given value.


### GetProjectName

`func (o *StandaloneVmListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *StandaloneVmListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *StandaloneVmListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetProjectId

`func (o *StandaloneVmListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StandaloneVmListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StandaloneVmListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetStandAloneProfile

`func (o *StandaloneVmListDto) GetStandAloneProfile() StandaloneProfileListDto`

GetStandAloneProfile returns the StandAloneProfile field if non-nil, zero value otherwise.

### GetStandAloneProfileOk

`func (o *StandaloneVmListDto) GetStandAloneProfileOk() (*StandaloneProfileListDto, bool)`

GetStandAloneProfileOk returns a tuple with the StandAloneProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfile

`func (o *StandaloneVmListDto) SetStandAloneProfile(v StandaloneProfileListDto)`

SetStandAloneProfile sets StandAloneProfile field to given value.


### GetCreatedAt

`func (o *StandaloneVmListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StandaloneVmListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StandaloneVmListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StandaloneVmListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StandaloneVmListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *StandaloneVmListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StandaloneVmListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StandaloneVmListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *StandaloneVmListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *StandaloneVmListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *StandaloneVmListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StandaloneVmListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StandaloneVmListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *StandaloneVmListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *StandaloneVmListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


