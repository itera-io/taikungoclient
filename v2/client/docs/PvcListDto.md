# PvcListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Status** | **NullableString** |  | 
**Volume** | **NullableString** |  | 
**Capacity** | **NullableString** |  | 
**AccessModes** | **NullableString** |  | 
**StorageClass** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewPvcListDto

`func NewPvcListDto(name NullableString, namespace NullableString, status NullableString, volume NullableString, capacity NullableString, accessModes NullableString, storageClass NullableString, createdAt NullableString, ) *PvcListDto`

NewPvcListDto instantiates a new PvcListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvcListDtoWithDefaults

`func NewPvcListDtoWithDefaults() *PvcListDto`

NewPvcListDtoWithDefaults instantiates a new PvcListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PvcListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvcListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvcListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PvcListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PvcListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *PvcListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PvcListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PvcListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *PvcListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *PvcListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStatus

`func (o *PvcListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PvcListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PvcListDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *PvcListDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PvcListDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVolume

`func (o *PvcListDto) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *PvcListDto) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *PvcListDto) SetVolume(v string)`

SetVolume sets Volume field to given value.


### SetVolumeNil

`func (o *PvcListDto) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *PvcListDto) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil
### GetCapacity

`func (o *PvcListDto) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *PvcListDto) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *PvcListDto) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *PvcListDto) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *PvcListDto) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetAccessModes

`func (o *PvcListDto) GetAccessModes() string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *PvcListDto) GetAccessModesOk() (*string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *PvcListDto) SetAccessModes(v string)`

SetAccessModes sets AccessModes field to given value.


### SetAccessModesNil

`func (o *PvcListDto) SetAccessModesNil(b bool)`

 SetAccessModesNil sets the value for AccessModes to be an explicit nil

### UnsetAccessModes
`func (o *PvcListDto) UnsetAccessModes()`

UnsetAccessModes ensures that no value is present for AccessModes, not even an explicit nil
### GetStorageClass

`func (o *PvcListDto) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *PvcListDto) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *PvcListDto) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *PvcListDto) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *PvcListDto) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetCreatedAt

`func (o *PvcListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PvcListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PvcListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PvcListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PvcListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


