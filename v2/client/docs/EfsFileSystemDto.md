# EfsFileSystemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemId** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**LifecycleState** | **NullableString** |  | 
**SizeInBytes** | **int64** |  | 

## Methods

### NewEfsFileSystemDto

`func NewEfsFileSystemDto(filesystemId NullableString, name NullableString, lifecycleState NullableString, sizeInBytes int64, ) *EfsFileSystemDto`

NewEfsFileSystemDto instantiates a new EfsFileSystemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEfsFileSystemDtoWithDefaults

`func NewEfsFileSystemDtoWithDefaults() *EfsFileSystemDto`

NewEfsFileSystemDtoWithDefaults instantiates a new EfsFileSystemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemId

`func (o *EfsFileSystemDto) GetFilesystemId() string`

GetFilesystemId returns the FilesystemId field if non-nil, zero value otherwise.

### GetFilesystemIdOk

`func (o *EfsFileSystemDto) GetFilesystemIdOk() (*string, bool)`

GetFilesystemIdOk returns a tuple with the FilesystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemId

`func (o *EfsFileSystemDto) SetFilesystemId(v string)`

SetFilesystemId sets FilesystemId field to given value.


### SetFilesystemIdNil

`func (o *EfsFileSystemDto) SetFilesystemIdNil(b bool)`

 SetFilesystemIdNil sets the value for FilesystemId to be an explicit nil

### UnsetFilesystemId
`func (o *EfsFileSystemDto) UnsetFilesystemId()`

UnsetFilesystemId ensures that no value is present for FilesystemId, not even an explicit nil
### GetName

`func (o *EfsFileSystemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EfsFileSystemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EfsFileSystemDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *EfsFileSystemDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EfsFileSystemDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLifecycleState

`func (o *EfsFileSystemDto) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *EfsFileSystemDto) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *EfsFileSystemDto) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.


### SetLifecycleStateNil

`func (o *EfsFileSystemDto) SetLifecycleStateNil(b bool)`

 SetLifecycleStateNil sets the value for LifecycleState to be an explicit nil

### UnsetLifecycleState
`func (o *EfsFileSystemDto) UnsetLifecycleState()`

UnsetLifecycleState ensures that no value is present for LifecycleState, not even an explicit nil
### GetSizeInBytes

`func (o *EfsFileSystemDto) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *EfsFileSystemDto) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *EfsFileSystemDto) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


