# DeploymentListDto

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

### NewDeploymentListDto

`func NewDeploymentListDto(state NullableString, name NullableString, ready NullableString, createdAt NullableString, namespace NullableString, images []string, ) *DeploymentListDto`

NewDeploymentListDto instantiates a new DeploymentListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentListDtoWithDefaults

`func NewDeploymentListDtoWithDefaults() *DeploymentListDto`

NewDeploymentListDtoWithDefaults instantiates a new DeploymentListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DeploymentListDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentListDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentListDto) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *DeploymentListDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *DeploymentListDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetName

`func (o *DeploymentListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeploymentListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReady

`func (o *DeploymentListDto) GetReady() string`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *DeploymentListDto) GetReadyOk() (*string, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *DeploymentListDto) SetReady(v string)`

SetReady sets Ready field to given value.


### SetReadyNil

`func (o *DeploymentListDto) SetReadyNil(b bool)`

 SetReadyNil sets the value for Ready to be an explicit nil

### UnsetReady
`func (o *DeploymentListDto) UnsetReady()`

UnsetReady ensures that no value is present for Ready, not even an explicit nil
### GetCreatedAt

`func (o *DeploymentListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DeploymentListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DeploymentListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetNamespace

`func (o *DeploymentListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeploymentListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeploymentListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *DeploymentListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *DeploymentListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetImages

`func (o *DeploymentListDto) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DeploymentListDto) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DeploymentListDto) SetImages(v []string)`

SetImages sets Images field to given value.


### SetImagesNil

`func (o *DeploymentListDto) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *DeploymentListDto) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


