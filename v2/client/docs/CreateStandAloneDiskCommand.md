# CreateStandAloneDiskCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandaloneVmId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateStandAloneDiskCommand

`func NewCreateStandAloneDiskCommand() *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommand instantiates a new CreateStandAloneDiskCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandAloneDiskCommandWithDefaults

`func NewCreateStandAloneDiskCommandWithDefaults() *CreateStandAloneDiskCommand`

NewCreateStandAloneDiskCommandWithDefaults instantiates a new CreateStandAloneDiskCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmId() int32`

GetStandaloneVmId returns the StandaloneVmId field if non-nil, zero value otherwise.

### GetStandaloneVmIdOk

`func (o *CreateStandAloneDiskCommand) GetStandaloneVmIdOk() (*int32, bool)`

GetStandaloneVmIdOk returns a tuple with the StandaloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmId

`func (o *CreateStandAloneDiskCommand) SetStandaloneVmId(v int32)`

SetStandaloneVmId sets StandaloneVmId field to given value.

### HasStandaloneVmId

`func (o *CreateStandAloneDiskCommand) HasStandaloneVmId() bool`

HasStandaloneVmId returns a boolean if a field has been set.

### GetName

`func (o *CreateStandAloneDiskCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandAloneDiskCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandAloneDiskCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStandAloneDiskCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateStandAloneDiskCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateStandAloneDiskCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *CreateStandAloneDiskCommand) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStandAloneDiskCommand) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStandAloneDiskCommand) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateStandAloneDiskCommand) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *CreateStandAloneDiskCommand) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateStandAloneDiskCommand) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreateStandAloneDiskCommand) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreateStandAloneDiskCommand) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CreateStandAloneDiskCommand) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CreateStandAloneDiskCommand) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


