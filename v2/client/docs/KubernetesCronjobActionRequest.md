# KubernetesCronjobActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 

## Methods

### NewKubernetesCronjobActionRequest

`func NewKubernetesCronjobActionRequest(name NullableString, namespace NullableString, ) *KubernetesCronjobActionRequest`

NewKubernetesCronjobActionRequest instantiates a new KubernetesCronjobActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCronjobActionRequestWithDefaults

`func NewKubernetesCronjobActionRequestWithDefaults() *KubernetesCronjobActionRequest`

NewKubernetesCronjobActionRequestWithDefaults instantiates a new KubernetesCronjobActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesCronjobActionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCronjobActionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCronjobActionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *KubernetesCronjobActionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesCronjobActionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *KubernetesCronjobActionRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesCronjobActionRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesCronjobActionRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *KubernetesCronjobActionRequest) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *KubernetesCronjobActionRequest) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


