# KubernetesActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKubernetesActionRequest

`func NewKubernetesActionRequest(name string, ) *KubernetesActionRequest`

NewKubernetesActionRequest instantiates a new KubernetesActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesActionRequestWithDefaults

`func NewKubernetesActionRequestWithDefaults() *KubernetesActionRequest`

NewKubernetesActionRequestWithDefaults instantiates a new KubernetesActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesActionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesActionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesActionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *KubernetesActionRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesActionRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesActionRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesActionRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *KubernetesActionRequest) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *KubernetesActionRequest) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


