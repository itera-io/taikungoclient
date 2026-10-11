# DaemonsetListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Desired** | **int32** |  | 
**Current** | **int32** |  | 
**Ready** | **int32** |  | 
**Available** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Image** | **NullableString** |  | 

## Methods

### NewDaemonsetListDto

`func NewDaemonsetListDto(status NullableString, name NullableString, desired int32, current int32, ready int32, available NullableString, createdAt NullableString, namespace NullableString, image NullableString, ) *DaemonsetListDto`

NewDaemonsetListDto instantiates a new DaemonsetListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonsetListDtoWithDefaults

`func NewDaemonsetListDtoWithDefaults() *DaemonsetListDto`

NewDaemonsetListDtoWithDefaults instantiates a new DaemonsetListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DaemonsetListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DaemonsetListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DaemonsetListDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *DaemonsetListDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DaemonsetListDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetName

`func (o *DaemonsetListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaemonsetListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaemonsetListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DaemonsetListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DaemonsetListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDesired

`func (o *DaemonsetListDto) GetDesired() int32`

GetDesired returns the Desired field if non-nil, zero value otherwise.

### GetDesiredOk

`func (o *DaemonsetListDto) GetDesiredOk() (*int32, bool)`

GetDesiredOk returns a tuple with the Desired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesired

`func (o *DaemonsetListDto) SetDesired(v int32)`

SetDesired sets Desired field to given value.


### GetCurrent

`func (o *DaemonsetListDto) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *DaemonsetListDto) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *DaemonsetListDto) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetReady

`func (o *DaemonsetListDto) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *DaemonsetListDto) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *DaemonsetListDto) SetReady(v int32)`

SetReady sets Ready field to given value.


### GetAvailable

`func (o *DaemonsetListDto) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DaemonsetListDto) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DaemonsetListDto) SetAvailable(v string)`

SetAvailable sets Available field to given value.


### SetAvailableNil

`func (o *DaemonsetListDto) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *DaemonsetListDto) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCreatedAt

`func (o *DaemonsetListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaemonsetListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaemonsetListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DaemonsetListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DaemonsetListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetNamespace

`func (o *DaemonsetListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DaemonsetListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DaemonsetListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *DaemonsetListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DaemonsetListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetImage

`func (o *DaemonsetListDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DaemonsetListDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DaemonsetListDto) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *DaemonsetListDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *DaemonsetListDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


