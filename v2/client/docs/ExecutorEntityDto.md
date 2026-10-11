# ExecutorEntityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Health** | [**ExecutorHealth**](ExecutorHealth.md) |  | 
**KubeConfig** | **NullableString** |  | 
**IsLocal** | **bool** |  | 

## Methods

### NewExecutorEntityDto

`func NewExecutorEntityDto(id int32, name NullableString, health ExecutorHealth, kubeConfig NullableString, isLocal bool, ) *ExecutorEntityDto`

NewExecutorEntityDto instantiates a new ExecutorEntityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutorEntityDtoWithDefaults

`func NewExecutorEntityDtoWithDefaults() *ExecutorEntityDto`

NewExecutorEntityDtoWithDefaults instantiates a new ExecutorEntityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutorEntityDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutorEntityDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutorEntityDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ExecutorEntityDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutorEntityDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutorEntityDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ExecutorEntityDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExecutorEntityDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHealth

`func (o *ExecutorEntityDto) GetHealth() ExecutorHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ExecutorEntityDto) GetHealthOk() (*ExecutorHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ExecutorEntityDto) SetHealth(v ExecutorHealth)`

SetHealth sets Health field to given value.


### GetKubeConfig

`func (o *ExecutorEntityDto) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *ExecutorEntityDto) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *ExecutorEntityDto) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.


### SetKubeConfigNil

`func (o *ExecutorEntityDto) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *ExecutorEntityDto) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetIsLocal

`func (o *ExecutorEntityDto) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *ExecutorEntityDto) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *ExecutorEntityDto) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


