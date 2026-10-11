# StatefulSetListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Ready** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Images** | **[]string** |  | 

## Methods

### NewStatefulSetListDto

`func NewStatefulSetListDto(state NullableString, name NullableString, ready NullableString, createdAt NullableString, namespace NullableString, images []string, ) *StatefulSetListDto`

NewStatefulSetListDto instantiates a new StatefulSetListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatefulSetListDtoWithDefaults

`func NewStatefulSetListDtoWithDefaults() *StatefulSetListDto`

NewStatefulSetListDtoWithDefaults instantiates a new StatefulSetListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StatefulSetListDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatefulSetListDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatefulSetListDto) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *StatefulSetListDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *StatefulSetListDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetName

`func (o *StatefulSetListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatefulSetListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatefulSetListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StatefulSetListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StatefulSetListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReady

`func (o *StatefulSetListDto) GetReady() string`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *StatefulSetListDto) GetReadyOk() (*string, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *StatefulSetListDto) SetReady(v string)`

SetReady sets Ready field to given value.


### SetReadyNil

`func (o *StatefulSetListDto) SetReadyNil(b bool)`

 SetReadyNil sets the value for Ready to be an explicit nil

### UnsetReady
`func (o *StatefulSetListDto) UnsetReady()`

UnsetReady ensures that no value is present for Ready, not even an explicit nil
### GetCreatedAt

`func (o *StatefulSetListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatefulSetListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatefulSetListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StatefulSetListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StatefulSetListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetNamespace

`func (o *StatefulSetListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StatefulSetListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StatefulSetListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *StatefulSetListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *StatefulSetListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetImages

`func (o *StatefulSetListDto) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *StatefulSetListDto) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *StatefulSetListDto) SetImages(v []string)`

SetImages sets Images field to given value.


### SetImagesNil

`func (o *StatefulSetListDto) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *StatefulSetListDto) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


