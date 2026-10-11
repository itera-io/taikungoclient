# NetworkPolicyListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**PodSelector** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewNetworkPolicyListDto

`func NewNetworkPolicyListDto(name NullableString, namespace NullableString, podSelector NullableString, createdAt NullableString, ) *NetworkPolicyListDto`

NewNetworkPolicyListDto instantiates a new NetworkPolicyListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyListDtoWithDefaults

`func NewNetworkPolicyListDtoWithDefaults() *NetworkPolicyListDto`

NewNetworkPolicyListDtoWithDefaults instantiates a new NetworkPolicyListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPolicyListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPolicyListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPolicyListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *NetworkPolicyListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NetworkPolicyListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *NetworkPolicyListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NetworkPolicyListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NetworkPolicyListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *NetworkPolicyListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *NetworkPolicyListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetPodSelector

`func (o *NetworkPolicyListDto) GetPodSelector() string`

GetPodSelector returns the PodSelector field if non-nil, zero value otherwise.

### GetPodSelectorOk

`func (o *NetworkPolicyListDto) GetPodSelectorOk() (*string, bool)`

GetPodSelectorOk returns a tuple with the PodSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSelector

`func (o *NetworkPolicyListDto) SetPodSelector(v string)`

SetPodSelector sets PodSelector field to given value.


### SetPodSelectorNil

`func (o *NetworkPolicyListDto) SetPodSelectorNil(b bool)`

 SetPodSelectorNil sets the value for PodSelector to be an explicit nil

### UnsetPodSelector
`func (o *NetworkPolicyListDto) UnsetPodSelector()`

UnsetPodSelector ensures that no value is present for PodSelector, not even an explicit nil
### GetCreatedAt

`func (o *NetworkPolicyListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkPolicyListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkPolicyListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *NetworkPolicyListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NetworkPolicyListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


