# CreateStandAloneVmCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**FlavorName** | Pointer to **NullableString** |  | [optional] 
**VolumeSize** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PublicIpEnabled** | Pointer to **bool** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**CloudInit** | Pointer to **NullableString** |  | [optional] 
**StandAloneProfileId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**SpotPrice** | Pointer to **NullableFloat64** |  | [optional] 
**SpotInstance** | Pointer to **bool** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**Hypervisor** | Pointer to **NullableString** |  | [optional] 
**UseLocalDisk** | Pointer to **bool** |  | [optional] 
**StandAloneVmDisks** | Pointer to [**[]StandAloneVmDiskDto**](StandAloneVmDiskDto.md) |  | [optional] 
**StandAloneMetaDatas** | Pointer to [**[]StandAloneMetaDataDto**](StandAloneMetaDataDto.md) |  | [optional] 

## Methods

### NewCreateStandAloneVmCommand

`func NewCreateStandAloneVmCommand() *CreateStandAloneVmCommand`

NewCreateStandAloneVmCommand instantiates a new CreateStandAloneVmCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandAloneVmCommandWithDefaults

`func NewCreateStandAloneVmCommandWithDefaults() *CreateStandAloneVmCommand`

NewCreateStandAloneVmCommandWithDefaults instantiates a new CreateStandAloneVmCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStandAloneVmCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandAloneVmCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandAloneVmCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStandAloneVmCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateStandAloneVmCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateStandAloneVmCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFlavorName

`func (o *CreateStandAloneVmCommand) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *CreateStandAloneVmCommand) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *CreateStandAloneVmCommand) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.

### HasFlavorName

`func (o *CreateStandAloneVmCommand) HasFlavorName() bool`

HasFlavorName returns a boolean if a field has been set.

### SetFlavorNameNil

`func (o *CreateStandAloneVmCommand) SetFlavorNameNil(b bool)`

 SetFlavorNameNil sets the value for FlavorName to be an explicit nil

### UnsetFlavorName
`func (o *CreateStandAloneVmCommand) UnsetFlavorName()`

UnsetFlavorName ensures that no value is present for FlavorName, not even an explicit nil
### GetVolumeSize

`func (o *CreateStandAloneVmCommand) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *CreateStandAloneVmCommand) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *CreateStandAloneVmCommand) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *CreateStandAloneVmCommand) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *CreateStandAloneVmCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateStandAloneVmCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateStandAloneVmCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreateStandAloneVmCommand) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CreateStandAloneVmCommand) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CreateStandAloneVmCommand) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetUsername

`func (o *CreateStandAloneVmCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateStandAloneVmCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateStandAloneVmCommand) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateStandAloneVmCommand) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateStandAloneVmCommand) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateStandAloneVmCommand) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateStandAloneVmCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateStandAloneVmCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateStandAloneVmCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateStandAloneVmCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateStandAloneVmCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateStandAloneVmCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPublicIpEnabled

`func (o *CreateStandAloneVmCommand) GetPublicIpEnabled() bool`

GetPublicIpEnabled returns the PublicIpEnabled field if non-nil, zero value otherwise.

### GetPublicIpEnabledOk

`func (o *CreateStandAloneVmCommand) GetPublicIpEnabledOk() (*bool, bool)`

GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpEnabled

`func (o *CreateStandAloneVmCommand) SetPublicIpEnabled(v bool)`

SetPublicIpEnabled sets PublicIpEnabled field to given value.

### HasPublicIpEnabled

`func (o *CreateStandAloneVmCommand) HasPublicIpEnabled() bool`

HasPublicIpEnabled returns a boolean if a field has been set.

### GetImage

`func (o *CreateStandAloneVmCommand) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateStandAloneVmCommand) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateStandAloneVmCommand) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateStandAloneVmCommand) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *CreateStandAloneVmCommand) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *CreateStandAloneVmCommand) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetCloudInit

`func (o *CreateStandAloneVmCommand) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *CreateStandAloneVmCommand) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *CreateStandAloneVmCommand) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *CreateStandAloneVmCommand) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### SetCloudInitNil

`func (o *CreateStandAloneVmCommand) SetCloudInitNil(b bool)`

 SetCloudInitNil sets the value for CloudInit to be an explicit nil

### UnsetCloudInit
`func (o *CreateStandAloneVmCommand) UnsetCloudInit()`

UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
### GetStandAloneProfileId

`func (o *CreateStandAloneVmCommand) GetStandAloneProfileId() int32`

GetStandAloneProfileId returns the StandAloneProfileId field if non-nil, zero value otherwise.

### GetStandAloneProfileIdOk

`func (o *CreateStandAloneVmCommand) GetStandAloneProfileIdOk() (*int32, bool)`

GetStandAloneProfileIdOk returns a tuple with the StandAloneProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfileId

`func (o *CreateStandAloneVmCommand) SetStandAloneProfileId(v int32)`

SetStandAloneProfileId sets StandAloneProfileId field to given value.

### HasStandAloneProfileId

`func (o *CreateStandAloneVmCommand) HasStandAloneProfileId() bool`

HasStandAloneProfileId returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateStandAloneVmCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateStandAloneVmCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateStandAloneVmCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateStandAloneVmCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCount

`func (o *CreateStandAloneVmCommand) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateStandAloneVmCommand) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateStandAloneVmCommand) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateStandAloneVmCommand) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSpotPrice

`func (o *CreateStandAloneVmCommand) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *CreateStandAloneVmCommand) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *CreateStandAloneVmCommand) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *CreateStandAloneVmCommand) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### SetSpotPriceNil

`func (o *CreateStandAloneVmCommand) SetSpotPriceNil(b bool)`

 SetSpotPriceNil sets the value for SpotPrice to be an explicit nil

### UnsetSpotPrice
`func (o *CreateStandAloneVmCommand) UnsetSpotPrice()`

UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
### GetSpotInstance

`func (o *CreateStandAloneVmCommand) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *CreateStandAloneVmCommand) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *CreateStandAloneVmCommand) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *CreateStandAloneVmCommand) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateStandAloneVmCommand) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateStandAloneVmCommand) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateStandAloneVmCommand) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateStandAloneVmCommand) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *CreateStandAloneVmCommand) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *CreateStandAloneVmCommand) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetHypervisor

`func (o *CreateStandAloneVmCommand) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *CreateStandAloneVmCommand) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *CreateStandAloneVmCommand) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *CreateStandAloneVmCommand) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### SetHypervisorNil

`func (o *CreateStandAloneVmCommand) SetHypervisorNil(b bool)`

 SetHypervisorNil sets the value for Hypervisor to be an explicit nil

### UnsetHypervisor
`func (o *CreateStandAloneVmCommand) UnsetHypervisor()`

UnsetHypervisor ensures that no value is present for Hypervisor, not even an explicit nil
### GetUseLocalDisk

`func (o *CreateStandAloneVmCommand) GetUseLocalDisk() bool`

GetUseLocalDisk returns the UseLocalDisk field if non-nil, zero value otherwise.

### GetUseLocalDiskOk

`func (o *CreateStandAloneVmCommand) GetUseLocalDiskOk() (*bool, bool)`

GetUseLocalDiskOk returns a tuple with the UseLocalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalDisk

`func (o *CreateStandAloneVmCommand) SetUseLocalDisk(v bool)`

SetUseLocalDisk sets UseLocalDisk field to given value.

### HasUseLocalDisk

`func (o *CreateStandAloneVmCommand) HasUseLocalDisk() bool`

HasUseLocalDisk returns a boolean if a field has been set.

### GetStandAloneVmDisks

`func (o *CreateStandAloneVmCommand) GetStandAloneVmDisks() []StandAloneVmDiskDto`

GetStandAloneVmDisks returns the StandAloneVmDisks field if non-nil, zero value otherwise.

### GetStandAloneVmDisksOk

`func (o *CreateStandAloneVmCommand) GetStandAloneVmDisksOk() (*[]StandAloneVmDiskDto, bool)`

GetStandAloneVmDisksOk returns a tuple with the StandAloneVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneVmDisks

`func (o *CreateStandAloneVmCommand) SetStandAloneVmDisks(v []StandAloneVmDiskDto)`

SetStandAloneVmDisks sets StandAloneVmDisks field to given value.

### HasStandAloneVmDisks

`func (o *CreateStandAloneVmCommand) HasStandAloneVmDisks() bool`

HasStandAloneVmDisks returns a boolean if a field has been set.

### SetStandAloneVmDisksNil

`func (o *CreateStandAloneVmCommand) SetStandAloneVmDisksNil(b bool)`

 SetStandAloneVmDisksNil sets the value for StandAloneVmDisks to be an explicit nil

### UnsetStandAloneVmDisks
`func (o *CreateStandAloneVmCommand) UnsetStandAloneVmDisks()`

UnsetStandAloneVmDisks ensures that no value is present for StandAloneVmDisks, not even an explicit nil
### GetStandAloneMetaDatas

`func (o *CreateStandAloneVmCommand) GetStandAloneMetaDatas() []StandAloneMetaDataDto`

GetStandAloneMetaDatas returns the StandAloneMetaDatas field if non-nil, zero value otherwise.

### GetStandAloneMetaDatasOk

`func (o *CreateStandAloneVmCommand) GetStandAloneMetaDatasOk() (*[]StandAloneMetaDataDto, bool)`

GetStandAloneMetaDatasOk returns a tuple with the StandAloneMetaDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneMetaDatas

`func (o *CreateStandAloneVmCommand) SetStandAloneMetaDatas(v []StandAloneMetaDataDto)`

SetStandAloneMetaDatas sets StandAloneMetaDatas field to given value.

### HasStandAloneMetaDatas

`func (o *CreateStandAloneVmCommand) HasStandAloneMetaDatas() bool`

HasStandAloneMetaDatas returns a boolean if a field has been set.

### SetStandAloneMetaDatasNil

`func (o *CreateStandAloneVmCommand) SetStandAloneMetaDatasNil(b bool)`

 SetStandAloneMetaDatasNil sets the value for StandAloneMetaDatas to be an explicit nil

### UnsetStandAloneMetaDatas
`func (o *CreateStandAloneVmCommand) UnsetStandAloneMetaDatas()`

UnsetStandAloneMetaDatas ensures that no value is present for StandAloneMetaDatas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


