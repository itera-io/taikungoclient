# StandaloneVmsListForDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ImageName** | **string** |  | 
**ImageId** | **string** |  | 
**Status** | **string** |  | 
**CloudInit** | **NullableString** |  | 
**VolumeType** | **string** |  | 
**VolumeSize** | **int64** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**SshPublicKey** | **string** |  | 
**CurrentFlavor** | **string** |  | 
**TargetFlavor** | **NullableString** |  | 
**PublicIpEnabled** | **bool** |  | 
**PublicIp** | **NullableString** |  | 
**Hypervisor** | **NullableString** |  | 
**HypervisorId** | **NullableString** |  | 
**IpAddress** | **NullableString** |  | 
**SpotPrice** | **float64** |  | 
**SpotInstance** | **bool** |  | 
**AvailabilityZone** | **NullableString** |  | 
**ActionButtons** | [**StandaloneVisibilityDto**](StandaloneVisibilityDto.md) |  | 
**IsWindows** | **bool** |  | 
**Disks** | [**[]StandAloneVmDiskForDetailsDto**](StandAloneVmDiskForDetailsDto.md) |  | 
**StandAloneMetaDatas** | [**[]StandAloneMetaDataDtoForVm**](StandAloneMetaDataDtoForVm.md) |  | 
**Profile** | [**StandAloneProfileForDetailsDto**](StandAloneProfileForDetailsDto.md) |  | 

## Methods

### NewStandaloneVmsListForDetailsDto

`func NewStandaloneVmsListForDetailsDto(id int32, name string, imageName string, imageId string, status string, cloudInit NullableString, volumeType string, volumeSize int64, createdAt string, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, sshPublicKey string, currentFlavor string, targetFlavor NullableString, publicIpEnabled bool, publicIp NullableString, hypervisor NullableString, hypervisorId NullableString, ipAddress NullableString, spotPrice float64, spotInstance bool, availabilityZone NullableString, actionButtons StandaloneVisibilityDto, isWindows bool, disks []StandAloneVmDiskForDetailsDto, standAloneMetaDatas []StandAloneMetaDataDtoForVm, profile StandAloneProfileForDetailsDto, ) *StandaloneVmsListForDetailsDto`

NewStandaloneVmsListForDetailsDto instantiates a new StandaloneVmsListForDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneVmsListForDetailsDtoWithDefaults

`func NewStandaloneVmsListForDetailsDtoWithDefaults() *StandaloneVmsListForDetailsDto`

NewStandaloneVmsListForDetailsDtoWithDefaults instantiates a new StandaloneVmsListForDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneVmsListForDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneVmsListForDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneVmsListForDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandaloneVmsListForDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneVmsListForDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneVmsListForDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetImageName

`func (o *StandaloneVmsListForDetailsDto) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *StandaloneVmsListForDetailsDto) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *StandaloneVmsListForDetailsDto) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageId

`func (o *StandaloneVmsListForDetailsDto) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *StandaloneVmsListForDetailsDto) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *StandaloneVmsListForDetailsDto) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetStatus

`func (o *StandaloneVmsListForDetailsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandaloneVmsListForDetailsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandaloneVmsListForDetailsDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCloudInit

`func (o *StandaloneVmsListForDetailsDto) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *StandaloneVmsListForDetailsDto) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *StandaloneVmsListForDetailsDto) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.


### SetCloudInitNil

`func (o *StandaloneVmsListForDetailsDto) SetCloudInitNil(b bool)`

 SetCloudInitNil sets the value for CloudInit to be an explicit nil

### UnsetCloudInit
`func (o *StandaloneVmsListForDetailsDto) UnsetCloudInit()`

UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
### GetVolumeType

`func (o *StandaloneVmsListForDetailsDto) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StandaloneVmsListForDetailsDto) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StandaloneVmsListForDetailsDto) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeSize

`func (o *StandaloneVmsListForDetailsDto) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *StandaloneVmsListForDetailsDto) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *StandaloneVmsListForDetailsDto) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.


### GetCreatedAt

`func (o *StandaloneVmsListForDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StandaloneVmsListForDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StandaloneVmsListForDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *StandaloneVmsListForDetailsDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StandaloneVmsListForDetailsDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StandaloneVmsListForDetailsDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *StandaloneVmsListForDetailsDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StandaloneVmsListForDetailsDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *StandaloneVmsListForDetailsDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *StandaloneVmsListForDetailsDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *StandaloneVmsListForDetailsDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *StandaloneVmsListForDetailsDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetSshPublicKey

`func (o *StandaloneVmsListForDetailsDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *StandaloneVmsListForDetailsDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *StandaloneVmsListForDetailsDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.


### GetCurrentFlavor

`func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavor() string`

GetCurrentFlavor returns the CurrentFlavor field if non-nil, zero value otherwise.

### GetCurrentFlavorOk

`func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavorOk() (*string, bool)`

GetCurrentFlavorOk returns a tuple with the CurrentFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFlavor

`func (o *StandaloneVmsListForDetailsDto) SetCurrentFlavor(v string)`

SetCurrentFlavor sets CurrentFlavor field to given value.


### GetTargetFlavor

`func (o *StandaloneVmsListForDetailsDto) GetTargetFlavor() string`

GetTargetFlavor returns the TargetFlavor field if non-nil, zero value otherwise.

### GetTargetFlavorOk

`func (o *StandaloneVmsListForDetailsDto) GetTargetFlavorOk() (*string, bool)`

GetTargetFlavorOk returns a tuple with the TargetFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFlavor

`func (o *StandaloneVmsListForDetailsDto) SetTargetFlavor(v string)`

SetTargetFlavor sets TargetFlavor field to given value.


### SetTargetFlavorNil

`func (o *StandaloneVmsListForDetailsDto) SetTargetFlavorNil(b bool)`

 SetTargetFlavorNil sets the value for TargetFlavor to be an explicit nil

### UnsetTargetFlavor
`func (o *StandaloneVmsListForDetailsDto) UnsetTargetFlavor()`

UnsetTargetFlavor ensures that no value is present for TargetFlavor, not even an explicit nil
### GetPublicIpEnabled

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabled() bool`

GetPublicIpEnabled returns the PublicIpEnabled field if non-nil, zero value otherwise.

### GetPublicIpEnabledOk

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabledOk() (*bool, bool)`

GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpEnabled

`func (o *StandaloneVmsListForDetailsDto) SetPublicIpEnabled(v bool)`

SetPublicIpEnabled sets PublicIpEnabled field to given value.


### GetPublicIp

`func (o *StandaloneVmsListForDetailsDto) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *StandaloneVmsListForDetailsDto) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *StandaloneVmsListForDetailsDto) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### SetPublicIpNil

`func (o *StandaloneVmsListForDetailsDto) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *StandaloneVmsListForDetailsDto) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetHypervisor

`func (o *StandaloneVmsListForDetailsDto) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *StandaloneVmsListForDetailsDto) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *StandaloneVmsListForDetailsDto) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.


### SetHypervisorNil

`func (o *StandaloneVmsListForDetailsDto) SetHypervisorNil(b bool)`

 SetHypervisorNil sets the value for Hypervisor to be an explicit nil

### UnsetHypervisor
`func (o *StandaloneVmsListForDetailsDto) UnsetHypervisor()`

UnsetHypervisor ensures that no value is present for Hypervisor, not even an explicit nil
### GetHypervisorId

`func (o *StandaloneVmsListForDetailsDto) GetHypervisorId() string`

GetHypervisorId returns the HypervisorId field if non-nil, zero value otherwise.

### GetHypervisorIdOk

`func (o *StandaloneVmsListForDetailsDto) GetHypervisorIdOk() (*string, bool)`

GetHypervisorIdOk returns a tuple with the HypervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorId

`func (o *StandaloneVmsListForDetailsDto) SetHypervisorId(v string)`

SetHypervisorId sets HypervisorId field to given value.


### SetHypervisorIdNil

`func (o *StandaloneVmsListForDetailsDto) SetHypervisorIdNil(b bool)`

 SetHypervisorIdNil sets the value for HypervisorId to be an explicit nil

### UnsetHypervisorId
`func (o *StandaloneVmsListForDetailsDto) UnsetHypervisorId()`

UnsetHypervisorId ensures that no value is present for HypervisorId, not even an explicit nil
### GetIpAddress

`func (o *StandaloneVmsListForDetailsDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *StandaloneVmsListForDetailsDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *StandaloneVmsListForDetailsDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *StandaloneVmsListForDetailsDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *StandaloneVmsListForDetailsDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetSpotPrice

`func (o *StandaloneVmsListForDetailsDto) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *StandaloneVmsListForDetailsDto) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *StandaloneVmsListForDetailsDto) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.


### GetSpotInstance

`func (o *StandaloneVmsListForDetailsDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *StandaloneVmsListForDetailsDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *StandaloneVmsListForDetailsDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.


### GetAvailabilityZone

`func (o *StandaloneVmsListForDetailsDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *StandaloneVmsListForDetailsDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *StandaloneVmsListForDetailsDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *StandaloneVmsListForDetailsDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *StandaloneVmsListForDetailsDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetActionButtons

`func (o *StandaloneVmsListForDetailsDto) GetActionButtons() StandaloneVisibilityDto`

GetActionButtons returns the ActionButtons field if non-nil, zero value otherwise.

### GetActionButtonsOk

`func (o *StandaloneVmsListForDetailsDto) GetActionButtonsOk() (*StandaloneVisibilityDto, bool)`

GetActionButtonsOk returns a tuple with the ActionButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionButtons

`func (o *StandaloneVmsListForDetailsDto) SetActionButtons(v StandaloneVisibilityDto)`

SetActionButtons sets ActionButtons field to given value.


### GetIsWindows

`func (o *StandaloneVmsListForDetailsDto) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *StandaloneVmsListForDetailsDto) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *StandaloneVmsListForDetailsDto) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.


### GetDisks

`func (o *StandaloneVmsListForDetailsDto) GetDisks() []StandAloneVmDiskForDetailsDto`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *StandaloneVmsListForDetailsDto) GetDisksOk() (*[]StandAloneVmDiskForDetailsDto, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *StandaloneVmsListForDetailsDto) SetDisks(v []StandAloneVmDiskForDetailsDto)`

SetDisks sets Disks field to given value.


### GetStandAloneMetaDatas

`func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatas() []StandAloneMetaDataDtoForVm`

GetStandAloneMetaDatas returns the StandAloneMetaDatas field if non-nil, zero value otherwise.

### GetStandAloneMetaDatasOk

`func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatasOk() (*[]StandAloneMetaDataDtoForVm, bool)`

GetStandAloneMetaDatasOk returns a tuple with the StandAloneMetaDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneMetaDatas

`func (o *StandaloneVmsListForDetailsDto) SetStandAloneMetaDatas(v []StandAloneMetaDataDtoForVm)`

SetStandAloneMetaDatas sets StandAloneMetaDatas field to given value.


### GetProfile

`func (o *StandaloneVmsListForDetailsDto) GetProfile() StandAloneProfileForDetailsDto`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *StandaloneVmsListForDetailsDto) GetProfileOk() (*StandAloneProfileForDetailsDto, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *StandaloneVmsListForDetailsDto) SetProfile(v StandAloneProfileForDetailsDto)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


