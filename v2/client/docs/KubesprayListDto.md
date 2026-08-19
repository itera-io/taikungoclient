# KubesprayListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Version** | **NullableString** |  | 
**KubernetesVersion** | **NullableString** |  | 
**IsDeprecated** | **bool** |  | 

## Methods

### NewKubesprayListDto

`func NewKubesprayListDto(id int32, version NullableString, kubernetesVersion NullableString, isDeprecated bool, ) *KubesprayListDto`

NewKubesprayListDto instantiates a new KubesprayListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubesprayListDtoWithDefaults

`func NewKubesprayListDtoWithDefaults() *KubesprayListDto`

NewKubesprayListDtoWithDefaults instantiates a new KubesprayListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubesprayListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubesprayListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubesprayListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *KubesprayListDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubesprayListDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubesprayListDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *KubesprayListDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *KubesprayListDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetKubernetesVersion

`func (o *KubesprayListDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubesprayListDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubesprayListDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *KubesprayListDto) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *KubesprayListDto) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetIsDeprecated

`func (o *KubesprayListDto) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *KubesprayListDto) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *KubesprayListDto) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


